package ast

import (
	"errors"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/token"
	"strconv"
)

type IterType int

const (
	Cross IterType = iota
	Dot
	Prefix
	Port
)

type Processor struct {
	Name, Type, Interpreter, Scheduler, Script string
	Sanitized                                  bool
	Parallel                                   int
	WaitList                                   []string
	Outputs                                    []PortDef
	Inputs                                     IterNode
}

type TypeDef struct {
	Type  string
	Depth int
}

type PortDef struct {
	Name string
	TypeDef
}

type SourceDef struct {
	Processor, Port string
}

type IterNode struct {
	Type     IterType
	Children []IterNode
	Port     PortDef
	Source   SourceDef
}

func (p TypeDef) String() string {
	switch {
	case p.Type == "" && p.Depth == 0:
		return ""
	case p.Type == "":
		return "(" + strconv.Itoa(p.Depth) + ")"
	case p.Depth == 0:
		return p.Type
	default:
		return p.Type + "(" + strconv.Itoa(p.Depth) + ")"
	}
}

func StrVal(val interface{}) (string, error) {
	return string(val.(*token.Token).Lit), nil
}

func StrLitVal(val interface{}) (string, error) {
	str := val.(*token.Token).Lit
	return string(str[1 : len(str)-1]), nil
}

func IntVal(val interface{}) (int, error) {
	res, err := strconv.ParseInt(string(val.(*token.Token).Lit), 10, 0)
	return int(res), err
}

func NewTypeDef(name, depth interface{}) (TypeDef, error) {
	return TypeDef{name.(string), depth.(int)}, nil
}

func NewPortDef(name, typeDef interface{}) (PortDef, error) {
	var res PortDef
	res.Name = name.(string)
	if typeDef != nil {
		res.TypeDef = typeDef.(TypeDef)
	}
	return res, nil
}

func NewSourceDef(processor, port interface{}) (SourceDef, error) {
	return SourceDef{processor.(string), port.(string)}, nil
}

func NewIterPort(port, source interface{}) (*IterNode, error) {
	return &IterNode{
		Type:   Port,
		Port:   port.(PortDef),
		Source: source.(SourceDef)}, nil
}

func NewIterNode(nodeType IterType, args ...interface{}) (*IterNode, error) {
	if nodeType == Port {
		return nil, errors.New("Can't create PortNode, use NewIterPort")
	}

	res := &IterNode{
		Type:     nodeType,
		Children: make([]IterNode, 0)}

	for _, item := range args {
		res.Children = append(res.Children, *item.(*IterNode))
	}
	return res, nil
}

func (n *IterNode) Append(val interface{}) (*IterNode, error) {
	if n.Type == Port {
		return nil, errors.New("Can't append children to PortNode")
	}
	n.Children = append(n.Children, *val.(*IterNode))
	return n, nil
}

func NewProcessor() *Processor {
	return &Processor{
		Parallel: 1}
}

func SetName(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Name = val.(string)
	return p, nil
}

func SetType(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Type = val.(string)
	return p, nil
}

func SetSanitize(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Sanitized = val.(bool)
	return p, nil
}

func SetParallel(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Parallel = val.(int)
	return p, nil
}

func SetInterpreter(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Interpreter = val.(string)
	return p, nil
}

func SetScheduler(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Scheduler = val.(string)
	return p, nil
}

func SetScript(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Script = val.(string)
	return p, nil
}

func SetOutputs(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Outputs = val.([]PortDef)
	return p, nil
}

func SetInputs(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.Inputs = *val.(*IterNode)
	return p, nil
}

func SetWaitList(proc, val interface{}) (*Processor, error) {
	var p *Processor
	if proc != nil {
		p = proc.(*Processor)
	} else {
		p = NewProcessor()
	}
	p.WaitList = val.([]string)
	return p, nil
}
