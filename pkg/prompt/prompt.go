package prompt

import (
	"fmt"
	"strings"
)

func GenerateCommitPrompt(gitDiff string, scope string) string {
	scope = strings.TrimSpace(scope)
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
2. **Scope Handling**: 
   - THE SCOPE TO USE IS: "%s"
   - If a scope is provided above, you MUST use it in the format: type(scope): summary.
   - If the scope above is empty, infer a scope from the changed files (e.g., api, ui, auth) or omit it if unsure.
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
	`, scope, gitDiff)
}
