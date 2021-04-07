package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func filepathJoin(m dsl.Matcher) {
	m.Match(
		`os.Open(path.Join($*_))`,
		`ioutil.ReadFile(path.Join($*_))`,
		`$p := path.Join($*_); $_, $_ := os.Open($p)`,
		`$p := path.Join($*_); $_, $_ := ioutil.ReadFile($p)`,
	).Report(`use filepath.Join for file paths`)
}
