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

// CloudServiceProviderMetadataItem Cloud Service Provider metadata item.
// Warning - In future this object can change to generic object with future Cloud Service Provider based on
// CloudServiceProvider field. This can be one of CSP provider type Azure, GCP and AWS.
type CloudServiceProviderMetadataItem interface {

	// CSP resource anchor ID or name.
	GetResourceAnchorName() *string

	// The Azure, AWS or GCP region.
	GetRegion() *string

	// CSP resource anchor Uri.
	GetResourceAnchorUri() *string
}

type cloudserviceprovidermetadataitem struct {
	JsonData           []byte
	Region             *string `mandatory:"false" json:"region"`
	ResourceAnchorUri  *string `mandatory:"false" json:"resourceAnchorUri"`
	ResourceAnchorName *string `mandatory:"true" json:"resourceAnchorName"`
	SubscriptionType   string  `json:"subscriptionType"`
}

// UnmarshalJSON unmarshals json
func (m *cloudserviceprovidermetadataitem) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercloudserviceprovidermetadataitem cloudserviceprovidermetadataitem
	s := struct {
		Model Unmarshalercloudserviceprovidermetadataitem
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ResourceAnchorName = s.Model.ResourceAnchorName
	m.Region = s.Model.Region
	m.ResourceAnchorUri = s.Model.ResourceAnchorUri
	m.SubscriptionType = s.Model.SubscriptionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cloudserviceprovidermetadataitem) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SubscriptionType {
	case "ORACLEDBATAZURE":
		mm := AzureCloudServiceProviderMetadataItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDBATGOOGLE":
		mm := GcpCloudServiceProviderMetadataItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLEDBATAWS":
		mm := AwsCloudServiceProviderMetadataItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloudServiceProviderMetadataItem: %s.", m.SubscriptionType)
		return *m, nil
	}
}

// GetRegion returns Region
func (m cloudserviceprovidermetadataitem) GetRegion() *string {
	return m.Region
}

// GetResourceAnchorUri returns ResourceAnchorUri
func (m cloudserviceprovidermetadataitem) GetResourceAnchorUri() *string {
	return m.ResourceAnchorUri
}

// GetResourceAnchorName returns ResourceAnchorName
func (m cloudserviceprovidermetadataitem) GetResourceAnchorName() *string {
	return m.ResourceAnchorName
}

func (m cloudserviceprovidermetadataitem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cloudserviceprovidermetadataitem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
