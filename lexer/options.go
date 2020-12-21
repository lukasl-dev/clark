package lexer

type Options struct {
	Prefixes         []string `json:"prefixes,omitempty"`
	PrefixIgnoreCase bool     `json:"prefixIgnoreCase,omitempty"`
	Labels           []string `json:"labels,omitempty"`
	LabelIgnoreCase  bool     `json:"labelIgnoreCase,omitempty"`
}
