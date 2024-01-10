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

// EffectivePropertySummary A property and its effective value details.
type EffectivePropertySummary struct {

	// The property name.
	Name *string `mandatory:"true" json:"name"`

	// The effective value of the property. This is determined by considering the value set at the most effective level.
	Value *string `mandatory:"false" json:"value"`

	// The level from which the effective value was determined.
	EffectiveLevel *string `mandatory:"false" json:"effectiveLevel"`

	// A list of pattern level override values for the property.
	Patterns []PatternOverride `mandatory:"false" json:"patterns"`
}

func (m EffectivePropertySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EffectivePropertySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
