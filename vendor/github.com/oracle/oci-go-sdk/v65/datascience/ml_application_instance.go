// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

	// Resource name which must be unique in the context of given ML Application and tenancy.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of ML Application. This resource is an instance of ML Application referenced by this OCID.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId).
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// The OCID of the compartment where the MlApplicationInstance is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// States whether the MlApplicationInstance is supposed to be in ACTIVE lifecycle state.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The time the the MlApplication was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MlApplicationInstance.
	LifecycleState MlApplicationInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	AuthConfiguration AuthConfiguration `mandatory:"false" json:"authConfiguration"`

	// Data that are used for provisioning of the given MlApplicationInstance. These are validated against configurationSchema defined in referenced MlApplication.
	Configuration []ConfigurationProperty `mandatory:"false" json:"configuration"`

	// Array of all prediction URIs per use-case.
	PredictionUris []PredictionUri `mandatory:"false" json:"predictionUris"`

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

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplicationInstance) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AuthConfiguration authconfiguration                       `json:"authConfiguration"`
		Configuration     []ConfigurationProperty                 `json:"configuration"`
		PredictionUris    []PredictionUri                         `json:"predictionUris"`
		SystemTags        map[string]map[string]interface{}       `json:"systemTags"`
		Id                *string                                 `json:"id"`
		Name              *string                                 `json:"name"`
		MlApplicationId   *string                                 `json:"mlApplicationId"`
		MlApplicationName *string                                 `json:"mlApplicationName"`
		CompartmentId     *string                                 `json:"compartmentId"`
		IsEnabled         *bool                                   `json:"isEnabled"`
		TimeCreated       *common.SDKTime                         `json:"timeCreated"`
		LifecycleState    MlApplicationInstanceLifecycleStateEnum `json:"lifecycleState"`
		LifecycleDetails  *string                                 `json:"lifecycleDetails"`
		FreeformTags      map[string]string                       `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{}       `json:"definedTags"`
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
	for i, n := range model.Configuration {
		m.Configuration[i] = n
	}

	m.PredictionUris = make([]PredictionUri, len(model.PredictionUris))
	for i, n := range model.PredictionUris {
		m.PredictionUris[i] = n
	}

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationName = model.MlApplicationName

	m.CompartmentId = model.CompartmentId

	m.IsEnabled = model.IsEnabled

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationInstanceLifecycleStateEnum Enum with underlying type: string
type MlApplicationInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationInstanceLifecycleStateEnum
const (
	MlApplicationInstanceLifecycleStateCreating MlApplicationInstanceLifecycleStateEnum = "CREATING"
	MlApplicationInstanceLifecycleStateUpdating MlApplicationInstanceLifecycleStateEnum = "UPDATING"
	MlApplicationInstanceLifecycleStateActive   MlApplicationInstanceLifecycleStateEnum = "ACTIVE"
	MlApplicationInstanceLifecycleStateInactive MlApplicationInstanceLifecycleStateEnum = "INACTIVE"
	MlApplicationInstanceLifecycleStateDeleting MlApplicationInstanceLifecycleStateEnum = "DELETING"
	MlApplicationInstanceLifecycleStateDeleted  MlApplicationInstanceLifecycleStateEnum = "DELETED"
	MlApplicationInstanceLifecycleStateFailed   MlApplicationInstanceLifecycleStateEnum = "FAILED"
)

var mappingMlApplicationInstanceLifecycleStateEnum = map[string]MlApplicationInstanceLifecycleStateEnum{
	"CREATING": MlApplicationInstanceLifecycleStateCreating,
	"UPDATING": MlApplicationInstanceLifecycleStateUpdating,
	"ACTIVE":   MlApplicationInstanceLifecycleStateActive,
	"INACTIVE": MlApplicationInstanceLifecycleStateInactive,
	"DELETING": MlApplicationInstanceLifecycleStateDeleting,
	"DELETED":  MlApplicationInstanceLifecycleStateDeleted,
	"FAILED":   MlApplicationInstanceLifecycleStateFailed,
}

var mappingMlApplicationInstanceLifecycleStateEnumLowerCase = map[string]MlApplicationInstanceLifecycleStateEnum{
	"creating": MlApplicationInstanceLifecycleStateCreating,
	"updating": MlApplicationInstanceLifecycleStateUpdating,
	"active":   MlApplicationInstanceLifecycleStateActive,
	"inactive": MlApplicationInstanceLifecycleStateInactive,
	"deleting": MlApplicationInstanceLifecycleStateDeleting,
	"deleted":  MlApplicationInstanceLifecycleStateDeleted,
	"failed":   MlApplicationInstanceLifecycleStateFailed,
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
		"FAILED",
	}
}

// GetMappingMlApplicationInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationInstanceLifecycleStateEnum(val string) (MlApplicationInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
