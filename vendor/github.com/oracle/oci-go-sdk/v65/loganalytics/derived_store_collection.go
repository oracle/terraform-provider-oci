// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DerivedStoreCollection Derived store collection which will post results as new log records.
type DerivedStoreCollection struct {

	// Log Analytics log group OCID to associate with the collection.
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// Name of the Log Analytics derived source to use for processing.
	LogSourceName *string `mandatory:"true" json:"logSourceName"`

	// Selected fields for the collection.
	Fields []DerivedStoreField `mandatory:"true" json:"fields"`

	// Name of the collection.
	Name *string `mandatory:"false" json:"name"`

	// The log set to be associated with the collection.
	LogSet *string `mandatory:"false" json:"logSet"`

	// Output table in the saved search query.
	QueryTableName *string `mandatory:"false" json:"queryTableName"`

	// Use as timestamp for each collected record.
	QueryTimeField *string `mandatory:"false" json:"queryTimeField"`
}

func (m DerivedStoreCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DerivedStoreCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
