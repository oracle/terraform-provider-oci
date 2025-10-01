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

// AzureCloudServiceProviderMetadataItem Azure Cloud Service Provider metadata item.
type AzureCloudServiceProviderMetadataItem struct {

	// CSP resource anchor ID or name.
	ResourceAnchorName *string `mandatory:"true" json:"resourceAnchorName"`

	// Azure resource group that was used for creating this resource.
	ResourceGroup *string `mandatory:"true" json:"resourceGroup"`

	// Azure subscription that was used for creating this resource.
	Subscription *string `mandatory:"true" json:"subscription"`

	// The Azure, AWS or GCP region.
	Region *string `mandatory:"false" json:"region"`

	// CSP resource anchor Uri.
	ResourceAnchorUri *string `mandatory:"false" json:"resourceAnchorUri"`
}

// GetRegion returns Region
func (m AzureCloudServiceProviderMetadataItem) GetRegion() *string {
	return m.Region
}

// GetResourceAnchorName returns ResourceAnchorName
func (m AzureCloudServiceProviderMetadataItem) GetResourceAnchorName() *string {
	return m.ResourceAnchorName
}

// GetResourceAnchorUri returns ResourceAnchorUri
func (m AzureCloudServiceProviderMetadataItem) GetResourceAnchorUri() *string {
	return m.ResourceAnchorUri
}

func (m AzureCloudServiceProviderMetadataItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AzureCloudServiceProviderMetadataItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AzureCloudServiceProviderMetadataItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAzureCloudServiceProviderMetadataItem AzureCloudServiceProviderMetadataItem
	s := struct {
		DiscriminatorParam string `json:"subscriptionType"`
		MarshalTypeAzureCloudServiceProviderMetadataItem
	}{
		"ORACLEDBATAZURE",
		(MarshalTypeAzureCloudServiceProviderMetadataItem)(m),
	}

	return json.Marshal(&s)
}
