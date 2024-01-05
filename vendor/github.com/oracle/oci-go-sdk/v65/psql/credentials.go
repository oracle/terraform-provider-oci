// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Credentials Initial database system credentials that the database system will be provisioned with.
// The password details are not visible on any subsequent operation, such as GET /dbSystems/{dbSystemId}.
type Credentials struct {

	// The database system administrator username.
	Username *string `mandatory:"true" json:"username"`

	PasswordDetails PasswordDetails `mandatory:"true" json:"passwordDetails"`
}

func (m Credentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Credentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Credentials) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Username        *string         `json:"username"`
		PasswordDetails passworddetails `json:"passwordDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Username = model.Username

	nn, e = model.PasswordDetails.UnmarshalPolymorphicJSON(model.PasswordDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PasswordDetails = nn.(PasswordDetails)
	} else {
		m.PasswordDetails = nil
	}

	return
}
