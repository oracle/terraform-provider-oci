// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CatalogResultPayload Catalog result payload.
type CatalogResultPayload interface {

	// working directory
	GetWorkingDirectory() *string
}

type catalogresultpayload struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigResultType string  `json:"configResultType"`
}

// UnmarshalJSON unmarshals json
func (m *catalogresultpayload) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercatalogresultpayload catalogresultpayload
	s := struct {
		Model Unmarshalercatalogresultpayload
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigResultType = s.Model.ConfigResultType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *catalogresultpayload) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigResultType {
	case "GIT_RESULT_CONFIG":
		mm := CatalogGitResultConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEMPLATE_RESULT_CONFIG":
		mm := CatalogTemplateResultConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PAR_RESULT_CONFIG":
		mm := CatalogParResultConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CatalogResultPayload: %s.", m.ConfigResultType)
		return *m, nil
	}
}

// GetWorkingDirectory returns WorkingDirectory
func (m catalogresultpayload) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m catalogresultpayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m catalogresultpayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CatalogResultPayloadConfigResultTypeEnum Enum with underlying type: string
type CatalogResultPayloadConfigResultTypeEnum string

// Set of constants representing the allowable values for CatalogResultPayloadConfigResultTypeEnum
const (
	CatalogResultPayloadConfigResultTypeParResultConfig      CatalogResultPayloadConfigResultTypeEnum = "PAR_RESULT_CONFIG"
	CatalogResultPayloadConfigResultTypeTemplateResultConfig CatalogResultPayloadConfigResultTypeEnum = "TEMPLATE_RESULT_CONFIG"
	CatalogResultPayloadConfigResultTypeGitResultConfig      CatalogResultPayloadConfigResultTypeEnum = "GIT_RESULT_CONFIG"
)

var mappingCatalogResultPayloadConfigResultTypeEnum = map[string]CatalogResultPayloadConfigResultTypeEnum{
	"PAR_RESULT_CONFIG":      CatalogResultPayloadConfigResultTypeParResultConfig,
	"TEMPLATE_RESULT_CONFIG": CatalogResultPayloadConfigResultTypeTemplateResultConfig,
	"GIT_RESULT_CONFIG":      CatalogResultPayloadConfigResultTypeGitResultConfig,
}

var mappingCatalogResultPayloadConfigResultTypeEnumLowerCase = map[string]CatalogResultPayloadConfigResultTypeEnum{
	"par_result_config":      CatalogResultPayloadConfigResultTypeParResultConfig,
	"template_result_config": CatalogResultPayloadConfigResultTypeTemplateResultConfig,
	"git_result_config":      CatalogResultPayloadConfigResultTypeGitResultConfig,
}

// GetCatalogResultPayloadConfigResultTypeEnumValues Enumerates the set of values for CatalogResultPayloadConfigResultTypeEnum
func GetCatalogResultPayloadConfigResultTypeEnumValues() []CatalogResultPayloadConfigResultTypeEnum {
	values := make([]CatalogResultPayloadConfigResultTypeEnum, 0)
	for _, v := range mappingCatalogResultPayloadConfigResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogResultPayloadConfigResultTypeEnumStringValues Enumerates the set of values in String for CatalogResultPayloadConfigResultTypeEnum
func GetCatalogResultPayloadConfigResultTypeEnumStringValues() []string {
	return []string{
		"PAR_RESULT_CONFIG",
		"TEMPLATE_RESULT_CONFIG",
		"GIT_RESULT_CONFIG",
	}
}

// GetMappingCatalogResultPayloadConfigResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogResultPayloadConfigResultTypeEnum(val string) (CatalogResultPayloadConfigResultTypeEnum, bool) {
	enum, ok := mappingCatalogResultPayloadConfigResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
