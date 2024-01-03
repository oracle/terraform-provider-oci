// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DnsRecordTypeEnum Enum with underlying type: string
type DnsRecordTypeEnum string

// Set of constants representing the allowable values for DnsRecordTypeEnum
const (
	DnsRecordTypeA       DnsRecordTypeEnum = "A"
	DnsRecordTypeAaaa    DnsRecordTypeEnum = "AAAA"
	DnsRecordTypeAny     DnsRecordTypeEnum = "ANY"
	DnsRecordTypeCname   DnsRecordTypeEnum = "CNAME"
	DnsRecordTypeDnskey  DnsRecordTypeEnum = "DNSKEY"
	DnsRecordTypeDs      DnsRecordTypeEnum = "DS"
	DnsRecordTypeMx      DnsRecordTypeEnum = "MX"
	DnsRecordTypeNs      DnsRecordTypeEnum = "NS"
	DnsRecordTypeNsec    DnsRecordTypeEnum = "NSEC"
	DnsRecordTypeNullRec DnsRecordTypeEnum = "NULL_REC"
	DnsRecordTypePtr     DnsRecordTypeEnum = "PTR"
	DnsRecordTypeRrsig   DnsRecordTypeEnum = "RRSIG"
	DnsRecordTypeSoa     DnsRecordTypeEnum = "SOA"
	DnsRecordTypeTxt     DnsRecordTypeEnum = "TXT"
)

var mappingDnsRecordTypeEnum = map[string]DnsRecordTypeEnum{
	"A":        DnsRecordTypeA,
	"AAAA":     DnsRecordTypeAaaa,
	"ANY":      DnsRecordTypeAny,
	"CNAME":    DnsRecordTypeCname,
	"DNSKEY":   DnsRecordTypeDnskey,
	"DS":       DnsRecordTypeDs,
	"MX":       DnsRecordTypeMx,
	"NS":       DnsRecordTypeNs,
	"NSEC":     DnsRecordTypeNsec,
	"NULL_REC": DnsRecordTypeNullRec,
	"PTR":      DnsRecordTypePtr,
	"RRSIG":    DnsRecordTypeRrsig,
	"SOA":      DnsRecordTypeSoa,
	"TXT":      DnsRecordTypeTxt,
}

var mappingDnsRecordTypeEnumLowerCase = map[string]DnsRecordTypeEnum{
	"a":        DnsRecordTypeA,
	"aaaa":     DnsRecordTypeAaaa,
	"any":      DnsRecordTypeAny,
	"cname":    DnsRecordTypeCname,
	"dnskey":   DnsRecordTypeDnskey,
	"ds":       DnsRecordTypeDs,
	"mx":       DnsRecordTypeMx,
	"ns":       DnsRecordTypeNs,
	"nsec":     DnsRecordTypeNsec,
	"null_rec": DnsRecordTypeNullRec,
	"ptr":      DnsRecordTypePtr,
	"rrsig":    DnsRecordTypeRrsig,
	"soa":      DnsRecordTypeSoa,
	"txt":      DnsRecordTypeTxt,
}

// GetDnsRecordTypeEnumValues Enumerates the set of values for DnsRecordTypeEnum
func GetDnsRecordTypeEnumValues() []DnsRecordTypeEnum {
	values := make([]DnsRecordTypeEnum, 0)
	for _, v := range mappingDnsRecordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDnsRecordTypeEnumStringValues Enumerates the set of values in String for DnsRecordTypeEnum
func GetDnsRecordTypeEnumStringValues() []string {
	return []string{
		"A",
		"AAAA",
		"ANY",
		"CNAME",
		"DNSKEY",
		"DS",
		"MX",
		"NS",
		"NSEC",
		"NULL_REC",
		"PTR",
		"RRSIG",
		"SOA",
		"TXT",
	}
}

// GetMappingDnsRecordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnsRecordTypeEnum(val string) (DnsRecordTypeEnum, bool) {
	enum, ok := mappingDnsRecordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
