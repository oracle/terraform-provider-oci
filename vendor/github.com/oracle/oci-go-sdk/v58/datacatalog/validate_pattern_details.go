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

// ValidatePatternDetails Validate pattern using the expression and file list.
type ValidatePatternDetails struct {

	// Input string which drives the selection process, allowing for fine-grained control using qualifiers.
	// Refer to the user documentation for details of the format and examples. A pattern cannot include both
	// a prefix and an expression.
	Expression *string `mandatory:"false" json:"expression"`

	// Input string which drives the selection process.
	// Refer to the user documentation for details of the format and examples. A pattern cannot include both
	// a prefix and an expression.
	FilePathPrefix *string `mandatory:"false" json:"filePathPrefix"`

	// List of file paths against which the pattern can be tried, as a check. This documents, for reference
	// purposes, some example objects a pattern is meant to work with.
	// If provided with the request,this overrides the list which already exists as part of the pattern, if any.
	CheckFilePathList []string `mandatory:"false" json:"checkFilePathList"`

	// The maximum number of UNMATCHED files, in checkFilePathList, above which the check fails.
	// Optional, if checkFilePathList is provided.
	// If provided with the request, this overrides the value which already exists as part of the pattern, if any.
	CheckFailureLimit *int `mandatory:"false" json:"checkFailureLimit"`
}

func (m ValidatePatternDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidatePatternDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
