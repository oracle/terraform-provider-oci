// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PlainPasswordSecurityUserPasswordDetails Password-source details for a plaintext master-user password.
type PlainPasswordSecurityUserPasswordDetails struct {

	// The plaintext password for the master user.
	PlainPassword *string `mandatory:"true" json:"plainPassword" sensitive:"true"`
}

func (m PlainPasswordSecurityUserPasswordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlainPasswordSecurityUserPasswordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PlainPasswordSecurityUserPasswordDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePlainPasswordSecurityUserPasswordDetails PlainPasswordSecurityUserPasswordDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypePlainPasswordSecurityUserPasswordDetails
	}{
		"PLAIN_PASSWORD",
		(MarshalTypePlainPasswordSecurityUserPasswordDetails)(m),
	}

	return json.Marshal(&s)
}
