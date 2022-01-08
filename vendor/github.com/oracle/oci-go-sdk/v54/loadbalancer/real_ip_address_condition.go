// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// RealIpAddressCondition An access control rule condition that requires a match for real/original client IP coming in the
// HTTP request with the HTTP header name and CidrBlocks resource or IP ranges
type RealIpAddressCondition struct {

	// A CidrBlocks resource name containing the CIDR Block or IP range for matching against the source IP in the request.
	AttributeValue *string `mandatory:"true" json:"attributeValue"`

	// A header name that conforms to RFC 7230.
	// Example: `example_header_name`
	HeaderName *string `mandatory:"false" json:"headerName"`

	// Index of the IP address from left to be matched when multiple IP addresses appears in the header value.
	// In case of multiple addresses, it is expected that addresses are separated by a comma and space eg. 172.31.4.1, 192.168.21.4.
	// Default is the left most IP address in the header value.
	Index *int `mandatory:"false" json:"index"`
}

func (m RealIpAddressCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealIpAddressCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RealIpAddressCondition) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRealIpAddressCondition RealIpAddressCondition
	s := struct {
		DiscriminatorParam string `json:"attributeName"`
		MarshalTypeRealIpAddressCondition
	}{
		"REAL_IP_ADDRESS",
		(MarshalTypeRealIpAddressCondition)(m),
	}

	return json.Marshal(&s)
}
