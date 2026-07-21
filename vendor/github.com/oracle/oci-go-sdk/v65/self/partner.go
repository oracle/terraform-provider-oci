// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Partner Marketplace publisher partner details.
type Partner struct {

	// The unique identifier of the marketplace publisher partner.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier of the compartment of partner.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the marketplace publisher partner.
	LifecycleState PartnerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The display name of marketplace publisher partner.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Partner) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Partner) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPartnerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPartnerLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PartnerLifecycleStateEnum Enum with underlying type: string
type PartnerLifecycleStateEnum string

// Set of constants representing the allowable values for PartnerLifecycleStateEnum
const (
	PartnerLifecycleStateActive  PartnerLifecycleStateEnum = "ACTIVE"
	PartnerLifecycleStateDeleted PartnerLifecycleStateEnum = "DELETED"
)

var mappingPartnerLifecycleStateEnum = map[string]PartnerLifecycleStateEnum{
	"ACTIVE":  PartnerLifecycleStateActive,
	"DELETED": PartnerLifecycleStateDeleted,
}

var mappingPartnerLifecycleStateEnumLowerCase = map[string]PartnerLifecycleStateEnum{
	"active":  PartnerLifecycleStateActive,
	"deleted": PartnerLifecycleStateDeleted,
}

// GetPartnerLifecycleStateEnumValues Enumerates the set of values for PartnerLifecycleStateEnum
func GetPartnerLifecycleStateEnumValues() []PartnerLifecycleStateEnum {
	values := make([]PartnerLifecycleStateEnum, 0)
	for _, v := range mappingPartnerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPartnerLifecycleStateEnumStringValues Enumerates the set of values in String for PartnerLifecycleStateEnum
func GetPartnerLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingPartnerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPartnerLifecycleStateEnum(val string) (PartnerLifecycleStateEnum, bool) {
	enum, ok := mappingPartnerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
