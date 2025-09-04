// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMySqlDbSystemReplicationUserDetails The credentials for the HeatWave MySQL DB System replication user, containing the username and the OCID of the vault secret that stores the password.
type UpdateMySqlDbSystemReplicationUserDetails struct {

	// The user name for connecting to the HeatWave MySQL DB System node.
	// Example: `user`
	Username *string `mandatory:"true" json:"username"`

	// The OCID of the vault secret where the HeatWave MySQL DB System password is stored.
	// Example: `ocid1.vaultsecret.oc1..uniqueID`
	PasswordVaultSecretId *string `mandatory:"true" json:"passwordVaultSecretId"`
}

func (m UpdateMySqlDbSystemReplicationUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMySqlDbSystemReplicationUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
