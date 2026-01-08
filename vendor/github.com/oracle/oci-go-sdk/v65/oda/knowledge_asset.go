// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// KnowledgeAsset Metadata for a Knowledge Asset resource.
type KnowledgeAsset struct {

	// Unique immutable identifier that was assigned when the resource was created.
	Id *string `mandatory:"true" json:"id"`

	// The resource's name. Taken from the relative path of the corresponding file.
	Name *string `mandatory:"true" json:"name"`

	// End-user-friendly name for this resource, to be used whenever the asset's name is displayed. Defaults to name if not provided.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// File content type of the asset
	FileContentType AssetFileContentTypeEnum `mandatory:"true" json:"fileContentType"`

	// MIME type of the asset's file
	Format *string `mandatory:"true" json:"format"`

	// Lifecycle state of the asset
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Origin of the asset's file
	Source AssetSourceEnum `mandatory:"true" json:"source"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Schema for the asset that is applicable for the file content type/format
	Schema *string `mandatory:"false" json:"schema"`

	// End-user-friendly URL for this resource, to be used whenever the asset's source is displayed.
	KnowledgeSourceUrl *string `mandatory:"false" json:"knowledgeSourceUrl"`

	// Custom unique identifier for this resource
	ExternalIdentifier *string `mandatory:"false" json:"externalIdentifier"`
}

func (m KnowledgeAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KnowledgeAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetFileContentTypeEnum(string(m.FileContentType)); !ok && m.FileContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FileContentType: %s. Supported values are: %s.", m.FileContentType, strings.Join(GetAssetFileContentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssetSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetAssetSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
