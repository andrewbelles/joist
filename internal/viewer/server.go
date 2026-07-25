// The HTTP server. It serves the embedded SPA and a read only artifact endpoint.
// There is one code path here regardless of build tag; the tag only decides what
// assets holds.

package viewer

import (
	"errors"
	"io/fs"
	"net/http"
)

// ErrNotImplemented marks the scaffold.
var ErrNotImplemented = errors.New("viewer: not implemented")

// DefaultAddr binds loopback. Binding any other interface would put the
// artifact on the network, which the design does not allow.
const DefaultAddr = "127.0.0.1:7373"

// Server serves one artifact. It never mutates anything.
type Server struct {
	Addr         string
	ArtifactPath string
}

// New returns a server bound to addr, serving the artifact at path.
func New(addr, artifactPath string) *Server {
	if addr == "" {
		addr = DefaultAddr
	}
	return &Server{Addr: addr, ArtifactPath: artifactPath}
}

// Handler returns the routes. The artifact endpoint is read only by
// construction: no method other than GET is registered.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/artifact", s.handleArtifact)
	mux.Handle("GET /", http.FileServerFS(assets))
	return mux
}

// ListenAndServe starts the server on s.Addr.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Addr, s.Handler())
}

func (s *Server) handleArtifact(w http.ResponseWriter, r *http.Request) {
	http.Error(w, ErrNotImplemented.Error(), http.StatusNotImplemented)
}

// mustSub roots an embedded tree at dir. The argument is a compile time constant
// in both build variants, so a failure here is a build error, not a runtime one.
func mustSub(f fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(f, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
