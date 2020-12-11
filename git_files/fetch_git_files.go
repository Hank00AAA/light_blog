package git_files

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// UpdateLocalCache
func UpdateLocalCache() *git.Repository {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/Hank00AAA/light_blog.git",
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("UpdateLocalCache succ")
	return r
}