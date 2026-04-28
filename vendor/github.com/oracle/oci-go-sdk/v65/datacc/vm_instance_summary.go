// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmInstanceSummary VM instance on Database Infrastructure summary.
type VmInstanceSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM instance.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM
	// instance.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores enabled for each VM instance.
	CpusEnabled *int `mandatory:"true" json:"cpusEnabled"`

	// The current state of the VM instance.
	LifecycleState VmInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time that the VM instance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the VM instance was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time zone to use for the VM instance.
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// VM instance display name. This name does not have to be unique, and is changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// VM instance description.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"false" json:"infrastructureId"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *float64 `mandatory:"false" json:"memorySizeInGBs"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM custom instance uploaded.
	ImageId *string `mandatory:"false" json:"imageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute node on which VM instance should be launched.
	ServerId *string `mandatory:"false" json:"serverId"`

	// Lifecycle state details of the VM instance.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The domain name of the VM instance.
	DomainName *string `mandatory:"false" json:"domainName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Network.
	VmNetworkId *string `mandatory:"false" json:"vmNetworkId"`
}

func (m VmInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
