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

// SavedCustomTable The custom table for Cost Analysis UI rendering.
type SavedCustomTable struct {

	// The name of the custom table.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The row groupBy key list.
	// example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "tenantId", "tenantName"]`
	RowGroupBy []string `mandatory:"false" json:"rowGroupBy"`

	// The column groupBy key list.
	// example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "tenantId", "tenantName"]`
	ColumnGroupBy []string `mandatory:"false" json:"columnGroupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only one tag in the list is supported.
	// For example:
	//   `[{"namespace":"oracle", "key":"createdBy"]`
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	// The compartment depth level.
	CompartmentDepth *float32 `mandatory:"false" json:"compartmentDepth"`

	// The version of the custom table.
	Version *float32 `mandatory:"false" json:"version"`
}

func (m SavedCustomTable) String() string {
	return common.PointerString(m)
}
