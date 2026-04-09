// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddonOptionSummary Summary information about a AddonOptions.
type AddonOptionSummary struct {

	// A user-friendly name.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the KafkaClusterAddon.
	LifecycleState AddonOptionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m AddonOptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddonOptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddonOptionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAddonOptionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddonOptionSummaryLifecycleStateEnum Enum with underlying type: string
type AddonOptionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AddonOptionSummaryLifecycleStateEnum
const (
	AddonOptionSummaryLifecycleStateActive   AddonOptionSummaryLifecycleStateEnum = "ACTIVE"
	AddonOptionSummaryLifecycleStateInactive AddonOptionSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingAddonOptionSummaryLifecycleStateEnum = map[string]AddonOptionSummaryLifecycleStateEnum{
	"ACTIVE":   AddonOptionSummaryLifecycleStateActive,
	"INACTIVE": AddonOptionSummaryLifecycleStateInactive,
}

var mappingAddonOptionSummaryLifecycleStateEnumLowerCase = map[string]AddonOptionSummaryLifecycleStateEnum{
	"active":   AddonOptionSummaryLifecycleStateActive,
	"inactive": AddonOptionSummaryLifecycleStateInactive,
}

// GetAddonOptionSummaryLifecycleStateEnumValues Enumerates the set of values for AddonOptionSummaryLifecycleStateEnum
func GetAddonOptionSummaryLifecycleStateEnumValues() []AddonOptionSummaryLifecycleStateEnum {
	values := make([]AddonOptionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAddonOptionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAddonOptionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AddonOptionSummaryLifecycleStateEnum
func GetAddonOptionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingAddonOptionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddonOptionSummaryLifecycleStateEnum(val string) (AddonOptionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAddonOptionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
