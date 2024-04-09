// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DnsHealthCheckerDetails DNS healthcheck configurations.
type DnsHealthCheckerDetails struct {

	// The absolute fully-qualified domain name to perform periodic DNS queries.
	// If not provided, an extra dot will be added at the end of a domain name during the query.
	DomainName *string `mandatory:"true" json:"domainName"`

	// DNS transport protocol; either UDP or TCP.
	// Example: `UDP`
	TransportProtocol DnsHealthCheckTransportProtocolsEnum `mandatory:"false" json:"transportProtocol,omitempty"`

	// The class the dns health check query to use; either IN or CH.
	// Example: `IN`
	QueryClass DnsHealthCheckQueryClassesEnum `mandatory:"false" json:"queryClass,omitempty"`

	// The type the dns health check query to use; A, AAAA, TXT.
	// Example: `A`
	QueryType DnsHealthCheckQueryTypesEnum `mandatory:"false" json:"queryType,omitempty"`

	// An array that represents accepetable RCODE values for DNS query response.
	// Example: ["NOERROR", "NXDOMAIN"]
	Rcodes []DnsHealthCheckRCodesEnum `mandatory:"false" json:"rcodes"`
}

func (m DnsHealthCheckerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DnsHealthCheckerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDnsHealthCheckTransportProtocolsEnum(string(m.TransportProtocol)); !ok && m.TransportProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportProtocol: %s. Supported values are: %s.", m.TransportProtocol, strings.Join(GetDnsHealthCheckTransportProtocolsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDnsHealthCheckQueryClassesEnum(string(m.QueryClass)); !ok && m.QueryClass != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryClass: %s. Supported values are: %s.", m.QueryClass, strings.Join(GetDnsHealthCheckQueryClassesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDnsHealthCheckQueryTypesEnum(string(m.QueryType)); !ok && m.QueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryType: %s. Supported values are: %s.", m.QueryType, strings.Join(GetDnsHealthCheckQueryTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
