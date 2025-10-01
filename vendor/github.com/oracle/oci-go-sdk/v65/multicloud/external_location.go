// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalLocation External location for CSP Region, CSP-Physical-AZ
type ExternalLocation struct {

	// CSP region corresponding to the given OCI region
	CspRegion *string `mandatory:"true" json:"cspRegion"`

	// CSP region display Name corresponding to the given OCI region
	CspRegionDisplayName *string `mandatory:"true" json:"cspRegionDisplayName"`

	// A mapping of OCI site group name to CSP physical availability zone name
	CspPhysicalAz *string `mandatory:"true" json:"cspPhysicalAz"`

	// User friendly display name for cspPhysicalAZ
	CspPhysicalAzDisplayName *string `mandatory:"true" json:"cspPhysicalAzDisplayName"`

	// The serviceName that externalLocation map object belongs to
	ServiceName SubscriptionTypeEnum `mandatory:"false" json:"serviceName,omitempty"`
}

func (m ExternalLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionTypeEnum(string(m.ServiceName)); !ok && m.ServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceName: %s. Supported values are: %s.", m.ServiceName, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
