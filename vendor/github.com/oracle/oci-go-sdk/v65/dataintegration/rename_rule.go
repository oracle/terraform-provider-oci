// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RenameRule Lets you rename an attribute.
type RenameRule struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Specifies whether the rule uses a java regex syntax.
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether to skip remaining rules when a match is found.
	IsSkipRemainingRulesOnMatch *bool `mandatory:"false" json:"isSkipRemainingRulesOnMatch"`

	// The attribute name that needs to be renamed.
	FromName *string `mandatory:"false" json:"fromName"`

	// The new attribute name.
	ToName *string `mandatory:"false" json:"toName"`
}

// GetKey returns Key
func (m RenameRule) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m RenameRule) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m RenameRule) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetIsJavaRegexSyntax returns IsJavaRegexSyntax
func (m RenameRule) GetIsJavaRegexSyntax() *bool {
	return m.IsJavaRegexSyntax
}

// GetConfigValues returns ConfigValues
func (m RenameRule) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m RenameRule) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m RenameRule) GetDescription() *string {
	return m.Description
}

func (m RenameRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RenameRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RenameRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRenameRule RenameRule
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeRenameRule
	}{
		"RENAME_RULE",
		(MarshalTypeRenameRule)(m),
	}

	return json.Marshal(&s)
}
