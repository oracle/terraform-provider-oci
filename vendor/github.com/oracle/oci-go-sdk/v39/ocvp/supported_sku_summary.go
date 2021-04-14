// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// SupportedSkuSummary A specific SKU. HOUR, MONTH, ONE_YEAR and THREE_YEARS supported by the Oracle Cloud VMware Solution.
type SupportedSkuSummary struct {

	// name of SKU
	Name SupportedSkuSummaryNameEnum `mandatory:"true" json:"name"`
}

func (m SupportedSkuSummary) String() string {
	return common.PointerString(m)
}

// SupportedSkuSummaryNameEnum Enum with underlying type: string
type SupportedSkuSummaryNameEnum string

// Set of constants representing the allowable values for SupportedSkuSummaryNameEnum
const (
	SupportedSkuSummaryNameHour       SupportedSkuSummaryNameEnum = "HOUR"
	SupportedSkuSummaryNameMonth      SupportedSkuSummaryNameEnum = "MONTH"
	SupportedSkuSummaryNameOneYear    SupportedSkuSummaryNameEnum = "ONE_YEAR"
	SupportedSkuSummaryNameThreeYears SupportedSkuSummaryNameEnum = "THREE_YEARS"
)

var mappingSupportedSkuSummaryName = map[string]SupportedSkuSummaryNameEnum{
	"HOUR":        SupportedSkuSummaryNameHour,
	"MONTH":       SupportedSkuSummaryNameMonth,
	"ONE_YEAR":    SupportedSkuSummaryNameOneYear,
	"THREE_YEARS": SupportedSkuSummaryNameThreeYears,
}

// GetSupportedSkuSummaryNameEnumValues Enumerates the set of values for SupportedSkuSummaryNameEnum
func GetSupportedSkuSummaryNameEnumValues() []SupportedSkuSummaryNameEnum {
	values := make([]SupportedSkuSummaryNameEnum, 0)
	for _, v := range mappingSupportedSkuSummaryName {
		values = append(values, v)
	}
	return values
}
