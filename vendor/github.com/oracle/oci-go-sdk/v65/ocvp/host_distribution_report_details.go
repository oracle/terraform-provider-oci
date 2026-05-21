// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostDistributionReportDetails Details for how ESXi hosts are distributed across Fault Domains in a Cluster.
type HostDistributionReportDetails struct {
	FaultDomainHostDistributionState HostDistributionReportDetailsFaultDomainHostDistributionStateEnum `mandatory:"true" json:"faultDomainHostDistributionState"`
}

func (m HostDistributionReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostDistributionReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnum(string(m.FaultDomainHostDistributionState)); !ok && m.FaultDomainHostDistributionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FaultDomainHostDistributionState: %s. Supported values are: %s.", m.FaultDomainHostDistributionState, strings.Join(GetHostDistributionReportDetailsFaultDomainHostDistributionStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostDistributionReportDetailsFaultDomainHostDistributionStateEnum Enum with underlying type: string
type HostDistributionReportDetailsFaultDomainHostDistributionStateEnum string

// Set of constants representing the allowable values for HostDistributionReportDetailsFaultDomainHostDistributionStateEnum
const (
	HostDistributionReportDetailsFaultDomainHostDistributionStateEvenlyDistributed   HostDistributionReportDetailsFaultDomainHostDistributionStateEnum = "EVENLY_DISTRIBUTED"
	HostDistributionReportDetailsFaultDomainHostDistributionStateUnevenlyDistributed HostDistributionReportDetailsFaultDomainHostDistributionStateEnum = "UNEVENLY_DISTRIBUTED"
	HostDistributionReportDetailsFaultDomainHostDistributionStateUnsupported         HostDistributionReportDetailsFaultDomainHostDistributionStateEnum = "UNSUPPORTED"
)

var mappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnum = map[string]HostDistributionReportDetailsFaultDomainHostDistributionStateEnum{
	"EVENLY_DISTRIBUTED":   HostDistributionReportDetailsFaultDomainHostDistributionStateEvenlyDistributed,
	"UNEVENLY_DISTRIBUTED": HostDistributionReportDetailsFaultDomainHostDistributionStateUnevenlyDistributed,
	"UNSUPPORTED":          HostDistributionReportDetailsFaultDomainHostDistributionStateUnsupported,
}

var mappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnumLowerCase = map[string]HostDistributionReportDetailsFaultDomainHostDistributionStateEnum{
	"evenly_distributed":   HostDistributionReportDetailsFaultDomainHostDistributionStateEvenlyDistributed,
	"unevenly_distributed": HostDistributionReportDetailsFaultDomainHostDistributionStateUnevenlyDistributed,
	"unsupported":          HostDistributionReportDetailsFaultDomainHostDistributionStateUnsupported,
}

// GetHostDistributionReportDetailsFaultDomainHostDistributionStateEnumValues Enumerates the set of values for HostDistributionReportDetailsFaultDomainHostDistributionStateEnum
func GetHostDistributionReportDetailsFaultDomainHostDistributionStateEnumValues() []HostDistributionReportDetailsFaultDomainHostDistributionStateEnum {
	values := make([]HostDistributionReportDetailsFaultDomainHostDistributionStateEnum, 0)
	for _, v := range mappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHostDistributionReportDetailsFaultDomainHostDistributionStateEnumStringValues Enumerates the set of values in String for HostDistributionReportDetailsFaultDomainHostDistributionStateEnum
func GetHostDistributionReportDetailsFaultDomainHostDistributionStateEnumStringValues() []string {
	return []string{
		"EVENLY_DISTRIBUTED",
		"UNEVENLY_DISTRIBUTED",
		"UNSUPPORTED",
	}
}

// GetMappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnum(val string) (HostDistributionReportDetailsFaultDomainHostDistributionStateEnum, bool) {
	enum, ok := mappingHostDistributionReportDetailsFaultDomainHostDistributionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
