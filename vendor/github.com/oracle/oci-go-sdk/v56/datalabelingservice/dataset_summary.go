// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatasetSummary A dataset summary is the representation returned in list views.  It is usually a subset of the full dataset entity and should not contain any potentially sensitive information.
type DatasetSummary struct {

	// The OCID of the Dataset
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the Dataset was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The state of a Dataset.
	LifecycleState DatasetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The annotation format name required for labeling records.
	AnnotationFormat *string `mandatory:"true" json:"annotationFormat"`

	DatasetFormatDetails DatasetFormatDetails `mandatory:"true" json:"datasetFormatDetails"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The usage of system tag keys. These predefined keys are scoped to namespaces.
	// For example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DatasetSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DatasetSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		LifecycleDetails     *string                           `json:"lifecycleDetails"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		SystemTags           map[string]map[string]interface{} `json:"systemTags"`
		Id                   *string                           `json:"id"`
		CompartmentId        *string                           `json:"compartmentId"`
		TimeCreated          *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState       DatasetLifecycleStateEnum         `json:"lifecycleState"`
		AnnotationFormat     *string                           `json:"annotationFormat"`
		DatasetFormatDetails datasetformatdetails              `json:"datasetFormatDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.AnnotationFormat = model.AnnotationFormat

	nn, e = model.DatasetFormatDetails.UnmarshalPolymorphicJSON(model.DatasetFormatDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetFormatDetails = nn.(DatasetFormatDetails)
	} else {
		m.DatasetFormatDetails = nil
	}

	return
}
