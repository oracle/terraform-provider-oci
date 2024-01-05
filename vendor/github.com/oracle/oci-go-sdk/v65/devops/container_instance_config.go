// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerInstanceConfig Specifies ContainerInstance configuration.
type ContainerInstanceConfig struct {

	// The shape of the ContainerInstance. The shape determines the resources available to the ContainerInstance.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	ShapeConfig *ShapeConfig `mandatory:"true" json:"shapeConfig"`

	NetworkChannel NetworkChannel `mandatory:"true" json:"networkChannel"`

	// The OCID of the compartment where the ContainerInstance will be created.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Availability domain where the ContainerInstance will be created.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
}

func (m ContainerInstanceConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerInstanceConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ContainerInstanceConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerInstanceConfig ContainerInstanceConfig
	s := struct {
		DiscriminatorParam string `json:"containerConfigType"`
		MarshalTypeContainerInstanceConfig
	}{
		"CONTAINER_INSTANCE_CONFIG",
		(MarshalTypeContainerInstanceConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ContainerInstanceConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId      *string        `json:"compartmentId"`
		AvailabilityDomain *string        `json:"availabilityDomain"`
		ShapeName          *string        `json:"shapeName"`
		ShapeConfig        *ShapeConfig   `json:"shapeConfig"`
		NetworkChannel     networkchannel `json:"networkChannel"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.ShapeName = model.ShapeName

	m.ShapeConfig = model.ShapeConfig

	nn, e = model.NetworkChannel.UnmarshalPolymorphicJSON(model.NetworkChannel.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkChannel = nn.(NetworkChannel)
	} else {
		m.NetworkChannel = nil
	}

	return
}
