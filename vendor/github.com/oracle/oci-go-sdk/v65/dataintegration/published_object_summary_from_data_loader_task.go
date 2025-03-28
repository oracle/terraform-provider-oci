// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublishedObjectSummaryFromDataLoaderTask The data loader task published object summary.
type PublishedObjectSummaryFromDataLoaderTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	DataFlow *DataFlow `mandatory:"false" json:"dataFlow"`
}

// GetKey returns Key
func (m PublishedObjectSummaryFromDataLoaderTask) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m PublishedObjectSummaryFromDataLoaderTask) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m PublishedObjectSummaryFromDataLoaderTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m PublishedObjectSummaryFromDataLoaderTask) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m PublishedObjectSummaryFromDataLoaderTask) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m PublishedObjectSummaryFromDataLoaderTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m PublishedObjectSummaryFromDataLoaderTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m PublishedObjectSummaryFromDataLoaderTask) GetIdentifier() *string {
	return m.Identifier
}

// GetMetadata returns Metadata
func (m PublishedObjectSummaryFromDataLoaderTask) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m PublishedObjectSummaryFromDataLoaderTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublishedObjectSummaryFromDataLoaderTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PublishedObjectSummaryFromDataLoaderTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePublishedObjectSummaryFromDataLoaderTask PublishedObjectSummaryFromDataLoaderTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypePublishedObjectSummaryFromDataLoaderTask
	}{
		"DATA_LOADER_TASK",
		(MarshalTypePublishedObjectSummaryFromDataLoaderTask)(m),
	}

	return json.Marshal(&s)
}
