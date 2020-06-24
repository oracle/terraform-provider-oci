// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbSystemSourceImportFromUrlDetails An Object Storage PAR from which to import the DB System initial data.
type CreateDbSystemSourceImportFromUrlDetails struct {

	// The Pre-Authenticated Request (PAR) URL of the file you want to import from Object Storage.
	SourceUrl *string `mandatory:"true" json:"sourceUrl"`
}

func (m CreateDbSystemSourceImportFromUrlDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbSystemSourceImportFromUrlDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbSystemSourceImportFromUrlDetails CreateDbSystemSourceImportFromUrlDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbSystemSourceImportFromUrlDetails
	}{
		"IMPORTURL",
		(MarshalTypeCreateDbSystemSourceImportFromUrlDetails)(m),
	}

	return json.Marshal(&s)
}
