// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatasetDetails Parameters needed to create a new Dataset. A Dataset allows a user to describe the data source that provides the Records and how Annotations should be applied to the Records.
type CreateDatasetDetails struct {

	// The OCID of the compartment of the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The annotation format name required for labeling records.
	AnnotationFormat *string `mandatory:"true" json:"annotationFormat"`

	DatasetSourceDetails DatasetSourceDetails `mandatory:"true" json:"datasetSourceDetails"`

	DatasetFormatDetails DatasetFormatDetails `mandatory:"true" json:"datasetFormatDetails"`

	LabelSet *LabelSet `mandatory:"true" json:"labelSet"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user provided description of the dataset
	Description *string `mandatory:"false" json:"description"`

	InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `mandatory:"false" json:"initialRecordGenerationConfiguration"`

	InitialImportDatasetConfiguration *InitialImportDatasetConfiguration `mandatory:"false" json:"initialImportDatasetConfiguration"`

	// The labeling instructions for human labelers in rich text format
	LabelingInstructions *string `mandatory:"false" json:"labelingInstructions"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDatasetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatasetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatasetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                          *string                               `json:"displayName"`
		Description                          *string                               `json:"description"`
		InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `json:"initialRecordGenerationConfiguration"`
		InitialImportDatasetConfiguration    *InitialImportDatasetConfiguration    `json:"initialImportDatasetConfiguration"`
		LabelingInstructions                 *string                               `json:"labelingInstructions"`
		FreeformTags                         map[string]string                     `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{}     `json:"definedTags"`
		CompartmentId                        *string                               `json:"compartmentId"`
		AnnotationFormat                     *string                               `json:"annotationFormat"`
		DatasetSourceDetails                 datasetsourcedetails                  `json:"datasetSourceDetails"`
		DatasetFormatDetails                 datasetformatdetails                  `json:"datasetFormatDetails"`
		LabelSet                             *LabelSet                             `json:"labelSet"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.InitialRecordGenerationConfiguration = model.InitialRecordGenerationConfiguration

	m.InitialImportDatasetConfiguration = model.InitialImportDatasetConfiguration

	m.LabelingInstructions = model.LabelingInstructions

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.AnnotationFormat = model.AnnotationFormat

	nn, e = model.DatasetSourceDetails.UnmarshalPolymorphicJSON(model.DatasetSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetSourceDetails = nn.(DatasetSourceDetails)
	} else {
		m.DatasetSourceDetails = nil
	}

	nn, e = model.DatasetFormatDetails.UnmarshalPolymorphicJSON(model.DatasetFormatDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetFormatDetails = nn.(DatasetFormatDetails)
	} else {
		m.DatasetFormatDetails = nil
	}

	m.LabelSet = model.LabelSet

	return
}
