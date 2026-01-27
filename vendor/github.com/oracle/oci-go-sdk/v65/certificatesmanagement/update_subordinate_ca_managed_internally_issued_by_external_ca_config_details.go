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

// UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails The configuration details for updating an internally managed subordinate certificate authority (CA) which is issued by an external CA.
type UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails struct {
	ActionDetails UpdateCertificateAuthorityActionDetails `mandatory:"true" json:"actionDetails"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	// The rotation state of the CA. The default is `PENDING`, meaning that the CA is staged and available for use. A CA version
	// that you mark as `CURRENT` is currently in use, but you don't yet want to rotate it into current, active use. For example,
	// you might create or update a CA and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	Stage UpdateCertificateAuthorityConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
}

// GetVersionName returns VersionName
func (m UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

// GetStage returns Stage
func (m UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) GetStage() UpdateCertificateAuthorityConfigDetailsStageEnum {
	return m.Stage
}

func (m UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCertificateAuthorityConfigDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetUpdateCertificateAuthorityConfigDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails
	}{
		"SUBORDINATE_CA_MANAGED_INTERNALLY_ISSUED_BY_EXTERNAL_CA",
		(MarshalTypeUpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSubordinateCaManagedInternallyIssuedByExternalCaConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		VersionName   *string                                          `json:"versionName"`
		Stage         UpdateCertificateAuthorityConfigDetailsStageEnum `json:"stage"`
		ActionDetails updatecertificateauthorityactiondetails          `json:"actionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.VersionName = model.VersionName

	m.Stage = model.Stage

	nn, e = model.ActionDetails.UnmarshalPolymorphicJSON(model.ActionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ActionDetails = nn.(UpdateCertificateAuthorityActionDetails)
	} else {
		m.ActionDetails = nil
	}

	return
}
