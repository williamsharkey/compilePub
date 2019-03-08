package gogit

import (
	"io/ioutil"

	"gopkg.in/src-d/go-billy.v3/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Commit points to a single tree, marking it as what the project looked like
// at a certain point in time. It contains meta-information about that point
// in time, such as a timestamp, the author of the changes since the last
// commit, a pointer to the previous commit(s), etc.
// http://schacon.github.io/gitbook/1_the_git_object_model.html
type Commit object.Commit

// CommitFile represents git file objects.
type CommitFile object.File

// var Files = gito.Commit.Files

// Clone a repository into the given Storer and worktree Filesystem with the
// given options, if worktree is nil a bare repository is created. If the given
// storer is not empty ErrRepositoryAlreadyExists is returned.
//
// The provided Context must be non-nil. If the context expires before the
// operation is complete, an error is returned. The context only affects to the
// transport operations.
func Clone(user, auth, url string, f func(c Commit) error) (err error) {
	ba := http.NewBasicAuth(user, auth)
	storer := memory.NewStorage()
	memFs := memfs.New()

	repo, err := git.Clone(storer, memFs, &git.CloneOptions{
		URL:  url,
		Auth: ba,
	})

	if err != nil {

		return
	}

	commitIter, err := repo.CommitObjects()
	if err != nil {
		return
	}

	err = commitIter.ForEach(func(c *object.Commit) error {
		myCommit := Commit(*c)
		err := f(myCommit)
		return err

	})

	return

}

//MakeFileAccessor makes a function to return files.
func MakeFileAccessor(c Commit) func(func(CommitFile, []byte)) {

	fn := func(process func(CommitFile, []byte)) {

		q := object.Commit(c)

		headFilesItr, err := q.Files()

		if err != nil {
			return
		}

		headFilesItr.ForEach(func(f *object.File) (err error) {

			reader, err := f.Reader()
			if err != nil {
				return
			}

			b, err := ioutil.ReadAll(reader)
			if err != nil {
				return
			}

			myFile := CommitFile(*f)
			process(myFile, b)
			return
		})

		return
	}

	return fn
}
