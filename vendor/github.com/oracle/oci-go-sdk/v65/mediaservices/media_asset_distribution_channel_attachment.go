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

// MediaAssetDistributionChannelAttachment Attachment between MediaAsset and streaming DistributionChannel.
type MediaAssetDistributionChannelAttachment struct {

	// OCID of associated Distribution Channel.
	DistributionChannelId *string `mandatory:"true" json:"distributionChannelId"`

	// Version of the attachment.
	Version *int64 `mandatory:"true" json:"version"`

	// Lifecycle state of the attachment.
	LifecycleState MediaAssetDistributionChannelAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The identifier for the metadata.
	MetadataRef *string `mandatory:"true" json:"metadataRef"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The ingest MediaWorkflowJob ID that created this attachment.
	MediaWorkflowJobId *string `mandatory:"false" json:"mediaWorkflowJobId"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m MediaAssetDistributionChannelAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaAssetDistributionChannelAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MediaAssetDistributionChannelAttachmentLifecycleStateEnum Enum with underlying type: string
type MediaAssetDistributionChannelAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for MediaAssetDistributionChannelAttachmentLifecycleStateEnum
const (
	MediaAssetDistributionChannelAttachmentLifecycleStateCreating       MediaAssetDistributionChannelAttachmentLifecycleStateEnum = "CREATING"
	MediaAssetDistributionChannelAttachmentLifecycleStateActive         MediaAssetDistributionChannelAttachmentLifecycleStateEnum = "ACTIVE"
	MediaAssetDistributionChannelAttachmentLifecycleStateNeedsAttention MediaAssetDistributionChannelAttachmentLifecycleStateEnum = "NEEDS_ATTENTION"
	MediaAssetDistributionChannelAttachmentLifecycleStateUpdating       MediaAssetDistributionChannelAttachmentLifecycleStateEnum = "UPDATING"
)

var mappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum = map[string]MediaAssetDistributionChannelAttachmentLifecycleStateEnum{
	"CREATING":        MediaAssetDistributionChannelAttachmentLifecycleStateCreating,
	"ACTIVE":          MediaAssetDistributionChannelAttachmentLifecycleStateActive,
	"NEEDS_ATTENTION": MediaAssetDistributionChannelAttachmentLifecycleStateNeedsAttention,
	"UPDATING":        MediaAssetDistributionChannelAttachmentLifecycleStateUpdating,
}

var mappingMediaAssetDistributionChannelAttachmentLifecycleStateEnumLowerCase = map[string]MediaAssetDistributionChannelAttachmentLifecycleStateEnum{
	"creating":        MediaAssetDistributionChannelAttachmentLifecycleStateCreating,
	"active":          MediaAssetDistributionChannelAttachmentLifecycleStateActive,
	"needs_attention": MediaAssetDistributionChannelAttachmentLifecycleStateNeedsAttention,
	"updating":        MediaAssetDistributionChannelAttachmentLifecycleStateUpdating,
}

// GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumValues Enumerates the set of values for MediaAssetDistributionChannelAttachmentLifecycleStateEnum
func GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumValues() []MediaAssetDistributionChannelAttachmentLifecycleStateEnum {
	values := make([]MediaAssetDistributionChannelAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for MediaAssetDistributionChannelAttachmentLifecycleStateEnum
func GetMediaAssetDistributionChannelAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaAssetDistributionChannelAttachmentLifecycleStateEnum(val string) (MediaAssetDistributionChannelAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingMediaAssetDistributionChannelAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
