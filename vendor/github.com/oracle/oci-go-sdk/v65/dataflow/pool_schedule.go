// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PoolSchedule Definition of when pool auto start or stop for a given day of a week.
type PoolSchedule struct {

	// Day of the week SUN-SAT
	DayOfWeek DayOfWeekEnum `mandatory:"false" json:"dayOfWeek,omitempty"`

	// Hour of the day to start or stop pool.
	StartTime *int `mandatory:"false" json:"startTime"`

	// Hour of the day to stop the pool.
	StopTime *int `mandatory:"false" json:"stopTime"`
}

func (m PoolSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PoolSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetDayOfWeekEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
