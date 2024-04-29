package parser

type Lexer interface {
	Validate([]byte) error
}

type lexer struct{}

func NewLexer() Lexer {
	return &lexer{}
}

func (l *lexer) Validate(data []byte) error {
	// An empty JSON is not a valid JSON
	if len(data) == 0 {
		return ErrEmptyData
	}

	return nil
}
