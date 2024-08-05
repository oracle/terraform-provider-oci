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

// OdhPatchingConfig Detailed configurations for defining the behavior when installing ODH patches. If not provided, nodes will be patched with down time.
type OdhPatchingConfig interface {
}

type odhpatchingconfig struct {
	JsonData               []byte
	PatchingConfigStrategy string `json:"patchingConfigStrategy"`
}

// UnmarshalJSON unmarshals json
func (m *odhpatchingconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerodhpatchingconfig odhpatchingconfig
	s := struct {
		Model Unmarshalerodhpatchingconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PatchingConfigStrategy = s.Model.PatchingConfigStrategy

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *odhpatchingconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PatchingConfigStrategy {
	case "DOWNTIME_BASED":
		mm := DowntimeBasedOdhPatchingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOMAIN_BASED":
		mm := DomainBasedOdhPatchingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BATCHING_BASED":
		mm := BatchingBasedOdhPatchingConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OdhPatchingConfig: %s.", m.PatchingConfigStrategy)
		return *m, nil
	}
}

func (m odhpatchingconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m odhpatchingconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdhPatchingConfigPatchingConfigStrategyEnum Enum with underlying type: string
type OdhPatchingConfigPatchingConfigStrategyEnum string

// Set of constants representing the allowable values for OdhPatchingConfigPatchingConfigStrategyEnum
const (
	OdhPatchingConfigPatchingConfigStrategyDowntimeBased OdhPatchingConfigPatchingConfigStrategyEnum = "DOWNTIME_BASED"
	OdhPatchingConfigPatchingConfigStrategyBatchingBased OdhPatchingConfigPatchingConfigStrategyEnum = "BATCHING_BASED"
	OdhPatchingConfigPatchingConfigStrategyDomainBased   OdhPatchingConfigPatchingConfigStrategyEnum = "DOMAIN_BASED"
)

var mappingOdhPatchingConfigPatchingConfigStrategyEnum = map[string]OdhPatchingConfigPatchingConfigStrategyEnum{
	"DOWNTIME_BASED": OdhPatchingConfigPatchingConfigStrategyDowntimeBased,
	"BATCHING_BASED": OdhPatchingConfigPatchingConfigStrategyBatchingBased,
	"DOMAIN_BASED":   OdhPatchingConfigPatchingConfigStrategyDomainBased,
}

var mappingOdhPatchingConfigPatchingConfigStrategyEnumLowerCase = map[string]OdhPatchingConfigPatchingConfigStrategyEnum{
	"downtime_based": OdhPatchingConfigPatchingConfigStrategyDowntimeBased,
	"batching_based": OdhPatchingConfigPatchingConfigStrategyBatchingBased,
	"domain_based":   OdhPatchingConfigPatchingConfigStrategyDomainBased,
}

// GetOdhPatchingConfigPatchingConfigStrategyEnumValues Enumerates the set of values for OdhPatchingConfigPatchingConfigStrategyEnum
func GetOdhPatchingConfigPatchingConfigStrategyEnumValues() []OdhPatchingConfigPatchingConfigStrategyEnum {
	values := make([]OdhPatchingConfigPatchingConfigStrategyEnum, 0)
	for _, v := range mappingOdhPatchingConfigPatchingConfigStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetOdhPatchingConfigPatchingConfigStrategyEnumStringValues Enumerates the set of values in String for OdhPatchingConfigPatchingConfigStrategyEnum
func GetOdhPatchingConfigPatchingConfigStrategyEnumStringValues() []string {
	return []string{
		"DOWNTIME_BASED",
		"BATCHING_BASED",
		"DOMAIN_BASED",
	}
}

// GetMappingOdhPatchingConfigPatchingConfigStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdhPatchingConfigPatchingConfigStrategyEnum(val string) (OdhPatchingConfigPatchingConfigStrategyEnum, bool) {
	enum, ok := mappingOdhPatchingConfigPatchingConfigStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
