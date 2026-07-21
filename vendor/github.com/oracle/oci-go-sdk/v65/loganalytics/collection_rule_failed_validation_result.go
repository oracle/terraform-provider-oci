// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CollectionRuleFailedValidationResult Result of a failed collection rule validation action.
type CollectionRuleFailedValidationResult struct {

	// explanation of the validation status.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// The violations causing validation failure.
	Violations []Violation `mandatory:"false" json:"violations"`
}

// GetStatusDescription returns StatusDescription
func (m CollectionRuleFailedValidationResult) GetStatusDescription() *string {
	return m.StatusDescription
}

func (m CollectionRuleFailedValidationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectionRuleFailedValidationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CollectionRuleFailedValidationResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCollectionRuleFailedValidationResult CollectionRuleFailedValidationResult
	s := struct {
		DiscriminatorParam string `json:"status"`
		MarshalTypeCollectionRuleFailedValidationResult
	}{
		"FAILED",
		(MarshalTypeCollectionRuleFailedValidationResult)(m),
	}

	return json.Marshal(&s)
}
