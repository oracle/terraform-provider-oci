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

// EnabledKerberosAuthDetails Enable or Update Existing Kerberos Authentication for the database system.
type EnabledKerberosAuthDetails struct {

	// List of Kerberos Credentials to be configured for the dbsystem. Currently supports only one entry.
	Credentials []KerberosCredential `mandatory:"true" json:"credentials"`

	// Optional. List of Kerberos Credentials previously configured for the dbsystem. Currently supports only one entry.
	BackupCredentials []KerberosCredential `mandatory:"false" json:"backupCredentials"`
}

func (m EnabledKerberosAuthDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnabledKerberosAuthDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnabledKerberosAuthDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnabledKerberosAuthDetails EnabledKerberosAuthDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeEnabledKerberosAuthDetails
	}{
		"ENABLED",
		(MarshalTypeEnabledKerberosAuthDetails)(m),
	}

	return json.Marshal(&s)
}
