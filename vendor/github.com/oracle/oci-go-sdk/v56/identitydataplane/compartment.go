// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
