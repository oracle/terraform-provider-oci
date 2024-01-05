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

// StreamDistributionChannel Channel used for delivering video streams to the end-users.
type StreamDistributionChannel struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Stream Distribution Channel display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique domain name of the Distribution Channel.
	DomainName *string `mandatory:"false" json:"domainName"`

	// The time when the Stream Distribution Channel was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the Stream Distribution Channel was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Stream Distribution Channel.
	LifecycleState StreamDistributionChannelLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m StreamDistributionChannel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamDistributionChannel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStreamDistributionChannelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamDistributionChannelLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamDistributionChannelLifecycleStateEnum Enum with underlying type: string
type StreamDistributionChannelLifecycleStateEnum string

// Set of constants representing the allowable values for StreamDistributionChannelLifecycleStateEnum
const (
	StreamDistributionChannelLifecycleStateActive         StreamDistributionChannelLifecycleStateEnum = "ACTIVE"
	StreamDistributionChannelLifecycleStateNeedsAttention StreamDistributionChannelLifecycleStateEnum = "NEEDS_ATTENTION"
	StreamDistributionChannelLifecycleStateDeleted        StreamDistributionChannelLifecycleStateEnum = "DELETED"
)

var mappingStreamDistributionChannelLifecycleStateEnum = map[string]StreamDistributionChannelLifecycleStateEnum{
	"ACTIVE":          StreamDistributionChannelLifecycleStateActive,
	"NEEDS_ATTENTION": StreamDistributionChannelLifecycleStateNeedsAttention,
	"DELETED":         StreamDistributionChannelLifecycleStateDeleted,
}

var mappingStreamDistributionChannelLifecycleStateEnumLowerCase = map[string]StreamDistributionChannelLifecycleStateEnum{
	"active":          StreamDistributionChannelLifecycleStateActive,
	"needs_attention": StreamDistributionChannelLifecycleStateNeedsAttention,
	"deleted":         StreamDistributionChannelLifecycleStateDeleted,
}

// GetStreamDistributionChannelLifecycleStateEnumValues Enumerates the set of values for StreamDistributionChannelLifecycleStateEnum
func GetStreamDistributionChannelLifecycleStateEnumValues() []StreamDistributionChannelLifecycleStateEnum {
	values := make([]StreamDistributionChannelLifecycleStateEnum, 0)
	for _, v := range mappingStreamDistributionChannelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamDistributionChannelLifecycleStateEnumStringValues Enumerates the set of values in String for StreamDistributionChannelLifecycleStateEnum
func GetStreamDistributionChannelLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingStreamDistributionChannelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamDistributionChannelLifecycleStateEnum(val string) (StreamDistributionChannelLifecycleStateEnum, bool) {
	enum, ok := mappingStreamDistributionChannelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
