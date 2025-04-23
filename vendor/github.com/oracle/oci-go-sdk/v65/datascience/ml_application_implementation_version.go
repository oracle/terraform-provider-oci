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

// MlApplicationImplementationVersion Read-only fully managed snapshot of MlApplicationImplementation taken when MlApplicationImplementation was updated with new ML Application package.
type MlApplicationImplementationVersion struct {

	// The OCID of the MlApplicationImplementationVersion. Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the MlApplicationImplementation for which this resource keeps the historical state.
	MlApplicationImplementationId *string `mandatory:"true" json:"mlApplicationImplementationId"`

	// ML Application Implementation name which is unique for given ML Application.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the ML Application implemented by this ML Application Implementation.
	MlApplicationId *string `mandatory:"true" json:"mlApplicationId"`

	// The name of ML Application (based on mlApplicationId)
	MlApplicationName *string `mandatory:"true" json:"mlApplicationName"`

	// Creation time of MlApplicationImplementationVersion in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MlApplicationImplementationVersion.
	LifecycleState MlApplicationImplementationVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplicationImplementationVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplicationImplementationVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationImplementationVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationImplementationVersionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplicationImplementationVersion) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                   *string                                              `json:"description"`
		PackageVersion                *string                                              `json:"packageVersion"`
		ApplicationComponents         []applicationcomponent                               `json:"applicationComponents"`
		ConfigurationSchema           []ConfigurationPropertySchema                        `json:"configurationSchema"`
		MlApplicationPackageArguments *MlApplicationPackageArguments                       `json:"mlApplicationPackageArguments"`
		AllowedMigrationDestinations  []string                                             `json:"allowedMigrationDestinations"`
		SystemTags                    map[string]map[string]interface{}                    `json:"systemTags"`
		Id                            *string                                              `json:"id"`
		MlApplicationImplementationId *string                                              `json:"mlApplicationImplementationId"`
		Name                          *string                                              `json:"name"`
		MlApplicationId               *string                                              `json:"mlApplicationId"`
		MlApplicationName             *string                                              `json:"mlApplicationName"`
		TimeCreated                   *common.SDKTime                                      `json:"timeCreated"`
		LifecycleState                MlApplicationImplementationVersionLifecycleStateEnum `json:"lifecycleState"`
		LifecycleDetails              *string                                              `json:"lifecycleDetails"`
		FreeformTags                  map[string]string                                    `json:"freeformTags"`
		DefinedTags                   map[string]map[string]interface{}                    `json:"definedTags"`
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
	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.MlApplicationImplementationId = model.MlApplicationImplementationId

	m.Name = model.Name

	m.MlApplicationId = model.MlApplicationId

	m.MlApplicationName = model.MlApplicationName

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationImplementationVersionLifecycleStateEnum Enum with underlying type: string
type MlApplicationImplementationVersionLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationImplementationVersionLifecycleStateEnum
const (
	MlApplicationImplementationVersionLifecycleStateCreating MlApplicationImplementationVersionLifecycleStateEnum = "CREATING"
	MlApplicationImplementationVersionLifecycleStateActive   MlApplicationImplementationVersionLifecycleStateEnum = "ACTIVE"
	MlApplicationImplementationVersionLifecycleStateFailed   MlApplicationImplementationVersionLifecycleStateEnum = "FAILED"
	MlApplicationImplementationVersionLifecycleStateDeleting MlApplicationImplementationVersionLifecycleStateEnum = "DELETING"
)

var mappingMlApplicationImplementationVersionLifecycleStateEnum = map[string]MlApplicationImplementationVersionLifecycleStateEnum{
	"CREATING": MlApplicationImplementationVersionLifecycleStateCreating,
	"ACTIVE":   MlApplicationImplementationVersionLifecycleStateActive,
	"FAILED":   MlApplicationImplementationVersionLifecycleStateFailed,
	"DELETING": MlApplicationImplementationVersionLifecycleStateDeleting,
}

var mappingMlApplicationImplementationVersionLifecycleStateEnumLowerCase = map[string]MlApplicationImplementationVersionLifecycleStateEnum{
	"creating": MlApplicationImplementationVersionLifecycleStateCreating,
	"active":   MlApplicationImplementationVersionLifecycleStateActive,
	"failed":   MlApplicationImplementationVersionLifecycleStateFailed,
	"deleting": MlApplicationImplementationVersionLifecycleStateDeleting,
}

// GetMlApplicationImplementationVersionLifecycleStateEnumValues Enumerates the set of values for MlApplicationImplementationVersionLifecycleStateEnum
func GetMlApplicationImplementationVersionLifecycleStateEnumValues() []MlApplicationImplementationVersionLifecycleStateEnum {
	values := make([]MlApplicationImplementationVersionLifecycleStateEnum, 0)
	for _, v := range mappingMlApplicationImplementationVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationImplementationVersionLifecycleStateEnumStringValues Enumerates the set of values in String for MlApplicationImplementationVersionLifecycleStateEnum
func GetMlApplicationImplementationVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"DELETING",
	}
}

// GetMappingMlApplicationImplementationVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationImplementationVersionLifecycleStateEnum(val string) (MlApplicationImplementationVersionLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationImplementationVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
