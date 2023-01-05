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

// MlApplication Description of MlApplication.
type MlApplication struct {

	// The OCID of the MlApplication. Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of MlApplication. It is unique in a given tenancy.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment where the MlApplication is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the MlApplication was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MlApplication.
	LifecycleState MlApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Optional description of the ML Application
	Description *string `mandatory:"false" json:"description"`

	PredictionContract *PredictionContract `mandatory:"false" json:"predictionContract"`

	// List of application components (OCI resources shared for all MlApplicationInstances).
	ApplicationComponents []ApplicationComponent `mandatory:"false" json:"applicationComponents"`

	// Schema of configuration which needs to be provided for each ML Application Instance
	ConfigurationSchema []ConfigurationPropertySchema `mandatory:"false" json:"configurationSchema"`

	// Vault ID used for secure persisting possible secrets from configuration
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MlApplication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MlApplication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMlApplicationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MlApplication) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description           *string                           `json:"description"`
		PredictionContract    *PredictionContract               `json:"predictionContract"`
		ApplicationComponents []applicationcomponent            `json:"applicationComponents"`
		ConfigurationSchema   []ConfigurationPropertySchema     `json:"configurationSchema"`
		VaultId               *string                           `json:"vaultId"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		Id                    *string                           `json:"id"`
		Name                  *string                           `json:"name"`
		CompartmentId         *string                           `json:"compartmentId"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		LifecycleState        MlApplicationLifecycleStateEnum   `json:"lifecycleState"`
		LifecycleDetails      *string                           `json:"lifecycleDetails"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.PredictionContract = model.PredictionContract

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
	for i, n := range model.ConfigurationSchema {
		m.ConfigurationSchema[i] = n
	}

	m.VaultId = model.VaultId

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// MlApplicationLifecycleStateEnum Enum with underlying type: string
type MlApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationLifecycleStateEnum
const (
	MlApplicationLifecycleStateCreating MlApplicationLifecycleStateEnum = "CREATING"
	MlApplicationLifecycleStateActive   MlApplicationLifecycleStateEnum = "ACTIVE"
	MlApplicationLifecycleStateUpdating MlApplicationLifecycleStateEnum = "UPDATING"
	MlApplicationLifecycleStateDeleting MlApplicationLifecycleStateEnum = "DELETING"
	MlApplicationLifecycleStateDeleted  MlApplicationLifecycleStateEnum = "DELETED"
	MlApplicationLifecycleStateFailed   MlApplicationLifecycleStateEnum = "FAILED"
)

var mappingMlApplicationLifecycleStateEnum = map[string]MlApplicationLifecycleStateEnum{
	"CREATING": MlApplicationLifecycleStateCreating,
	"ACTIVE":   MlApplicationLifecycleStateActive,
	"UPDATING": MlApplicationLifecycleStateUpdating,
	"DELETING": MlApplicationLifecycleStateDeleting,
	"DELETED":  MlApplicationLifecycleStateDeleted,
	"FAILED":   MlApplicationLifecycleStateFailed,
}

var mappingMlApplicationLifecycleStateEnumLowerCase = map[string]MlApplicationLifecycleStateEnum{
	"creating": MlApplicationLifecycleStateCreating,
	"active":   MlApplicationLifecycleStateActive,
	"updating": MlApplicationLifecycleStateUpdating,
	"deleting": MlApplicationLifecycleStateDeleting,
	"deleted":  MlApplicationLifecycleStateDeleted,
	"failed":   MlApplicationLifecycleStateFailed,
}

// GetMlApplicationLifecycleStateEnumValues Enumerates the set of values for MlApplicationLifecycleStateEnum
func GetMlApplicationLifecycleStateEnumValues() []MlApplicationLifecycleStateEnum {
	values := make([]MlApplicationLifecycleStateEnum, 0)
	for _, v := range mappingMlApplicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMlApplicationLifecycleStateEnumStringValues Enumerates the set of values in String for MlApplicationLifecycleStateEnum
func GetMlApplicationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMlApplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationLifecycleStateEnum(val string) (MlApplicationLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
