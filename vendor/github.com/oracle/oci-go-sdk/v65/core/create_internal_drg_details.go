// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateInternalDrgDetails The request to create a new InternalDrg.
type CreateInternalDrgDetails struct {

	// The DRG's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the DRG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Anycast IP of the El Paso fleet handling the ingress traffic.
	IngressIP *string `mandatory:"true" json:"ingressIP"`

	// Anycast IP of the El Paso fleet handling the egress traffic.
	EgressIP *string `mandatory:"true" json:"egressIP"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// The DP ID of the DRG
	DrgDpId *int `mandatory:"false" json:"drgDpId"`

	// Route data for the Drg.
	RouteData *string `mandatory:"false" json:"routeData"`

	// Indicates if Drg is Global or Regional
	IsGlobal *bool `mandatory:"false" json:"isGlobal"`

	// Indicates if Drg is Substrate Access or not
	IsSubstrateAccess *bool `mandatory:"false" json:"isSubstrateAccess"`

	// The type of the Substrate Access DRG. Can only be specified when isSubstrateAccess = TRUE.
	SubstrateAccessDrgType CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum `mandatory:"false" json:"substrateAccessDrgType,omitempty"`
}

func (m CreateInternalDrgDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalDrgDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnum(string(m.SubstrateAccessDrgType)); !ok && m.SubstrateAccessDrgType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubstrateAccessDrgType: %s. Supported values are: %s.", m.SubstrateAccessDrgType, strings.Join(GetCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum Enum with underlying type: string
type CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum string

// Set of constants representing the allowable values for CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum
const (
	CreateInternalDrgDetailsSubstrateAccessDrgTypeProd CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum = "SUBSTRATE_ACCESS_DRG_PROD"
	CreateInternalDrgDetailsSubstrateAccessDrgTypeTest CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum = "SUBSTRATE_ACCESS_DRG_TEST"
)

var mappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnum = map[string]CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum{
	"SUBSTRATE_ACCESS_DRG_PROD": CreateInternalDrgDetailsSubstrateAccessDrgTypeProd,
	"SUBSTRATE_ACCESS_DRG_TEST": CreateInternalDrgDetailsSubstrateAccessDrgTypeTest,
}

var mappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumLowerCase = map[string]CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum{
	"substrate_access_drg_prod": CreateInternalDrgDetailsSubstrateAccessDrgTypeProd,
	"substrate_access_drg_test": CreateInternalDrgDetailsSubstrateAccessDrgTypeTest,
}

// GetCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumValues Enumerates the set of values for CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum
func GetCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumValues() []CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum {
	values := make([]CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum, 0)
	for _, v := range mappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumStringValues Enumerates the set of values in String for CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum
func GetCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumStringValues() []string {
	return []string{
		"SUBSTRATE_ACCESS_DRG_PROD",
		"SUBSTRATE_ACCESS_DRG_TEST",
	}
}

// GetMappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnum(val string) (CreateInternalDrgDetailsSubstrateAccessDrgTypeEnum, bool) {
	enum, ok := mappingCreateInternalDrgDetailsSubstrateAccessDrgTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
