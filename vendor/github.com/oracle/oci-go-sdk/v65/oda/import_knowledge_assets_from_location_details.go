// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportKnowledgeAssetsFromLocationDetails Properties that are required to create Knowledge Assets for files imported from a given location. If the import operation affects a single asset, specify assetMetadata to also update the asset.
type ImportKnowledgeAssetsFromLocationDetails struct {

	// The publicly-accessible object storage location
	Location *string `mandatory:"true" json:"location"`

	// Whether to override the assets with matching relative paths
	IsOverride *bool `mandatory:"false" json:"isOverride"`

	AssetMetadata *UpdateKnowledgeAssetDetails `mandatory:"false" json:"assetMetadata"`
}

func (m ImportKnowledgeAssetsFromLocationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportKnowledgeAssetsFromLocationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
