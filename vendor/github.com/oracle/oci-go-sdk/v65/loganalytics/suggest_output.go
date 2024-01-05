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

// SuggestOutput Typeahead results
type SuggestOutput struct {

	// Character position suggestion should be placed in queryString provided as input.
	Position *int `mandatory:"true" json:"position"`

	// Context specific list of querylanguage commands if input is seeking command suggestions.
	Commands []string `mandatory:"false" json:"commands"`

	// Context specific list of querylanguage fields / columns if input is seeking field / column suggestions.
	Fields []string `mandatory:"false" json:"fields"`

	// Context specific list of field values if input is seeking field value suggestions.
	FieldValues []string `mandatory:"false" json:"fieldValues"`

	// Context specific list of terms / phrases if input is seeking terms / phrase suggestions.
	Terms []string `mandatory:"false" json:"terms"`

	// Context specific list of querylanguage command options if input is seeking command option suggestions.
	Options []string `mandatory:"false" json:"options"`

	// Context specific list of querylanguage querystring examples if input is seeking queryString example suggestions.
	Examples []string `mandatory:"false" json:"examples"`
}

func (m SuggestOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SuggestOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
