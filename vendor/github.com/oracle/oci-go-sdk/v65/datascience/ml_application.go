// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MlApplication Resource representing a definition of an AI/ML use-case
type MlApplication struct {

	// The OCID of the MlApplication. Unique identifier that is immutable after creation.
	Id *string `mandatory:"true" json:"id"`

	// The name of MlApplication. It is unique in a given tenancy.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment where the MlApplication is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Creation time of MlApplication in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time of last MlApplication update in the format defined by RFC 3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the MlApplication.
	LifecycleState MlApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Optional description of the ML Application
	Description *string `mandatory:"false" json:"description"`

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

// MlApplicationLifecycleStateEnum Enum with underlying type: string
type MlApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for MlApplicationLifecycleStateEnum
const (
	MlApplicationLifecycleStateActive MlApplicationLifecycleStateEnum = "ACTIVE"
	MlApplicationLifecycleStateFailed MlApplicationLifecycleStateEnum = "FAILED"
)

var mappingMlApplicationLifecycleStateEnum = map[string]MlApplicationLifecycleStateEnum{
	"ACTIVE": MlApplicationLifecycleStateActive,
	"FAILED": MlApplicationLifecycleStateFailed,
}

var mappingMlApplicationLifecycleStateEnumLowerCase = map[string]MlApplicationLifecycleStateEnum{
	"active": MlApplicationLifecycleStateActive,
	"failed": MlApplicationLifecycleStateFailed,
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
		"ACTIVE",
		"FAILED",
	}
}

// GetMappingMlApplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlApplicationLifecycleStateEnum(val string) (MlApplicationLifecycleStateEnum, bool) {
	enum, ok := mappingMlApplicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
