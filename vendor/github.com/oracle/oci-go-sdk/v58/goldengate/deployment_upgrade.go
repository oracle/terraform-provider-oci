// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DeploymentUpgrade A container for your OCI GoldenGate Upgrade information.
type DeploymentUpgrade struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment upgrade being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The type of the deployment upgrade: MANUAL or AUTOMATIC
	DeploymentUpgradeType DeploymentUpgradeTypeEnum `mandatory:"true" json:"deploymentUpgradeType"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the request was started. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the request was finished. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Version of OGG
	OggVersion *string `mandatory:"false" json:"oggVersion"`

	// The time the resource was created. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Possible GGS lifecycle sub-states.
	LifecycleSubState LifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DeploymentUpgrade) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentUpgrade) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentUpgradeTypeEnum(string(m.DeploymentUpgradeType)); !ok && m.DeploymentUpgradeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentUpgradeType: %s. Supported values are: %s.", m.DeploymentUpgradeType, strings.Join(GetDeploymentUpgradeTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
