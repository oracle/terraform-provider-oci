// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RollbackDetails Rollback details specified for the action.
type RollbackDetails interface {
}

type rollbackdetails struct {
	JsonData []byte
	Strategy string `json:"strategy"`
}

// UnmarshalJSON unmarshals json
func (m *rollbackdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrollbackdetails rollbackdetails
	s := struct {
		Model Unmarshalerrollbackdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Strategy = s.Model.Strategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *rollbackdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Strategy {
	case "LIST_OF_TARGETS":
		mm := ListOfTargetsRollbackDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILED_JOBS":
		mm := FailedJobsRollbackDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RollbackDetails: %s.", m.Strategy)
		return *m, nil
	}
}

func (m rollbackdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m rollbackdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RollbackDetailsStrategyEnum Enum with underlying type: string
type RollbackDetailsStrategyEnum string

// Set of constants representing the allowable values for RollbackDetailsStrategyEnum
const (
	RollbackDetailsStrategyFailedJobs    RollbackDetailsStrategyEnum = "FAILED_JOBS"
	RollbackDetailsStrategyListOfTargets RollbackDetailsStrategyEnum = "LIST_OF_TARGETS"
)

var mappingRollbackDetailsStrategyEnum = map[string]RollbackDetailsStrategyEnum{
	"FAILED_JOBS":     RollbackDetailsStrategyFailedJobs,
	"LIST_OF_TARGETS": RollbackDetailsStrategyListOfTargets,
}

var mappingRollbackDetailsStrategyEnumLowerCase = map[string]RollbackDetailsStrategyEnum{
	"failed_jobs":     RollbackDetailsStrategyFailedJobs,
	"list_of_targets": RollbackDetailsStrategyListOfTargets,
}

// GetRollbackDetailsStrategyEnumValues Enumerates the set of values for RollbackDetailsStrategyEnum
func GetRollbackDetailsStrategyEnumValues() []RollbackDetailsStrategyEnum {
	values := make([]RollbackDetailsStrategyEnum, 0)
	for _, v := range mappingRollbackDetailsStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetRollbackDetailsStrategyEnumStringValues Enumerates the set of values in String for RollbackDetailsStrategyEnum
func GetRollbackDetailsStrategyEnumStringValues() []string {
	return []string{
		"FAILED_JOBS",
		"LIST_OF_TARGETS",
	}
}

// GetMappingRollbackDetailsStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRollbackDetailsStrategyEnum(val string) (RollbackDetailsStrategyEnum, bool) {
	enum, ok := mappingRollbackDetailsStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
