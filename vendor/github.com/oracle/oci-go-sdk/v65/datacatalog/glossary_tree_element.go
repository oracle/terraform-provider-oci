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

// GlossaryTreeElement Glossary tree element with child terms.
type GlossaryTreeElement struct {

	// Unique term key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the term.
	Description *string `mandatory:"false" json:"description"`

	// Unique id of the parent glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// URI to the term instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Indicates whether a term may contain child terms.
	IsAllowedToHaveChildTerms *bool `mandatory:"false" json:"isAllowedToHaveChildTerms"`

	// Absolute path of the term.
	Path *string `mandatory:"false" json:"path"`

	// The date and time the term was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Status of the approval process workflow for this business term in the glossary.
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`

	// The number of objects tagged with this term.
	AssociatedObjectCount *int `mandatory:"false" json:"associatedObjectCount"`

	// State of the term.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An array of child terms.
	ChildTerms []GlossaryTreeElement `mandatory:"false" json:"childTerms"`
}

func (m GlossaryTreeElement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GlossaryTreeElement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTermWorkflowStatusEnum(string(m.WorkflowStatus)); !ok && m.WorkflowStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkflowStatus: %s. Supported values are: %s.", m.WorkflowStatus, strings.Join(GetTermWorkflowStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
