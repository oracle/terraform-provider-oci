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

// CollectionRuleSuccessValidationResult Result of a successful collection rule validation action.
type CollectionRuleSuccessValidationResult struct {

	// explanation of the validation status.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	Template *CollectionRule `mandatory:"false" json:"template"`
}

// GetStatusDescription returns StatusDescription
func (m CollectionRuleSuccessValidationResult) GetStatusDescription() *string {
	return m.StatusDescription
}

func (m CollectionRuleSuccessValidationResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectionRuleSuccessValidationResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CollectionRuleSuccessValidationResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCollectionRuleSuccessValidationResult CollectionRuleSuccessValidationResult
	s := struct {
		DiscriminatorParam string `json:"status"`
		MarshalTypeCollectionRuleSuccessValidationResult
	}{
		"SUCCESS",
		(MarshalTypeCollectionRuleSuccessValidationResult)(m),
	}

	return json.Marshal(&s)
}
