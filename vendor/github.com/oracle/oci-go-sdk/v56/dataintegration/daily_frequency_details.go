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

// DailyFrequencyDetails Frequency details model to set daily frequency
type DailyFrequencyDetails struct {

	// This hold the repeatability aspect of a schedule. i.e. in a monhtly frequency, a task can be scheduled for every month, once in two months, once in tree months etc.
	Interval *int `mandatory:"false" json:"interval"`

	Time *Time `mandatory:"false" json:"time"`

	// the frequency of the schedule.
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
}

//GetFrequency returns Frequency
func (m DailyFrequencyDetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m DailyFrequencyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DailyFrequencyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDailyFrequencyDetails DailyFrequencyDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDailyFrequencyDetails
	}{
		"DAILY",
		(MarshalTypeDailyFrequencyDetails)(m),
	}

	return json.Marshal(&s)
}
