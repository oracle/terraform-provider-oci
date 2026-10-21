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

// CreateSparkSecretKeyOciObjectStorageConfigDetails Secret-key authentication and location details for creating an OCI Object Storage Spark artifact location.
// Provide exactly one of `secretAccessKey` or `secretAccessKeySecretId`.
type CreateSparkSecretKeyOciObjectStorageConfigDetails struct {

	// The OCI Object Storage namespace containing the bucket.
	// Example: "mynamespace"
	Namespace *string `mandatory:"true" json:"namespace"`

	// Bucket containing Spark application artifacts.
	// Example: "spark-artifacts"
	Bucket *string `mandatory:"true" json:"bucket"`

	// Object path used for Spark application artifacts.
	// Example: "spark-apps/order-analysis/1.2.0"
	Path *string `mandatory:"true" json:"path"`

	// Access key identifier used for OCI Object Storage secret-key authentication.
	// Note: Despite the "Id" suffix, this value is not an OCI OCID.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the current region.
	Region *string `mandatory:"false" json:"region"`

	// Secret access key used for OCI Object Storage secret-key authentication.
	// Deprecated: This field is deprecated and replaced by "secretAccessKeySecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the
	// secret access key. Note: When provided, `secretAccessKey` must not be provided.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`
}

func (m CreateSparkSecretKeyOciObjectStorageConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSparkSecretKeyOciObjectStorageConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSparkSecretKeyOciObjectStorageConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSparkSecretKeyOciObjectStorageConfigDetails CreateSparkSecretKeyOciObjectStorageConfigDetails
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeCreateSparkSecretKeyOciObjectStorageConfigDetails
	}{
		"CUSTOMER_SECRET_KEY",
		(MarshalTypeCreateSparkSecretKeyOciObjectStorageConfigDetails)(m),
	}

	return json.Marshal(&s)
}
