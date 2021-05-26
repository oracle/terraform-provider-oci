// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// PemCaCertificate The CA certificate in PEM format.
type PemCaCertificate struct {

	// The string containing the CA certificate in PEM format.
	Contents *string `mandatory:"true" json:"contents"`
}

func (m PemCaCertificate) String() string {
	return common.PointerString(m)
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
