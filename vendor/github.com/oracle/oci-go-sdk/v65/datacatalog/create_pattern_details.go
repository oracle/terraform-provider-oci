// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePatternDetails Properties used in pattern create operations.
type CreatePatternDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the Pattern.
	Description *string `mandatory:"false" json:"description"`

	// Input string which drives the selection process, allowing for fine-grained control using qualifiers.
	// Refer to the user documentation for details of the format and examples. A pattern cannot include both
	// a prefix and an expression.
	Expression *string `mandatory:"false" json:"expression"`

	// Input string which drives the selection process.
	// Refer to the user documentation for details of the format and examples. A pattern cannot include both
	// a prefix and an expression.
	FilePathPrefix *string `mandatory:"false" json:"filePathPrefix"`

	// List of file paths against which the pattern can be tried, as a check. This documents, for reference
	// purposes, some example objects a pattern is meant to work with. If isEnableCheckFailureLimit is set to true,
	// this will be run as a validation during the request, such that if the check fails the request fails. If
	// isEnableCheckFailureLimit instead is set to (the default) false, a pattern will still be created or updated even
	// if the check fails, with a lifecycleState of FAILED.
	CheckFilePathList []string `mandatory:"false" json:"checkFilePathList"`

	// Indicates whether the pattern check, against the checkFilePathList, will fail the request if the count of
	// UNMATCHED files is above the checkFailureLimit.
	IsEnableCheckFailureLimit *bool `mandatory:"false" json:"isEnableCheckFailureLimit"`

	// The maximum number of UNMATCHED files, in checkFilePathList, above which the check fails. Optional, if
	// checkFilePathList is provided - but if isEnableCheckFailureLimit is set to true it is required.
	CheckFailureLimit *int `mandatory:"false" json:"checkFailureLimit"`

	// A map of maps that contains the properties which are specific to the pattern type. Each pattern type
	// definition defines it's set of required and optional properties.
	// Example: `{"properties": { "default": { "tbd"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreatePatternDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePatternDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
