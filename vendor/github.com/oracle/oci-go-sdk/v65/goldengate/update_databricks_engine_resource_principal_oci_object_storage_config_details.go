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

// UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails Resource-principal authentication and location details for updating an OCI Object Storage Databricks Engine artifact location.
type UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails struct {

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the current region.
	Region *string `mandatory:"false" json:"region"`

	// The OCI Object Storage namespace containing the bucket.
	// Example: "mynamespace"
	Namespace *string `mandatory:"false" json:"namespace"`

	// Bucket containing Databricks application artifacts.
	Bucket *string `mandatory:"false" json:"bucket"`

	// Object path used for Databricks application artifacts.
	Path *string `mandatory:"false" json:"path"`
}

func (m UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeUpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails
	}{
		"RESOURCE_PRINCIPAL",
		(MarshalTypeUpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails)(m),
	}

	return json.Marshal(&s)
}
