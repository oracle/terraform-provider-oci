// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwsCloudServiceProviderMetadataItem Aws Cloud Service Provider metadata item.
type AwsCloudServiceProviderMetadataItem struct {

	// OCI resource anchor name.
	ResourceAnchorName *string `mandatory:"true" json:"resourceAnchorName"`

	// AWS accountId that was used for creating this resource anchor resource.
	AccountId *string `mandatory:"true" json:"accountId"`

	// The Azure, AWS or GCP region.
	Region *string `mandatory:"false" json:"region"`

	// CSP resource anchor ID.
	CspResourceAnchorId *string `mandatory:"false" json:"cspResourceAnchorId"`

	// CSP resource anchor name.
	CspResourceAnchorName *string `mandatory:"false" json:"cspResourceAnchorName"`

	// CSP resource anchor Uri.
	ResourceAnchorUri *string `mandatory:"false" json:"resourceAnchorUri"`

	// CSP Specific Additional Properties, AzureSubnetId for Azure
	CspAdditionalProperties map[string]string `mandatory:"false" json:"cspAdditionalProperties"`
}

// GetRegion returns Region
func (m AwsCloudServiceProviderMetadataItem) GetRegion() *string {
	return m.Region
}

// GetResourceAnchorName returns ResourceAnchorName
func (m AwsCloudServiceProviderMetadataItem) GetResourceAnchorName() *string {
	return m.ResourceAnchorName
}

// GetCspResourceAnchorId returns CspResourceAnchorId
func (m AwsCloudServiceProviderMetadataItem) GetCspResourceAnchorId() *string {
	return m.CspResourceAnchorId
}

// GetCspResourceAnchorName returns CspResourceAnchorName
func (m AwsCloudServiceProviderMetadataItem) GetCspResourceAnchorName() *string {
	return m.CspResourceAnchorName
}

// GetResourceAnchorUri returns ResourceAnchorUri
func (m AwsCloudServiceProviderMetadataItem) GetResourceAnchorUri() *string {
	return m.ResourceAnchorUri
}

// GetCspAdditionalProperties returns CspAdditionalProperties
func (m AwsCloudServiceProviderMetadataItem) GetCspAdditionalProperties() map[string]string {
	return m.CspAdditionalProperties
}

func (m AwsCloudServiceProviderMetadataItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsCloudServiceProviderMetadataItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AwsCloudServiceProviderMetadataItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAwsCloudServiceProviderMetadataItem AwsCloudServiceProviderMetadataItem
	s := struct {
		DiscriminatorParam string `json:"subscriptionType"`
		MarshalTypeAwsCloudServiceProviderMetadataItem
	}{
		"ORACLEDBATAWS",
		(MarshalTypeAwsCloudServiceProviderMetadataItem)(m),
	}

	return json.Marshal(&s)
}
