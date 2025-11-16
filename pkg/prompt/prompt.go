package prompt

import "fmt"

func GenerateCommitPrompt(gitDiff string) string {
	return fmt.Sprintf(`
You are an assistant that writes git commit messages following the Conventional Commits specification.
Read the following git diff and generate a single commit message that follows these rules:
1. Use one of these types:
   - feat: for new features
   - fix: for bug fixes
   - refactor: for code refactors that don’t fix a bug or add a feature
   - docs: for documentation only changes
   - test: for tests only
   - style: for formatting, whitespace, or non-functional style changes
   - chore: for build tools, config, dependencies, or other non-app changes
   - perf: for performance improvements
   - ci: for CI-related changes
   - build: for build system changes
   - revert: for reverting a previous commit
2. If possible, infer a scope from the changed files or directories.
   Example scopes: api, ui, auth, db, config, tests, etc.
   Format: type(scope): summary
   If you’re not sure, omit the scope.
3. The summary:
   - Must be in the imperative mood (e.g. "add user validation", not "added" or "adds").
   - Max ~72 characters.
   - No trailing period.
4. Add an optional body only if it adds useful context:
   - Use short paragraphs or bullet points.
   - Explain *why* the change was made or any important details.
   - Mention notable side effects, refactors, or risks.
5. If there is a breaking change:
   - Add a line at the end of the body:
     BREAKING CHANGE: <description of the breaking change>
6. Do NOT mention git diff, file names literally, or internal paths unless needed for clarity.
   Focus on what the change does at a feature/behavior level.
Now, here is the git diff:
%s
Return ONLY the commit message, nothing else.
	`, gitDiff)
}
