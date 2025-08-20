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

// DscpOverride DSCP override used to map DSCP values to respective OCI mapped values for queues: `PREMIUM` (P1), `DEFAULT` (P2), `BULK` (P3), or `SCAVENGER` (P4).
type DscpOverride struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DSCP override.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the DSCP override.
	Id *string `mandatory:"true" json:"id"`

	// The Tenancy's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for which the DSCP override is applicable.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The current state of the DSCP override.
	LifecycleState DscpOverrideLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of `DscpOverrides` which consist of DSCP values and a respective `ClassOfService`. Example: `{43 - PREMIUM}`
	DscpOverrides []DscpOverrides `mandatory:"true" json:"dscpOverrides"`

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

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the DSCP override was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m DscpOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DscpOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDscpOverrideLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDscpOverrideLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DscpOverrideLifecycleStateEnum Enum with underlying type: string
type DscpOverrideLifecycleStateEnum string

// Set of constants representing the allowable values for DscpOverrideLifecycleStateEnum
const (
	DscpOverrideLifecycleStateProvisioning DscpOverrideLifecycleStateEnum = "PROVISIONING"
	DscpOverrideLifecycleStateAvailable    DscpOverrideLifecycleStateEnum = "AVAILABLE"
	DscpOverrideLifecycleStateTerminating  DscpOverrideLifecycleStateEnum = "TERMINATING"
	DscpOverrideLifecycleStateTerminated   DscpOverrideLifecycleStateEnum = "TERMINATED"
	DscpOverrideLifecycleStateUpdating     DscpOverrideLifecycleStateEnum = "UPDATING"
)

var mappingDscpOverrideLifecycleStateEnum = map[string]DscpOverrideLifecycleStateEnum{
	"PROVISIONING": DscpOverrideLifecycleStateProvisioning,
	"AVAILABLE":    DscpOverrideLifecycleStateAvailable,
	"TERMINATING":  DscpOverrideLifecycleStateTerminating,
	"TERMINATED":   DscpOverrideLifecycleStateTerminated,
	"UPDATING":     DscpOverrideLifecycleStateUpdating,
}

var mappingDscpOverrideLifecycleStateEnumLowerCase = map[string]DscpOverrideLifecycleStateEnum{
	"provisioning": DscpOverrideLifecycleStateProvisioning,
	"available":    DscpOverrideLifecycleStateAvailable,
	"terminating":  DscpOverrideLifecycleStateTerminating,
	"terminated":   DscpOverrideLifecycleStateTerminated,
	"updating":     DscpOverrideLifecycleStateUpdating,
}

// GetDscpOverrideLifecycleStateEnumValues Enumerates the set of values for DscpOverrideLifecycleStateEnum
func GetDscpOverrideLifecycleStateEnumValues() []DscpOverrideLifecycleStateEnum {
	values := make([]DscpOverrideLifecycleStateEnum, 0)
	for _, v := range mappingDscpOverrideLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDscpOverrideLifecycleStateEnumStringValues Enumerates the set of values in String for DscpOverrideLifecycleStateEnum
func GetDscpOverrideLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingDscpOverrideLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDscpOverrideLifecycleStateEnum(val string) (DscpOverrideLifecycleStateEnum, bool) {
	enum, ok := mappingDscpOverrideLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
