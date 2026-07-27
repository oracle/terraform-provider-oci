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

// CreateDbClusterSourceFromImportUrlDetails Specifies the pre-authenticated request (PAR) URL of an Object Storage bucket from which the initial data of the shared-storage DB cluster is imported.
type CreateDbClusterSourceFromImportUrlDetails struct {

	// The pre-authenticated request (PAR) URL of an Object Storage bucket, prefix, or a @.manifest.json object.
	// For a bucket or prefix, create the PAR with "Permit object reads" access type and "Enable Object Listing" permission.
	// For a @.manifest.json object, create the PAR with "Permit object reads" access type.
	// To learn how to create a PAR, see Using Pre-Authenticated Requests (https://docs.oracle.com/en-us/iaas/Content/Object/Tasks/usingpreauthenticatedrequests.htm).
	SourceUrl *string `mandatory:"true" json:"sourceUrl"`
}

func (m CreateDbClusterSourceFromImportUrlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbClusterSourceFromImportUrlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDbClusterSourceFromImportUrlDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbClusterSourceFromImportUrlDetails CreateDbClusterSourceFromImportUrlDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbClusterSourceFromImportUrlDetails
	}{
		"IMPORTURL",
		(MarshalTypeCreateDbClusterSourceFromImportUrlDetails)(m),
	}

	return json.Marshal(&s)
}
