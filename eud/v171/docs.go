// Package eud provides a compatibility layer for Panic's EUD scripts.
//
// This package reimplements APIs from the Panic's EUD project:
// https://gitlab.com/happysoft3/eud-maps-project/-/tree/master/eud_project/libs.
//
// OpenNox cannot support direct memory access, thus any memory manipulation APIs are unavailable.
// Stubs for these APIs may still be present in the library. Calling them will cause a panic.
package eud
