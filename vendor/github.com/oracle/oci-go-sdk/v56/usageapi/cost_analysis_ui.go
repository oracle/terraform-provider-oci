// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CostAnalysisUi The common fields for Cost Analysis UI rendering.
type CostAnalysisUi struct {

	// The graph type.
	Graph CostAnalysisUiGraphEnum `mandatory:"false" json:"graph,omitempty"`

	// A cumulative graph.
	IsCumulativeGraph *bool `mandatory:"false" json:"isCumulativeGraph"`
}

func (m CostAnalysisUi) String() string {
	return common.PointerString(m)
}

// CostAnalysisUiGraphEnum Enum with underlying type: string
type CostAnalysisUiGraphEnum string

// Set of constants representing the allowable values for CostAnalysisUiGraphEnum
const (
	CostAnalysisUiGraphBars         CostAnalysisUiGraphEnum = "BARS"
	CostAnalysisUiGraphLines        CostAnalysisUiGraphEnum = "LINES"
	CostAnalysisUiGraphStackedLines CostAnalysisUiGraphEnum = "STACKED_LINES"
)

var mappingCostAnalysisUiGraph = map[string]CostAnalysisUiGraphEnum{
	"BARS":          CostAnalysisUiGraphBars,
	"LINES":         CostAnalysisUiGraphLines,
	"STACKED_LINES": CostAnalysisUiGraphStackedLines,
}

// GetCostAnalysisUiGraphEnumValues Enumerates the set of values for CostAnalysisUiGraphEnum
func GetCostAnalysisUiGraphEnumValues() []CostAnalysisUiGraphEnum {
	values := make([]CostAnalysisUiGraphEnum, 0)
	for _, v := range mappingCostAnalysisUiGraph {
		values = append(values, v)
	}
	return values
}
