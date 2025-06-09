package main

import (
    "testing"
)

func TestEvalInvalid(t *testing.T) {
    line := "(ike)"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 0

    if res != expected || err == nil {
        t.Errorf(`%s = %d, %v, want %d, error`, line, res, err, expected)
    }
}

func TestEvalAdd(t *testing.T) {
    line := "(+ 1 2)"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 3

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalNestedExpr(t *testing.T) {
    line := "(+ (+ 1 2) (+ 2 3))"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 8

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalIf(t *testing.T) {
    line := "(la (1) 1)"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 1

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalIfComplex(t *testing.T) {
    line := "(la (= 1 1) (+ 6 6))"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 12

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalToki(t *testing.T) {
    line := "(toki 32)"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 1

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalEmptyExpr(t *testing.T) {
    line := "()"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    scope := makeScope(nil)
    res, err := eval(scope, expr)

    expected := 0

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

func TestEvalFunctions(t *testing.T) {
    scope := makeScope(nil)

    line := "(lon test 32)"
    tokens, err := lexLine(line) 
    parser := makeParser(tokens)
    expr := parser.GetExpr()
    res, err := eval(scope, expr)

    expected := 0

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }

    line = "(test)"
    tokens, err = lexLine(line) 
    parser = makeParser(tokens)
    expr = parser.GetExpr()

    res, err = eval(scope, expr)

    expected = 12

    if res != expected || err != nil {
        t.Errorf(`%s = %d, %v, want %d, nil`, line, res, err, expected)
    }
}

