// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by the Cost Analysis and Carbon Emissions Analysis tools in the Console. See Cost Analysis Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm) and Using the Usage API (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CostAnalysisUi) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCostAnalysisUiGraphEnum(string(m.Graph)); !ok && m.Graph != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Graph: %s. Supported values are: %s.", m.Graph, strings.Join(GetCostAnalysisUiGraphEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CostAnalysisUiGraphEnum Enum with underlying type: string
type CostAnalysisUiGraphEnum string

// Set of constants representing the allowable values for CostAnalysisUiGraphEnum
const (
	CostAnalysisUiGraphBars         CostAnalysisUiGraphEnum = "BARS"
	CostAnalysisUiGraphLines        CostAnalysisUiGraphEnum = "LINES"
	CostAnalysisUiGraphStackedLines CostAnalysisUiGraphEnum = "STACKED_LINES"
)

var mappingCostAnalysisUiGraphEnum = map[string]CostAnalysisUiGraphEnum{
	"BARS":          CostAnalysisUiGraphBars,
	"LINES":         CostAnalysisUiGraphLines,
	"STACKED_LINES": CostAnalysisUiGraphStackedLines,
}

var mappingCostAnalysisUiGraphEnumLowerCase = map[string]CostAnalysisUiGraphEnum{
	"bars":          CostAnalysisUiGraphBars,
	"lines":         CostAnalysisUiGraphLines,
	"stacked_lines": CostAnalysisUiGraphStackedLines,
}

// GetCostAnalysisUiGraphEnumValues Enumerates the set of values for CostAnalysisUiGraphEnum
func GetCostAnalysisUiGraphEnumValues() []CostAnalysisUiGraphEnum {
	values := make([]CostAnalysisUiGraphEnum, 0)
	for _, v := range mappingCostAnalysisUiGraphEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAnalysisUiGraphEnumStringValues Enumerates the set of values in String for CostAnalysisUiGraphEnum
func GetCostAnalysisUiGraphEnumStringValues() []string {
	return []string{
		"BARS",
		"LINES",
		"STACKED_LINES",
	}
}

// GetMappingCostAnalysisUiGraphEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAnalysisUiGraphEnum(val string) (CostAnalysisUiGraphEnum, bool) {
	enum, ok := mappingCostAnalysisUiGraphEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
