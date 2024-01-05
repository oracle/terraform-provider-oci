// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociatedResourcesSummary The information about monitored resource.
type AssociatedResourcesSummary struct {

	// Monitored resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// Monitored Resource Name.
	Name *string `mandatory:"false" json:"name"`

	// Monitored resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Monitored Resource Type.
	Type *string `mandatory:"false" json:"type"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Monitored Resource Host Name.
	HostName *string `mandatory:"false" json:"hostName"`

	// External resource is any OCI resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	// which is not a Stack Monitoring service resource.
	// Currently supports only following resource types - Container database, non-container database,
	// pluggable database and OCI compute instance.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The current state of the monitored resource.
	LifecycleState ResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// License edition of the monitored resource.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	// List of associated monitored resources.
	AssociatedResources []AssociatedMonitoredResource `mandatory:"false" json:"associatedResources"`
}

func (m AssociatedResourcesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedResourcesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
