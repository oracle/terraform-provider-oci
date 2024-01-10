// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddmDbRecommendationsTimeSeriesSummary ADDM recommendation
type AddmDbRecommendationsTimeSeriesSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Unique ADDM task id
	TaskId *int `mandatory:"true" json:"taskId"`

	// ADDM task name
	TaskName *string `mandatory:"true" json:"taskName"`

	// Timestamp when recommendation was generated
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Start Timestamp of snapshot
	TimeAnalysisStarted *common.SDKTime `mandatory:"false" json:"timeAnalysisStarted"`

	// End Timestamp of snapshot
	TimeAnalysisEnded *common.SDKTime `mandatory:"false" json:"timeAnalysisEnded"`

	// Type of recommendation
	Type *string `mandatory:"false" json:"type"`

	// DB time in seconds for the snapshot
	AnalysisDbTimeInSecs *float64 `mandatory:"false" json:"analysisDbTimeInSecs"`

	// DB avg active sessions for the snapshot
	AnalysisAvgActiveSessions *float64 `mandatory:"false" json:"analysisAvgActiveSessions"`

	// Maximum estimated benefit in terms of percentage of total activity
	MaxBenefitPercent *float64 `mandatory:"false" json:"maxBenefitPercent"`

	// Maximum estimated benefit in terms of seconds
	MaxBenefitDbTimeInSecs *float64 `mandatory:"false" json:"maxBenefitDbTimeInSecs"`

	// Maximum estimated benefit in terms of average active sessions
	MaxBenefitAvgActiveSessions *float64 `mandatory:"false" json:"maxBenefitAvgActiveSessions"`

	RelatedObject RelatedObjectTypeDetails `mandatory:"false" json:"relatedObject"`
}

func (m AddmDbRecommendationsTimeSeriesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbRecommendationsTimeSeriesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AddmDbRecommendationsTimeSeriesSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeAnalysisStarted         *common.SDKTime          `json:"timeAnalysisStarted"`
		TimeAnalysisEnded           *common.SDKTime          `json:"timeAnalysisEnded"`
		Type                        *string                  `json:"type"`
		AnalysisDbTimeInSecs        *float64                 `json:"analysisDbTimeInSecs"`
		AnalysisAvgActiveSessions   *float64                 `json:"analysisAvgActiveSessions"`
		MaxBenefitPercent           *float64                 `json:"maxBenefitPercent"`
		MaxBenefitDbTimeInSecs      *float64                 `json:"maxBenefitDbTimeInSecs"`
		MaxBenefitAvgActiveSessions *float64                 `json:"maxBenefitAvgActiveSessions"`
		RelatedObject               relatedobjecttypedetails `json:"relatedObject"`
		Id                          *string                  `json:"id"`
		TaskId                      *int                     `json:"taskId"`
		TaskName                    *string                  `json:"taskName"`
		Timestamp                   *common.SDKTime          `json:"timestamp"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeAnalysisStarted = model.TimeAnalysisStarted

	m.TimeAnalysisEnded = model.TimeAnalysisEnded

	m.Type = model.Type

	m.AnalysisDbTimeInSecs = model.AnalysisDbTimeInSecs

	m.AnalysisAvgActiveSessions = model.AnalysisAvgActiveSessions

	m.MaxBenefitPercent = model.MaxBenefitPercent

	m.MaxBenefitDbTimeInSecs = model.MaxBenefitDbTimeInSecs

	m.MaxBenefitAvgActiveSessions = model.MaxBenefitAvgActiveSessions

	nn, e = model.RelatedObject.UnmarshalPolymorphicJSON(model.RelatedObject.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RelatedObject = nn.(RelatedObjectTypeDetails)
	} else {
		m.RelatedObject = nil
	}

	m.Id = model.Id

	m.TaskId = model.TaskId

	m.TaskName = model.TaskName

	m.Timestamp = model.Timestamp

	return
}
