// Package viewer serves the read only SPA and the artifact behind it over HTTP.
// The SPA is compiled into the binary under the spa build tag, so a release
// build has no static file dependency.
// Bind loopback only. This server exists so nothing has to leave the machine.
package viewer
