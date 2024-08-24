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

// CreateScheduledActionDetails Request to create Scheduled Action resource.
type CreateScheduledActionDetails struct {

	// The type of the scheduled action being performed
	ActionType CreateScheduledActionDetailsActionTypeEnum `mandatory:"true" json:"actionType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
	SchedulingPlanId *string `mandatory:"true" json:"schedulingPlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
	SchedulingWindowId *string `mandatory:"true" json:"schedulingWindowId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"false" json:"actionParams"`

	// The list of action members in a scheduled action.
	ActionMembers []ActionMember `mandatory:"false" json:"actionMembers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduledActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduledActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateScheduledActionDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetCreateScheduledActionDetailsActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateScheduledActionDetailsActionTypeEnum Enum with underlying type: string
type CreateScheduledActionDetailsActionTypeEnum string

// Set of constants representing the allowable values for CreateScheduledActionDetailsActionTypeEnum
const (
	CreateScheduledActionDetailsActionTypeDbServerFullSoftwareUpdate      CreateScheduledActionDetailsActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	CreateScheduledActionDetailsActionTypeStorageServerFullSoftwareUpdate CreateScheduledActionDetailsActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	CreateScheduledActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate CreateScheduledActionDetailsActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingCreateScheduledActionDetailsActionTypeEnum = map[string]CreateScheduledActionDetailsActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      CreateScheduledActionDetailsActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": CreateScheduledActionDetailsActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": CreateScheduledActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingCreateScheduledActionDetailsActionTypeEnumLowerCase = map[string]CreateScheduledActionDetailsActionTypeEnum{
	"db_server_full_software_update":      CreateScheduledActionDetailsActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": CreateScheduledActionDetailsActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": CreateScheduledActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetCreateScheduledActionDetailsActionTypeEnumValues Enumerates the set of values for CreateScheduledActionDetailsActionTypeEnum
func GetCreateScheduledActionDetailsActionTypeEnumValues() []CreateScheduledActionDetailsActionTypeEnum {
	values := make([]CreateScheduledActionDetailsActionTypeEnum, 0)
	for _, v := range mappingCreateScheduledActionDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduledActionDetailsActionTypeEnumStringValues Enumerates the set of values in String for CreateScheduledActionDetailsActionTypeEnum
func GetCreateScheduledActionDetailsActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingCreateScheduledActionDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduledActionDetailsActionTypeEnum(val string) (CreateScheduledActionDetailsActionTypeEnum, bool) {
	enum, ok := mappingCreateScheduledActionDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
