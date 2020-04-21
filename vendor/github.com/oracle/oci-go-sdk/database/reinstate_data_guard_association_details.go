// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReinstateDataGuardAssociationDetails The Data Guard association reinstate parameters.
type ReinstateDataGuardAssociationDetails struct {

	// The DB system administrator password.
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`
}

func (m ReinstateDataGuardAssociationDetails) String() string {
	return common.PointerString(m)
}
