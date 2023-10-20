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

// RemediationRecipe An Application Dependency Management (ADM) remediation recipe contains the basic configuration and the
// details of each of the remediation stages (Detect, Recommend, Verify, and Apply).
type RemediationRecipe struct {

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Remediation Recipe.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The creation date and time of the Remediation Recipe (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Remediation Recipe was last updated (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the Remediation Recipe.
	LifecycleState RemediationRecipeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The compartment Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the remediation recipe.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"true" json:"knowledgeBaseId"`

	// Boolean indicating if a run should be automatically triggered once the Knowledge Base contents are updated.
	IsRunTriggeredOnKbChange *bool `mandatory:"true" json:"isRunTriggeredOnKbChange"`

	ScmConfiguration ScmConfiguration `mandatory:"true" json:"scmConfiguration"`

	VerifyConfiguration VerifyConfiguration `mandatory:"true" json:"verifyConfiguration"`

	DetectConfiguration *DetectConfiguration `mandatory:"true" json:"detectConfiguration"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RemediationRecipe) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemediationRecipe) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemediationRecipeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRemediationRecipeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RemediationRecipe) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SystemTags               map[string]map[string]interface{}   `json:"systemTags"`
		Id                       *string                             `json:"id"`
		DisplayName              *string                             `json:"displayName"`
		TimeCreated              *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated              *common.SDKTime                     `json:"timeUpdated"`
		LifecycleState           RemediationRecipeLifecycleStateEnum `json:"lifecycleState"`
		CompartmentId            *string                             `json:"compartmentId"`
		KnowledgeBaseId          *string                             `json:"knowledgeBaseId"`
		IsRunTriggeredOnKbChange *bool                               `json:"isRunTriggeredOnKbChange"`
		ScmConfiguration         scmconfiguration                    `json:"scmConfiguration"`
		VerifyConfiguration      verifyconfiguration                 `json:"verifyConfiguration"`
		DetectConfiguration      *DetectConfiguration                `json:"detectConfiguration"`
		NetworkConfiguration     *NetworkConfiguration               `json:"networkConfiguration"`
		FreeformTags             map[string]string                   `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}   `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.CompartmentId = model.CompartmentId

	m.KnowledgeBaseId = model.KnowledgeBaseId

	m.IsRunTriggeredOnKbChange = model.IsRunTriggeredOnKbChange

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// RemediationRecipeLifecycleStateEnum Enum with underlying type: string
type RemediationRecipeLifecycleStateEnum string

// Set of constants representing the allowable values for RemediationRecipeLifecycleStateEnum
const (
	RemediationRecipeLifecycleStateCreating       RemediationRecipeLifecycleStateEnum = "CREATING"
	RemediationRecipeLifecycleStateActive         RemediationRecipeLifecycleStateEnum = "ACTIVE"
	RemediationRecipeLifecycleStateUpdating       RemediationRecipeLifecycleStateEnum = "UPDATING"
	RemediationRecipeLifecycleStateInactive       RemediationRecipeLifecycleStateEnum = "INACTIVE"
	RemediationRecipeLifecycleStateFailed         RemediationRecipeLifecycleStateEnum = "FAILED"
	RemediationRecipeLifecycleStateDeleting       RemediationRecipeLifecycleStateEnum = "DELETING"
	RemediationRecipeLifecycleStateDeleted        RemediationRecipeLifecycleStateEnum = "DELETED"
	RemediationRecipeLifecycleStateNeedsAttention RemediationRecipeLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingRemediationRecipeLifecycleStateEnum = map[string]RemediationRecipeLifecycleStateEnum{
	"CREATING":        RemediationRecipeLifecycleStateCreating,
	"ACTIVE":          RemediationRecipeLifecycleStateActive,
	"UPDATING":        RemediationRecipeLifecycleStateUpdating,
	"INACTIVE":        RemediationRecipeLifecycleStateInactive,
	"FAILED":          RemediationRecipeLifecycleStateFailed,
	"DELETING":        RemediationRecipeLifecycleStateDeleting,
	"DELETED":         RemediationRecipeLifecycleStateDeleted,
	"NEEDS_ATTENTION": RemediationRecipeLifecycleStateNeedsAttention,
}

var mappingRemediationRecipeLifecycleStateEnumLowerCase = map[string]RemediationRecipeLifecycleStateEnum{
	"creating":        RemediationRecipeLifecycleStateCreating,
	"active":          RemediationRecipeLifecycleStateActive,
	"updating":        RemediationRecipeLifecycleStateUpdating,
	"inactive":        RemediationRecipeLifecycleStateInactive,
	"failed":          RemediationRecipeLifecycleStateFailed,
	"deleting":        RemediationRecipeLifecycleStateDeleting,
	"deleted":         RemediationRecipeLifecycleStateDeleted,
	"needs_attention": RemediationRecipeLifecycleStateNeedsAttention,
}

// GetRemediationRecipeLifecycleStateEnumValues Enumerates the set of values for RemediationRecipeLifecycleStateEnum
func GetRemediationRecipeLifecycleStateEnumValues() []RemediationRecipeLifecycleStateEnum {
	values := make([]RemediationRecipeLifecycleStateEnum, 0)
	for _, v := range mappingRemediationRecipeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRemediationRecipeLifecycleStateEnumStringValues Enumerates the set of values in String for RemediationRecipeLifecycleStateEnum
func GetRemediationRecipeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingRemediationRecipeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemediationRecipeLifecycleStateEnum(val string) (RemediationRecipeLifecycleStateEnum, bool) {
	enum, ok := mappingRemediationRecipeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
