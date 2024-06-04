// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Certificate Certificate data.
type Certificate struct {

	// The identifier key (unique name in the scope of the deployment) of the certificate being referenced.
	// It must be 1 to 32 characters long, must contain only alphanumeric characters and must start with a letter.
	Key *string `mandatory:"true" json:"key"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The base64 encoded content of the PEM file containing the SSL certificate.
	CertificateContent *string `mandatory:"true" json:"certificateContent"`

	// The Certificate issuer.
	Issuer *string `mandatory:"true" json:"issuer"`

	// Indicates if the certificate is self signed.
	IsSelfSigned *bool `mandatory:"true" json:"isSelfSigned"`

	// The Certificate md5Hash.
	Md5Hash *string `mandatory:"true" json:"md5Hash"`

	// The Certificate public key.
	PublicKey *string `mandatory:"true" json:"publicKey"`

	// The Certificate public key algorithm.
	PublicKeyAlgorithm *string `mandatory:"true" json:"publicKeyAlgorithm"`

	// The Certificate public key size.
	PublicKeySize *int64 `mandatory:"true" json:"publicKeySize"`

	// The Certificate serial.
	Serial *string `mandatory:"true" json:"serial"`

	// The Certificate subject.
	Subject *string `mandatory:"true" json:"subject"`

	// The time the certificate is valid from. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeValidFrom *common.SDKTime `mandatory:"true" json:"timeValidFrom"`

	// The time the certificate is valid to. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeValidTo *common.SDKTime `mandatory:"true" json:"timeValidTo"`

	// The Certificate version.
	Version *string `mandatory:"true" json:"version"`

	// The Certificate sha1 hash.
	Sha1Hash *string `mandatory:"true" json:"sha1Hash"`

	// The Certificate authority key id.
	AuthorityKeyId *string `mandatory:"true" json:"authorityKeyId"`

	// Indicates if the certificate is ca.
	IsCa *bool `mandatory:"true" json:"isCa"`

	// The Certificate subject key id.
	SubjectKeyId *string `mandatory:"true" json:"subjectKeyId"`

	// Possible certificate lifecycle states.
	LifecycleState CertificateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m Certificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Certificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCertificateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCertificateLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
