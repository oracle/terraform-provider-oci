// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SecretReuseRule A rule that disallows reuse of previously used secret content by the specified secret.
type SecretReuseRule struct {

	// A property indicating whether the rule is applied even if the secret version with the content you are trying to reuse was deleted.
	IsEnforcedOnDeletedSecretVersions *bool `mandatory:"false" json:"isEnforcedOnDeletedSecretVersions"`
}

func (m SecretReuseRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecretReuseRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SecretReuseRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSecretReuseRule SecretReuseRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeSecretReuseRule
	}{
		"SECRET_REUSE_RULE",
		(MarshalTypeSecretReuseRule)(m),
	}

	return json.Marshal(&s)
}
