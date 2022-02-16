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

// HourlyFrequencyDetails Frequency details model to set hourly frequency
type HourlyFrequencyDetails struct {

	// This hold the repeatability aspect of a schedule. i.e. in a monhtly frequency, a task can be scheduled for every month, once in two months, once in tree months etc.
	Interval *int `mandatory:"false" json:"interval"`

	Time *Time `mandatory:"false" json:"time"`

	// the frequency of the schedule.
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
}

//GetFrequency returns Frequency
func (m HourlyFrequencyDetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m HourlyFrequencyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HourlyFrequencyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAbstractFrequencyDetailsFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetAbstractFrequencyDetailsFrequencyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HourlyFrequencyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHourlyFrequencyDetails HourlyFrequencyDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeHourlyFrequencyDetails
	}{
		"HOURLY",
		(MarshalTypeHourlyFrequencyDetails)(m),
	}

	return json.Marshal(&s)
}
