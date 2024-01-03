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

// DnsTraceMonitorConfiguration Request configuration details for the DNS Trace monitor type.
type DnsTraceMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// Verify response content against regular expression based string.
	// If response content does not match the verifyResponseContent value, then it will be considered a failure.
	VerifyResponseContent *string `mandatory:"false" json:"verifyResponseContent"`

	// DNS record type.
	RecordType DnsRecordTypeEnum `mandatory:"false" json:"recordType,omitempty"`

	// Type of protocol.
	Protocol DnsTransportProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m DnsTraceMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m DnsTraceMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m DnsTraceMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DnsTraceMonitorConfiguration) ValidateEnumValue() (bool, error) {
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
func (m DnsTraceMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDnsTraceMonitorConfiguration DnsTraceMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeDnsTraceMonitorConfiguration
	}{
		"DNS_TRACE_CONFIG",
		(MarshalTypeDnsTraceMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
