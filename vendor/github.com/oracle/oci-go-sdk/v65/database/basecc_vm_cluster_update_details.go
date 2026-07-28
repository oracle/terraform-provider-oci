// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseccVmClusterUpdateDetails Details specifying which maintenance update to apply to the BaseDB-C@C VM cluster and which action is to be performed by the maintenance update. Applies to Base Database Service on Cloud@Customer instances only.
type BaseccVmClusterUpdateDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"false" json:"updateId"`

	// The update action to perform.
	UpdateAction BaseccVmClusterUpdateDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a grid infrastructure software image. This is a database software image of the type `GRID_IMAGE`.
	GiSoftwareImageId *string `mandatory:"false" json:"giSoftwareImageId"`
}

func (m BaseccVmClusterUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBaseccVmClusterUpdateDetailsUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetBaseccVmClusterUpdateDetailsUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterUpdateDetailsUpdateActionEnum Enum with underlying type: string
type BaseccVmClusterUpdateDetailsUpdateActionEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateDetailsUpdateActionEnum
const (
	BaseccVmClusterUpdateDetailsUpdateActionRollingApply BaseccVmClusterUpdateDetailsUpdateActionEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateDetailsUpdateActionPrecheck     BaseccVmClusterUpdateDetailsUpdateActionEnum = "PRECHECK"
	BaseccVmClusterUpdateDetailsUpdateActionRollback     BaseccVmClusterUpdateDetailsUpdateActionEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateDetailsUpdateActionEnum = map[string]BaseccVmClusterUpdateDetailsUpdateActionEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateDetailsUpdateActionRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateDetailsUpdateActionPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateDetailsUpdateActionRollback,
}

var mappingBaseccVmClusterUpdateDetailsUpdateActionEnumLowerCase = map[string]BaseccVmClusterUpdateDetailsUpdateActionEnum{
	"rolling_apply": BaseccVmClusterUpdateDetailsUpdateActionRollingApply,
	"precheck":      BaseccVmClusterUpdateDetailsUpdateActionPrecheck,
	"rollback":      BaseccVmClusterUpdateDetailsUpdateActionRollback,
}

// GetBaseccVmClusterUpdateDetailsUpdateActionEnumValues Enumerates the set of values for BaseccVmClusterUpdateDetailsUpdateActionEnum
func GetBaseccVmClusterUpdateDetailsUpdateActionEnumValues() []BaseccVmClusterUpdateDetailsUpdateActionEnum {
	values := make([]BaseccVmClusterUpdateDetailsUpdateActionEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateDetailsUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateDetailsUpdateActionEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateDetailsUpdateActionEnum
func GetBaseccVmClusterUpdateDetailsUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateDetailsUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateDetailsUpdateActionEnum(val string) (BaseccVmClusterUpdateDetailsUpdateActionEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateDetailsUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
