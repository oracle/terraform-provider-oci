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

// UpdateCertificateAuthorityConfigDetails The configuration details for updating a certificate authority (CA).
type UpdateCertificateAuthorityConfigDetails interface {

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	GetVersionName() *string

	// The rotation state of the CA. The default is `PENDING`, meaning that the CA is staged and available for use. A CA version
	// that you mark as `CURRENT` is currently in use, but you don't yet want to rotate it into current, active use. For example,
	// you might create or update a CA and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	GetStage() UpdateCertificateAuthorityConfigDetailsStageEnum
}

type updatecertificateauthorityconfigdetails struct {
	JsonData    []byte
	VersionName *string                                          `mandatory:"false" json:"versionName"`
	Stage       UpdateCertificateAuthorityConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
	ConfigType  string                                           `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *updatecertificateauthorityconfigdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatecertificateauthorityconfigdetails updatecertificateauthorityconfigdetails
	s := struct {
		Model Unmarshalerupdatecertificateauthorityconfigdetails
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
func (m *updatecertificateauthorityconfigdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA":
		mm := UpdateSubordinateCaIssuedByInternalCaConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROOT_CA_GENERATED_INTERNALLY":
		mm := UpdateRootCaByGeneratingInternallyConfigDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetVersionName returns VersionName
func (m updatecertificateauthorityconfigdetails) GetVersionName() *string {
	return m.VersionName
}

//GetStage returns Stage
func (m updatecertificateauthorityconfigdetails) GetStage() UpdateCertificateAuthorityConfigDetailsStageEnum {
	return m.Stage
}

func (m updatecertificateauthorityconfigdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatecertificateauthorityconfigdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCertificateAuthorityConfigDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetUpdateCertificateAuthorityConfigDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCertificateAuthorityConfigDetailsStageEnum Enum with underlying type: string
type UpdateCertificateAuthorityConfigDetailsStageEnum string

// Set of constants representing the allowable values for UpdateCertificateAuthorityConfigDetailsStageEnum
const (
	UpdateCertificateAuthorityConfigDetailsStageCurrent UpdateCertificateAuthorityConfigDetailsStageEnum = "CURRENT"
	UpdateCertificateAuthorityConfigDetailsStagePending UpdateCertificateAuthorityConfigDetailsStageEnum = "PENDING"
)

var mappingUpdateCertificateAuthorityConfigDetailsStageEnum = map[string]UpdateCertificateAuthorityConfigDetailsStageEnum{
	"CURRENT": UpdateCertificateAuthorityConfigDetailsStageCurrent,
	"PENDING": UpdateCertificateAuthorityConfigDetailsStagePending,
}

// GetUpdateCertificateAuthorityConfigDetailsStageEnumValues Enumerates the set of values for UpdateCertificateAuthorityConfigDetailsStageEnum
func GetUpdateCertificateAuthorityConfigDetailsStageEnumValues() []UpdateCertificateAuthorityConfigDetailsStageEnum {
	values := make([]UpdateCertificateAuthorityConfigDetailsStageEnum, 0)
	for _, v := range mappingUpdateCertificateAuthorityConfigDetailsStageEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCertificateAuthorityConfigDetailsStageEnumStringValues Enumerates the set of values in String for UpdateCertificateAuthorityConfigDetailsStageEnum
func GetUpdateCertificateAuthorityConfigDetailsStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
	}
}

// GetMappingUpdateCertificateAuthorityConfigDetailsStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCertificateAuthorityConfigDetailsStageEnum(val string) (UpdateCertificateAuthorityConfigDetailsStageEnum, bool) {
	mappingUpdateCertificateAuthorityConfigDetailsStageEnumIgnoreCase := make(map[string]UpdateCertificateAuthorityConfigDetailsStageEnum)
	for k, v := range mappingUpdateCertificateAuthorityConfigDetailsStageEnum {
		mappingUpdateCertificateAuthorityConfigDetailsStageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateCertificateAuthorityConfigDetailsStageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
