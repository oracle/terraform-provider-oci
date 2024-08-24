// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HsmPartitionSummary An object which encapsulates the details of a given HSM.
type HsmPartitionSummary struct {

	// The OCID of the HSM resource. Each HSM resource will have a unique OCID identifier.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the HSMCluster that contains a particular HSM resource.
	HsmClusterId *string `mandatory:"true" json:"hsmClusterId"`

	// A HSMCluster resource's current lifecycle state.
	// Example: `ACTIVE`
	LifecycleState HsmPartitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the compartment that contains a particular HSM resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Details of a single portInformation item include the PortNumber (an integer used as an identifier) and the PortType (refers to either an enum value of Managementutility,Clientutility, or null)
	PortInformation []PortInformation `mandatory:"true" json:"portInformation"`

	// The date and time an HSM was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time an HSM was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m HsmPartitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HsmPartitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHsmPartitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHsmPartitionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
