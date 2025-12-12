//go:build !testsmallbatch

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

const (
	IdealBatchSize = 128 // maximum number of packets handled per read and write
)
