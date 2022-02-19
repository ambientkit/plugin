// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytesize

import (
	"fmt"
)

// ByteSize is a float64.
type ByteSize float64

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KB is kilobytes.
	KB ByteSize = 1 << (10 * iota)
	// MB is megabytes.
	MB
	// GB is gigabytes
	GB
	// TB is terabytes
	TB
	// PB is petabytes.
	PB
	// EB is exabytes.
	EB
	// ZB is zettabytes.
	ZB
	// YB is yottabytes.
	YB
)

// String returns float64 in bytes.
func String(b float64, decimals int) string {
	return ByteSize(b).String(decimals)
}

func (b ByteSize) String(decimals int) string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f YB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f ZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f EB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f KB", b/KB)
	}
	return fmt.Sprintf("%."+fmt.Sprint(decimals)+"f B", b)
}
