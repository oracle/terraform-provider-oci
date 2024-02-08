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

// AddmDbRecommendationAggregation Summarizes a specific ADDM recommendation
type AddmDbRecommendationAggregation struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Recommendation message
	Message *string `mandatory:"true" json:"message"`

	// Type of recommendation
	Type *string `mandatory:"false" json:"type"`

	// Indicates implementation of the recommended action requires a database restart in order for it
	// to take effect. Possible values "Y", "N" and null.
	RequiresDbRestart *string `mandatory:"false" json:"requiresDbRestart"`

	// Actions that can be performed to implement the recommendation (such as 'ALTER PARAMETER',
	// 'RUN SQL TUNING ADVISOR')
	ImplementActions []string `mandatory:"false" json:"implementActions"`

	// Recommendation message
	Rationale *string `mandatory:"false" json:"rationale"`

	// Maximum estimated benefit in terms of percentage of total activity
	MaxBenefitPercent *float64 `mandatory:"false" json:"maxBenefitPercent"`

	// Overall estimated benefit in terms of percentage of total activity
	OverallBenefitPercent *float64 `mandatory:"false" json:"overallBenefitPercent"`

	// Maximum estimated benefit in terms of average active sessions
	MaxBenefitAvgActiveSessions *float64 `mandatory:"false" json:"maxBenefitAvgActiveSessions"`

	// Number of occurrences for this recommendation
	FrequencyCount *int `mandatory:"false" json:"frequencyCount"`

	RelatedObject RelatedObjectTypeDetails `mandatory:"false" json:"relatedObject"`
}

func (m AddmDbRecommendationAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbRecommendationAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AddmDbRecommendationAggregation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Type                        *string                  `json:"type"`
		RequiresDbRestart           *string                  `json:"requiresDbRestart"`
		ImplementActions            []string                 `json:"implementActions"`
		Rationale                   *string                  `json:"rationale"`
		MaxBenefitPercent           *float64                 `json:"maxBenefitPercent"`
		OverallBenefitPercent       *float64                 `json:"overallBenefitPercent"`
		MaxBenefitAvgActiveSessions *float64                 `json:"maxBenefitAvgActiveSessions"`
		FrequencyCount              *int                     `json:"frequencyCount"`
		RelatedObject               relatedobjecttypedetails `json:"relatedObject"`
		Id                          *string                  `json:"id"`
		Message                     *string                  `json:"message"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Type = model.Type

	m.RequiresDbRestart = model.RequiresDbRestart

	m.ImplementActions = make([]string, len(model.ImplementActions))
	copy(m.ImplementActions, model.ImplementActions)
	m.Rationale = model.Rationale

	m.MaxBenefitPercent = model.MaxBenefitPercent

	m.OverallBenefitPercent = model.OverallBenefitPercent

	m.MaxBenefitAvgActiveSessions = model.MaxBenefitAvgActiveSessions

	m.FrequencyCount = model.FrequencyCount

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

	m.Message = model.Message

	return
}
