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

// UpdateZfsIcebergStorageDetails The information to update the ZFS storage used in the Iceberg connection for an Oracle AI Data Catalog Iceberg catalogs.
type UpdateZfsIcebergStorageDetails struct {

	// The access key ID used to access the ZFS storage.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the ZFS secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// The secret access key used to access the warehouse storage.
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the ZFS storage trust store.
	// Note: When provided, 'trustStore' field must not be provided.
	TrustStoreSecretId *string `mandatory:"false" json:"trustStoreSecretId"`

	// The base64 encoded content of the ZFS storage TrustStore file.
	// Deprecated: This field is deprecated and replaced by "trustStoreSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	TrustStore *string `mandatory:"false" json:"trustStore"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content
	// of the configuration file containing additional properties for the REST catalog.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PropertiesSecretId *string `mandatory:"false" json:"propertiesSecretId"`

	// The base64 encoded content of the configuration file containing additional properties for the REST catalog.
	Properties *string `mandatory:"false" json:"properties"`
}

func (m UpdateZfsIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateZfsIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateZfsIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateZfsIcebergStorageDetails UpdateZfsIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateZfsIcebergStorageDetails
	}{
		"ZFS",
		(MarshalTypeUpdateZfsIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
