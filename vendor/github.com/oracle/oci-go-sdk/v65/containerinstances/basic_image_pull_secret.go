// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BasicImagePullSecret A BasicImagePullSecret is a ImagePullSecret which accepts username and password as credentials information.
type BasicImagePullSecret struct {

	// The registry endpoint of the container image.
	RegistryEndpoint *string `mandatory:"true" json:"registryEndpoint"`
}

// GetRegistryEndpoint returns RegistryEndpoint
func (m BasicImagePullSecret) GetRegistryEndpoint() *string {
	return m.RegistryEndpoint
}

func (m BasicImagePullSecret) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicImagePullSecret) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicImagePullSecret) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicImagePullSecret BasicImagePullSecret
	s := struct {
		DiscriminatorParam string `json:"secretType"`
		MarshalTypeBasicImagePullSecret
	}{
		"BASIC",
		(MarshalTypeBasicImagePullSecret)(m),
	}

	return json.Marshal(&s)
}
