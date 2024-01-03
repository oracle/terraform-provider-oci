// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PurgeWebAppAccelerationCacheDetails Specifies options for a cache purge.
type PurgeWebAppAccelerationCacheDetails interface {
}

type purgewebappaccelerationcachedetails struct {
	JsonData  []byte
	PurgeType string `json:"purgeType"`
}

// UnmarshalJSON unmarshals json
func (m *purgewebappaccelerationcachedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpurgewebappaccelerationcachedetails purgewebappaccelerationcachedetails
	s := struct {
		Model Unmarshalerpurgewebappaccelerationcachedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PurgeType = s.Model.PurgeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *purgewebappaccelerationcachedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PurgeType {
	case "ENTIRE_CACHE":
		mm := PurgeEntireWebAppAccelerationCacheDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PurgeWebAppAccelerationCacheDetails: %s.", m.PurgeType)
		return *m, nil
	}
}

func (m purgewebappaccelerationcachedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m purgewebappaccelerationcachedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum Enum with underlying type: string
type PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum string

// Set of constants representing the allowable values for PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum
const (
	PurgeWebAppAccelerationCacheDetailsPurgeTypeEntireCache PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum = "ENTIRE_CACHE"
)

var mappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnum = map[string]PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum{
	"ENTIRE_CACHE": PurgeWebAppAccelerationCacheDetailsPurgeTypeEntireCache,
}

var mappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumLowerCase = map[string]PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum{
	"entire_cache": PurgeWebAppAccelerationCacheDetailsPurgeTypeEntireCache,
}

// GetPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumValues Enumerates the set of values for PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum
func GetPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumValues() []PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum {
	values := make([]PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum, 0)
	for _, v := range mappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumStringValues Enumerates the set of values in String for PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum
func GetPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumStringValues() []string {
	return []string{
		"ENTIRE_CACHE",
	}
}

// GetMappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnum(val string) (PurgeWebAppAccelerationCacheDetailsPurgeTypeEnum, bool) {
	enum, ok := mappingPurgeWebAppAccelerationCacheDetailsPurgeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
