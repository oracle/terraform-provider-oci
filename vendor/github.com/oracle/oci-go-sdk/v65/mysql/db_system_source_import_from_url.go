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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystemSourceImportFromUrl An Object Storage PAR from which to import the DB System initial data.
type DbSystemSourceImportFromUrl struct {
}

func (m DbSystemSourceImportFromUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSourceImportFromUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceImportFromUrl) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceImportFromUrl DbSystemSourceImportFromUrl
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceImportFromUrl
	}{
		"IMPORTURL",
		(MarshalTypeDbSystemSourceImportFromUrl)(m),
	}

	return json.Marshal(&s)
}
