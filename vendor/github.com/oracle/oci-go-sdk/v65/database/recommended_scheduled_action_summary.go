// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecommendedScheduledActionSummary Details of the scheduled action that is used in getRecommendedScheduledActions.
type RecommendedScheduledActionSummary struct {

	// The order of the scheduled action.
	ActionOrder *int `mandatory:"true" json:"actionOrder"`

	// The type of the scheduled action being performed
	ActionType RecommendedScheduledActionSummaryActionTypeEnum `mandatory:"true" json:"actionType"`

	// The id of the scheduling window this scheduled action belongs to.
	SchedulingWindowId *string `mandatory:"true" json:"schedulingWindowId"`

	// Description of the scheduled action being performed, i.e. apply full update to DB Servers 1,2,3,4.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The estimated patching time in minutes for the entire scheduled action.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// The list of action members in a scheduled action.
	ActionMembers []ActionMember `mandatory:"false" json:"actionMembers"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"false" json:"actionParams"`
}

func (m RecommendedScheduledActionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendedScheduledActionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecommendedScheduledActionSummaryActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetRecommendedScheduledActionSummaryActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecommendedScheduledActionSummaryActionTypeEnum Enum with underlying type: string
type RecommendedScheduledActionSummaryActionTypeEnum string

// Set of constants representing the allowable values for RecommendedScheduledActionSummaryActionTypeEnum
const (
	RecommendedScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate      RecommendedScheduledActionSummaryActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	RecommendedScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate RecommendedScheduledActionSummaryActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	RecommendedScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate RecommendedScheduledActionSummaryActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingRecommendedScheduledActionSummaryActionTypeEnum = map[string]RecommendedScheduledActionSummaryActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      RecommendedScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": RecommendedScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": RecommendedScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingRecommendedScheduledActionSummaryActionTypeEnumLowerCase = map[string]RecommendedScheduledActionSummaryActionTypeEnum{
	"db_server_full_software_update":      RecommendedScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": RecommendedScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": RecommendedScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetRecommendedScheduledActionSummaryActionTypeEnumValues Enumerates the set of values for RecommendedScheduledActionSummaryActionTypeEnum
func GetRecommendedScheduledActionSummaryActionTypeEnumValues() []RecommendedScheduledActionSummaryActionTypeEnum {
	values := make([]RecommendedScheduledActionSummaryActionTypeEnum, 0)
	for _, v := range mappingRecommendedScheduledActionSummaryActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendedScheduledActionSummaryActionTypeEnumStringValues Enumerates the set of values in String for RecommendedScheduledActionSummaryActionTypeEnum
func GetRecommendedScheduledActionSummaryActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingRecommendedScheduledActionSummaryActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendedScheduledActionSummaryActionTypeEnum(val string) (RecommendedScheduledActionSummaryActionTypeEnum, bool) {
	enum, ok := mappingRecommendedScheduledActionSummaryActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
