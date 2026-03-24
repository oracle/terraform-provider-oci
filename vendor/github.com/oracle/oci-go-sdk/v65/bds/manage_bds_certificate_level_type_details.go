// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManageBdsCertificateLevelTypeDetails Details of the type of level used to trigger certificate generation or renewal.
type ManageBdsCertificateLevelTypeDetails interface {
}

type managebdscertificateleveltypedetails struct {
	JsonData    []byte
	TriggerType string `json:"triggerType"`
}

// UnmarshalJSON unmarshals json
func (m *managebdscertificateleveltypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagebdscertificateleveltypedetails managebdscertificateleveltypedetails
	s := struct {
		Model Unmarshalermanagebdscertificateleveltypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TriggerType = s.Model.TriggerType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managebdscertificateleveltypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerType {
	case "NODE_LEVEL":
		mm := NodeLevelManageBdsCertificateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONFIG_LEVEL":
		mm := ConfigLevelManageBdsCertificateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManageBdsCertificateLevelTypeDetails: %s.", m.TriggerType)
		return *m, nil
	}
}

func (m managebdscertificateleveltypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managebdscertificateleveltypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum Enum with underlying type: string
type ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum string

// Set of constants representing the allowable values for ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum
const (
	ManageBdsCertificateLevelTypeDetailsTriggerTypeNodeLevel   ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum = "NODE_LEVEL"
	ManageBdsCertificateLevelTypeDetailsTriggerTypeConfigLevel ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum = "CONFIG_LEVEL"
)

var mappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnum = map[string]ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum{
	"NODE_LEVEL":   ManageBdsCertificateLevelTypeDetailsTriggerTypeNodeLevel,
	"CONFIG_LEVEL": ManageBdsCertificateLevelTypeDetailsTriggerTypeConfigLevel,
}

var mappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnumLowerCase = map[string]ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum{
	"node_level":   ManageBdsCertificateLevelTypeDetailsTriggerTypeNodeLevel,
	"config_level": ManageBdsCertificateLevelTypeDetailsTriggerTypeConfigLevel,
}

// GetManageBdsCertificateLevelTypeDetailsTriggerTypeEnumValues Enumerates the set of values for ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum
func GetManageBdsCertificateLevelTypeDetailsTriggerTypeEnumValues() []ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum {
	values := make([]ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum, 0)
	for _, v := range mappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManageBdsCertificateLevelTypeDetailsTriggerTypeEnumStringValues Enumerates the set of values in String for ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum
func GetManageBdsCertificateLevelTypeDetailsTriggerTypeEnumStringValues() []string {
	return []string{
		"NODE_LEVEL",
		"CONFIG_LEVEL",
	}
}

// GetMappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnum(val string) (ManageBdsCertificateLevelTypeDetailsTriggerTypeEnum, bool) {
	enum, ok := mappingManageBdsCertificateLevelTypeDetailsTriggerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
