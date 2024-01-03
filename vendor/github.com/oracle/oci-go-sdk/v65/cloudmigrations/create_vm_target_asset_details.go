// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVmTargetAssetDetails Description of the VM target asset.
type CreateVmTargetAssetDetails struct {

	// OCID of the associated migration plan.
	MigrationPlanId *string `mandatory:"true" json:"migrationPlanId"`

	// A boolean indicating whether the asset should be migrated.
	IsExcludedFromExecution *bool `mandatory:"true" json:"isExcludedFromExecution"`

	UserSpec *LaunchInstanceDetails `mandatory:"true" json:"userSpec"`

	// Performance of the block volumes.
	BlockVolumesPerformance *int `mandatory:"false" json:"blockVolumesPerformance"`

	// Microsoft license for the VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`

	// Preferred VM shape type that you provide.
	PreferredShapeType VmTargetAssetPreferredShapeTypeEnum `mandatory:"true" json:"preferredShapeType"`
}

// GetMigrationPlanId returns MigrationPlanId
func (m CreateVmTargetAssetDetails) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m CreateVmTargetAssetDetails) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

func (m CreateVmTargetAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVmTargetAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmTargetAssetPreferredShapeTypeEnum(string(m.PreferredShapeType)); !ok && m.PreferredShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredShapeType: %s. Supported values are: %s.", m.PreferredShapeType, strings.Join(GetVmTargetAssetPreferredShapeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateVmTargetAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateVmTargetAssetDetails CreateVmTargetAssetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateVmTargetAssetDetails
	}{
		"INSTANCE",
		(MarshalTypeCreateVmTargetAssetDetails)(m),
	}

	return json.Marshal(&s)
}
