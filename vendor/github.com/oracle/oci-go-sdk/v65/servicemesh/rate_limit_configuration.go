// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RateLimitConfiguration Rate limiting configuration which is used to control the number of requests a service is allowed to handle within a given time period.
type RateLimitConfiguration interface {

	// The interval for which requests are limited.
	GetIntervalInMs() *int64
}

type ratelimitconfiguration struct {
	JsonData     []byte
	IntervalInMs *int64 `mandatory:"true" json:"intervalInMs"`
	Type         string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *ratelimitconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerratelimitconfiguration ratelimitconfiguration
	s := struct {
		Model Unmarshalerratelimitconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IntervalInMs = s.Model.IntervalInMs
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ratelimitconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TCP":
		mm := TcpRateLimitConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpRateLimitConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RateLimitConfiguration: %s.", m.Type)
		return *m, nil
	}
}

// GetIntervalInMs returns IntervalInMs
func (m ratelimitconfiguration) GetIntervalInMs() *int64 {
	return m.IntervalInMs
}

func (m ratelimitconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ratelimitconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RateLimitConfigurationTypeEnum Enum with underlying type: string
type RateLimitConfigurationTypeEnum string

// Set of constants representing the allowable values for RateLimitConfigurationTypeEnum
const (
	RateLimitConfigurationTypeHttp RateLimitConfigurationTypeEnum = "HTTP"
	RateLimitConfigurationTypeTcp  RateLimitConfigurationTypeEnum = "TCP"
)

var mappingRateLimitConfigurationTypeEnum = map[string]RateLimitConfigurationTypeEnum{
	"HTTP": RateLimitConfigurationTypeHttp,
	"TCP":  RateLimitConfigurationTypeTcp,
}

var mappingRateLimitConfigurationTypeEnumLowerCase = map[string]RateLimitConfigurationTypeEnum{
	"http": RateLimitConfigurationTypeHttp,
	"tcp":  RateLimitConfigurationTypeTcp,
}

// GetRateLimitConfigurationTypeEnumValues Enumerates the set of values for RateLimitConfigurationTypeEnum
func GetRateLimitConfigurationTypeEnumValues() []RateLimitConfigurationTypeEnum {
	values := make([]RateLimitConfigurationTypeEnum, 0)
	for _, v := range mappingRateLimitConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRateLimitConfigurationTypeEnumStringValues Enumerates the set of values in String for RateLimitConfigurationTypeEnum
func GetRateLimitConfigurationTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TCP",
	}
}

// GetMappingRateLimitConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRateLimitConfigurationTypeEnum(val string) (RateLimitConfigurationTypeEnum, bool) {
	enum, ok := mappingRateLimitConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
