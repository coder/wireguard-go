//go:build testsmallbuffers

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 *
 * Modified by Coder for test environments to reduce memory usage.
 */

package device

import "github.com/tailscale/wireguard-go/conn"

const (
	QueueStagedSize            = conn.IdealBatchSize
	QueueOutboundSize          = 64   // Reduced from 1024 for tests
	QueueInboundSize           = 64   // Reduced from 1024 for tests
	QueueHandshakeSize         = 64   // Reduced from 1024 for tests
	MaxSegmentSize             = 2048 // Reduced from 65535 for tests (enough for typical packets)
	PreallocatedBuffersPerPool = 0    // Disable and allow for infinite memory growth
)
