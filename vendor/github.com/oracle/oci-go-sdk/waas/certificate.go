// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Certificate The details of the SSL certificate.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Certificate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the certificate.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the certificate's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name of the certificate.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// The version of the encoded certificate.
	Version *int `mandatory:"true" json:"version"`

	// The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
	SignatureAlgorithm *string `mandatory:"true" json:"signatureAlgorithm"`

	// The date and time the certificate will become valid, expressed in RFC 3339 timestamp format.
	TimeNotValidBefore *common.SDKTime `mandatory:"true" json:"timeNotValidBefore"`

	// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
	TimeNotValidAfter *common.SDKTime `mandatory:"true" json:"timeNotValidAfter"`

	IssuedBy *string `mandatory:"false" json:"issuedBy"`

	SubjectName *CertificateSubjectName `mandatory:"false" json:"subjectName"`

	IssuerName *CertificateIssuerName `mandatory:"false" json:"issuerName"`

	PublicKeyInfo *CertificatePublicKeyInfo `mandatory:"false" json:"publicKeyInfo"`

	// Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
	Extensions []CertificateExtensions `mandatory:"false" json:"extensions"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current lifecycle state of the SSL certificate.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the certificate was created, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Certificate) String() string {
	return common.PointerString(m)
}

// CertificateLifecycleStateEnum is an alias to type: LifecycleStatesEnum
// Consider using LifecycleStatesEnum instead
// Deprecated
type CertificateLifecycleStateEnum = LifecycleStatesEnum

// Set of constants representing the allowable values for LifecycleStatesEnum
// Deprecated
const (
	CertificateLifecycleStateCreating LifecycleStatesEnum = "CREATING"
	CertificateLifecycleStateActive   LifecycleStatesEnum = "ACTIVE"
	CertificateLifecycleStateFailed   LifecycleStatesEnum = "FAILED"
	CertificateLifecycleStateUpdating LifecycleStatesEnum = "UPDATING"
	CertificateLifecycleStateDeleting LifecycleStatesEnum = "DELETING"
	CertificateLifecycleStateDeleted  LifecycleStatesEnum = "DELETED"
)

// GetCertificateLifecycleStateEnumValues Enumerates the set of values for LifecycleStatesEnum
// Consider using GetLifecycleStatesEnumValue
// Deprecated
var GetCertificateLifecycleStateEnumValues = GetLifecycleStatesEnumValues
