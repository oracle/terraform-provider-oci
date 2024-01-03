// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AllVirtualServicesAccessPolicyTarget An internal virtual service directs traffic to all virtual services in a mesh using this target type or vice versa.
type AllVirtualServicesAccessPolicyTarget struct {
}

func (m AllVirtualServicesAccessPolicyTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllVirtualServicesAccessPolicyTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AllVirtualServicesAccessPolicyTarget) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllVirtualServicesAccessPolicyTarget AllVirtualServicesAccessPolicyTarget
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAllVirtualServicesAccessPolicyTarget
	}{
		"ALL_VIRTUAL_SERVICES",
		(MarshalTypeAllVirtualServicesAccessPolicyTarget)(m),
	}

	return json.Marshal(&s)
}
