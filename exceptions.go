package alteryx_formulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

func NotifyError(notifier ErrorNotifier, msg string) {
	exception := formulasException{
		Message:        ``,
		InputStream:    notifier.GetStart().GetInputStream(),
		OffendingToken: notifier.GetStart(),
	}
	notifier.GetParser().NotifyErrorListeners(msg, notifier.GetStart(), exception)
}

type ErrorNotifier interface {
	GetParser() antlr.Parser
	antlr.ParserRuleContext
}

type formulasException struct {
	Message        string
	InputStream    antlr.IntStream
	OffendingToken antlr.Token
}

func (e formulasException) GetOffendingToken() antlr.Token {
	return e.OffendingToken
}

func (e formulasException) GetMessage() string {
	return e.Message
}

func (e formulasException) GetInputStream() antlr.IntStream {
	return e.InputStream
}
