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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamCdnConfig Configuration used for integrating with a CDN.
type StreamCdnConfig struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The CDN Configuration identifier or display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Distribution Channel Identifier.
	DistributionChannelId *string `mandatory:"true" json:"distributionChannelId"`

	// Whether publishing to CDN is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	Config StreamCdnConfigSection `mandatory:"true" json:"config"`

	// The time when the CDN Config was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the CDN Config was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the CDN Configuration.
	LifecycleState StreamCdnConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m StreamCdnConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamCdnConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStreamCdnConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamCdnConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StreamCdnConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState        StreamCdnConfigLifecycleStateEnum `json:"lifecycleState"`
		LifecyleDetails       *string                           `json:"lifecyleDetails"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		Locks                 []ResourceLock                    `json:"locks"`
		Id                    *string                           `json:"id"`
		DisplayName           *string                           `json:"displayName"`
		CompartmentId         *string                           `json:"compartmentId"`
		DistributionChannelId *string                           `json:"distributionChannelId"`
		IsEnabled             *bool                             `json:"isEnabled"`
		Config                streamcdnconfigsection            `json:"config"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecyleDetails = model.LifecyleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.DistributionChannelId = model.DistributionChannelId

	m.IsEnabled = model.IsEnabled

	nn, e = model.Config.UnmarshalPolymorphicJSON(model.Config.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Config = nn.(StreamCdnConfigSection)
	} else {
		m.Config = nil
	}

	return
}

// StreamCdnConfigLifecycleStateEnum Enum with underlying type: string
type StreamCdnConfigLifecycleStateEnum string

// Set of constants representing the allowable values for StreamCdnConfigLifecycleStateEnum
const (
	StreamCdnConfigLifecycleStateActive         StreamCdnConfigLifecycleStateEnum = "ACTIVE"
	StreamCdnConfigLifecycleStateNeedsAttention StreamCdnConfigLifecycleStateEnum = "NEEDS_ATTENTION"
	StreamCdnConfigLifecycleStateDeleted        StreamCdnConfigLifecycleStateEnum = "DELETED"
)

var mappingStreamCdnConfigLifecycleStateEnum = map[string]StreamCdnConfigLifecycleStateEnum{
	"ACTIVE":          StreamCdnConfigLifecycleStateActive,
	"NEEDS_ATTENTION": StreamCdnConfigLifecycleStateNeedsAttention,
	"DELETED":         StreamCdnConfigLifecycleStateDeleted,
}

var mappingStreamCdnConfigLifecycleStateEnumLowerCase = map[string]StreamCdnConfigLifecycleStateEnum{
	"active":          StreamCdnConfigLifecycleStateActive,
	"needs_attention": StreamCdnConfigLifecycleStateNeedsAttention,
	"deleted":         StreamCdnConfigLifecycleStateDeleted,
}

// GetStreamCdnConfigLifecycleStateEnumValues Enumerates the set of values for StreamCdnConfigLifecycleStateEnum
func GetStreamCdnConfigLifecycleStateEnumValues() []StreamCdnConfigLifecycleStateEnum {
	values := make([]StreamCdnConfigLifecycleStateEnum, 0)
	for _, v := range mappingStreamCdnConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamCdnConfigLifecycleStateEnumStringValues Enumerates the set of values in String for StreamCdnConfigLifecycleStateEnum
func GetStreamCdnConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingStreamCdnConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamCdnConfigLifecycleStateEnum(val string) (StreamCdnConfigLifecycleStateEnum, bool) {
	enum, ok := mappingStreamCdnConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
