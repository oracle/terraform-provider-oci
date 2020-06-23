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

// VersionSummary A summary of the supported MySQL Versions families, and a list of their supported minor versions.
type VersionSummary struct {

	// The list of supported MySQL Versions.
	Versions []Version `mandatory:"true" json:"versions"`

	// A descriptive summary of a group of versions.
	VersionFamily *string `mandatory:"false" json:"versionFamily"`
}

func (m VersionSummary) String() string {
	return common.PointerString(m)
}
