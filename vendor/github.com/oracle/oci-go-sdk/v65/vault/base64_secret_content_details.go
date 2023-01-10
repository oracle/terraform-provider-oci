// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Base64SecretContentDetails Base64-encoded secret content.
type Base64SecretContentDetails struct {

	// Names should be unique within a secret. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"false" json:"name"`

	// The base64-encoded content of the secret.
	Content *string `mandatory:"false" json:"content"`

	// The rotation state of the secret content. The default is `CURRENT`, meaning that the secret is currently in use. A secret version
	// that you mark as `PENDING` is staged and available for use, but you don't yet want to rotate it into current, active use. For example,
	// you might create or update a secret and mark its rotation state as `PENDING` if you haven't yet updated the secret on the target system.
	// When creating a secret, only the value `CURRENT` is applicable, although the value `LATEST` is also automatically applied. When updating
	// a secret, you can specify a version's rotation state as either `CURRENT` or `PENDING`.
	Stage SecretContentDetailsStageEnum `mandatory:"false" json:"stage,omitempty"`
}

//GetName returns Name
func (m Base64SecretContentDetails) GetName() *string {
	return m.Name
}

//GetStage returns Stage
func (m Base64SecretContentDetails) GetStage() SecretContentDetailsStageEnum {
	return m.Stage
}

func (m Base64SecretContentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Base64SecretContentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecretContentDetailsStageEnum(string(m.Stage)); !ok && m.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", m.Stage, strings.Join(GetSecretContentDetailsStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Base64SecretContentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBase64SecretContentDetails Base64SecretContentDetails
	s := struct {
		DiscriminatorParam string `json:"contentType"`
		MarshalTypeBase64SecretContentDetails
	}{
		"BASE64",
		(MarshalTypeBase64SecretContentDetails)(m),
	}

	return json.Marshal(&s)
}
