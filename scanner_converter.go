package sqlconv

// ScannerConverter is the interface providing the ScanConvert method.
//
// ScannerConverter converts SQL value to Go value.
type ScannerConverter interface {
	// ScanConvert converts a SQL value src to a Go value dest.
	ScanConvert(dest, src any) error
}

// DefaultScannerConverter is the default implementation of ScannerConverter.
//
// DefaultScannerConverter converts value by calling convertAssign.
var DefaultScannerConverter defaultScannerConverter

type defaultScannerConverter struct{}

var _ ScannerConverter = defaultScannerConverter{}

func (defaultScannerConverter) ScanConvert(dest, src any) error {
	return convertAssign(dest, src)
}
