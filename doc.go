// Package sqlconv converts between Go values and SQL values.
//
// Currently, [sqlconv] provides a interface [ScannerConverter], which is used
// to convert SQL values to Go values:
//
//	type ScannerConverter interface {
//	  ScanConvert(dest, src any) error
//	}
//
// [sqlconv.ScannerConverter.ScanConvert] converts SQL value src to Go value,
// then writes the Go value into dest.
//
// A predefined variable [DefaultScannerConverter] converts SQL values based
// on the rules defined by the function [database/sql.convertAssign].
package sqlconv
