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

// CreatePipelineDetails Details with which to create a pipeline.
type CreatePipelineDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	GetCompartmentId() *string

	// The Oracle license model that applies to a Deployment.
	GetLicenseModel() LicenseModelEnum

	GetSourceConnectionDetails() *SourcePipelineConnectionDetails

	GetTargetConnectionDetails() *TargetPipelineConnectionDetails

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createpipelinedetails struct {
	JsonData                []byte
	Description             *string                           `mandatory:"false" json:"description"`
	FreeformTags            map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags             map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Locks                   []ResourceLock                    `mandatory:"false" json:"locks"`
	DisplayName             *string                           `mandatory:"true" json:"displayName"`
	CompartmentId           *string                           `mandatory:"true" json:"compartmentId"`
	LicenseModel            LicenseModelEnum                  `mandatory:"true" json:"licenseModel"`
	SourceConnectionDetails *SourcePipelineConnectionDetails  `mandatory:"true" json:"sourceConnectionDetails"`
	TargetConnectionDetails *TargetPipelineConnectionDetails  `mandatory:"true" json:"targetConnectionDetails"`
	RecipeType              string                            `json:"recipeType"`
}

// UnmarshalJSON unmarshals json
func (m *createpipelinedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatepipelinedetails createpipelinedetails
	s := struct {
		Model Unmarshalercreatepipelinedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LicenseModel = s.Model.LicenseModel
	m.SourceConnectionDetails = s.Model.SourceConnectionDetails
	m.TargetConnectionDetails = s.Model.TargetConnectionDetails
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Locks = s.Model.Locks
	m.RecipeType = s.Model.RecipeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createpipelinedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RecipeType {
	case "ZERO_ETL":
		mm := CreateZeroEtlPipelineDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreatePipelineDetails: %s.", m.RecipeType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createpipelinedetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createpipelinedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createpipelinedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m createpipelinedetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetDisplayName returns DisplayName
func (m createpipelinedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createpipelinedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLicenseModel returns LicenseModel
func (m createpipelinedetails) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetSourceConnectionDetails returns SourceConnectionDetails
func (m createpipelinedetails) GetSourceConnectionDetails() *SourcePipelineConnectionDetails {
	return m.SourceConnectionDetails
}

// GetTargetConnectionDetails returns TargetConnectionDetails
func (m createpipelinedetails) GetTargetConnectionDetails() *TargetPipelineConnectionDetails {
	return m.TargetConnectionDetails
}

func (m createpipelinedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createpipelinedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
