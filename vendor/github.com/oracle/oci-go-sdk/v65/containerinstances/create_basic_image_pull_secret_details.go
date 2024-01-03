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

// CreateBasicImagePullSecretDetails A CreateBasicImagePullSecretDetails is a ImagePullSecret which accepts username and password as credentials information.
type CreateBasicImagePullSecretDetails struct {

	// The registry endpoint of the container image.
	RegistryEndpoint *string `mandatory:"true" json:"registryEndpoint"`

	// The username which should be used with the registry for authentication. The value is expected in base64 format.
	Username *string `mandatory:"true" json:"username"`

	// The password which should be used with the registry for authentication. The value is expected in base64 format.
	Password *string `mandatory:"true" json:"password"`
}

// GetRegistryEndpoint returns RegistryEndpoint
func (m CreateBasicImagePullSecretDetails) GetRegistryEndpoint() *string {
	return m.RegistryEndpoint
}

func (m CreateBasicImagePullSecretDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBasicImagePullSecretDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateBasicImagePullSecretDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateBasicImagePullSecretDetails CreateBasicImagePullSecretDetails
	s := struct {
		DiscriminatorParam string `json:"secretType"`
		MarshalTypeCreateBasicImagePullSecretDetails
	}{
		"BASIC",
		(MarshalTypeCreateBasicImagePullSecretDetails)(m),
	}

	return json.Marshal(&s)
}
