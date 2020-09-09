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

// PatchDetails The details about what actions to perform and using what patch to the specified target.
// This is part of an update request that is applied to a version field on the target such
// as DB system, Database Home, etc.
type PatchDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId *string `mandatory:"false" json:"patchId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database software image.
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// The action to perform on the patch.
	Action PatchDetailsActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m PatchDetails) String() string {
	return common.PointerString(m)
}

// PatchDetailsActionEnum Enum with underlying type: string
type PatchDetailsActionEnum string

// Set of constants representing the allowable values for PatchDetailsActionEnum
const (
	PatchDetailsActionApply    PatchDetailsActionEnum = "APPLY"
	PatchDetailsActionPrecheck PatchDetailsActionEnum = "PRECHECK"
)

var mappingPatchDetailsAction = map[string]PatchDetailsActionEnum{
	"APPLY":    PatchDetailsActionApply,
	"PRECHECK": PatchDetailsActionPrecheck,
}

// GetPatchDetailsActionEnumValues Enumerates the set of values for PatchDetailsActionEnum
func GetPatchDetailsActionEnumValues() []PatchDetailsActionEnum {
	values := make([]PatchDetailsActionEnum, 0)
	for _, v := range mappingPatchDetailsAction {
		values = append(values, v)
	}
	return values
}
