	"code.gitea.io/gitea/models/db"
			_, err := io.Copy(io.Discard, reader)
				count, err := db.Count(m)
func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
			workdir, err := os.MkdirTemp("", "empty-work-dir")
	separator := "..."
	if directComparison {
		separator = ".."
	}

	shortstatArgs := []string{beforeCommitID + separator + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior, directComparison)