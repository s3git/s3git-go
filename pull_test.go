package s3git

import (
	"fmt"
	"github.com/s3git/s3git-go/internal/core"
	_ "github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestPull(t *testing.T) {
	repo, path := setupRepo()
	fmt.Println(path)
	//defer teardownRepo(path)

	for i := 0; i < 10; i++ {
		repo.Add(strings.NewReader(fmt.Sprintf("hello s3git: %d, %s", i, time.Now())))
	}

	hash, _, _ := repo.Commit("1st commit")

	core.GetCommitObject(hash)

	repo.Push()

	repo.Pull()
}
