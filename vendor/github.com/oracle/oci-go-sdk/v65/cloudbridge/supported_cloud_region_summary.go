// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SupportedCloudRegionSummary Summary of the supported cloud region.
type SupportedCloudRegionSummary struct {

	// The asset source type associated with the supported cloud region.
	AssetSourceType AssetSourceTypeEnum `mandatory:"true" json:"assetSourceType"`

	// The supported cloud region name.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the supported cloud region.
	LifecycleState SupportedCloudRegionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m SupportedCloudRegionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportedCloudRegionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetSourceTypeEnum(string(m.AssetSourceType)); !ok && m.AssetSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetSourceType: %s. Supported values are: %s.", m.AssetSourceType, strings.Join(GetAssetSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSupportedCloudRegionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSupportedCloudRegionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
