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

// MlApplicationImplementation Resource representing solution for AI/ML use-case defined by ML Application
type MlApplicationImplementation struct {

	// The OCID of the MlApplicationImplementation. Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// ML Application Implementation name which is unique for given ML Application.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the ML Application implemented by this ML Application Implementation.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId)
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// The OCID of the compartment where ML Application Implementation is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Creation time of MlApplicationImplementation creation in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time of last MlApplicationImplementation update in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the MlApplicationImplementation.
	LifecycleState MlApplicationImplementationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Description of ML Application Implementation defined in ML Application package descriptor
	Description *string `mandatory:"false" json:"description"`

	// The version of ML Application Package (e.g. "1.2" or "2.0.4") defined in ML Application package descriptor. Value is not mandatory only for CREATING state otherwise it must be always presented.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// List of application components (OCI resources shared for all MlApplicationInstances). These have been created automatically based on their definitions in the ML Application package.
	ApplicationComponents []ApplicationComponent `mandatory:"false" json:"applicationComponents"`

	// Schema of configuration which needs to be provided for each ML Application Instance. It is defined in the ML Application package descriptor.
	ConfigurationSchema []ConfigurationPropertySchema `mandatory:"false" json:"configurationSchema"`

	MlApplicationPackageArguments *MlApplicationPackageArguments `mandatory:"false" json:"mlApplicationPackageArguments"`

	// List of ML Application Implementation OCIDs for which migration from this implementation is allowed. Migration means that if consumers change implementation for their instances to implementation with OCID from this list, instance components will be updated in place otherwise new instance components are created based on the new implementation and old instance components are removed.
	AllowedMigrationDestinations []string `mandatory:"false" json:"allowedMigrationDestinations"`

	Logging *ImplementationLogging `mandatory:"false" json:"logging"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationImplementation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationImplementation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationImplementationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationImplementationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplicationImplementation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                   *string                                       `json:"description"`
		PackageVersion                *string                                       `json:"packageVersion"`
		ApplicationComponents         []applicationcomponent                        `json:"applicationComponents"`
		ConfigurationSchema           []ConfigurationPropertySchema                 `json:"configurationSchema"`
		MlApplicationPackageArguments *MlApplicationPackageArguments                `json:"mlApplicationPackageArguments"`
		AllowedMigrationDestinations  []string                                      `json:"allowedMigrationDestinations"`
		Logging                       *ImplementationLogging                        `json:"logging"`
		SystemTags                    map[string]map[string]interface{}             `json:"systemTags"`
		Id                            *string                                       `json:"id"`
		Name                          *string                                       `json:"name"`
		MlApplicationId               *string                                       `json:"mlApplicationId"`
		MlApplicationName             *string                                       `json:"mlApplicationName"`
		CompartmentId                 *string                                       `json:"compartmentId"`
		TimeCreated                   *common.SDKTime                               `json:"timeCreated"`
		TimeUpdated                   *common.SDKTime                               `json:"timeUpdated"`
		LifecycleState                MlApplicationImplementationLifecycleStateEnum `json:"lifecycleState"`
		LifecycleDetails              *string                                       `json:"lifecycleDetails"`
		FreeformTags                  map[string]string                             `json:"freeformTags"`
		DefinedTags                   map[string]map[string]interface{}             `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.PackageVersion = model.PackageVersion

	m.ApplicationComponents = make([]ApplicationComponent, len(model.ApplicationComponents))
	for i, n := range model.ApplicationComponents {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ApplicationComponents[i] = nn.(ApplicationComponent)
		} else {
			m.ApplicationComponents[i] = nil
		}
	}
	m.ConfigurationSchema = make([]ConfigurationPropertySchema, len(model.ConfigurationSchema))
	copy(m.ConfigurationSchema, model.ConfigurationSchema)
	m.MlApplicationPackageArguments = model.MlApplicationPackageArguments

	m.AllowedMigrationDestinations = make([]string, len(model.AllowedMigrationDestinations))
	copy(m.AllowedMigrationDestinations, model.AllowedMigrationDestinations)
	m.Logging = model.Logging

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationName = model.MlApplicationName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationImplementationLifecycleStateEnum Enum with underlying type: string
type MlApplicationImplementationLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationImplementationLifecycleStateEnum
const (
	MlApplicationImplementationLifecycleStateCreating       MlApplicationImplementationLifecycleStateEnum = "CREATING"
	MlApplicationImplementationLifecycleStateActive         MlApplicationImplementationLifecycleStateEnum = "ACTIVE"
	MlApplicationImplementationLifecycleStateNeedsAttention MlApplicationImplementationLifecycleStateEnum = "NEEDS_ATTENTION"
	MlApplicationImplementationLifecycleStateFailed         MlApplicationImplementationLifecycleStateEnum = "FAILED"
	MlApplicationImplementationLifecycleStateDeleting       MlApplicationImplementationLifecycleStateEnum = "DELETING"
	MlApplicationImplementationLifecycleStateUpdating       MlApplicationImplementationLifecycleStateEnum = "UPDATING"
)

var mappingMlApplicationImplementationLifecycleStateEnum = map[string]MlApplicationImplementationLifecycleStateEnum{
	"CREATING":        MlApplicationImplementationLifecycleStateCreating,
	"ACTIVE":          MlApplicationImplementationLifecycleStateActive,
	"NEEDS_ATTENTION": MlApplicationImplementationLifecycleStateNeedsAttention,
	"FAILED":          MlApplicationImplementationLifecycleStateFailed,
	"DELETING":        MlApplicationImplementationLifecycleStateDeleting,
	"UPDATING":        MlApplicationImplementationLifecycleStateUpdating,
}

var mappingMlApplicationImplementationLifecycleStateEnumLowerCase = map[string]MlApplicationImplementationLifecycleStateEnum{
	"creating":        MlApplicationImplementationLifecycleStateCreating,
	"active":          MlApplicationImplementationLifecycleStateActive,
	"needs_attention": MlApplicationImplementationLifecycleStateNeedsAttention,
	"failed":          MlApplicationImplementationLifecycleStateFailed,
	"deleting":        MlApplicationImplementationLifecycleStateDeleting,
	"updating":        MlApplicationImplementationLifecycleStateUpdating,
}

// GetMlApplicationImplementationLifecycleStateEnumValues Enumerates the set of values for MlApplicationImplementationLifecycleStateEnum
func GetMlApplicationImplementationLifecycleStateEnumValues() []MlApplicationImplementationLifecycleStateEnum {
	values := make([]MlApplicationImplementationLifecycleStateEnum, 0)
	for _, v := range mappingMlApplicationImplementationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationImplementationLifecycleStateEnumStringValues Enumerates the set of values in String for MlApplicationImplementationLifecycleStateEnum
func GetMlApplicationImplementationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"UPDATING",
	}
}

// GetMappingMlApplicationImplementationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationImplementationLifecycleStateEnum(val string) (MlApplicationImplementationLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationImplementationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
