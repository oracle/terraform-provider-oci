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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CollaborativeCapabilityWeightOverride Collaborative capability key and overriding weight.
type CollaborativeCapabilityWeightOverride struct {

	// Unique key of collaborative capability for which weight will be overridden.
	Key *string `mandatory:"true" json:"key"`

	// The value of weight to set.
	Weight *int `mandatory:"true" json:"weight"`
}

func (m CollaborativeCapabilityWeightOverride) String() string {
	return common.PointerString(m)
}
