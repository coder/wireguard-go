module github.com/tailscale/wireguard-go

go 1.23.1

replace gvisor.dev/gvisor => github.com/coder/gvisor v0.0.0-20260313191224-2e3ee1119d41

require (
	golang.org/x/crypto v0.28.0
	golang.org/x/net v0.30.0
	golang.org/x/sys v0.26.0
	golang.zx2c4.com/wintun v0.0.0-20230126152724-0fa3db229ce2
	gvisor.dev/gvisor v0.0.0-20250416071322-d75576f0af2b
)

require (
	github.com/google/btree v1.1.2 // indirect
	golang.org/x/time v0.7.0 // indirect
)
