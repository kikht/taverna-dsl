package t2

import "encoding/xml"
import "strings"

type Port struct {
	Name          string `xml:"name"`
	Depth         int    `xml:"depth"`
	GranularDepth int    `xml:"granularDepth"`
	//	Annotations   string `xml:"annotations,omitempty"`
}

type ClassRef struct {
	Group    string `xml:"raven>group"`
	Artifact string `xml:"raven>artifact"`
	Version  string `xml:"raven>version"`
	Class    string `xml:"class"`
}

type Map struct {
	From string `xml:"from,attr"`
	To   string `xml:"to,attr"`
}

type ConfigBean struct {
	Encoding                string                   `xml:"encoding,attr"`
	NullConfig              *EmptyActivityConfig     `xml:"null"`
	BashActivityConfig      *BashActivityConfig      `xml:"ru.nsc.ict.taverna.activity.BashFileActivityConfig"`
	ParallelizeConfig       *ParallelizeConfig       `xml:"net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.ParallelizeConfig"`
	RetryConfig             *RetryConfig             `xml:"net.sf.taverna.t2.workflowmodel.processor.dispatch.layers.RetryConfig"`
	EmptyActivityConfig     *EmptyActivityConfig     `xml:"ru.nsc.ict.taverna.activity.EmptyActivityConfig"`
	BeanshellActivityConfig *BeanshellActivityConfig `xml:"net.sf.taverna.t2.activities.beanshell.BeanshellActivityConfigurationBean"`
	//TODO: other configs
}

type Activity struct {
	ClassRef
	InputMap   []Map      `xml:"inputMap>map"`
	OutputMap  []Map      `xml:"outputMap>map"`
	ConfigBean ConfigBean `xml:"configBean"`
}

type DispatchLayer struct {
	ClassRef
	ConfigBean ConfigBean `xml:"configBean"`
}

type IterationNode struct {
	XMLName  xml.Name
	Children []IterationNode `xml:",any"`
	Name     string          `xml:"name,attr"`
	Depth    int             `xml:"depth,attr"`
}

type IterationStrategy struct {
	Root *IterationNode `xml:",any"`
}

type DispatchStack []DispatchLayer

type Processor struct {
	Name              string            `xml:"name"`
	InputPorts        []Port            `xml:"inputPorts>port"`
	OutputPorts       []Port            `xml:"outputPorts>port"`
	Activity          Activity          `xml:"activities>activity"`
	DispatchStack     DispatchStack     `xml:"dispatchStack>dispatchLayer"`
	IterationStrategy IterationStrategy `xml:"iterationStrategyStack>iteration>strategy"`
}

type Condition struct {
	Control string `xml:"control,attr"`
	Target  string `xml:"target,attr"`
}

type DatalinkEdge struct {
	Type      string `xml:"type,attr"`
	Processor string `xml:"processor"`
	Port      string `xml:"port"`
}

type Datalink struct {
	Source DatalinkEdge `xml:"source"`
	Sink   DatalinkEdge `xml:"sink"`
}

type Dataflow struct {
	Id          string      `xml:"id,attr"`
	Role        string      `xml:"role,attr"`
	Name        string      `xml:"name"`
	InputPorts  []Port      `xml:"inputPorts>port"`
	OutputPorts []Port      `xml:"outputPorts>port"`
	Processors  []Processor `xml:"processors>processor"`
	Conditions  []Condition `xml:"conditions>condition"`
	Datalinks   []Datalink  `xml:"datalinks>datalink"`
}

type Workflow struct {
	XMLName    xml.Name `xml:"workflow"`
	XMLNs      string   `xml:"xmlns,attr"`
	Version    int      `xml:"version,attr"`
	ProducedBy string   `xml:"producedBy,attr"`
	Dataflow   Dataflow `xml:"dataflow"`
}

type EmptyActivityConfig struct {
	XMLNs string `xml:"xmlns,attr"`
}

type ParallelizeConfig struct {
	XMLNs   string `xml:"xmlns,attr"`
	MaxJobs int    `xml:"maxJobs"`
}

type RetryConfig struct {
	XMLNs         string  `xml:"xmlns,attr"`
	BackoffFactor float64 `xml:"backoffFactor"`
	InitialDelay  int     `xml:"initialDelay"`
	MaxDelay      int     `xml:"maxDelay"`
	MaxRetries    int     `xml:"maxRetries"`
}

type BashInputPort struct {
	Name  string `xml:"name"`
	Depth int    `xml:"depth"`
	Type  string `xml:"type"`
}

type BashOutputPort struct {
	Name string `xml:"name"`
	File bool   `xml:"file"`
	List bool   `xml:"list"`
}

type BashActivityConfig struct {
	XMLNs       string           `xml:"xmlns,attr"`
	Script      string           `xml:"script"`
	Interpreter string           `xml:"interpreter"`
	Scheduler   string           `xml:"schedulerName"`
	InputPorts  []BashInputPort  `xml:"inputPorts>ru.nsc.ict.taverna.activity.BashFileActivityInputPort"`
	OutputPorts []BashOutputPort `xml:"outputPorts>ru.nsc.ict.taverna.activity.BashFileActivityOutputPort"`
}

type ActivityInputPort struct {
	Name                    string   `xml:"name"`
	Depth                   int      `xml:"depth"`
	MimeTypes               []string `xml:"mimeTypes>string"`
	HandledReferenceSchemes string   `xml:"handledReferenceSchemes,omitempty"`
	TranslatedElementType   string   `xml:"translatedElementType,omitempty"`
	AllowsLiteralValues     bool     `xml:"allowsLiteralValues"`
}

type ActivityOutputPort struct {
	Name          string   `xml:"name"`
	Depth         int      `xml:"depth"`
	MimeTypes     []string `xml:"mimeTypes>string"`
	GranularDepth int      `xml:"granularDepth"`
}

type BeanshellActivityConfig struct {
	//TODO:inputs, outputs
	InputPorts           []ActivityInputPort  `xml:"inputs>net.sf.taverna.t2.workflowmodel.processor.activity.config.ActivityInputPortDefinitionBean"`
	OutputPorts          []ActivityOutputPort `xml:"outputs>net.sf.taverna.t2.workflowmodel.processor.activity.config.ActivityOutputPortDefinitionBean"`
	ClassLoaderSharing   string               `xml:"classLoaderSharing"`
	LocalDependencies    string               `xml:"localDependencies"`
	ArtifactDependencies string               `xml:"artifactDependencies"`
	Script               string               `xml:"script"`
	Dependencies         string               `xml:"dependencies"`
}

func (p DispatchStack) GetLayer(name string) *DispatchLayer {
	for _, l := range p {
		if strings.Contains(l.Class, name) {
			return &l
		}
	}
	return nil
}
