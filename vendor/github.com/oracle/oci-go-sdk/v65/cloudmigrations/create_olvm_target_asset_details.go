// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateOlvmTargetAssetDetails Description of the OLVM target asset.
type CreateOlvmTargetAssetDetails struct {

	// OCID of the associated migration plan.
	MigrationPlanId *string `mandatory:"true" json:"migrationPlanId"`

	// A boolean indicating whether the asset should be migrated.
	IsExcludedFromExecution *bool `mandatory:"true" json:"isExcludedFromExecution"`

	// Microsoft license for the VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`
}

// GetMigrationPlanId returns MigrationPlanId
func (m CreateOlvmTargetAssetDetails) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m CreateOlvmTargetAssetDetails) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

func (m CreateOlvmTargetAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOlvmTargetAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOlvmTargetAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOlvmTargetAssetDetails CreateOlvmTargetAssetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateOlvmTargetAssetDetails
	}{
		"OLVM_INSTANCE",
		(MarshalTypeCreateOlvmTargetAssetDetails)(m),
	}

	return json.Marshal(&s)
}
