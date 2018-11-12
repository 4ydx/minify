package benchmarks

import (
	"io/ioutil"
    "bytes"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/parse/v2/buffer"
)

var m = minify.New()
var r = map[string]*bytes.Buffer{}
var w = map[string]*buffer.Writer{}

func load(filename string) {
	sample, _ := ioutil.ReadFile(filename)
	r[filename] = bytes.NewBuffer(sample)
	w[filename] = buffer.NewWriter(make([]byte, 0, len(sample)))
}
