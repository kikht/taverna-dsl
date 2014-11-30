package main

import (
	"encoding/xml"
	//	"github.com/gobs/pretty"
	"log"
	"os"
)

type t2Port struct {
	Name          string `xml:"name"`
	Depth         int    `xml:"depth"`
	GranularDepth int    `xml:"granularDepth"`
	//	Annotations   string `xml:"annotations,omitempty"`
}

type t2ClassRef struct {
	Group    string `xml:"raven>group"`
	Artifact string `xml:"raven>artifact"`
	Version  string `xml:"raven>version"`
	Class    string `xml:"class"`
}

type t2Map struct {
	From string `xml:"from,attr"`
	To   string `xml:"to,attr"`
}

type t2ConfigBean struct {
	Encoding                string                     `xml:"encoding,attr"`
	NullConfig              *t2EmptyActivityConfig     `xml:"null"`
	BashActivityConfig      *t2BashActivityConfig      `xml:"ru.nsc.ict.taverna.activity.BashFileActivityConfig"`
	ParallelizeConfig       *t2ParallelizeConfig       `xml:"net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.ParallelizeConfig"`
	RetryConfig             *t2RetryConfig             `xml:"net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.RetryConfig"`
	EmptyActivityConfig     *t2EmptyActivityConfig     `xml:"ru.nsc.ict.taverna.activity.EmptyActivityConfig"`
	BeanshellActivityConfig *t2BeanshellActivityConfig `xml:"net.sf.taverna.t2.activities.beanshell.BeanshellActivityConfigurationBean"`
	//TODO: other configs
}

type t2Activity struct {
	t2ClassRef
	InputMap   []t2Map      `xml:"inputMap>map"`
	OutputMap  []t2Map      `xml:"outputMap>map"`
	ConfigBean t2ConfigBean `xml:"configBean"`
}

type t2DispatchLayer struct {
	t2ClassRef
	ConfigBean t2ConfigBean `xml:"configBean"`
}

type t2IterationNode struct {
	XMLName  xml.Name
	Children []t2IterationNode `xml:",any"`
	Name     string            `xml:"name,attr"`
	Depth    int               `xml:"depth,attr"`
}

type t2IterationStrategy struct {
	Root *t2IterationNode `xml:",any"`
}

type t2Processor struct {
	Name              string              `xml:"name"`
	InputPorts        []t2Port            `xml:"inputPorts>port"`
	OutputPorts       []t2Port            `xml:"outputPorts>port"`
	Activity          t2Activity          `xml:"activities>activity"`
	DispatchStack     []t2DispatchLayer   `xml:"dispatchStack>dispatchLayer"`
	IterationStrategy t2IterationStrategy `xml:"iterationStrategyStack>iteration>strategy"`
}

type t2Condition struct {
	Control string `xml:"control,attr"`
	Target  string `xml:"target,attr"`
}

type t2DatalinkEdge struct {
	Type      string `xml:"type,attr"`
	Processor string `xml:"processor"`
	Port      string `xml:"port"`
}

type t2Datalink struct {
	Source t2DatalinkEdge `xml:"source"`
	Sink   t2DatalinkEdge `xml:"sink"`
}

type t2Dataflow struct {
	Id          string        `xml:"id,attr"`
	Role        string        `xml:"role,attr"`
	Name        string        `xml:"name"`
	InputPorts  []t2Port      `xml:"inputPorts>port"`
	OutputPorts []t2Port      `xml:"outputPorts>port"`
	Processors  []t2Processor `xml:"processors>processor"`
	Conditions  []t2Condition `xml:"conditions>condition"`
	Datalinks   []t2Datalink  `xml:"datalinks>datalink"`
}

type t2Workflow struct {
	XMLName    xml.Name   `xml:"workflow"`
	XMLNs      string     `xml:"xmlns,attr"`
	Version    int        `xml:"version,attr"`
	ProducedBy string     `xml:"producedBy,attr"`
	Dataflow   t2Dataflow `xml:"dataflow"`
}

type t2EmptyActivityConfig struct {
	XMLNs string `xml:"xmlns,attr"`
}

type t2ParallelizeConfig struct {
	XMLNs   string `xml:"xmlns,attr"`
	MaxJobs int    `xml:"maxJobs"`
}

type t2RetryConfig struct {
	XMLNs         string  `xml:"xmlns,attr"`
	BackoffFactor float64 `xml:"backoffFactor"`
	InitialDelay  int     `xml:"initialDelay"`
	MaxDelay      int     `xml:"maxDelay"`
	MaxRetries    int     `xml:"maxRetries"`
}

type t2BashInputPort struct {
	Name  string `xml:"name"`
	Depth int    `xml:"depth"`
	Type  string `xml:"type"`
}

type t2BashOutputPort struct {
	Name string `xml:"name"`
	File bool   `xml:"file"`
	List bool   `xml:"list"`
}

type t2BashActivityConfig struct {
	XMLNs       string             `xml:"xmlns,attr"`
	Script      string             `xml:"script"`
	Interpreter string             `xml:"interpreter"`
	Scheduler   string             `xml:"schedulerName"`
	InputPorts  []t2BashInputPort  `xml:"inputPorts>ru.nsc.ict.taverna.activity.BashFileActivityInputPort"`
	OutputPorts []t2BashOutputPort `xml:"outputPorts>ru.nsc.ict.taverna.activity.BashFileActivityOutputPort"`
}

type t2ActivityInputPort struct {
	Name                    string   `xml:"name"`
	Depth                   int      `xml:"depth"`
	MimeTypes               []string `xml:"mimeTypes>string"`
	HandledReferenceSchemes string   `xml:"handledReferenceSchemes,omitempty"`
	TranslatedElementType   string   `xml:"translatedElementType,omitempty"`
	AllowsLiteralValues     bool     `xml:"allowsLiteralValues"`
}

type t2ActivityOutputPort struct {
	Name          string   `xml:"name"`
	Depth         int      `xml:"depth"`
	MimeTypes     []string `xml:"mimeTypes>string"`
	GranularDepth int      `xml:"granularDepth"`
}

type t2BeanshellActivityConfig struct {
	//TODO:inputs, outputs
	Inputs               []t2ActivityInputPort  `xml:"inputs>net.sf.taverna.t2.workflowmodel.processor.activity.config.ActivityInputPortDefinitionBean"`
	Outputs              []t2ActivityOutputPort `xml:"outputs>net.sf.taverna.t2.workflowmodel.processor.activity.config.ActivityOutputPortDefinitionBean"`
	ClassLoaderSharing   string                 `xml:"classLoaderSharing"`
	LocalDependencies    string                 `xml:"localDependencies"`
	ArtifactDependencies string                 `xml:"artifactDependencies"`
	Script               string                 `xml:"script"`
	Dependencies         string                 `xml:"dependencies"`
}

func main() {
	decoder := xml.NewDecoder(os.Stdin)
	var workflow t2Workflow
	err := decoder.Decode(&workflow)
	if err == nil {
		//pretty.PrettyPrint(workflow)
		encoder := xml.NewEncoder(os.Stdout)
		encoder.Indent("", "  ")
		err := encoder.Encode(workflow)
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Print(err)
	}
}
