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

// CreateZfsWarehouseStorageDetails The information about a new ZFS warehouse storage configuration for Oracle AI Data Catalog Storage connection.
type CreateZfsWarehouseStorageDetails struct {

	// The access key ID used to access the ZFS storage.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// The endpoint URL of the ZFS storage service.
	EndpointUrl *string `mandatory:"true" json:"endpointUrl"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the ZFS secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// The secret access key used to access the warehouse storage.
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`

	// The base64 encoded content of the ZFS storage TrustStore file.
	// Deprecated: This field is deprecated and replaced by "trustStoreSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	TrustStore *string `mandatory:"false" json:"trustStore"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the ZFS storage trust store.
	// Note: When provided, 'trustStore' field must not be provided.
	TrustStoreSecretId *string `mandatory:"false" json:"trustStoreSecretId"`
}

func (m CreateZfsWarehouseStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateZfsWarehouseStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateZfsWarehouseStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZfsWarehouseStorageDetails CreateZfsWarehouseStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeCreateZfsWarehouseStorageDetails
	}{
		"ZFS",
		(MarshalTypeCreateZfsWarehouseStorageDetails)(m),
	}

	return json.Marshal(&s)
}
