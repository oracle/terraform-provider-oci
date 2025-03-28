// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMediaAssetDetails The information about new MediaAsset.
type CreateMediaAssetDetails struct {

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the media asset.
	Type AssetTypeEnum `mandatory:"true" json:"type"`

	// The ID of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowId *string `mandatory:"false" json:"sourceMediaWorkflowId"`

	// The ID of the MediaWorkflowJob used to produce this asset.
	MediaWorkflowJobId *string `mandatory:"false" json:"mediaWorkflowJobId"`

	// The version of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowVersion *int64 `mandatory:"false" json:"sourceMediaWorkflowVersion"`

	// Display name for the Media Asset. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The ID of the parent asset from which this asset is derived.
	ParentMediaAssetId *string `mandatory:"false" json:"parentMediaAssetId"`

	// The ID of the senior most asset from which this asset is derived.
	MasterMediaAssetId *string `mandatory:"false" json:"masterMediaAssetId"`

	// The name of the object storage bucket where this asset is located.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The object storage namespace where this asset is located.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The object storage object name that identifies this asset.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// eTag of the underlying object storage object.
	ObjectEtag *string `mandatory:"false" json:"objectEtag"`

	// List of Metadata.
	Metadata []Metadata `mandatory:"false" json:"metadata"`

	// The start index for video segment files.
	SegmentRangeStartIndex *int64 `mandatory:"false" json:"segmentRangeStartIndex"`

	// The end index for video segment files.
	SegmentRangeEndIndex *int64 `mandatory:"false" json:"segmentRangeEndIndex"`

	// list of tags for the MediaAsset.
	MediaAssetTags []MediaAssetTag `mandatory:"false" json:"mediaAssetTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m CreateMediaAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMediaAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAssetTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
