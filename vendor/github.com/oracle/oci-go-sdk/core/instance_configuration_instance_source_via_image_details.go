// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceConfigurationInstanceSourceViaImageDetails The representation of InstanceConfigurationInstanceSourceViaImageDetails
type InstanceConfigurationInstanceSourceViaImageDetails struct {

	// The OCID of the image used to boot the instance.
	ImageId *string `mandatory:"false" json:"imageId"`
}

func (m InstanceConfigurationInstanceSourceViaImageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationInstanceSourceViaImageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationInstanceSourceViaImageDetails InstanceConfigurationInstanceSourceViaImageDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceConfigurationInstanceSourceViaImageDetails
	}{
		"image",
		(MarshalTypeInstanceConfigurationInstanceSourceViaImageDetails)(m),
	}

	return json.Marshal(&s)
}
