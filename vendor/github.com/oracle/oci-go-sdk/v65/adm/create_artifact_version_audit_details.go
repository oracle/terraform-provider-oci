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

// CreateArtifactVersionAuditDetails Details to create an Artifact Version Audit.
type CreateArtifactVersionAuditDetails struct {

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"true" json:"knowledgeBaseId"`

	// The compartment Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Artifact Version Audit.
	// If compartment identifier is not provided the compartment of the associated knowledge base will be used instead.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// List of artifact version.
	ArtifactVersions []AuditArtifactVersion `mandatory:"false" json:"artifactVersions"`

	LifecyclePolicy AuditLifecyclePolicy `mandatory:"false" json:"lifecyclePolicy"`

	// The name of the Artifact Version Audit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateArtifactVersionAuditDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateArtifactVersionAuditDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateArtifactVersionAuditDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId    *string                           `json:"compartmentId"`
		ArtifactVersions []AuditArtifactVersion            `json:"artifactVersions"`
		LifecyclePolicy  auditlifecyclepolicy              `json:"lifecyclePolicy"`
		DisplayName      *string                           `json:"displayName"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		KnowledgeBaseId  *string                           `json:"knowledgeBaseId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.ArtifactVersions = make([]AuditArtifactVersion, len(model.ArtifactVersions))
	copy(m.ArtifactVersions, model.ArtifactVersions)
	nn, e = model.LifecyclePolicy.UnmarshalPolymorphicJSON(model.LifecyclePolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LifecyclePolicy = nn.(AuditLifecyclePolicy)
	} else {
		m.LifecyclePolicy = nil
	}

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.KnowledgeBaseId = model.KnowledgeBaseId

	return
}
