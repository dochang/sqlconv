package sqlconv

import (
	"fmt"
	"time"
)

func ExampleScannerConverter() {
	var deststr string
	srctime := time.Unix(1, 0).UTC()
	err := DefaultScannerConverter.ScanConvert(&deststr, srctime)
	fmt.Println(err)
	fmt.Println(deststr)
	// Output:
	// <nil>
	// 1970-01-01T00:00:01Z
}
