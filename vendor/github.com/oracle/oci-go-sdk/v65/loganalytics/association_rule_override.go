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

// AssociationRuleOverride Criteria and parameters that apply to specific entity type or log source(s) or both.
type AssociationRuleOverride struct {

	// The list of log source names to which this override applies to.
	Sources []string `mandatory:"false" json:"sources"`

	// The entity type to which this override applies to.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The criteria used in selecting entities to which this override applies to.
	EntityCriteria *EntityCriteria `mandatory:"false" json:"entityCriteria"`

	// The parameters to be used in associations that are created by applying this override.
	Parameters *AssociationRuleOverrideParameters `mandatory:"false" json:"parameters"`
}

func (m AssociationRuleOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociationRuleOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
