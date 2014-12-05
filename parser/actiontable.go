
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(1),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(1),		/* type, reduce: Processor */
			reduce(1),		/* sanitized, reduce: Processor */
			reduce(1),		/* parallel, reduce: Processor */
			reduce(1),		/* interpreter, reduce: Processor */
			reduce(1),		/* scheduler, reduce: Processor */
			reduce(1),		/* script, reduce: Processor */
			reduce(1),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(1),		/* wait, reduce: Processor */
			reduce(1),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* empty */
			shift(12),		/* name */
			nil,		/* : */
			shift(13),		/* type */
			shift(14),		/* sanitized */
			shift(15),		/* parallel */
			shift(16),		/* interpreter */
			shift(17),		/* scheduler */
			shift(18),		/* script */
			shift(19),		/* outputs */
			nil,		/* , */
			shift(20),		/* wait */
			shift(21),		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(2),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(2),		/* type, reduce: Processor */
			reduce(2),		/* sanitized, reduce: Processor */
			reduce(2),		/* parallel, reduce: Processor */
			reduce(2),		/* interpreter, reduce: Processor */
			reduce(2),		/* scheduler, reduce: Processor */
			reduce(2),		/* script, reduce: Processor */
			reduce(2),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(2),		/* wait, reduce: Processor */
			reduce(2),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(3),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(3),		/* type, reduce: Processor */
			reduce(3),		/* sanitized, reduce: Processor */
			reduce(3),		/* parallel, reduce: Processor */
			reduce(3),		/* interpreter, reduce: Processor */
			reduce(3),		/* scheduler, reduce: Processor */
			reduce(3),		/* script, reduce: Processor */
			reduce(3),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(3),		/* wait, reduce: Processor */
			reduce(3),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(4),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(4),		/* type, reduce: Processor */
			reduce(4),		/* sanitized, reduce: Processor */
			reduce(4),		/* parallel, reduce: Processor */
			reduce(4),		/* interpreter, reduce: Processor */
			reduce(4),		/* scheduler, reduce: Processor */
			reduce(4),		/* script, reduce: Processor */
			reduce(4),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(4),		/* wait, reduce: Processor */
			reduce(4),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(5),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(5),		/* type, reduce: Processor */
			reduce(5),		/* sanitized, reduce: Processor */
			reduce(5),		/* parallel, reduce: Processor */
			reduce(5),		/* interpreter, reduce: Processor */
			reduce(5),		/* scheduler, reduce: Processor */
			reduce(5),		/* script, reduce: Processor */
			reduce(5),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(5),		/* wait, reduce: Processor */
			reduce(5),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(6),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(6),		/* type, reduce: Processor */
			reduce(6),		/* sanitized, reduce: Processor */
			reduce(6),		/* parallel, reduce: Processor */
			reduce(6),		/* interpreter, reduce: Processor */
			reduce(6),		/* scheduler, reduce: Processor */
			reduce(6),		/* script, reduce: Processor */
			reduce(6),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(6),		/* wait, reduce: Processor */
			reduce(6),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(7),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(7),		/* type, reduce: Processor */
			reduce(7),		/* sanitized, reduce: Processor */
			reduce(7),		/* parallel, reduce: Processor */
			reduce(7),		/* interpreter, reduce: Processor */
			reduce(7),		/* scheduler, reduce: Processor */
			reduce(7),		/* script, reduce: Processor */
			reduce(7),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(7),		/* wait, reduce: Processor */
			reduce(7),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(8),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(8),		/* type, reduce: Processor */
			reduce(8),		/* sanitized, reduce: Processor */
			reduce(8),		/* parallel, reduce: Processor */
			reduce(8),		/* interpreter, reduce: Processor */
			reduce(8),		/* scheduler, reduce: Processor */
			reduce(8),		/* script, reduce: Processor */
			reduce(8),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(8),		/* wait, reduce: Processor */
			reduce(8),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(9),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(9),		/* type, reduce: Processor */
			reduce(9),		/* sanitized, reduce: Processor */
			reduce(9),		/* parallel, reduce: Processor */
			reduce(9),		/* interpreter, reduce: Processor */
			reduce(9),		/* scheduler, reduce: Processor */
			reduce(9),		/* script, reduce: Processor */
			reduce(9),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(9),		/* wait, reduce: Processor */
			reduce(9),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(10),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(10),		/* type, reduce: Processor */
			reduce(10),		/* sanitized, reduce: Processor */
			reduce(10),		/* parallel, reduce: Processor */
			reduce(10),		/* interpreter, reduce: Processor */
			reduce(10),		/* scheduler, reduce: Processor */
			reduce(10),		/* script, reduce: Processor */
			reduce(10),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(10),		/* wait, reduce: Processor */
			reduce(10),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Processor */
			nil,		/* empty */
			reduce(11),		/* name, reduce: Processor */
			nil,		/* : */
			reduce(11),		/* type, reduce: Processor */
			reduce(11),		/* sanitized, reduce: Processor */
			reduce(11),		/* parallel, reduce: Processor */
			reduce(11),		/* interpreter, reduce: Processor */
			reduce(11),		/* scheduler, reduce: Processor */
			reduce(11),		/* script, reduce: Processor */
			reduce(11),		/* outputs, reduce: Processor */
			nil,		/* , */
			reduce(11),		/* wait, reduce: Processor */
			reduce(11),		/* inputs, reduce: Processor */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(22),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(23),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(24),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(25),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(26),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(27),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(28),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(29),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(30),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(31),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(33),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(33),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			shift(36),		/* true */
			shift(37),		/* false */
			shift(38),		/* t */
			shift(39),		/* f */
			shift(40),		/* 0 */
			shift(41),		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			shift(43),		/* 0 */
			shift(44),		/* 1 */
			shift(45),		/* pos_int */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(47),		/* id */
			shift(48),		/* str_lit */
			shift(49),		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(33),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(47),		/* id */
			shift(48),		/* str_lit */
			shift(49),		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(55),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(58),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: NameStmt */
			nil,		/* empty */
			reduce(12),		/* name, reduce: NameStmt */
			nil,		/* : */
			reduce(12),		/* type, reduce: NameStmt */
			reduce(12),		/* sanitized, reduce: NameStmt */
			reduce(12),		/* parallel, reduce: NameStmt */
			reduce(12),		/* interpreter, reduce: NameStmt */
			reduce(12),		/* scheduler, reduce: NameStmt */
			reduce(12),		/* script, reduce: NameStmt */
			reduce(12),		/* outputs, reduce: NameStmt */
			nil,		/* , */
			reduce(12),		/* wait, reduce: NameStmt */
			reduce(12),		/* inputs, reduce: NameStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			nil,		/* , */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: TypeStmt */
			nil,		/* empty */
			reduce(13),		/* name, reduce: TypeStmt */
			nil,		/* : */
			reduce(13),		/* type, reduce: TypeStmt */
			reduce(13),		/* sanitized, reduce: TypeStmt */
			reduce(13),		/* parallel, reduce: TypeStmt */
			reduce(13),		/* interpreter, reduce: TypeStmt */
			reduce(13),		/* scheduler, reduce: TypeStmt */
			reduce(13),		/* script, reduce: TypeStmt */
			reduce(13),		/* outputs, reduce: TypeStmt */
			nil,		/* , */
			reduce(13),		/* wait, reduce: TypeStmt */
			reduce(13),		/* inputs, reduce: TypeStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: SanitInputStmt */
			nil,		/* empty */
			reduce(14),		/* name, reduce: SanitInputStmt */
			nil,		/* : */
			reduce(14),		/* type, reduce: SanitInputStmt */
			reduce(14),		/* sanitized, reduce: SanitInputStmt */
			reduce(14),		/* parallel, reduce: SanitInputStmt */
			reduce(14),		/* interpreter, reduce: SanitInputStmt */
			reduce(14),		/* scheduler, reduce: SanitInputStmt */
			reduce(14),		/* script, reduce: SanitInputStmt */
			reduce(14),		/* outputs, reduce: SanitInputStmt */
			nil,		/* , */
			reduce(14),		/* wait, reduce: SanitInputStmt */
			reduce(14),		/* inputs, reduce: SanitInputStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(48),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(48),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(48),		/* type, reduce: BoolVal */
			reduce(48),		/* sanitized, reduce: BoolVal */
			reduce(48),		/* parallel, reduce: BoolVal */
			reduce(48),		/* interpreter, reduce: BoolVal */
			reduce(48),		/* scheduler, reduce: BoolVal */
			reduce(48),		/* script, reduce: BoolVal */
			reduce(48),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(48),		/* wait, reduce: BoolVal */
			reduce(48),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(49),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(49),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(49),		/* type, reduce: BoolVal */
			reduce(49),		/* sanitized, reduce: BoolVal */
			reduce(49),		/* parallel, reduce: BoolVal */
			reduce(49),		/* interpreter, reduce: BoolVal */
			reduce(49),		/* scheduler, reduce: BoolVal */
			reduce(49),		/* script, reduce: BoolVal */
			reduce(49),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(49),		/* wait, reduce: BoolVal */
			reduce(49),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(50),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(50),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(50),		/* type, reduce: BoolVal */
			reduce(50),		/* sanitized, reduce: BoolVal */
			reduce(50),		/* parallel, reduce: BoolVal */
			reduce(50),		/* interpreter, reduce: BoolVal */
			reduce(50),		/* scheduler, reduce: BoolVal */
			reduce(50),		/* script, reduce: BoolVal */
			reduce(50),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(50),		/* wait, reduce: BoolVal */
			reduce(50),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(51),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(51),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(51),		/* type, reduce: BoolVal */
			reduce(51),		/* sanitized, reduce: BoolVal */
			reduce(51),		/* parallel, reduce: BoolVal */
			reduce(51),		/* interpreter, reduce: BoolVal */
			reduce(51),		/* scheduler, reduce: BoolVal */
			reduce(51),		/* script, reduce: BoolVal */
			reduce(51),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(51),		/* wait, reduce: BoolVal */
			reduce(51),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(52),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(52),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(52),		/* type, reduce: BoolVal */
			reduce(52),		/* sanitized, reduce: BoolVal */
			reduce(52),		/* parallel, reduce: BoolVal */
			reduce(52),		/* interpreter, reduce: BoolVal */
			reduce(52),		/* scheduler, reduce: BoolVal */
			reduce(52),		/* script, reduce: BoolVal */
			reduce(52),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(52),		/* wait, reduce: BoolVal */
			reduce(52),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(53),		/* $, reduce: BoolVal */
			nil,		/* empty */
			reduce(53),		/* name, reduce: BoolVal */
			nil,		/* : */
			reduce(53),		/* type, reduce: BoolVal */
			reduce(53),		/* sanitized, reduce: BoolVal */
			reduce(53),		/* parallel, reduce: BoolVal */
			reduce(53),		/* interpreter, reduce: BoolVal */
			reduce(53),		/* scheduler, reduce: BoolVal */
			reduce(53),		/* script, reduce: BoolVal */
			reduce(53),		/* outputs, reduce: BoolVal */
			nil,		/* , */
			reduce(53),		/* wait, reduce: BoolVal */
			reduce(53),		/* inputs, reduce: BoolVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: ParallelStmt */
			nil,		/* empty */
			reduce(15),		/* name, reduce: ParallelStmt */
			nil,		/* : */
			reduce(15),		/* type, reduce: ParallelStmt */
			reduce(15),		/* sanitized, reduce: ParallelStmt */
			reduce(15),		/* parallel, reduce: ParallelStmt */
			reduce(15),		/* interpreter, reduce: ParallelStmt */
			reduce(15),		/* scheduler, reduce: ParallelStmt */
			reduce(15),		/* script, reduce: ParallelStmt */
			reduce(15),		/* outputs, reduce: ParallelStmt */
			nil,		/* , */
			reduce(15),		/* wait, reduce: ParallelStmt */
			reduce(15),		/* inputs, reduce: ParallelStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(55),		/* $, reduce: IntVal */
			nil,		/* empty */
			reduce(55),		/* name, reduce: IntVal */
			nil,		/* : */
			reduce(55),		/* type, reduce: IntVal */
			reduce(55),		/* sanitized, reduce: IntVal */
			reduce(55),		/* parallel, reduce: IntVal */
			reduce(55),		/* interpreter, reduce: IntVal */
			reduce(55),		/* scheduler, reduce: IntVal */
			reduce(55),		/* script, reduce: IntVal */
			reduce(55),		/* outputs, reduce: IntVal */
			nil,		/* , */
			reduce(55),		/* wait, reduce: IntVal */
			reduce(55),		/* inputs, reduce: IntVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(56),		/* $, reduce: IntVal */
			nil,		/* empty */
			reduce(56),		/* name, reduce: IntVal */
			nil,		/* : */
			reduce(56),		/* type, reduce: IntVal */
			reduce(56),		/* sanitized, reduce: IntVal */
			reduce(56),		/* parallel, reduce: IntVal */
			reduce(56),		/* interpreter, reduce: IntVal */
			reduce(56),		/* scheduler, reduce: IntVal */
			reduce(56),		/* script, reduce: IntVal */
			reduce(56),		/* outputs, reduce: IntVal */
			nil,		/* , */
			reduce(56),		/* wait, reduce: IntVal */
			reduce(56),		/* inputs, reduce: IntVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(54),		/* $, reduce: IntVal */
			nil,		/* empty */
			reduce(54),		/* name, reduce: IntVal */
			nil,		/* : */
			reduce(54),		/* type, reduce: IntVal */
			reduce(54),		/* sanitized, reduce: IntVal */
			reduce(54),		/* parallel, reduce: IntVal */
			reduce(54),		/* interpreter, reduce: IntVal */
			reduce(54),		/* scheduler, reduce: IntVal */
			reduce(54),		/* script, reduce: IntVal */
			reduce(54),		/* outputs, reduce: IntVal */
			nil,		/* , */
			reduce(54),		/* wait, reduce: IntVal */
			reduce(54),		/* inputs, reduce: IntVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(16),		/* $, reduce: InterpreterStmt */
			nil,		/* empty */
			reduce(16),		/* name, reduce: InterpreterStmt */
			nil,		/* : */
			reduce(16),		/* type, reduce: InterpreterStmt */
			reduce(16),		/* sanitized, reduce: InterpreterStmt */
			reduce(16),		/* parallel, reduce: InterpreterStmt */
			reduce(16),		/* interpreter, reduce: InterpreterStmt */
			reduce(16),		/* scheduler, reduce: InterpreterStmt */
			reduce(16),		/* script, reduce: InterpreterStmt */
			reduce(16),		/* outputs, reduce: InterpreterStmt */
			nil,		/* , */
			reduce(16),		/* wait, reduce: InterpreterStmt */
			reduce(16),		/* inputs, reduce: InterpreterStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(46),		/* $, reduce: ScriptVal */
			nil,		/* empty */
			reduce(46),		/* name, reduce: ScriptVal */
			nil,		/* : */
			reduce(46),		/* type, reduce: ScriptVal */
			reduce(46),		/* sanitized, reduce: ScriptVal */
			reduce(46),		/* parallel, reduce: ScriptVal */
			reduce(46),		/* interpreter, reduce: ScriptVal */
			reduce(46),		/* scheduler, reduce: ScriptVal */
			reduce(46),		/* script, reduce: ScriptVal */
			reduce(46),		/* outputs, reduce: ScriptVal */
			nil,		/* , */
			reduce(46),		/* wait, reduce: ScriptVal */
			reduce(46),		/* inputs, reduce: ScriptVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(45),		/* $, reduce: ScriptVal */
			nil,		/* empty */
			reduce(45),		/* name, reduce: ScriptVal */
			nil,		/* : */
			reduce(45),		/* type, reduce: ScriptVal */
			reduce(45),		/* sanitized, reduce: ScriptVal */
			reduce(45),		/* parallel, reduce: ScriptVal */
			reduce(45),		/* interpreter, reduce: ScriptVal */
			reduce(45),		/* scheduler, reduce: ScriptVal */
			reduce(45),		/* script, reduce: ScriptVal */
			reduce(45),		/* outputs, reduce: ScriptVal */
			nil,		/* , */
			reduce(45),		/* wait, reduce: ScriptVal */
			reduce(45),		/* inputs, reduce: ScriptVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(47),		/* $, reduce: ScriptVal */
			nil,		/* empty */
			reduce(47),		/* name, reduce: ScriptVal */
			nil,		/* : */
			reduce(47),		/* type, reduce: ScriptVal */
			reduce(47),		/* sanitized, reduce: ScriptVal */
			reduce(47),		/* parallel, reduce: ScriptVal */
			reduce(47),		/* interpreter, reduce: ScriptVal */
			reduce(47),		/* scheduler, reduce: ScriptVal */
			reduce(47),		/* script, reduce: ScriptVal */
			reduce(47),		/* outputs, reduce: ScriptVal */
			nil,		/* , */
			reduce(47),		/* wait, reduce: ScriptVal */
			reduce(47),		/* inputs, reduce: ScriptVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: SchedulerStmt */
			nil,		/* empty */
			reduce(17),		/* name, reduce: SchedulerStmt */
			nil,		/* : */
			reduce(17),		/* type, reduce: SchedulerStmt */
			reduce(17),		/* sanitized, reduce: SchedulerStmt */
			reduce(17),		/* parallel, reduce: SchedulerStmt */
			reduce(17),		/* interpreter, reduce: SchedulerStmt */
			reduce(17),		/* scheduler, reduce: SchedulerStmt */
			reduce(17),		/* script, reduce: SchedulerStmt */
			reduce(17),		/* outputs, reduce: SchedulerStmt */
			nil,		/* , */
			reduce(17),		/* wait, reduce: SchedulerStmt */
			reduce(17),		/* inputs, reduce: SchedulerStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: ScriptStmt */
			nil,		/* empty */
			reduce(18),		/* name, reduce: ScriptStmt */
			nil,		/* : */
			reduce(18),		/* type, reduce: ScriptStmt */
			reduce(18),		/* sanitized, reduce: ScriptStmt */
			reduce(18),		/* parallel, reduce: ScriptStmt */
			reduce(18),		/* interpreter, reduce: ScriptStmt */
			reduce(18),		/* scheduler, reduce: ScriptStmt */
			reduce(18),		/* script, reduce: ScriptStmt */
			reduce(18),		/* outputs, reduce: ScriptStmt */
			nil,		/* , */
			reduce(18),		/* wait, reduce: ScriptStmt */
			reduce(18),		/* inputs, reduce: ScriptStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(40),		/* $, reduce: PortDef */
			nil,		/* empty */
			reduce(40),		/* name, reduce: PortDef */
			nil,		/* : */
			reduce(40),		/* type, reduce: PortDef */
			reduce(40),		/* sanitized, reduce: PortDef */
			reduce(40),		/* parallel, reduce: PortDef */
			reduce(40),		/* interpreter, reduce: PortDef */
			reduce(40),		/* scheduler, reduce: PortDef */
			reduce(40),		/* script, reduce: PortDef */
			reduce(40),		/* outputs, reduce: PortDef */
			reduce(40),		/* ,, reduce: PortDef */
			reduce(40),		/* wait, reduce: PortDef */
			reduce(40),		/* inputs, reduce: PortDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(71),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(19),		/* $, reduce: OutputStmt */
			nil,		/* empty */
			reduce(19),		/* name, reduce: OutputStmt */
			nil,		/* : */
			reduce(19),		/* type, reduce: OutputStmt */
			reduce(19),		/* sanitized, reduce: OutputStmt */
			reduce(19),		/* parallel, reduce: OutputStmt */
			reduce(19),		/* interpreter, reduce: OutputStmt */
			reduce(19),		/* scheduler, reduce: OutputStmt */
			reduce(19),		/* script, reduce: OutputStmt */
			reduce(19),		/* outputs, reduce: OutputStmt */
			shift(72),		/* , */
			reduce(19),		/* wait, reduce: OutputStmt */
			reduce(19),		/* inputs, reduce: OutputStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(20),		/* $, reduce: OutputList */
			nil,		/* empty */
			reduce(20),		/* name, reduce: OutputList */
			nil,		/* : */
			reduce(20),		/* type, reduce: OutputList */
			reduce(20),		/* sanitized, reduce: OutputList */
			reduce(20),		/* parallel, reduce: OutputList */
			reduce(20),		/* interpreter, reduce: OutputList */
			reduce(20),		/* scheduler, reduce: OutputList */
			reduce(20),		/* script, reduce: OutputList */
			reduce(20),		/* outputs, reduce: OutputList */
			reduce(20),		/* ,, reduce: OutputList */
			reduce(20),		/* wait, reduce: OutputList */
			reduce(20),		/* inputs, reduce: OutputList */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			reduce(44),		/* id, reduce: IdVal */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(23),		/* $, reduce: WaitList */
			nil,		/* empty */
			reduce(23),		/* name, reduce: WaitList */
			nil,		/* : */
			reduce(23),		/* type, reduce: WaitList */
			reduce(23),		/* sanitized, reduce: WaitList */
			reduce(23),		/* parallel, reduce: WaitList */
			reduce(23),		/* interpreter, reduce: WaitList */
			reduce(23),		/* scheduler, reduce: WaitList */
			reduce(23),		/* script, reduce: WaitList */
			reduce(23),		/* outputs, reduce: WaitList */
			reduce(23),		/* ,, reduce: WaitList */
			reduce(23),		/* wait, reduce: WaitList */
			reduce(23),		/* inputs, reduce: WaitList */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(22),		/* $, reduce: WaitStmt */
			nil,		/* empty */
			reduce(22),		/* name, reduce: WaitStmt */
			nil,		/* : */
			reduce(22),		/* type, reduce: WaitStmt */
			reduce(22),		/* sanitized, reduce: WaitStmt */
			reduce(22),		/* parallel, reduce: WaitStmt */
			reduce(22),		/* interpreter, reduce: WaitStmt */
			reduce(22),		/* scheduler, reduce: WaitStmt */
			reduce(22),		/* script, reduce: WaitStmt */
			reduce(22),		/* outputs, reduce: WaitStmt */
			shift(73),		/* , */
			reduce(22),		/* wait, reduce: WaitStmt */
			reduce(22),		/* inputs, reduce: WaitStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			reduce(40),		/* =, reduce: PortDef */
			shift(76),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(77),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(25),		/* $, reduce: IntputStmt */
			nil,		/* empty */
			reduce(25),		/* name, reduce: IntputStmt */
			nil,		/* : */
			reduce(25),		/* type, reduce: IntputStmt */
			reduce(25),		/* sanitized, reduce: IntputStmt */
			reduce(25),		/* parallel, reduce: IntputStmt */
			reduce(25),		/* interpreter, reduce: IntputStmt */
			reduce(25),		/* scheduler, reduce: IntputStmt */
			reduce(25),		/* script, reduce: IntputStmt */
			reduce(25),		/* outputs, reduce: IntputStmt */
			nil,		/* , */
			reduce(25),		/* wait, reduce: IntputStmt */
			reduce(25),		/* inputs, reduce: IntputStmt */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(26),		/* $, reduce: IterOper */
			nil,		/* empty */
			reduce(26),		/* name, reduce: IterOper */
			nil,		/* : */
			reduce(26),		/* type, reduce: IterOper */
			reduce(26),		/* sanitized, reduce: IterOper */
			reduce(26),		/* parallel, reduce: IterOper */
			reduce(26),		/* interpreter, reduce: IterOper */
			reduce(26),		/* scheduler, reduce: IterOper */
			reduce(26),		/* script, reduce: IterOper */
			reduce(26),		/* outputs, reduce: IterOper */
			nil,		/* , */
			reduce(26),		/* wait, reduce: IterOper */
			reduce(26),		/* inputs, reduce: IterOper */
			shift(78),		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(27),		/* $, reduce: IterOper */
			nil,		/* empty */
			reduce(27),		/* name, reduce: IterOper */
			nil,		/* : */
			reduce(27),		/* type, reduce: IterOper */
			reduce(27),		/* sanitized, reduce: IterOper */
			reduce(27),		/* parallel, reduce: IterOper */
			reduce(27),		/* interpreter, reduce: IterOper */
			reduce(27),		/* scheduler, reduce: IterOper */
			reduce(27),		/* script, reduce: IterOper */
			reduce(27),		/* outputs, reduce: IterOper */
			shift(79),		/* , */
			reduce(27),		/* wait, reduce: IterOper */
			reduce(27),		/* inputs, reduce: IterOper */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(28),		/* $, reduce: IterOper */
			nil,		/* empty */
			reduce(28),		/* name, reduce: IterOper */
			nil,		/* : */
			reduce(28),		/* type, reduce: IterOper */
			reduce(28),		/* sanitized, reduce: IterOper */
			reduce(28),		/* parallel, reduce: IterOper */
			reduce(28),		/* interpreter, reduce: IterOper */
			reduce(28),		/* scheduler, reduce: IterOper */
			reduce(28),		/* script, reduce: IterOper */
			reduce(28),		/* outputs, reduce: IterOper */
			nil,		/* , */
			reduce(28),		/* wait, reduce: IterOper */
			reduce(28),		/* inputs, reduce: IterOper */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(29),		/* $, reduce: IterOper */
			nil,		/* empty */
			reduce(29),		/* name, reduce: IterOper */
			nil,		/* : */
			reduce(29),		/* type, reduce: IterOper */
			reduce(29),		/* sanitized, reduce: IterOper */
			reduce(29),		/* parallel, reduce: IterOper */
			reduce(29),		/* interpreter, reduce: IterOper */
			reduce(29),		/* scheduler, reduce: IterOper */
			reduce(29),		/* script, reduce: IterOper */
			reduce(29),		/* outputs, reduce: IterOper */
			reduce(35),		/* ,, reduce: IterNode */
			reduce(29),		/* wait, reduce: IterOper */
			reduce(29),		/* inputs, reduce: IterOper */
			reduce(35),		/* *, reduce: IterNode */
			reduce(35),		/* <, reduce: IterNode */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			shift(80),		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			shift(81),		/* * */
			shift(82),		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			reduce(44),		/* =, reduce: IdVal */
			reduce(44),		/* id, reduce: IdVal */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(42),		/* $, reduce: TypeDef */
			nil,		/* empty */
			reduce(42),		/* name, reduce: TypeDef */
			nil,		/* : */
			reduce(42),		/* type, reduce: TypeDef */
			reduce(42),		/* sanitized, reduce: TypeDef */
			reduce(42),		/* parallel, reduce: TypeDef */
			reduce(42),		/* interpreter, reduce: TypeDef */
			reduce(42),		/* scheduler, reduce: TypeDef */
			reduce(42),		/* script, reduce: TypeDef */
			reduce(42),		/* outputs, reduce: TypeDef */
			reduce(42),		/* ,, reduce: TypeDef */
			reduce(42),		/* wait, reduce: TypeDef */
			reduce(42),		/* inputs, reduce: TypeDef */
			nil,		/* * */
			nil,		/* < */
			shift(90),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(41),		/* $, reduce: PortDef */
			nil,		/* empty */
			reduce(41),		/* name, reduce: PortDef */
			nil,		/* : */
			reduce(41),		/* type, reduce: PortDef */
			reduce(41),		/* sanitized, reduce: PortDef */
			reduce(41),		/* parallel, reduce: PortDef */
			reduce(41),		/* interpreter, reduce: PortDef */
			reduce(41),		/* scheduler, reduce: PortDef */
			reduce(41),		/* script, reduce: PortDef */
			reduce(41),		/* outputs, reduce: PortDef */
			reduce(41),		/* ,, reduce: PortDef */
			reduce(41),		/* wait, reduce: PortDef */
			reduce(41),		/* inputs, reduce: PortDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			reduce(44),		/* (, reduce: IdVal */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(55),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(58),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(93),		/* ( */
			nil,		/* ) */
			reduce(42),		/* =, reduce: TypeDef */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			reduce(41),		/* =, reduce: PortDef */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			reduce(44),		/* (, reduce: IdVal */
			nil,		/* ) */
			reduce(44),		/* =, reduce: IdVal */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(96),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S78
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(100),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S79
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(104),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S80
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(104),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(100),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S82
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(110),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S83
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(111),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S84
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(112),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S85
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			shift(113),		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(26),		/* ), reduce: IterOper */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S86
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			shift(114),		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(27),		/* ), reduce: IterOper */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S87
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(28),		/* ), reduce: IterOper */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S88
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(35),		/* ,, reduce: IterNode */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(35),		/* *, reduce: IterNode */
			reduce(35),		/* <, reduce: IterNode */
			nil,		/* ( */
			reduce(29),		/* ), reduce: IterOper */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S89
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			shift(115),		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			shift(116),		/* * */
			shift(117),		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S90
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			shift(119),		/* 0 */
			shift(120),		/* 1 */
			shift(121),		/* pos_int */
			
		},

	},
	actionRow{ // S91
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(21),		/* $, reduce: OutputList */
			nil,		/* empty */
			reduce(21),		/* name, reduce: OutputList */
			nil,		/* : */
			reduce(21),		/* type, reduce: OutputList */
			reduce(21),		/* sanitized, reduce: OutputList */
			reduce(21),		/* parallel, reduce: OutputList */
			reduce(21),		/* interpreter, reduce: OutputList */
			reduce(21),		/* scheduler, reduce: OutputList */
			reduce(21),		/* script, reduce: OutputList */
			reduce(21),		/* outputs, reduce: OutputList */
			reduce(21),		/* ,, reduce: OutputList */
			reduce(21),		/* wait, reduce: OutputList */
			reduce(21),		/* inputs, reduce: OutputList */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S92
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(24),		/* $, reduce: WaitList */
			nil,		/* empty */
			reduce(24),		/* name, reduce: WaitList */
			nil,		/* : */
			reduce(24),		/* type, reduce: WaitList */
			reduce(24),		/* sanitized, reduce: WaitList */
			reduce(24),		/* parallel, reduce: WaitList */
			reduce(24),		/* interpreter, reduce: WaitList */
			reduce(24),		/* scheduler, reduce: WaitList */
			reduce(24),		/* script, reduce: WaitList */
			reduce(24),		/* outputs, reduce: WaitList */
			reduce(24),		/* ,, reduce: WaitList */
			reduce(24),		/* wait, reduce: WaitList */
			reduce(24),		/* inputs, reduce: WaitList */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S93
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			shift(119),		/* 0 */
			shift(120),		/* 1 */
			shift(121),		/* pos_int */
			
		},

	},
	actionRow{ // S94
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(38),		/* name, reduce: SourceDef */
			shift(123),		/* : */
			reduce(38),		/* type, reduce: SourceDef */
			reduce(38),		/* sanitized, reduce: SourceDef */
			reduce(38),		/* parallel, reduce: SourceDef */
			reduce(38),		/* interpreter, reduce: SourceDef */
			reduce(38),		/* scheduler, reduce: SourceDef */
			reduce(38),		/* script, reduce: SourceDef */
			reduce(38),		/* outputs, reduce: SourceDef */
			reduce(38),		/* ,, reduce: SourceDef */
			reduce(38),		/* wait, reduce: SourceDef */
			reduce(38),		/* inputs, reduce: SourceDef */
			reduce(38),		/* *, reduce: SourceDef */
			reduce(38),		/* <, reduce: SourceDef */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S95
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(37),		/* $, reduce: IterPort */
			nil,		/* empty */
			reduce(37),		/* name, reduce: IterPort */
			nil,		/* : */
			reduce(37),		/* type, reduce: IterPort */
			reduce(37),		/* sanitized, reduce: IterPort */
			reduce(37),		/* parallel, reduce: IterPort */
			reduce(37),		/* interpreter, reduce: IterPort */
			reduce(37),		/* scheduler, reduce: IterPort */
			reduce(37),		/* script, reduce: IterPort */
			reduce(37),		/* outputs, reduce: IterPort */
			reduce(37),		/* ,, reduce: IterPort */
			reduce(37),		/* wait, reduce: IterPort */
			reduce(37),		/* inputs, reduce: IterPort */
			reduce(37),		/* *, reduce: IterPort */
			reduce(37),		/* <, reduce: IterPort */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S96
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			reduce(44),		/* :, reduce: IdVal */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			reduce(44),		/* *, reduce: IdVal */
			reduce(44),		/* <, reduce: IdVal */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S97
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(124),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S98
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(35),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(35),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(35),		/* type, reduce: IterNode */
			reduce(35),		/* sanitized, reduce: IterNode */
			reduce(35),		/* parallel, reduce: IterNode */
			reduce(35),		/* interpreter, reduce: IterNode */
			reduce(35),		/* scheduler, reduce: IterNode */
			reduce(35),		/* script, reduce: IterNode */
			reduce(35),		/* outputs, reduce: IterNode */
			nil,		/* , */
			reduce(35),		/* wait, reduce: IterNode */
			reduce(35),		/* inputs, reduce: IterNode */
			reduce(35),		/* *, reduce: IterNode */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S99
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(31),		/* $, reduce: IterOperCross */
			nil,		/* empty */
			reduce(31),		/* name, reduce: IterOperCross */
			nil,		/* : */
			reduce(31),		/* type, reduce: IterOperCross */
			reduce(31),		/* sanitized, reduce: IterOperCross */
			reduce(31),		/* parallel, reduce: IterOperCross */
			reduce(31),		/* interpreter, reduce: IterOperCross */
			reduce(31),		/* scheduler, reduce: IterOperCross */
			reduce(31),		/* script, reduce: IterOperCross */
			reduce(31),		/* outputs, reduce: IterOperCross */
			nil,		/* , */
			reduce(31),		/* wait, reduce: IterOperCross */
			reduce(31),		/* inputs, reduce: IterOperCross */
			reduce(31),		/* *, reduce: IterOperCross */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S100
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S101
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(126),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S102
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(35),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(35),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(35),		/* type, reduce: IterNode */
			reduce(35),		/* sanitized, reduce: IterNode */
			reduce(35),		/* parallel, reduce: IterNode */
			reduce(35),		/* interpreter, reduce: IterNode */
			reduce(35),		/* scheduler, reduce: IterNode */
			reduce(35),		/* script, reduce: IterNode */
			reduce(35),		/* outputs, reduce: IterNode */
			reduce(35),		/* ,, reduce: IterNode */
			reduce(35),		/* wait, reduce: IterNode */
			reduce(35),		/* inputs, reduce: IterNode */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S103
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(33),		/* $, reduce: IterOperDot */
			nil,		/* empty */
			reduce(33),		/* name, reduce: IterOperDot */
			nil,		/* : */
			reduce(33),		/* type, reduce: IterOperDot */
			reduce(33),		/* sanitized, reduce: IterOperDot */
			reduce(33),		/* parallel, reduce: IterOperDot */
			reduce(33),		/* interpreter, reduce: IterOperDot */
			reduce(33),		/* scheduler, reduce: IterOperDot */
			reduce(33),		/* script, reduce: IterOperDot */
			reduce(33),		/* outputs, reduce: IterOperDot */
			reduce(33),		/* ,, reduce: IterOperDot */
			reduce(33),		/* wait, reduce: IterOperDot */
			reduce(33),		/* inputs, reduce: IterOperDot */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S104
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S105
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(32),		/* $, reduce: IterOperDot */
			nil,		/* empty */
			reduce(32),		/* name, reduce: IterOperDot */
			nil,		/* : */
			reduce(32),		/* type, reduce: IterOperDot */
			reduce(32),		/* sanitized, reduce: IterOperDot */
			reduce(32),		/* parallel, reduce: IterOperDot */
			reduce(32),		/* interpreter, reduce: IterOperDot */
			reduce(32),		/* scheduler, reduce: IterOperDot */
			reduce(32),		/* script, reduce: IterOperDot */
			reduce(32),		/* outputs, reduce: IterOperDot */
			reduce(32),		/* ,, reduce: IterOperDot */
			reduce(32),		/* wait, reduce: IterOperDot */
			reduce(32),		/* inputs, reduce: IterOperDot */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S106
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(30),		/* $, reduce: IterOperCross */
			nil,		/* empty */
			reduce(30),		/* name, reduce: IterOperCross */
			nil,		/* : */
			reduce(30),		/* type, reduce: IterOperCross */
			reduce(30),		/* sanitized, reduce: IterOperCross */
			reduce(30),		/* parallel, reduce: IterOperCross */
			reduce(30),		/* interpreter, reduce: IterOperCross */
			reduce(30),		/* scheduler, reduce: IterOperCross */
			reduce(30),		/* script, reduce: IterOperCross */
			reduce(30),		/* outputs, reduce: IterOperCross */
			nil,		/* , */
			reduce(30),		/* wait, reduce: IterOperCross */
			reduce(30),		/* inputs, reduce: IterOperCross */
			reduce(30),		/* *, reduce: IterOperCross */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S107
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(128),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S108
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(35),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(35),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(35),		/* type, reduce: IterNode */
			reduce(35),		/* sanitized, reduce: IterNode */
			reduce(35),		/* parallel, reduce: IterNode */
			reduce(35),		/* interpreter, reduce: IterNode */
			reduce(35),		/* scheduler, reduce: IterNode */
			reduce(35),		/* script, reduce: IterNode */
			reduce(35),		/* outputs, reduce: IterNode */
			nil,		/* , */
			reduce(35),		/* wait, reduce: IterNode */
			reduce(35),		/* inputs, reduce: IterNode */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S109
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(34),		/* $, reduce: IterOperPrefix */
			nil,		/* empty */
			reduce(34),		/* name, reduce: IterOperPrefix */
			nil,		/* : */
			reduce(34),		/* type, reduce: IterOperPrefix */
			reduce(34),		/* sanitized, reduce: IterOperPrefix */
			reduce(34),		/* parallel, reduce: IterOperPrefix */
			reduce(34),		/* interpreter, reduce: IterOperPrefix */
			reduce(34),		/* scheduler, reduce: IterOperPrefix */
			reduce(34),		/* script, reduce: IterOperPrefix */
			reduce(34),		/* outputs, reduce: IterOperPrefix */
			nil,		/* , */
			reduce(34),		/* wait, reduce: IterOperPrefix */
			reduce(34),		/* inputs, reduce: IterOperPrefix */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S110
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S111
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(132),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S112
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(36),		/* ,, reduce: IterNode */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(36),		/* *, reduce: IterNode */
			reduce(36),		/* <, reduce: IterNode */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S113
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(136),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S114
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(140),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S115
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(140),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S116
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(136),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S117
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(146),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S118
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(147),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S119
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(55),		/* ), reduce: IntVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S120
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(56),		/* ), reduce: IntVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S121
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(54),		/* ), reduce: IntVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S122
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(148),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S123
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(150),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S124
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(153),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S125
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(154),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S126
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(157),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S127
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(158),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S128
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(161),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S129
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(162),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S130
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(163),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(38),		/* ,, reduce: SourceDef */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(38),		/* *, reduce: SourceDef */
			reduce(38),		/* <, reduce: SourceDef */
			nil,		/* ( */
			reduce(38),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S131
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(37),		/* ,, reduce: IterPort */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(37),		/* *, reduce: IterPort */
			reduce(37),		/* <, reduce: IterPort */
			nil,		/* ( */
			reduce(37),		/* ), reduce: IterPort */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S132
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			reduce(44),		/* :, reduce: IdVal */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(44),		/* ,, reduce: IdVal */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(44),		/* *, reduce: IdVal */
			reduce(44),		/* <, reduce: IdVal */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S133
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(164),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S134
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(35),		/* *, reduce: IterNode */
			nil,		/* < */
			nil,		/* ( */
			reduce(35),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S135
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(31),		/* *, reduce: IterOperCross */
			nil,		/* < */
			nil,		/* ( */
			reduce(31),		/* ), reduce: IterOperCross */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S136
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S137
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(166),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S138
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(35),		/* ,, reduce: IterNode */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(35),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S139
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(33),		/* ,, reduce: IterOperDot */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(33),		/* ), reduce: IterOperDot */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S140
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S141
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(32),		/* ,, reduce: IterOperDot */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(32),		/* ), reduce: IterOperDot */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S142
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(30),		/* *, reduce: IterOperCross */
			nil,		/* < */
			nil,		/* ( */
			reduce(30),		/* ), reduce: IterOperCross */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S143
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			shift(168),		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S144
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(35),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S145
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(34),		/* ), reduce: IterOperPrefix */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S146
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			shift(67),		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(68),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S147
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(43),		/* $, reduce: TypeDef */
			nil,		/* empty */
			reduce(43),		/* name, reduce: TypeDef */
			nil,		/* : */
			reduce(43),		/* type, reduce: TypeDef */
			reduce(43),		/* sanitized, reduce: TypeDef */
			reduce(43),		/* parallel, reduce: TypeDef */
			reduce(43),		/* interpreter, reduce: TypeDef */
			reduce(43),		/* scheduler, reduce: TypeDef */
			reduce(43),		/* script, reduce: TypeDef */
			reduce(43),		/* outputs, reduce: TypeDef */
			reduce(43),		/* ,, reduce: TypeDef */
			reduce(43),		/* wait, reduce: TypeDef */
			reduce(43),		/* inputs, reduce: TypeDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S148
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			reduce(43),		/* =, reduce: TypeDef */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S149
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(39),		/* name, reduce: SourceDef */
			nil,		/* : */
			reduce(39),		/* type, reduce: SourceDef */
			reduce(39),		/* sanitized, reduce: SourceDef */
			reduce(39),		/* parallel, reduce: SourceDef */
			reduce(39),		/* interpreter, reduce: SourceDef */
			reduce(39),		/* scheduler, reduce: SourceDef */
			reduce(39),		/* script, reduce: SourceDef */
			reduce(39),		/* outputs, reduce: SourceDef */
			reduce(39),		/* ,, reduce: SourceDef */
			reduce(39),		/* wait, reduce: SourceDef */
			reduce(39),		/* inputs, reduce: SourceDef */
			reduce(39),		/* *, reduce: SourceDef */
			reduce(39),		/* <, reduce: SourceDef */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S150
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			reduce(44),		/* *, reduce: IdVal */
			reduce(44),		/* <, reduce: IdVal */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S151
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(38),		/* name, reduce: SourceDef */
			shift(170),		/* : */
			reduce(38),		/* type, reduce: SourceDef */
			reduce(38),		/* sanitized, reduce: SourceDef */
			reduce(38),		/* parallel, reduce: SourceDef */
			reduce(38),		/* interpreter, reduce: SourceDef */
			reduce(38),		/* scheduler, reduce: SourceDef */
			reduce(38),		/* script, reduce: SourceDef */
			reduce(38),		/* outputs, reduce: SourceDef */
			nil,		/* , */
			reduce(38),		/* wait, reduce: SourceDef */
			reduce(38),		/* inputs, reduce: SourceDef */
			reduce(38),		/* *, reduce: SourceDef */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S152
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(37),		/* $, reduce: IterPort */
			nil,		/* empty */
			reduce(37),		/* name, reduce: IterPort */
			nil,		/* : */
			reduce(37),		/* type, reduce: IterPort */
			reduce(37),		/* sanitized, reduce: IterPort */
			reduce(37),		/* parallel, reduce: IterPort */
			reduce(37),		/* interpreter, reduce: IterPort */
			reduce(37),		/* scheduler, reduce: IterPort */
			reduce(37),		/* script, reduce: IterPort */
			reduce(37),		/* outputs, reduce: IterPort */
			nil,		/* , */
			reduce(37),		/* wait, reduce: IterPort */
			reduce(37),		/* inputs, reduce: IterPort */
			reduce(37),		/* *, reduce: IterPort */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S153
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			reduce(44),		/* :, reduce: IdVal */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			nil,		/* , */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			reduce(44),		/* *, reduce: IdVal */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S154
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(36),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(36),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(36),		/* type, reduce: IterNode */
			reduce(36),		/* sanitized, reduce: IterNode */
			reduce(36),		/* parallel, reduce: IterNode */
			reduce(36),		/* interpreter, reduce: IterNode */
			reduce(36),		/* scheduler, reduce: IterNode */
			reduce(36),		/* script, reduce: IterNode */
			reduce(36),		/* outputs, reduce: IterNode */
			nil,		/* , */
			reduce(36),		/* wait, reduce: IterNode */
			reduce(36),		/* inputs, reduce: IterNode */
			reduce(36),		/* *, reduce: IterNode */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S155
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(38),		/* name, reduce: SourceDef */
			shift(171),		/* : */
			reduce(38),		/* type, reduce: SourceDef */
			reduce(38),		/* sanitized, reduce: SourceDef */
			reduce(38),		/* parallel, reduce: SourceDef */
			reduce(38),		/* interpreter, reduce: SourceDef */
			reduce(38),		/* scheduler, reduce: SourceDef */
			reduce(38),		/* script, reduce: SourceDef */
			reduce(38),		/* outputs, reduce: SourceDef */
			reduce(38),		/* ,, reduce: SourceDef */
			reduce(38),		/* wait, reduce: SourceDef */
			reduce(38),		/* inputs, reduce: SourceDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S156
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(37),		/* $, reduce: IterPort */
			nil,		/* empty */
			reduce(37),		/* name, reduce: IterPort */
			nil,		/* : */
			reduce(37),		/* type, reduce: IterPort */
			reduce(37),		/* sanitized, reduce: IterPort */
			reduce(37),		/* parallel, reduce: IterPort */
			reduce(37),		/* interpreter, reduce: IterPort */
			reduce(37),		/* scheduler, reduce: IterPort */
			reduce(37),		/* script, reduce: IterPort */
			reduce(37),		/* outputs, reduce: IterPort */
			reduce(37),		/* ,, reduce: IterPort */
			reduce(37),		/* wait, reduce: IterPort */
			reduce(37),		/* inputs, reduce: IterPort */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S157
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			reduce(44),		/* :, reduce: IdVal */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			reduce(44),		/* ,, reduce: IdVal */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S158
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(36),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(36),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(36),		/* type, reduce: IterNode */
			reduce(36),		/* sanitized, reduce: IterNode */
			reduce(36),		/* parallel, reduce: IterNode */
			reduce(36),		/* interpreter, reduce: IterNode */
			reduce(36),		/* scheduler, reduce: IterNode */
			reduce(36),		/* script, reduce: IterNode */
			reduce(36),		/* outputs, reduce: IterNode */
			reduce(36),		/* ,, reduce: IterNode */
			reduce(36),		/* wait, reduce: IterNode */
			reduce(36),		/* inputs, reduce: IterNode */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S159
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(38),		/* name, reduce: SourceDef */
			shift(172),		/* : */
			reduce(38),		/* type, reduce: SourceDef */
			reduce(38),		/* sanitized, reduce: SourceDef */
			reduce(38),		/* parallel, reduce: SourceDef */
			reduce(38),		/* interpreter, reduce: SourceDef */
			reduce(38),		/* scheduler, reduce: SourceDef */
			reduce(38),		/* script, reduce: SourceDef */
			reduce(38),		/* outputs, reduce: SourceDef */
			nil,		/* , */
			reduce(38),		/* wait, reduce: SourceDef */
			reduce(38),		/* inputs, reduce: SourceDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S160
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(37),		/* $, reduce: IterPort */
			nil,		/* empty */
			reduce(37),		/* name, reduce: IterPort */
			nil,		/* : */
			reduce(37),		/* type, reduce: IterPort */
			reduce(37),		/* sanitized, reduce: IterPort */
			reduce(37),		/* parallel, reduce: IterPort */
			reduce(37),		/* interpreter, reduce: IterPort */
			reduce(37),		/* scheduler, reduce: IterPort */
			reduce(37),		/* script, reduce: IterPort */
			reduce(37),		/* outputs, reduce: IterPort */
			nil,		/* , */
			reduce(37),		/* wait, reduce: IterPort */
			reduce(37),		/* inputs, reduce: IterPort */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S161
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			reduce(44),		/* :, reduce: IdVal */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			nil,		/* , */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S162
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(36),		/* $, reduce: IterNode */
			nil,		/* empty */
			reduce(36),		/* name, reduce: IterNode */
			nil,		/* : */
			reduce(36),		/* type, reduce: IterNode */
			reduce(36),		/* sanitized, reduce: IterNode */
			reduce(36),		/* parallel, reduce: IterNode */
			reduce(36),		/* interpreter, reduce: IterNode */
			reduce(36),		/* scheduler, reduce: IterNode */
			reduce(36),		/* script, reduce: IterNode */
			reduce(36),		/* outputs, reduce: IterNode */
			nil,		/* , */
			reduce(36),		/* wait, reduce: IterNode */
			reduce(36),		/* inputs, reduce: IterNode */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S163
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(174),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S164
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(177),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S165
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(178),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S166
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(181),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S167
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(182),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S168
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(185),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S169
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			shift(186),		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S170
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(188),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S171
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(58),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S172
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(33),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S173
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(39),		/* ,, reduce: SourceDef */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(39),		/* *, reduce: SourceDef */
			reduce(39),		/* <, reduce: SourceDef */
			nil,		/* ( */
			reduce(39),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S174
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(44),		/* ,, reduce: IdVal */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(44),		/* *, reduce: IdVal */
			reduce(44),		/* <, reduce: IdVal */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S175
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(191),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(38),		/* *, reduce: SourceDef */
			nil,		/* < */
			nil,		/* ( */
			reduce(38),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S176
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(37),		/* *, reduce: IterPort */
			nil,		/* < */
			nil,		/* ( */
			reduce(37),		/* ), reduce: IterPort */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S177
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			reduce(44),		/* :, reduce: IdVal */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(44),		/* *, reduce: IdVal */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S178
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(36),		/* *, reduce: IterNode */
			nil,		/* < */
			nil,		/* ( */
			reduce(36),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S179
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(192),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(38),		/* ,, reduce: SourceDef */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(38),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S180
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(37),		/* ,, reduce: IterPort */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(37),		/* ), reduce: IterPort */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S181
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			reduce(44),		/* :, reduce: IdVal */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(44),		/* ,, reduce: IdVal */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S182
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(36),		/* ,, reduce: IterNode */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(36),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S183
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			shift(193),		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(38),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S184
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(37),		/* ), reduce: IterPort */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S185
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			reduce(44),		/* :, reduce: IdVal */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S186
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(36),		/* ), reduce: IterNode */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S187
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(39),		/* name, reduce: SourceDef */
			nil,		/* : */
			reduce(39),		/* type, reduce: SourceDef */
			reduce(39),		/* sanitized, reduce: SourceDef */
			reduce(39),		/* parallel, reduce: SourceDef */
			reduce(39),		/* interpreter, reduce: SourceDef */
			reduce(39),		/* scheduler, reduce: SourceDef */
			reduce(39),		/* script, reduce: SourceDef */
			reduce(39),		/* outputs, reduce: SourceDef */
			nil,		/* , */
			reduce(39),		/* wait, reduce: SourceDef */
			reduce(39),		/* inputs, reduce: SourceDef */
			reduce(39),		/* *, reduce: SourceDef */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S188
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(44),		/* $, reduce: IdVal */
			nil,		/* empty */
			reduce(44),		/* name, reduce: IdVal */
			nil,		/* : */
			reduce(44),		/* type, reduce: IdVal */
			reduce(44),		/* sanitized, reduce: IdVal */
			reduce(44),		/* parallel, reduce: IdVal */
			reduce(44),		/* interpreter, reduce: IdVal */
			reduce(44),		/* scheduler, reduce: IdVal */
			reduce(44),		/* script, reduce: IdVal */
			reduce(44),		/* outputs, reduce: IdVal */
			nil,		/* , */
			reduce(44),		/* wait, reduce: IdVal */
			reduce(44),		/* inputs, reduce: IdVal */
			reduce(44),		/* *, reduce: IdVal */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S189
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(39),		/* name, reduce: SourceDef */
			nil,		/* : */
			reduce(39),		/* type, reduce: SourceDef */
			reduce(39),		/* sanitized, reduce: SourceDef */
			reduce(39),		/* parallel, reduce: SourceDef */
			reduce(39),		/* interpreter, reduce: SourceDef */
			reduce(39),		/* scheduler, reduce: SourceDef */
			reduce(39),		/* script, reduce: SourceDef */
			reduce(39),		/* outputs, reduce: SourceDef */
			reduce(39),		/* ,, reduce: SourceDef */
			reduce(39),		/* wait, reduce: SourceDef */
			reduce(39),		/* inputs, reduce: SourceDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S190
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: SourceDef */
			nil,		/* empty */
			reduce(39),		/* name, reduce: SourceDef */
			nil,		/* : */
			reduce(39),		/* type, reduce: SourceDef */
			reduce(39),		/* sanitized, reduce: SourceDef */
			reduce(39),		/* parallel, reduce: SourceDef */
			reduce(39),		/* interpreter, reduce: SourceDef */
			reduce(39),		/* scheduler, reduce: SourceDef */
			reduce(39),		/* script, reduce: SourceDef */
			reduce(39),		/* outputs, reduce: SourceDef */
			nil,		/* , */
			reduce(39),		/* wait, reduce: SourceDef */
			reduce(39),		/* inputs, reduce: SourceDef */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S191
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(195),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S192
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(197),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S193
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* = */
			shift(199),		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S194
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(39),		/* *, reduce: SourceDef */
			nil,		/* < */
			nil,		/* ( */
			reduce(39),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S195
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			reduce(44),		/* *, reduce: IdVal */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S196
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(39),		/* ,, reduce: SourceDef */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(39),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S197
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			reduce(44),		/* ,, reduce: IdVal */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S198
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(39),		/* ), reduce: SourceDef */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	actionRow{ // S199
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* name */
			nil,		/* : */
			nil,		/* type */
			nil,		/* sanitized */
			nil,		/* parallel */
			nil,		/* interpreter */
			nil,		/* scheduler */
			nil,		/* script */
			nil,		/* outputs */
			nil,		/* , */
			nil,		/* wait */
			nil,		/* inputs */
			nil,		/* * */
			nil,		/* < */
			nil,		/* ( */
			reduce(44),		/* ), reduce: IdVal */
			nil,		/* = */
			nil,		/* id */
			nil,		/* str_lit */
			nil,		/* path */
			nil,		/* true */
			nil,		/* false */
			nil,		/* t */
			nil,		/* f */
			nil,		/* 0 */
			nil,		/* 1 */
			nil,		/* pos_int */
			
		},

	},
	
}

