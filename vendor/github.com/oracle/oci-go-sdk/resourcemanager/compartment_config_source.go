// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CompartmentConfigSource Compartment to use for creating the stack.
// The new stack will include definitions for supported resource types in this compartment.
type CompartmentConfigSource struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to use
	// for creating the stack. The new stack will include definitions for supported
	// resource types in this compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The region to use for creating the stack. The new stack will include definitions for
	// supported resource types in this region.
	Region *string `mandatory:"true" json:"region"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// Filter for services to use with Resource Discovery (https://www.terraform.io/docs/providers/oci/guides/resource_discovery.html#services).
	// For example, "database" limits resource discovery to resource types within the Database service.
	// The specified services must be in scope of the given compartment OCID (tenancy level for root compartment, compartment level otherwise).
	// If not specified, then all services at the scope of the given compartment OCID are used.
	ServicesToDiscover []string `mandatory:"false" json:"servicesToDiscover"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CompartmentConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CompartmentConfigSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CompartmentConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompartmentConfigSource CompartmentConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCompartmentConfigSource
	}{
		"COMPARTMENT_CONFIG_SOURCE",
		(MarshalTypeCompartmentConfigSource)(m),
	}

	return json.Marshal(&s)
}
