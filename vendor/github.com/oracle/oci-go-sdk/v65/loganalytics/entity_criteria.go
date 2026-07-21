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

// EntityCriteria The criteria used to select entities.
type EntityCriteria struct {

	// Compartment Identifier OCID  (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Flag to include entities in all child compartments of the compartment Id specified in compartmentId.
	IncludeCompartmentsInSubtree *bool `mandatory:"false" json:"includeCompartmentsInSubtree"`

	// The entity name regular expression.
	EntityNameRegex *string `mandatory:"false" json:"entityNameRegex"`

	// The tags criteria to be applied to entities. Multiple criteria could be AND'ed or OR'ed.
	// Example: "(System = PROD) OR ((System=PERF) AND (Target.Env=PROD))"
	TagsCriteria *string `mandatory:"false" json:"tagsCriteria"`
}

func (m EntityCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
