// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Retrieval API
//
// Use the Secret Retrieval API to retrieve secrets and secret versions from vaults. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package secrets

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Base64SecretBundleContentDetails The contents of the secret.
type Base64SecretBundleContentDetails struct {

	// The base64-encoded content of the secret.
	Content *string `mandatory:"false" json:"content"`
}

func (m Base64SecretBundleContentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Base64SecretBundleContentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Base64SecretBundleContentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBase64SecretBundleContentDetails Base64SecretBundleContentDetails
	s := struct {
		DiscriminatorParam string `json:"contentType"`
		MarshalTypeBase64SecretBundleContentDetails
	}{
		"BASE64",
		(MarshalTypeBase64SecretBundleContentDetails)(m),
	}

	return json.Marshal(&s)
}
