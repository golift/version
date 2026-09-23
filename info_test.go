package version_test

import (
	"strings"
	"testing"

	"golift.io/version"
)

func TestPrintInfoAndBuildContext(t *testing.T) { //nolint:paralleltest // mutates link-time variables.
	origVersion := version.Version
	origRevision := version.Revision
	origBranch := version.Branch
	origUser := version.BuildUser
	origDate := version.BuildDate
	t.Cleanup(func() {
		version.Version = origVersion
		version.Revision = origRevision
		version.Branch = origBranch
		version.BuildUser = origUser
		version.BuildDate = origDate
	})

	version.Version = "1.2.3"
	version.Revision = "abc123"
	version.Branch = "main"
	version.BuildUser = "tester"
	version.BuildDate = "2026-09-22"

	info := version.Info()
	if !strings.Contains(info, "version=1.2.3") ||
		!strings.Contains(info, "branch=main") ||
		!strings.Contains(info, "revision=abc123") {
		t.Fatalf("Info() %s", info)
	}

	build := version.BuildContext()
	if !strings.Contains(build, "user=tester") ||
		!strings.Contains(build, "date=2026-09-22") ||
		!strings.Contains(build, "go=") {
		t.Fatalf("BuildContext() %s", build)
	}

	printed := version.Print("notifiarr")
	wants := []string{
		"notifiarr, version 1.2.3",
		"branch: main",
		"revision: abc123",
		"build user:       tester",
		"build date:       2026-09-22",
	}
	for _, want := range wants {
		if !strings.Contains(printed, want) {
			t.Fatalf("Print() missing %q in %s", want, printed)
		}
	}
}
