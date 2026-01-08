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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KerberosCredential Kerberos Credential details as OCI Vault Secret for the keytab.
type KerberosCredential struct {

	// Kerberos realm name.
	// https://docs.oracle.com/cd/E36784_01/html/E37126/kplanning-27.html
	// Realm names can consist of any ASCII string. Usually, the realm name is the same as your DNS domain name
	// except that the realm name is in uppercase. This convention helps differentiate problems with the Kerberos
	// service from problems with the DNS namespace, while keeping a name that is familiar. You can use any string,
	// but configuration and maintenance might then require more work. Use realm names that follow the standard
	// Internet naming structure.
	RealmName *string `mandatory:"true" json:"realmName"`

	// The OCID of the secret where the Kerberos keytab file is stored as base64 text.
	KeytabSecretId *string `mandatory:"true" json:"keytabSecretId"`

	// The secret version of the stored Kerberos keytab file.
	KeytabSecretVersion *int64 `mandatory:"true" json:"keytabSecretVersion"`
}

func (m KerberosCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KerberosCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
