// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// Pattern Pattern representation. A Pattern is defined using an expression and can be used as data selectors or filters
// to provide a singular view of an entity across multiple physical data artifacts.
type Pattern struct {

	// Unique pattern key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the pattern.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// The current state of the data asset.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the pattern was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the pattern. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the pattern.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who last modified the pattern.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// The expression used in the pattern that may include qualifiers. Refer to the user documentation for details of the format and examples.
	Expression *string `mandatory:"false" json:"expression"`

	// List of file paths against which the expression can be tried, as a check. This documents, for reference
	// purposes, some example objects a pattern is meant to work with. If isEnableCheckFailureLimit is set to true,
	// this will be run as a validation during the request, such that if the check fails the request fails. If
	// isEnableCheckFailureLimit instead is set to (the default) false, a pattern will still be created or updated even
	// if the check fails, with a lifecycleState of FAILED.
	CheckFilePathList []string `mandatory:"false" json:"checkFilePathList"`

	// Indicates whether the expression check, against the checkFilePathList, will fail the request if the count of
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

func (m Pattern) String() string {
	return common.PointerString(m)
}
