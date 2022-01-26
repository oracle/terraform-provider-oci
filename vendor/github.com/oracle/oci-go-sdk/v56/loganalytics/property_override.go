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

// PropertyOverride Property overrides at the scope of objects.
// For example, if you want to use logSourceName as 'xyz' for all objects that conatins string 'abc/' then
// define matchType as 'contains', matchValue as 'abc/', propertyName as 'logSourceName' and propertyValue as 'xyz'.
type PropertyOverride struct {

	// Match Type. Accepted values are: contains.
	MatchType *string `mandatory:"false" json:"matchType"`

	// Match Value.
	MatchValue *string `mandatory:"false" json:"matchValue"`

	// Property to override. Accepted values are: logSourceName, charEncoding.
	PropertyName *string `mandatory:"false" json:"propertyName"`

	// Value of the property.
	PropertyValue *string `mandatory:"false" json:"propertyValue"`
}

func (m PropertyOverride) String() string {
	return common.PointerString(m)
}
