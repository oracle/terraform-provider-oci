// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDbSystemConnectorSummary The summary of an external DB system connector.
type ExternalDbSystemConnectorSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system connector.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the external connector. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of connector.
	ConnectorType ExternalDbSystemConnectorSummaryConnectorTypeEnum `mandatory:"true" json:"connectorType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the connector is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the external DB system connector.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The current lifecycle state of the external DB system connector.
	LifecycleState ExternalDbSystemConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB system connector was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB system connector was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExternalDbSystemConnectorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbSystemConnectorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbSystemConnectorSummaryConnectorTypeEnum(string(m.ConnectorType)); !ok && m.ConnectorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectorType: %s. Supported values are: %s.", m.ConnectorType, strings.Join(GetExternalDbSystemConnectorSummaryConnectorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExternalDbSystemConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbSystemConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDbSystemConnectorSummaryConnectorTypeEnum Enum with underlying type: string
type ExternalDbSystemConnectorSummaryConnectorTypeEnum string

// Set of constants representing the allowable values for ExternalDbSystemConnectorSummaryConnectorTypeEnum
const (
	ExternalDbSystemConnectorSummaryConnectorTypeMacs ExternalDbSystemConnectorSummaryConnectorTypeEnum = "MACS"
)

var mappingExternalDbSystemConnectorSummaryConnectorTypeEnum = map[string]ExternalDbSystemConnectorSummaryConnectorTypeEnum{
	"MACS": ExternalDbSystemConnectorSummaryConnectorTypeMacs,
}

var mappingExternalDbSystemConnectorSummaryConnectorTypeEnumLowerCase = map[string]ExternalDbSystemConnectorSummaryConnectorTypeEnum{
	"macs": ExternalDbSystemConnectorSummaryConnectorTypeMacs,
}

// GetExternalDbSystemConnectorSummaryConnectorTypeEnumValues Enumerates the set of values for ExternalDbSystemConnectorSummaryConnectorTypeEnum
func GetExternalDbSystemConnectorSummaryConnectorTypeEnumValues() []ExternalDbSystemConnectorSummaryConnectorTypeEnum {
	values := make([]ExternalDbSystemConnectorSummaryConnectorTypeEnum, 0)
	for _, v := range mappingExternalDbSystemConnectorSummaryConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemConnectorSummaryConnectorTypeEnumStringValues Enumerates the set of values in String for ExternalDbSystemConnectorSummaryConnectorTypeEnum
func GetExternalDbSystemConnectorSummaryConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingExternalDbSystemConnectorSummaryConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemConnectorSummaryConnectorTypeEnum(val string) (ExternalDbSystemConnectorSummaryConnectorTypeEnum, bool) {
	enum, ok := mappingExternalDbSystemConnectorSummaryConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
