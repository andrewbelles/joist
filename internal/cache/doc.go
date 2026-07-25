// Package cache is the content addressed blob store and the per-SHA manifest.
// A key covers file content, compile flags, toolchain version and indexer
// version, so an entry is valid on any machine that computes the same key.
// A key must fully determine its value. Reading anything undeclared here turns
// every cache hit into a guess.
package cache
