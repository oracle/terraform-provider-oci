// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Summary Summary of the audit report.
type Summary struct {

	// Name of the report summary.
	Name *string `mandatory:"true" json:"name"`

	// Specifies the order in which the summary must be displayed.
	DisplayOrder *int `mandatory:"true" json:"displayOrder"`

	// Indicates if the summary is hidden. Values can either be 'true' or 'false'.
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// A comma-delimited string that specifies the names of the fields by which the records must be aggregated to get the summary.
	GroupByFieldName *string `mandatory:"false" json:"groupByFieldName"`

	// Name of the key or count of object.
	CountOf *string `mandatory:"false" json:"countOf"`

	// Additional scim filters used to get the specific summary.
	ScimFilter *string `mandatory:"false" json:"scimFilter"`
}

func (m Summary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Summary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
