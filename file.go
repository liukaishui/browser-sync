package main

import (
	"io/ioutil"
	"time"
)

var (
	errors = make(chan interface{})
	change = make(chan bool)
)

type Tree struct {
	Path  string
	IsDir bool
	Hash  string
	Tree  map[string]*Tree
}

func DirTree(tree *Tree, ignore []string) (*Tree, error) {
	fs, err := ioutil.ReadDir(tree.Path)
	if err != nil {
		return tree, err
	}

	for _, v := range fs {
		vTree := &Tree{
			Path: tree.Path + "/" + v.Name(),
			Tree: make(map[string]*Tree),
		}

		if InSlice(vTree.Path, ignore) {
			continue
		}

		if v.IsDir() {
			vTree.IsDir = true
			if _, err := DirTree(vTree, ignore); err != nil {
				return tree, err
			}
		} else {
			vHash, err := Md5File(vTree.Path)
			if err != nil {
				continue
			}
			vTree.Hash = vHash
		}

		tree.Tree[vTree.Path] = vTree
	}

	return tree, nil
}

func TreeDiff(old *Tree, new *Tree) bool {
	if old.Path != new.Path {
		return false
	}
	if old.IsDir != new.IsDir {
		return false
	}
	if old.Hash != new.Hash {
		return false
	}
	if len(old.Tree) != len(new.Tree) {
		return false
	}

	for k, _ := range old.Tree {
		if _, ok := new.Tree[k]; !ok {
			return false
		}
		ok := TreeDiff(old.Tree[k], new.Tree[k])
		if !ok {
			return false
		}
	}

	return true
}

func WatcherDir(path string, d time.Duration, ignore []string) {
	for k, v := range ignore {
		ignore[k] = path + "/" + v
	}

	oldTree, err := DirTree(
		&Tree{
			Path:  path,
			IsDir: true,
			Tree:  make(map[string]*Tree),
		},
		ignore,
	)
	if err != nil {
		errors <- err
		return
	}

	for range time.Tick(d) {
		newTree, err := DirTree(
			&Tree{
				Path:  path,
				IsDir: true,
				Tree:  make(map[string]*Tree),
			},
			ignore,
		)
		if err != nil {
			errors <- err
			return
		}

		ok := TreeDiff(oldTree, newTree)
		if !ok {
			change <- ok
		}

		oldTree = newTree
	}
}
