package tokenizer

import (
	"fmt"
	"unicode"
)

type Tokenizer struct {
	source    []rune
	tokens    []Token
	peek      int //index of current string reader
	sourceLen int //cache value of source lenght to optimize computing
}

func (this *Tokenizer) peekRune() rune {
	return this.source[this.peek]
}

func (this *Tokenizer) startToken(tokenType TokenType) Token {
	token := Token{Type: tokenType, startIndex: this.peek}
	this.peek++
	return token
}

func (this *Tokenizer) isEOL() bool {
	return this.peek >= this.sourceLen
}

func (this *Tokenizer) isWordStart() bool {
	return unicode.IsLetter(this.peekRune())
}

// isWordTail check current symbol is it belongs to letter characters
// or it other symbols which could be in the work "_", "-"
func (this *Tokenizer) isWordTail() bool {
	return unicode.IsLetter(this.peekRune()) ||
		this.peekRune() == '_' ||
		this.peekRune() == '-'
}

func (this *Tokenizer) isNumStart() bool {
	return unicode.IsDigit(this.peekRune())
}

func (this *Tokenizer) isNumTail() bool {
	return unicode.IsDigit(this.peekRune()) ||
		this.peekRune() == '.' ||
		this.peekRune() == 'e' ||
		this.peekRune() == 'E'
}

func (this *Tokenizer) isNewline() bool {
	return unicode.IsSpace(this.peekRune()) && this.peekRune() == '\n'
}

func (this *Tokenizer) isSpace() bool {
	return unicode.IsSpace(this.peekRune()) && this.peekRune() != '\n'
}

func (this *Tokenizer) String() string {
	var result string
	for i := 0; i < len(this.tokens); i++ {
		token := this.tokens[i]
		tokenEndIndex := token.startIndex + 1
		if i < len(this.tokens)-1 {
			nextToken := this.tokens[i+1]
			tokenEndIndex = nextToken.startIndex
		}
		tokenName, exits := TokenTypeNames[token.Type]
		if !exits {
			tokenName = "UNKNOWN"
		}
		val := string(this.source[token.startIndex:tokenEndIndex])
		result += fmt.Sprintf("%3d %s\t%v\n", token.startIndex, tokenName, val)
	}
	return result
}

func New(input string) Tokenizer {
	var result Tokenizer
	result.source = []rune(input)
	result.sourceLen = len(result.source)

	var token Token
	for !result.isEOL() {
		switch {
		case result.isWordStart():
			token = result.startToken(TOKEN_WORD)
			for !result.isEOL() && result.isWordTail() {
				result.peek++
			}
		case result.isNumStart():
			token = result.startToken(TOKEN_NUMBER)
			for !result.isEOL() && result.isNumTail() {
				result.peek++
			}
		case result.isSpace():
			token = result.startToken(TOKEN_SPACE)
		case result.isNewline():
			token = result.startToken(TOKEN_NEWLINE)
		case result.peekRune() == ',':
			token = result.startToken(TOKEN_COMMA)
		case result.peekRune() == '.':
			token = result.startToken(TOKEN_DOT)
		case result.peekRune() == '*':
			token = result.startToken(TOKEN_ASTERISK)
		case result.peekRune() == '_':
			token = result.startToken(TOKEN_UNDERSCORE)
		case result.peekRune() == '[':
			token = result.startToken(TOKEN_SQBRACKET_BEGIN)
		case result.peekRune() == ']':
			token = result.startToken(TOKEN_SQBRACKET_END)
		case result.peekRune() == '(':
			token = result.startToken(TOKEN_PHARENTHESIS_BEGIN)
		case result.peekRune() == ')':
			token = result.startToken(TOKEN_PHARENTHESIS_END)
		case result.peekRune() == '`':
			token = result.startToken(TOKEN_BACKTICK)
		case result.peekRune() == '&':
			token = result.startToken(TOKEN_AMPERSAND)
		case result.peekRune() == '%':
			token = result.startToken(TOKEN_PROCENT)
		case result.peekRune() == '$':
			token = result.startToken(TOKEN_DOLLAR)
		case result.peekRune() == '#':
			token = result.startToken(TOKEN_SHARP)
		case result.peekRune() == '{':
			token = result.startToken(TOKEN_CBRACKET_BEGIN)
		case result.peekRune() == '}':
			token = result.startToken(TOKEN_CBRACKET_END)
		case result.peekRune() == '~':
			token = result.startToken(TOKEN_TILDA)
		case result.peekRune() == '^':
			token = result.startToken(TOKEN_CARAT)
		case result.peekRune() == '\\':
			token = result.startToken(TOKEN_BACKSLASH)
		case result.peekRune() == '/':
			token = result.startToken(TOKEN_SLASH)
		case result.peekRune() == ':':
			token = result.startToken(TOKEN_SEMICOLON)
		case result.peekRune() == '"':
			token = result.startToken(TOKEN_DOUBLE_QUOTE)
		case result.peekRune() == '\'':
			token = result.startToken(TOKEN_SINGLE_QUOTE)
		case result.peekRune() == '-':
			token = result.startToken(TOKEN_HYPHEN)
		default:
			token = result.startToken(TOKEN_ILLEGAL)
		}
		result.tokens = append(result.tokens, token)
	}

	return result
}
