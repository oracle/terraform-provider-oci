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

// MlApplicationInstance Resource representing instance of ML Application.
type MlApplicationInstance struct {

	// The OCID of the MlApplicationInstance. Unique identifier that is immutable after creation
	Id *string `mandatory:"true" json:"id"`

	// The name of MlApplicationInstance. System will generate displayName when not provided during creation.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId).
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// The OCID of ML Application Implementation selected as a certain solution for a given ML problem (ML Application)
	MlApplicationImplementationId *string `mandatory:"true" json:"mlApplicationImplementationId"`

	// The name of Ml Application Implementation (based on mlApplicationImplementationId)
	MlApplicationImplementationName *string `mandatory:"true" json:"mlApplicationImplementationName"`

	// States whether the MlApplicationInstance is supposed to be in ACTIVE lifecycle state.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID of the compartment where the MlApplicationInstance is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the MlApplication was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time of last MlApplicationInstance update in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the MlApplicationInstance.
	LifecycleState MlApplicationInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current substate of the MlApplicationInstance. The substate has MlApplicationInstance specific values in comparison with lifecycleState which has standard values common for all OCI resources.
	// The NEEDS_ATTENTION and FAILED substates are deprecated in favor of (NON_)?RECOVERABLE_(PROVIDER|SERVICE)_ISSUE and will be removed in next release.
	LifecycleSubstate MlApplicationInstanceLifecycleSubstateEnum `mandatory:"true" json:"lifecycleSubstate"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	AuthConfiguration AuthConfiguration `mandatory:"false" json:"authConfiguration"`

	// Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplicationImplementation.
	Configuration []ConfigurationProperty `mandatory:"false" json:"configuration"`

	PredictionEndpointDetails *PredictionEndpointDetails `mandatory:"false" json:"predictionEndpointDetails"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMlApplicationInstanceLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetMlApplicationInstanceLifecycleSubstateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplicationInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AuthConfiguration               authconfiguration                          `json:"authConfiguration"`
		Configuration                   []ConfigurationProperty                    `json:"configuration"`
		PredictionEndpointDetails       *PredictionEndpointDetails                 `json:"predictionEndpointDetails"`
		SystemTags                      map[string]map[string]interface{}          `json:"systemTags"`
		Id                              *string                                    `json:"id"`
		DisplayName                     *string                                    `json:"displayName"`
		MlApplicationId                 *string                                    `json:"mlApplicationId"`
		MlApplicationName               *string                                    `json:"mlApplicationName"`
		MlApplicationImplementationId   *string                                    `json:"mlApplicationImplementationId"`
		MlApplicationImplementationName *string                                    `json:"mlApplicationImplementationName"`
		IsEnabled                       *bool                                      `json:"isEnabled"`
		CompartmentId                   *string                                    `json:"compartmentId"`
		TimeCreated                     *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated                     *common.SDKTime                            `json:"timeUpdated"`
		LifecycleState                  MlApplicationInstanceLifecycleStateEnum    `json:"lifecycleState"`
		LifecycleSubstate               MlApplicationInstanceLifecycleSubstateEnum `json:"lifecycleSubstate"`
		LifecycleDetails                *string                                    `json:"lifecycleDetails"`
		FreeformTags                    map[string]string                          `json:"freeformTags"`
		DefinedTags                     map[string]map[string]interface{}          `json:"definedTags"`
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
	m.PredictionEndpointDetails = model.PredictionEndpointDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationName = model.MlApplicationName

	m.MlApplicationImplementationId = model.MlApplicationImplementationId

	m.MlApplicationImplementationName = model.MlApplicationImplementationName

	m.IsEnabled = model.IsEnabled

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleSubstate = model.LifecycleSubstate

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationInstanceLifecycleStateEnum Enum with underlying type: string
type MlApplicationInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationInstanceLifecycleStateEnum
const (
	MlApplicationInstanceLifecycleStateCreating       MlApplicationInstanceLifecycleStateEnum = "CREATING"
	MlApplicationInstanceLifecycleStateUpdating       MlApplicationInstanceLifecycleStateEnum = "UPDATING"
	MlApplicationInstanceLifecycleStateActive         MlApplicationInstanceLifecycleStateEnum = "ACTIVE"
	MlApplicationInstanceLifecycleStateInactive       MlApplicationInstanceLifecycleStateEnum = "INACTIVE"
	MlApplicationInstanceLifecycleStateDeleting       MlApplicationInstanceLifecycleStateEnum = "DELETING"
	MlApplicationInstanceLifecycleStateDeleted        MlApplicationInstanceLifecycleStateEnum = "DELETED"
	MlApplicationInstanceLifecycleStateNeedsAttention MlApplicationInstanceLifecycleStateEnum = "NEEDS_ATTENTION"
	MlApplicationInstanceLifecycleStateFailed         MlApplicationInstanceLifecycleStateEnum = "FAILED"
)

var mappingMlApplicationInstanceLifecycleStateEnum = map[string]MlApplicationInstanceLifecycleStateEnum{
	"CREATING":        MlApplicationInstanceLifecycleStateCreating,
	"UPDATING":        MlApplicationInstanceLifecycleStateUpdating,
	"ACTIVE":          MlApplicationInstanceLifecycleStateActive,
	"INACTIVE":        MlApplicationInstanceLifecycleStateInactive,
	"DELETING":        MlApplicationInstanceLifecycleStateDeleting,
	"DELETED":         MlApplicationInstanceLifecycleStateDeleted,
	"NEEDS_ATTENTION": MlApplicationInstanceLifecycleStateNeedsAttention,
	"FAILED":          MlApplicationInstanceLifecycleStateFailed,
}

var mappingMlApplicationInstanceLifecycleStateEnumLowerCase = map[string]MlApplicationInstanceLifecycleStateEnum{
	"creating":        MlApplicationInstanceLifecycleStateCreating,
	"updating":        MlApplicationInstanceLifecycleStateUpdating,
	"active":          MlApplicationInstanceLifecycleStateActive,
	"inactive":        MlApplicationInstanceLifecycleStateInactive,
	"deleting":        MlApplicationInstanceLifecycleStateDeleting,
	"deleted":         MlApplicationInstanceLifecycleStateDeleted,
	"needs_attention": MlApplicationInstanceLifecycleStateNeedsAttention,
	"failed":          MlApplicationInstanceLifecycleStateFailed,
}

// GetMlApplicationInstanceLifecycleStateEnumValues Enumerates the set of values for MlApplicationInstanceLifecycleStateEnum
func GetMlApplicationInstanceLifecycleStateEnumValues() []MlApplicationInstanceLifecycleStateEnum {
	values := make([]MlApplicationInstanceLifecycleStateEnum, 0)
	for _, v := range mappingMlApplicationInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for MlApplicationInstanceLifecycleStateEnum
func GetMlApplicationInstanceLifecycleStateEnumStringValues() []string {
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

// GetMappingMlApplicationInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationInstanceLifecycleStateEnum(val string) (MlApplicationInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MlApplicationInstanceLifecycleSubstateEnum Enum with underlying type: string
type MlApplicationInstanceLifecycleSubstateEnum string

// Set of constants representing the allowable values for MlApplicationInstanceLifecycleSubstateEnum
const (
	MlApplicationInstanceLifecycleSubstateCreating                    MlApplicationInstanceLifecycleSubstateEnum = "CREATING"
	MlApplicationInstanceLifecycleSubstateUpdating                    MlApplicationInstanceLifecycleSubstateEnum = "UPDATING"
	MlApplicationInstanceLifecycleSubstateUpgrading                   MlApplicationInstanceLifecycleSubstateEnum = "UPGRADING"
	MlApplicationInstanceLifecycleSubstateActive                      MlApplicationInstanceLifecycleSubstateEnum = "ACTIVE"
	MlApplicationInstanceLifecycleSubstateInactive                    MlApplicationInstanceLifecycleSubstateEnum = "INACTIVE"
	MlApplicationInstanceLifecycleSubstateDeleting                    MlApplicationInstanceLifecycleSubstateEnum = "DELETING"
	MlApplicationInstanceLifecycleSubstateDeleted                     MlApplicationInstanceLifecycleSubstateEnum = "DELETED"
	MlApplicationInstanceLifecycleSubstateNeedsAttention              MlApplicationInstanceLifecycleSubstateEnum = "NEEDS_ATTENTION"
	MlApplicationInstanceLifecycleSubstateFailed                      MlApplicationInstanceLifecycleSubstateEnum = "FAILED"
	MlApplicationInstanceLifecycleSubstateNonRecoverableProviderIssue MlApplicationInstanceLifecycleSubstateEnum = "NON_RECOVERABLE_PROVIDER_ISSUE"
	MlApplicationInstanceLifecycleSubstateRecoverableProviderIssue    MlApplicationInstanceLifecycleSubstateEnum = "RECOVERABLE_PROVIDER_ISSUE"
	MlApplicationInstanceLifecycleSubstateNonRecoverableServiceIssue  MlApplicationInstanceLifecycleSubstateEnum = "NON_RECOVERABLE_SERVICE_ISSUE"
	MlApplicationInstanceLifecycleSubstateRecoverableServiceIssue     MlApplicationInstanceLifecycleSubstateEnum = "RECOVERABLE_SERVICE_ISSUE"
)

var mappingMlApplicationInstanceLifecycleSubstateEnum = map[string]MlApplicationInstanceLifecycleSubstateEnum{
	"CREATING":                       MlApplicationInstanceLifecycleSubstateCreating,
	"UPDATING":                       MlApplicationInstanceLifecycleSubstateUpdating,
	"UPGRADING":                      MlApplicationInstanceLifecycleSubstateUpgrading,
	"ACTIVE":                         MlApplicationInstanceLifecycleSubstateActive,
	"INACTIVE":                       MlApplicationInstanceLifecycleSubstateInactive,
	"DELETING":                       MlApplicationInstanceLifecycleSubstateDeleting,
	"DELETED":                        MlApplicationInstanceLifecycleSubstateDeleted,
	"NEEDS_ATTENTION":                MlApplicationInstanceLifecycleSubstateNeedsAttention,
	"FAILED":                         MlApplicationInstanceLifecycleSubstateFailed,
	"NON_RECOVERABLE_PROVIDER_ISSUE": MlApplicationInstanceLifecycleSubstateNonRecoverableProviderIssue,
	"RECOVERABLE_PROVIDER_ISSUE":     MlApplicationInstanceLifecycleSubstateRecoverableProviderIssue,
	"NON_RECOVERABLE_SERVICE_ISSUE":  MlApplicationInstanceLifecycleSubstateNonRecoverableServiceIssue,
	"RECOVERABLE_SERVICE_ISSUE":      MlApplicationInstanceLifecycleSubstateRecoverableServiceIssue,
}

var mappingMlApplicationInstanceLifecycleSubstateEnumLowerCase = map[string]MlApplicationInstanceLifecycleSubstateEnum{
	"creating":                       MlApplicationInstanceLifecycleSubstateCreating,
	"updating":                       MlApplicationInstanceLifecycleSubstateUpdating,
	"upgrading":                      MlApplicationInstanceLifecycleSubstateUpgrading,
	"active":                         MlApplicationInstanceLifecycleSubstateActive,
	"inactive":                       MlApplicationInstanceLifecycleSubstateInactive,
	"deleting":                       MlApplicationInstanceLifecycleSubstateDeleting,
	"deleted":                        MlApplicationInstanceLifecycleSubstateDeleted,
	"needs_attention":                MlApplicationInstanceLifecycleSubstateNeedsAttention,
	"failed":                         MlApplicationInstanceLifecycleSubstateFailed,
	"non_recoverable_provider_issue": MlApplicationInstanceLifecycleSubstateNonRecoverableProviderIssue,
	"recoverable_provider_issue":     MlApplicationInstanceLifecycleSubstateRecoverableProviderIssue,
	"non_recoverable_service_issue":  MlApplicationInstanceLifecycleSubstateNonRecoverableServiceIssue,
	"recoverable_service_issue":      MlApplicationInstanceLifecycleSubstateRecoverableServiceIssue,
}

// GetMlApplicationInstanceLifecycleSubstateEnumValues Enumerates the set of values for MlApplicationInstanceLifecycleSubstateEnum
func GetMlApplicationInstanceLifecycleSubstateEnumValues() []MlApplicationInstanceLifecycleSubstateEnum {
	values := make([]MlApplicationInstanceLifecycleSubstateEnum, 0)
	for _, v := range mappingMlApplicationInstanceLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationInstanceLifecycleSubstateEnumStringValues Enumerates the set of values in String for MlApplicationInstanceLifecycleSubstateEnum
func GetMlApplicationInstanceLifecycleSubstateEnumStringValues() []string {
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

// GetMappingMlApplicationInstanceLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationInstanceLifecycleSubstateEnum(val string) (MlApplicationInstanceLifecycleSubstateEnum, bool) {
	enum, ok := mappingMlApplicationInstanceLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
