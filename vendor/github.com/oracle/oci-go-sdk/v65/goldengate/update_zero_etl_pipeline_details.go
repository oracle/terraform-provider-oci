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

// UpdateZeroEtlPipelineDetails Information to update for an existing ZeroETL pipeline.
type UpdateZeroEtlPipelineDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ProcessOptions *ProcessOptions `mandatory:"false" json:"processOptions"`

	// Mapping for source/target schema/tables for the pipeline data replication.
	MappingRules []MappingRule `mandatory:"false" json:"mappingRules"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
}

// GetDisplayName returns DisplayName
func (m UpdateZeroEtlPipelineDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateZeroEtlPipelineDetails) GetDescription() *string {
	return m.Description
}

// GetLicenseModel returns LicenseModel
func (m UpdateZeroEtlPipelineDetails) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetFreeformTags returns FreeformTags
func (m UpdateZeroEtlPipelineDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateZeroEtlPipelineDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateZeroEtlPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateZeroEtlPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateZeroEtlPipelineDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateZeroEtlPipelineDetails UpdateZeroEtlPipelineDetails
	s := struct {
		DiscriminatorParam string `json:"recipeType"`
		MarshalTypeUpdateZeroEtlPipelineDetails
	}{
		"ZERO_ETL",
		(MarshalTypeUpdateZeroEtlPipelineDetails)(m),
	}

	return json.Marshal(&s)
}
