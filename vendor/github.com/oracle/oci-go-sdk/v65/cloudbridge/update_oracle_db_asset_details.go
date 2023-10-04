// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleDbAssetDetails The information of Oracle DB asset to be updated.
type UpdateOracleDbAssetDetails struct {

	// Asset display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// List of asset source OCID.
	AssetSourceIds []string `mandatory:"false" json:"assetSourceIds"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	OracleDbProperties *OracleDbProperties `mandatory:"false" json:"oracleDbProperties"`
}

//GetDisplayName returns DisplayName
func (m UpdateOracleDbAssetDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetAssetSourceIds returns AssetSourceIds
func (m UpdateOracleDbAssetDetails) GetAssetSourceIds() []string {
	return m.AssetSourceIds
}

//GetFreeformTags returns FreeformTags
func (m UpdateOracleDbAssetDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateOracleDbAssetDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOracleDbAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleDbAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleDbAssetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleDbAssetDetails UpdateOracleDbAssetDetails
	s := struct {
		DiscriminatorParam string `json:"assetType"`
		MarshalTypeUpdateOracleDbAssetDetails
	}{
		"ORACLE_DB",
		(MarshalTypeUpdateOracleDbAssetDetails)(m),
	}

	return json.Marshal(&s)
}
