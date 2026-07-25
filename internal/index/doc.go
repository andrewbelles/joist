// Package index orchestrates the subprocesses that produce raw graph input and
// defines the two contracts they speak. Indexer covers precise indexers, whose
// wire format is SCIP and therefore is not redefined here. Adapter covers
// framework adapters, which synthesize edges SCIP cannot express.
// This package invokes and collects. It does not decode SCIP or resolve symbols;
// that is the scip package.
package index
