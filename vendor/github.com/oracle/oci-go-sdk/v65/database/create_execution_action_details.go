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

// CreateExecutionActionDetails Request to create execution action resource.
type CreateExecutionActionDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the execution window resource the execution action belongs to.
	ExecutionWindowId *string `mandatory:"true" json:"executionWindowId"`

	// The action type of the execution action being performed
	ActionType CreateExecutionActionDetailsActionTypeEnum `mandatory:"true" json:"actionType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"false" json:"actionParams"`

	// List of action members of this execution action.
	ActionMembers []ExecutionActionMember `mandatory:"false" json:"actionMembers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateExecutionActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExecutionActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateExecutionActionDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetCreateExecutionActionDetailsActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExecutionActionDetailsActionTypeEnum Enum with underlying type: string
type CreateExecutionActionDetailsActionTypeEnum string

// Set of constants representing the allowable values for CreateExecutionActionDetailsActionTypeEnum
const (
	CreateExecutionActionDetailsActionTypeDbServerFullSoftwareUpdate      CreateExecutionActionDetailsActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	CreateExecutionActionDetailsActionTypeStorageServerFullSoftwareUpdate CreateExecutionActionDetailsActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	CreateExecutionActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate CreateExecutionActionDetailsActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingCreateExecutionActionDetailsActionTypeEnum = map[string]CreateExecutionActionDetailsActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      CreateExecutionActionDetailsActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": CreateExecutionActionDetailsActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": CreateExecutionActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingCreateExecutionActionDetailsActionTypeEnumLowerCase = map[string]CreateExecutionActionDetailsActionTypeEnum{
	"db_server_full_software_update":      CreateExecutionActionDetailsActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": CreateExecutionActionDetailsActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": CreateExecutionActionDetailsActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetCreateExecutionActionDetailsActionTypeEnumValues Enumerates the set of values for CreateExecutionActionDetailsActionTypeEnum
func GetCreateExecutionActionDetailsActionTypeEnumValues() []CreateExecutionActionDetailsActionTypeEnum {
	values := make([]CreateExecutionActionDetailsActionTypeEnum, 0)
	for _, v := range mappingCreateExecutionActionDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExecutionActionDetailsActionTypeEnumStringValues Enumerates the set of values in String for CreateExecutionActionDetailsActionTypeEnum
func GetCreateExecutionActionDetailsActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingCreateExecutionActionDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExecutionActionDetailsActionTypeEnum(val string) (CreateExecutionActionDetailsActionTypeEnum, bool) {
	enum, ok := mappingCreateExecutionActionDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
