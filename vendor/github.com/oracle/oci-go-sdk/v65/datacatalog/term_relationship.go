// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TermRelationship Full term relationship definition. Business term relationship between two terms in a business glossary.
type TermRelationship struct {

	// Unique term relationship key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.This is the same as relationshipType for termRelationship
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the term relationship usually defined at the time of creation.
	Description *string `mandatory:"false" json:"description"`

	// Unique id of the related term.
	RelatedTermKey *string `mandatory:"false" json:"relatedTermKey"`

	// Name of the related term.
	RelatedTermDisplayName *string `mandatory:"false" json:"relatedTermDisplayName"`

	// Description of the related term.
	RelatedTermDescription *string `mandatory:"false" json:"relatedTermDescription"`

	// Full path of the related term.
	RelatedTermPath *string `mandatory:"false" json:"relatedTermPath"`

	// Glossary key of the related term.
	RelatedTermGlossaryKey *string `mandatory:"false" json:"relatedTermGlossaryKey"`

	// URI to the term relationship instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// This relationships parent term key.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term.
	ParentTermDisplayName *string `mandatory:"false" json:"parentTermDisplayName"`

	// Description of the parent term.
	ParentTermDescription *string `mandatory:"false" json:"parentTermDescription"`

	// Full path of the parent term.
	ParentTermPath *string `mandatory:"false" json:"parentTermPath"`

	// Glossary key of the parent term.
	ParentTermGlossaryKey *string `mandatory:"false" json:"parentTermGlossaryKey"`

	// The date and time the term relationship was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// State of the term relationship.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TermRelationship) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TermRelationship) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
