// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateServiceManagedFleetDetails Service managed Fleet configuration of the batch context.
type CreateServiceManagedFleetDetails struct {

	// Name of the service managed fleet.
	Name *string `mandatory:"true" json:"name"`

	Shape *FleetShape `mandatory:"true" json:"shape"`

	// Maximum number of concurrent tasks for the service managed fleet.
	MaxConcurrentTasks *int `mandatory:"true" json:"maxConcurrentTasks"`
}

func (m CreateServiceManagedFleetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateServiceManagedFleetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateServiceManagedFleetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateServiceManagedFleetDetails CreateServiceManagedFleetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateServiceManagedFleetDetails
	}{
		"SERVICE_MANAGED_FLEET",
		(MarshalTypeCreateServiceManagedFleetDetails)(m),
	}

	return json.Marshal(&s)
}
