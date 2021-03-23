// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// Forecast Forcast configuration of usage/cost.
type Forecast struct {

	// forecast start time.
	TimeForecastStarted *common.SDKTime `mandatory:"true" json:"timeForecastStarted"`

	// forecast end time.
	TimeForecastEnded *common.SDKTime `mandatory:"true" json:"timeForecastEnded"`

	// BASIC uses ETS to project future usage/cost based on history data. The basis for projections will be a rolling set of equivalent historical days for which projection is being made.
	ForcastType ForecastForcastTypeEnum `mandatory:"false" json:"forcastType,omitempty"`
}

func (m Forecast) String() string {
	return common.PointerString(m)
}

// ForecastForcastTypeEnum Enum with underlying type: string
type ForecastForcastTypeEnum string

// Set of constants representing the allowable values for ForecastForcastTypeEnum
const (
	ForecastForcastTypeBasic ForecastForcastTypeEnum = "BASIC"
)

var mappingForecastForcastType = map[string]ForecastForcastTypeEnum{
	"BASIC": ForecastForcastTypeBasic,
}

// GetForecastForcastTypeEnumValues Enumerates the set of values for ForecastForcastTypeEnum
func GetForecastForcastTypeEnumValues() []ForecastForcastTypeEnum {
	values := make([]ForecastForcastTypeEnum, 0)
	for _, v := range mappingForecastForcastType {
		values = append(values, v)
	}
	return values
}
