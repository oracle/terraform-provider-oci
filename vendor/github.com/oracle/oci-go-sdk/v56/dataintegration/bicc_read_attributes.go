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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingBiccReadAttributesExtractStrategy = map[string]BiccReadAttributesExtractStrategyEnum{
	"FULL":        BiccReadAttributesExtractStrategyFull,
	"INCREMENTAL": BiccReadAttributesExtractStrategyIncremental,
}

// GetBiccReadAttributesExtractStrategyEnumValues Enumerates the set of values for BiccReadAttributesExtractStrategyEnum
func GetBiccReadAttributesExtractStrategyEnumValues() []BiccReadAttributesExtractStrategyEnum {
	values := make([]BiccReadAttributesExtractStrategyEnum, 0)
	for _, v := range mappingBiccReadAttributesExtractStrategy {
		values = append(values, v)
	}
	return values
}
