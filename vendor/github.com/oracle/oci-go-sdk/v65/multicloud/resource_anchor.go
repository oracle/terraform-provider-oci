// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceAnchor A ResourceAnchor is a description of a ResourceAnchor.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ResourceAnchor struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the ResourceAnchor was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ResourceAnchor.
	LifecycleState ResourceAnchorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Oracle Cloud Infrastructure Subscription Id
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// OCI Region that resource is created.
	Region *string `mandatory:"false" json:"region"`

	// The name assigned to the compartment during creation.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The date and time the ResourceAnchor was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the ResourceAnchor in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// AUTO_BIND - when passed compartment will be created on-behalf of customer and bind to this resource anchor
	// NO_AUTO_BIND - compartment will not be created and later customer can bind existing compartment.
	// to this resource anchor. This is for future use only
	SetupMode ResourceAnchorSetupModeEnum `mandatory:"false" json:"setupMode,omitempty"`

	// Optional - Oracle Cloud Infrastructure compartment Id (OCID) which was created or linked by customer with resource anchor.
	// This compartmentId is different from where resource Anchor live.
	LinkedCompartmentId *string `mandatory:"false" json:"linkedCompartmentId"`

	// The name assigned to the compartment which was created or linked by customer with resource anchor. This compartment is different from where resource Anchor live.
	LinkedCompartmentName *string `mandatory:"false" json:"linkedCompartmentName"`

	// subscription type
	SubscriptionType SubscriptionTypeEnum `mandatory:"false" json:"subscriptionType,omitempty"`

	CloudServiceProviderMetadataItem CloudServiceProviderMetadataItem `mandatory:"false" json:"cloudServiceProviderMetadataItem"`
}

func (m ResourceAnchor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceAnchor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAnchorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceAnchorLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingResourceAnchorSetupModeEnum(string(m.SetupMode)); !ok && m.SetupMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SetupMode: %s. Supported values are: %s.", m.SetupMode, strings.Join(GetResourceAnchorSetupModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionTypeEnum(string(m.SubscriptionType)); !ok && m.SubscriptionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionType: %s. Supported values are: %s.", m.SubscriptionType, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ResourceAnchor) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Region                           *string                           `json:"region"`
		CompartmentName                  *string                           `json:"compartmentName"`
		TimeUpdated                      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails                 *string                           `json:"lifecycleDetails"`
		SetupMode                        ResourceAnchorSetupModeEnum       `json:"setupMode"`
		LinkedCompartmentId              *string                           `json:"linkedCompartmentId"`
		LinkedCompartmentName            *string                           `json:"linkedCompartmentName"`
		SubscriptionType                 SubscriptionTypeEnum              `json:"subscriptionType"`
		CloudServiceProviderMetadataItem cloudserviceprovidermetadataitem  `json:"cloudServiceProviderMetadataItem"`
		Id                               *string                           `json:"id"`
		DisplayName                      *string                           `json:"displayName"`
		CompartmentId                    *string                           `json:"compartmentId"`
		TimeCreated                      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState                   ResourceAnchorLifecycleStateEnum  `json:"lifecycleState"`
		FreeformTags                     map[string]string                 `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                       map[string]map[string]interface{} `json:"systemTags"`
		SubscriptionId                   *string                           `json:"subscriptionId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Region = model.Region

	m.CompartmentName = model.CompartmentName

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.SetupMode = model.SetupMode

	m.LinkedCompartmentId = model.LinkedCompartmentId

	m.LinkedCompartmentName = model.LinkedCompartmentName

	m.SubscriptionType = model.SubscriptionType

	nn, e = model.CloudServiceProviderMetadataItem.UnmarshalPolymorphicJSON(model.CloudServiceProviderMetadataItem.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CloudServiceProviderMetadataItem = nn.(CloudServiceProviderMetadataItem)
	} else {
		m.CloudServiceProviderMetadataItem = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.SubscriptionId = model.SubscriptionId

	return
}

// ResourceAnchorLifecycleStateEnum Enum with underlying type: string
type ResourceAnchorLifecycleStateEnum string

// Set of constants representing the allowable values for ResourceAnchorLifecycleStateEnum
const (
	ResourceAnchorLifecycleStateCreating ResourceAnchorLifecycleStateEnum = "CREATING"
	ResourceAnchorLifecycleStateUpdating ResourceAnchorLifecycleStateEnum = "UPDATING"
	ResourceAnchorLifecycleStateActive   ResourceAnchorLifecycleStateEnum = "ACTIVE"
	ResourceAnchorLifecycleStateDeleting ResourceAnchorLifecycleStateEnum = "DELETING"
	ResourceAnchorLifecycleStateDeleted  ResourceAnchorLifecycleStateEnum = "DELETED"
	ResourceAnchorLifecycleStateFailed   ResourceAnchorLifecycleStateEnum = "FAILED"
)

var mappingResourceAnchorLifecycleStateEnum = map[string]ResourceAnchorLifecycleStateEnum{
	"CREATING": ResourceAnchorLifecycleStateCreating,
	"UPDATING": ResourceAnchorLifecycleStateUpdating,
	"ACTIVE":   ResourceAnchorLifecycleStateActive,
	"DELETING": ResourceAnchorLifecycleStateDeleting,
	"DELETED":  ResourceAnchorLifecycleStateDeleted,
	"FAILED":   ResourceAnchorLifecycleStateFailed,
}

var mappingResourceAnchorLifecycleStateEnumLowerCase = map[string]ResourceAnchorLifecycleStateEnum{
	"creating": ResourceAnchorLifecycleStateCreating,
	"updating": ResourceAnchorLifecycleStateUpdating,
	"active":   ResourceAnchorLifecycleStateActive,
	"deleting": ResourceAnchorLifecycleStateDeleting,
	"deleted":  ResourceAnchorLifecycleStateDeleted,
	"failed":   ResourceAnchorLifecycleStateFailed,
}

// GetResourceAnchorLifecycleStateEnumValues Enumerates the set of values for ResourceAnchorLifecycleStateEnum
func GetResourceAnchorLifecycleStateEnumValues() []ResourceAnchorLifecycleStateEnum {
	values := make([]ResourceAnchorLifecycleStateEnum, 0)
	for _, v := range mappingResourceAnchorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnchorLifecycleStateEnumStringValues Enumerates the set of values in String for ResourceAnchorLifecycleStateEnum
func GetResourceAnchorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingResourceAnchorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnchorLifecycleStateEnum(val string) (ResourceAnchorLifecycleStateEnum, bool) {
	enum, ok := mappingResourceAnchorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResourceAnchorSetupModeEnum Enum with underlying type: string
type ResourceAnchorSetupModeEnum string

// Set of constants representing the allowable values for ResourceAnchorSetupModeEnum
const (
	ResourceAnchorSetupModeAutoBind   ResourceAnchorSetupModeEnum = "AUTO_BIND"
	ResourceAnchorSetupModeNoAutoBind ResourceAnchorSetupModeEnum = "NO_AUTO_BIND"
)

var mappingResourceAnchorSetupModeEnum = map[string]ResourceAnchorSetupModeEnum{
	"AUTO_BIND":    ResourceAnchorSetupModeAutoBind,
	"NO_AUTO_BIND": ResourceAnchorSetupModeNoAutoBind,
}

var mappingResourceAnchorSetupModeEnumLowerCase = map[string]ResourceAnchorSetupModeEnum{
	"auto_bind":    ResourceAnchorSetupModeAutoBind,
	"no_auto_bind": ResourceAnchorSetupModeNoAutoBind,
}

// GetResourceAnchorSetupModeEnumValues Enumerates the set of values for ResourceAnchorSetupModeEnum
func GetResourceAnchorSetupModeEnumValues() []ResourceAnchorSetupModeEnum {
	values := make([]ResourceAnchorSetupModeEnum, 0)
	for _, v := range mappingResourceAnchorSetupModeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnchorSetupModeEnumStringValues Enumerates the set of values in String for ResourceAnchorSetupModeEnum
func GetResourceAnchorSetupModeEnumStringValues() []string {
	return []string{
		"AUTO_BIND",
		"NO_AUTO_BIND",
	}
}

// GetMappingResourceAnchorSetupModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnchorSetupModeEnum(val string) (ResourceAnchorSetupModeEnum, bool) {
	enum, ok := mappingResourceAnchorSetupModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
