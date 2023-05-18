// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// PublicIpCapacity A public IP capacity CIDR owned by VCN CP with its allocation status and metadata.
type PublicIpCapacity struct {

	// CIDR of the public IP capacity
	Cidr *string `mandatory:"true" json:"cidr"`

	// Public IP pool. This is the actual pool name not the ID (e.g. "marathon" is a valid pool ID where we would map
	// it to "oracle-prod" pool. For IPv6, which is not applicable now, the pool name is "default-v6" but the pool ID
	// is "default".)
	Pool *string `mandatory:"true" json:"pool"`

	// AD of the capacity. If it's in a regional pool, the scope is "reserved_no_ad".
	Scope *string `mandatory:"true" json:"scope"`

	// Status of the public IP CIDR whether it is currently available to allocate to users.
	CidrState PublicIpCapacityCidrStateEnum `mandatory:"true" json:"cidrState"`

	// Size of the address/CIDR configured to be allocated from this capacity.
	PrefixLength *int `mandatory:"true" json:"prefixLength"`

	// Public IP capacity metadata, such as capacity request JIRA ID.
	Metadata map[string]string `mandatory:"true" json:"metadata"`

	// Bitmap of the allocated IPs in big-endian order displayed in Base64 format. For example, once decoded a byte
	// array of [0,0,0,2] mean the second address of a /27 CIDR is allocated (Note that 2 is on array index 3).
	AllocatedBits *string `mandatory:"true" json:"allocatedBits"`
}

func (m PublicIpCapacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicIpCapacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPublicIpCapacityCidrStateEnum(string(m.CidrState)); !ok && m.CidrState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CidrState: %s. Supported values are: %s.", m.CidrState, strings.Join(GetPublicIpCapacityCidrStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicIpCapacityCidrStateEnum Enum with underlying type: string
type PublicIpCapacityCidrStateEnum string

// Set of constants representing the allowable values for PublicIpCapacityCidrStateEnum
const (
	PublicIpCapacityCidrStateReadyForTest PublicIpCapacityCidrStateEnum = "READY_FOR_TEST"
	PublicIpCapacityCidrStateFailedTest   PublicIpCapacityCidrStateEnum = "FAILED_TEST"
	PublicIpCapacityCidrStatePassedTest   PublicIpCapacityCidrStateEnum = "PASSED_TEST"
	PublicIpCapacityCidrStateAvailable    PublicIpCapacityCidrStateEnum = "AVAILABLE"
	PublicIpCapacityCidrStateDraining     PublicIpCapacityCidrStateEnum = "DRAINING"
	PublicIpCapacityCidrStateDrained      PublicIpCapacityCidrStateEnum = "DRAINED"
	PublicIpCapacityCidrStateMigrating    PublicIpCapacityCidrStateEnum = "MIGRATING"
	PublicIpCapacityCidrStateMigrated     PublicIpCapacityCidrStateEnum = "MIGRATED"
)

var mappingPublicIpCapacityCidrStateEnum = map[string]PublicIpCapacityCidrStateEnum{
	"READY_FOR_TEST": PublicIpCapacityCidrStateReadyForTest,
	"FAILED_TEST":    PublicIpCapacityCidrStateFailedTest,
	"PASSED_TEST":    PublicIpCapacityCidrStatePassedTest,
	"AVAILABLE":      PublicIpCapacityCidrStateAvailable,
	"DRAINING":       PublicIpCapacityCidrStateDraining,
	"DRAINED":        PublicIpCapacityCidrStateDrained,
	"MIGRATING":      PublicIpCapacityCidrStateMigrating,
	"MIGRATED":       PublicIpCapacityCidrStateMigrated,
}

var mappingPublicIpCapacityCidrStateEnumLowerCase = map[string]PublicIpCapacityCidrStateEnum{
	"ready_for_test": PublicIpCapacityCidrStateReadyForTest,
	"failed_test":    PublicIpCapacityCidrStateFailedTest,
	"passed_test":    PublicIpCapacityCidrStatePassedTest,
	"available":      PublicIpCapacityCidrStateAvailable,
	"draining":       PublicIpCapacityCidrStateDraining,
	"drained":        PublicIpCapacityCidrStateDrained,
	"migrating":      PublicIpCapacityCidrStateMigrating,
	"migrated":       PublicIpCapacityCidrStateMigrated,
}

// GetPublicIpCapacityCidrStateEnumValues Enumerates the set of values for PublicIpCapacityCidrStateEnum
func GetPublicIpCapacityCidrStateEnumValues() []PublicIpCapacityCidrStateEnum {
	values := make([]PublicIpCapacityCidrStateEnum, 0)
	for _, v := range mappingPublicIpCapacityCidrStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicIpCapacityCidrStateEnumStringValues Enumerates the set of values in String for PublicIpCapacityCidrStateEnum
func GetPublicIpCapacityCidrStateEnumStringValues() []string {
	return []string{
		"READY_FOR_TEST",
		"FAILED_TEST",
		"PASSED_TEST",
		"AVAILABLE",
		"DRAINING",
		"DRAINED",
		"MIGRATING",
		"MIGRATED",
	}
}

// GetMappingPublicIpCapacityCidrStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicIpCapacityCidrStateEnum(val string) (PublicIpCapacityCidrStateEnum, bool) {
	enum, ok := mappingPublicIpCapacityCidrStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
