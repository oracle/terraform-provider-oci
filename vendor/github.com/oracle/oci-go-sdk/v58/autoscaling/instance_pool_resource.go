// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InstancePoolResource A Compute instance pool.
type InstancePoolResource struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource that is managed by the autoscaling configuration.
	Id *string `mandatory:"true" json:"id"`
}

//GetId returns Id
func (m InstancePoolResource) GetId() *string {
	return m.Id
}

func (m InstancePoolResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstancePoolResource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstancePoolResource InstancePoolResource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstancePoolResource
	}{
		"instancePool",
		(MarshalTypeInstancePoolResource)(m),
	}

	return json.Marshal(&s)
}
