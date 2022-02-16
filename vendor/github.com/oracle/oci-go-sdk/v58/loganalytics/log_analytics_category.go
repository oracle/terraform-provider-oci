// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// LogAnalyticsCategory A category into which resources can be placed.
type LogAnalyticsCategory struct {

	// The unique name that identifies the category.
	Name *string `mandatory:"false" json:"name"`

	// The category description.
	Description *string `mandatory:"false" json:"description"`

	// The category display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
	Type *string `mandatory:"false" json:"type"`

	// The system flag. A value of false denotes a user-created
	// category. A value of true denotes an Oracle-defined category.
	IsSystem *bool `mandatory:"false" json:"isSystem"`
}

func (m LogAnalyticsCategory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsCategory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
