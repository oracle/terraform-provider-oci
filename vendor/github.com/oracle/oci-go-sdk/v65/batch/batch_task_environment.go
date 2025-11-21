// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTaskEnvironment A batch task environment resource defines the software that is used to execute tasks.
// It mainly contains a URL to a container image along with other configurations that are needed by the image.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type BatchTaskEnvironment struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task environment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// If not specified or provided as null or empty string, it be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchtaskenvironment20250914115623.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The URL of the ocir image.
	ImageUrl *string `mandatory:"true" json:"imageUrl"`

	// List of volumes attached to the image.
	// The use cases of the volumes are but not limited to: read the input of the task and write the output.
	Volumes []BatchTaskEnvironmentVolume `mandatory:"true" json:"volumes"`

	// The current state of the batch task environment.
	LifecycleState BatchTaskEnvironmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the batch task environment was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The batch task environment description.
	Description *string `mandatory:"false" json:"description"`

	SecurityContext *SecurityContext `mandatory:"false" json:"securityContext"`

	// Container's working directory.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The date and time the batch task environment was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BatchTaskEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchTaskEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskEnvironmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskEnvironmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BatchTaskEnvironment) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                `json:"description"`
		SecurityContext  *SecurityContext                       `json:"securityContext"`
		WorkingDirectory *string                                `json:"workingDirectory"`
		TimeUpdated      *common.SDKTime                        `json:"timeUpdated"`
		Id               *string                                `json:"id"`
		CompartmentId    *string                                `json:"compartmentId"`
		DisplayName      *string                                `json:"displayName"`
		ImageUrl         *string                                `json:"imageUrl"`
		Volumes          []batchtaskenvironmentvolume           `json:"volumes"`
		LifecycleState   BatchTaskEnvironmentLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                        `json:"timeCreated"`
		DefinedTags      map[string]map[string]interface{}      `json:"definedTags"`
		FreeformTags     map[string]string                      `json:"freeformTags"`
		SystemTags       map[string]map[string]interface{}      `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.SecurityContext = model.SecurityContext

	m.WorkingDirectory = model.WorkingDirectory

	m.TimeUpdated = model.TimeUpdated

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.ImageUrl = model.ImageUrl

	m.Volumes = make([]BatchTaskEnvironmentVolume, len(model.Volumes))
	for i, n := range model.Volumes {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Volumes[i] = nn.(BatchTaskEnvironmentVolume)
		} else {
			m.Volumes[i] = nil
		}
	}
	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	return
}

// BatchTaskEnvironmentLifecycleStateEnum Enum with underlying type: string
type BatchTaskEnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for BatchTaskEnvironmentLifecycleStateEnum
const (
	BatchTaskEnvironmentLifecycleStateActive  BatchTaskEnvironmentLifecycleStateEnum = "ACTIVE"
	BatchTaskEnvironmentLifecycleStateDeleted BatchTaskEnvironmentLifecycleStateEnum = "DELETED"
)

var mappingBatchTaskEnvironmentLifecycleStateEnum = map[string]BatchTaskEnvironmentLifecycleStateEnum{
	"ACTIVE":  BatchTaskEnvironmentLifecycleStateActive,
	"DELETED": BatchTaskEnvironmentLifecycleStateDeleted,
}

var mappingBatchTaskEnvironmentLifecycleStateEnumLowerCase = map[string]BatchTaskEnvironmentLifecycleStateEnum{
	"active":  BatchTaskEnvironmentLifecycleStateActive,
	"deleted": BatchTaskEnvironmentLifecycleStateDeleted,
}

// GetBatchTaskEnvironmentLifecycleStateEnumValues Enumerates the set of values for BatchTaskEnvironmentLifecycleStateEnum
func GetBatchTaskEnvironmentLifecycleStateEnumValues() []BatchTaskEnvironmentLifecycleStateEnum {
	values := make([]BatchTaskEnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingBatchTaskEnvironmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskEnvironmentLifecycleStateEnumStringValues Enumerates the set of values in String for BatchTaskEnvironmentLifecycleStateEnum
func GetBatchTaskEnvironmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingBatchTaskEnvironmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskEnvironmentLifecycleStateEnum(val string) (BatchTaskEnvironmentLifecycleStateEnum, bool) {
	enum, ok := mappingBatchTaskEnvironmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
