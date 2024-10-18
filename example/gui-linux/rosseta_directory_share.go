//go:build darwin && !arm64
// +build darwin,!arm64

package main

import "github.com/lima-vm/vz/v4"

func createRosettaDirectoryShareConfiguration() (*vz.VirtioFileSystemDeviceConfiguration, error) {
	return nil, errIgnoreInstall
}
