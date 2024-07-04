// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// QosMappings List of QosMappings consisting of DSCP values with their respective ClassOfService. Eg {43 - PREMIUM}
type QosMappings struct {

	// Differentiated Services Code Point(DSCP) Values for QoS. DSCP uses the 6 bits, thereby giving 2^6 = 64 different values (0 to 63)
	DscpValues []int `mandatory:"true" json:"dscpValues"`

	// The type of Class Of Service for each DSCP values. PREMIUM (P1), DEFAULT (P2), BULK (P3), SCAVENGER (P4)
	ClassOfService QosMappingsClassOfServiceEnum `mandatory:"true" json:"classOfService"`
}

func (m QosMappings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QosMappings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQosMappingsClassOfServiceEnum(string(m.ClassOfService)); !ok && m.ClassOfService != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassOfService: %s. Supported values are: %s.", m.ClassOfService, strings.Join(GetQosMappingsClassOfServiceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QosMappingsClassOfServiceEnum Enum with underlying type: string
type QosMappingsClassOfServiceEnum string

// Set of constants representing the allowable values for QosMappingsClassOfServiceEnum
const (
	QosMappingsClassOfServicePremium   QosMappingsClassOfServiceEnum = "PREMIUM"
	QosMappingsClassOfServiceDefault   QosMappingsClassOfServiceEnum = "DEFAULT"
	QosMappingsClassOfServiceBulk      QosMappingsClassOfServiceEnum = "BULK"
	QosMappingsClassOfServiceScavenger QosMappingsClassOfServiceEnum = "SCAVENGER"
)

var mappingQosMappingsClassOfServiceEnum = map[string]QosMappingsClassOfServiceEnum{
	"PREMIUM":   QosMappingsClassOfServicePremium,
	"DEFAULT":   QosMappingsClassOfServiceDefault,
	"BULK":      QosMappingsClassOfServiceBulk,
	"SCAVENGER": QosMappingsClassOfServiceScavenger,
}

var mappingQosMappingsClassOfServiceEnumLowerCase = map[string]QosMappingsClassOfServiceEnum{
	"premium":   QosMappingsClassOfServicePremium,
	"default":   QosMappingsClassOfServiceDefault,
	"bulk":      QosMappingsClassOfServiceBulk,
	"scavenger": QosMappingsClassOfServiceScavenger,
}

// GetQosMappingsClassOfServiceEnumValues Enumerates the set of values for QosMappingsClassOfServiceEnum
func GetQosMappingsClassOfServiceEnumValues() []QosMappingsClassOfServiceEnum {
	values := make([]QosMappingsClassOfServiceEnum, 0)
	for _, v := range mappingQosMappingsClassOfServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetQosMappingsClassOfServiceEnumStringValues Enumerates the set of values in String for QosMappingsClassOfServiceEnum
func GetQosMappingsClassOfServiceEnumStringValues() []string {
	return []string{
		"PREMIUM",
		"DEFAULT",
		"BULK",
		"SCAVENGER",
	}
}

// GetMappingQosMappingsClassOfServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQosMappingsClassOfServiceEnum(val string) (QosMappingsClassOfServiceEnum, bool) {
	enum, ok := mappingQosMappingsClassOfServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
