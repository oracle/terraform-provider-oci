// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArtifactVersionAuditSummary Artifact Version Audit summary, contains basic details of the audit.
type ArtifactVersionAuditSummary struct {

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Artifact Version Audit.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"true" json:"knowledgeBaseId"`

	// The creation date and time of the Artifact Version Audit (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The update date and time of the Artifact Version Audit (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the Artifact Version Audit.
	LifecycleState ArtifactVersionAuditLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The compartment Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Artifact Version Audit.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the Artifact Version Audit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ArtifactVersionAuditSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArtifactVersionAuditSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactVersionAuditLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetArtifactVersionAuditLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
