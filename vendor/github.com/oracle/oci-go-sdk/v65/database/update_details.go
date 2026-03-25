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

// UpdateDetails Details specifying which maintenance update to apply to the cloud VM cluster and which actions are to be performed by the maintenance update. Applies to Exadata Cloud Service instances only.
type UpdateDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"false" json:"updateId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Grid Infrastructure Home.
	// Specify this field for out of place Grid Infrastructure Home patching and upgrade of the Cloud VM Cluster.
	// This is mutually exclusive option to `updateId` and `giSoftwareImageId` which are used for in place patching and upgrade using Oracle supplied and custom images respectively.
	GiHomeId *string `mandatory:"false" json:"giHomeId"`

	// The update action.
	UpdateAction UpdateDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// The update mode to perform for OS Update.
	UpdateMode UpdateDetailsUpdateModeEnum `mandatory:"false" json:"updateMode,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a grid infrastructure software image. This is a database software image of the type `GRID_IMAGE`.
	GiSoftwareImageId *string `mandatory:"false" json:"giSoftwareImageId"`
}

func (m UpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDetailsUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetUpdateDetailsUpdateActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDetailsUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetUpdateDetailsUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingUpdateDetailsUpdateActionEnum = map[string]UpdateDetailsUpdateActionEnum{
	"ROLLING_APPLY":     UpdateDetailsUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateDetailsUpdateActionNonRollingApply,
	"PRECHECK":          UpdateDetailsUpdateActionPrecheck,
	"ROLLBACK":          UpdateDetailsUpdateActionRollback,
}

var mappingUpdateDetailsUpdateActionEnumLowerCase = map[string]UpdateDetailsUpdateActionEnum{
	"rolling_apply":     UpdateDetailsUpdateActionRollingApply,
	"non_rolling_apply": UpdateDetailsUpdateActionNonRollingApply,
	"precheck":          UpdateDetailsUpdateActionPrecheck,
	"rollback":          UpdateDetailsUpdateActionRollback,
}

// GetUpdateDetailsUpdateActionEnumValues Enumerates the set of values for UpdateDetailsUpdateActionEnum
func GetUpdateDetailsUpdateActionEnumValues() []UpdateDetailsUpdateActionEnum {
	values := make([]UpdateDetailsUpdateActionEnum, 0)
	for _, v := range mappingUpdateDetailsUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDetailsUpdateActionEnumStringValues Enumerates the set of values in String for UpdateDetailsUpdateActionEnum
func GetUpdateDetailsUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateDetailsUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDetailsUpdateActionEnum(val string) (UpdateDetailsUpdateActionEnum, bool) {
	enum, ok := mappingUpdateDetailsUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDetailsUpdateModeEnum Enum with underlying type: string
type UpdateDetailsUpdateModeEnum string

// Set of constants representing the allowable values for UpdateDetailsUpdateModeEnum
const (
	UpdateDetailsUpdateModeOnlineHighcvss   UpdateDetailsUpdateModeEnum = "ONLINE_HIGHCVSS"
	UpdateDetailsUpdateModeOnlineAllcvss    UpdateDetailsUpdateModeEnum = "ONLINE_ALLCVSS"
	UpdateDetailsUpdateModeOnlineAllUpdates UpdateDetailsUpdateModeEnum = "ONLINE_ALL_UPDATES"
	UpdateDetailsUpdateModePendingUpdates   UpdateDetailsUpdateModeEnum = "PENDING_UPDATES"
	UpdateDetailsUpdateModeFullUpdate       UpdateDetailsUpdateModeEnum = "FULL_UPDATE"
)

var mappingUpdateDetailsUpdateModeEnum = map[string]UpdateDetailsUpdateModeEnum{
	"ONLINE_HIGHCVSS":    UpdateDetailsUpdateModeOnlineHighcvss,
	"ONLINE_ALLCVSS":     UpdateDetailsUpdateModeOnlineAllcvss,
	"ONLINE_ALL_UPDATES": UpdateDetailsUpdateModeOnlineAllUpdates,
	"PENDING_UPDATES":    UpdateDetailsUpdateModePendingUpdates,
	"FULL_UPDATE":        UpdateDetailsUpdateModeFullUpdate,
}

var mappingUpdateDetailsUpdateModeEnumLowerCase = map[string]UpdateDetailsUpdateModeEnum{
	"online_highcvss":    UpdateDetailsUpdateModeOnlineHighcvss,
	"online_allcvss":     UpdateDetailsUpdateModeOnlineAllcvss,
	"online_all_updates": UpdateDetailsUpdateModeOnlineAllUpdates,
	"pending_updates":    UpdateDetailsUpdateModePendingUpdates,
	"full_update":        UpdateDetailsUpdateModeFullUpdate,
}

// GetUpdateDetailsUpdateModeEnumValues Enumerates the set of values for UpdateDetailsUpdateModeEnum
func GetUpdateDetailsUpdateModeEnumValues() []UpdateDetailsUpdateModeEnum {
	values := make([]UpdateDetailsUpdateModeEnum, 0)
	for _, v := range mappingUpdateDetailsUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDetailsUpdateModeEnumStringValues Enumerates the set of values in String for UpdateDetailsUpdateModeEnum
func GetUpdateDetailsUpdateModeEnumStringValues() []string {
	return []string{
		"ONLINE_HIGHCVSS",
		"ONLINE_ALLCVSS",
		"ONLINE_ALL_UPDATES",
		"PENDING_UPDATES",
		"FULL_UPDATE",
	}
}

// GetMappingUpdateDetailsUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDetailsUpdateModeEnum(val string) (UpdateDetailsUpdateModeEnum, bool) {
	enum, ok := mappingUpdateDetailsUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
