package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/ast"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/t2"
	"io"
	"log"
	"os"
	"strings"
)

func fixId(id string) string {
	if '0' <= id[0] && id[0] <= '9' {
		return "_" + id
	}
	if id == "name" {
		return "_name"
	}
	return id
}

func printOutputNode(out io.Writer, node *t2.IterationNode,
	workflow *t2.Workflow, processor string, types map[string]ast.TypeDef,
	level int) error {
	switch len(node.Children) {
	case 0:
		if node.XMLName.Local != "port" {
			return errors.New("Invalid terminal node: " + node.XMLName.Local +
				" at processor " + processor + ":" + node.Name)
		}

		typeDef, found := types[node.Name]
		if !found {
			return errors.New(
				"Output port is not found: " + processor + ":" + node.Name)
		}

		typeStr := typeDef.String()
		fmt.Fprintf(out, "\n%s%s %s", strings.Repeat("\t", level), 
			fixId(node.Name), typeStr)
		if typeStr != "" {
			fmt.Fprint(out, " ")
		}

		var srcProc, srcPort string
		for _, link := range workflow.Dataflow.Datalinks {
			if link.Sink.Type == "processor" &&
				link.Sink.Processor == processor &&
				link.Sink.Port == node.Name {

				if srcProc != "" || srcPort != "" {
					return errors.New(
						"Duplicate input for " + processor + ":" + node.Name)
				}

				srcPort = link.Source.Port
				switch link.Source.Type {
				case "processor":
					srcProc = link.Source.Processor
				case "dataflow":
					srcProc = "dataflow"
				default:
					return errors.New(
						"Unknown datalink type: " + link.Source.Type)
				}

			}
		}

		if srcProc == "" || srcPort == "" {
			return errors.New(
				"Can't find source for port: " + processor + ":" + node.Name)
		}

		fmt.Fprintf(out, "= %s", fixId(srcProc))
		if srcPort != node.Name {
			fmt.Fprintf(out, ":%s", fixId(srcPort))
		}

		return nil
	case 1:
		return printOutputNode(out, &node.Children[0], workflow, processor,
			types, level)
	default:
		prefix := strings.Repeat("\t", level)

		var op string
		switch node.XMLName.Local {
		case "cross":
			op = " *"
		case "dot":
			op = ","
		case "prefix":
			op = " <"
			if len(node.Children) != 2 {
				return errors.New("Invalid prefix node on " + processor + ":" + 
					node.Name)
			}
		default:
			return errors.New("Unknown node: " + node.XMLName.Local +
				" at processor " + processor + ":" + node.Name)
		}

		if level > 0 {
			fmt.Fprint(out, " (")
		}

		for i, child := range node.Children {
			if i != 0 {
				fmt.Fprintf(out, "%s", op)
			}
			err := printOutputNode(out, &child, workflow, processor, types, 
				level+1)
			if err != nil {
				return err
			}
		}

		if level > 0 {
			fmt.Fprintf(out, "\n%s)", prefix)
		}
	}
	return nil
}

func writeProcessor(workflow *t2.Workflow, processor *t2.Processor) error {
	out, err := os.Create(fixId(processor.Name) + ".processor")
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Fprintf(out, "name: %s\n", fixId(processor.Name))

	if processor.DispatchStack.GetLayer("ListSanitizer") != nil {
		fmt.Fprint(out, "sanitized: true\n")
	}

	parallel := processor.DispatchStack.GetLayer("Parallelize")
	if parallel == nil {
		return errors.New("Can't find parallelize layer for " + processor.Name)
	}
	parallelConfig := parallel.ConfigBean.ParallelizeConfig
	if parallelConfig == nil {
		return errors.New("Invalid parallelize config for " + processor.Name)
	}
	fmt.Fprintf(out, "parallel: %d\n", parallelConfig.MaxJobs)

	waitList := make([]string, 0)
	for _, cond := range workflow.Dataflow.Conditions {
		if cond.Target == processor.Name {
			waitList = append(waitList, fixId(cond.Control))
		}
	}

	var script, ext string
	inputMap := make(map[string]ast.TypeDef)
	outputMap := make(map[string]ast.TypeDef)

	switch processor.Activity.Class {
	case "ru.nsc.ict.taverna.activity.ConfigParserActivity":
		fmt.Fprint(out, "type: config\n")
		fmt.Fprint(out, "inputs: config_data = dataflow\n")
		fmt.Fprint(out, "outputs: config\n")
		return nil
	case "ru.nsc.ict.taverna.activity.WaitForCopyActivity":
		fmt.Fprint(out, "type: waitForCopy\n")
		fmt.Fprintf(out, "wait: %s\n", strings.Join(waitList, ", "))
		return nil
	case "ru.nsc.ict.taverna.activity.BashFileActivity":
		fmt.Fprintf(out, "type: bash\n")
		ext = "sh"
		config := processor.Activity.ConfigBean.BashActivityConfig
		if config == nil {
			return errors.New("Invalid bash activity config for " + processor.Name)
		}

		script = config.Script

		if config.Scheduler != "" {
			fmt.Fprintf(out, "scheduler: %s\n", config.Scheduler)
		}

		if config.Interpreter != "" {
			fmt.Fprintf(out, "interpreter: %s\n", config.Interpreter)
		}

		for _, port := range config.InputPorts {
			portType := ""
			if port.Type != "SIMPLE" {
				portType = strings.ToLower(port.Type)
			}
			inputMap[port.Name] = ast.TypeDef{portType, port.Depth}
		}

		for _, port := range config.OutputPorts {
			portType := ""
			if port.File {
				portType = "file"
			}

			portDepth := 0
			if port.List {
				portDepth = 1
			}

			outputMap[port.Name] = ast.TypeDef{portType, portDepth}
		}
	case "net.sf.taverna.t2.activities.beanshell.BeanshellActivity":
		fmt.Fprintf(out, "type: beanshell\n")
		ext = "bsh"
		config := processor.Activity.ConfigBean.BeanshellActivityConfig
		if config == nil {
			return errors.New("Invalid beanshell activity config for " + processor.Name)
		}

		script = config.Script

		for _, port := range config.InputPorts {
			inputMap[port.Name] = ast.TypeDef{"", port.Depth}
		}

		for _, port := range config.OutputPorts {
			outputMap[port.Name] = ast.TypeDef{"", port.Depth}
		}
	default:
		return errors.New("Unknown processor type " + processor.Activity.Class)
	}

	if script != "" {
		if ext != "" {
			scriptFile, err := os.Create(fixId(processor.Name) + "." + ext)
			if err != nil {
				return err
			}
			defer scriptFile.Close()

			fmt.Fprint(scriptFile, script)
			fmt.Fprintf(out, "script: %s.%s\n", fixId(processor.Name), ext)
		} else {
			fmt.Fprintf(out, "script: \"%s\"\n", script)
		}
	}

	if len(inputMap) > 0 {
		fmt.Fprint(out, "inputs:")
		err := printOutputNode(out, processor.IterationStrategy.Root, workflow,
			processor.Name, inputMap, 0)
		if err != nil {
			return err
		}
		fmt.Fprint(out, "\n")
	}

	if len(outputMap) > 0 {
		fmt.Fprint(out, "outputs:\n")
		next := false
		for name, desc := range outputMap {
			if next {
				fmt.Fprint(out, ",\n")
			} else {
				next = true
			}

			fmt.Fprintf(out, "\t%s %v", fixId(name), desc)
		}
		fmt.Fprint(out, "\n")
	}

	if len(waitList) > 0 {
		fmt.Fprintf(out, "wait: %s\n", strings.Join(waitList, ", "))
	}
	return nil
}

func writeDSL(workflow *t2.Workflow, dir string) error {
	oldDir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	defer os.Chdir(oldDir)

	for _, processor := range workflow.Dataflow.Processors {
		err := writeProcessor(workflow, &processor)
		if err != nil {
			return err
		}
	}

	return nil
}

func run() error {
	if len(os.Args) != 2 {
		return errors.New(fmt.Sprintf(
			"Invalid arguments. Usage %s <output_dir>", os.Args[0]))
	}
	outDir := os.Args[1]

	decoder := xml.NewDecoder(os.Stdin)
	var workflow t2.Workflow
	err := decoder.Decode(&workflow)
	if err != nil {
		return errors.New("Error while decoding xml:" + err.Error())
	}

	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	return writeDSL(&workflow, outDir)
}

func main() {
	err := run()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
