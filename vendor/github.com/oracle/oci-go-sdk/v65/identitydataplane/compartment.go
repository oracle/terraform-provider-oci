// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Compartment The representation of Compartment
type Compartment struct {

	// The id of the compartment.
	Id *string `mandatory:"true" json:"id"`

	// The name of the compartment.
	Name *string `mandatory:"true" json:"name"`

	// The display name of the compartment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The full name of the compartment.
	FullName *string `mandatory:"true" json:"fullName"`

	// The id of the parent compartment.
	ParentCompartmentId *string `mandatory:"true" json:"parentCompartmentId"`

	// The status of the compartment.
	Status *EntityStatus `mandatory:"true" json:"status"`

	// The extended properties.
	PropertyMap map[string]string `mandatory:"true" json:"propertyMap"`
}

func (m Compartment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Compartment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
