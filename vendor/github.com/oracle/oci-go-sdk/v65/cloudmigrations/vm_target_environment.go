// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmTargetEnvironment Description of the VM target environment.
type VmTargetEnvironment struct {

	// OCID of the VM configuration VCN.
	Vcn *string `mandatory:"true" json:"vcn"`

	// OCID of the VM configuration subnet.
	Subnet *string `mandatory:"true" json:"subnet"`

	// Target compartment identifier
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// Availability Domain of the VM configuration.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Fault domain of the VM configuration.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// OCID of the dedicated VM configuration host.
	DedicatedVmHost *string `mandatory:"false" json:"dedicatedVmHost"`

	// Microsoft license for the VM configuration.
	MsLicense *string `mandatory:"false" json:"msLicense"`

	// Preferred VM shape type provided by the customer.
	PreferredShapeType VmTargetAssetPreferredShapeTypeEnum `mandatory:"false" json:"preferredShapeType,omitempty"`
}

// GetTargetCompartmentId returns TargetCompartmentId
func (m VmTargetEnvironment) GetTargetCompartmentId() *string {
	return m.TargetCompartmentId
}

func (m VmTargetEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmTargetEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmTargetAssetPreferredShapeTypeEnum(string(m.PreferredShapeType)); !ok && m.PreferredShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredShapeType: %s. Supported values are: %s.", m.PreferredShapeType, strings.Join(GetVmTargetAssetPreferredShapeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VmTargetEnvironment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVmTargetEnvironment VmTargetEnvironment
	s := struct {
		DiscriminatorParam string `json:"targetEnvironmentType"`
		MarshalTypeVmTargetEnvironment
	}{
		"VM_TARGET_ENV",
		(MarshalTypeVmTargetEnvironment)(m),
	}

	return json.Marshal(&s)
}
