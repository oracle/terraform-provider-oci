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

// UpdateOlvmTargetAssetDetails Description of the OLVM target asset.
type UpdateOlvmTargetAssetDetails struct {

	// A boolean indicating whether the asset should be migrated.
	IsExcludedFromExecution *bool `mandatory:"false" json:"isExcludedFromExecution"`

	// Microsoft license for the VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m UpdateOlvmTargetAssetDetails) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

func (m UpdateOlvmTargetAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOlvmTargetAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOlvmTargetAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOlvmTargetAssetDetails UpdateOlvmTargetAssetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOlvmTargetAssetDetails
	}{
		"OLVM_INSTANCE",
		(MarshalTypeUpdateOlvmTargetAssetDetails)(m),
	}

	return json.Marshal(&s)
}
