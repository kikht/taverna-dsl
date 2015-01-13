package main

import (
	"errors"
	"fmt"
	"github.com/twinj/uuid"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/ast"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/lexer"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/parser"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/t2"
	"io/ioutil"
	"log"
	"os"
	"encoding/xml"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func newDispatchStack(parallel int, sanitized bool) t2.DispatchStack {
	if parallel < 1 {
		parallel = 1
	}
	res := []t2.DispatchLayer{
		t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.core",
				Artifact: "workflowmodel-impl",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.Parallelize"},
			ConfigBean: t2.ConfigBean{
				Encoding:          "xstream",
				ParallelizeConfig: &t2.ParallelizeConfig{MaxJobs: parallel}}},
		t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.core",
				Artifact: "workflowmodel-impl",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.ErrorBounce"},
			ConfigBean: t2.ConfigBean{
				Encoding:   "xstream",
				NullConfig: &t2.EmptyActivityConfig{}}},
		t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.core",
				Artifact: "workflowmodel-impl",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.Failover"},
			ConfigBean: t2.ConfigBean{
				Encoding:   "xstream",
				NullConfig: &t2.EmptyActivityConfig{}}},
		t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.core",
				Artifact: "workflowmodel-impl",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.Retry"},
			ConfigBean: t2.ConfigBean{
				Encoding: "xstream",
				RetryConfig: &t2.RetryConfig{
					BackoffFactor: 1,
					InitialDelay:  1000,
					MaxDelay:      5000,
					MaxRetries:    0}}},
		t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.core",
				Artifact: "workflowmodel-impl",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.Invoke"},
			ConfigBean: t2.ConfigBean{
				Encoding:   "xstream",
				NullConfig: &t2.EmptyActivityConfig{}}}}
	if sanitized {
		sanitizer := t2.DispatchLayer{
			ClassRef: t2.ClassRef{
				Group:    "ru.nsc.ict.taverna",
				Artifact: "modis-activity-core",
				Version:  "0.0.1-SNAPSHOT",
				Class:    "ru.nsc.ict.taverna.dispatch.ListSanitizer"},
			ConfigBean: t2.ConfigBean{
				Encoding:   "xstream",
				NullConfig: &t2.EmptyActivityConfig{}}}
		res = append(res, t2.DispatchLayer{})
		copy(res[2:], res[1:])
		res[1] = sanitizer
	}
	return res
}

func extractInputs(root ast.IterNode) []ast.IterNode {
	if root.Type == ast.Port {
		return []ast.IterNode{root}
	} else {
		res := []ast.IterNode{}
		for _, child := range root.Children {
			res = append(res, extractInputs(child)...)
		}
		return res
	}
}

func convertIterationStrategy(node ast.IterNode) (t2.IterationNode, error) {
	if node.Type == ast.Port {
		return t2.IterationNode{
			XMLName: xml.Name{"http://taverna.sf.net/2008/xml/t2flow", "port"},
			Name: node.Port.Name,
			Depth: node.Port.Depth}, nil
	} else {
		var res t2.IterationNode
		switch node.Type {
		case ast.Cross:
			res.XMLName.Local = "cross"
		case ast.Dot:
			res.XMLName.Local = "dot"
		case ast.Prefix:
			res.XMLName.Local = "prefix"
		default:
			return res, errors.New("Invalid iteration node type")
		}
		
		res.XMLName.Space = "http://taverna.sf.net/2008/xml/t2flow"
		res.Children = make([]t2.IterationNode, len(node.Children))
		for i, child := range node.Children {
			var err error
			res.Children[i], err = convertIterationStrategy(child)
			if err != nil {
				return res, err
			}
		}

		return res, nil
	}
}

func parseProcessor(p *parser.Parser, filename string, workflow *t2.Workflow, 
	workflowInputs map[string]int) error {
	lex, err := lexer.NewLexerFile(filename)
	if err != nil {
		return err
	}

	res, err := p.Parse(lex)
	if err != nil {
		return err
	}

	astProc := res.(*ast.Processor)
	inputs := extractInputs(astProc.Inputs)
	var activity t2.Activity

	switch astProc.Type {
	case "beanshell":
		config := t2.BeanshellActivityConfig{
			InputPorts:         make([]t2.ActivityInputPort, len(inputs)),
			OutputPorts:        make([]t2.ActivityOutputPort, len(astProc.Outputs)),
			ClassLoaderSharing: "workflow"}

		script, err := ioutil.ReadFile(astProc.Script)
		if err == nil {
			config.Script = string(script)
		} else {
			log.Printf("Can't read script file for %s, assuming inline", 
				filename)
			config.Script = astProc.Script
		}

		for i, port := range inputs {
			config.InputPorts[i] = t2.ActivityInputPort{
				Name:                  port.Port.Name,
				Depth:                 port.Port.Depth,
				MimeTypes:             []string{"text/plain"},
				TranslatedElementType: "java.lang.String",
				AllowsLiteralValues:   true}
		}

		for i, port := range astProc.Outputs {
			config.OutputPorts[i] = t2.ActivityOutputPort{
				Name:          port.Name,
				Depth:         port.Depth,
				GranularDepth: port.Depth}
		}

		activity = t2.Activity{
			ClassRef: t2.ClassRef{
				Group:    "net.sf.taverna.t2.activities",
				Artifact: "beanshell-activity",
				Version:  "1.3",
				Class:    "net.sf.taverna.t2.activities.beanshell.BeanshellActivity"},
			ConfigBean: t2.ConfigBean{
				Encoding:                "xstream",
				BeanshellActivityConfig: &config}}
	case "bash":
		config := t2.BashActivityConfig{
			Interpreter: astProc.Interpreter,
			Scheduler:   astProc.Scheduler,
			InputPorts:  make([]t2.BashInputPort, len(inputs)),
			OutputPorts: make([]t2.BashOutputPort, len(astProc.Outputs))}
		
		script, err := ioutil.ReadFile(astProc.Script)
		if err == nil {
			config.Script = string(script)
		} else {
			log.Printf("Can't read script file for %s, assuming inline", 
				filename)
			config.Script = astProc.Script
		}

		for i, port := range inputs {
			typeStr := strings.ToUpper(port.Port.Type)
			if typeStr == "" {
				typeStr = "SIMPLE"
			}
			config.InputPorts[i] = t2.BashInputPort{
				Name:  port.Port.Name,
				Depth: port.Port.Depth,
				Type:  typeStr}
		}

		anyFiles := false
		for i, port := range astProc.Outputs {
			file := port.Type == "file"
			list := port.Depth > 0
			anyFiles = anyFiles || file
			config.OutputPorts[i] = t2.BashOutputPort{
				Name: port.Name,
				File: file,
				List: list}

		}

		if anyFiles {
			workflow.Dataflow.Conditions = append(
				workflow.Dataflow.Conditions,
				t2.Condition{astProc.Name, "Wait_for_copy"})
		}

		activity = t2.Activity{
			ClassRef: t2.ClassRef{
				Group:    "ru.nsc.ict.taverna",
				Artifact: "modis-activity-core",
				Version:  "0.0.1-SNAPSHOT",
				Class:    "ru.nsc.ict.taverna.activity.BashFileActivity"},
			ConfigBean: t2.ConfigBean{
				Encoding:           "xstream",
				BashActivityConfig: &config}}
	case "config":
		activity = t2.Activity{
			ClassRef: t2.ClassRef{
				Group:    "ru.nsc.ict.taverna",
				Artifact: "modis-activity-core",
				Version:  "0.0.1-SNAPSHOT",
				Class:    "ru.nsc.ict.taverna.activity.ConfigParserActivity"},
			ConfigBean: t2.ConfigBean{
				Encoding:            "xstream",
				EmptyActivityConfig: &t2.EmptyActivityConfig{}}}
	case "waitForCopy":
		//just skip it
		return nil
	}

	activity.InputMap = make([]t2.Map, len(inputs))
	activity.OutputMap = make([]t2.Map, len(astProc.Outputs))
	processor := t2.Processor{
		Name:          astProc.Name,
		InputPorts:    make([]t2.Port, len(inputs)),
		OutputPorts:   make([]t2.Port, len(astProc.Outputs)),
		Activity:      activity,
		DispatchStack: newDispatchStack(astProc.Parallel, 
			astProc.Sanitized)}

	for i, port := range inputs {
		processor.Activity.InputMap[i] = 
			t2.Map{port.Port.Name, port.Port.Name}
		processor.InputPorts[i] = 
			t2.Port{port.Port.Name, port.Port.Depth, port.Port.Depth}

		link := t2.Datalink{
			Source: t2.DatalinkEdge{
				Type:      "processor",
				Processor: port.Source.Processor,
				Port:      port.Source.Port},
			Sink: t2.DatalinkEdge{
				Type:      "processor",
				Processor: astProc.Name,
				Port:      port.Port.Name}}
		if link.Source.Port == "" {
			link.Source.Port = port.Port.Name
		}
		if link.Source.Processor == "dataflow" {
			link.Source.Processor = ""
			link.Source.Type = "dataflow"

			oldDepth, portExists := workflowInputs[link.Source.Port]
			if !portExists || oldDepth < port.Port.Depth {
				workflowInputs[link.Source.Port] = port.Port.Depth
			}
		}
		workflow.Dataflow.Datalinks = append(
			workflow.Dataflow.Datalinks, link)
	}

	for i, port := range astProc.Outputs {
		processor.Activity.OutputMap[i] = t2.Map{port.Name, port.Name}
		processor.OutputPorts[i] = t2.Port{port.Name, port.Depth, port.Depth}
	}

	for _, name := range astProc.WaitList {
		workflow.Dataflow.Conditions = append(
			workflow.Dataflow.Conditions,
			t2.Condition{name, astProc.Name})
	}

	iter, err := convertIterationStrategy(astProc.Inputs)
	if err != nil {
		return err
	}
	if astProc.Inputs.Type == ast.Port {
		iter = t2.IterationNode{
			XMLName: xml.Name{"http://taverna.sf.net/2008/xml/t2flow", 
				"cross"},
			Children: []t2.IterationNode{iter}}
	}
	processor.IterationStrategy.Root = &iter

	workflow.Dataflow.Processors = append(
		workflow.Dataflow.Processors,
		processor)

	return nil
}

func run() error {
	if len(os.Args) != 2 {
		return errors.New(fmt.Sprintf(
			"Invalid arguments. Usage %s <input_dir>", os.Args[0]))
	}
	inputDir := os.Args[1]

	oldDir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(inputDir)
	if err != nil {
		return err
	}
	defer os.Chdir(oldDir)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	waitForCopy := t2.Processor{
		Name: "Wait_for_copy",
		Activity: t2.Activity{
			ClassRef: t2.ClassRef{
				Group:    "ru.nsc.ict.taverna",
				Artifact: "modis-activity-core",
				Version:  "0.0.1-SNAPSHOT",
				Class:    "ru.nsc.ict.taverna.activity.WaitForCopyActivity"},
			ConfigBean: t2.ConfigBean{
				Encoding:            "xstream",
				EmptyActivityConfig: &t2.EmptyActivityConfig{}}},
		DispatchStack: newDispatchStack(1, false)}

	uuid := uuid.NewV4().String()
	uuid = uuid[1:len(uuid)-1]
	workflow := t2.Workflow{
		XMLNs:      "http://taverna.sf.net/2008/xml/t2flow",
		Version:    1,
		ProducedBy: "t2Compile",
		Dataflow: t2.Dataflow{
			Id:         uuid,
			Role:       "top",
			Name:       inputDir,
			Processors: []t2.Processor{waitForCopy}}}
	workflowInputs := make(map[string]int)

	p := parser.NewParser()
	for _, file := range files {
		if file.Mode().IsRegular() && 
			strings.HasSuffix(file.Name(), ".processor") {
			err := parseProcessor(p, file.Name(), &workflow, workflowInputs)
			if err != nil {
				return errors.New(fmt.Sprintf("Error while processing %s: %s",
					file.Name(), err.Error())) 
			}
		}
	}
	
	workflow.Dataflow.InputPorts = make([]t2.Port, len(workflowInputs))
	i := 0
	for name, depth := range workflowInputs {
		workflow.Dataflow.InputPorts[i] = t2.Port{name, depth, depth}
		i++
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")
	err = encoder.Encode(workflow)
	if err != nil {
		return err
	}

	return nil
}
