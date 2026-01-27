// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecretDetail Details of the kubernetes secrets to be created or updated.
type SecretDetail struct {

	// Name of the kubernetes secret of max length 63 and contain only lowercase alphanumeric characters or '-' and start and end with an alphabetic character.
	SecretName *string `mandatory:"true" json:"secretName"`

	// List of kubernetes secret data.
	SecretData []SecretData `mandatory:"true" json:"secretData"`

	// Object storage path for the secret template to be used for creating secret otherwise it will be created with default template.
	TemplateObjectStoragePath *string `mandatory:"false" json:"templateObjectStoragePath"`
}

func (m SecretDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecretDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
