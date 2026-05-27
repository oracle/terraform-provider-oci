// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdditionalCapabilityDetails Additional details describing the selected capability.
type AdditionalCapabilityDetails interface {
}

type additionalcapabilitydetails struct {
	JsonData    []byte
	ServiceType string `json:"serviceType"`
}

// UnmarshalJSON unmarshals json
func (m *additionalcapabilitydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleradditionalcapabilitydetails additionalcapabilitydetails
	s := struct {
		Model Unmarshaleradditionalcapabilitydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ServiceType = s.Model.ServiceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *additionalcapabilitydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ServiceType {
	case "COMPUTE":
		mm := AdditionalComputeCapabilityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AdditionalCapabilityDetails: %s.", m.ServiceType)
		return *m, nil
	}
}

func (m additionalcapabilitydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m additionalcapabilitydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdditionalCapabilityDetailsServiceTypeEnum Enum with underlying type: string
type AdditionalCapabilityDetailsServiceTypeEnum string

// Set of constants representing the allowable values for AdditionalCapabilityDetailsServiceTypeEnum
const (
	AdditionalCapabilityDetailsServiceTypeCompute AdditionalCapabilityDetailsServiceTypeEnum = "COMPUTE"
)

var mappingAdditionalCapabilityDetailsServiceTypeEnum = map[string]AdditionalCapabilityDetailsServiceTypeEnum{
	"COMPUTE": AdditionalCapabilityDetailsServiceTypeCompute,
}

var mappingAdditionalCapabilityDetailsServiceTypeEnumLowerCase = map[string]AdditionalCapabilityDetailsServiceTypeEnum{
	"compute": AdditionalCapabilityDetailsServiceTypeCompute,
}

// GetAdditionalCapabilityDetailsServiceTypeEnumValues Enumerates the set of values for AdditionalCapabilityDetailsServiceTypeEnum
func GetAdditionalCapabilityDetailsServiceTypeEnumValues() []AdditionalCapabilityDetailsServiceTypeEnum {
	values := make([]AdditionalCapabilityDetailsServiceTypeEnum, 0)
	for _, v := range mappingAdditionalCapabilityDetailsServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAdditionalCapabilityDetailsServiceTypeEnumStringValues Enumerates the set of values in String for AdditionalCapabilityDetailsServiceTypeEnum
func GetAdditionalCapabilityDetailsServiceTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingAdditionalCapabilityDetailsServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdditionalCapabilityDetailsServiceTypeEnum(val string) (AdditionalCapabilityDetailsServiceTypeEnum, bool) {
	enum, ok := mappingAdditionalCapabilityDetailsServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
