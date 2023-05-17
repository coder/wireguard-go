module github.com/tailscale/wireguard-go

go 1.20

replace golang.zx2c4.com/wireguard => github.com/coder/wireguard-go v0.0.0-20230517155151-b078c3c3adac

require (
	golang.org/x/crypto v0.6.0
	golang.org/x/net v0.7.0
	golang.org/x/sys v0.5.1-0.20230222185716-a3b23cc77e89
	golang.zx2c4.com/wintun v0.0.0-20230126152724-0fa3db229ce2
	golang.zx2c4.com/wireguard v0.0.0-20230325221338-052af4a8072b
	gvisor.dev/gvisor v0.0.0-20230328175328-162ed5ef888d
)

require (
	github.com/google/btree v1.0.1 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
)
