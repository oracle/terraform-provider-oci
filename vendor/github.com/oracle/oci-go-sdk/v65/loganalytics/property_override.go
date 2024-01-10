// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PropertyOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
