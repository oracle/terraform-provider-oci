// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateZeroEtlPipelineDetails Creation details for a new ZeroETL pipeline.
type CreateZeroEtlPipelineDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceConnectionDetails *SourcePipelineConnectionDetails `mandatory:"true" json:"sourceConnectionDetails"`

	TargetConnectionDetails *TargetPipelineConnectionDetails `mandatory:"true" json:"targetConnectionDetails"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	ProcessOptions *ProcessOptions `mandatory:"false" json:"processOptions"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`
}

// GetDisplayName returns DisplayName
func (m CreateZeroEtlPipelineDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateZeroEtlPipelineDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateZeroEtlPipelineDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLicenseModel returns LicenseModel
func (m CreateZeroEtlPipelineDetails) GetLicenseModel() LicenseModelEnum {
	return m.LicenseModel
}

// GetFreeformTags returns FreeformTags
func (m CreateZeroEtlPipelineDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateZeroEtlPipelineDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m CreateZeroEtlPipelineDetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetSourceConnectionDetails returns SourceConnectionDetails
func (m CreateZeroEtlPipelineDetails) GetSourceConnectionDetails() *SourcePipelineConnectionDetails {
	return m.SourceConnectionDetails
}

// GetTargetConnectionDetails returns TargetConnectionDetails
func (m CreateZeroEtlPipelineDetails) GetTargetConnectionDetails() *TargetPipelineConnectionDetails {
	return m.TargetConnectionDetails
}

func (m CreateZeroEtlPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateZeroEtlPipelineDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateZeroEtlPipelineDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZeroEtlPipelineDetails CreateZeroEtlPipelineDetails
	s := struct {
		DiscriminatorParam string `json:"recipeType"`
		MarshalTypeCreateZeroEtlPipelineDetails
	}{
		"ZERO_ETL",
		(MarshalTypeCreateZeroEtlPipelineDetails)(m),
	}

	return json.Marshal(&s)
}
