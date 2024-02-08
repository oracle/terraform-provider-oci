// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Forecast Forecast configuration of usage/cost.
type Forecast struct {

	// The forecast end time.
	TimeForecastEnded *common.SDKTime `mandatory:"true" json:"timeForecastEnded"`

	// BASIC uses the exponential smoothing (ETS) model to project future usage/costs based on history data. The basis for projections is a periodic set of equivalent historical days for which the projection is being made.
	ForecastType ForecastForecastTypeEnum `mandatory:"false" json:"forecastType,omitempty"`

	// The forecast start time. Defaults to UTC-1 if not specified.
	TimeForecastStarted *common.SDKTime `mandatory:"false" json:"timeForecastStarted"`
}

func (m Forecast) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Forecast) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingForecastForecastTypeEnum(string(m.ForecastType)); !ok && m.ForecastType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ForecastType: %s. Supported values are: %s.", m.ForecastType, strings.Join(GetForecastForecastTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ForecastForecastTypeEnum Enum with underlying type: string
type ForecastForecastTypeEnum string

// Set of constants representing the allowable values for ForecastForecastTypeEnum
const (
	ForecastForecastTypeBasic ForecastForecastTypeEnum = "BASIC"
)

var mappingForecastForecastTypeEnum = map[string]ForecastForecastTypeEnum{
	"BASIC": ForecastForecastTypeBasic,
}

var mappingForecastForecastTypeEnumLowerCase = map[string]ForecastForecastTypeEnum{
	"basic": ForecastForecastTypeBasic,
}

// GetForecastForecastTypeEnumValues Enumerates the set of values for ForecastForecastTypeEnum
func GetForecastForecastTypeEnumValues() []ForecastForecastTypeEnum {
	values := make([]ForecastForecastTypeEnum, 0)
	for _, v := range mappingForecastForecastTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetForecastForecastTypeEnumStringValues Enumerates the set of values in String for ForecastForecastTypeEnum
func GetForecastForecastTypeEnumStringValues() []string {
	return []string{
		"BASIC",
	}
}

// GetMappingForecastForecastTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingForecastForecastTypeEnum(val string) (ForecastForecastTypeEnum, bool) {
	enum, ok := mappingForecastForecastTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
