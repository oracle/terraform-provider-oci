// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Process Automation
//
// Process Automation helps you to rapidly design, automate, and manage business processes in the cloud. With the Process Automation design-time (Designer) and the runtime (Workspace) environments, you can easily create, develop, manage, test, and monitor process applications and their components.
//

package opa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpaInstanceSummary Summary of the OpaInstance.
type OpaInstanceSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OpaInstance Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape of the instance.
	ShapeName OpaInstanceShapeNameEnum `mandatory:"true" json:"shapeName"`

	// The time the the OpaInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the OpaInstance.
	LifecycleState OpaInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the Process Automation instance.
	Description *string `mandatory:"false" json:"description"`

	// OPA Instance URL
	InstanceUrl *string `mandatory:"false" json:"instanceUrl"`

	// Parameter specifying which entitlement to use for billing purposes
	ConsumptionModel OpaInstanceConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	// MeteringType Identifier
	MeteringType OpaInstanceMeteringTypeEnum `mandatory:"false" json:"meteringType,omitempty"`

	// The time the OpaInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// indicates if breakGlass is enabled for the opa instance.
	IsBreakglassEnabled *bool `mandatory:"false" json:"isBreakglassEnabled"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OpaInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpaInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpaInstanceShapeNameEnum(string(m.ShapeName)); !ok && m.ShapeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeName: %s. Supported values are: %s.", m.ShapeName, strings.Join(GetOpaInstanceShapeNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOpaInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpaInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOpaInstanceConsumptionModelEnum(string(m.ConsumptionModel)); !ok && m.ConsumptionModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumptionModel: %s. Supported values are: %s.", m.ConsumptionModel, strings.Join(GetOpaInstanceConsumptionModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOpaInstanceMeteringTypeEnum(string(m.MeteringType)); !ok && m.MeteringType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MeteringType: %s. Supported values are: %s.", m.MeteringType, strings.Join(GetOpaInstanceMeteringTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
