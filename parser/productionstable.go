
package parser

import ( "gitlab.ict.sbras.ru/taverna/dsl/ast" )

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Processor	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Processor : empty	<<  >>`,
		Id: "Processor",
		NTType: 1,
		Index: 1,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Processor : Processor NameStmt	<< ast.SetName(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetName(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor TypeStmt	<< ast.SetType(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetType(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor SanitInputStmt	<< ast.SetSanitize(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetSanitize(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor ParallelStmt	<< ast.SetParallel(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetParallel(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor InterpreterStmt	<< ast.SetInterpreter(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetInterpreter(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor SchedulerStmt	<< ast.SetScheduler(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetScheduler(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor ScriptStmt	<< ast.SetScript(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetScript(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor OutputStmt	<< ast.SetOutputs(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetOutputs(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor IntputStmt	<< ast.SetInputs(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetInputs(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Processor : Processor WaitStmt	<< ast.SetWaitList(X[0], X[1]) >>`,
		Id: "Processor",
		NTType: 1,
		Index: 11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SetWaitList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `NameStmt : "name" ":" IdVal	<< X[2], nil >>`,
		Id: "NameStmt",
		NTType: 2,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `TypeStmt : "type" ":" IdVal	<< X[2], nil >>`,
		Id: "TypeStmt",
		NTType: 3,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `SanitInputStmt : "sanitized" ":" BoolVal	<< X[2], nil >>`,
		Id: "SanitInputStmt",
		NTType: 4,
		Index: 14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `ParallelStmt : "parallel" ":" IntVal	<< X[2], nil >>`,
		Id: "ParallelStmt",
		NTType: 5,
		Index: 15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `InterpreterStmt : "interpreter" ":" ScriptVal	<< X[2], nil >>`,
		Id: "InterpreterStmt",
		NTType: 6,
		Index: 16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `SchedulerStmt : "scheduler" ":" IdVal	<< X[2], nil >>`,
		Id: "SchedulerStmt",
		NTType: 7,
		Index: 17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `ScriptStmt : "script" ":" ScriptVal	<< X[2], nil >>`,
		Id: "ScriptStmt",
		NTType: 8,
		Index: 18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `OutputStmt : "outputs" ":" OutputList	<< X[2], nil >>`,
		Id: "OutputStmt",
		NTType: 9,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `OutputList : PortDef	<< []ast.PortDef{ X[0].(ast.PortDef) }, nil >>`,
		Id: "OutputList",
		NTType: 10,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.PortDef{ X[0].(ast.PortDef) }, nil
		},
	},
	ProdTabEntry{
		String: `OutputList : OutputList "," PortDef	<< append(X[0].([]ast.PortDef), X[2].(ast.PortDef)), nil >>`,
		Id: "OutputList",
		NTType: 10,
		Index: 21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]ast.PortDef), X[2].(ast.PortDef)), nil
		},
	},
	ProdTabEntry{
		String: `WaitStmt : "wait" ":" WaitList	<< X[2], nil >>`,
		Id: "WaitStmt",
		NTType: 11,
		Index: 22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `WaitList : IdVal	<< []string{ X[0].(string) }, nil >>`,
		Id: "WaitList",
		NTType: 12,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []string{ X[0].(string) }, nil
		},
	},
	ProdTabEntry{
		String: `WaitList : WaitList "," IdVal	<< append(X[0].([]string), X[2].(string)), nil >>`,
		Id: "WaitList",
		NTType: 12,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]string), X[2].(string)), nil
		},
	},
	ProdTabEntry{
		String: `IntputStmt : "inputs" ":" IterOper	<< X[2], nil >>`,
		Id: "IntputStmt",
		NTType: 13,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `IterOper : IterOperCross	<<  >>`,
		Id: "IterOper",
		NTType: 14,
		Index: 26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IterOper : IterOperDot	<<  >>`,
		Id: "IterOper",
		NTType: 14,
		Index: 27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IterOper : IterOperPrefix	<<  >>`,
		Id: "IterOper",
		NTType: 14,
		Index: 28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IterOper : IterPort	<< ast.NewIterNode(ast.Cross, X[0]) >>`,
		Id: "IterOper",
		NTType: 14,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIterNode(ast.Cross, X[0])
		},
	},
	ProdTabEntry{
		String: `IterOperCross : IterNode "*" IterNode	<< ast.NewIterNode(ast.Cross, X[0], X[2]) >>`,
		Id: "IterOperCross",
		NTType: 15,
		Index: 30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIterNode(ast.Cross, X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `IterOperCross : IterOperCross "*" IterNode	<< X[0].(*ast.IterNode).Append(X[2]) >>`,
		Id: "IterOperCross",
		NTType: 15,
		Index: 31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(*ast.IterNode).Append(X[2])
		},
	},
	ProdTabEntry{
		String: `IterOperDot : IterNode "," IterNode	<< ast.NewIterNode(ast.Dot, X[0], X[2]) >>`,
		Id: "IterOperDot",
		NTType: 16,
		Index: 32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIterNode(ast.Dot, X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `IterOperDot : IterOperDot "," IterNode	<< X[0].(*ast.IterNode).Append(X[2]) >>`,
		Id: "IterOperDot",
		NTType: 16,
		Index: 33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(*ast.IterNode).Append(X[2])
		},
	},
	ProdTabEntry{
		String: `IterOperPrefix : IterNode "<" IterNode	<< ast.NewIterNode(ast.Prefix, X[0], X[2]) >>`,
		Id: "IterOperPrefix",
		NTType: 17,
		Index: 34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIterNode(ast.Prefix, X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `IterNode : IterPort	<<  >>`,
		Id: "IterNode",
		NTType: 18,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IterNode : "(" IterOper ")"	<< X[1], nil >>`,
		Id: "IterNode",
		NTType: 18,
		Index: 36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `IterPort : PortDef "=" SourceDef	<< ast.NewIterPort(X[0], X[2]) >>`,
		Id: "IterPort",
		NTType: 19,
		Index: 37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIterPort(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `SourceDef : IdVal	<< ast.NewSourceDef(X[0], "") >>`,
		Id: "SourceDef",
		NTType: 20,
		Index: 38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSourceDef(X[0], "")
		},
	},
	ProdTabEntry{
		String: `SourceDef : IdVal ":" IdVal	<< ast.NewSourceDef(X[0], X[2]) >>`,
		Id: "SourceDef",
		NTType: 20,
		Index: 39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSourceDef(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `PortDef : IdVal	<< ast.NewPortDef(X[0], nil) >>`,
		Id: "PortDef",
		NTType: 21,
		Index: 40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPortDef(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `PortDef : IdVal TypeDef	<< ast.NewPortDef(X[0], X[1]) >>`,
		Id: "PortDef",
		NTType: 21,
		Index: 41,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPortDef(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `TypeDef : IdVal	<< ast.NewTypeDef(X[0], 0) >>`,
		Id: "TypeDef",
		NTType: 22,
		Index: 42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTypeDef(X[0], 0)
		},
	},
	ProdTabEntry{
		String: `TypeDef : IdVal "(" IntVal ")"	<< ast.NewTypeDef(X[0], X[2]) >>`,
		Id: "TypeDef",
		NTType: 22,
		Index: 43,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTypeDef(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `TypeDef : "(" IntVal ")"	<< ast.NewTypeDef("", X[1]) >>`,
		Id: "TypeDef",
		NTType: 22,
		Index: 44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTypeDef("", X[1])
		},
	},
	ProdTabEntry{
		String: `IdVal : id	<< ast.StrVal(X[0]) >>`,
		Id: "IdVal",
		NTType: 23,
		Index: 45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StrVal(X[0])
		},
	},
	ProdTabEntry{
		String: `ScriptVal : str_lit	<< ast.StrLitVal(X[0]) >>`,
		Id: "ScriptVal",
		NTType: 24,
		Index: 46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StrLitVal(X[0])
		},
	},
	ProdTabEntry{
		String: `ScriptVal : id	<< ast.StrVal(X[0]) >>`,
		Id: "ScriptVal",
		NTType: 24,
		Index: 47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StrVal(X[0])
		},
	},
	ProdTabEntry{
		String: `ScriptVal : path	<< ast.StrVal(X[0]) >>`,
		Id: "ScriptVal",
		NTType: 24,
		Index: 48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StrVal(X[0])
		},
	},
	ProdTabEntry{
		String: `BoolVal : "true"	<< true,  nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true,  nil
		},
	},
	ProdTabEntry{
		String: `BoolVal : "false"	<< false, nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `BoolVal : "t"	<< true,  nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true,  nil
		},
	},
	ProdTabEntry{
		String: `BoolVal : "f"	<< false, nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `BoolVal : "0"	<< false, nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return false, nil
		},
	},
	ProdTabEntry{
		String: `BoolVal : "1"	<< true,  nil >>`,
		Id: "BoolVal",
		NTType: 25,
		Index: 54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return true,  nil
		},
	},
	ProdTabEntry{
		String: `IntVal : pos_int	<< ast.IntVal(X[0]) >>`,
		Id: "IntVal",
		NTType: 26,
		Index: 55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.IntVal(X[0])
		},
	},
	ProdTabEntry{
		String: `IntVal : "0"	<< 0, nil >>`,
		Id: "IntVal",
		NTType: 26,
		Index: 56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return 0, nil
		},
	},
	ProdTabEntry{
		String: `IntVal : "1"	<< 1, nil >>`,
		Id: "IntVal",
		NTType: 26,
		Index: 57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return 1, nil
		},
	},
	
}
