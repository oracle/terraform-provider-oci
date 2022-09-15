// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddonSummary The properties that define a cluster addons summary.
type AddonSummary struct {

	// The name of the cluster addon.
	Name *string `mandatory:"true" json:"name"`

	// selected addon version, or null indicates autoUpdate
	Version *string `mandatory:"true" json:"version"`

	// current installed version of the addon
	CurrentInstalledVersion *string `mandatory:"true" json:"currentInstalledVersion"`

	// The state of the cluster addons.
	LifecycleState AddonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The error info of the addon.
	AddonError *AddonError `mandatory:"false" json:"addonError"`
}

func (m AddonSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddonSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAddonLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
