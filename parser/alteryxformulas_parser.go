// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 222,
	4, 2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 69, 10, 2, 12, 2,
	14, 2, 72, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 6, 2, 121, 10, 2, 13, 2, 14, 2, 122, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 6, 2, 132, 10, 2, 13, 2, 14, 2, 133, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2,
	150, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 194, 10, 2, 12, 2,
	14, 2, 197, 11, 2, 5, 2, 199, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 7, 2, 209, 10, 2, 12, 2, 14, 2, 212, 11, 2, 5, 2, 214, 10,
	2, 3, 2, 7, 2, 217, 10, 2, 12, 2, 14, 2, 220, 11, 2, 3, 2, 2, 3, 2, 3,
	2, 2, 4, 4, 2, 16, 16, 37, 37, 4, 2, 17, 17, 38, 38, 2, 269, 2, 149, 3,
	2, 2, 2, 4, 5, 8, 2, 1, 2, 5, 6, 7, 3, 2, 2, 6, 7, 5, 2, 2, 2, 7, 8, 7,
	4, 2, 2, 8, 150, 3, 2, 2, 2, 9, 10, 7, 39, 2, 2, 10, 11, 5, 2, 2, 2, 11,
	12, 7, 40, 2, 2, 12, 20, 5, 2, 2, 2, 13, 14, 7, 42, 2, 2, 14, 15, 5, 2,
	2, 2, 15, 16, 7, 40, 2, 2, 16, 17, 5, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18,
	13, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2,
	2, 21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7, 41, 2, 2, 24, 25,
	5, 2, 2, 2, 25, 26, 7, 43, 2, 2, 26, 150, 3, 2, 2, 2, 27, 28, 7, 34, 2,
	2, 28, 29, 7, 3, 2, 2, 29, 30, 5, 2, 2, 2, 30, 31, 7, 15, 2, 2, 31, 32,
	5, 2, 2, 2, 32, 33, 7, 15, 2, 2, 33, 34, 5, 2, 2, 2, 34, 35, 7, 4, 2, 2,
	35, 150, 3, 2, 2, 2, 36, 37, 7, 19, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5,
	2, 2, 2, 39, 40, 7, 4, 2, 2, 40, 150, 3, 2, 2, 2, 41, 42, 7, 20, 2, 2,
	42, 43, 7, 3, 2, 2, 43, 44, 5, 2, 2, 2, 44, 45, 7, 4, 2, 2, 45, 150, 3,
	2, 2, 2, 46, 47, 7, 21, 2, 2, 47, 48, 7, 3, 2, 2, 48, 49, 5, 2, 2, 2, 49,
	50, 7, 4, 2, 2, 50, 150, 3, 2, 2, 2, 51, 52, 7, 22, 2, 2, 52, 53, 7, 3,
	2, 2, 53, 54, 5, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 150, 3, 2, 2, 2, 56,
	57, 7, 23, 2, 2, 57, 58, 7, 3, 2, 2, 58, 59, 5, 2, 2, 2, 59, 60, 7, 15,
	2, 2, 60, 61, 5, 2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 150, 3, 2, 2, 2, 63,
	64, 7, 24, 2, 2, 64, 65, 7, 3, 2, 2, 65, 70, 5, 2, 2, 2, 66, 67, 7, 15,
	2, 2, 67, 69, 5, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2,
	73, 74, 7, 4, 2, 2, 74, 150, 3, 2, 2, 2, 75, 76, 7, 25, 2, 2, 76, 77, 7,
	3, 2, 2, 77, 78, 5, 2, 2, 2, 78, 79, 7, 4, 2, 2, 79, 150, 3, 2, 2, 2, 80,
	81, 7, 26, 2, 2, 81, 82, 7, 3, 2, 2, 82, 83, 5, 2, 2, 2, 83, 84, 7, 4,
	2, 2, 84, 150, 3, 2, 2, 2, 85, 86, 7, 27, 2, 2, 86, 87, 7, 3, 2, 2, 87,
	88, 5, 2, 2, 2, 88, 89, 7, 4, 2, 2, 89, 150, 3, 2, 2, 2, 90, 91, 7, 28,
	2, 2, 91, 92, 7, 3, 2, 2, 92, 93, 5, 2, 2, 2, 93, 94, 7, 15, 2, 2, 94,
	95, 5, 2, 2, 2, 95, 96, 7, 15, 2, 2, 96, 97, 5, 2, 2, 2, 97, 98, 7, 15,
	2, 2, 98, 99, 5, 2, 2, 2, 99, 100, 7, 4, 2, 2, 100, 150, 3, 2, 2, 2, 101,
	102, 7, 29, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 5, 2, 2, 2, 104, 105,
	7, 4, 2, 2, 105, 150, 3, 2, 2, 2, 106, 107, 7, 33, 2, 2, 107, 150, 7, 18,
	2, 2, 108, 109, 7, 30, 2, 2, 109, 110, 7, 3, 2, 2, 110, 111, 5, 2, 2, 2,
	111, 112, 7, 15, 2, 2, 112, 113, 5, 2, 2, 2, 113, 114, 7, 4, 2, 2, 114,
	150, 3, 2, 2, 2, 115, 116, 7, 31, 2, 2, 116, 117, 7, 3, 2, 2, 117, 120,
	5, 2, 2, 2, 118, 119, 7, 15, 2, 2, 119, 121, 5, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	123, 124, 3, 2, 2, 2, 124, 125, 7, 4, 2, 2, 125, 150, 3, 2, 2, 2, 126,
	127, 7, 32, 2, 2, 127, 128, 7, 3, 2, 2, 128, 131, 5, 2, 2, 2, 129, 130,
	7, 15, 2, 2, 130, 132, 5, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 136, 7, 4, 2, 2, 136, 150, 3, 2, 2, 2, 137, 150, 7, 45, 2, 2, 138,
	139, 7, 8, 2, 2, 139, 150, 7, 45, 2, 2, 140, 150, 7, 46, 2, 2, 141, 142,
	7, 8, 2, 2, 142, 150, 7, 46, 2, 2, 143, 150, 7, 50, 2, 2, 144, 150, 7,
	51, 2, 2, 145, 150, 7, 48, 2, 2, 146, 150, 7, 47, 2, 2, 147, 150, 7, 44,
	2, 2, 148, 150, 7, 49, 2, 2, 149, 4, 3, 2, 2, 2, 149, 9, 3, 2, 2, 2, 149,
	27, 3, 2, 2, 2, 149, 36, 3, 2, 2, 2, 149, 41, 3, 2, 2, 2, 149, 46, 3, 2,
	2, 2, 149, 51, 3, 2, 2, 2, 149, 56, 3, 2, 2, 2, 149, 63, 3, 2, 2, 2, 149,
	75, 3, 2, 2, 2, 149, 80, 3, 2, 2, 2, 149, 85, 3, 2, 2, 2, 149, 90, 3, 2,
	2, 2, 149, 101, 3, 2, 2, 2, 149, 106, 3, 2, 2, 2, 149, 108, 3, 2, 2, 2,
	149, 115, 3, 2, 2, 2, 149, 126, 3, 2, 2, 2, 149, 137, 3, 2, 2, 2, 149,
	138, 3, 2, 2, 2, 149, 140, 3, 2, 2, 2, 149, 141, 3, 2, 2, 2, 149, 143,
	3, 2, 2, 2, 149, 144, 3, 2, 2, 2, 149, 145, 3, 2, 2, 2, 149, 146, 3, 2,
	2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 218, 3, 2, 2, 2,
	151, 152, 12, 43, 2, 2, 152, 153, 7, 5, 2, 2, 153, 217, 5, 2, 2, 44, 154,
	155, 12, 42, 2, 2, 155, 156, 7, 6, 2, 2, 156, 217, 5, 2, 2, 43, 157, 158,
	12, 41, 2, 2, 158, 159, 7, 7, 2, 2, 159, 217, 5, 2, 2, 42, 160, 161, 12,
	40, 2, 2, 161, 162, 7, 8, 2, 2, 162, 217, 5, 2, 2, 41, 163, 164, 12, 39,
	2, 2, 164, 165, 7, 9, 2, 2, 165, 217, 5, 2, 2, 40, 166, 167, 12, 38, 2,
	2, 167, 168, 7, 10, 2, 2, 168, 217, 5, 2, 2, 39, 169, 170, 12, 37, 2, 2,
	170, 171, 7, 11, 2, 2, 171, 217, 5, 2, 2, 38, 172, 173, 12, 36, 2, 2, 173,
	174, 7, 12, 2, 2, 174, 217, 5, 2, 2, 37, 175, 176, 12, 35, 2, 2, 176, 177,
	7, 13, 2, 2, 177, 217, 5, 2, 2, 36, 178, 179, 12, 34, 2, 2, 179, 180, 7,
	14, 2, 2, 180, 217, 5, 2, 2, 35, 181, 182, 12, 31, 2, 2, 182, 183, 9, 2,
	2, 2, 183, 217, 5, 2, 2, 32, 184, 185, 12, 30, 2, 2, 185, 186, 9, 3, 2,
	2, 186, 217, 5, 2, 2, 31, 187, 188, 12, 33, 2, 2, 188, 189, 7, 35, 2, 2,
	189, 198, 7, 3, 2, 2, 190, 195, 5, 2, 2, 2, 191, 192, 7, 15, 2, 2, 192,
	194, 5, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193,
	3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2,
	2, 2, 198, 190, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2,
	200, 217, 7, 4, 2, 2, 201, 202, 12, 32, 2, 2, 202, 203, 7, 36, 2, 2, 203,
	204, 7, 35, 2, 2, 204, 213, 7, 3, 2, 2, 205, 210, 5, 2, 2, 2, 206, 207,
	7, 15, 2, 2, 207, 209, 5, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 212, 3, 2,
	2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2,
	212, 210, 3, 2, 2, 2, 213, 205, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	215, 3, 2, 2, 2, 215, 217, 7, 4, 2, 2, 216, 151, 3, 2, 2, 2, 216, 154,
	3, 2, 2, 2, 216, 157, 3, 2, 2, 2, 216, 160, 3, 2, 2, 2, 216, 163, 3, 2,
	2, 2, 216, 166, 3, 2, 2, 2, 216, 169, 3, 2, 2, 2, 216, 172, 3, 2, 2, 2,
	216, 175, 3, 2, 2, 2, 216, 178, 3, 2, 2, 2, 216, 181, 3, 2, 2, 2, 216,
	184, 3, 2, 2, 2, 216, 187, 3, 2, 2, 2, 216, 201, 3, 2, 2, 2, 217, 220,
	3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 3, 3, 2, 2,
	2, 220, 218, 3, 2, 2, 2, 13, 20, 70, 122, 133, 149, 195, 198, 210, 213,
	216, 218,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'>='", "'<'",
	"'<='", "'!='", "','", "'&&'", "'||'", "'()'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Abs",
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "Cos", "Cosh", "Distance",
	"Exp", "Pow", "Min", "Max", "Null", "Iif", "In", "Not", "And", "Or", "If",
	"Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal", "Date",
	"Datetime", "Field", "SingleQuoteString", "DoubleQuoteString", "Whitespace",
}

var ruleNames = []string{
	"expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AlteryxFormulasParser struct {
	*antlr.BaseParser
}

func NewAlteryxFormulasParser(input antlr.TokenStream) *AlteryxFormulasParser {
	this := new(AlteryxFormulasParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AlteryxFormulas.g4"

	return this
}

// AlteryxFormulasParser tokens.
const (
	AlteryxFormulasParserEOF               = antlr.TokenEOF
	AlteryxFormulasParserT__0              = 1
	AlteryxFormulasParserT__1              = 2
	AlteryxFormulasParserT__2              = 3
	AlteryxFormulasParserT__3              = 4
	AlteryxFormulasParserT__4              = 5
	AlteryxFormulasParserT__5              = 6
	AlteryxFormulasParserT__6              = 7
	AlteryxFormulasParserT__7              = 8
	AlteryxFormulasParserT__8              = 9
	AlteryxFormulasParserT__9              = 10
	AlteryxFormulasParserT__10             = 11
	AlteryxFormulasParserT__11             = 12
	AlteryxFormulasParserT__12             = 13
	AlteryxFormulasParserT__13             = 14
	AlteryxFormulasParserT__14             = 15
	AlteryxFormulasParserT__15             = 16
	AlteryxFormulasParserAbs               = 17
	AlteryxFormulasParserAcos              = 18
	AlteryxFormulasParserAsin              = 19
	AlteryxFormulasParserAtan              = 20
	AlteryxFormulasParserAtan2             = 21
	AlteryxFormulasParserAverage           = 22
	AlteryxFormulasParserCeil              = 23
	AlteryxFormulasParserCos               = 24
	AlteryxFormulasParserCosh              = 25
	AlteryxFormulasParserDistance          = 26
	AlteryxFormulasParserExp               = 27
	AlteryxFormulasParserPow               = 28
	AlteryxFormulasParserMin               = 29
	AlteryxFormulasParserMax               = 30
	AlteryxFormulasParserNull              = 31
	AlteryxFormulasParserIif               = 32
	AlteryxFormulasParserIn                = 33
	AlteryxFormulasParserNot               = 34
	AlteryxFormulasParserAnd               = 35
	AlteryxFormulasParserOr                = 36
	AlteryxFormulasParserIf                = 37
	AlteryxFormulasParserThen              = 38
	AlteryxFormulasParserElse              = 39
	AlteryxFormulasParserElseif            = 40
	AlteryxFormulasParserEndif             = 41
	AlteryxFormulasParserBool              = 42
	AlteryxFormulasParserInteger           = 43
	AlteryxFormulasParserDecimal           = 44
	AlteryxFormulasParserDate              = 45
	AlteryxFormulasParserDatetime          = 46
	AlteryxFormulasParserField             = 47
	AlteryxFormulasParserSingleQuoteString = 48
	AlteryxFormulasParserDoubleQuoteString = 49
	AlteryxFormulasParserWhitespace        = 50
)

// AlteryxFormulasParserRULE_expr is the AlteryxFormulasParser rule.
const AlteryxFormulasParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MaxFuncContext struct {
	*ExprContext
}

func NewMaxFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxFuncContext {
	var p = new(MaxFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MaxFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxFuncContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *MaxFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MaxFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MaxFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMaxFunc(s)
	}
}

func (s *MaxFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMaxFunc(s)
	}
}

type CeilFuncContext struct {
	*ExprContext
}

func NewCeilFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilFuncContext {
	var p = new(CeilFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CeilFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilFuncContext) Ceil() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCeil, 0)
}

func (s *CeilFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CeilFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCeilFunc(s)
	}
}

func (s *CeilFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCeilFunc(s)
	}
}

type CosFuncContext struct {
	*ExprContext
}

func NewCosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CosFuncContext {
	var p = new(CosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosFuncContext) Cos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCos, 0)
}

func (s *CosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCosFunc(s)
	}
}

func (s *CosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCosFunc(s)
	}
}

type DistanceFuncContext struct {
	*ExprContext
}

func NewDistanceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistanceFuncContext {
	var p = new(DistanceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DistanceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceFuncContext) Distance() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDistance, 0)
}

func (s *DistanceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DistanceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DistanceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDistanceFunc(s)
	}
}

func (s *DistanceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDistanceFunc(s)
	}
}

type BoolLiteralContext struct {
	*ExprContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) Bool() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserBool, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

type AcosFuncContext struct {
	*ExprContext
}

func NewAcosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AcosFuncContext {
	var p = new(AcosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AcosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcosFuncContext) Acos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAcos, 0)
}

func (s *AcosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AcosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAcosFunc(s)
	}
}

func (s *AcosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAcosFunc(s)
	}
}

type MinFuncContext struct {
	*ExprContext
}

func NewMinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinFuncContext {
	var p = new(MinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinFuncContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *MinFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MinFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMinFunc(s)
	}
}

func (s *MinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMinFunc(s)
	}
}

type AndContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetLeft() IExprContext { return s.left }

func (s *AndContext) GetRight() IExprContext { return s.right }

func (s *AndContext) SetLeft(v IExprContext) { s.left = v }

func (s *AndContext) SetRight(v IExprContext) { s.right = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) And() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAnd, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAnd(s)
	}
}

type LessThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessThanContext) GetLeft() IExprContext { return s.left }

func (s *LessThanContext) GetRight() IExprContext { return s.right }

func (s *LessThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessThanContext) SetRight(v IExprContext) { s.right = v }

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessThan(s)
	}
}

type DivideContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivideContext) GetLeft() IExprContext { return s.left }

func (s *DivideContext) GetRight() IExprContext { return s.right }

func (s *DivideContext) SetLeft(v IExprContext) { s.left = v }

func (s *DivideContext) SetRight(v IExprContext) { s.right = v }

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivideContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDivide(s)
	}
}

type NotInContext struct {
	*ExprContext
}

func NewNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInContext {
	var p = new(NotInContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotInContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *NotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *NotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotIn(s)
	}
}

func (s *NotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotIn(s)
	}
}

type ExprIfContext struct {
	*ExprContext
	ifStmnt   IExprContext
	thenStmnt IExprContext
	elseStmnt IExprContext
}

func NewExprIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIfContext {
	var p = new(ExprIfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprIfContext) GetIfStmnt() IExprContext { return s.ifStmnt }

func (s *ExprIfContext) GetThenStmnt() IExprContext { return s.thenStmnt }

func (s *ExprIfContext) GetElseStmnt() IExprContext { return s.elseStmnt }

func (s *ExprIfContext) SetIfStmnt(v IExprContext) { s.ifStmnt = v }

func (s *ExprIfContext) SetThenStmnt(v IExprContext) { s.thenStmnt = v }

func (s *ExprIfContext) SetElseStmnt(v IExprContext) { s.elseStmnt = v }

func (s *ExprIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *ExprIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *ExprIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *ExprIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *ExprIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *ExprIfContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprIfContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *ExprIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *ExprIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprIf(s)
	}
}

func (s *ExprIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprIf(s)
	}
}

type MultiplyContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultiplyContext) GetLeft() IExprContext { return s.left }

func (s *MultiplyContext) GetRight() IExprContext { return s.right }

func (s *MultiplyContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultiplyContext) SetRight(v IExprContext) { s.right = v }

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultiplyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMultiply(s)
	}
}

func (s *MultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMultiply(s)
	}
}

type AtanFuncContext struct {
	*ExprContext
}

func NewAtanFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtanFuncContext {
	var p = new(AtanFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtanFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtanFuncContext) Atan() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan, 0)
}

func (s *AtanFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtanFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtanFunc(s)
	}
}

func (s *AtanFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtanFunc(s)
	}
}

type ExprFieldContext struct {
	*ExprContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprField(s)
	}
}

type GreaterThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterThanContext) GetLeft() IExprContext { return s.left }

func (s *GreaterThanContext) GetRight() IExprContext { return s.right }

func (s *GreaterThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterThanContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

type AsinFuncContext struct {
	*ExprContext
}

func NewAsinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsinFuncContext {
	var p = new(AsinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AsinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsinFuncContext) Asin() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAsin, 0)
}

func (s *AsinFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAsinFunc(s)
	}
}

func (s *AsinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAsinFunc(s)
	}
}

type AddContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddContext) GetLeft() IExprContext { return s.left }

func (s *AddContext) GetRight() IExprContext { return s.right }

func (s *AddContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddContext) SetRight(v IExprContext) { s.right = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAdd(s)
	}
}

type OrContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetLeft() IExprContext { return s.left }

func (s *OrContext) GetRight() IExprContext { return s.right }

func (s *OrContext) SetLeft(v IExprContext) { s.left = v }

func (s *OrContext) SetRight(v IExprContext) { s.right = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) Or() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserOr, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitOr(s)
	}
}

type CoshFuncContext struct {
	*ExprContext
}

func NewCoshFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoshFuncContext {
	var p = new(CoshFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CoshFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoshFuncContext) Cosh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCosh, 0)
}

func (s *CoshFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CoshFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCoshFunc(s)
	}
}

func (s *CoshFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCoshFunc(s)
	}
}

type InContext struct {
	*ExprContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIn(s)
	}
}

type ExpFuncContext struct {
	*ExprContext
}

func NewExpFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFuncContext {
	var p = new(ExpFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFuncContext) Exp() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserExp, 0)
}

func (s *ExpFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExpFunc(s)
	}
}

func (s *ExpFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExpFunc(s)
	}
}

type SubtractContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewSubtractContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractContext {
	var p = new(SubtractContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubtractContext) GetLeft() IExprContext { return s.left }

func (s *SubtractContext) GetRight() IExprContext { return s.right }

func (s *SubtractContext) SetLeft(v IExprContext) { s.left = v }

func (s *SubtractContext) SetRight(v IExprContext) { s.right = v }

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubtractContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubtractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSubtract(s)
	}
}

func (s *SubtractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSubtract(s)
	}
}

type NotEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqualContext) GetLeft() IExprContext { return s.left }

func (s *NotEqualContext) GetRight() IExprContext { return s.right }

func (s *NotEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *NotEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotEqual(s)
	}
}

func (s *NotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotEqual(s)
	}
}

type ParenthesisContext struct {
	*ExprContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type PowFuncContext struct {
	*ExprContext
}

func NewPowFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowFuncContext {
	var p = new(PowFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowFuncContext) Pow() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPow, 0)
}

func (s *PowFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPowFunc(s)
	}
}

func (s *PowFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPowFunc(s)
	}
}

type IifFuncContext struct {
	*ExprContext
}

func NewIifFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IifFuncContext {
	var p = new(IifFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IifFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IifFuncContext) Iif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIif, 0)
}

func (s *IifFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IifFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IifFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIifFunc(s)
	}
}

func (s *IifFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIifFunc(s)
	}
}

type EqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualContext) GetLeft() IExprContext { return s.left }

func (s *EqualContext) GetRight() IExprContext { return s.right }

func (s *EqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitEqual(s)
	}
}

type NullFuncContext struct {
	*ExprContext
}

func NewNullFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullFuncContext {
	var p = new(NullFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NullFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullFuncContext) Null() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNull, 0)
}

func (s *NullFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNullFunc(s)
	}
}

func (s *NullFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNullFunc(s)
	}
}

type DatetimeLiteralContext struct {
	*ExprContext
}

func NewDatetimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatetimeLiteralContext {
	var p = new(DatetimeLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DatetimeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeLiteralContext) Datetime() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDatetime, 0)
}

func (s *DatetimeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDatetimeLiteral(s)
	}
}

func (s *DatetimeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDatetimeLiteral(s)
	}
}

type StringLiteralContext struct {
	*ExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) SingleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSingleQuoteString, 0)
}

func (s *StringLiteralContext) DoubleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDoubleQuoteString, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type AverageFuncContext struct {
	*ExprContext
}

func NewAverageFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AverageFuncContext {
	var p = new(AverageFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AverageFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AverageFuncContext) Average() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAverage, 0)
}

func (s *AverageFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AverageFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AverageFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAverageFunc(s)
	}
}

func (s *AverageFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAverageFunc(s)
	}
}

type DateLiteralContext struct {
	*ExprContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DateLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLiteralContext) Date() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDate, 0)
}

func (s *DateLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateLiteral(s)
	}
}

func (s *DateLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateLiteral(s)
	}
}

type GreaterEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterEqualContext {
	var p = new(GreaterEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterEqualContext) GetLeft() IExprContext { return s.left }

func (s *GreaterEqualContext) GetRight() IExprContext { return s.right }

func (s *GreaterEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterEqual(s)
	}
}

func (s *GreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterEqual(s)
	}
}

type AbsFuncContext struct {
	*ExprContext
}

func NewAbsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbsFuncContext {
	var p = new(AbsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AbsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsFuncContext) Abs() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAbs, 0)
}

func (s *AbsFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAbsFunc(s)
	}
}

func (s *AbsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAbsFunc(s)
	}
}

type LessEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessEqualContext {
	var p = new(LessEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessEqualContext) GetLeft() IExprContext { return s.left }

func (s *LessEqualContext) GetRight() IExprContext { return s.right }

func (s *LessEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *LessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessEqual(s)
	}
}

func (s *LessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessEqual(s)
	}
}

type Atan2FuncContext struct {
	*ExprContext
}

func NewAtan2FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atan2FuncContext {
	var p = new(Atan2FuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Atan2FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atan2FuncContext) Atan2() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan2, 0)
}

func (s *Atan2FuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Atan2FuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Atan2FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtan2Func(s)
	}
}

func (s *Atan2FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtan2Func(s)
	}
}

type NumberLiteralContext struct {
	*ExprContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) Integer() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserInteger, 0)
}

func (s *NumberLiteralContext) Decimal() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDecimal, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *AlteryxFormulasParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *AlteryxFormulasParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, AlteryxFormulasParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(4)
			p.expr(0)
		}
		{
			p.SetState(5)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewExprIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(7)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(8)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).ifStmnt = _x
		}
		{
			p.SetState(9)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(10)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).thenStmnt = _x
		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserElseif {
			{
				p.SetState(11)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(12)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).ifStmnt = _x
			}
			{
				p.SetState(13)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(14)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).thenStmnt = _x
			}

			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(21)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(22)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).elseStmnt = _x
		}
		{
			p.SetState(23)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 3:
		localctx = NewIifFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(AlteryxFormulasParserIif)
		}
		{
			p.SetState(26)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(27)
			p.expr(0)
		}
		{
			p.SetState(28)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(29)
			p.expr(0)
		}
		{
			p.SetState(30)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(31)
			p.expr(0)
		}
		{
			p.SetState(32)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 4:
		localctx = NewAbsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(AlteryxFormulasParserAbs)
		}
		{
			p.SetState(35)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(36)
			p.expr(0)
		}
		{
			p.SetState(37)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 5:
		localctx = NewAcosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			p.Match(AlteryxFormulasParserAcos)
		}
		{
			p.SetState(40)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(41)
			p.expr(0)
		}
		{
			p.SetState(42)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 6:
		localctx = NewAsinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(AlteryxFormulasParserAsin)
		}
		{
			p.SetState(45)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(46)
			p.expr(0)
		}
		{
			p.SetState(47)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 7:
		localctx = NewAtanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(AlteryxFormulasParserAtan)
		}
		{
			p.SetState(50)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(51)
			p.expr(0)
		}
		{
			p.SetState(52)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 8:
		localctx = NewAtan2FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(AlteryxFormulasParserAtan2)
		}
		{
			p.SetState(55)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(56)
			p.expr(0)
		}
		{
			p.SetState(57)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		{
			p.SetState(59)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 9:
		localctx = NewAverageFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(AlteryxFormulasParserAverage)
		}
		{
			p.SetState(62)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(63)
			p.expr(0)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(64)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(65)
				p.expr(0)
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 10:
		localctx = NewCeilFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(AlteryxFormulasParserCeil)
		}
		{
			p.SetState(74)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(75)
			p.expr(0)
		}
		{
			p.SetState(76)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 11:
		localctx = NewCosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(AlteryxFormulasParserCos)
		}
		{
			p.SetState(79)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(80)
			p.expr(0)
		}
		{
			p.SetState(81)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 12:
		localctx = NewCoshFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(AlteryxFormulasParserCosh)
		}
		{
			p.SetState(84)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(85)
			p.expr(0)
		}
		{
			p.SetState(86)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 13:
		localctx = NewDistanceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(AlteryxFormulasParserDistance)
		}
		{
			p.SetState(89)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(90)
			p.expr(0)
		}
		{
			p.SetState(91)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(92)
			p.expr(0)
		}
		{
			p.SetState(93)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(96)
			p.expr(0)
		}
		{
			p.SetState(97)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 14:
		localctx = NewExpFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(AlteryxFormulasParserExp)
		}
		{
			p.SetState(100)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(101)
			p.expr(0)
		}
		{
			p.SetState(102)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 15:
		localctx = NewNullFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.Match(AlteryxFormulasParserNull)
		}
		{
			p.SetState(105)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 16:
		localctx = NewPowFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(106)
			p.Match(AlteryxFormulasParserPow)
		}
		{
			p.SetState(107)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(108)
			p.expr(0)
		}
		{
			p.SetState(109)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(110)
			p.expr(0)
		}
		{
			p.SetState(111)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 17:
		localctx = NewMinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(114)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(115)
			p.expr(0)
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(116)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(117)
				p.expr(0)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 18:
		localctx = NewMaxFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(124)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(125)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(126)
			p.expr(0)
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(127)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(128)
				p.expr(0)
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(133)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 19:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 20:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(137)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 21:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 22:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(140)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 23:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(AlteryxFormulasParserSingleQuoteString)
		}

	case 24:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(142)
			p.Match(AlteryxFormulasParserDoubleQuoteString)
		}

	case 25:
		localctx = NewDatetimeLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(AlteryxFormulasParserDatetime)
		}

	case 26:
		localctx = NewDateLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(AlteryxFormulasParserDate)
		}

	case 27:
		localctx = NewBoolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.Match(AlteryxFormulasParserBool)
		}

	case 28:
		localctx = NewExprFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(146)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplyContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 41)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 41)", ""))
				}
				{
					p.SetState(150)
					p.Match(AlteryxFormulasParserT__2)
				}
				{
					p.SetState(151)

					var _x = p.expr(42)

					localctx.(*MultiplyContext).right = _x
				}

			case 2:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivideContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 40)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 40)", ""))
				}
				{
					p.SetState(153)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(154)

					var _x = p.expr(41)

					localctx.(*DivideContext).right = _x
				}

			case 3:
				localctx = NewAddContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 39)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 39)", ""))
				}
				{
					p.SetState(156)
					p.Match(AlteryxFormulasParserT__4)
				}
				{
					p.SetState(157)

					var _x = p.expr(40)

					localctx.(*AddContext).right = _x
				}

			case 4:
				localctx = NewSubtractContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubtractContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 38)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 38)", ""))
				}
				{
					p.SetState(159)
					p.Match(AlteryxFormulasParserT__5)
				}
				{
					p.SetState(160)

					var _x = p.expr(39)

					localctx.(*SubtractContext).right = _x
				}

			case 5:
				localctx = NewEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 37)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 37)", ""))
				}
				{
					p.SetState(162)
					p.Match(AlteryxFormulasParserT__6)
				}
				{
					p.SetState(163)

					var _x = p.expr(38)

					localctx.(*EqualContext).right = _x
				}

			case 6:
				localctx = NewGreaterThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 36)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 36)", ""))
				}
				{
					p.SetState(165)
					p.Match(AlteryxFormulasParserT__7)
				}
				{
					p.SetState(166)

					var _x = p.expr(37)

					localctx.(*GreaterThanContext).right = _x
				}

			case 7:
				localctx = NewGreaterEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 35)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 35)", ""))
				}
				{
					p.SetState(168)
					p.Match(AlteryxFormulasParserT__8)
				}
				{
					p.SetState(169)

					var _x = p.expr(36)

					localctx.(*GreaterEqualContext).right = _x
				}

			case 8:
				localctx = NewLessThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
				}
				{
					p.SetState(171)
					p.Match(AlteryxFormulasParserT__9)
				}
				{
					p.SetState(172)

					var _x = p.expr(35)

					localctx.(*LessThanContext).right = _x
				}

			case 9:
				localctx = NewLessEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
				}
				{
					p.SetState(174)
					p.Match(AlteryxFormulasParserT__10)
				}
				{
					p.SetState(175)

					var _x = p.expr(34)

					localctx.(*LessEqualContext).right = _x
				}

			case 10:
				localctx = NewNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
				}
				{
					p.SetState(177)
					p.Match(AlteryxFormulasParserT__11)
				}
				{
					p.SetState(178)

					var _x = p.expr(33)

					localctx.(*NotEqualContext).right = _x
				}

			case 11:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
				}
				{
					p.SetState(180)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__13 || _la == AlteryxFormulasParserAnd) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(181)

					var _x = p.expr(30)

					localctx.(*AndContext).right = _x
				}

			case 12:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
				}
				{
					p.SetState(183)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__14 || _la == AlteryxFormulasParserOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(184)

					var _x = p.expr(29)

					localctx.(*OrContext).right = _x
				}

			case 13:
				localctx = NewInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
				}
				{
					p.SetState(186)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(187)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(196)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserExp)|(1<<AlteryxFormulasParserPow)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserNull))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIf-32))|(1<<(AlteryxFormulasParserBool-32))|(1<<(AlteryxFormulasParserInteger-32))|(1<<(AlteryxFormulasParserDecimal-32))|(1<<(AlteryxFormulasParserDate-32))|(1<<(AlteryxFormulasParserDatetime-32))|(1<<(AlteryxFormulasParserField-32))|(1<<(AlteryxFormulasParserSingleQuoteString-32))|(1<<(AlteryxFormulasParserDoubleQuoteString-32)))) != 0) {
					{
						p.SetState(188)
						p.expr(0)
					}
					p.SetState(193)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(189)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(190)
							p.expr(0)
						}

						p.SetState(195)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(198)
					p.Match(AlteryxFormulasParserT__1)
				}

			case 14:
				localctx = NewNotInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
				}
				{
					p.SetState(200)
					p.Match(AlteryxFormulasParserNot)
				}
				{
					p.SetState(201)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(202)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(211)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserExp)|(1<<AlteryxFormulasParserPow)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserNull))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIf-32))|(1<<(AlteryxFormulasParserBool-32))|(1<<(AlteryxFormulasParserInteger-32))|(1<<(AlteryxFormulasParserDecimal-32))|(1<<(AlteryxFormulasParserDate-32))|(1<<(AlteryxFormulasParserDatetime-32))|(1<<(AlteryxFormulasParserField-32))|(1<<(AlteryxFormulasParserSingleQuoteString-32))|(1<<(AlteryxFormulasParserDoubleQuoteString-32)))) != 0) {
					{
						p.SetState(203)
						p.expr(0)
					}
					p.SetState(208)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(204)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(205)
							p.expr(0)
						}

						p.SetState(210)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(213)
					p.Match(AlteryxFormulasParserT__1)
				}

			}

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

func (p *AlteryxFormulasParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlteryxFormulasParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 41)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 40)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 39)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 38)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 37)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 36)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 35)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 33)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 30)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
