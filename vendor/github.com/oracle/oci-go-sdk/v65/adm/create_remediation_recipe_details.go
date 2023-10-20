// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRemediationRecipeDetails Details to create a new Remediation Recipe.
type CreateRemediationRecipeDetails struct {

	// The compartment Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ScmConfiguration ScmConfiguration `mandatory:"true" json:"scmConfiguration"`

	VerifyConfiguration VerifyConfiguration `mandatory:"true" json:"verifyConfiguration"`

	DetectConfiguration *DetectConfiguration `mandatory:"true" json:"detectConfiguration"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"true" json:"knowledgeBaseId"`

	// Boolean indicating if a run should be automatically triggered once the knowledge base is updated.
	IsRunTriggeredOnKbChange *bool `mandatory:"true" json:"isRunTriggeredOnKbChange"`

	// The name of the remediation recipe.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateRemediationRecipeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRemediationRecipeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateRemediationRecipeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId            *string                           `json:"compartmentId"`
		ScmConfiguration         scmconfiguration                  `json:"scmConfiguration"`
		VerifyConfiguration      verifyconfiguration               `json:"verifyConfiguration"`
		DetectConfiguration      *DetectConfiguration              `json:"detectConfiguration"`
		NetworkConfiguration     *NetworkConfiguration             `json:"networkConfiguration"`
		KnowledgeBaseId          *string                           `json:"knowledgeBaseId"`
		IsRunTriggeredOnKbChange *bool                             `json:"isRunTriggeredOnKbChange"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

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

	return
}
