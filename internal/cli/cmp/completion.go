package cmp

import (
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
)

var cmdTag = &complete.Command{}
var cmdIew = &complete.Command{}
var cmdLsp = &complete.Command{}
var cmdList = &complete.Command{}
var cmdInit = &complete.Command{}
var cmdIndex = &complete.Command{}
var cmdGraph = &complete.Command{}
var cmdNew = &complete.Command{}
var cmdEdit = &complete.Command{
	Args: getFiles(),
}

var CmpZk = &complete.Command{
	Sub: map[string]*complete.Command{
		"edit":  cmdEdit,
		"graph": cmdGraph,
		"index": cmdIndex,
		"init":  cmdIndex,
		"list":  cmdList,
		"lsp":   cmdLsp,
		"new":   cmdNew,
		"tag":   cmdTag,
	},
}

func getFiles() predict.Set {
	files := predict.Set{}
	files = append(files, "foo")
	files = append(files, "bar")
	files = append(files, "baz")
	return files
}
