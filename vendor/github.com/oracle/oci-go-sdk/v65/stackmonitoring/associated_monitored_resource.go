// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AssociatedMonitoredResource The information about monitored resource.
type AssociatedMonitoredResource struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of monitored resource.
	Id *string `mandatory:"true" json:"id"`

	// Name of the monitored resource
	Name *string `mandatory:"false" json:"name"`

	// Monitored resource display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of the monitored resource
	Type *string `mandatory:"false" json:"type"`

	// Resource Host Name
	HostName *string `mandatory:"false" json:"hostName"`

	// External resource is any OCI resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	// which is not a Stack Monitoring service resource.
	// Currently supports only following resource type identifiers - externalcontainerdatabase,
	// externalnoncontainerdatabase, externalpluggabledatabase and OCI compute instance.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Management Agent Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The current state of the monitored resource.
	LifecycleState ResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Association details of the resource
	Association *interface{} `mandatory:"false" json:"association"`
}

func (m AssociatedMonitoredResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedMonitoredResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
