// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VerifyDrPlanDetails The details for verifying a DR plan.
type VerifyDrPlanDetails interface {
}

type verifydrplandetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *verifydrplandetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerverifydrplandetails verifydrplandetails
	s := struct {
		Model Unmarshalerverifydrplandetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *verifydrplandetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := VerifyDrPlanDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VerifyDrPlanDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m verifydrplandetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m verifydrplandetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VerifyDrPlanDetailsTypeEnum Enum with underlying type: string
type VerifyDrPlanDetailsTypeEnum string

// Set of constants representing the allowable values for VerifyDrPlanDetailsTypeEnum
const (
	VerifyDrPlanDetailsTypeDefault VerifyDrPlanDetailsTypeEnum = "DEFAULT"
)

var mappingVerifyDrPlanDetailsTypeEnum = map[string]VerifyDrPlanDetailsTypeEnum{
	"DEFAULT": VerifyDrPlanDetailsTypeDefault,
}

var mappingVerifyDrPlanDetailsTypeEnumLowerCase = map[string]VerifyDrPlanDetailsTypeEnum{
	"default": VerifyDrPlanDetailsTypeDefault,
}

// GetVerifyDrPlanDetailsTypeEnumValues Enumerates the set of values for VerifyDrPlanDetailsTypeEnum
func GetVerifyDrPlanDetailsTypeEnumValues() []VerifyDrPlanDetailsTypeEnum {
	values := make([]VerifyDrPlanDetailsTypeEnum, 0)
	for _, v := range mappingVerifyDrPlanDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVerifyDrPlanDetailsTypeEnumStringValues Enumerates the set of values in String for VerifyDrPlanDetailsTypeEnum
func GetVerifyDrPlanDetailsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingVerifyDrPlanDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVerifyDrPlanDetailsTypeEnum(val string) (VerifyDrPlanDetailsTypeEnum, bool) {
	enum, ok := mappingVerifyDrPlanDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
