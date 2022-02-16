// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateCertificateAuthorityConfigDetails The configuration details for creating a certificate authority (CA).
type CreateCertificateAuthorityConfigDetails interface {

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	GetVersionName() *string
}

type createcertificateauthorityconfigdetails struct {
	JsonData    []byte
	VersionName *string `mandatory:"false" json:"versionName"`
	ConfigType  string  `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *createcertificateauthorityconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatecertificateauthorityconfigdetails createcertificateauthorityconfigdetails
	s := struct {
		Model Unmarshalercreatecertificateauthorityconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VersionName = s.Model.VersionName
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createcertificateauthorityconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "ROOT_CA_GENERATED_INTERNALLY":
		mm := CreateRootCaByGeneratingInternallyConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA":
		mm := CreateSubordinateCaIssuedByInternalCaConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetVersionName returns VersionName
func (m createcertificateauthorityconfigdetails) GetVersionName() *string {
	return m.VersionName
}

func (m createcertificateauthorityconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createcertificateauthorityconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
