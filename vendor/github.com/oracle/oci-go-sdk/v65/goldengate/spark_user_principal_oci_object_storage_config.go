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

// SparkUserPrincipalOciObjectStorageConfig User-principal authentication and location details for an OCI Object Storage Spark artifact location.
// The private key is represented by `privateKeyFileSecretId`; the private key content is
// never returned.
type SparkUserPrincipalOciObjectStorageConfig struct {

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the current region.
	Region *string `mandatory:"true" json:"region"`

	// The OCI Object Storage namespace containing the bucket.
	// Example: "mynamespace"
	Namespace *string `mandatory:"true" json:"namespace"`

	// Bucket containing Spark application artifacts.
	// Example: "spark-artifacts"
	Bucket *string `mandatory:"true" json:"bucket"`

	// Object path used for Spark application artifacts.
	// Example: "spark-apps/order-analysis/1.2.0"
	Path *string `mandatory:"true" json:"path"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI user whose API key
	// is used to access the artifact location.
	// If not provided it will default to the current user.
	UserId *string `mandatory:"true" json:"userId"`

	// The fingerprint of the API Key of the user specified by the userId.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PublicKeyFingerprint *string `mandatory:"true" json:"publicKeyFingerprint"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related OCI tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content of the private key file (PEM file) corresponding to the API key of the fingerprint.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	// Note: When provided, 'privateKeyFile' field must not be provided.
	PrivateKeyFileSecretId *string `mandatory:"false" json:"privateKeyFileSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the passphrase of the private key.
	// Note: When provided, 'privateKeyPassphrase' field must not be provided.
	PrivateKeyPassphraseSecretId *string `mandatory:"false" json:"privateKeyPassphraseSecretId"`
}

func (m SparkUserPrincipalOciObjectStorageConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SparkUserPrincipalOciObjectStorageConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SparkUserPrincipalOciObjectStorageConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSparkUserPrincipalOciObjectStorageConfig SparkUserPrincipalOciObjectStorageConfig
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeSparkUserPrincipalOciObjectStorageConfig
	}{
		"USER_PRINCIPAL",
		(MarshalTypeSparkUserPrincipalOciObjectStorageConfig)(m),
	}

	return json.Marshal(&s)
}
