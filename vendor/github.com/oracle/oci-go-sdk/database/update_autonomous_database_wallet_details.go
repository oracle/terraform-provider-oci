// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutonomousDatabaseWalletDetails Details to update an Autonomous Database wallet.
type UpdateAutonomousDatabaseWalletDetails struct {

	// Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate *bool `mandatory:"false" json:"shouldRotate"`
}

func (m UpdateAutonomousDatabaseWalletDetails) String() string {
	return common.PointerString(m)
}
