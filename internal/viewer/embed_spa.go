// SPA assets for a release build. The dist tree is produced by the web build and
// is not committed, so this file only compiles under the spa tag. That tag is
// what keeps a plain go build from needing Node installed.

//go:build spa

package viewer

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var dist embed.FS

// assets is the served file tree, rooted so that index.html is at the top.
var assets fs.FS = mustSub(dist, "dist")

// Built reports whether the real SPA is compiled in.
const Built = true
