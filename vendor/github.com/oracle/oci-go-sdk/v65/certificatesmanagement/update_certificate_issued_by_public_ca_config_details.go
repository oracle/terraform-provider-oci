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

// UpdateCertificateIssuedByPublicCaConfigDetails The details of the configuration for updating a certificate issued by public CA.
type UpdateCertificateIssuedByPublicCaConfigDetails struct {

	// A name for the certificate version. When the value is not null, a name is unique across versions of a given certificate.
	VersionName *string `mandatory:"false" json:"versionName"`

	Validity *PublicCaCertificateValidity `mandatory:"false" json:"validity"`

	// The rotation state of the certificate. The default is `CURRENT`, meaning that the certificate is currently in use. A certificate version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might update a certificate and mark its rotation state as `PENDING` if you haven't yet updated the certificate on the target system.
	Stage UpdateCertificateConfigDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`

	// The name of the public CA issuing the certificate.
	PublicCertificateAuthority PublicCertificateAuthorityEnum `mandatory:"false" json:"publicCertificateAuthority,omitempty"`
}

// GetVersionName returns VersionName
func (m UpdateCertificateIssuedByPublicCaConfigDetails) GetVersionName() *string {
	return m.VersionName
}

// GetStage returns Stage
func (m UpdateCertificateIssuedByPublicCaConfigDetails) GetStage() UpdateCertificateConfigDetailsStageEnum {
	return m.Stage
}

func (m UpdateCertificateIssuedByPublicCaConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCertificateIssuedByPublicCaConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateCertificateConfigDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetUpdateCertificateConfigDetailsStageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPublicCertificateAuthorityEnum(string(m.PublicCertificateAuthority)); !ok && m.PublicCertificateAuthority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublicCertificateAuthority: %s. Supported values are: %s.", m.PublicCertificateAuthority, strings.Join(GetPublicCertificateAuthorityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateCertificateIssuedByPublicCaConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCertificateIssuedByPublicCaConfigDetails UpdateCertificateIssuedByPublicCaConfigDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeUpdateCertificateIssuedByPublicCaConfigDetails
	}{
		"ISSUED_BY_PUBLIC_CA",
		(MarshalTypeUpdateCertificateIssuedByPublicCaConfigDetails)(m),
	}

	return json.Marshal(&s)
}
