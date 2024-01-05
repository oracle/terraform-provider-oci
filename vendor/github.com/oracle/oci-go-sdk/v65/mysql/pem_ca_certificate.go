// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PemCaCertificate The CA certificate in PEM format.
type PemCaCertificate struct {

	// The string containing the CA certificate in PEM format.
	Contents *string `mandatory:"true" json:"contents"`
}

func (m PemCaCertificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PemCaCertificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PemCaCertificate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePemCaCertificate PemCaCertificate
	s := struct {
		DiscriminatorParam string `json:"certificateType"`
		MarshalTypePemCaCertificate
	}{
		"PEM",
		(MarshalTypePemCaCertificate)(m),
	}

	return json.Marshal(&s)
}
