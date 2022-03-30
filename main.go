package main

import (
	"os"
	"path/filepath"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	maxCacheAge     = 180 * time.Minute
	repo            = "mjhuber/alfred-repos"
	query           string
	paths           []string
	// aw.Workflow is the main API
	wf *aw.Workflow
)

// Options contains options for connecting to the gitlab API
type Options struct {
	Directory string `env:"REPOS_DIRECTORY"`
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	paths = []string{}
}

func main() {
	wf.Run(run)
}

// run executes the Script Filter.
func run() {
	query := wf.Args()[0]
	opts := &Options{}
	cfg := aw.NewConfig()
	if err := cfg.To(opts); err != nil {
		wf.Fatalf("Error loading variables: %v", err)
		return
	}

	err := filepath.Walk(opts.Directory, visit)
	if err != nil {
		wf.Fatalf("Error traversing path: %s", err.Error())
	}

	for _, path := range paths {

		// Convenience method. Sets Item title to filename, subtitle
		// to shortened path, arg to full path, and icon to file icon.
		it := wf.NewFileItem(path)

		// Alternate actions
		it.NewModifier(aw.ModCmd).
			Subtitle("Open in VSCode").
			Var("action", "code")

		it.NewModifier(aw.ModOpt).
			Subtitle("Open in SublimeText").
			Var("action", "sublime")

		it.NewModifier(aw.ModCtrl).
			Subtitle("Open in iTerm").
			Var("action", "iterm")
	}

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching folders found", "Try a different query?")
	wf.SendFeedback()
}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() && f.Name() == ".git" {
		parent := filepath.Dir(path)
		paths = append(paths, parent)
	}
	return nil
}
