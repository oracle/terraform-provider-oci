// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MembershipSummary Summary of an Exadata Fleet Update Collection containing a target.
type MembershipSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Collection.
	FsuCollectionId *string `mandatory:"true" json:"fsuCollectionId"`

	// Type of Exadata Fleet Update Collection.
	FsuCollectionType CollectionTypesEnum `mandatory:"true" json:"fsuCollectionType"`

	// The user-friendly name for the Exadata Fleet Update Collection.
	FsuCollectionName *string `mandatory:"true" json:"fsuCollectionName"`
}

func (m MembershipSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MembershipSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectionTypesEnum(string(m.FsuCollectionType)); !ok && m.FsuCollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FsuCollectionType: %s. Supported values are: %s.", m.FsuCollectionType, strings.Join(GetCollectionTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
