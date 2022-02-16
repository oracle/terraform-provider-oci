// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateDbSystemSourceImportFromUrlDetails An Object Storage PAR from which to import the DB System initial data.
type CreateDbSystemSourceImportFromUrlDetails struct {

	// The Pre-Authenticated Request (PAR) URL of the file you want to import from Object Storage.
	SourceUrl *string `mandatory:"true" json:"sourceUrl"`
}

func (m CreateDbSystemSourceImportFromUrlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemSourceImportFromUrlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
