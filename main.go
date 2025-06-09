package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"
    "bufio"
    "os"
    "errors"
    "unsafe"
)

func boolToInt(b bool) int {
    return int(*(*byte)(unsafe.Pointer(&b)))
}

type Token struct {
    t string
    value string
}

const (
    LPARAN = "Open Parenthesis"
    RPARAN = "Close Parenthesis"
    SYMBOL = "Symbol"
    INTEGER = "Integer"
    FLOAT = "Float"
    STRING = "String"
    EOF = "EOF"
)

func readSymbol(symbol string) Token {

    if _, err := strconv.Atoi(symbol); err == nil {
        return Token{ INTEGER, symbol }
    } else if _, err := strconv.ParseFloat(symbol, 32); err == nil {
        return Token{ FLOAT, symbol}
    }

    return Token{ SYMBOL, symbol}
}

func lexLine(line string) ([]Token, error) {
    var tokens = make([]Token, 0)
    var buffer strings.Builder

    if line[0] == ')' {
        return nil, errors.New("nasa-pini") // unmatched parenthesis
    }

    for _, r := range line {
        switch r {
        case '(':
            if buffer.Len() > 0 {
                tokens = append(tokens, readSymbol(buffer.String()))
                buffer.Reset()
            }

            tokens = append(tokens, Token{ LPARAN, "(" } )
            continue

        case ')':
            if buffer.Len() > 0 {
                tokens = append(tokens, readSymbol(buffer.String()))
                buffer.Reset()
            }

            tokens = append(tokens, Token{ RPARAN, ")" })
            continue;

        case ' ': 
            if buffer.String() != "" {
                tokens = append(tokens, readSymbol(buffer.String()))
                buffer.Reset()
            }
            continue
        }

        buffer.WriteRune(r)
    }
    return tokens, nil
}

type Parser struct {
    head int 
    tokens []Token 
}

func makeParser(t []Token) *Parser {
    return &Parser { head: 0, tokens: t} 
}

func (p *Parser) GetToken() Token {
    if p.head + 1 > len(p.tokens) {
        return Token{ EOF, "" }
    }
    tok := p.tokens[p.head]
    p.head += 1
    return tok
}

const (
    EXPR_KEWORD = iota
    EXPR_CONSTANT
    EXPR_SYMBOL
    EXPR_VALUE
)

type Expr struct {
    value_kind int
    value string
    left *Expr
    right *Expr
}

func (p *Parser) GetExpr() *Expr {
    var expr Expr

    tok := p.GetToken()
    if tok.t == EOF {
        return nil
    }

    for tok.value == " " {
        tok = p.GetToken()
    }

    switch tok.t {
    case RPARAN:
        tmp := p.GetExpr()
        if tmp != nil {
            expr = *tmp
        }
    case LPARAN:
        expr = *p.GetExpr()
    case SYMBOL:
        switch tok.value {
        default:
            expr.value_kind = EXPR_SYMBOL
            expr.value = tok.value
            expr.left = p.GetExpr()
            expr.right = p.GetExpr()
        }
    case INTEGER:
        expr.value_kind = EXPR_CONSTANT
        expr.value = tok.value
        expr.left = nil
        expr.right = nil
    }

    return &expr
}

func eval(expr *Expr) (int, error) {
    var ret int = 0

    if expr == nil {
        return 0, nil
    }

    switch expr.value_kind {
    case EXPR_SYMBOL:
        switch expr.value {
        case "+":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            y, err := eval(expr.right)
            if err != nil {
                return 0, err
            }

            res := x + y
            return res, nil
        case "-":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            y, err := eval(expr.right)
            if err != nil {
                return 0, err
            }

            res := x - y
            return res, nil
        case "*":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            y, err := eval(expr.right)
            if err != nil {
                return 0, err
            }

            res := x * y
            return res, nil
        case "/":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            y, err := eval(expr.right)
            if err != nil {
                return 0, err
            }

            res := x / y
            return res, nil
        case "=":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            y, err := eval(expr.right)
            if err != nil {
                return 0, err
            }

            res := boolToInt(x == y)
            return res, nil

        case "la":
            x, err := eval(expr.left)
            if err != nil {
                return 0, err
            }

            if x != 0 {
                return eval(expr.right)
            } else {
                return 0, nil
            }
        default:
            fmt.Printf("nasa ijo: << %s >>\n", expr.value)
            return 0, errors.New("nasa-ijo")
        }
    case EXPR_CONSTANT:
        val, err := strconv.Atoi(expr.value)
        if err != nil {
            return 0, errors.New("nasa-nanpa")
        }
        return val, nil
    default:
        fmt.Printf("nasa nimi: << %s >>\n", expr.value)
        return 0, errors.New("nasa-nimi")
    }

    return ret, nil
}

func main() {
    quit := false
    reader := bufio.NewReader(os.Stdin)

    for !quit {
        fmt.Print("|: ")
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        line = strings.Join(strings.Fields(strings.TrimSpace(line)), " ")

        if len(line) == 0 {
            fmt.Printf("ike: poki nimi li ala\n")
            continue
        }

        tokens, err := lexLine(line) 
        if err != nil {
            fmt.Printf("ike: %v\n", err)
            continue
        }

        parser := makeParser(tokens)
        expr := parser.GetExpr()
        res, err := eval(expr)

        if err != nil {
            fmt.Printf("ike: %v\n", err)
            continue
        }

        fmt.Printf("pona: %d\n", res)
    }
}
