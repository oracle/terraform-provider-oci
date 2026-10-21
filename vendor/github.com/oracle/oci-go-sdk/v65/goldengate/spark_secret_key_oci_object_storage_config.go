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

// SparkSecretKeyOciObjectStorageConfig Secret-key authentication and location details for an OCI Object Storage Spark artifact location.
type SparkSecretKeyOciObjectStorageConfig struct {

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

	// Access key identifier used for OCI Object Storage secret-key authentication.
	// Note: Despite the "Id" suffix, this value is not an OCI OCID.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the
	// secret access key. Note: When provided, `secretAccessKey` must not be provided.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`
}

func (m SparkSecretKeyOciObjectStorageConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SparkSecretKeyOciObjectStorageConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SparkSecretKeyOciObjectStorageConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSparkSecretKeyOciObjectStorageConfig SparkSecretKeyOciObjectStorageConfig
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeSparkSecretKeyOciObjectStorageConfig
	}{
		"CUSTOMER_SECRET_KEY",
		(MarshalTypeSparkSecretKeyOciObjectStorageConfig)(m),
	}

	return json.Marshal(&s)
}
