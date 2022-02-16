// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ApplicationSummary Summary of an application.
type ApplicationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the application.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The display name of the application. The display name is unique within the compartment containing the application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the application.
	LifecycleState ApplicationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application.
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Network Security Groups to add the application to.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	TraceConfig *ApplicationTraceConfig `mandatory:"false" json:"traceConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The time the application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the application was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	ImagePolicyConfig *ImagePolicyConfig `mandatory:"false" json:"imagePolicyConfig"`
}

func (m ApplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplicationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
