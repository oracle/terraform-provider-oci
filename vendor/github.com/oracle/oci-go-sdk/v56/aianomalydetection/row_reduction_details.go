// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RowReductionDetails Information regarding how/what row reduction methods will be applied. If this property is not present or is null, then it means row reduction is not applied.
type RowReductionDetails struct {

	// A boolean value to indicate if row reduction is applied
	IsReductionEnabled *bool `mandatory:"true" json:"isReductionEnabled"`

	// A percentage to reduce data size down to on top of original data
	ReductionPercentage *float64 `mandatory:"true" json:"reductionPercentage"`

	// Method for row reduction:
	//   * DELETE_ROW - delete rows with equal intervals
	//   * AVERAGE_ROW - average multiple rows to one row
	ReductionMethod RowReductionDetailsReductionMethodEnum `mandatory:"true" json:"reductionMethod"`
}

func (m RowReductionDetails) String() string {
	return common.PointerString(m)
}

// RowReductionDetailsReductionMethodEnum Enum with underlying type: string
type RowReductionDetailsReductionMethodEnum string

// Set of constants representing the allowable values for RowReductionDetailsReductionMethodEnum
const (
	RowReductionDetailsReductionMethodDeleteRow  RowReductionDetailsReductionMethodEnum = "DELETE_ROW"
	RowReductionDetailsReductionMethodAverageRow RowReductionDetailsReductionMethodEnum = "AVERAGE_ROW"
)

var mappingRowReductionDetailsReductionMethod = map[string]RowReductionDetailsReductionMethodEnum{
	"DELETE_ROW":  RowReductionDetailsReductionMethodDeleteRow,
	"AVERAGE_ROW": RowReductionDetailsReductionMethodAverageRow,
}

// GetRowReductionDetailsReductionMethodEnumValues Enumerates the set of values for RowReductionDetailsReductionMethodEnum
func GetRowReductionDetailsReductionMethodEnumValues() []RowReductionDetailsReductionMethodEnum {
	values := make([]RowReductionDetailsReductionMethodEnum, 0)
	for _, v := range mappingRowReductionDetailsReductionMethod {
		values = append(values, v)
	}
	return values
}
