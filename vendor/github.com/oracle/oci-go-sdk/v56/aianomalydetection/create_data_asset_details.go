// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateDataAssetDetails Parameters needed to create a new data asset.
type CreateDataAssetDetails struct {

	// The OCID for the data asset's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId *string `mandatory:"true" json:"projectId"`

	DataSourceDetails DataSourceDetails `mandatory:"true" json:"dataSourceDetails"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the Ai data asset
	Description *string `mandatory:"false" json:"description"`

	// OCID of Private Endpoint.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDataAssetDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDataAssetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                           `json:"displayName"`
		Description       *string                           `json:"description"`
		PrivateEndpointId *string                           `json:"privateEndpointId"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId     *string                           `json:"compartmentId"`
		ProjectId         *string                           `json:"projectId"`
		DataSourceDetails datasourcedetails                 `json:"dataSourceDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.PrivateEndpointId = model.PrivateEndpointId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	nn, e = model.DataSourceDetails.UnmarshalPolymorphicJSON(model.DataSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataSourceDetails = nn.(DataSourceDetails)
	} else {
		m.DataSourceDetails = nil
	}

	return
}
