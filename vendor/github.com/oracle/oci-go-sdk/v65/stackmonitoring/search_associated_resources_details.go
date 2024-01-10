// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchAssociatedResourcesDetails The criteria for searching associated monitored resources.
type SearchAssociatedResourcesDetails struct {

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A filter to return associated resources that match resources of type.
	// Either resourceId or resourceType should be provided.
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Monitored resource identifier for which the associated resources should be fetched.
	// Either resourceId or resourceType should be provided.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The field which determines the depth of hierarchy while searching for associated resources.
	// Possible values - 0 for all levels. And positive number to indicate different levels.
	// Default value is 1, which indicates 1st level associations.
	LimitLevel *int `mandatory:"false" json:"limitLevel"`

	// Association types filter to be searched for finding associated resources.
	AssociationTypes []string `mandatory:"false" json:"associationTypes"`
}

func (m SearchAssociatedResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchAssociatedResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
