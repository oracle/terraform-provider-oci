// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceConfigurationHostGroupPlacementConstraintDetails The details for providing placement constraints using the compute host group OCID.
type InstanceConfigurationHostGroupPlacementConstraintDetails struct {

	// The OCID of the compute host group. This is only available for dedicated capacity customers.
	ComputeHostGroupId *string `mandatory:"true" json:"computeHostGroupId"`
}

func (m InstanceConfigurationHostGroupPlacementConstraintDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfigurationHostGroupPlacementConstraintDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationHostGroupPlacementConstraintDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationHostGroupPlacementConstraintDetails InstanceConfigurationHostGroupPlacementConstraintDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationHostGroupPlacementConstraintDetails
	}{
		"HOST_GROUP",
		(MarshalTypeInstanceConfigurationHostGroupPlacementConstraintDetails)(m),
	}

	return json.Marshal(&s)
}
