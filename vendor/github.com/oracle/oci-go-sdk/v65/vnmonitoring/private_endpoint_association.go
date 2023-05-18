// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointAssociation A summary of private endpoint information. This object is returned when listing the private
// endpoints associated with a given endpoint service.
type PrivateEndpointAssociation struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	Id *string `mandatory:"false" json:"id"`

	// The private IP address (in the customer's VCN) that represents the access point for the
	// associated endpoint service.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records use
	// this FQDN.
	// For important information about how this attribute is used, see the discussion
	// of DNS and FQDNs in PrivateEndpoint.
	// Example: `xyz.oraclecloud.com`
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`

	// A list of additional FQDNs that you can provide along with endpointFqdn. These FQDNs are added to the
	// customer VCN's DNS record. For more information, see the discussion of DNS and FQDNs in
	// PrivateEndpoint.
	AdditionalFqdns []string `mandatory:"false" json:"additionalFqdns"`

	ReverseConnectionConfiguration *ReverseConnectionConfiguration `mandatory:"false" json:"reverseConnectionConfiguration"`
}

func (m PrivateEndpointAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
