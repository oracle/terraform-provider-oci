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

// DashStreamPackagingConfig Configuration fields for a DASH Packaging Configuration.
type DashStreamPackagingConfig struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier of the Distribution Channel that this stream packaging configuration belongs to.
	DistributionChannelId *string `mandatory:"true" json:"distributionChannelId"`

	// The name of the stream packaging configuration. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The duration in seconds for each fragment.
	SegmentTimeInSeconds *int `mandatory:"true" json:"segmentTimeInSeconds"`

	Encryption StreamPackagingConfigEncryption `mandatory:"false" json:"encryption"`

	// The time when the Packaging Configuration was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when the Packaging Configuration was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Packaging Configuration.
	LifecycleState StreamPackagingConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m DashStreamPackagingConfig) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DashStreamPackagingConfig) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDistributionChannelId returns DistributionChannelId
func (m DashStreamPackagingConfig) GetDistributionChannelId() *string {
	return m.DistributionChannelId
}

// GetDisplayName returns DisplayName
func (m DashStreamPackagingConfig) GetDisplayName() *string {
	return m.DisplayName
}

// GetSegmentTimeInSeconds returns SegmentTimeInSeconds
func (m DashStreamPackagingConfig) GetSegmentTimeInSeconds() *int {
	return m.SegmentTimeInSeconds
}

// GetEncryption returns Encryption
func (m DashStreamPackagingConfig) GetEncryption() StreamPackagingConfigEncryption {
	return m.Encryption
}

// GetTimeCreated returns TimeCreated
func (m DashStreamPackagingConfig) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DashStreamPackagingConfig) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m DashStreamPackagingConfig) GetLifecycleState() StreamPackagingConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m DashStreamPackagingConfig) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m DashStreamPackagingConfig) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m DashStreamPackagingConfig) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DashStreamPackagingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DashStreamPackagingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStreamPackagingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStreamPackagingConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DashStreamPackagingConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDashStreamPackagingConfig DashStreamPackagingConfig
	s := struct {
		DiscriminatorParam string `json:"streamPackagingFormat"`
		MarshalTypeDashStreamPackagingConfig
	}{
		"DASH",
		(MarshalTypeDashStreamPackagingConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DashStreamPackagingConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Encryption            streampackagingconfigencryption         `json:"encryption"`
		TimeCreated           *common.SDKTime                         `json:"timeCreated"`
		TimeUpdated           *common.SDKTime                         `json:"timeUpdated"`
		LifecycleState        StreamPackagingConfigLifecycleStateEnum `json:"lifecycleState"`
		FreeformTags          map[string]string                       `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}       `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}       `json:"systemTags"`
		Id                    *string                                 `json:"id"`
		CompartmentId         *string                                 `json:"compartmentId"`
		DistributionChannelId *string                                 `json:"distributionChannelId"`
		DisplayName           *string                                 `json:"displayName"`
		SegmentTimeInSeconds  *int                                    `json:"segmentTimeInSeconds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Encryption.UnmarshalPolymorphicJSON(model.Encryption.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Encryption = nn.(StreamPackagingConfigEncryption)
	} else {
		m.Encryption = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DistributionChannelId = model.DistributionChannelId

	m.DisplayName = model.DisplayName

	m.SegmentTimeInSeconds = model.SegmentTimeInSeconds

	return
}
