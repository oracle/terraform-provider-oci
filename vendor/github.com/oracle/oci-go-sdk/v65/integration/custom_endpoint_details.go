// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomEndpointDetails Details for a custom endpoint for the integration instance.
type CustomEndpointDetails struct {

	// A custom hostname to be used for the integration instance URL, in FQDN format.
	Hostname *string `mandatory:"true" json:"hostname"`

	// Indicates if custom endpoint is managed by oracle or customer.
	ManagedType CustomEndpointDetailsManagedTypeEnum `mandatory:"false" json:"managedType,omitempty"`

	// DNS Zone name
	DnsZoneName *string `mandatory:"false" json:"dnsZoneName"`

	// Type of DNS.
	DnsType CustomEndpointDetailsDnsTypeEnum `mandatory:"false" json:"dnsType,omitempty"`

	// Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname.
	CertificateSecretId *string `mandatory:"false" json:"certificateSecretId"`

	// The secret version used for the certificate-secret-id (if certificate-secret-id is specified).
	CertificateSecretVersion *int `mandatory:"false" json:"certificateSecretVersion"`

	// When creating the DNS CNAME record for the custom hostname, this value must be specified in the rdata.
	Alias *string `mandatory:"false" json:"alias"`
}

func (m CustomEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomEndpointDetailsManagedTypeEnum(string(m.ManagedType)); !ok && m.ManagedType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedType: %s. Supported values are: %s.", m.ManagedType, strings.Join(GetCustomEndpointDetailsManagedTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomEndpointDetailsDnsTypeEnum(string(m.DnsType)); !ok && m.DnsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsType: %s. Supported values are: %s.", m.DnsType, strings.Join(GetCustomEndpointDetailsDnsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomEndpointDetailsManagedTypeEnum Enum with underlying type: string
type CustomEndpointDetailsManagedTypeEnum string

// Set of constants representing the allowable values for CustomEndpointDetailsManagedTypeEnum
const (
	CustomEndpointDetailsManagedTypeOracleManaged   CustomEndpointDetailsManagedTypeEnum = "ORACLE_MANAGED"
	CustomEndpointDetailsManagedTypeCustomerManaged CustomEndpointDetailsManagedTypeEnum = "CUSTOMER_MANAGED"
)

var mappingCustomEndpointDetailsManagedTypeEnum = map[string]CustomEndpointDetailsManagedTypeEnum{
	"ORACLE_MANAGED":   CustomEndpointDetailsManagedTypeOracleManaged,
	"CUSTOMER_MANAGED": CustomEndpointDetailsManagedTypeCustomerManaged,
}

var mappingCustomEndpointDetailsManagedTypeEnumLowerCase = map[string]CustomEndpointDetailsManagedTypeEnum{
	"oracle_managed":   CustomEndpointDetailsManagedTypeOracleManaged,
	"customer_managed": CustomEndpointDetailsManagedTypeCustomerManaged,
}

// GetCustomEndpointDetailsManagedTypeEnumValues Enumerates the set of values for CustomEndpointDetailsManagedTypeEnum
func GetCustomEndpointDetailsManagedTypeEnumValues() []CustomEndpointDetailsManagedTypeEnum {
	values := make([]CustomEndpointDetailsManagedTypeEnum, 0)
	for _, v := range mappingCustomEndpointDetailsManagedTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomEndpointDetailsManagedTypeEnumStringValues Enumerates the set of values in String for CustomEndpointDetailsManagedTypeEnum
func GetCustomEndpointDetailsManagedTypeEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingCustomEndpointDetailsManagedTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomEndpointDetailsManagedTypeEnum(val string) (CustomEndpointDetailsManagedTypeEnum, bool) {
	enum, ok := mappingCustomEndpointDetailsManagedTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CustomEndpointDetailsDnsTypeEnum Enum with underlying type: string
type CustomEndpointDetailsDnsTypeEnum string

// Set of constants representing the allowable values for CustomEndpointDetailsDnsTypeEnum
const (
	CustomEndpointDetailsDnsTypeOci CustomEndpointDetailsDnsTypeEnum = "OCI"
)

var mappingCustomEndpointDetailsDnsTypeEnum = map[string]CustomEndpointDetailsDnsTypeEnum{
	"OCI": CustomEndpointDetailsDnsTypeOci,
}

var mappingCustomEndpointDetailsDnsTypeEnumLowerCase = map[string]CustomEndpointDetailsDnsTypeEnum{
	"oci": CustomEndpointDetailsDnsTypeOci,
}

// GetCustomEndpointDetailsDnsTypeEnumValues Enumerates the set of values for CustomEndpointDetailsDnsTypeEnum
func GetCustomEndpointDetailsDnsTypeEnumValues() []CustomEndpointDetailsDnsTypeEnum {
	values := make([]CustomEndpointDetailsDnsTypeEnum, 0)
	for _, v := range mappingCustomEndpointDetailsDnsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomEndpointDetailsDnsTypeEnumStringValues Enumerates the set of values in String for CustomEndpointDetailsDnsTypeEnum
func GetCustomEndpointDetailsDnsTypeEnumStringValues() []string {
	return []string{
		"OCI",
	}
}

// GetMappingCustomEndpointDetailsDnsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomEndpointDetailsDnsTypeEnum(val string) (CustomEndpointDetailsDnsTypeEnum, bool) {
	enum, ok := mappingCustomEndpointDetailsDnsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
