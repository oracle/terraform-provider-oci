// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSparkUserPrincipalOciObjectStorageConfigDetails User-principal authentication and location details for creating an OCI Object Storage Spark artifact location.
// Provide exactly one of `privateKeyFile` or `privateKeyFileSecretId`. If the private key
// is passphrase-protected, provide at most one of `privateKeyPassphrase` or
// `privateKeyPassphraseSecretId`.
type CreateSparkUserPrincipalOciObjectStorageConfigDetails struct {

	// The OCI Object Storage namespace containing the bucket.
	// Example: "mynamespace"
	Namespace *string `mandatory:"true" json:"namespace"`

	// Bucket containing Spark application artifacts.
	// Example: "spark-artifacts"
	Bucket *string `mandatory:"true" json:"bucket"`

	// Object path used for Spark application artifacts.
	// Example: "spark-apps/order-analysis/1.2.0"
	Path *string `mandatory:"true" json:"path"`

	// The fingerprint of the API Key of the user specified by the userId.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PublicKeyFingerprint *string `mandatory:"true" json:"publicKeyFingerprint"`

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the current region.
	Region *string `mandatory:"false" json:"region"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related OCI tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI user whose API key
	// is used to access the artifact location.
	// If not provided it will default to the current user.
	UserId *string `mandatory:"false" json:"userId"`

	// The base64 encoded content of the private key file (PEM file) corresponding to the API key of the fingerprint.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	// Deprecated: This field is deprecated and replaced by "privateKeyFileSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	PrivateKeyFile *string `mandatory:"false" json:"privateKeyFile" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	// Note: When provided, 'privateKeyFile' field must not be provided.
	PrivateKeyFileSecretId *string `mandatory:"false" json:"privateKeyFileSecretId"`

	// The passphrase of the private key.
	// Deprecated: This field is deprecated and replaced by "privateKeyPassphraseSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	PrivateKeyPassphrase *string `mandatory:"false" json:"privateKeyPassphrase" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the passphrase of the private key.
	// Note: When provided, 'privateKeyPassphrase' field must not be provided.
	PrivateKeyPassphraseSecretId *string `mandatory:"false" json:"privateKeyPassphraseSecretId"`
}

func (m CreateSparkUserPrincipalOciObjectStorageConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSparkUserPrincipalOciObjectStorageConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSparkUserPrincipalOciObjectStorageConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSparkUserPrincipalOciObjectStorageConfigDetails CreateSparkUserPrincipalOciObjectStorageConfigDetails
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeCreateSparkUserPrincipalOciObjectStorageConfigDetails
	}{
		"USER_PRINCIPAL",
		(MarshalTypeCreateSparkUserPrincipalOciObjectStorageConfigDetails)(m),
	}

	return json.Marshal(&s)
}
