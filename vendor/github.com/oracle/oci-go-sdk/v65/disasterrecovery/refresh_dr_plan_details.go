// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RefreshDrPlanDetails The details for refreshing a DR plan.
type RefreshDrPlanDetails interface {
}

type refreshdrplandetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *refreshdrplandetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrefreshdrplandetails refreshdrplandetails
	s := struct {
		Model Unmarshalerrefreshdrplandetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *refreshdrplandetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := RefreshDrPlanDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RefreshDrPlanDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m refreshdrplandetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m refreshdrplandetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RefreshDrPlanDetailsTypeEnum Enum with underlying type: string
type RefreshDrPlanDetailsTypeEnum string

// Set of constants representing the allowable values for RefreshDrPlanDetailsTypeEnum
const (
	RefreshDrPlanDetailsTypeDefault RefreshDrPlanDetailsTypeEnum = "DEFAULT"
)

var mappingRefreshDrPlanDetailsTypeEnum = map[string]RefreshDrPlanDetailsTypeEnum{
	"DEFAULT": RefreshDrPlanDetailsTypeDefault,
}

var mappingRefreshDrPlanDetailsTypeEnumLowerCase = map[string]RefreshDrPlanDetailsTypeEnum{
	"default": RefreshDrPlanDetailsTypeDefault,
}

// GetRefreshDrPlanDetailsTypeEnumValues Enumerates the set of values for RefreshDrPlanDetailsTypeEnum
func GetRefreshDrPlanDetailsTypeEnumValues() []RefreshDrPlanDetailsTypeEnum {
	values := make([]RefreshDrPlanDetailsTypeEnum, 0)
	for _, v := range mappingRefreshDrPlanDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRefreshDrPlanDetailsTypeEnumStringValues Enumerates the set of values in String for RefreshDrPlanDetailsTypeEnum
func GetRefreshDrPlanDetailsTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingRefreshDrPlanDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRefreshDrPlanDetailsTypeEnum(val string) (RefreshDrPlanDetailsTypeEnum, bool) {
	enum, ok := mappingRefreshDrPlanDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
