// Package mcpsrv serves the artifact to a model over MCP on stdio. Queries the
// model initiates are tools; the artifact itself is a resource; canned review
// workflows are prompts.
// It is a read only query surface. It holds no session state, never sees the
// conversation, and must write nothing but MCP messages to stdout.
package mcpsrv
