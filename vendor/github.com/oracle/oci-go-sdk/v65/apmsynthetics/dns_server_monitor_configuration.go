// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DnsServerMonitorConfiguration Request configuration details for the DNS Server monitor type.
type DnsServerMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// Name of the server that will be used to perform DNS lookup.
	NameServer *string `mandatory:"false" json:"nameServer"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// Verify response content against regular expression based string.
	// If response content does not match the verifyResponseContent value, then it will be considered a failure.
	VerifyResponseContent *string `mandatory:"false" json:"verifyResponseContent"`

	// If isQueryRecursive is enabled, then queries will be sent recursively to the target server.
	IsQueryRecursive *bool `mandatory:"false" json:"isQueryRecursive"`

	// DNS record type.
	RecordType DnsRecordTypeEnum `mandatory:"false" json:"recordType,omitempty"`

	// Type of protocol.
	Protocol DnsTransportProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m DnsServerMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m DnsServerMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m DnsServerMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DnsServerMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDnsRecordTypeEnum(string(m.RecordType)); !ok && m.RecordType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecordType: %s. Supported values are: %s.", m.RecordType, strings.Join(GetDnsRecordTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDnsTransportProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetDnsTransportProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DnsServerMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDnsServerMonitorConfiguration DnsServerMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeDnsServerMonitorConfiguration
	}{
		"DNS_SERVER_CONFIG",
		(MarshalTypeDnsServerMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
