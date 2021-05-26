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

// UpdateDetails Details specifying which maintenance update to apply to the cloud VM cluster and which actions are to be performed by the maintenance update. Applies to Exadata Cloud Service instances only.
type UpdateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"false" json:"updateId"`

	// The update action.
	UpdateAction UpdateDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`
}

func (m UpdateDetails) String() string {
	return common.PointerString(m)
}

// UpdateDetailsUpdateActionEnum Enum with underlying type: string
type UpdateDetailsUpdateActionEnum string

// Set of constants representing the allowable values for UpdateDetailsUpdateActionEnum
const (
	UpdateDetailsUpdateActionRollingApply    UpdateDetailsUpdateActionEnum = "ROLLING_APPLY"
	UpdateDetailsUpdateActionNonRollingApply UpdateDetailsUpdateActionEnum = "NON_ROLLING_APPLY"
	UpdateDetailsUpdateActionPrecheck        UpdateDetailsUpdateActionEnum = "PRECHECK"
	UpdateDetailsUpdateActionRollback        UpdateDetailsUpdateActionEnum = "ROLLBACK"
)

var mappingUpdateDetailsUpdateAction = map[string]UpdateDetailsUpdateActionEnum{
	"ROLLING_APPLY":     UpdateDetailsUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateDetailsUpdateActionNonRollingApply,
	"PRECHECK":          UpdateDetailsUpdateActionPrecheck,
	"ROLLBACK":          UpdateDetailsUpdateActionRollback,
}

// GetUpdateDetailsUpdateActionEnumValues Enumerates the set of values for UpdateDetailsUpdateActionEnum
func GetUpdateDetailsUpdateActionEnumValues() []UpdateDetailsUpdateActionEnum {
	values := make([]UpdateDetailsUpdateActionEnum, 0)
	for _, v := range mappingUpdateDetailsUpdateAction {
		values = append(values, v)
	}
	return values
}
