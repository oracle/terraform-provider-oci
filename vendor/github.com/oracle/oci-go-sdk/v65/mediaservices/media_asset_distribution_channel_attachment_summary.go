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

// MediaAssetDistributionChannelAttachmentSummary Summary of the MediaAssetDistributionChannelAttachment.
type MediaAssetDistributionChannelAttachmentSummary struct {

	// OCID of associated media asset.
	MediaAssetId *string `mandatory:"true" json:"mediaAssetId"`

	// OCID of associated Distribution Channel.
	DistributionChannelId *string `mandatory:"true" json:"distributionChannelId"`

	// Version number of the attachment.
	Version *int64 `mandatory:"true" json:"version"`

	// Lifecycle state of the attachment.
	LifecycleState MediaAssetDistributionChannelAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The identifier for the metadata.
	MetadataRef *string `mandatory:"true" json:"metadataRef"`

	// Display name for the MediaAssetDistributionChannelAttachment. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The ingest MediaWorkflowJob ID that created this attachment.
	MediaWorkflowJobId *string `mandatory:"false" json:"mediaWorkflowJobId"`
}

func (m MediaAssetDistributionChannelAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaAssetDistributionChannelAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
