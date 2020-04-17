// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
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
