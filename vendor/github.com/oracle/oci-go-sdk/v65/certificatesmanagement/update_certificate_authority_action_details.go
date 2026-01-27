// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateCertificateAuthorityActionDetails The details of the type of certificate authority (CA) update request. Updates can be performed by updating certificate contents or by generating a certificate signing request (CSR).
type UpdateCertificateAuthorityActionDetails interface {
}

type updatecertificateauthorityactiondetails struct {
	JsonData   []byte
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *updatecertificateauthorityactiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatecertificateauthorityactiondetails updatecertificateauthorityactiondetails
	s := struct {
		Model Unmarshalerupdatecertificateauthorityactiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatecertificateauthorityactiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "UPDATE_CERTIFICATE":
		mm := UpdateCertificateAuthorityCertificateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERATE_CSR":
		mm := UpdateCertificateAuthorityGenerateCsrDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateCertificateAuthorityActionDetails: %s.", m.ActionType)
		return *m, nil
	}
}

func (m updatecertificateauthorityactiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatecertificateauthorityactiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
