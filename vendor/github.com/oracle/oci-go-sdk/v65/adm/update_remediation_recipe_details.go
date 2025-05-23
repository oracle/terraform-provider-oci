// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateRemediationRecipeDetails Details to update an existing remediation recipe.
type UpdateRemediationRecipeDetails struct {

	// The name of the remediation recipe.
	DisplayName *string `mandatory:"false" json:"displayName"`

	ScmConfiguration ScmConfiguration `mandatory:"false" json:"scmConfiguration"`

	VerifyConfiguration VerifyConfiguration `mandatory:"false" json:"verifyConfiguration"`

	DetectConfiguration *DetectConfiguration `mandatory:"false" json:"detectConfiguration"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"false" json:"networkConfiguration"`

	// The Oracle Cloud Identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"false" json:"knowledgeBaseId"`

	// Boolean indicating if a run should be automatically triggered once the knowledge base is updated.
	IsRunTriggeredOnKbChange *bool `mandatory:"false" json:"isRunTriggeredOnKbChange"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateRemediationRecipeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateRemediationRecipeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateRemediationRecipeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		ScmConfiguration         scmconfiguration                  `json:"scmConfiguration"`
		VerifyConfiguration      verifyconfiguration               `json:"verifyConfiguration"`
		DetectConfiguration      *DetectConfiguration              `json:"detectConfiguration"`
		NetworkConfiguration     *NetworkConfiguration             `json:"networkConfiguration"`
		KnowledgeBaseId          *string                           `json:"knowledgeBaseId"`
		IsRunTriggeredOnKbChange *bool                             `json:"isRunTriggeredOnKbChange"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.ScmConfiguration.UnmarshalPolymorphicJSON(model.ScmConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScmConfiguration = nn.(ScmConfiguration)
	} else {
		m.ScmConfiguration = nil
	}

	nn, e = model.VerifyConfiguration.UnmarshalPolymorphicJSON(model.VerifyConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.VerifyConfiguration = nn.(VerifyConfiguration)
	} else {
		m.VerifyConfiguration = nil
	}

	m.DetectConfiguration = model.DetectConfiguration

	m.NetworkConfiguration = model.NetworkConfiguration

	m.KnowledgeBaseId = model.KnowledgeBaseId

	m.IsRunTriggeredOnKbChange = model.IsRunTriggeredOnKbChange

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
