package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/PatrikOlin/monkey_interpreter/lexer"
	"github.com/PatrikOlin/monkey_interpreter/parser"
	"github.com/PatrikOlin/monkey_interpreter/compiler"
	"github.com/PatrikOlin/monkey_interpreter/vm"
)

const PROMPT = ">> "
const TABLE_FLIPPED = `
(╯°□°）╯︵ ┻━┻
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
			
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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "God damnit! Compilation failed, what is it this time? \n %s\n",
				err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
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
