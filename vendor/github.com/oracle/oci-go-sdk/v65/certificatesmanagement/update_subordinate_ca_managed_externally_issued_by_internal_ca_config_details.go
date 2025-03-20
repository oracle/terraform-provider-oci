// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails The details for updating an externally managed subordinate certificate authority (CA) which is issued by an internal CA.
type UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails struct {

	// The certificate signing request (in PEM format).
	CsrPem *string `mandatory:"true" json:"csrPem"`

	// The name of the CA version. When the value is not null, a name is unique across versions of a given CA.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *Validity `mandatory:"false" json:"validity"`

	// The rotation state of the CA. The default is `PENDING`, meaning that the CA is staged and available for use. A CA version
	// that you mark as `CURRENT` is currently in use, but you don't yet want to rotate it into current, active use. For example,
	// you might create or update a CA and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	Stage UpdateCertificateAuthorityConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
}

// GetVersionName returns VersionName
func (m UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

// GetStage returns Stage
func (m UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) GetStage() UpdateCertificateAuthorityConfigDetailsStageEnum {
	return m.Stage
}

func (m UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCertificateAuthorityConfigDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetUpdateCertificateAuthorityConfigDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails UpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails
	}{
		"SUBORDINATE_CA_MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA",
		(MarshalTypeUpdateSubordinateCaManagedExternallyIssuedByInternalCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
