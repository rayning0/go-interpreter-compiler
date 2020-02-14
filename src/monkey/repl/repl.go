package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

// Read from the input source until encountering a \n,
// take the just read line and pass it to an instance of our
// lexer and finally print all the tokens the lexer gives us
// until we encounter EOF.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok) // print struct w/ field names
		}
	}
}

// Sample REPL output, showing all tokens of input line:

// >> let add = fn(x, y) { x + y; };

// {Type:LET Literal:let}
// {Type:IDENT Literal:add}
// {Type:= Literal:=}
// {Type:FUNCTION Literal:fn}
// {Type:( Literal:(}
// {Type:IDENT Literal:x}
// {Type:, Literal:,}
// {Type:IDENT Literal:y}
// {Type:) Literal:)}
// {Type:{ Literal:{}
// {Type:IDENT Literal:x}
// {Type:+ Literal:+}
// {Type:IDENT Literal:y}
// {Type:; Literal:;}
// {Type:} Literal:}}
// {Type:; Literal:;}
