// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Visual Builder API
//
// Oracle Visual Builder enables developers to quickly build web and mobile applications. With a visual development environment that makes it easy to connect to Oracle data and third-party REST services, developers can build modern, consumer-grade applications in a fraction of the time it would take in other tools.
// The Visual Builder Instance Management API allows users to create and manage a Visual Builder instance.
//

package visualbuilder

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VbInstance Description of Vb Instance.
type VbInstance struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Vb Instance Identifier, can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the vb instance.
	LifecycleState VbInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Vb Instance URL.
	InstanceUrl *string `mandatory:"true" json:"instanceUrl"`

	// The number of Nodes
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The time the the VbInstance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the VbInstance was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Visual Builder is enabled or not.
	IsVisualBuilderEnabled *bool `mandatory:"false" json:"isVisualBuilderEnabled"`

	CustomEndpoint *CustomEndpointDetails `mandatory:"false" json:"customEndpoint"`

	// A list of alternate custom endpoints used for the vb instance URL.
	AlternateCustomEndpoints []CustomEndpointDetails `mandatory:"false" json:"alternateCustomEndpoints"`

	// The entitlement used for billing purposes.
	ConsumptionModel VbInstanceConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	IdcsInfo *IdcsInfoDetails `mandatory:"false" json:"idcsInfo"`

	// A list of associated attachments to other services
	Attachments []AttachmentDetails `mandatory:"false" json:"attachments"`

	// The NAT gateway IP address for the VB service VCN
	ServiceNatGatewayIp *string `mandatory:"false" json:"serviceNatGatewayIp"`

	// The NAT gateway IP address for the VB management VCN
	ManagementNatGatewayIp *string `mandatory:"false" json:"managementNatGatewayIp"`

	// The Oracle Cloud ID (OCID) of the Visual Builder service VCN
	ServiceVcnId *string `mandatory:"false" json:"serviceVcnId"`

	// The Oracle Cloud ID (OCID) of the Visual Builder management VCN
	ManagementVcnId *string `mandatory:"false" json:"managementVcnId"`
}

func (m VbInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVbInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVbInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVbInstanceConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetVbInstanceConsumptionModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VbInstanceLifecycleStateEnum Enum with underlying type: string
type VbInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for VbInstanceLifecycleStateEnum
const (
	VbInstanceLifecycleStateCreating VbInstanceLifecycleStateEnum = "CREATING"
	VbInstanceLifecycleStateUpdating VbInstanceLifecycleStateEnum = "UPDATING"
	VbInstanceLifecycleStateActive   VbInstanceLifecycleStateEnum = "ACTIVE"
	VbInstanceLifecycleStateInactive VbInstanceLifecycleStateEnum = "INACTIVE"
	VbInstanceLifecycleStateDeleting VbInstanceLifecycleStateEnum = "DELETING"
	VbInstanceLifecycleStateDeleted  VbInstanceLifecycleStateEnum = "DELETED"
	VbInstanceLifecycleStateFailed   VbInstanceLifecycleStateEnum = "FAILED"
)

var mappingVbInstanceLifecycleStateEnum = map[string]VbInstanceLifecycleStateEnum{
	"CREATING": VbInstanceLifecycleStateCreating,
	"UPDATING": VbInstanceLifecycleStateUpdating,
	"ACTIVE":   VbInstanceLifecycleStateActive,
	"INACTIVE": VbInstanceLifecycleStateInactive,
	"DELETING": VbInstanceLifecycleStateDeleting,
	"DELETED":  VbInstanceLifecycleStateDeleted,
	"FAILED":   VbInstanceLifecycleStateFailed,
}

var mappingVbInstanceLifecycleStateEnumLowerCase = map[string]VbInstanceLifecycleStateEnum{
	"creating": VbInstanceLifecycleStateCreating,
	"updating": VbInstanceLifecycleStateUpdating,
	"active":   VbInstanceLifecycleStateActive,
	"inactive": VbInstanceLifecycleStateInactive,
	"deleting": VbInstanceLifecycleStateDeleting,
	"deleted":  VbInstanceLifecycleStateDeleted,
	"failed":   VbInstanceLifecycleStateFailed,
}

// GetVbInstanceLifecycleStateEnumValues Enumerates the set of values for VbInstanceLifecycleStateEnum
func GetVbInstanceLifecycleStateEnumValues() []VbInstanceLifecycleStateEnum {
	values := make([]VbInstanceLifecycleStateEnum, 0)
	for _, v := range mappingVbInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVbInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for VbInstanceLifecycleStateEnum
func GetVbInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVbInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVbInstanceLifecycleStateEnum(val string) (VbInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingVbInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VbInstanceConsumptionModelEnum Enum with underlying type: string
type VbInstanceConsumptionModelEnum string

// Set of constants representing the allowable values for VbInstanceConsumptionModelEnum
const (
	VbInstanceConsumptionModelUcm     VbInstanceConsumptionModelEnum = "UCM"
	VbInstanceConsumptionModelGov     VbInstanceConsumptionModelEnum = "GOV"
	VbInstanceConsumptionModelVb4saas VbInstanceConsumptionModelEnum = "VB4SAAS"
)

var mappingVbInstanceConsumptionModelEnum = map[string]VbInstanceConsumptionModelEnum{
	"UCM":     VbInstanceConsumptionModelUcm,
	"GOV":     VbInstanceConsumptionModelGov,
	"VB4SAAS": VbInstanceConsumptionModelVb4saas,
}

var mappingVbInstanceConsumptionModelEnumLowerCase = map[string]VbInstanceConsumptionModelEnum{
	"ucm":     VbInstanceConsumptionModelUcm,
	"gov":     VbInstanceConsumptionModelGov,
	"vb4saas": VbInstanceConsumptionModelVb4saas,
}

// GetVbInstanceConsumptionModelEnumValues Enumerates the set of values for VbInstanceConsumptionModelEnum
func GetVbInstanceConsumptionModelEnumValues() []VbInstanceConsumptionModelEnum {
	values := make([]VbInstanceConsumptionModelEnum, 0)
	for _, v := range mappingVbInstanceConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetVbInstanceConsumptionModelEnumStringValues Enumerates the set of values in String for VbInstanceConsumptionModelEnum
func GetVbInstanceConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"VB4SAAS",
	}
}

// GetMappingVbInstanceConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVbInstanceConsumptionModelEnum(val string) (VbInstanceConsumptionModelEnum, bool) {
	enum, ok := mappingVbInstanceConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
