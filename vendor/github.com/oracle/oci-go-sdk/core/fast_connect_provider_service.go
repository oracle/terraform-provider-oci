// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// FastConnectProviderService A service offering from a supported provider. For more information,
// see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
type FastConnectProviderService struct {

	// The OCID of the service offered by the provider.
	Id *string `mandatory:"true" json:"id"`

	// Private peering BGP management.
	PrivatePeeringBgpManagement FastConnectProviderServicePrivatePeeringBgpManagementEnum `mandatory:"true" json:"privatePeeringBgpManagement"`

	// The name of the provider.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// The name of the service offered by the provider.
	ProviderServiceName *string `mandatory:"true" json:"providerServiceName"`

	// Public peering BGP management.
	PublicPeeringBgpManagement FastConnectProviderServicePublicPeeringBgpManagementEnum `mandatory:"true" json:"publicPeeringBgpManagement"`

	// Provider service type.
	Type FastConnectProviderServiceTypeEnum `mandatory:"true" json:"type"`

	// A description of the service offered by the provider.
	Description *string `mandatory:"false" json:"description"`

	// An array of virtual circuit types supported by this service.
	SupportedVirtualCircuitTypes []FastConnectProviderServiceSupportedVirtualCircuitTypesEnum `mandatory:"false" json:"supportedVirtualCircuitTypes"`
}

func (m FastConnectProviderService) String() string {
	return common.PointerString(m)
}

// FastConnectProviderServicePrivatePeeringBgpManagementEnum Enum with underlying type: string
type FastConnectProviderServicePrivatePeeringBgpManagementEnum string

// Set of constants representing the allowable values for FastConnectProviderServicePrivatePeeringBgpManagement
const (
	FastConnectProviderServicePrivatePeeringBgpManagementCustomerManaged FastConnectProviderServicePrivatePeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementProviderManaged FastConnectProviderServicePrivatePeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementOracleManaged   FastConnectProviderServicePrivatePeeringBgpManagementEnum = "ORACLE_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementUnknown         FastConnectProviderServicePrivatePeeringBgpManagementEnum = "UNKNOWN"
)

var mappingFastConnectProviderServicePrivatePeeringBgpManagement = map[string]FastConnectProviderServicePrivatePeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FastConnectProviderServicePrivatePeeringBgpManagementCustomerManaged,
	"PROVIDER_MANAGED": FastConnectProviderServicePrivatePeeringBgpManagementProviderManaged,
	"ORACLE_MANAGED":   FastConnectProviderServicePrivatePeeringBgpManagementOracleManaged,
	"UNKNOWN":          FastConnectProviderServicePrivatePeeringBgpManagementUnknown,
}

// GetFastConnectProviderServicePrivatePeeringBgpManagementEnumValues Enumerates the set of values for FastConnectProviderServicePrivatePeeringBgpManagement
func GetFastConnectProviderServicePrivatePeeringBgpManagementEnumValues() []FastConnectProviderServicePrivatePeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePrivatePeeringBgpManagementEnum, 0)
	for _, v := range mappingFastConnectProviderServicePrivatePeeringBgpManagement {
		if v != FastConnectProviderServicePrivatePeeringBgpManagementUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServicePublicPeeringBgpManagementEnum Enum with underlying type: string
type FastConnectProviderServicePublicPeeringBgpManagementEnum string

// Set of constants representing the allowable values for FastConnectProviderServicePublicPeeringBgpManagement
const (
	FastConnectProviderServicePublicPeeringBgpManagementCustomerManaged FastConnectProviderServicePublicPeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementProviderManaged FastConnectProviderServicePublicPeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementOracleManaged   FastConnectProviderServicePublicPeeringBgpManagementEnum = "ORACLE_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementUnknown         FastConnectProviderServicePublicPeeringBgpManagementEnum = "UNKNOWN"
)

var mappingFastConnectProviderServicePublicPeeringBgpManagement = map[string]FastConnectProviderServicePublicPeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FastConnectProviderServicePublicPeeringBgpManagementCustomerManaged,
	"PROVIDER_MANAGED": FastConnectProviderServicePublicPeeringBgpManagementProviderManaged,
	"ORACLE_MANAGED":   FastConnectProviderServicePublicPeeringBgpManagementOracleManaged,
	"UNKNOWN":          FastConnectProviderServicePublicPeeringBgpManagementUnknown,
}

// GetFastConnectProviderServicePublicPeeringBgpManagementEnumValues Enumerates the set of values for FastConnectProviderServicePublicPeeringBgpManagement
func GetFastConnectProviderServicePublicPeeringBgpManagementEnumValues() []FastConnectProviderServicePublicPeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePublicPeeringBgpManagementEnum, 0)
	for _, v := range mappingFastConnectProviderServicePublicPeeringBgpManagement {
		if v != FastConnectProviderServicePublicPeeringBgpManagementUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServiceSupportedVirtualCircuitTypesEnum Enum with underlying type: string
type FastConnectProviderServiceSupportedVirtualCircuitTypesEnum string

// Set of constants representing the allowable values for FastConnectProviderServiceSupportedVirtualCircuitTypes
const (
	FastConnectProviderServiceSupportedVirtualCircuitTypesPublic  FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "PUBLIC"
	FastConnectProviderServiceSupportedVirtualCircuitTypesPrivate FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "PRIVATE"
	FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "UNKNOWN"
)

var mappingFastConnectProviderServiceSupportedVirtualCircuitTypes = map[string]FastConnectProviderServiceSupportedVirtualCircuitTypesEnum{
	"PUBLIC":  FastConnectProviderServiceSupportedVirtualCircuitTypesPublic,
	"PRIVATE": FastConnectProviderServiceSupportedVirtualCircuitTypesPrivate,
	"UNKNOWN": FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown,
}

// GetFastConnectProviderServiceSupportedVirtualCircuitTypesEnumValues Enumerates the set of values for FastConnectProviderServiceSupportedVirtualCircuitTypes
func GetFastConnectProviderServiceSupportedVirtualCircuitTypesEnumValues() []FastConnectProviderServiceSupportedVirtualCircuitTypesEnum {
	values := make([]FastConnectProviderServiceSupportedVirtualCircuitTypesEnum, 0)
	for _, v := range mappingFastConnectProviderServiceSupportedVirtualCircuitTypes {
		if v != FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServiceTypeEnum Enum with underlying type: string
type FastConnectProviderServiceTypeEnum string

// Set of constants representing the allowable values for FastConnectProviderServiceType
const (
	FastConnectProviderServiceTypeLayer2  FastConnectProviderServiceTypeEnum = "LAYER2"
	FastConnectProviderServiceTypeLayer3  FastConnectProviderServiceTypeEnum = "LAYER3"
	FastConnectProviderServiceTypeUnknown FastConnectProviderServiceTypeEnum = "UNKNOWN"
)

var mappingFastConnectProviderServiceType = map[string]FastConnectProviderServiceTypeEnum{
	"LAYER2":  FastConnectProviderServiceTypeLayer2,
	"LAYER3":  FastConnectProviderServiceTypeLayer3,
	"UNKNOWN": FastConnectProviderServiceTypeUnknown,
}

// GetFastConnectProviderServiceTypeEnumValues Enumerates the set of values for FastConnectProviderServiceType
func GetFastConnectProviderServiceTypeEnumValues() []FastConnectProviderServiceTypeEnum {
	values := make([]FastConnectProviderServiceTypeEnum, 0)
	for _, v := range mappingFastConnectProviderServiceType {
		if v != FastConnectProviderServiceTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}
