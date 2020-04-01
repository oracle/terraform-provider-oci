// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
