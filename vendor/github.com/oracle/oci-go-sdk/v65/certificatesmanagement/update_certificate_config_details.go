// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateCertificateConfigDetails The details of the contents of the certificate and certificate metadata.
type UpdateCertificateConfigDetails interface {

	// A name for the certificate version. When the value is not null, a name is unique across versions of a given certificate.
	GetVersionName() *string

	// The rotation state of the certificate. The default is `CURRENT`, meaning that the certificate is currently in use. A certificate version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might update a certificate and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	GetStage() UpdateCertificateConfigDetailsStageEnum
}

type updatecertificateconfigdetails struct {
	JsonData    []byte
	VersionName *string                                 `mandatory:"false" json:"versionName"`
	Stage       UpdateCertificateConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
	ConfigType  string                                  `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *updatecertificateconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatecertificateconfigdetails updatecertificateconfigdetails
	s := struct {
		Model Unmarshalerupdatecertificateconfigdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VersionName = s.Model.VersionName
	m.Stage = s.Model.Stage
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatecertificateconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "IMPORTED":
		mm := UpdateCertificateByImportingConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ISSUED_BY_INTERNAL_CA":
		mm := UpdateCertificateIssuedByInternalCaConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA":
		mm := UpdateCertificateManagedExternallyIssuedByInternalCaConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateCertificateConfigDetails: %s.", m.ConfigType)
		return *m, nil
	}
}

// GetVersionName returns VersionName
func (m updatecertificateconfigdetails) GetVersionName() *string {
	return m.VersionName
}

// GetStage returns Stage
func (m updatecertificateconfigdetails) GetStage() UpdateCertificateConfigDetailsStageEnum {
	return m.Stage
}

func (m updatecertificateconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatecertificateconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCertificateConfigDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetUpdateCertificateConfigDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCertificateConfigDetailsStageEnum Enum with underlying type: string
type UpdateCertificateConfigDetailsStageEnum string

// Set of constants representing the allowable values for UpdateCertificateConfigDetailsStageEnum
const (
	UpdateCertificateConfigDetailsStageCurrent UpdateCertificateConfigDetailsStageEnum = "CURRENT"
	UpdateCertificateConfigDetailsStagePending UpdateCertificateConfigDetailsStageEnum = "PENDING"
)

var mappingUpdateCertificateConfigDetailsStageEnum = map[string]UpdateCertificateConfigDetailsStageEnum{
	"CURRENT": UpdateCertificateConfigDetailsStageCurrent,
	"PENDING": UpdateCertificateConfigDetailsStagePending,
}

var mappingUpdateCertificateConfigDetailsStageEnumLowerCase = map[string]UpdateCertificateConfigDetailsStageEnum{
	"current": UpdateCertificateConfigDetailsStageCurrent,
	"pending": UpdateCertificateConfigDetailsStagePending,
}

// GetUpdateCertificateConfigDetailsStageEnumValues Enumerates the set of values for UpdateCertificateConfigDetailsStageEnum
func GetUpdateCertificateConfigDetailsStageEnumValues() []UpdateCertificateConfigDetailsStageEnum {
	values := make([]UpdateCertificateConfigDetailsStageEnum, 0)
	for _, v := range mappingUpdateCertificateConfigDetailsStageEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCertificateConfigDetailsStageEnumStringValues Enumerates the set of values in String for UpdateCertificateConfigDetailsStageEnum
func GetUpdateCertificateConfigDetailsStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
	}
}

// GetMappingUpdateCertificateConfigDetailsStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCertificateConfigDetailsStageEnum(val string) (UpdateCertificateConfigDetailsStageEnum, bool) {
	enum, ok := mappingUpdateCertificateConfigDetailsStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
