package blog_data

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// BlogData
type BlogData struct {
	Repo *git.Repository
}

// GetFileContent
func (data *BlogData)GetFileContent(file string) (string, error)  {
	filePtr, err := data.toTree().File(file)
	if err != nil {
		fmt.Printf("GetFile err:%v", err)
		return "", err
	}

	content, err := filePtr.Contents()
	if err != nil {
		fmt.Printf("Content err:%v", err)
		return "", err
	}

	return content, err
}

// toTree
func (data *BlogData)toTree() *object.Tree {
	ref, err := data.Repo.Head()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	commit, err := data.Repo.CommitObject(ref.Hash())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	tree, err := commit.Tree()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return tree
}

// ListBlogFiles
func (data *BlogData)ListBlogFiles() []string {
	res := []string{}
	err := data.toTree().Files().ForEach(func(f *object.File) error {
		res = append(res, f.Name)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return res
}
