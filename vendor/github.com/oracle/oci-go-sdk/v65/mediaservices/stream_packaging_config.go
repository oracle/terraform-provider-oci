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

// StreamPackagingConfig A stream packaging configuration for a Distribution Channel.
type StreamPackagingConfig interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// Compartment Identifier
	GetCompartmentId() *string

	// Unique identifier of the Distribution Channel that this stream packaging configuration belongs to.
	GetDistributionChannelId() *string

	// The name of the stream packaging configuration. Avoid entering confidential information.
	GetDisplayName() *string

	// The duration in seconds for each fragment.
	GetSegmentTimeInSeconds() *int

	GetEncryption() StreamPackagingConfigEncryption

	// The time when the Packaging Configuration was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time when the Packaging Configuration was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The current state of the Packaging Configuration.
	GetLifecycleState() StreamPackagingConfigLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type streampackagingconfig struct {
	JsonData              []byte
	Encryption            streampackagingconfigencryption         `mandatory:"false" json:"encryption"`
	TimeCreated           *common.SDKTime                         `mandatory:"false" json:"timeCreated"`
	TimeUpdated           *common.SDKTime                         `mandatory:"false" json:"timeUpdated"`
	LifecycleState        StreamPackagingConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags          map[string]string                       `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{}       `mandatory:"false" json:"definedTags"`
	SystemTags            map[string]map[string]interface{}       `mandatory:"false" json:"systemTags"`
	Locks                 []ResourceLock                          `mandatory:"false" json:"locks"`
	Id                    *string                                 `mandatory:"true" json:"id"`
	CompartmentId         *string                                 `mandatory:"true" json:"compartmentId"`
	DistributionChannelId *string                                 `mandatory:"true" json:"distributionChannelId"`
	DisplayName           *string                                 `mandatory:"true" json:"displayName"`
	SegmentTimeInSeconds  *int                                    `mandatory:"true" json:"segmentTimeInSeconds"`
	StreamPackagingFormat string                                  `json:"streamPackagingFormat"`
}

// UnmarshalJSON unmarshals json
func (m *streampackagingconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreampackagingconfig streampackagingconfig
	s := struct {
		Model Unmarshalerstreampackagingconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DistributionChannelId = s.Model.DistributionChannelId
	m.DisplayName = s.Model.DisplayName
	m.SegmentTimeInSeconds = s.Model.SegmentTimeInSeconds
	m.Encryption = s.Model.Encryption
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.StreamPackagingFormat = s.Model.StreamPackagingFormat

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streampackagingconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StreamPackagingFormat {
	case "HLS":
		mm := HlsStreamPackagingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DASH":
		mm := DashStreamPackagingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StreamPackagingConfig: %s.", m.StreamPackagingFormat)
		return *m, nil
	}
}

// GetEncryption returns Encryption
func (m streampackagingconfig) GetEncryption() streampackagingconfigencryption {
	return m.Encryption
}

// GetTimeCreated returns TimeCreated
func (m streampackagingconfig) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m streampackagingconfig) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m streampackagingconfig) GetLifecycleState() StreamPackagingConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m streampackagingconfig) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m streampackagingconfig) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m streampackagingconfig) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m streampackagingconfig) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m streampackagingconfig) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m streampackagingconfig) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDistributionChannelId returns DistributionChannelId
func (m streampackagingconfig) GetDistributionChannelId() *string {
	return m.DistributionChannelId
}

// GetDisplayName returns DisplayName
func (m streampackagingconfig) GetDisplayName() *string {
	return m.DisplayName
}

// GetSegmentTimeInSeconds returns SegmentTimeInSeconds
func (m streampackagingconfig) GetSegmentTimeInSeconds() *int {
	return m.SegmentTimeInSeconds
}

func (m streampackagingconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streampackagingconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStreamPackagingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamPackagingConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamPackagingConfigLifecycleStateEnum Enum with underlying type: string
type StreamPackagingConfigLifecycleStateEnum string

// Set of constants representing the allowable values for StreamPackagingConfigLifecycleStateEnum
const (
	StreamPackagingConfigLifecycleStateActive         StreamPackagingConfigLifecycleStateEnum = "ACTIVE"
	StreamPackagingConfigLifecycleStateNeedsAttention StreamPackagingConfigLifecycleStateEnum = "NEEDS_ATTENTION"
	StreamPackagingConfigLifecycleStateDeleted        StreamPackagingConfigLifecycleStateEnum = "DELETED"
)

var mappingStreamPackagingConfigLifecycleStateEnum = map[string]StreamPackagingConfigLifecycleStateEnum{
	"ACTIVE":          StreamPackagingConfigLifecycleStateActive,
	"NEEDS_ATTENTION": StreamPackagingConfigLifecycleStateNeedsAttention,
	"DELETED":         StreamPackagingConfigLifecycleStateDeleted,
}

var mappingStreamPackagingConfigLifecycleStateEnumLowerCase = map[string]StreamPackagingConfigLifecycleStateEnum{
	"active":          StreamPackagingConfigLifecycleStateActive,
	"needs_attention": StreamPackagingConfigLifecycleStateNeedsAttention,
	"deleted":         StreamPackagingConfigLifecycleStateDeleted,
}

// GetStreamPackagingConfigLifecycleStateEnumValues Enumerates the set of values for StreamPackagingConfigLifecycleStateEnum
func GetStreamPackagingConfigLifecycleStateEnumValues() []StreamPackagingConfigLifecycleStateEnum {
	values := make([]StreamPackagingConfigLifecycleStateEnum, 0)
	for _, v := range mappingStreamPackagingConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamPackagingConfigLifecycleStateEnumStringValues Enumerates the set of values in String for StreamPackagingConfigLifecycleStateEnum
func GetStreamPackagingConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingStreamPackagingConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamPackagingConfigLifecycleStateEnum(val string) (StreamPackagingConfigLifecycleStateEnum, bool) {
	enum, ok := mappingStreamPackagingConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// StreamPackagingConfigStreamPackagingFormatEnum Enum with underlying type: string
type StreamPackagingConfigStreamPackagingFormatEnum string

// Set of constants representing the allowable values for StreamPackagingConfigStreamPackagingFormatEnum
const (
	StreamPackagingConfigStreamPackagingFormatHls  StreamPackagingConfigStreamPackagingFormatEnum = "HLS"
	StreamPackagingConfigStreamPackagingFormatDash StreamPackagingConfigStreamPackagingFormatEnum = "DASH"
)

var mappingStreamPackagingConfigStreamPackagingFormatEnum = map[string]StreamPackagingConfigStreamPackagingFormatEnum{
	"HLS":  StreamPackagingConfigStreamPackagingFormatHls,
	"DASH": StreamPackagingConfigStreamPackagingFormatDash,
}

var mappingStreamPackagingConfigStreamPackagingFormatEnumLowerCase = map[string]StreamPackagingConfigStreamPackagingFormatEnum{
	"hls":  StreamPackagingConfigStreamPackagingFormatHls,
	"dash": StreamPackagingConfigStreamPackagingFormatDash,
}

// GetStreamPackagingConfigStreamPackagingFormatEnumValues Enumerates the set of values for StreamPackagingConfigStreamPackagingFormatEnum
func GetStreamPackagingConfigStreamPackagingFormatEnumValues() []StreamPackagingConfigStreamPackagingFormatEnum {
	values := make([]StreamPackagingConfigStreamPackagingFormatEnum, 0)
	for _, v := range mappingStreamPackagingConfigStreamPackagingFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamPackagingConfigStreamPackagingFormatEnumStringValues Enumerates the set of values in String for StreamPackagingConfigStreamPackagingFormatEnum
func GetStreamPackagingConfigStreamPackagingFormatEnumStringValues() []string {
	return []string{
		"HLS",
		"DASH",
	}
}

// GetMappingStreamPackagingConfigStreamPackagingFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamPackagingConfigStreamPackagingFormatEnum(val string) (StreamPackagingConfigStreamPackagingFormatEnum, bool) {
	enum, ok := mappingStreamPackagingConfigStreamPackagingFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
