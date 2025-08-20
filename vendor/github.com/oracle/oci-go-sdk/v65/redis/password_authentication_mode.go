// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PasswordAuthenticationMode child class of AuthenticationMode.
type PasswordAuthenticationMode struct {

	// SHA-256 hashed passwords for OCI Cache user,required if authenticationType is set to PASSWORD.
	HashedPasswords []string `mandatory:"true" json:"hashedPasswords"`
}

func (m PasswordAuthenticationMode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PasswordAuthenticationMode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PasswordAuthenticationMode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePasswordAuthenticationMode PasswordAuthenticationMode
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypePasswordAuthenticationMode
	}{
		"PASSWORD",
		(MarshalTypePasswordAuthenticationMode)(m),
	}

	return json.Marshal(&s)
}
