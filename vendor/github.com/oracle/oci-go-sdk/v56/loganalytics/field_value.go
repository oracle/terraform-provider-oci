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

// FieldValue Field value representing and entry in a list-of-values field.
type FieldValue struct {

	// Display representation of the field value.
	DisplayValue *string `mandatory:"false" json:"displayValue"`

	// Internal representation of the field value.
	InternalValue *interface{} `mandatory:"false" json:"internalValue"`

	// Denotes if this list-of-values value has been marked as deleted.
	IsDeleted *bool `mandatory:"false" json:"isDeleted"`
}

func (m FieldValue) String() string {
	return common.PointerString(m)
}
