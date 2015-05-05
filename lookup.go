package abbr

import (
	"strings"
)

var Names = map[string]string{
	"abbreviat":            "abbrev",
	"absolut":              "abs",
	"account":              "acct",
	"accumulat":            "accum",
	"accuracy":             "acc",
	"accurat":              "acc",
	"actual":               "act",
	"addend":               "adn",
	"address":              "addr",
	"aggregat":             "agg",
	"algebra":              "alg",
	"algorithm":            "algo",
	"ancestor":             "anc",
	"approximat":           "approx",
	"architect":            "arch",
	"argument":             "arg",
	"ascending":            "asc",
	"attribut":             "attr",
	"averag":               "avg",
	"binary":               "bin",
	"binomial":             "binom",
	"boolean":              "bool",
	"buffer":               "buf",
	"calculat":             "calc",
	"calculus":             "calc",
	"calendar":             "cal",
	"capacity":             "cap",
	"capitalization":       "cap",
	"capitaliz":            "cap",
	"ceiling":              "ceil",
	"certificat":           "cert",
	"certify":              "cert",
	"channel":              "chan",
	"character":            "char",
	"coefficient":          "coef",
	"collect":              "coll",
	"column":               "col",
	"combination":          "comb",
	"combin":               "comb",
	"command":              "cmd",
	"compar":               "comp",
	"compensat":            "comp",
	"complex":              "cop",
	"condition":            "cond",
	"configuration":        "conf",
	"configur":             "conf",
	"connect":              "conn",
	"constant":             "const",
	"contain":              "cont",
	"context":              "ctx",
	"continuation":         "cont",
	"continu":              "cont",
	"control":              "ctl",
	"conversion":           "conv",
	"convert":              "conv",
	"cosecant":             "csc",
	"cosin":                "cos",
	"cotangent":            "cot",
	"count":                "cnt",
	"current":              "cur",
	"decimal":              "dec",
	"declaration":          "decl",
	"declar":               "decl",
	"decod":                "dec",
	"decrement":            "dec",
	"decrypt":              "dec",
	"defin":                "def",
	"definition":           "def",
	"degre":                "deg",
	"delet":                "del",
	"deletion":             "del",
	"delimiter":            "delim",
	"denominat":            "den",
	"depend":               "dep",
	"descendant":           "des",
	"descending":           "desc",
	"describ":              "desc",
	"description":          "desc",
	"destination":          "dest",
	"destin":               "dest",
	"determinant":          "det",
	"develop":              "dev",
	"deviat":               "dev",
	"diagonal":             "diag",
	"diameter":             "diam",
	"dictionary":           "dict",
	"differenc":            "diff",
	"dimension":            "dim",
	"direct":               "dir",
	"discriminant":         "disc",
	"distanc":              "dist",
	"distribut":            "dist",
	"divid":                "div",
	"document":             "doc",
	"domain":               "dom",
	"duplicat":             "dupe",
	"element":              "elem",
	"employe":              "emp",
	"encod":                "enc",
	"encrypt":              "enc",
	"entropy":              "ent",
	"enumerat":             "enum",
	"environment":          "env",
	"equivalenc":           "equiv",
	"equivalent":           "equiv",
	"error":                "err",
	"escap":                "esc",
	"estimat":              "est",
	"evaluat":              "eval",
	"except":               "excp",
	"exclud":               "excl",
	"exclusion":            "excl",
	"exclusiveor":          "xor",
	"execut":               "exec",
	"execution":            "exec",
	"expect":               "exp",
	"exponent":             "exp",
	"express":              "expr",
	"extend":               "ext",
	"extension":            "ext",
	"external":             "ext",
	"factor":               "fac",
	"factorial":            "facl",
	"field":                "fld",
	"figur":                "fig",
	"floating":             "float",
	"format":               "fmt",
	"forward":              "fwd",
	"fraction":             "frac",
	"frequency":            "freq",
	"frequent":             "freq",
	"function":             "func",
	"general":              "gen",
	"generat":              "gen",
	"generic":              "gen",
	"geometric":            "geom",
	"guarante":             "guar",
	"handl":                "hdl",
	"harmonic":             "har",
	"header":               "hdr",
	"hexadecimal":          "hex",
	"horizontal":           "horiz",
	"identification":       "ident",
	"identifier":           "ident",
	"imag":                 "img",
	"imaginary":            "imag",
	"immediat":             "imm",
	"implement":            "imp",
	"includ":               "incl",
	"inclusion":            "incl",
	"increment":            "inc",
	"independent":          "indep",
	"indicat":              "ind",
	"indirect":             "ind",
	"information":          "info",
	"initial":              "init",
	"initialization":       "init",
	"input":                "in",
	"insert":               "ins",
	"instanc":              "inst",
	"instantiat":           "inst",
	"instruct":             "inst",
	"integer":              "int",
	"intercept":            "intc",
	"interfac":             "intf",
	"internal":             "int",
	"internationalization": "i18n",
	"intersection":         "intx",
	"invers":               "inv",
	"iterat":               "iter",
	"kurtosis":             "kurt",
	"languag":              "lang",
	"length":               "len",
	"lexicon":              "lex",
	"library":              "lib",
	"limit":                "lim",
	"linear":               "lin",
	"literal":              "lit",
	"localization":         "l10n",
	"locat":                "loc",
	"logarithm":            "log",
	"lowercas":             "low",
	"machin":               "mach",
	"mantissa":             "mant",
	"maximum":              "max",
	"median":               "med",
	"medium":               "med",
	"memory":               "mem",
	"messag":               "msg",
	"minimum":              "min",
	"minut":                "min",
	"modification":         "mod",
	"modify":               "mod",
	"modul":                "mod",
	"modulo":               "mod",
	"multipl":              "multi",
	"multiply":             "mul",
	"mutabl":               "mut",
	"mutat":                "mut",
	"mutual":               "mutl",
	"mutualexclusion":      "mutex",
	"natural":              "nat",
	"negat":                "neg",
	"negativ":              "neg",
	"network":              "net",
	"number":               "num",
	"numerat":              "num",
	"object":               "obj",
	"octal":                "oct",
	"operat":               "op",
	"option":               "opt",
	"original":             "orig",
	"originat":             "orig",
	"output":               "out",
	"overflow":             "ovfl",
	"packag":               "pkg",
	"paragraph":            "para",
	"parallel":             "parl",
	"parameter":            "param",
	"parent":               "par",
	"pattern":              "pat",
	"perform":              "perf",
	"perimeter":            "perim",
	"permanent":            "perm",
	"permutat":             "perm",
	"perpendicular":        "perp",
	"phonetic":             "phon",
	"pictur":               "pic",
	"pointer":              "ptr",
	"position":             "pos",
	"positiv":              "pos",
	"possibl":              "poss",
	"possibly":             "poss",
	"precis":               "prec",
	"precision":            "prec",
	"preparation":          "prep",
	"prepar":               "prep",
	"previous":             "prev",
	"primary":              "pri",
	"prioritiz":            "prio",
	"priority":             "prio",
	"privat":               "pri",
	"probability":          "prob",
	"procedur":             "proc",
	"process":              "proc",
	"product":              "prod",
	"profil":               "prof",
	"program":              "prog",
	"progress":             "prog",
	"project":              "proj",
	"pronounc":             "pron",
	"pronunciation":        "pron",
	"proportion":           "prop",
	"protocol":             "prot",
	"prototyp":             "proto",
	"punctuat":             "punct",
	"public":               "pub",
	"quadratic":            "quad",
	"query":                "qry",
	"quotient":             "quot",
	"random":               "rand",
	"ratio":                "rat",
	"receiv":               "recv",
	"record":               "rec",
	"rectangl":             "rect",
	"recursiv":             "rec",
	"referenc":             "ref",
	"register":             "reg",
	"regular":              "reg",
	"regularexpression":    "regexp",
	"relat":                "rel",
	"relativ":              "rel",
	"remainder":            "rem",
	"repeat":               "rep",
	"repetition":           "rep",
	"replac":               "repl",
	"repository":           "repo",
	"represent":            "rep",
	"requir":               "req",
	"resourc":              "res",
	"respond":              "resp",
	"respons":              "resp",
	"responsiv":            "resp",
	"result":               "res",
	"return":               "ret",
	"revers":               "rev",
	"sanitization":         "san",
	"sanitiz":              "san",
	"schedul":              "sched",
	"scienc":               "sci",
	"scientific":           "sci",
	"secant":               "sec",
	"second":               "sec",
	"section":              "sect",
	"secur":                "sec",
	"security":             "sec",
	"segment":              "seg",
	"select":               "sel",
	"sentenc":              "sent",
	"separat":              "sep",
	"sequenc":              "seq",
	"server":               "srv",
	"signal":               "sig",
	"signatur":             "sig",
	"sin":                  "sin",
	"socket":               "sock",
	"sourc":                "src",
	"specific":             "spec",
	"specify":              "spec",
	"standard":             "std",
	"statement":            "stm",
	"statistic":            "stat",
	"string":               "str",
	"structur":             "struct",
	"subtract":             "sub",
	"surfac":               "surf",
	"symbol":               "sym",
	"symmetry":             "symm",
	"synchroniz":           "sync",
	"system":               "sys",
	"tabl":                 "tbl",
	"tangent":              "tan",
	"technology":           "tech",
	"templat":              "temp",
	"temporary":            "temp",
	"terminal":             "term",
	"terminat":             "term",
	"token":                "tok",
	"transfer":             "xfer",
	"translat":             "xlat",
	"transmission":         "xmis",
	"transmit":             "xmit",
	"trigonometry":         "trig",
	"updat":                "upd",
	"uppercas":             "up",
	"utility":              "util",
	"validat":              "val",
	"valu":                 "val",
	"variabl":              "var",
	"varianc":              "var",
	"vector":               "vec",
	"verification":         "ver",
	"verify":               "ver",
	"version":              "ver",
	"vertical":             "vert",
	"visibility":           "vis",
	"visibl":               "vis",
	"volum":                "vol",
	"window":               "win",
}

// LookUp looks up a word to find its short form, if it is known.
func LookUp(word string) string {
	lowWord := strings.ToLower(word)
	for long, short := range Names {
		if strings.Contains(lowWord, long) {
			return short
		}
	}
	return ""
}
