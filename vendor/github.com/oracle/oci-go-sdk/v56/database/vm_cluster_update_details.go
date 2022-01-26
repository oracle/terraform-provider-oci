// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// VmClusterUpdateDetails Details specifying which maintenance update to apply to the VM Cluster and which action is to be performed by the maintenance update. Applies to Exadata Cloud@Customer instances only.
type VmClusterUpdateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"false" json:"updateId"`

	// The update action to perform.
	UpdateAction VmClusterUpdateDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`
}

func (m VmClusterUpdateDetails) String() string {
	return common.PointerString(m)
}

// VmClusterUpdateDetailsUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateDetailsUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateDetailsUpdateActionEnum
const (
	VmClusterUpdateDetailsUpdateActionRollingApply VmClusterUpdateDetailsUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateDetailsUpdateActionPrecheck     VmClusterUpdateDetailsUpdateActionEnum = "PRECHECK"
	VmClusterUpdateDetailsUpdateActionRollback     VmClusterUpdateDetailsUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateDetailsUpdateAction = map[string]VmClusterUpdateDetailsUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateDetailsUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateDetailsUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateDetailsUpdateActionRollback,
}

// GetVmClusterUpdateDetailsUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateDetailsUpdateActionEnum
func GetVmClusterUpdateDetailsUpdateActionEnumValues() []VmClusterUpdateDetailsUpdateActionEnum {
	values := make([]VmClusterUpdateDetailsUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateDetailsUpdateAction {
		values = append(values, v)
	}
	return values
}
