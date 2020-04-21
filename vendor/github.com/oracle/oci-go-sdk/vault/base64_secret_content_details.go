// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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
