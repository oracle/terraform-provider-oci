// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVanityDomainDetails Vanity domain request detail
type CreateVanityDomainDetails struct {

	// The origin request type for which the certificate is generated
	OriginCertRequestType CreateVanityDomainDetailsOriginCertRequestTypeEnum `mandatory:"false" json:"originCertRequestType,omitempty"`

	// The cdn request type for which the certificate is generated
	CdnCertRequestType CreateVanityDomainDetailsCdnCertRequestTypeEnum `mandatory:"false" json:"cdnCertRequestType,omitempty"`

	// Vanity domain
	VanityDomain *string `mandatory:"false" json:"vanityDomain"`

	// The dns is managed by the customer or Oracle
	DnsManagedBy CreateVanityDomainDetailsDnsManagedByEnum `mandatory:"false" json:"dnsManagedBy,omitempty"`

	// The prefix value of the DnsPrefix. Can't be changed after creation
	Prefix *string `mandatory:"false" json:"prefix"`

	CertificateInfo *CertificateInfo `mandatory:"false" json:"certificateInfo"`

	// The cm link that was used to create the DNS prefix
	ChangeManagementLink *string `mandatory:"false" json:"changeManagementLink"`
}

func (m CreateVanityDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVanityDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateVanityDomainDetailsOriginCertRequestTypeEnum(string(m.OriginCertRequestType)); !ok && m.OriginCertRequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginCertRequestType: %s. Supported values are: %s.", m.OriginCertRequestType, strings.Join(GetCreateVanityDomainDetailsOriginCertRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateVanityDomainDetailsCdnCertRequestTypeEnum(string(m.CdnCertRequestType)); !ok && m.CdnCertRequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CdnCertRequestType: %s. Supported values are: %s.", m.CdnCertRequestType, strings.Join(GetCreateVanityDomainDetailsCdnCertRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateVanityDomainDetailsDnsManagedByEnum(string(m.DnsManagedBy)); !ok && m.DnsManagedBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsManagedBy: %s. Supported values are: %s.", m.DnsManagedBy, strings.Join(GetCreateVanityDomainDetailsDnsManagedByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateVanityDomainDetailsOriginCertRequestTypeEnum Enum with underlying type: string
type CreateVanityDomainDetailsOriginCertRequestTypeEnum string

// Set of constants representing the allowable values for CreateVanityDomainDetailsOriginCertRequestTypeEnum
const (
	CreateVanityDomainDetailsOriginCertRequestTypeCsr CreateVanityDomainDetailsOriginCertRequestTypeEnum = "REQUEST_CSR"
	CreateVanityDomainDetailsOriginCertRequestTypeDv  CreateVanityDomainDetailsOriginCertRequestTypeEnum = "REQUEST_DV"
)

var mappingCreateVanityDomainDetailsOriginCertRequestTypeEnum = map[string]CreateVanityDomainDetailsOriginCertRequestTypeEnum{
	"REQUEST_CSR": CreateVanityDomainDetailsOriginCertRequestTypeCsr,
	"REQUEST_DV":  CreateVanityDomainDetailsOriginCertRequestTypeDv,
}

var mappingCreateVanityDomainDetailsOriginCertRequestTypeEnumLowerCase = map[string]CreateVanityDomainDetailsOriginCertRequestTypeEnum{
	"request_csr": CreateVanityDomainDetailsOriginCertRequestTypeCsr,
	"request_dv":  CreateVanityDomainDetailsOriginCertRequestTypeDv,
}

// GetCreateVanityDomainDetailsOriginCertRequestTypeEnumValues Enumerates the set of values for CreateVanityDomainDetailsOriginCertRequestTypeEnum
func GetCreateVanityDomainDetailsOriginCertRequestTypeEnumValues() []CreateVanityDomainDetailsOriginCertRequestTypeEnum {
	values := make([]CreateVanityDomainDetailsOriginCertRequestTypeEnum, 0)
	for _, v := range mappingCreateVanityDomainDetailsOriginCertRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVanityDomainDetailsOriginCertRequestTypeEnumStringValues Enumerates the set of values in String for CreateVanityDomainDetailsOriginCertRequestTypeEnum
func GetCreateVanityDomainDetailsOriginCertRequestTypeEnumStringValues() []string {
	return []string{
		"REQUEST_CSR",
		"REQUEST_DV",
	}
}

// GetMappingCreateVanityDomainDetailsOriginCertRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVanityDomainDetailsOriginCertRequestTypeEnum(val string) (CreateVanityDomainDetailsOriginCertRequestTypeEnum, bool) {
	enum, ok := mappingCreateVanityDomainDetailsOriginCertRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateVanityDomainDetailsCdnCertRequestTypeEnum Enum with underlying type: string
type CreateVanityDomainDetailsCdnCertRequestTypeEnum string

// Set of constants representing the allowable values for CreateVanityDomainDetailsCdnCertRequestTypeEnum
const (
	CreateVanityDomainDetailsCdnCertRequestTypeCsr CreateVanityDomainDetailsCdnCertRequestTypeEnum = "REQUEST_CSR"
	CreateVanityDomainDetailsCdnCertRequestTypeDv  CreateVanityDomainDetailsCdnCertRequestTypeEnum = "REQUEST_DV"
)

var mappingCreateVanityDomainDetailsCdnCertRequestTypeEnum = map[string]CreateVanityDomainDetailsCdnCertRequestTypeEnum{
	"REQUEST_CSR": CreateVanityDomainDetailsCdnCertRequestTypeCsr,
	"REQUEST_DV":  CreateVanityDomainDetailsCdnCertRequestTypeDv,
}

var mappingCreateVanityDomainDetailsCdnCertRequestTypeEnumLowerCase = map[string]CreateVanityDomainDetailsCdnCertRequestTypeEnum{
	"request_csr": CreateVanityDomainDetailsCdnCertRequestTypeCsr,
	"request_dv":  CreateVanityDomainDetailsCdnCertRequestTypeDv,
}

// GetCreateVanityDomainDetailsCdnCertRequestTypeEnumValues Enumerates the set of values for CreateVanityDomainDetailsCdnCertRequestTypeEnum
func GetCreateVanityDomainDetailsCdnCertRequestTypeEnumValues() []CreateVanityDomainDetailsCdnCertRequestTypeEnum {
	values := make([]CreateVanityDomainDetailsCdnCertRequestTypeEnum, 0)
	for _, v := range mappingCreateVanityDomainDetailsCdnCertRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVanityDomainDetailsCdnCertRequestTypeEnumStringValues Enumerates the set of values in String for CreateVanityDomainDetailsCdnCertRequestTypeEnum
func GetCreateVanityDomainDetailsCdnCertRequestTypeEnumStringValues() []string {
	return []string{
		"REQUEST_CSR",
		"REQUEST_DV",
	}
}

// GetMappingCreateVanityDomainDetailsCdnCertRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVanityDomainDetailsCdnCertRequestTypeEnum(val string) (CreateVanityDomainDetailsCdnCertRequestTypeEnum, bool) {
	enum, ok := mappingCreateVanityDomainDetailsCdnCertRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateVanityDomainDetailsDnsManagedByEnum Enum with underlying type: string
type CreateVanityDomainDetailsDnsManagedByEnum string

// Set of constants representing the allowable values for CreateVanityDomainDetailsDnsManagedByEnum
const (
	CreateVanityDomainDetailsDnsManagedByOracleManaged   CreateVanityDomainDetailsDnsManagedByEnum = "ORACLE_MANAGED"
	CreateVanityDomainDetailsDnsManagedByCustomerManaged CreateVanityDomainDetailsDnsManagedByEnum = "CUSTOMER_MANAGED"
)

var mappingCreateVanityDomainDetailsDnsManagedByEnum = map[string]CreateVanityDomainDetailsDnsManagedByEnum{
	"ORACLE_MANAGED":   CreateVanityDomainDetailsDnsManagedByOracleManaged,
	"CUSTOMER_MANAGED": CreateVanityDomainDetailsDnsManagedByCustomerManaged,
}

var mappingCreateVanityDomainDetailsDnsManagedByEnumLowerCase = map[string]CreateVanityDomainDetailsDnsManagedByEnum{
	"oracle_managed":   CreateVanityDomainDetailsDnsManagedByOracleManaged,
	"customer_managed": CreateVanityDomainDetailsDnsManagedByCustomerManaged,
}

// GetCreateVanityDomainDetailsDnsManagedByEnumValues Enumerates the set of values for CreateVanityDomainDetailsDnsManagedByEnum
func GetCreateVanityDomainDetailsDnsManagedByEnumValues() []CreateVanityDomainDetailsDnsManagedByEnum {
	values := make([]CreateVanityDomainDetailsDnsManagedByEnum, 0)
	for _, v := range mappingCreateVanityDomainDetailsDnsManagedByEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateVanityDomainDetailsDnsManagedByEnumStringValues Enumerates the set of values in String for CreateVanityDomainDetailsDnsManagedByEnum
func GetCreateVanityDomainDetailsDnsManagedByEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingCreateVanityDomainDetailsDnsManagedByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateVanityDomainDetailsDnsManagedByEnum(val string) (CreateVanityDomainDetailsDnsManagedByEnum, bool) {
	enum, ok := mappingCreateVanityDomainDetailsDnsManagedByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
