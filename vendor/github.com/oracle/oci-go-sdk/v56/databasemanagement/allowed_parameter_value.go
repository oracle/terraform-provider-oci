// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AllowedParameterValue A valid value for a database parameter.
type AllowedParameterValue struct {

	// The ordinal number in the list (1-based).
	Ordinal *float32 `mandatory:"false" json:"ordinal"`

	// The parameter value at ordinal.
	Value *string `mandatory:"false" json:"value"`

	// Indicates whether the given ordinal value is the default value for the parameter.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m AllowedParameterValue) String() string {
	return common.PointerString(m)
}
