// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PollNcpJobDetails Poll Ncp job
type PollNcpJobDetails struct {

	// Current state of NCP job
	JobState PollNcpJobDetailsJobStateEnum `mandatory:"true" json:"jobState"`
}

func (m PollNcpJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PollNcpJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPollNcpJobDetailsJobStateEnum(string(m.JobState)); !ok && m.JobState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JobState: %s. Supported values are: %s.", m.JobState, strings.Join(GetPollNcpJobDetailsJobStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PollNcpJobDetailsJobStateEnum Enum with underlying type: string
type PollNcpJobDetailsJobStateEnum string

// Set of constants representing the allowable values for PollNcpJobDetailsJobStateEnum
const (
	PollNcpJobDetailsJobStateNew       PollNcpJobDetailsJobStateEnum = "NEW"
	PollNcpJobDetailsJobStatePending   PollNcpJobDetailsJobStateEnum = "PENDING"
	PollNcpJobDetailsJobStateStarted   PollNcpJobDetailsJobStateEnum = "STARTED"
	PollNcpJobDetailsJobStateScheduled PollNcpJobDetailsJobStateEnum = "SCHEDULED"
	PollNcpJobDetailsJobStateCanceled  PollNcpJobDetailsJobStateEnum = "CANCELED"
	PollNcpJobDetailsJobStateFailed    PollNcpJobDetailsJobStateEnum = "FAILED"
	PollNcpJobDetailsJobStateError     PollNcpJobDetailsJobStateEnum = "ERROR"
	PollNcpJobDetailsJobStateTimeout   PollNcpJobDetailsJobStateEnum = "TIMEOUT"
	PollNcpJobDetailsJobStateSucceeded PollNcpJobDetailsJobStateEnum = "SUCCEEDED"
)

var mappingPollNcpJobDetailsJobStateEnum = map[string]PollNcpJobDetailsJobStateEnum{
	"NEW":       PollNcpJobDetailsJobStateNew,
	"PENDING":   PollNcpJobDetailsJobStatePending,
	"STARTED":   PollNcpJobDetailsJobStateStarted,
	"SCHEDULED": PollNcpJobDetailsJobStateScheduled,
	"CANCELED":  PollNcpJobDetailsJobStateCanceled,
	"FAILED":    PollNcpJobDetailsJobStateFailed,
	"ERROR":     PollNcpJobDetailsJobStateError,
	"TIMEOUT":   PollNcpJobDetailsJobStateTimeout,
	"SUCCEEDED": PollNcpJobDetailsJobStateSucceeded,
}

var mappingPollNcpJobDetailsJobStateEnumLowerCase = map[string]PollNcpJobDetailsJobStateEnum{
	"new":       PollNcpJobDetailsJobStateNew,
	"pending":   PollNcpJobDetailsJobStatePending,
	"started":   PollNcpJobDetailsJobStateStarted,
	"scheduled": PollNcpJobDetailsJobStateScheduled,
	"canceled":  PollNcpJobDetailsJobStateCanceled,
	"failed":    PollNcpJobDetailsJobStateFailed,
	"error":     PollNcpJobDetailsJobStateError,
	"timeout":   PollNcpJobDetailsJobStateTimeout,
	"succeeded": PollNcpJobDetailsJobStateSucceeded,
}

// GetPollNcpJobDetailsJobStateEnumValues Enumerates the set of values for PollNcpJobDetailsJobStateEnum
func GetPollNcpJobDetailsJobStateEnumValues() []PollNcpJobDetailsJobStateEnum {
	values := make([]PollNcpJobDetailsJobStateEnum, 0)
	for _, v := range mappingPollNcpJobDetailsJobStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPollNcpJobDetailsJobStateEnumStringValues Enumerates the set of values in String for PollNcpJobDetailsJobStateEnum
func GetPollNcpJobDetailsJobStateEnumStringValues() []string {
	return []string{
		"NEW",
		"PENDING",
		"STARTED",
		"SCHEDULED",
		"CANCELED",
		"FAILED",
		"ERROR",
		"TIMEOUT",
		"SUCCEEDED",
	}
}

// GetMappingPollNcpJobDetailsJobStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPollNcpJobDetailsJobStateEnum(val string) (PollNcpJobDetailsJobStateEnum, bool) {
	enum, ok := mappingPollNcpJobDetailsJobStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
