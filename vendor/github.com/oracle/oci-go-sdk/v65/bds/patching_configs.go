// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchingConfigs Detailed configurations for defining the behavior when installing os patches. If not provided, nodes will be patched and rebooted AD/FD by AD/FD.
type PatchingConfigs interface {
}

type patchingconfigs struct {
	JsonData               []byte
	PatchingConfigStrategy string `json:"patchingConfigStrategy"`
}

// UnmarshalJSON unmarshals json
func (m *patchingconfigs) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchingconfigs patchingconfigs
	s := struct {
		Model Unmarshalerpatchingconfigs
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PatchingConfigStrategy = s.Model.PatchingConfigStrategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *patchingconfigs) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PatchingConfigStrategy {
	case "BATCHING_BASED":
		mm := BatchingBasedPatchingConfigs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOWNTIME_BASED":
		mm := DowntimeBasedPatchingConfigs{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PatchingConfigs: %s.", m.PatchingConfigStrategy)
		return *m, nil
	}
}

func (m patchingconfigs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchingconfigs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchingConfigsPatchingConfigStrategyEnum Enum with underlying type: string
type PatchingConfigsPatchingConfigStrategyEnum string

// Set of constants representing the allowable values for PatchingConfigsPatchingConfigStrategyEnum
const (
	PatchingConfigsPatchingConfigStrategyDowntimeBased PatchingConfigsPatchingConfigStrategyEnum = "DOWNTIME_BASED"
	PatchingConfigsPatchingConfigStrategyBatchingBased PatchingConfigsPatchingConfigStrategyEnum = "BATCHING_BASED"
)

var mappingPatchingConfigsPatchingConfigStrategyEnum = map[string]PatchingConfigsPatchingConfigStrategyEnum{
	"DOWNTIME_BASED": PatchingConfigsPatchingConfigStrategyDowntimeBased,
	"BATCHING_BASED": PatchingConfigsPatchingConfigStrategyBatchingBased,
}

var mappingPatchingConfigsPatchingConfigStrategyEnumLowerCase = map[string]PatchingConfigsPatchingConfigStrategyEnum{
	"downtime_based": PatchingConfigsPatchingConfigStrategyDowntimeBased,
	"batching_based": PatchingConfigsPatchingConfigStrategyBatchingBased,
}

// GetPatchingConfigsPatchingConfigStrategyEnumValues Enumerates the set of values for PatchingConfigsPatchingConfigStrategyEnum
func GetPatchingConfigsPatchingConfigStrategyEnumValues() []PatchingConfigsPatchingConfigStrategyEnum {
	values := make([]PatchingConfigsPatchingConfigStrategyEnum, 0)
	for _, v := range mappingPatchingConfigsPatchingConfigStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchingConfigsPatchingConfigStrategyEnumStringValues Enumerates the set of values in String for PatchingConfigsPatchingConfigStrategyEnum
func GetPatchingConfigsPatchingConfigStrategyEnumStringValues() []string {
	return []string{
		"DOWNTIME_BASED",
		"BATCHING_BASED",
	}
}

// GetMappingPatchingConfigsPatchingConfigStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchingConfigsPatchingConfigStrategyEnum(val string) (PatchingConfigsPatchingConfigStrategyEnum, bool) {
	enum, ok := mappingPatchingConfigsPatchingConfigStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
