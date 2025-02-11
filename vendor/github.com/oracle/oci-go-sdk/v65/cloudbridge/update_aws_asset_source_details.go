// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAwsAssetSourceDetails AWS asset source update request.
type UpdateAwsAssetSourceDetails struct {

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	AssetsCompartmentId *string `mandatory:"false" json:"assetsCompartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the discovery schedule that is going to be assigned to an asset source.
	DiscoveryScheduleId *string `mandatory:"false" json:"discoveryScheduleId"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace/scope. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	DiscoveryCredentials *AssetSourceCredentials `mandatory:"false" json:"discoveryCredentials"`

	ReplicationCredentials *AssetSourceCredentials `mandatory:"false" json:"replicationCredentials"`

	// Flag indicating whether historical metrics are collected for assets, originating from this asset source.
	AreHistoricalMetricsCollected *bool `mandatory:"false" json:"areHistoricalMetricsCollected"`

	// Flag indicating whether real-time metrics are collected for assets, originating from this asset source.
	AreRealtimeMetricsCollected *bool `mandatory:"false" json:"areRealtimeMetricsCollected"`

	// Flag indicating whether cost data collection is enabled for assets, originating from this asset source.
	IsCostInformationCollected *bool `mandatory:"false" json:"isCostInformationCollected"`
}

// GetDisplayName returns DisplayName
func (m UpdateAwsAssetSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m UpdateAwsAssetSourceDetails) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

// GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m UpdateAwsAssetSourceDetails) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

// GetFreeformTags returns FreeformTags
func (m UpdateAwsAssetSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateAwsAssetSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m UpdateAwsAssetSourceDetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m UpdateAwsAssetSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAwsAssetSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAwsAssetSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAwsAssetSourceDetails UpdateAwsAssetSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateAwsAssetSourceDetails
	}{
		"AWS",
		(MarshalTypeUpdateAwsAssetSourceDetails)(m),
	}

	return json.Marshal(&s)
}
