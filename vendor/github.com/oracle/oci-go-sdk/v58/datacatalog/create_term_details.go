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

// CreateTermDetails Properties used in term create operations.
type CreateTermDetails struct {

	// A user-friendly display name. Is changeable. The combination of 'displayName' and 'parentTermKey'
	// must be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the term.
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether a term may contain child terms.
	IsAllowedToHaveChildTerms *bool `mandatory:"false" json:"isAllowedToHaveChildTerms"`

	// The parent key of the term. In the case of a root-level category only, the term would have no parent and this should be left unset.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// OCID of the user who is the owner of this business terminology.
	Owner *string `mandatory:"false" json:"owner"`

	// Status of the approval process workflow for this business term in the glossary.
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertySetUsage `mandatory:"false" json:"customPropertyMembers"`
}

func (m CreateTermDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTermDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTermWorkflowStatusEnum(string(m.WorkflowStatus)); !ok && m.WorkflowStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkflowStatus: %s. Supported values are: %s.", m.WorkflowStatus, strings.Join(GetTermWorkflowStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
