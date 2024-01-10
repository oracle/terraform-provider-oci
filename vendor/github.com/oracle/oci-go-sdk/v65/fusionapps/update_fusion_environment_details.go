// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateFusionEnvironmentDetails The information to be updated.
type UpdateFusionEnvironmentDetails struct {

	// FusionEnvironment Identifier, can be renamed
	DisplayName *string `mandatory:"false" json:"displayName"`

	// byok kms keyId
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	MaintenancePolicy *MaintenancePolicy `mandatory:"false" json:"maintenancePolicy"`

	// Language packs
	AdditionalLanguagePacks []string `mandatory:"false" json:"additionalLanguagePacks"`

	// Network access control rules to limit internet traffic that can access the environment. For more information, see AllowRule.
	Rules []Rule `mandatory:"false" json:"rules"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateFusionEnvironmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateFusionEnvironmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateFusionEnvironmentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		KmsKeyId                *string                           `json:"kmsKeyId"`
		MaintenancePolicy       *MaintenancePolicy                `json:"maintenancePolicy"`
		AdditionalLanguagePacks []string                          `json:"additionalLanguagePacks"`
		Rules                   []rule                            `json:"rules"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.KmsKeyId = model.KmsKeyId

	m.MaintenancePolicy = model.MaintenancePolicy

	m.AdditionalLanguagePacks = make([]string, len(model.AdditionalLanguagePacks))
	copy(m.AdditionalLanguagePacks, model.AdditionalLanguagePacks)
	m.Rules = make([]Rule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(Rule)
		} else {
			m.Rules[i] = nil
		}
	}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
