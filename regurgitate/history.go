package regurgitate

import (
	"github.com/functionserver/file"
	"github.com/functionserver/gogit"
	"github.com/functionserver/structure"
	"github.com/functionserver/trace"
)

// GetHistory returns a list of commits from newest to oldest.
func GetHistory(url, user, auth string) (commits map[string]structure.Commit, sortedHashes []structure.CommitTime, err error) {

	commits = map[string]structure.Commit{}

	err = gogit.Clone(user, auth, url, func(c gogit.Commit) (err error) {

		parentHashes := []string{}
		for _, h := range c.ParentHashes {
			parentHashes = append(parentHashes, h.String())
		}

		commitHash := c.Hash.String()
		commits[commitHash] =
			structure.Commit{
				Hash:         commitHash,
				Time:         c.Author.When,
				AuthorName:   c.Author.Name,
				AuthorEmail:  c.Author.Email,
				Message:      c.Message,
				ParentHashes: parentHashes,
				GetFiles:     makeFileLoader(c),
			}

		return
	})

	sortedHashes = structure.SortByTime(commits)

	return
}

// GetFiles gets the files for a commit
func GetFiles(commitHashMatch, url, user, auth string) (files []file.File, err error) {
	foundMatch := false
	err = gogit.Clone(user, auth, url, func(c gogit.Commit) (errIter error) {
		if foundMatch {
			return
		}

		commitHash := c.Hash.String()
		if commitHash == commitHashMatch {
			files = makeFileLoader(c)()
			foundMatch = true

		}
		return
	})

	if !foundMatch {
		err = trace.ErrorNew("didn't find match for commitHash " + commitHashMatch)
	}
	return

}

func makeFileLoader(c gogit.Commit) func() []file.File {
	return func() []file.File {
		files := []file.File{}
		gogit.MakeFileAccessor(c)(func(commitFile gogit.CommitFile, fileBytes []byte) {

			newFile := file.File{
				Bytes: fileBytes,
				Name:  commitFile.Name,
				Size:  commitFile.Size,
				Hash:  commitFile.Hash.String(),
			}

			files = append(files, newFile)

		})
		return files
	}
}
