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

// DscpOverrides List of DSCP values for each ClassOfService.
// Example: `{43 - PREMIUM}`, `{20 - DEFAULT}`, `{15 - BULK}`, `{25 - SCAVENGER}`
type DscpOverrides struct {

	// DSCP values. DSCP uses 6 bits in the IP packet header, thereby giving 2^6 = 64 possible values (0 to 63). See RFC 4594 (https://datatracker.ietf.org/doc/html/rfc4594).
	// type: integer
	DscpValue *int `mandatory:"true" json:"dscpValue"`

	// The type of Class Of Service or DSCP queue for each DSCP value. `PREMIUM` (P1), `DEFAULT` (P2), `BULK` (P3), `SCAVENGER` (P4)
	ClassOfService DscpOverridesClassOfServiceEnum `mandatory:"true" json:"classOfService"`
}

func (m DscpOverrides) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DscpOverrides) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDscpOverridesClassOfServiceEnum(string(m.ClassOfService)); !ok && m.ClassOfService != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassOfService: %s. Supported values are: %s.", m.ClassOfService, strings.Join(GetDscpOverridesClassOfServiceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DscpOverridesClassOfServiceEnum Enum with underlying type: string
type DscpOverridesClassOfServiceEnum string

// Set of constants representing the allowable values for DscpOverridesClassOfServiceEnum
const (
	DscpOverridesClassOfServicePremium   DscpOverridesClassOfServiceEnum = "PREMIUM"
	DscpOverridesClassOfServiceDefault   DscpOverridesClassOfServiceEnum = "DEFAULT"
	DscpOverridesClassOfServiceBulk      DscpOverridesClassOfServiceEnum = "BULK"
	DscpOverridesClassOfServiceScavenger DscpOverridesClassOfServiceEnum = "SCAVENGER"
)

var mappingDscpOverridesClassOfServiceEnum = map[string]DscpOverridesClassOfServiceEnum{
	"PREMIUM":   DscpOverridesClassOfServicePremium,
	"DEFAULT":   DscpOverridesClassOfServiceDefault,
	"BULK":      DscpOverridesClassOfServiceBulk,
	"SCAVENGER": DscpOverridesClassOfServiceScavenger,
}

var mappingDscpOverridesClassOfServiceEnumLowerCase = map[string]DscpOverridesClassOfServiceEnum{
	"premium":   DscpOverridesClassOfServicePremium,
	"default":   DscpOverridesClassOfServiceDefault,
	"bulk":      DscpOverridesClassOfServiceBulk,
	"scavenger": DscpOverridesClassOfServiceScavenger,
}

// GetDscpOverridesClassOfServiceEnumValues Enumerates the set of values for DscpOverridesClassOfServiceEnum
func GetDscpOverridesClassOfServiceEnumValues() []DscpOverridesClassOfServiceEnum {
	values := make([]DscpOverridesClassOfServiceEnum, 0)
	for _, v := range mappingDscpOverridesClassOfServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetDscpOverridesClassOfServiceEnumStringValues Enumerates the set of values in String for DscpOverridesClassOfServiceEnum
func GetDscpOverridesClassOfServiceEnumStringValues() []string {
	return []string{
		"PREMIUM",
		"DEFAULT",
		"BULK",
		"SCAVENGER",
	}
}

// GetMappingDscpOverridesClassOfServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDscpOverridesClassOfServiceEnum(val string) (DscpOverridesClassOfServiceEnum, bool) {
	enum, ok := mappingDscpOverridesClassOfServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
