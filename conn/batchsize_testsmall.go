//go:build testsmallbatch

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 *
 * Modified by Coder for test environments to reduce memory usage.
 */

package conn

// IdealBatchSize is reduced for tests to minimize memory usage.
// In production, this is 128 (see batchsize_default.go), but tests don't need
// the full batch capacity and this significantly reduces memory overhead
// from wireguard buffer pools (~8MB per connection down to ~256KB).
const IdealBatchSize = 4
