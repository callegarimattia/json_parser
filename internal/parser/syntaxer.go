package parser

type Syntaxer interface {
	Validate([]byte) error
}

type syntaxer struct{}

func NewSyntaxer() Syntaxer {
	return &syntaxer{}
}

func (s *syntaxer) Validate(data []byte) error {
	return nil
}
