// Package lib is create for vendoring reasons, see: https://github.com/golang/go/issues/26366.
package lib

import (
	_ "github.com/NMAC427/wasmer-go/wasmer/packaged/lib/darwin-amd64"
	_ "github.com/NMAC427/wasmer-go/wasmer/packaged/lib/linux-aarch64"
	_ "github.com/NMAC427/wasmer-go/wasmer/packaged/lib/linux-amd64"
)
