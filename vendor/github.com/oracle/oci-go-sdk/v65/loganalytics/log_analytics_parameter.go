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

// LogAnalyticsParameter LogAnalyticsParameter
type LogAnalyticsParameter struct {

	// The default value of the parameter.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// The parameter description.
	Description *string `mandatory:"false" json:"description"`

	// A flag indicating whether or not the parameter is active.
	IsActive *bool `mandatory:"false" json:"isActive"`

	// The parameter name.
	Name *string `mandatory:"false" json:"name"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`
}

func (m LogAnalyticsParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
