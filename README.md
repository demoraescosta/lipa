# toki Lipa

#### e toki tawa ilo sona

Lipa is a toki pona variant of the [Lisp](https://en.wikipedia.org/wiki/Lisp_(programming_language)) programming language written in Go.

Currently the program consists of a simple repl that can evaluate addition and simple conditionals and functions,
this of course won't stay this way for long as I do plan to improve on the interpreter further.

## Keywords
- '+' = sin : add
- '-' = weka : subtract
- '\*' = li mute : multiply
- '/' = tu : divide
- 'la' = if
- 'lon' = defun / define

## Example usage

```
$ go run main.go

|: (+ (+ 1 2) 2)
pona: 5
|: (la (= 1 1) 3)
pona: 3
|: (lon luka 5)
pona: 0
|: (luka)
pona: 5
```

