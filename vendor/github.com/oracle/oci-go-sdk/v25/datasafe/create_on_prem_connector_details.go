// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateOnPremConnectorDetails The details used to create a new on-premises connector.
type CreateOnPremConnectorDetails struct {

	// The OCID of the compartment where you want to create the on-premises connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the on-premises connector. Allowed values:
	// - CMAN - Represents Connection Manager.
	// - AGENT - Represents Management Agent.
	ConnectorType CreateOnPremConnectorDetailsConnectorTypeEnum `mandatory:"true" json:"connectorType"`

	// The display name of the on-premises connector. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the management agent if on-premises connector's type is AGENT.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The description of the on-premises connector.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOnPremConnectorDetails) String() string {
	return common.PointerString(m)
}

// CreateOnPremConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type CreateOnPremConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for CreateOnPremConnectorDetailsConnectorTypeEnum
const (
	CreateOnPremConnectorDetailsConnectorTypeCman  CreateOnPremConnectorDetailsConnectorTypeEnum = "CMAN"
	CreateOnPremConnectorDetailsConnectorTypeAgent CreateOnPremConnectorDetailsConnectorTypeEnum = "AGENT"
)

var mappingCreateOnPremConnectorDetailsConnectorType = map[string]CreateOnPremConnectorDetailsConnectorTypeEnum{
	"CMAN":  CreateOnPremConnectorDetailsConnectorTypeCman,
	"AGENT": CreateOnPremConnectorDetailsConnectorTypeAgent,
}

// GetCreateOnPremConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for CreateOnPremConnectorDetailsConnectorTypeEnum
func GetCreateOnPremConnectorDetailsConnectorTypeEnumValues() []CreateOnPremConnectorDetailsConnectorTypeEnum {
	values := make([]CreateOnPremConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingCreateOnPremConnectorDetailsConnectorType {
		values = append(values, v)
	}
	return values
}
