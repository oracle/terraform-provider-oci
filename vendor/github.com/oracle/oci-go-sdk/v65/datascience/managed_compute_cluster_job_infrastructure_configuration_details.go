// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedComputeClusterJobInfrastructureConfigurationDetails The job infrastructure configuration for compute target.
type ManagedComputeClusterJobInfrastructureConfigurationDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute target.
	ComputeTargetId *string `mandatory:"true" json:"computeTargetId"`

	ResourceConfiguration *ManagedComputeClusterJobResourceConfiguration `mandatory:"true" json:"resourceConfiguration"`
}

func (m ManagedComputeClusterJobInfrastructureConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedComputeClusterJobInfrastructureConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedComputeClusterJobInfrastructureConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedComputeClusterJobInfrastructureConfigurationDetails ManagedComputeClusterJobInfrastructureConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"jobInfrastructureType"`
		MarshalTypeManagedComputeClusterJobInfrastructureConfigurationDetails
	}{
		"MANAGED_COMPUTE_CLUSTER",
		(MarshalTypeManagedComputeClusterJobInfrastructureConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
