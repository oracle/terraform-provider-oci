// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDbNodeDetails Details used to update a DB node.
type UpdateDbNodeDetails struct {

	// Description of the DB node.
	Description *string `mandatory:"false" json:"description"`

	// Name of the DB node. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Preference of a DB node as a potential failover target.
	// This is an integer  that defines the relative weight/priority
	// of a particular DB node versus another, where the lower the number,
	// the higher the preference for that DB node to become a new primary on a failover.
	PromotionTier *int `mandatory:"false" json:"promotionTier"`

	ReadEndpoint *UpdateDbNodeReadEndpointDetails `mandatory:"false" json:"readEndpoint"`
}

func (m UpdateDbNodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbNodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
