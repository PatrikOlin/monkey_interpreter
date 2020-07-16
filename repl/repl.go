package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/PatrikOlin/monkey_interpreter/evaluator"
	"github.com/PatrikOlin/monkey_interpreter/lexer"
	"github.com/PatrikOlin/monkey_interpreter/parser"
	"github.com/PatrikOlin/monkey_interpreter/object"
)

const PROMPT = ">> "
const TABLE_FLIPPED = `
(╯°□°）╯︵ ┻━┻
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, TABLE_FLIPPED)
	io.WriteString(out, "Woops, we ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
