package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func fileWriteSimplify(m dsl.Matcher) {
	m.Match(`os.Stderr.WriteString(fmt.Sprintf($*args))`).
		Suggest(`fmt.Fprintf(os.Stderr, $args)`)

	m.Match(`fmt.Fprint(os.Stdout, $*args)`).Suggest(`fmt.Print($args)`)
	m.Match(`fmt.Fprintln(os.Stdout, $*args)`).Suggest(`fmt.Println($args)`)
	m.Match(`fmt.Fprintf(os.Stdout, $*args)`).Suggest(`fmt.Printf($args)`)
}
