package parser

import "log/slog"

type Parser interface {
	Parse([]byte) (int, error)
}

var _ Parser = (*parser)(nil)

type parser struct {
	lexer    Lexer
	syntaxer Syntaxer
}

func NewParser() Parser {
	return &parser{
		lexer:    NewLexer(),
		syntaxer: NewSyntaxer(),
	}
}

func (p *parser) Parse(data []byte) (int, error) {
	code, err := errorParsing(p.lexer.Validate(data))
	if err != nil {
		return code, err
	}

	return 0, nil
}

func errorParsing(err error) (int, error) {
	switch err {
	case ErrEmptyData:
		slog.Warn("Empty data")
		return 1, nil
	default:
		return 0, err
	}
}
