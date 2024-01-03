// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AffectedResource The resource affected by the event described in the announcement.
type AffectedResource struct {

	// The OCID of the affected resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The friendly name of the resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The region where the affected resource exists.
	Region *string `mandatory:"true" json:"region"`

	// Additional properties associated with the resource.
	AdditionalProperties []Property `mandatory:"false" json:"additionalProperties"`
}

func (m AffectedResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AffectedResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
