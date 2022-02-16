// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ProtectionCapability References an OCI-managed protection capability. Checks if HTTP requests/responses are malicious.
type ProtectionCapability struct {

	// Unique key of referenced protection capability.
	Key *string `mandatory:"true" json:"key"`

	// Version of referenced protection capability.
	Version *int `mandatory:"true" json:"version"`

	Exclusions *ProtectionCapabilityExclusions `mandatory:"false" json:"exclusions"`

	// Override action to take if capability was triggered, defined in Protection Rule for this capability.
	// Only actions of type CHECK are allowed.
	ActionName *string `mandatory:"false" json:"actionName"`

	// The minimum sum of weights of associated collaborative protection capabilities that have triggered which
	// must be reached in order for _this_ capability to trigger.
	// This field is ignored for non-collaborative capabilities.
	CollaborativeActionThreshold *int `mandatory:"false" json:"collaborativeActionThreshold"`

	// Explicit weight values to use for associated collaborative protection capabilities.
	CollaborativeWeights []CollaborativeCapabilityWeightOverride `mandatory:"false" json:"collaborativeWeights"`
}

func (m ProtectionCapability) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectionCapability) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
