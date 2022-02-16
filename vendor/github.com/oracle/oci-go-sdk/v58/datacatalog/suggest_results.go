// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SuggestResults The list of potential matches returned from the suggest operation for the given input text. The size of the list will be determined
// by the limit parameter.
type SuggestResults struct {

	// Total number of items returned.
	TotalCount *int `mandatory:"true" json:"totalCount"`

	// Input string for which the potential matches are computed.
	InputText *string `mandatory:"true" json:"inputText"`

	// Time taken to compute the result, in milliseconds.
	SearchLatencyInMs *int `mandatory:"false" json:"searchLatencyInMs"`

	// List of suggestions.
	Items []SuggestListItem `mandatory:"false" json:"items"`
}

func (m SuggestResults) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SuggestResults) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
