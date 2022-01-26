// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResultColumn Querylanguage result column.
type ResultColumn struct {

	// Internal identifier for the column.
	InternalName *string `mandatory:"false" json:"internalName"`

	// Display name - will be alias if result column is renamed by queryString.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Field denoting column data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

func (m ResultColumn) String() string {
	return common.PointerString(m)
}
