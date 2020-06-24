// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Version A supported MySQL Version.
type Version struct {

	// The specific version identifier
	Version *string `mandatory:"false" json:"version"`

	// A link to a page describing the version.
	Description *string `mandatory:"false" json:"description"`
}

func (m Version) String() string {
	return common.PointerString(m)
}
