// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigCategoryDetails Config Category Details.
type ConfigCategoryDetails interface {
}

type configcategorydetails struct {
	JsonData       []byte
	ConfigCategory string `json:"configCategory"`
}

// UnmarshalJSON unmarshals json
func (m *configcategorydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigcategorydetails configcategorydetails
	s := struct {
		Model Unmarshalerconfigcategorydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigCategory = s.Model.ConfigCategory

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configcategorydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigCategory {
	case "PRODUCT_STACK":
		mm := ProductStackConfigCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ENVIRONMENT":
		mm := EnvironmentConfigCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREDENTIAL":
		mm := CredentialConfigCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PATCH_TYPE":
		mm := PatchTypeConfigCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRODUCT":
		mm := ProductConfigCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConfigCategoryDetails: %s.", m.ConfigCategory)
		return *m, nil
	}
}

func (m configcategorydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configcategorydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigCategoryDetailsConfigCategoryEnum Enum with underlying type: string
type ConfigCategoryDetailsConfigCategoryEnum string

// Set of constants representing the allowable values for ConfigCategoryDetailsConfigCategoryEnum
const (
	ConfigCategoryDetailsConfigCategoryProduct      ConfigCategoryDetailsConfigCategoryEnum = "PRODUCT"
	ConfigCategoryDetailsConfigCategoryProductStack ConfigCategoryDetailsConfigCategoryEnum = "PRODUCT_STACK"
	ConfigCategoryDetailsConfigCategoryEnvironment  ConfigCategoryDetailsConfigCategoryEnum = "ENVIRONMENT"
	ConfigCategoryDetailsConfigCategoryPatchType    ConfigCategoryDetailsConfigCategoryEnum = "PATCH_TYPE"
	ConfigCategoryDetailsConfigCategoryCredential   ConfigCategoryDetailsConfigCategoryEnum = "CREDENTIAL"
)

var mappingConfigCategoryDetailsConfigCategoryEnum = map[string]ConfigCategoryDetailsConfigCategoryEnum{
	"PRODUCT":       ConfigCategoryDetailsConfigCategoryProduct,
	"PRODUCT_STACK": ConfigCategoryDetailsConfigCategoryProductStack,
	"ENVIRONMENT":   ConfigCategoryDetailsConfigCategoryEnvironment,
	"PATCH_TYPE":    ConfigCategoryDetailsConfigCategoryPatchType,
	"CREDENTIAL":    ConfigCategoryDetailsConfigCategoryCredential,
}

var mappingConfigCategoryDetailsConfigCategoryEnumLowerCase = map[string]ConfigCategoryDetailsConfigCategoryEnum{
	"product":       ConfigCategoryDetailsConfigCategoryProduct,
	"product_stack": ConfigCategoryDetailsConfigCategoryProductStack,
	"environment":   ConfigCategoryDetailsConfigCategoryEnvironment,
	"patch_type":    ConfigCategoryDetailsConfigCategoryPatchType,
	"credential":    ConfigCategoryDetailsConfigCategoryCredential,
}

// GetConfigCategoryDetailsConfigCategoryEnumValues Enumerates the set of values for ConfigCategoryDetailsConfigCategoryEnum
func GetConfigCategoryDetailsConfigCategoryEnumValues() []ConfigCategoryDetailsConfigCategoryEnum {
	values := make([]ConfigCategoryDetailsConfigCategoryEnum, 0)
	for _, v := range mappingConfigCategoryDetailsConfigCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigCategoryDetailsConfigCategoryEnumStringValues Enumerates the set of values in String for ConfigCategoryDetailsConfigCategoryEnum
func GetConfigCategoryDetailsConfigCategoryEnumStringValues() []string {
	return []string{
		"PRODUCT",
		"PRODUCT_STACK",
		"ENVIRONMENT",
		"PATCH_TYPE",
		"CREDENTIAL",
	}
}

// GetMappingConfigCategoryDetailsConfigCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigCategoryDetailsConfigCategoryEnum(val string) (ConfigCategoryDetailsConfigCategoryEnum, bool) {
	enum, ok := mappingConfigCategoryDetailsConfigCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
