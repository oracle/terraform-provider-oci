// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePipelineDetails Information with which to update a pipeline.
type UpdatePipelineDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// Metadata about this specific object.
	GetDescription() *string

	// The Oracle license model that applies to a Deployment.
	GetLicenseModel() LicenseModelEnum

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatepipelinedetails struct {
	JsonData     []byte
	DisplayName  *string                           `mandatory:"false" json:"displayName"`
	Description  *string                           `mandatory:"false" json:"description"`
	LicenseModel LicenseModelEnum                  `mandatory:"false" json:"licenseModel,omitempty"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	RecipeType   string                            `json:"recipeType"`
}

// UnmarshalJSON unmarshals json
func (m *updatepipelinedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatepipelinedetails updatepipelinedetails
	s := struct {
		Model Unmarshalerupdatepipelinedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.LicenseModel = s.Model.LicenseModel
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.RecipeType = s.Model.RecipeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatepipelinedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RecipeType {
	case "ZERO_ETL":
		mm := UpdateZeroEtlPipelineDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdatePipelineDetails: %s.", m.RecipeType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatepipelinedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updatepipelinedetails) GetDescription() *string {
	return m.Description
}

// GetLicenseModel returns LicenseModel
func (m updatepipelinedetails) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetFreeformTags returns FreeformTags
func (m updatepipelinedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatepipelinedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatepipelinedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatepipelinedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
