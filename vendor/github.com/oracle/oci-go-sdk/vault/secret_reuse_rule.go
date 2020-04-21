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

// SecretReuseRule A rule that disallows reuse of previously used secret content by the specified secret.
type SecretReuseRule struct {

	// A property indicating whether the rule is applied even if the secret version with the content you are trying to reuse was deleted.
	IsEnforcedOnDeletedSecretVersions *bool `mandatory:"false" json:"isEnforcedOnDeletedSecretVersions"`
}

func (m SecretReuseRule) String() string {
	return common.PointerString(m)
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
