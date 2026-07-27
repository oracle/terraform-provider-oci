// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DbClusterSourceFromImportUrl Specifies the pre-authenticated request (PAR) URL of the Object Storage bucket, prefix, or object from which the initial data of the shared-storage DB cluster was imported.
type DbClusterSourceFromImportUrl struct {
}

func (m DbClusterSourceFromImportUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterSourceFromImportUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbClusterSourceFromImportUrl) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbClusterSourceFromImportUrl DbClusterSourceFromImportUrl
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbClusterSourceFromImportUrl
	}{
		"IMPORTURL",
		(MarshalTypeDbClusterSourceFromImportUrl)(m),
	}

	return json.Marshal(&s)
}
