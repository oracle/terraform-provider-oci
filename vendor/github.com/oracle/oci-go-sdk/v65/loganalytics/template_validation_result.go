// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TemplateValidationResult Template validation result
type TemplateValidationResult interface {

	// explanation of the validation status.
	GetStatusDescription() *string
}

type templatevalidationresult struct {
	JsonData          []byte
	StatusDescription *string `mandatory:"false" json:"statusDescription"`
	Status            string  `json:"status"`
}

// UnmarshalJSON unmarshals json
func (m *templatevalidationresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertemplatevalidationresult templatevalidationresult
	s := struct {
		Model Unmarshalertemplatevalidationresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StatusDescription = s.Model.StatusDescription
	m.Status = s.Model.Status

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *templatevalidationresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Status {
	case "SUCCESS":
		mm := TemplateSuccessValidationResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAILED":
		mm := TemplateFailedValidationResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TemplateValidationResult: %s.", m.Status)
		return *m, nil
	}
}

// GetStatusDescription returns StatusDescription
func (m templatevalidationresult) GetStatusDescription() *string {
	return m.StatusDescription
}

func (m templatevalidationresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m templatevalidationresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TemplateValidationResultStatusEnum Enum with underlying type: string
type TemplateValidationResultStatusEnum string

// Set of constants representing the allowable values for TemplateValidationResultStatusEnum
const (
	TemplateValidationResultStatusSuccess TemplateValidationResultStatusEnum = "SUCCESS"
	TemplateValidationResultStatusFailed  TemplateValidationResultStatusEnum = "FAILED"
)

var mappingTemplateValidationResultStatusEnum = map[string]TemplateValidationResultStatusEnum{
	"SUCCESS": TemplateValidationResultStatusSuccess,
	"FAILED":  TemplateValidationResultStatusFailed,
}

var mappingTemplateValidationResultStatusEnumLowerCase = map[string]TemplateValidationResultStatusEnum{
	"success": TemplateValidationResultStatusSuccess,
	"failed":  TemplateValidationResultStatusFailed,
}

// GetTemplateValidationResultStatusEnumValues Enumerates the set of values for TemplateValidationResultStatusEnum
func GetTemplateValidationResultStatusEnumValues() []TemplateValidationResultStatusEnum {
	values := make([]TemplateValidationResultStatusEnum, 0)
	for _, v := range mappingTemplateValidationResultStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTemplateValidationResultStatusEnumStringValues Enumerates the set of values in String for TemplateValidationResultStatusEnum
func GetTemplateValidationResultStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILED",
	}
}

// GetMappingTemplateValidationResultStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemplateValidationResultStatusEnum(val string) (TemplateValidationResultStatusEnum, bool) {
	enum, ok := mappingTemplateValidationResultStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
