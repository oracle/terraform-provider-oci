// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateInternalDnsRecordDetails This structure is used when creating DnsRecord for internal clients.
type CreateInternalDnsRecordDetails struct {

	// The OCID of the compartment to contain the DnsRecord.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Internal Hosted Zone the DnsRecord belongs to.
	InternalHostedZoneId *string `mandatory:"true" json:"internalHostedZoneId"`

	// Name of the DnsRecord.
	// -*A:* Partially Qualified DNS Name that will be mapped to the IPv4 address
	Name *string `mandatory:"true" json:"name"`

	// Type of Dns Record according to RFC 1035 (https://tools.ietf.org/html/rfc1035).
	// Currently supported list of types are the following.
	// -*A:* Type 1, a hostname to IPv4 address
	Type CreateInternalDnsRecordDetailsTypeEnum `mandatory:"true" json:"type"`

	// Value for the DnsRecord.
	// -*A:* One or more IPv4 addresses. Comma separated.
	Value *string `mandatory:"true" json:"value"`

	// Time to live value in seconds for the DnsRecord, according to RFC 1035 (https://tools.ietf.org/html/rfc1035).
	// Defaults to 86400.
	Ttl *int `mandatory:"false" json:"ttl"`
}

func (m CreateInternalDnsRecordDetails) String() string {
	return common.PointerString(m)
}

// CreateInternalDnsRecordDetailsTypeEnum Enum with underlying type: string
type CreateInternalDnsRecordDetailsTypeEnum string

// Set of constants representing the allowable values for CreateInternalDnsRecordDetailsTypeEnum
const (
	CreateInternalDnsRecordDetailsTypeA CreateInternalDnsRecordDetailsTypeEnum = "A"
)

var mappingCreateInternalDnsRecordDetailsType = map[string]CreateInternalDnsRecordDetailsTypeEnum{
	"A": CreateInternalDnsRecordDetailsTypeA,
}

// GetCreateInternalDnsRecordDetailsTypeEnumValues Enumerates the set of values for CreateInternalDnsRecordDetailsTypeEnum
func GetCreateInternalDnsRecordDetailsTypeEnumValues() []CreateInternalDnsRecordDetailsTypeEnum {
	values := make([]CreateInternalDnsRecordDetailsTypeEnum, 0)
	for _, v := range mappingCreateInternalDnsRecordDetailsType {
		values = append(values, v)
	}
	return values
}
