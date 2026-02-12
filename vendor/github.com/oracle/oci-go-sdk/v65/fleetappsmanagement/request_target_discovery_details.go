// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestTargetDiscoveryDetails Request to initiate target discovery.
type RequestTargetDiscoveryDetails struct {

	// A boolean flag that decides if all resources within the fleet should be part of discovery.
	IsApplicableToAllResources *bool `mandatory:"false" json:"isApplicableToAllResources"`

	// Resource OCIDS to be included for discovery.
	ResourceIds []string `mandatory:"false" json:"resourceIds"`

	// A list of discovery process statuses. Each status represents a specific state in the workflow.
	Statuses []RequestTargetDiscoveryDetailsStatusesEnum `mandatory:"false" json:"statuses,omitempty"`
}

func (m RequestTargetDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestTargetDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Statuses {
		if _, ok := GetMappingRequestTargetDiscoveryDetailsStatusesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statuses: %s. Supported values are: %s.", val, strings.Join(GetRequestTargetDiscoveryDetailsStatusesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestTargetDiscoveryDetailsStatusesEnum Enum with underlying type: string
type RequestTargetDiscoveryDetailsStatusesEnum string

// Set of constants representing the allowable values for RequestTargetDiscoveryDetailsStatusesEnum
const (
	RequestTargetDiscoveryDetailsStatusesInitiated          RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_INITIATED"
	RequestTargetDiscoveryDetailsStatusesAccepted           RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_ACCEPTED"
	RequestTargetDiscoveryDetailsStatusesSubmitted          RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_SUBMITTED"
	RequestTargetDiscoveryDetailsStatusesInProgress         RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_IN_PROGRESS"
	RequestTargetDiscoveryDetailsStatusesFailed             RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_FAILED"
	RequestTargetDiscoveryDetailsStatusesSkipped            RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_SKIPPED"
	RequestTargetDiscoveryDetailsStatusesPartiallySucceeded RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_PARTIALLY_SUCCEEDED"
	RequestTargetDiscoveryDetailsStatusesTimedOut           RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_TIMED_OUT"
	RequestTargetDiscoveryDetailsStatusesSucceeded          RequestTargetDiscoveryDetailsStatusesEnum = "DISCOVERY_SUCCEEDED"
)

var mappingRequestTargetDiscoveryDetailsStatusesEnum = map[string]RequestTargetDiscoveryDetailsStatusesEnum{
	"DISCOVERY_INITIATED":           RequestTargetDiscoveryDetailsStatusesInitiated,
	"DISCOVERY_ACCEPTED":            RequestTargetDiscoveryDetailsStatusesAccepted,
	"DISCOVERY_SUBMITTED":           RequestTargetDiscoveryDetailsStatusesSubmitted,
	"DISCOVERY_IN_PROGRESS":         RequestTargetDiscoveryDetailsStatusesInProgress,
	"DISCOVERY_FAILED":              RequestTargetDiscoveryDetailsStatusesFailed,
	"DISCOVERY_SKIPPED":             RequestTargetDiscoveryDetailsStatusesSkipped,
	"DISCOVERY_PARTIALLY_SUCCEEDED": RequestTargetDiscoveryDetailsStatusesPartiallySucceeded,
	"DISCOVERY_TIMED_OUT":           RequestTargetDiscoveryDetailsStatusesTimedOut,
	"DISCOVERY_SUCCEEDED":           RequestTargetDiscoveryDetailsStatusesSucceeded,
}

var mappingRequestTargetDiscoveryDetailsStatusesEnumLowerCase = map[string]RequestTargetDiscoveryDetailsStatusesEnum{
	"discovery_initiated":           RequestTargetDiscoveryDetailsStatusesInitiated,
	"discovery_accepted":            RequestTargetDiscoveryDetailsStatusesAccepted,
	"discovery_submitted":           RequestTargetDiscoveryDetailsStatusesSubmitted,
	"discovery_in_progress":         RequestTargetDiscoveryDetailsStatusesInProgress,
	"discovery_failed":              RequestTargetDiscoveryDetailsStatusesFailed,
	"discovery_skipped":             RequestTargetDiscoveryDetailsStatusesSkipped,
	"discovery_partially_succeeded": RequestTargetDiscoveryDetailsStatusesPartiallySucceeded,
	"discovery_timed_out":           RequestTargetDiscoveryDetailsStatusesTimedOut,
	"discovery_succeeded":           RequestTargetDiscoveryDetailsStatusesSucceeded,
}

// GetRequestTargetDiscoveryDetailsStatusesEnumValues Enumerates the set of values for RequestTargetDiscoveryDetailsStatusesEnum
func GetRequestTargetDiscoveryDetailsStatusesEnumValues() []RequestTargetDiscoveryDetailsStatusesEnum {
	values := make([]RequestTargetDiscoveryDetailsStatusesEnum, 0)
	for _, v := range mappingRequestTargetDiscoveryDetailsStatusesEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestTargetDiscoveryDetailsStatusesEnumStringValues Enumerates the set of values in String for RequestTargetDiscoveryDetailsStatusesEnum
func GetRequestTargetDiscoveryDetailsStatusesEnumStringValues() []string {
	return []string{
		"DISCOVERY_INITIATED",
		"DISCOVERY_ACCEPTED",
		"DISCOVERY_SUBMITTED",
		"DISCOVERY_IN_PROGRESS",
		"DISCOVERY_FAILED",
		"DISCOVERY_SKIPPED",
		"DISCOVERY_PARTIALLY_SUCCEEDED",
		"DISCOVERY_TIMED_OUT",
		"DISCOVERY_SUCCEEDED",
	}
}

// GetMappingRequestTargetDiscoveryDetailsStatusesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestTargetDiscoveryDetailsStatusesEnum(val string) (RequestTargetDiscoveryDetailsStatusesEnum, bool) {
	enum, ok := mappingRequestTargetDiscoveryDetailsStatusesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
