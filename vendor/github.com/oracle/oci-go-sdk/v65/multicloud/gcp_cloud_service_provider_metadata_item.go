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

// GcpCloudServiceProviderMetadataItem GCP Cloud Service Provider metadata item.
type GcpCloudServiceProviderMetadataItem struct {

	// CSP resource anchor ID or name.
	ResourceAnchorName *string `mandatory:"true" json:"resourceAnchorName"`

	// GCP project number that was used for creating this resource anchor resource.
	ProjectNumber *string `mandatory:"true" json:"projectNumber"`

	// The Azure, AWS or GCP region.
	Region *string `mandatory:"false" json:"region"`

	// CSP resource anchor Uri.
	ResourceAnchorUri *string `mandatory:"false" json:"resourceAnchorUri"`
}

// GetRegion returns Region
func (m GcpCloudServiceProviderMetadataItem) GetRegion() *string {
	return m.Region
}

// GetResourceAnchorName returns ResourceAnchorName
func (m GcpCloudServiceProviderMetadataItem) GetResourceAnchorName() *string {
	return m.ResourceAnchorName
}

// GetResourceAnchorUri returns ResourceAnchorUri
func (m GcpCloudServiceProviderMetadataItem) GetResourceAnchorUri() *string {
	return m.ResourceAnchorUri
}

func (m GcpCloudServiceProviderMetadataItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GcpCloudServiceProviderMetadataItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GcpCloudServiceProviderMetadataItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGcpCloudServiceProviderMetadataItem GcpCloudServiceProviderMetadataItem
	s := struct {
		DiscriminatorParam string `json:"subscriptionType"`
		MarshalTypeGcpCloudServiceProviderMetadataItem
	}{
		"ORACLEDBATGOOGLE",
		(MarshalTypeGcpCloudServiceProviderMetadataItem)(m),
	}

	return json.Marshal(&s)
}
