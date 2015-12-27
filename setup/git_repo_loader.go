package setup

import (
	"log"
	"os"
	"path/filepath"

	git "github.com/nildev/project/Godeps/_workspace/src/github.com/libgit2/git2go"
)

type (
	gitRepoLoader struct {
		branch string
	}
)

// NewGitRepoLoader returns instance of new repo loader
func NewGitRepoLoader(branch string) RepositoryLoader {
	return &gitRepoLoader{
		branch: branch,
	}
}

func credentialsCallback(url string, username_from_url string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	ret, cred := git.NewCredSshKeyFromAgent(username_from_url)
	return git.ErrorCode(ret), &cred
}

func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

// Load repo over git
func (grl *gitRepoLoader) Load(repoPath string, destDir string) {

	options := &git.CloneOptions{
		FetchOptions: &git.FetchOptions{
			RemoteCallbacks: git.RemoteCallbacks{
				CredentialsCallback:      credentialsCallback,
				CertificateCheckCallback: certificateCheckCallback,
			},
		},
	}

	_, err := git.Clone(repoPath, destDir, options)
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll(destDir + string(filepath.Separator) + ".git")
	if err != nil {
		log.Fatal(err)
	}
}
