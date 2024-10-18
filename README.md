vz - Go binding with Apple [Virtualization.framework](https://developer.apple.com/documentation/virtualization?language=objc)
=======

[![Build](https://github.com/lima-vm/vz/actions/workflows/compile.yml/badge.svg)](https://github.com/lima-vm/vz/actions/workflows/compile.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/lima-vm/vz/v4.svg)](https://pkg.go.dev/github.com/lima-vm/vz/v4)

vz provides the power of the Apple Virtualization.framework in Go. Put here is block quote of overreview which is written what is Virtualization.framework from the document.

> The Virtualization framework provides high-level APIs for creating and managing virtual machines (VM) on Apple silicon and Intel-based Mac computers. Use this framework to boot and run macOS or Linux-based operating systems in custom environments that you define. The framework supports the [Virtual I/O Device (VIRTIO)](https://docs.oasis-open.org/virtio/virtio/v1.1/csprd01/virtio-v1.1-csprd01.html) specification, which defines standard interfaces for many device types, including network, socket, serial port, storage, entropy, and memory-balloon devices.

This repository <https://github.com/lima-vm/vz> was forked from <https://github.com/Code-Hex/vz/tree/c3198942b1b797ac424b673e06769f45bef0d275>.

## Usage

Please see the [example](./example) directory.

## Requirements

- Higher or equal to macOS Big Sur (11.0.0).
- Latest version of vz supports last two Go major [releases](https://go.dev/doc/devel/release) and might work with older versions.

## Installation

Initialize your project by creating a folder and then running `go mod init github.com/your/repo` ([learn more](https://go.dev/blog/using-go-modules)) inside the folder. Then install vz with the go get command:

```
$ go get github.com/lima-vm/vz/v4
```

## Feature Overview

- ✅ Virtualize Linux on a Mac **(x86_64, arm64)**
  - GUI Support
  - Boot Extensible Firmware Interface (EFI) ROM
  - Clipboard sharing through the SPICE agent
- ✅ Virtualize macOS on Apple Silicon Macs **(arm64)**
    - Fetches the latest restore image supported by this host from the network
  - Start in recovery mode
- ✅ Running Intel Binaries in Linux VMs with Rosetta **(arm64)**
- ✅ [Shared Directories](https://github.com/Code-Hex/vz/wiki/Shared-Directories)
- ✅ [Virtio Sockets](https://github.com/Code-Hex/vz/wiki/Sockets)
- ✅ Less dependent (only under golang.org/x/*)

## Important

For binaries used in this package, you need to create an entitlements file like the one below and apply the following command.

<details>
<summary>vz.entitlements</summary>

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>com.apple.security.virtualization</key>
	<true/>
</dict>
</plist>
```

</details>

```sh
$ codesign --entitlements vz.entitlements -s - <YOUR BINARY PATH>
```

> A process must have the com.apple.security.virtualization entitlement to use the Virtualization APIs.

If you want to use [`VZBridgedNetworkDeviceAttachment`](https://developer.apple.com/documentation/virtualization/vzbridgednetworkdeviceattachment?language=objc), you need to add also `com.apple.vm.networking` entitlement.

## Known compile-time warnings

If you compile using an older Xcode SDK, you will get the following warnings.

This example warns that macOS 12.3 API and macOS 13 API are not available in the binary build. This means these APIs are not available even if you are running this binary on a modern OS (macOS 12.3 or macOS 13). 

<!-- If you wish to change Code-Hex/vz/v3 below to lima-vm/vz/v4, make sure to obtain the actual console output with lima-vm/vz/v4 -->
```
$ go build .
# github.com/Code-Hex/vz/v3
In file included from _cgo_export.c:4:
In file included from socket.go:6:
In file included from ./virtualization_11.h:9:
./virtualization_helper.h:25:9: warning: macOS 12.3 API has been disabled [-W#pragma-messages]
./virtualization_helper.h:32:9: warning: macOS 13 API has been disabled [-W#pragma-messages]
```

If you want to build a binary that can use the API on all operating systems, make sure the Xcode SDK is up-to-date.

You can check the version of the Xcode SDK available for each macOS on this site.

https://xcodereleases.com/

## Version compatibility check

The package provides a mechanism for checking the availability of the respective API through error handling:

```go
bootLoader, err := vz.NewEFIBootLoader()
if errors.Is(err, vz.ErrUnsupportedOSVersion) || errors.Is(err, vz.ErrBuildTargetOSVersion) {
  return fallbackBootLoader()
}
if err != nil {
  return nil, err
}
return bootLoader, nil
```

There are two items to check.

1. API is compatible with the version of macOS
2. The binary was built with the API enabled

## Knowledge for the Apple Virtualization.framework

Quote from [@Code-Hex](https://github.com/Code-Hex)'s upstream repository <https://github.com/Code-Hex/vz/tree/c3198942b1b797ac424b673e06769f45bef0d275>:

> There is a lot of knowledge required to use this Apple Virtualization.framework, but the information is too scattered and very difficult to understand. In most cases, this can be found in [the official documentation](https://developer.apple.com/documentation/virtualization?language=objc). However, the Linux kernel knowledge required to use the feature provided by this framework is not documented. Therefore, I have compiled the knowledge I have gathered so far into this wiki.
>
> https://github.com/Code-Hex/vz/wiki
>
> Anyone is free to edit this wiki. It would help someone if you could add information not listed here. Let's make a good wiki together!

## Testing

If you want to contribute some code, you will need to add tests.

[PUI PUI Linux](https://github.com/Code-Hex/puipui-linux) is used to test this library. This Linux is designed to provide only the minimum functionality required for the Apple Virtualization.framework (Virtio), so the kernel file size is very small.

The test code uses the `Makefile` in the project root.

```
$ # Download PUI PUI Linux, Only required the first time.
$ make download_kernel
$ make test
```

## Which projects use this library?

- [vfkit](https://github.com/crc-org/vfkit) is a macOS command-line hypervisor for Apple and Intel CPUs that supports most of Apple's Virtualization Framework features.
- [Lima](https://lima-vm.io/) launches Linux virtual machines with automatic file sharing and port forwarding (similar to WSL2).
- [linuxkit](https://github.com/linuxkit/linuxkit) is a toolkit for building custom minimal, immutable Linux distributions.

## LICENSE

MIT License
