//go:build mage
// +build mage

// +agentlint:enforce

// Package main provides build, test, and release automation for VerdiPitchEngine.
// @layer: Infrastructure
// @ref-rule: 100-CORE
// @constraint: Automated lifecycle management and quality gates
// @complexity: Low
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/AIgorLabs/Archon/pkg/tasksync"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Help

var Aliases = map[string]interface{}{
	"lint-markdown":        LintMarkdown,
	"fix-markdown":         FixMarkdown,
	"fmt":                  Fmt,
	"vet":                  Vet,
	"vulncheck":            Vulncheck,
	"lint":                 Lint,
	"test":                 Test,
	"ci-test":              CiTest,
	"check-coverage":       CheckCoverage,
	"check":                Check,
	"deps":                 Deps,
	"deploy":               Deploy,
	"build":                Build,
	"syncissues":           SyncIssues,
	"sync-issues":          SyncIssues,
	"syncissuestargeted":   SyncIssuesTargeted,
	"sync-issues-targeted": SyncIssuesTargeted,
	"syncbacklogs":         SyncBacklogs,
	"sync-backlogs":        SyncBacklogs,
	"audit-task-sync":      AuditTaskSync,
	"audittasksync":        AuditTaskSync,
	"audit-backlog":        AuditBacklog,
	"auditbacklog":         AuditBacklog,
	"next-task":            NextTask,
	"nexttask":             NextTask,
	"commit":               Commit,
	"commitmain":           CommitMain,
	"brutal":               Brutal,
	"brutal-consensus":     BrutalConsensus,
	"brutalconsensus":      BrutalConsensus,
	"brutal-attest":        BrutalAttest,
	"brutalattest":         BrutalAttest,
	"brutal-staged":        BrutalStaged,
	"brutalstaged":         BrutalStaged,
	"brutal-tasks":         BrutalTasks,
	"brutaltasks":          BrutalTasks,
	"brutal-tasks-force":   BrutalTasksForce,
	"brutaltasksforce":     BrutalTasksForce,
	"brutal-tasks-open":    BrutalTasksOpen,
	"brutaltasksopen":      BrutalTasksOpen,
	"brutal-scaffold":      BrutalScaffold,
	"brutalscaffold":       BrutalScaffold,
	"check-brutal":         CheckBrutal,
	"checkbrutal":          CheckBrutal,
	"ops-gap":              OpsGap,
	"opsgap":               OpsGap,
	"ops-swot":             OpsSwot,
	"opsswot":              OpsSwot,
	"ops-root":             OpsRoot,
	"opsroot":              OpsRoot,
}

var (
	Go = "go"
)

// Help displays the available targets.
func Help() {
	fmt.Println("Run 'mage -l' to see available targets.")
}

// LintMarkdown lints Markdown files for 1000-KEYS compliance using central script.
func LintMarkdown() error {
	lintScript := "../Forge/scripts/lint-markdown.py"
	if _, err := os.Stat(lintScript); os.IsNotExist(err) {
		lintScript = "../AIgorLabs-github/scripts/lint-markdown.py"
	}
	if _, err := os.Stat(lintScript); err == nil {
		fmt.Printf("Running central markdown linter from local workspace (%s)...\n", lintScript)
		return sh.RunV("python3", lintScript)
	}
	fmt.Println("WARNING: ../Forge and ../AIgorLabs-github not found. Skipping local markdown linting.")
	return nil
}

// FixMarkdown automatically adds or updates the 1000-KEYS YAML frontmatter on all markdown files.
func FixMarkdown() error {
	fmt.Println("Running Markdown fixer...")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path == ".git" || path == "vendor" || path == "node_modules" || path == "bin" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		if info.Name() == "CHANGELOG.md" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		contentStr := string(content)
		var body string
		var frontmatter string
		hasFrontmatter := false

		// Regex to match existing frontmatter
		re := regexp.MustCompile(`(?s)^---\n(.*?)\n---(?:\n|$)(.*)`)
		matches := re.FindStringSubmatch(contentStr)

		if len(matches) == 3 {
			hasFrontmatter = true
			frontmatter = matches[1]
			body = matches[2]
		} else {
			body = contentStr
		}

		// Calculate body hash
		hash := sha256.Sum256([]byte(body))
		bodyHash := fmt.Sprintf("%x", hash)[:16]

		dateStr := time.Now().Format("2006-01-02")

		if !hasFrontmatter {
			// Generate new frontmatter
			frontmatter = fmt.Sprintf(`project_name: VerdiPitchEngine
version: 0.1.0
status: active
priority: high
dev_stage: development
agent_role: lead_engineer
agent_weight: 1.0
asset_scope: backend
platform: qnap
tech_stack: [go, ffmpeg]
dependencies: []
created: %s
updated: %s
tags: [audio, dsp]
body_hash: %s`, dateStr, dateStr, bodyHash)
		} else {
			// Update body_hash in existing frontmatter
			reHash := regexp.MustCompile(`(?m)^body_hash\s*:.*$`)
			if reHash.MatchString(frontmatter) {
				frontmatter = reHash.ReplaceAllString(frontmatter, fmt.Sprintf(`body_hash: %s`, bodyHash))
			} else {
				frontmatter += fmt.Sprintf("\nbody_hash: %s", bodyHash)
			}
			// Update updated date
			reUpdated := regexp.MustCompile(`(?m)^updated\s*:.*$`)
			if reUpdated.MatchString(frontmatter) {
				frontmatter = reUpdated.ReplaceAllString(frontmatter, fmt.Sprintf(`updated: %s`, dateStr))
			}
		}

		newContent := fmt.Sprintf("---\n%s\n---\n%s", frontmatter, body)
		if newContent != contentStr {
			fmt.Printf("Fixed frontmatter for %s\n", path)
			return ioutil.WriteFile(path, []byte(newContent), info.Mode())
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("fix-markdown failed: %w", err)
	}
	fmt.Println("Markdown fixing complete!")
	return nil
}

// Deps installs project dependencies.
func Deps() error {
	fmt.Println("Installing dependencies...")
	return sh.RunV(Go, "mod", "download")
}

// Fmt formats code.
func Fmt() error {
	fmt.Println("Formatting code...")
	if err := sh.RunV(Go, "fmt", "./cmd/...", "./internal/...", "./pkg/..."); err != nil {
		return err
	}
	fmt.Println("Formatting complete!")
	return nil
}

// Vet runs go vet.
func Vet() error {
	fmt.Println("Running go vet...")
	if err := sh.RunV(Go, "vet", "./cmd/...", "./internal/...", "./pkg/..."); err != nil {
		return err
	}
	fmt.Println("Vet complete!")
	return nil
}

// Vulncheck runs govulncheck for dependency vulnerability scanning.
func Vulncheck() error {
	fmt.Println("Running govulncheck...")
	err := sh.RunV("govulncheck", "./...")
	if err != nil {
		fmt.Println("govulncheck failed or not installed. Installing...")
		if installErr := sh.RunV(Go, "install", "golang.org/x/vuln/cmd/govulncheck@latest"); installErr != nil {
			return installErr
		}
		if runErr := sh.RunV("govulncheck", "./..."); runErr != nil {
			return runErr
		}
	}
	fmt.Println("Vulncheck complete!")
	return nil
}

// Lint runs linters using strict .golangci.yml.
func Lint() error {
	fmt.Println("Running linters...")
	if err := sh.RunV("golangci-lint", "run", "--timeout", "5m", "./..."); err != nil {
		fmt.Println("golangci-lint failed or not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
		return err
	}
	fmt.Println("Lint complete!")
	return nil
}

// Test runs all tests.
func Test() error {
	fmt.Println("Running tests...")
	if err := sh.RunV(Go, "test", "-v", "-race", "-coverprofile=coverage.out", "./cmd/...", "./internal/...", "./pkg/..."); err != nil {
		return err
	}
	if err := sh.RunV(Go, "tool", "cover", "-html=coverage.out", "-o", "coverage.html"); err != nil {
		return err
	}
	fmt.Println("Tests complete! Coverage report: coverage.html")
	return nil
}

// CiTest runs tests without race detector for CI speed.
func CiTest() error {
	fmt.Println("Running CI tests...")
	return sh.RunV(Go, "test", "-v", "-coverprofile=coverage.out", "./cmd/...", "./internal/...", "./pkg/...")
}

// CheckCoverage enforces test coverage minimum targets across the workspace.
func CheckCoverage() error {
	mg.Deps(Test)
	fmt.Println("Enforcing test coverage targets...")
	return sh.RunV(Go, "run", "scripts/check_coverage.go", "coverage.out")
}

// Check runs all checks (format, vet, lint, vulncheck, check-coverage).
func Check() {
	mg.Deps(Fmt, Vet, Lint, LintMarkdown, Vulncheck, CheckCoverage)
}

// Deploy deploys current workspace to QNAP Container Station.
// You can pass a specific NAS target directory via VERDI_TARGET_DIR environment variable.
func Deploy() error {
	fmt.Printf("\n🚀 \033[1;35mINITIATING ZERO-TRUST QNAP SYNCHRONIZATION\033[0m 🚀\n")
	fmt.Printf("\033[1;30m------------------------------------------------\033[0m\n")
	fmt.Printf("📦 \033[1;34mSyncing VerdiPitchEngine to Remote Enclave Mirror...\033[0m\n")
	
	targetDir := os.Getenv("VERDI_TARGET_DIR")
	
	commit, err := sh.Output("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		commit = "latest"
	}
	imageTag := fmt.Sprintf("v0.3.0-%s", commit)
	fmt.Printf("🏷️  \033[1;36mVersioning Image as:\033[0m %s\n", imageTag)
	
	// Write to .env to safely pass to docker-compose without violating the SSH proxy strict whitelisting
	envContent := fmt.Sprintf("IMAGE_TAG=%s\nVERSION=0.3.0\nBUILD_HASH=%s\n", imageTag, commit)
	if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
		return err
	}

	rsyncCmd := `rsync -az --progress --chmod=a+rx -e "ssh -i ~/.ssh/id_ed25519_antigravity" --exclude='.git' --exclude='bin' . antigravity_agent@$(grep '^QNAP_HOST=' ../OpenBrain/.env | cut -d'"' -f2):/share/AIgorLabs/enclaves/VerdiPitchEngine/`
	if err := sh.RunV("bash", "-c", rsyncCmd); err != nil {
		return err
	}
	
	fmt.Printf("🔄 \033[1;34mRestarting VerdiPitchEngine container via Zero-Trust Proxy...\033[0m\n")
	
	var deployArg string
	if targetDir != "" {
		fmt.Printf("📂 \033[1;36mTarget NAS Directory Override:\033[0m %s\n", targetDir)
		deployArg = " " + targetDir
	}

	sshCmd := fmt.Sprintf(`ssh -i ~/.ssh/id_ed25519_antigravity antigravity_agent@$(grep '^QNAP_HOST=' ../OpenBrain/.env | cut -d'"' -f2) "deploy verdipitchengine%s"`, deployArg)
	if err := sh.RunV("bash", "-c", sshCmd); err != nil {
		return err
	}
	
	fmt.Printf("\n✨ \033[1;32mDEPLOYMENT COMPLETE\033[0m ✨\n\n")
	return nil
}

// Build compiles the project.
func Build() error {
	fmt.Println("Building binaries...")
	err := os.MkdirAll("bin", 0755)
	if err != nil {
		return err
	}
	commit, err := sh.Output("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		commit = "unknown"
	}

	ldflags := fmt.Sprintf("-X main.Version=0.3.0 -X main.Build=%s", commit)

	if err := sh.RunV(Go, "build", "-ldflags", ldflags, "-o", "bin/verdi", "./cmd/verdi"); err != nil {
		return err
	}
	fmt.Println("Build complete!")
	return nil
}

// Brutal runs an unsparing 1-to-10 architectural evaluation on the current workspace with Spinal Tap 11/10 mode.
func Brutal() error {
	return sh.RunV("arc-brutal", "-spinal-tap")
}

// BrutalConsensus runs a heterogeneous multi-model consensus brutal audit with 2/3 approval threshold.
func BrutalConsensus() error {
	return sh.RunV("arc-brutal", "-consensus", "-unconstrained", "-spinal-tap")
}

// BrutalAttest runs a brutal evaluation and generates a hardware attestation receipt.
func BrutalAttest() error {
	return sh.RunV("arc-brutal", "-attest", "-spinal-tap")
}

// BrutalStaged runs an incremental brutal evaluation restricted only to currently staged git files.
func BrutalStaged() error {
	return sh.RunV("arc-brutal", "-staged-only")
}

// BrutalTasks evaluates Master Task Dossiers (MTDs) across .agent/tasks/.
func BrutalTasks() error {
	return sh.RunV("arc-brutal", "-mtd")
}

// BrutalTasksOpen evaluates only open and in-progress Master Task Dossiers (MTDs).
func BrutalTasksOpen() error {
	return sh.RunV("arc-brutal", "-mtd", "-open")
}

// BrutalTasksForce forces re-evaluation of open Master Task Dossiers (MTDs), bypassing the cache.
func BrutalTasksForce() error {
	return sh.RunV("arc-brutal", "-mtd", "-open", "-force")
}

// BrutalScaffold runs an unsparing 1-to-10 architectural evaluation and auto-scaffolds Master Task Dossiers for findings.
func BrutalScaffold() error {
	return sh.RunV("arc-brutal", "-scaffold-remediation", "-spinal-tap")
}

// CheckBrutal executes the brutal architectural quality gate ensuring the target meets the minimum score threshold (8.0).
func CheckBrutal() error {
	return sh.RunV("arc-brutal", "-min-score", "8.0")
}

// OpsGap runs an unsparing /gen-brutal operational gap analysis against a specified target artifact with Spinal Tap 11/10 mode.
func OpsGap(target string) error {
	args := []string{"-spinal-tap"}
	if target != "" {
		args = append(args, target)
	}
	return sh.RunV("arc-brutal", args...)
}

// OpsRoot runs a forensic Root Cause Analysis (RCA) and 5-Whys defect autopsy against a specified issue or incident topic.
func OpsRoot(target string) error {
	args := []string{"-spinal-tap"}
	if target != "" {
		args = append(args, target)
	}
	return sh.RunV("arc-brutal", args...)
}

// OpsSwot runs an unsparing /gen-brutal SWOT analysis against a specified target artifact with Spinal Tap 11/10 mode.
func OpsSwot(target string) error {
	args := []string{"-spinal-tap"}
	if target != "" {
		args = append(args, target)
	}
	return sh.RunV("arc-brutal", args...)
}

// extractTargetArgsFromCLI extracts optional target arguments from os.Args for specific target commands.
func extractTargetArgsFromCLI(targetNames ...string) []string {
	var targets []string
	foundTarget := false
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		if !foundTarget {
			for _, name := range targetNames {
				if strings.EqualFold(arg, name) || strings.EqualFold(arg, strings.ReplaceAll(name, "-", "")) {
					foundTarget = true
					break
				}
			}
			continue
		}
		clean := strings.TrimSpace(arg)
		if clean != "" && !strings.EqualFold(clean, "all") {
			targets = append(targets, strings.Fields(strings.ReplaceAll(clean, ",", " "))...)
		}
	}
	return targets
}

// SyncIssues synchronizes local BACKLOG.md checklists and native blockers with remote GitHub issue states via native Go tasksync package.
// When called with target arguments (e.g. mage syncissues VPE-001 or mage syncissues 1447 1448), only specified tasks/issues are synced.
func SyncIssues() error {
	targets := extractTargetArgsFromCLI("syncissues", "sync-issues", "syncbacklogs", "sync-backlogs")
	if len(targets) > 0 {
		fmt.Printf("Synchronizing targeted project backlog items %v with GitHub (Native Go)...\n", targets)
		if err := tasksync.SyncIssuesTargeted(targets); err != nil {
			return err
		}
		os.Exit(0)
	}
	fmt.Println("Synchronizing project backlog with GitHub (Native Go)...")
	return tasksync.SyncIssuesTargeted(nil)
}

// SyncIssuesTargeted limits synchronization to target issue numbers or task IDs (e.g. mage syncIssuesTargeted "20 21").
func SyncIssuesTargeted(target string) error {
	var targets []string
	if strings.TrimSpace(target) != "" {
		targets = strings.Fields(strings.ReplaceAll(target, ",", " "))
	}
	if len(targets) > 0 {
		fmt.Printf("Synchronizing targeted project backlog items %v with GitHub (Native Go)...\n", targets)
		return tasksync.SyncIssuesTargeted(targets)
	}
	fmt.Println("Synchronizing project backlog with GitHub (Native Go)...")
	return tasksync.SyncIssuesTargeted(nil)
}

// SyncBacklogs is an alias for SyncIssues.
func SyncBacklogs() error {
	return SyncIssues()
}

// AuditTaskSync runs automated check for drift between Markdown tasks and GitHub Issues via native Go tasksync package.
func AuditTaskSync() error {
	return tasksync.AuditTaskSync()
}

// AuditBacklog runs static schema, header, and Rule 000 taxonomy validation across project backlogs via native Go tasksync package.
func AuditBacklog() error {
	return tasksync.AuditBacklog()
}

// NextTask runs the automated Prudence Task Selection Engine via shared tasksync package.
func NextTask() error {
	return tasksync.NextTask()
}

// Commit executes the 020-GIT compliant commit wrapper via shared tasksync package.
func Commit() error {
	fmt.Println("🚀 Executing 020-GIT compliant commit wrapper...")
	return tasksync.Commit(false)
}

// CommitMain executes the 020-GIT compliant commit wrapper with main branch override.
func CommitMain() error {
	fmt.Println("🚀 Executing 020-GIT compliant commit wrapper (Main Branch Override)...")
	return tasksync.Commit(true)
}

