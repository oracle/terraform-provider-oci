// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BiccReadAttributes Properties to configure reading from BICC.
type BiccReadAttributes struct {

	// The fetch size for reading.
	FetchSize *int `mandatory:"false" json:"fetchSize"`

	ExternalStorage *ExternalStorage `mandatory:"false" json:"externalStorage"`

	// Date from where extract should start
	InitialExtractDate *common.SDKTime `mandatory:"false" json:"initialExtractDate"`

	// Date last extracted
	LastExtractDate *common.SDKTime `mandatory:"false" json:"lastExtractDate"`

	// Extraction Strategy - FULL|INCREMENTAL
	ExtractStrategy BiccReadAttributesExtractStrategyEnum `mandatory:"false" json:"extractStrategy,omitempty"`
}

func (m BiccReadAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BiccReadAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBiccReadAttributesExtractStrategyEnum(string(m.ExtractStrategy)); !ok && m.ExtractStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExtractStrategy: %s. Supported values are: %s.", m.ExtractStrategy, strings.Join(GetBiccReadAttributesExtractStrategyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BiccReadAttributes) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBiccReadAttributes BiccReadAttributes
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeBiccReadAttributes
	}{
		"BICC_READ_ATTRIBUTE",
		(MarshalTypeBiccReadAttributes)(m),
	}

	return json.Marshal(&s)
}

// BiccReadAttributesExtractStrategyEnum Enum with underlying type: string
type BiccReadAttributesExtractStrategyEnum string

// Set of constants representing the allowable values for BiccReadAttributesExtractStrategyEnum
const (
	BiccReadAttributesExtractStrategyFull        BiccReadAttributesExtractStrategyEnum = "FULL"
	BiccReadAttributesExtractStrategyIncremental BiccReadAttributesExtractStrategyEnum = "INCREMENTAL"
)

var mappingBiccReadAttributesExtractStrategyEnum = map[string]BiccReadAttributesExtractStrategyEnum{
	"FULL":        BiccReadAttributesExtractStrategyFull,
	"INCREMENTAL": BiccReadAttributesExtractStrategyIncremental,
}

// GetBiccReadAttributesExtractStrategyEnumValues Enumerates the set of values for BiccReadAttributesExtractStrategyEnum
func GetBiccReadAttributesExtractStrategyEnumValues() []BiccReadAttributesExtractStrategyEnum {
	values := make([]BiccReadAttributesExtractStrategyEnum, 0)
	for _, v := range mappingBiccReadAttributesExtractStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetBiccReadAttributesExtractStrategyEnumStringValues Enumerates the set of values in String for BiccReadAttributesExtractStrategyEnum
func GetBiccReadAttributesExtractStrategyEnumStringValues() []string {
	return []string{
		"FULL",
		"INCREMENTAL",
	}
}

// GetMappingBiccReadAttributesExtractStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBiccReadAttributesExtractStrategyEnum(val string) (BiccReadAttributesExtractStrategyEnum, bool) {
	mappingBiccReadAttributesExtractStrategyEnumIgnoreCase := make(map[string]BiccReadAttributesExtractStrategyEnum)
	for k, v := range mappingBiccReadAttributesExtractStrategyEnum {
		mappingBiccReadAttributesExtractStrategyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBiccReadAttributesExtractStrategyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
