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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RowReductionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRowReductionDetailsReductionMethodEnum(string(m.ReductionMethod)); !ok && m.ReductionMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReductionMethod: %s. Supported values are: %s.", m.ReductionMethod, strings.Join(GetRowReductionDetailsReductionMethodEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RowReductionDetailsReductionMethodEnum Enum with underlying type: string
type RowReductionDetailsReductionMethodEnum string

// Set of constants representing the allowable values for RowReductionDetailsReductionMethodEnum
const (
	RowReductionDetailsReductionMethodDeleteRow  RowReductionDetailsReductionMethodEnum = "DELETE_ROW"
	RowReductionDetailsReductionMethodAverageRow RowReductionDetailsReductionMethodEnum = "AVERAGE_ROW"
)

var mappingRowReductionDetailsReductionMethodEnum = map[string]RowReductionDetailsReductionMethodEnum{
	"DELETE_ROW":  RowReductionDetailsReductionMethodDeleteRow,
	"AVERAGE_ROW": RowReductionDetailsReductionMethodAverageRow,
}

// GetRowReductionDetailsReductionMethodEnumValues Enumerates the set of values for RowReductionDetailsReductionMethodEnum
func GetRowReductionDetailsReductionMethodEnumValues() []RowReductionDetailsReductionMethodEnum {
	values := make([]RowReductionDetailsReductionMethodEnum, 0)
	for _, v := range mappingRowReductionDetailsReductionMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetRowReductionDetailsReductionMethodEnumStringValues Enumerates the set of values in String for RowReductionDetailsReductionMethodEnum
func GetRowReductionDetailsReductionMethodEnumStringValues() []string {
	return []string{
		"DELETE_ROW",
		"AVERAGE_ROW",
	}
}

// GetMappingRowReductionDetailsReductionMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRowReductionDetailsReductionMethodEnum(val string) (RowReductionDetailsReductionMethodEnum, bool) {
	mappingRowReductionDetailsReductionMethodEnumIgnoreCase := make(map[string]RowReductionDetailsReductionMethodEnum)
	for k, v := range mappingRowReductionDetailsReductionMethodEnum {
		mappingRowReductionDetailsReductionMethodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRowReductionDetailsReductionMethodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
