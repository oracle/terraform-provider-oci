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

// CreateDatabricksEngineParOciObjectStorageConfigDetails Pre-authenticated request URL details for creating an OCI Object Storage Databricks Engine artifact location.
// Provide exactly one of `parUrl` or `parUrlSecretId`.
type CreateDatabricksEngineParOciObjectStorageConfigDetails struct {

	// A pre-authenticated request (PAR) URL that provides access to the artifact location.
	// Deprecated: This field is deprecated and replaced by "parUrlSecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	ParUrl *string `mandatory:"false" json:"parUrl" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the
	// pre-authenticated request (PAR) URL. Note: When provided, `parUrl` must not be provided.
	ParUrlSecretId *string `mandatory:"false" json:"parUrlSecretId"`
}

func (m CreateDatabricksEngineParOciObjectStorageConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabricksEngineParOciObjectStorageConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabricksEngineParOciObjectStorageConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabricksEngineParOciObjectStorageConfigDetails CreateDatabricksEngineParOciObjectStorageConfigDetails
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeCreateDatabricksEngineParOciObjectStorageConfigDetails
	}{
		"PRE_AUTHENTICATED_REQUEST",
		(MarshalTypeCreateDatabricksEngineParOciObjectStorageConfigDetails)(m),
	}

	return json.Marshal(&s)
}
