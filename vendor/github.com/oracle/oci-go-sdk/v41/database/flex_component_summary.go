// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// FlexComponentSummary The Flex Components for a DB system. The Flex Component determines resources to allocate to the DB system -  CPU cores, memory and storage for Flex shapes.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type FlexComponentSummary struct {

	// The name of the Flex Component used for the DB system.
	Name *string `mandatory:"true" json:"name"`

	// The minimum number of CPU cores that can be enabled on the DB Server for this Flex Component.
	MinimumCoreCount *int `mandatory:"false" json:"minimumCoreCount"`

	// The maximum number of CPU cores that can ben enabled on the DB Server for this Flex Component.
	AvailableCoreCount *int `mandatory:"false" json:"availableCoreCount"`

	// The maximum  storage that can be enabled on the Storage Server for this Flex Component.
	AvailableDbStorageInGBs *int `mandatory:"false" json:"availableDbStorageInGBs"`
}

func (m FlexComponentSummary) String() string {
	return common.PointerString(m)
}
