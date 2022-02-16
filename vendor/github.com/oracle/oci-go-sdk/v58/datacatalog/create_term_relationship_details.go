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

// CreateTermRelationshipDetails Properties used in term relationship create operations.
type CreateTermRelationshipDetails struct {

	// A user-friendly display name. Is changeable. The combination of 'displayName' and 'parentTermKey'
	// must be unique. Avoid entering confidential information. This is the same as 'relationshipType' for 'termRelationship'.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique id of the related term.
	RelatedTermKey *string `mandatory:"true" json:"relatedTermKey"`

	// Detailed description of the term relationship usually defined at the time of creation.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreateTermRelationshipDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTermRelationshipDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
