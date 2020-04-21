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

// UpdateTermDetails Properties used in term update operations.
type UpdateTermDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the term.
	Description *string `mandatory:"false" json:"description"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// OCID of the user who is the owner of this business terminology.
	Owner *string `mandatory:"false" json:"owner"`

	// Status of the approval process workflow for this business term in the glossary
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`
}

func (m UpdateTermDetails) String() string {
	return common.PointerString(m)
}
