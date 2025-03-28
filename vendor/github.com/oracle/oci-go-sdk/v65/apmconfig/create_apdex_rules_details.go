// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateApdexRulesDetails The set of Apdex rules to be used in Apdex computation. In the current version, only one rule set may exist per
// configuration, and attempting to create a rule set if it already exists results in an error.
type CreateApdexRulesDetails struct {
	Rules []Apdex `mandatory:"true" json:"rules"`

	// The name by which a configuration entity is displayed to the end user.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetFreeformTags returns FreeformTags
func (m CreateApdexRulesDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateApdexRulesDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateApdexRulesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateApdexRulesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateApdexRulesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateApdexRulesDetails CreateApdexRulesDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateApdexRulesDetails
	}{
		"APDEX",
		(MarshalTypeCreateApdexRulesDetails)(m),
	}

	return json.Marshal(&s)
}
