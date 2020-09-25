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

// OnPremConnector A Data Safe on-premises connector that enables Data Safe to connect to on-premises databases.
type OnPremConnector struct {

	// The OCID of the on-premises connector.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the on-premises connector.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the on-premises connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the on-premises connector. Allowed values:
	// - CMAN - Represents Connection Manager.
	// - AGENT - Represents Management Agent.
	ConnectorType OnPremConnectorConnectorTypeEnum `mandatory:"true" json:"connectorType"`

	// The date and time the on-premises connector was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the on-premises connector.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the on-premises connector.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the management agent if on-premises connector's type is AGENT.
	AgentId *string `mandatory:"false" json:"agentId"`

	// Details about the current state of the on-premises connector.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OnPremConnector) String() string {
	return common.PointerString(m)
}

// OnPremConnectorConnectorTypeEnum Enum with underlying type: string
type OnPremConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for OnPremConnectorConnectorTypeEnum
const (
	OnPremConnectorConnectorTypeCman  OnPremConnectorConnectorTypeEnum = "CMAN"
	OnPremConnectorConnectorTypeAgent OnPremConnectorConnectorTypeEnum = "AGENT"
)

var mappingOnPremConnectorConnectorType = map[string]OnPremConnectorConnectorTypeEnum{
	"CMAN":  OnPremConnectorConnectorTypeCman,
	"AGENT": OnPremConnectorConnectorTypeAgent,
}

// GetOnPremConnectorConnectorTypeEnumValues Enumerates the set of values for OnPremConnectorConnectorTypeEnum
func GetOnPremConnectorConnectorTypeEnumValues() []OnPremConnectorConnectorTypeEnum {
	values := make([]OnPremConnectorConnectorTypeEnum, 0)
	for _, v := range mappingOnPremConnectorConnectorType {
		values = append(values, v)
	}
	return values
}
