// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MlApplicationInstanceView Representation of ML Application Instance which providers use for instance observability.
type MlApplicationInstanceView struct {

	// The OCID of the MlApplicationInstanceView. Unique identifier that is immutable after creation
	Id *string `mandatory:"true" json:"id"`

	// The name of the MlApplicationInstance (created by the consumer) which this MlApplicationInstanceView is mirroring.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the MlApplicationInstance (created by the consumer) which this MlApplicationInstanceView is mirroring.
	MlApplicationInstanceId *string `mandatory:"true" json:"mlApplicationInstanceId"`

	// This field is a copy from MlApplicationInstance created by the consumer. The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId).
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// This field is a copy from MlApplicationInstance created by the consumer. The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)
	MlApplicationImplementationId *string `mandatory:"true" json:"mlApplicationImplementationId"`

	// This field is a copy from MlApplicationInstance created by the consumer. The name of Ml Application Implemenation (based on mlApplicationImplementationId)
	MlApplicationImplementationName *string `mandatory:"true" json:"mlApplicationImplementationName"`

	// Version of MlApplication package which is currently used by this MlApplicationInstance.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// The OCID of the MlApplicationImplementationVersion
	MlApplicationImplementationVersionId *string `mandatory:"true" json:"mlApplicationImplementationVersionId"`

	// This field is a copy from MlApplicationInstance created by the consumer. States whether the MlApplicationInstance is supposed to be in ACTIVE lifecycle state.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID of the compartment where the MlApplicationInstanceView is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the MlApplicationInstanceView was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MlApplicationInstance(View).
	LifecycleState MlApplicationInstanceViewLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current substate of the MlApplicationInstance(View). The substate has MlApplicationInstance(View) specific values in comparison with lifecycleState which has standard values common for all OCI resources.
	// The NEEDS_ATTENTION and FAILED substates are deprecated in favor of (NON_)?RECOVERABLE_(PROVIDER|SERVICE)_ISSUE and will be removed in next release.
	LifecycleSubstate MlApplicationInstanceViewLifecycleSubstateEnum `mandatory:"true" json:"lifecycleSubstate"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	AuthConfiguration AuthConfiguration `mandatory:"false" json:"authConfiguration"`

	// This field is a copy from MlApplicationInstance created by the consumer. Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplication.
	Configuration []ConfigurationProperty `mandatory:"false" json:"configuration"`

	// References (Identifiers) for components dedicated to this instance.
	InstanceComponents []InstanceComponent `mandatory:"false" json:"instanceComponents"`

	PredictionEndpointDetails *PredictionEndpointDetails `mandatory:"false" json:"predictionEndpointDetails"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationInstanceView) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationInstanceView) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationInstanceViewLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationInstanceViewLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMlApplicationInstanceViewLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetMlApplicationInstanceViewLifecycleSubstateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplicationInstanceView) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AuthConfiguration                    authconfiguration                              `json:"authConfiguration"`
		Configuration                        []ConfigurationProperty                        `json:"configuration"`
		InstanceComponents                   []instancecomponent                            `json:"instanceComponents"`
		PredictionEndpointDetails            *PredictionEndpointDetails                     `json:"predictionEndpointDetails"`
		SystemTags                           map[string]map[string]interface{}              `json:"systemTags"`
		Id                                   *string                                        `json:"id"`
		DisplayName                          *string                                        `json:"displayName"`
		MlApplicationInstanceId              *string                                        `json:"mlApplicationInstanceId"`
		MlApplicationId                      *string                                        `json:"mlApplicationId"`
		MlApplicationName                    *string                                        `json:"mlApplicationName"`
		MlApplicationImplementationId        *string                                        `json:"mlApplicationImplementationId"`
		MlApplicationImplementationName      *string                                        `json:"mlApplicationImplementationName"`
		PackageVersion                       *string                                        `json:"packageVersion"`
		MlApplicationImplementationVersionId *string                                        `json:"mlApplicationImplementationVersionId"`
		IsEnabled                            *bool                                          `json:"isEnabled"`
		CompartmentId                        *string                                        `json:"compartmentId"`
		TimeCreated                          *common.SDKTime                                `json:"timeCreated"`
		LifecycleState                       MlApplicationInstanceViewLifecycleStateEnum    `json:"lifecycleState"`
		LifecycleSubstate                    MlApplicationInstanceViewLifecycleSubstateEnum `json:"lifecycleSubstate"`
		LifecycleDetails                     *string                                        `json:"lifecycleDetails"`
		FreeformTags                         map[string]string                              `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{}              `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.AuthConfiguration.UnmarshalPolymorphicJSON(model.AuthConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthConfiguration = nn.(AuthConfiguration)
	} else {
		m.AuthConfiguration = nil
	}

	m.Configuration = make([]ConfigurationProperty, len(model.Configuration))
	copy(m.Configuration, model.Configuration)
	m.InstanceComponents = make([]InstanceComponent, len(model.InstanceComponents))
	for i, n := range model.InstanceComponents {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.InstanceComponents[i] = nn.(InstanceComponent)
		} else {
			m.InstanceComponents[i] = nil
		}
	}
	m.PredictionEndpointDetails = model.PredictionEndpointDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.MlApplicationInstanceId = model.MlApplicationInstanceId

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationName = model.MlApplicationName

	m.MlApplicationImplementationId = model.MlApplicationImplementationId

	m.MlApplicationImplementationName = model.MlApplicationImplementationName

	m.PackageVersion = model.PackageVersion

	m.MlApplicationImplementationVersionId = model.MlApplicationImplementationVersionId

	m.IsEnabled = model.IsEnabled

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.LifecycleSubstate = model.LifecycleSubstate

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationInstanceViewLifecycleStateEnum Enum with underlying type: string
type MlApplicationInstanceViewLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationInstanceViewLifecycleStateEnum
const (
	MlApplicationInstanceViewLifecycleStateCreating       MlApplicationInstanceViewLifecycleStateEnum = "CREATING"
	MlApplicationInstanceViewLifecycleStateUpdating       MlApplicationInstanceViewLifecycleStateEnum = "UPDATING"
	MlApplicationInstanceViewLifecycleStateActive         MlApplicationInstanceViewLifecycleStateEnum = "ACTIVE"
	MlApplicationInstanceViewLifecycleStateInactive       MlApplicationInstanceViewLifecycleStateEnum = "INACTIVE"
	MlApplicationInstanceViewLifecycleStateDeleting       MlApplicationInstanceViewLifecycleStateEnum = "DELETING"
	MlApplicationInstanceViewLifecycleStateDeleted        MlApplicationInstanceViewLifecycleStateEnum = "DELETED"
	MlApplicationInstanceViewLifecycleStateNeedsAttention MlApplicationInstanceViewLifecycleStateEnum = "NEEDS_ATTENTION"
	MlApplicationInstanceViewLifecycleStateFailed         MlApplicationInstanceViewLifecycleStateEnum = "FAILED"
)

var mappingMlApplicationInstanceViewLifecycleStateEnum = map[string]MlApplicationInstanceViewLifecycleStateEnum{
	"CREATING":        MlApplicationInstanceViewLifecycleStateCreating,
	"UPDATING":        MlApplicationInstanceViewLifecycleStateUpdating,
	"ACTIVE":          MlApplicationInstanceViewLifecycleStateActive,
	"INACTIVE":        MlApplicationInstanceViewLifecycleStateInactive,
	"DELETING":        MlApplicationInstanceViewLifecycleStateDeleting,
	"DELETED":         MlApplicationInstanceViewLifecycleStateDeleted,
	"NEEDS_ATTENTION": MlApplicationInstanceViewLifecycleStateNeedsAttention,
	"FAILED":          MlApplicationInstanceViewLifecycleStateFailed,
}

var mappingMlApplicationInstanceViewLifecycleStateEnumLowerCase = map[string]MlApplicationInstanceViewLifecycleStateEnum{
	"creating":        MlApplicationInstanceViewLifecycleStateCreating,
	"updating":        MlApplicationInstanceViewLifecycleStateUpdating,
	"active":          MlApplicationInstanceViewLifecycleStateActive,
	"inactive":        MlApplicationInstanceViewLifecycleStateInactive,
	"deleting":        MlApplicationInstanceViewLifecycleStateDeleting,
	"deleted":         MlApplicationInstanceViewLifecycleStateDeleted,
	"needs_attention": MlApplicationInstanceViewLifecycleStateNeedsAttention,
	"failed":          MlApplicationInstanceViewLifecycleStateFailed,
}

// GetMlApplicationInstanceViewLifecycleStateEnumValues Enumerates the set of values for MlApplicationInstanceViewLifecycleStateEnum
func GetMlApplicationInstanceViewLifecycleStateEnumValues() []MlApplicationInstanceViewLifecycleStateEnum {
	values := make([]MlApplicationInstanceViewLifecycleStateEnum, 0)
	for _, v := range mappingMlApplicationInstanceViewLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationInstanceViewLifecycleStateEnumStringValues Enumerates the set of values in String for MlApplicationInstanceViewLifecycleStateEnum
func GetMlApplicationInstanceViewLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingMlApplicationInstanceViewLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationInstanceViewLifecycleStateEnum(val string) (MlApplicationInstanceViewLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationInstanceViewLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MlApplicationInstanceViewLifecycleSubstateEnum Enum with underlying type: string
type MlApplicationInstanceViewLifecycleSubstateEnum string

// Set of constants representing the allowable values for MlApplicationInstanceViewLifecycleSubstateEnum
const (
	MlApplicationInstanceViewLifecycleSubstateCreating                    MlApplicationInstanceViewLifecycleSubstateEnum = "CREATING"
	MlApplicationInstanceViewLifecycleSubstateUpdating                    MlApplicationInstanceViewLifecycleSubstateEnum = "UPDATING"
	MlApplicationInstanceViewLifecycleSubstateUpgrading                   MlApplicationInstanceViewLifecycleSubstateEnum = "UPGRADING"
	MlApplicationInstanceViewLifecycleSubstateActive                      MlApplicationInstanceViewLifecycleSubstateEnum = "ACTIVE"
	MlApplicationInstanceViewLifecycleSubstateInactive                    MlApplicationInstanceViewLifecycleSubstateEnum = "INACTIVE"
	MlApplicationInstanceViewLifecycleSubstateDeleting                    MlApplicationInstanceViewLifecycleSubstateEnum = "DELETING"
	MlApplicationInstanceViewLifecycleSubstateDeleted                     MlApplicationInstanceViewLifecycleSubstateEnum = "DELETED"
	MlApplicationInstanceViewLifecycleSubstateNeedsAttention              MlApplicationInstanceViewLifecycleSubstateEnum = "NEEDS_ATTENTION"
	MlApplicationInstanceViewLifecycleSubstateFailed                      MlApplicationInstanceViewLifecycleSubstateEnum = "FAILED"
	MlApplicationInstanceViewLifecycleSubstateNonRecoverableProviderIssue MlApplicationInstanceViewLifecycleSubstateEnum = "NON_RECOVERABLE_PROVIDER_ISSUE"
	MlApplicationInstanceViewLifecycleSubstateRecoverableProviderIssue    MlApplicationInstanceViewLifecycleSubstateEnum = "RECOVERABLE_PROVIDER_ISSUE"
	MlApplicationInstanceViewLifecycleSubstateNonRecoverableServiceIssue  MlApplicationInstanceViewLifecycleSubstateEnum = "NON_RECOVERABLE_SERVICE_ISSUE"
	MlApplicationInstanceViewLifecycleSubstateRecoverableServiceIssue     MlApplicationInstanceViewLifecycleSubstateEnum = "RECOVERABLE_SERVICE_ISSUE"
)

var mappingMlApplicationInstanceViewLifecycleSubstateEnum = map[string]MlApplicationInstanceViewLifecycleSubstateEnum{
	"CREATING":                       MlApplicationInstanceViewLifecycleSubstateCreating,
	"UPDATING":                       MlApplicationInstanceViewLifecycleSubstateUpdating,
	"UPGRADING":                      MlApplicationInstanceViewLifecycleSubstateUpgrading,
	"ACTIVE":                         MlApplicationInstanceViewLifecycleSubstateActive,
	"INACTIVE":                       MlApplicationInstanceViewLifecycleSubstateInactive,
	"DELETING":                       MlApplicationInstanceViewLifecycleSubstateDeleting,
	"DELETED":                        MlApplicationInstanceViewLifecycleSubstateDeleted,
	"NEEDS_ATTENTION":                MlApplicationInstanceViewLifecycleSubstateNeedsAttention,
	"FAILED":                         MlApplicationInstanceViewLifecycleSubstateFailed,
	"NON_RECOVERABLE_PROVIDER_ISSUE": MlApplicationInstanceViewLifecycleSubstateNonRecoverableProviderIssue,
	"RECOVERABLE_PROVIDER_ISSUE":     MlApplicationInstanceViewLifecycleSubstateRecoverableProviderIssue,
	"NON_RECOVERABLE_SERVICE_ISSUE":  MlApplicationInstanceViewLifecycleSubstateNonRecoverableServiceIssue,
	"RECOVERABLE_SERVICE_ISSUE":      MlApplicationInstanceViewLifecycleSubstateRecoverableServiceIssue,
}

var mappingMlApplicationInstanceViewLifecycleSubstateEnumLowerCase = map[string]MlApplicationInstanceViewLifecycleSubstateEnum{
	"creating":                       MlApplicationInstanceViewLifecycleSubstateCreating,
	"updating":                       MlApplicationInstanceViewLifecycleSubstateUpdating,
	"upgrading":                      MlApplicationInstanceViewLifecycleSubstateUpgrading,
	"active":                         MlApplicationInstanceViewLifecycleSubstateActive,
	"inactive":                       MlApplicationInstanceViewLifecycleSubstateInactive,
	"deleting":                       MlApplicationInstanceViewLifecycleSubstateDeleting,
	"deleted":                        MlApplicationInstanceViewLifecycleSubstateDeleted,
	"needs_attention":                MlApplicationInstanceViewLifecycleSubstateNeedsAttention,
	"failed":                         MlApplicationInstanceViewLifecycleSubstateFailed,
	"non_recoverable_provider_issue": MlApplicationInstanceViewLifecycleSubstateNonRecoverableProviderIssue,
	"recoverable_provider_issue":     MlApplicationInstanceViewLifecycleSubstateRecoverableProviderIssue,
	"non_recoverable_service_issue":  MlApplicationInstanceViewLifecycleSubstateNonRecoverableServiceIssue,
	"recoverable_service_issue":      MlApplicationInstanceViewLifecycleSubstateRecoverableServiceIssue,
}

// GetMlApplicationInstanceViewLifecycleSubstateEnumValues Enumerates the set of values for MlApplicationInstanceViewLifecycleSubstateEnum
func GetMlApplicationInstanceViewLifecycleSubstateEnumValues() []MlApplicationInstanceViewLifecycleSubstateEnum {
	values := make([]MlApplicationInstanceViewLifecycleSubstateEnum, 0)
	for _, v := range mappingMlApplicationInstanceViewLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationInstanceViewLifecycleSubstateEnumStringValues Enumerates the set of values in String for MlApplicationInstanceViewLifecycleSubstateEnum
func GetMlApplicationInstanceViewLifecycleSubstateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"UPGRADING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
		"NON_RECOVERABLE_PROVIDER_ISSUE",
		"RECOVERABLE_PROVIDER_ISSUE",
		"NON_RECOVERABLE_SERVICE_ISSUE",
		"RECOVERABLE_SERVICE_ISSUE",
	}
}

// GetMappingMlApplicationInstanceViewLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationInstanceViewLifecycleSubstateEnum(val string) (MlApplicationInstanceViewLifecycleSubstateEnum, bool) {
	enum, ok := mappingMlApplicationInstanceViewLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
