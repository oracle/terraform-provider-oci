// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigureIotDomainDataAccessDetails The configuration details for IoT Domain Data Access.
type ConfigureIotDomainDataAccessDetails interface {
}

type configureiotdomaindataaccessdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *configureiotdomaindataaccessdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigureiotdomaindataaccessdetails configureiotdomaindataaccessdetails
	s := struct {
		Model Unmarshalerconfigureiotdomaindataaccessdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configureiotdomaindataaccessdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "APEX":
		mm := ApexDataAccessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DIRECT":
		mm := DirectDataAccessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORDS":
		mm := OrdsDataAccessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ConfigureIotDomainDataAccessDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m configureiotdomaindataaccessdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configureiotdomaindataaccessdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigureIotDomainDataAccessDetailsTypeEnum Enum with underlying type: string
type ConfigureIotDomainDataAccessDetailsTypeEnum string

// Set of constants representing the allowable values for ConfigureIotDomainDataAccessDetailsTypeEnum
const (
	ConfigureIotDomainDataAccessDetailsTypeDirect ConfigureIotDomainDataAccessDetailsTypeEnum = "DIRECT"
	ConfigureIotDomainDataAccessDetailsTypeOrds   ConfigureIotDomainDataAccessDetailsTypeEnum = "ORDS"
	ConfigureIotDomainDataAccessDetailsTypeApex   ConfigureIotDomainDataAccessDetailsTypeEnum = "APEX"
)

var mappingConfigureIotDomainDataAccessDetailsTypeEnum = map[string]ConfigureIotDomainDataAccessDetailsTypeEnum{
	"DIRECT": ConfigureIotDomainDataAccessDetailsTypeDirect,
	"ORDS":   ConfigureIotDomainDataAccessDetailsTypeOrds,
	"APEX":   ConfigureIotDomainDataAccessDetailsTypeApex,
}

var mappingConfigureIotDomainDataAccessDetailsTypeEnumLowerCase = map[string]ConfigureIotDomainDataAccessDetailsTypeEnum{
	"direct": ConfigureIotDomainDataAccessDetailsTypeDirect,
	"ords":   ConfigureIotDomainDataAccessDetailsTypeOrds,
	"apex":   ConfigureIotDomainDataAccessDetailsTypeApex,
}

// GetConfigureIotDomainDataAccessDetailsTypeEnumValues Enumerates the set of values for ConfigureIotDomainDataAccessDetailsTypeEnum
func GetConfigureIotDomainDataAccessDetailsTypeEnumValues() []ConfigureIotDomainDataAccessDetailsTypeEnum {
	values := make([]ConfigureIotDomainDataAccessDetailsTypeEnum, 0)
	for _, v := range mappingConfigureIotDomainDataAccessDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigureIotDomainDataAccessDetailsTypeEnumStringValues Enumerates the set of values in String for ConfigureIotDomainDataAccessDetailsTypeEnum
func GetConfigureIotDomainDataAccessDetailsTypeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"ORDS",
		"APEX",
	}
}

// GetMappingConfigureIotDomainDataAccessDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigureIotDomainDataAccessDetailsTypeEnum(val string) (ConfigureIotDomainDataAccessDetailsTypeEnum, bool) {
	enum, ok := mappingConfigureIotDomainDataAccessDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
