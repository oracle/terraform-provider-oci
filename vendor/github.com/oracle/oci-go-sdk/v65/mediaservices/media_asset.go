// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MediaAsset Represents the metadata associated with an asset that has been either produced by or registered with Media Services.
type MediaAsset struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The ID of the compartment containing the MediaAsset.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the MediaAsset.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the media asset.
	Type AssetTypeEnum `mandatory:"true" json:"type"`

	// The ID of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowId *string `mandatory:"false" json:"sourceMediaWorkflowId"`

	// The ID of the MediaWorkflowJob used to produce this asset.
	MediaWorkflowJobId *string `mandatory:"false" json:"mediaWorkflowJobId"`

	// The version of the MediaWorkflow used to produce this asset.
	SourceMediaWorkflowVersion *int64 `mandatory:"false" json:"sourceMediaWorkflowVersion"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time when the MediaAsset was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The ID of the parent asset from which this asset is derived.
	ParentMediaAssetId *string `mandatory:"false" json:"parentMediaAssetId"`

	// The ID of the senior most asset from which this asset is derived.
	MasterMediaAssetId *string `mandatory:"false" json:"masterMediaAssetId"`

	// The name of the object storage bucket where this represented asset is located.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The object storage namespace where this asset is located.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The object storage object name that identifies this asset.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// eTag of the underlying object storage object.
	ObjectEtag *string `mandatory:"false" json:"objectEtag"`

	// The time when the MediaAsset was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The start index for video segment files.
	SegmentRangeStartIndex *int64 `mandatory:"false" json:"segmentRangeStartIndex"`

	// The end index of video segment files.
	SegmentRangeEndIndex *int64 `mandatory:"false" json:"segmentRangeEndIndex"`

	// List of Metadata.
	Metadata []Metadata `mandatory:"false" json:"metadata"`

	// List of tags for the MediaAsset.
	MediaAssetTags []MediaAssetTag `mandatory:"false" json:"mediaAssetTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MediaAsset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaAsset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssetTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAssetTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
