// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DvpStackDetails Details of a Dedicated Vantage Point (DVP) stack in Resource Manager.
type DvpStackDetails interface {

	// Version of the dedicated vantage point.
	GetDvpVersion() *string
}

type dvpstackdetails struct {
	JsonData     []byte
	DvpVersion   *string `mandatory:"true" json:"dvpVersion"`
	DvpStackType string  `json:"dvpStackType"`
}

// UnmarshalJSON unmarshals json
func (m *dvpstackdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdvpstackdetails dvpstackdetails
	s := struct {
		Model Unmarshalerdvpstackdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DvpVersion = s.Model.DvpVersion
	m.DvpStackType = s.Model.DvpStackType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dvpstackdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DvpStackType {
	case "ORACLE_RM_STACK":
		mm := OracleRmStack{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DvpStackDetails: %s.", m.DvpStackType)
		return *m, nil
	}
}

// GetDvpVersion returns DvpVersion
func (m dvpstackdetails) GetDvpVersion() *string {
	return m.DvpVersion
}

func (m dvpstackdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dvpstackdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DvpStackDetailsDvpStackTypeEnum Enum with underlying type: string
type DvpStackDetailsDvpStackTypeEnum string

// Set of constants representing the allowable values for DvpStackDetailsDvpStackTypeEnum
const (
	DvpStackDetailsDvpStackTypeOracleRmStack DvpStackDetailsDvpStackTypeEnum = "ORACLE_RM_STACK"
)

var mappingDvpStackDetailsDvpStackTypeEnum = map[string]DvpStackDetailsDvpStackTypeEnum{
	"ORACLE_RM_STACK": DvpStackDetailsDvpStackTypeOracleRmStack,
}

var mappingDvpStackDetailsDvpStackTypeEnumLowerCase = map[string]DvpStackDetailsDvpStackTypeEnum{
	"oracle_rm_stack": DvpStackDetailsDvpStackTypeOracleRmStack,
}

// GetDvpStackDetailsDvpStackTypeEnumValues Enumerates the set of values for DvpStackDetailsDvpStackTypeEnum
func GetDvpStackDetailsDvpStackTypeEnumValues() []DvpStackDetailsDvpStackTypeEnum {
	values := make([]DvpStackDetailsDvpStackTypeEnum, 0)
	for _, v := range mappingDvpStackDetailsDvpStackTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDvpStackDetailsDvpStackTypeEnumStringValues Enumerates the set of values in String for DvpStackDetailsDvpStackTypeEnum
func GetDvpStackDetailsDvpStackTypeEnumStringValues() []string {
	return []string{
		"ORACLE_RM_STACK",
	}
}

// GetMappingDvpStackDetailsDvpStackTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDvpStackDetailsDvpStackTypeEnum(val string) (DvpStackDetailsDvpStackTypeEnum, bool) {
	enum, ok := mappingDvpStackDetailsDvpStackTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
