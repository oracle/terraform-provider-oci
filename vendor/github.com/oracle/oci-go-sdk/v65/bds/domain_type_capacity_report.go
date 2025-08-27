// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DomainTypeCapacityReport Detailed information of the capacity under a domain.
type DomainTypeCapacityReport interface {
}

type domaintypecapacityreport struct {
	JsonData   []byte
	DomainType string `json:"domainType"`
}

// UnmarshalJSON unmarshals json
func (m *domaintypecapacityreport) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdomaintypecapacityreport domaintypecapacityreport
	s := struct {
		Model Unmarshalerdomaintypecapacityreport
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DomainType = s.Model.DomainType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *domaintypecapacityreport) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DomainType {
	case "SINGLE_AD":
		mm := SingleAdCapacityReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MULTI_AD":
		mm := MultiAdCapacityReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DomainTypeCapacityReport: %s.", m.DomainType)
		return *m, nil
	}
}

func (m domaintypecapacityreport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m domaintypecapacityreport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DomainTypeCapacityReportDomainTypeEnum Enum with underlying type: string
type DomainTypeCapacityReportDomainTypeEnum string

// Set of constants representing the allowable values for DomainTypeCapacityReportDomainTypeEnum
const (
	DomainTypeCapacityReportDomainTypeMultiAd  DomainTypeCapacityReportDomainTypeEnum = "MULTI_AD"
	DomainTypeCapacityReportDomainTypeSingleAd DomainTypeCapacityReportDomainTypeEnum = "SINGLE_AD"
)

var mappingDomainTypeCapacityReportDomainTypeEnum = map[string]DomainTypeCapacityReportDomainTypeEnum{
	"MULTI_AD":  DomainTypeCapacityReportDomainTypeMultiAd,
	"SINGLE_AD": DomainTypeCapacityReportDomainTypeSingleAd,
}

var mappingDomainTypeCapacityReportDomainTypeEnumLowerCase = map[string]DomainTypeCapacityReportDomainTypeEnum{
	"multi_ad":  DomainTypeCapacityReportDomainTypeMultiAd,
	"single_ad": DomainTypeCapacityReportDomainTypeSingleAd,
}

// GetDomainTypeCapacityReportDomainTypeEnumValues Enumerates the set of values for DomainTypeCapacityReportDomainTypeEnum
func GetDomainTypeCapacityReportDomainTypeEnumValues() []DomainTypeCapacityReportDomainTypeEnum {
	values := make([]DomainTypeCapacityReportDomainTypeEnum, 0)
	for _, v := range mappingDomainTypeCapacityReportDomainTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainTypeCapacityReportDomainTypeEnumStringValues Enumerates the set of values in String for DomainTypeCapacityReportDomainTypeEnum
func GetDomainTypeCapacityReportDomainTypeEnumStringValues() []string {
	return []string{
		"MULTI_AD",
		"SINGLE_AD",
	}
}

// GetMappingDomainTypeCapacityReportDomainTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainTypeCapacityReportDomainTypeEnum(val string) (DomainTypeCapacityReportDomainTypeEnum, bool) {
	enum, ok := mappingDomainTypeCapacityReportDomainTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
