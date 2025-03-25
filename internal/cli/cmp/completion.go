package cmp

import (
	"fmt"

	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"

	"github.com/zk-org/zk/internal/cli"
	"github.com/zk-org/zk/internal/core"
)

func Completer(c *cli.Container) *complete.Command {

	cmdTag := &complete.Command{
		Args: predict.Set{"list"},
	}
	cmdLsp := &complete.Command{}
	cmdList := &complete.Command{}
	cmdInit := &complete.Command{}
	cmdIndex := &complete.Command{}
	cmdGraph := &complete.Command{}
	cmdNew := &complete.Command{}
	cmdEdit := &complete.Command{
		Args: getFiles(c),
		Flags: map[string]complete.Predictor{
			"t": getTags(c),
		},
	}

	cmpZk := &complete.Command{
		Sub: map[string]*complete.Command{
			"edit":  cmdEdit,
			"graph": cmdGraph,
			"index": cmdIndex,
			"init":  cmdInit,
			"list":  cmdList,
			"lsp":   cmdLsp,
			"new":   cmdNew,
			"tag":   cmdTag,
		},
	}
	return cmpZk
}

func getTags(c *cli.Container) predict.Set {
	tagSet := predict.Set{}
	notebook, err := c.CurrentNotebook()
	if notebook == nil || err != nil {
		panic("Notebook is nil, what are you doing?")
	}
	mins, _ := notebook.FindNote(core.NoteFindOpts{})

	tagSet = append(tagSet, mins.Tags...)

	return tagSet
}

func getFiles(c *cli.Container) predict.Set {
	filesSet := predict.Set{}

	notebook, err := c.CurrentNotebook()
	if notebook == nil {
		panic("Notebook is nil, what are you doing?")
	}

	notes, err := notebook.FindMinimalNotes(core.NoteFindOpts{})
	if err != nil {
		panic("Couldn't get mins")
	}
	for _, m := range notes {
		//TODO: format w title
		strRep := fmt.Sprintf("%s", m.Path)
		filesSet = append(filesSet, strRep)
	}
	return filesSet
}
