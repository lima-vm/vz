module github.com/lima-vm/vz/example/linux

go 1.19

replace github.com/lima-vm/vz/v4 => ../../

require (
	github.com/lima-vm/vz/v4 v4.0.0-00010101000000-000000000000
	github.com/pkg/term v1.1.0
	golang.org/x/sys v0.16.0
)

require (
	github.com/Code-Hex/go-infinity-channel v1.0.0 // indirect
	golang.org/x/mod v0.6.0 // indirect
)
