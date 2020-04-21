// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTermDetails Properties used in term create operations.
type CreateTermDetails struct {

	// A user-friendly display name. Is changeable. The combination of 'displayName' and 'parentTermKey'
	// must be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the term.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether a term may contain child terms.
	IsAllowedToHaveChildTerms *bool `mandatory:"false" json:"isAllowedToHaveChildTerms"`

	// The terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// OCID of the user who is the owner of this business terminology.
	Owner *string `mandatory:"false" json:"owner"`

	// Status of the approval process workflow for this business term in the glossary.
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`
}

func (m CreateTermDetails) String() string {
	return common.PointerString(m)
}
