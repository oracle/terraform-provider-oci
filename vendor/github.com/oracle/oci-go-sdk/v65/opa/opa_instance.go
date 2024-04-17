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

// OpaInstance Description of OpaInstance.
type OpaInstance struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OpaInstance Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape of the instance.
	ShapeName OpaInstanceShapeNameEnum `mandatory:"true" json:"shapeName"`

	// The time when OpaInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the OpaInstance.
	LifecycleState OpaInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the Process Automation instance.
	Description *string `mandatory:"false" json:"description"`

	// OPA Instance URL
	InstanceUrl *string `mandatory:"false" json:"instanceUrl"`

	// The entitlement used for billing purposes
	ConsumptionModel OpaInstanceConsumptionModelEnum `mandatory:"false" json:"consumptionModel,omitempty"`

	// MeteringType Identifier
	MeteringType OpaInstanceMeteringTypeEnum `mandatory:"false" json:"meteringType,omitempty"`

	// The time the OpaInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This property specifies the GUID of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user role mappings to grant access to this OPA instance for users within the identity domain.
	IdentityAppGuid *string `mandatory:"false" json:"identityAppGuid"`

	// This property specifies the name of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
	IdentityAppDisplayName *string `mandatory:"false" json:"identityAppDisplayName"`

	// This property specifies the domain url of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
	IdentityDomainUrl *string `mandatory:"false" json:"identityDomainUrl"`

	// This property specifies the OPC Service Instance GUID of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
	IdentityAppOpcServiceInstanceGuid *string `mandatory:"false" json:"identityAppOpcServiceInstanceGuid"`

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

	// A list of associated attachments to other services
	Attachments []AttachmentDetails `mandatory:"false" json:"attachments"`
}

func (m OpaInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpaInstance) ValidateEnumValue() (bool, error) {
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

// OpaInstanceConsumptionModelEnum Enum with underlying type: string
type OpaInstanceConsumptionModelEnum string

// Set of constants representing the allowable values for OpaInstanceConsumptionModelEnum
const (
	OpaInstanceConsumptionModelUcm  OpaInstanceConsumptionModelEnum = "UCM"
	OpaInstanceConsumptionModelGov  OpaInstanceConsumptionModelEnum = "GOV"
	OpaInstanceConsumptionModelSaas OpaInstanceConsumptionModelEnum = "SAAS"
)

var mappingOpaInstanceConsumptionModelEnum = map[string]OpaInstanceConsumptionModelEnum{
	"UCM":  OpaInstanceConsumptionModelUcm,
	"GOV":  OpaInstanceConsumptionModelGov,
	"SAAS": OpaInstanceConsumptionModelSaas,
}

var mappingOpaInstanceConsumptionModelEnumLowerCase = map[string]OpaInstanceConsumptionModelEnum{
	"ucm":  OpaInstanceConsumptionModelUcm,
	"gov":  OpaInstanceConsumptionModelGov,
	"saas": OpaInstanceConsumptionModelSaas,
}

// GetOpaInstanceConsumptionModelEnumValues Enumerates the set of values for OpaInstanceConsumptionModelEnum
func GetOpaInstanceConsumptionModelEnumValues() []OpaInstanceConsumptionModelEnum {
	values := make([]OpaInstanceConsumptionModelEnum, 0)
	for _, v := range mappingOpaInstanceConsumptionModelEnum {
		values = append(values, v)
	}
	return values
}

// GetOpaInstanceConsumptionModelEnumStringValues Enumerates the set of values in String for OpaInstanceConsumptionModelEnum
func GetOpaInstanceConsumptionModelEnumStringValues() []string {
	return []string{
		"UCM",
		"GOV",
		"SAAS",
	}
}

// GetMappingOpaInstanceConsumptionModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpaInstanceConsumptionModelEnum(val string) (OpaInstanceConsumptionModelEnum, bool) {
	enum, ok := mappingOpaInstanceConsumptionModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OpaInstanceShapeNameEnum Enum with underlying type: string
type OpaInstanceShapeNameEnum string

// Set of constants representing the allowable values for OpaInstanceShapeNameEnum
const (
	OpaInstanceShapeNameDevelopment OpaInstanceShapeNameEnum = "DEVELOPMENT"
	OpaInstanceShapeNameProduction  OpaInstanceShapeNameEnum = "PRODUCTION"
)

var mappingOpaInstanceShapeNameEnum = map[string]OpaInstanceShapeNameEnum{
	"DEVELOPMENT": OpaInstanceShapeNameDevelopment,
	"PRODUCTION":  OpaInstanceShapeNameProduction,
}

var mappingOpaInstanceShapeNameEnumLowerCase = map[string]OpaInstanceShapeNameEnum{
	"development": OpaInstanceShapeNameDevelopment,
	"production":  OpaInstanceShapeNameProduction,
}

// GetOpaInstanceShapeNameEnumValues Enumerates the set of values for OpaInstanceShapeNameEnum
func GetOpaInstanceShapeNameEnumValues() []OpaInstanceShapeNameEnum {
	values := make([]OpaInstanceShapeNameEnum, 0)
	for _, v := range mappingOpaInstanceShapeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetOpaInstanceShapeNameEnumStringValues Enumerates the set of values in String for OpaInstanceShapeNameEnum
func GetOpaInstanceShapeNameEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingOpaInstanceShapeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpaInstanceShapeNameEnum(val string) (OpaInstanceShapeNameEnum, bool) {
	enum, ok := mappingOpaInstanceShapeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OpaInstanceMeteringTypeEnum Enum with underlying type: string
type OpaInstanceMeteringTypeEnum string

// Set of constants representing the allowable values for OpaInstanceMeteringTypeEnum
const (
	OpaInstanceMeteringTypeExecutionPack OpaInstanceMeteringTypeEnum = "EXECUTION_PACK"
	OpaInstanceMeteringTypeUsers         OpaInstanceMeteringTypeEnum = "USERS"
	OpaInstanceMeteringTypeEmployee      OpaInstanceMeteringTypeEnum = "EMPLOYEE"
	OpaInstanceMeteringTypeNamedUser     OpaInstanceMeteringTypeEnum = "NAMED_USER"
)

var mappingOpaInstanceMeteringTypeEnum = map[string]OpaInstanceMeteringTypeEnum{
	"EXECUTION_PACK": OpaInstanceMeteringTypeExecutionPack,
	"USERS":          OpaInstanceMeteringTypeUsers,
	"EMPLOYEE":       OpaInstanceMeteringTypeEmployee,
	"NAMED_USER":     OpaInstanceMeteringTypeNamedUser,
}

var mappingOpaInstanceMeteringTypeEnumLowerCase = map[string]OpaInstanceMeteringTypeEnum{
	"execution_pack": OpaInstanceMeteringTypeExecutionPack,
	"users":          OpaInstanceMeteringTypeUsers,
	"employee":       OpaInstanceMeteringTypeEmployee,
	"named_user":     OpaInstanceMeteringTypeNamedUser,
}

// GetOpaInstanceMeteringTypeEnumValues Enumerates the set of values for OpaInstanceMeteringTypeEnum
func GetOpaInstanceMeteringTypeEnumValues() []OpaInstanceMeteringTypeEnum {
	values := make([]OpaInstanceMeteringTypeEnum, 0)
	for _, v := range mappingOpaInstanceMeteringTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOpaInstanceMeteringTypeEnumStringValues Enumerates the set of values in String for OpaInstanceMeteringTypeEnum
func GetOpaInstanceMeteringTypeEnumStringValues() []string {
	return []string{
		"EXECUTION_PACK",
		"USERS",
		"EMPLOYEE",
		"NAMED_USER",
	}
}

// GetMappingOpaInstanceMeteringTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpaInstanceMeteringTypeEnum(val string) (OpaInstanceMeteringTypeEnum, bool) {
	enum, ok := mappingOpaInstanceMeteringTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OpaInstanceLifecycleStateEnum Enum with underlying type: string
type OpaInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OpaInstanceLifecycleStateEnum
const (
	OpaInstanceLifecycleStateCreating OpaInstanceLifecycleStateEnum = "CREATING"
	OpaInstanceLifecycleStateUpdating OpaInstanceLifecycleStateEnum = "UPDATING"
	OpaInstanceLifecycleStateActive   OpaInstanceLifecycleStateEnum = "ACTIVE"
	OpaInstanceLifecycleStateInactive OpaInstanceLifecycleStateEnum = "INACTIVE"
	OpaInstanceLifecycleStateDeleting OpaInstanceLifecycleStateEnum = "DELETING"
	OpaInstanceLifecycleStateDeleted  OpaInstanceLifecycleStateEnum = "DELETED"
	OpaInstanceLifecycleStateFailed   OpaInstanceLifecycleStateEnum = "FAILED"
)

var mappingOpaInstanceLifecycleStateEnum = map[string]OpaInstanceLifecycleStateEnum{
	"CREATING": OpaInstanceLifecycleStateCreating,
	"UPDATING": OpaInstanceLifecycleStateUpdating,
	"ACTIVE":   OpaInstanceLifecycleStateActive,
	"INACTIVE": OpaInstanceLifecycleStateInactive,
	"DELETING": OpaInstanceLifecycleStateDeleting,
	"DELETED":  OpaInstanceLifecycleStateDeleted,
	"FAILED":   OpaInstanceLifecycleStateFailed,
}

var mappingOpaInstanceLifecycleStateEnumLowerCase = map[string]OpaInstanceLifecycleStateEnum{
	"creating": OpaInstanceLifecycleStateCreating,
	"updating": OpaInstanceLifecycleStateUpdating,
	"active":   OpaInstanceLifecycleStateActive,
	"inactive": OpaInstanceLifecycleStateInactive,
	"deleting": OpaInstanceLifecycleStateDeleting,
	"deleted":  OpaInstanceLifecycleStateDeleted,
	"failed":   OpaInstanceLifecycleStateFailed,
}

// GetOpaInstanceLifecycleStateEnumValues Enumerates the set of values for OpaInstanceLifecycleStateEnum
func GetOpaInstanceLifecycleStateEnumValues() []OpaInstanceLifecycleStateEnum {
	values := make([]OpaInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOpaInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpaInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for OpaInstanceLifecycleStateEnum
func GetOpaInstanceLifecycleStateEnumStringValues() []string {
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

// GetMappingOpaInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpaInstanceLifecycleStateEnum(val string) (OpaInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingOpaInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
