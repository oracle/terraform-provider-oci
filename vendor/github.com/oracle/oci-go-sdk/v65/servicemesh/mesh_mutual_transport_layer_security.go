// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MeshMutualTransportLayerSecurity Sets a minimum level of mTLS authentication for all virtual services within the mesh.
type MeshMutualTransportLayerSecurity struct {

	// DISABLED: No minimum virtual services within this mesh can use any mTLS authentication mode.
	// PERMISSIVE: Virtual services within this mesh can use either PERMISSIVE or STRICT modes.
	// STRICT: All virtual services within this mesh must use STRICT mode.
	Minimum MutualTransportLayerSecurityModeEnum `mandatory:"true" json:"minimum"`
}

func (m MeshMutualTransportLayerSecurity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeshMutualTransportLayerSecurity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMutualTransportLayerSecurityModeEnum(string(m.Minimum)); !ok && m.Minimum != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Minimum: %s. Supported values are: %s.", m.Minimum, strings.Join(GetMutualTransportLayerSecurityModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
