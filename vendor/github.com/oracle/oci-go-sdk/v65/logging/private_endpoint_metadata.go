// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointMetadata Response-time metadata for private endpoint connectivity.
// Present only in GET/LIST responses when endpointMode=PRIVATE_ENDPOINT and provisioning succeeded.
// Customers use rceIpAddress to configure NSG/SecurityList ingress rules.
type PrivateEndpointMetadata struct {

	// RCE IP address (customer should allow ingress from this IP).
	RceIpAddress *string `mandatory:"true" json:"rceIpAddress"`

	// DNS proxy IP used by Logging DP custom resolver.
	DnsProxyIp *string `mandatory:"true" json:"dnsProxyIp"`

	// OCID of the managed Private Endpoint created in the customer tenancy.
	PrivateEndpointId *string `mandatory:"false" json:"privateEndpointId"`

	// Customer subnet OCID where the PE/RCE is provisioned.
	SubnetId *string `mandatory:"false" json:"subnetId"`
}

func (m PrivateEndpointMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
