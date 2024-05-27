package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT = "INT" // 12346

    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    SLASH = "/"
    ASTERISK = "*"
    BANG = "!"

    LT = "<"
    GT = ">"

    EQ = "=="
    NOT_EQ = "!="

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    RETURN = "RETURN"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "return": RETURN,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENT
}
