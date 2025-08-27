// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceInfo The Resource Info.
type ResourceInfo struct {

	// The compartmentId of the resource.
	ResourceCompartmentId *string `mandatory:"true" json:"resourceCompartmentId"`

	// The name of the resource.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The Id of the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The type of the resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The deployment type of the resource.
	DeploymentType *string `mandatory:"true" json:"deploymentType"`

	// The connector Id of the resource.
	ConnectorId *string `mandatory:"true" json:"connectorId"`

	// host info objects
	HostInfo []HostInfo `mandatory:"true" json:"hostInfo"`

	// The version of the resource.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The platform type of the resource.
	DbPlatformType *string `mandatory:"true" json:"dbPlatformType"`

	// The License Type of the resource.
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// True if it is a cluster db.
	IsClusterDb *bool `mandatory:"false" json:"isClusterDb"`

	// The agent Id of the agent managing the resource.
	AgentId *string `mandatory:"false" json:"agentId"`
}

func (m ResourceInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
