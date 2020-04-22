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

// SecretExpiryRule A rule that helps enforce the expiration of a secret's contents.
type SecretExpiryRule struct {

	// A property indicating how long the secret contents will be considered valid, expressed in
	// ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. The secret needs to be
	// updated when the secret content expires. No enforcement mechanism exists at this time, but audit logs
	// record the expiration on the appropriate date, according to the time interval specified in the rule.
	// The timer resets after you update the secret contents.
	// The minimum value is 1 day and the maximum value is 90 days for this property. Currently, only intervals expressed in days are supported.
	// For example, pass `P3D` to have the secret version expire every 3 days.
	SecretVersionExpiryInterval *string `mandatory:"false" json:"secretVersionExpiryInterval"`

	// An optional property indicating the absolute time when this secret will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// The minimum number of days from current time is 1 day and the maximum number of days from current time is 365 days.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfAbsoluteExpiry *common.SDKTime `mandatory:"false" json:"timeOfAbsoluteExpiry"`

	// A property indicating whether to block retrieval of the secret content, on expiry. The default is false.
	// If the secret has already expired and you would like to retrieve the secret contents,
	// you need to edit the secret rule to disable this property, to allow reading the secret content.
	IsSecretContentRetrievalBlockedOnExpiry *bool `mandatory:"false" json:"isSecretContentRetrievalBlockedOnExpiry"`
}

func (m SecretExpiryRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SecretExpiryRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSecretExpiryRule SecretExpiryRule
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypeSecretExpiryRule
	}{
		"SECRET_EXPIRY_RULE",
		(MarshalTypeSecretExpiryRule)(m),
	}

	return json.Marshal(&s)
}
