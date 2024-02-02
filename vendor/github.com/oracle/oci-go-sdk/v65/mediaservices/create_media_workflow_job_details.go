// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMediaWorkflowJobDetails Information to run the MediaWorkflow.
type CreateMediaWorkflowJobDetails interface {

	// ID of the compartment in which the job should be created.
	GetCompartmentId() *string

	// Configurations to be applied to this run of the workflow.
	GetMediaWorkflowConfigurationIds() []string

	// Name of the Media Workflow Job. Does not have to be unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Parameters that override parameters specified in MediaWorkflowTaskDeclarations, the MediaWorkflow,
	// the MediaWorkflow's MediaWorkflowConfigurations and the MediaWorkflowConfigurations of this
	// MediaWorkflowJob. The parameters are given as JSON. The top level and 2nd level elements must be
	// JSON objects (vs arrays, scalars, etc). The top level keys refer to a task's key and the 2nd level
	// keys refer to a parameter's name.
	GetParameters() map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createmediaworkflowjobdetails struct {
	JsonData                      []byte
	MediaWorkflowConfigurationIds []string                          `mandatory:"false" json:"mediaWorkflowConfigurationIds"`
	DisplayName                   *string                           `mandatory:"false" json:"displayName"`
	Parameters                    map[string]interface{}            `mandatory:"false" json:"parameters"`
	FreeformTags                  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Locks                         []ResourceLock                    `mandatory:"false" json:"locks"`
	CompartmentId                 *string                           `mandatory:"true" json:"compartmentId"`
	WorkflowIdentifierType        string                            `json:"workflowIdentifierType"`
}

// UnmarshalJSON unmarshals json
func (m *createmediaworkflowjobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatemediaworkflowjobdetails createmediaworkflowjobdetails
	s := struct {
		Model Unmarshalercreatemediaworkflowjobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.MediaWorkflowConfigurationIds = s.Model.MediaWorkflowConfigurationIds
	m.DisplayName = s.Model.DisplayName
	m.Parameters = s.Model.Parameters
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Locks = s.Model.Locks
	m.WorkflowIdentifierType = s.Model.WorkflowIdentifierType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createmediaworkflowjobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.WorkflowIdentifierType {
	case "NAME":
		mm := CreateMediaWorkflowJobByNameDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ID":
		mm := CreateMediaWorkflowJobByIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateMediaWorkflowJobDetails: %s.", m.WorkflowIdentifierType)
		return *m, nil
	}
}

// GetMediaWorkflowConfigurationIds returns MediaWorkflowConfigurationIds
func (m createmediaworkflowjobdetails) GetMediaWorkflowConfigurationIds() []string {
	return m.MediaWorkflowConfigurationIds
}

// GetDisplayName returns DisplayName
func (m createmediaworkflowjobdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetParameters returns Parameters
func (m createmediaworkflowjobdetails) GetParameters() map[string]interface{} {
	return m.Parameters
}

// GetFreeformTags returns FreeformTags
func (m createmediaworkflowjobdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createmediaworkflowjobdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m createmediaworkflowjobdetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetCompartmentId returns CompartmentId
func (m createmediaworkflowjobdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createmediaworkflowjobdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createmediaworkflowjobdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum Enum with underlying type: string
type CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum string

// Set of constants representing the allowable values for CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum
const (
	CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeId   CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum = "ID"
	CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeName CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum = "NAME"
)

var mappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum = map[string]CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum{
	"ID":   CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeId,
	"NAME": CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeName,
}

var mappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumLowerCase = map[string]CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum{
	"id":   CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeId,
	"name": CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeName,
}

// GetCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumValues Enumerates the set of values for CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum
func GetCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumValues() []CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum {
	values := make([]CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum, 0)
	for _, v := range mappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumStringValues Enumerates the set of values in String for CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum
func GetCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumStringValues() []string {
	return []string{
		"ID",
		"NAME",
	}
}

// GetMappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum(val string) (CreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnum, bool) {
	enum, ok := mappingCreateMediaWorkflowJobDetailsWorkflowIdentifierTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
