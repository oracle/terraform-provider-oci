// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgRedundancyStatus The redundancy status of the DRG. For more information, see
// Redundancy Remedies (https://docs.cloud.oracle.com/iaas/Content/Network/Troubleshoot/drgredundancy.htm).
type DrgRedundancyStatus struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	Id *string `mandatory:"false" json:"id"`

	// The redundancy status of the DRG.
	Status DrgRedundancyStatusStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m DrgRedundancyStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgRedundancyStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDrgRedundancyStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDrgRedundancyStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgRedundancyStatusStatusEnum Enum with underlying type: string
type DrgRedundancyStatusStatusEnum string

// Set of constants representing the allowable values for DrgRedundancyStatusStatusEnum
const (
	DrgRedundancyStatusStatusNotAvailable                        DrgRedundancyStatusStatusEnum = "NOT_AVAILABLE"
	DrgRedundancyStatusStatusRedundant                           DrgRedundancyStatusStatusEnum = "REDUNDANT"
	DrgRedundancyStatusStatusNotRedundantSingleIpsec             DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_SINGLE_IPSEC"
	DrgRedundancyStatusStatusNotRedundantSingleVirtualcircuit    DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_SINGLE_VIRTUALCIRCUIT"
	DrgRedundancyStatusStatusNotRedundantMultipleIpsecs          DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_MULTIPLE_IPSECS"
	DrgRedundancyStatusStatusNotRedundantMultipleVirtualcircuits DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_MULTIPLE_VIRTUALCIRCUITS"
	DrgRedundancyStatusStatusNotRedundantMixConnections          DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_MIX_CONNECTIONS"
	DrgRedundancyStatusStatusNotRedundantNoConnection            DrgRedundancyStatusStatusEnum = "NOT_REDUNDANT_NO_CONNECTION"
)

var mappingDrgRedundancyStatusStatusEnum = map[string]DrgRedundancyStatusStatusEnum{
	"NOT_AVAILABLE":                          DrgRedundancyStatusStatusNotAvailable,
	"REDUNDANT":                              DrgRedundancyStatusStatusRedundant,
	"NOT_REDUNDANT_SINGLE_IPSEC":             DrgRedundancyStatusStatusNotRedundantSingleIpsec,
	"NOT_REDUNDANT_SINGLE_VIRTUALCIRCUIT":    DrgRedundancyStatusStatusNotRedundantSingleVirtualcircuit,
	"NOT_REDUNDANT_MULTIPLE_IPSECS":          DrgRedundancyStatusStatusNotRedundantMultipleIpsecs,
	"NOT_REDUNDANT_MULTIPLE_VIRTUALCIRCUITS": DrgRedundancyStatusStatusNotRedundantMultipleVirtualcircuits,
	"NOT_REDUNDANT_MIX_CONNECTIONS":          DrgRedundancyStatusStatusNotRedundantMixConnections,
	"NOT_REDUNDANT_NO_CONNECTION":            DrgRedundancyStatusStatusNotRedundantNoConnection,
}

var mappingDrgRedundancyStatusStatusEnumLowerCase = map[string]DrgRedundancyStatusStatusEnum{
	"not_available":                          DrgRedundancyStatusStatusNotAvailable,
	"redundant":                              DrgRedundancyStatusStatusRedundant,
	"not_redundant_single_ipsec":             DrgRedundancyStatusStatusNotRedundantSingleIpsec,
	"not_redundant_single_virtualcircuit":    DrgRedundancyStatusStatusNotRedundantSingleVirtualcircuit,
	"not_redundant_multiple_ipsecs":          DrgRedundancyStatusStatusNotRedundantMultipleIpsecs,
	"not_redundant_multiple_virtualcircuits": DrgRedundancyStatusStatusNotRedundantMultipleVirtualcircuits,
	"not_redundant_mix_connections":          DrgRedundancyStatusStatusNotRedundantMixConnections,
	"not_redundant_no_connection":            DrgRedundancyStatusStatusNotRedundantNoConnection,
}

// GetDrgRedundancyStatusStatusEnumValues Enumerates the set of values for DrgRedundancyStatusStatusEnum
func GetDrgRedundancyStatusStatusEnumValues() []DrgRedundancyStatusStatusEnum {
	values := make([]DrgRedundancyStatusStatusEnum, 0)
	for _, v := range mappingDrgRedundancyStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgRedundancyStatusStatusEnumStringValues Enumerates the set of values in String for DrgRedundancyStatusStatusEnum
func GetDrgRedundancyStatusStatusEnumStringValues() []string {
	return []string{
		"NOT_AVAILABLE",
		"REDUNDANT",
		"NOT_REDUNDANT_SINGLE_IPSEC",
		"NOT_REDUNDANT_SINGLE_VIRTUALCIRCUIT",
		"NOT_REDUNDANT_MULTIPLE_IPSECS",
		"NOT_REDUNDANT_MULTIPLE_VIRTUALCIRCUITS",
		"NOT_REDUNDANT_MIX_CONNECTIONS",
		"NOT_REDUNDANT_NO_CONNECTION",
	}
}

// GetMappingDrgRedundancyStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgRedundancyStatusStatusEnum(val string) (DrgRedundancyStatusStatusEnum, bool) {
	enum, ok := mappingDrgRedundancyStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
