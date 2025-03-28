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

// RemediationRunSummary remediation run summary.
type RemediationRunSummary struct {

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle Cloud identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the Remediation Recipe.
	RemediationRecipeId *string `mandatory:"true" json:"remediationRecipeId"`

	// The creation date and time of the remediation run (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the remediation run was last updated (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The compartment Oracle Cloud Identifier (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the remediation run.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the remediation run.
	LifecycleState RemediationRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The source from which the remediation run was triggered.
	RemediationRunSource RemediationRunRemediationRunSourceEnum `mandatory:"true" json:"remediationRunSource"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The name of the remediation run.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time of the start of the remediation run (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time of the finish of the remediation run (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The type of the current stage of the remediation run.
	CurrentStageType RemediationRunStageTypeEnum `mandatory:"false" json:"currentStageType,omitempty"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RemediationRunSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemediationRunSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemediationRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRemediationRunLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRunRemediationRunSourceEnum(string(m.RemediationRunSource)); !ok && m.RemediationRunSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RemediationRunSource: %s. Supported values are: %s.", m.RemediationRunSource, strings.Join(GetRemediationRunRemediationRunSourceEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRemediationRunStageTypeEnum(string(m.CurrentStageType)); !ok && m.CurrentStageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentStageType: %s. Supported values are: %s.", m.CurrentStageType, strings.Join(GetRemediationRunStageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
