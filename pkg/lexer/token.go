package lexer

// Token represents a lexical token.
type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	IDENT // identifier

	BRACKETLEFT        // (
	BRACKETRIGHT       // )
	CURLYBRACKETLEFT   // {
	CURLYBRACKETRIGHT  // }
	SQUAREBRACKETLEFT  // [
	SQUAREBRACKETRIGHT // ]
	BACKSLASH          // \
	DOT                // .
	COMMA              // ,
	DOUBLEDOT          // :
	PLUS               // +
	MINUS              // -
	ASTERISK           //*
	CIRCUMFLEX         // ^
	DOLAR              // $
	BIGGER             // >
	HASH               // #

	VALUE
	START
)

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

// MiscCharMap is a map from the rune to the Token
var MiscCharMap = map[rune]Token{
	'(':  BRACKETLEFT,
	')':  BRACKETRIGHT,
	'{':  CURLYBRACKETLEFT,
	'}':  CURLYBRACKETRIGHT,
	'[':  SQUAREBRACKETLEFT,
	']':  SQUAREBRACKETRIGHT,
	'\\': BACKSLASH,
	'.':  DOT,
	',':  COMMA,
	':':  DOUBLEDOT,
	'+':  PLUS,
	'-':  MINUS,
	'*':  ASTERISK,
	'^':  CIRCUMFLEX,
	'$':  DOLAR,
	'>':  BIGGER,
	'#':  HASH,
}

// KeyWordMap is a map from the string to the Token
var KeyWordMap = map[string]Token{
	"Value": VALUE,
	"Start": START,
}
