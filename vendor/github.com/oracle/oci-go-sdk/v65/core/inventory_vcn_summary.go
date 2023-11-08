// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InventoryVcnSummary A Summary of IP Inventory data of the VCN under given compartments
type InventoryVcnSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN .
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The Name of  VCN
	VcnName *string `mandatory:"false" json:"vcnName"`

	// Resource types of the VCN
	ResourceType InventoryVcnSummaryResourceTypeEnum `mandatory:"false" json:"resourceType,omitempty"`

	// A list of `InventoryVcnCidrBlockSummary` objects.
	InventoryVcnCidrBlockCollection []InventoryVcnCidrBlockSummary `mandatory:"false" json:"inventoryVcnCidrBlockCollection"`

	// The DNS Domain Name of the VCN
	DnsDomainName *string `mandatory:"false" json:"dnsDomainName"`

	// The Region Name of the VCN
	Region *string `mandatory:"false" json:"region"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A List of `Subnetcollection` objects
	InventorySubnetcollection []InventorySubnetSummary `mandatory:"false" json:"inventorySubnetcollection"`
}

func (m InventoryVcnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InventoryVcnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInventoryVcnSummaryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetInventoryVcnSummaryResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InventoryVcnSummaryResourceTypeEnum Enum with underlying type: string
type InventoryVcnSummaryResourceTypeEnum string

// Set of constants representing the allowable values for InventoryVcnSummaryResourceTypeEnum
const (
	InventoryVcnSummaryResourceTypeVcn InventoryVcnSummaryResourceTypeEnum = "VCN"
)

var mappingInventoryVcnSummaryResourceTypeEnum = map[string]InventoryVcnSummaryResourceTypeEnum{
	"VCN": InventoryVcnSummaryResourceTypeVcn,
}

var mappingInventoryVcnSummaryResourceTypeEnumLowerCase = map[string]InventoryVcnSummaryResourceTypeEnum{
	"vcn": InventoryVcnSummaryResourceTypeVcn,
}

// GetInventoryVcnSummaryResourceTypeEnumValues Enumerates the set of values for InventoryVcnSummaryResourceTypeEnum
func GetInventoryVcnSummaryResourceTypeEnumValues() []InventoryVcnSummaryResourceTypeEnum {
	values := make([]InventoryVcnSummaryResourceTypeEnum, 0)
	for _, v := range mappingInventoryVcnSummaryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInventoryVcnSummaryResourceTypeEnumStringValues Enumerates the set of values in String for InventoryVcnSummaryResourceTypeEnum
func GetInventoryVcnSummaryResourceTypeEnumStringValues() []string {
	return []string{
		"VCN",
	}
}

// GetMappingInventoryVcnSummaryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInventoryVcnSummaryResourceTypeEnum(val string) (InventoryVcnSummaryResourceTypeEnum, bool) {
	enum, ok := mappingInventoryVcnSummaryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
