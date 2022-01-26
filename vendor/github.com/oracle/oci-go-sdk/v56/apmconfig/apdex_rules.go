// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ApdexRules The set of Apdex rules to be used in Apdex computation. In the current version, only one rule set can exist in the
// configuration. This may change in the future.
type ApdexRules struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID will be generated
	// when the item is created.
	Id *string `mandatory:"false" json:"id"`

	// The time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-13T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The name by which this rule set can be displayed to the user.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Rules []Apdex `mandatory:"false" json:"rules"`
}

//GetId returns Id
func (m ApdexRules) GetId() *string {
	return m.Id
}

//GetTimeCreated returns TimeCreated
func (m ApdexRules) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ApdexRules) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetFreeformTags returns FreeformTags
func (m ApdexRules) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ApdexRules) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m ApdexRules) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ApdexRules) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApdexRules ApdexRules
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeApdexRules
	}{
		"APDEX",
		(MarshalTypeApdexRules)(m),
	}

	return json.Marshal(&s)
}
