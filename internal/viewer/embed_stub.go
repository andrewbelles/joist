// Placeholder assets for a build without the spa tag. Serving a page that says
// how to build the SPA is better than failing to compile, which is what a bare
// go:embed of an absent dist tree would do.

//go:build !spa

package viewer

import (
	"embed"
	"io/fs"
)

//go:embed stubui
var stub embed.FS

// assets is the served file tree, rooted so that index.html is at the top.
var assets fs.FS = mustSub(stub, "stubui")

// Built reports whether the real SPA is compiled in.
const Built = false
