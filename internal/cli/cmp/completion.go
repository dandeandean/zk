package cmp

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	//"strings"

	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
	fz "github.com/zk-org/zk/internal/adapter/fs"
	//"github.com/zk-org/zk/internal/cli"
	"github.com/zk-org/zk/internal/util"
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
	// Flags: map[string]complete.Predictor{
	// 	"--working-dir": predict.Set{"foo", "bar", "baz"},
	// },
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
	filesSet := predict.Set{}
	notebookDir := os.Getenv("ZK_NOTEBOOK_DIR")
	fmt.Println(notebookDir)

	//_ = cli.Filtering
	err := filepath.Walk(notebookDir,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			zkfs, _ := fz.NewFileStorage(path, util.StdLogger{})
			filesSet = append(filesSet, zkfs.Canonical(path))
			return err
		})
	if err != nil {
		panic("Something horrible happened!")
	}
	return filesSet
}
