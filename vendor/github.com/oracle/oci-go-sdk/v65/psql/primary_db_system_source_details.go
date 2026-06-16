// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrimaryDbSystemSourceDetails Details of the primary database system.
type PrimaryDbSystemSourceDetails struct {

	// The [OCID] of the primary database system.
	PrimaryDbSystemId *string `mandatory:"true" json:"primaryDbSystemId"`
}

func (m PrimaryDbSystemSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrimaryDbSystemSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PrimaryDbSystemSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrimaryDbSystemSourceDetails PrimaryDbSystemSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypePrimaryDbSystemSourceDetails
	}{
		"DB_SYSTEM",
		(MarshalTypePrimaryDbSystemSourceDetails)(m),
	}

	return json.Marshal(&s)
}
