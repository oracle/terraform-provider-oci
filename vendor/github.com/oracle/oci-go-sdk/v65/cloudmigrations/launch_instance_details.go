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

// LaunchInstanceDetails Instance launch details.
// Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.
type LaunchInstanceDetails struct {

	// The availability domain of the instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the compute capacity reservation under which this instance is launched.
	// You can opt out of all default reservations by specifying an empty string as input for this field.
	// For more information, see Capacity Reservations (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	CreateVnicDetails *CreateVnicDetails `mandatory:"false" json:"createVnicDetails"`

	// The OCID of the dedicated VM host.
	DedicatedVmHostId *string `mandatory:"false" json:"dedicatedVmHostId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A fault domain is a grouping of hardware and infrastructure within an availability domain.
	// Each availability domain contains three fault domains. Fault domains lets you distribute your
	// instances so that they are not on the same physical hardware within a single availability domain.
	// A hardware failure or Compute hardware maintenance that affects one fault domain does not affect
	// instances in other fault domains.
	// If you do not specify the fault domain, the system selects one for you.
	//
	// To get a list of fault domains, use the
	// ListFaultDomains operation in the
	// Identity and Access Management Service API.
	// Example: `FAULT-DOMAIN-1`
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Deprecated. Instead use `hostnameLabel` in
	// CreateVnicDetails.
	// If you provide both, the values must match.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// This is an advanced option.
	// When a bare metal or virtual machine
	// instance boots, the iPXE firmware that runs on the instance is
	// configured to run an iPXE script to continue the boot process.
	// If you want more control over the boot process, you can provide
	// your own custom iPXE script that will run when the instance boots.
	// Be aware that the same iPXE script will run
	// every time an instance boots, not only after the initial
	// LaunchInstance call.
	// By default, the iPXE script connects to the instance's local boot
	// volume over iSCSI and performs a network boot. If you use a custom iPXE
	// script and want to network-boot from the instance's local boot volume
	// over iSCSI in the same way as the default iPXE script, use the
	// following iSCSI IP address: 169.254.0.2, and boot volume IQN:
	// iqn.2015-02.oracle.boot.
	// If your instance boot volume type is paravirtualized,
	// the boot volume is attached to the instance through virtio-scsi and no iPXE script is used.
	// If your instance boot volume type is paravirtualized
	// and you use custom iPXE to perform network-boot into your instance,
	// the primary boot volume is attached as a data volume through the virtio-scsi drive.
	// For more information about the Bring Your Own Image feature of
	// Oracle Cloud Infrastructure, see
	// Bring Your Own Image (https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).
	// For more information about iPXE, see http://ipxe.org.
	IpxeScript *string `mandatory:"false" json:"ipxeScript"`

	InstanceOptions *InstanceOptions `mandatory:"false" json:"instanceOptions"`

	PreemptibleInstanceConfig *PreemptibleInstanceConfigDetails `mandatory:"false" json:"preemptibleInstanceConfig"`

	AgentConfig *LaunchInstanceAgentConfigDetails `mandatory:"false" json:"agentConfig"`

	// The shape of an instance. The shape determines the number of CPUs, amount of memory,
	// and other resources allocated to the instance.
	// You can enumerate all available shapes by calling ListShapes.
	Shape *string `mandatory:"false" json:"shape"`

	ShapeConfig *LaunchInstanceShapeConfigDetails `mandatory:"false" json:"shapeConfig"`

	SourceDetails InstanceSourceDetails `mandatory:"false" json:"sourceDetails"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. By default, the value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`
}

func (m LaunchInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LaunchInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LaunchInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AvailabilityDomain             *string                           `json:"availabilityDomain"`
		CapacityReservationId          *string                           `json:"capacityReservationId"`
		CompartmentId                  *string                           `json:"compartmentId"`
		CreateVnicDetails              *CreateVnicDetails                `json:"createVnicDetails"`
		DedicatedVmHostId              *string                           `json:"dedicatedVmHostId"`
		DefinedTags                    map[string]map[string]interface{} `json:"definedTags"`
		DisplayName                    *string                           `json:"displayName"`
		FaultDomain                    *string                           `json:"faultDomain"`
		FreeformTags                   map[string]string                 `json:"freeformTags"`
		HostnameLabel                  *string                           `json:"hostnameLabel"`
		IpxeScript                     *string                           `json:"ipxeScript"`
		InstanceOptions                *InstanceOptions                  `json:"instanceOptions"`
		PreemptibleInstanceConfig      *PreemptibleInstanceConfigDetails `json:"preemptibleInstanceConfig"`
		AgentConfig                    *LaunchInstanceAgentConfigDetails `json:"agentConfig"`
		Shape                          *string                           `json:"shape"`
		ShapeConfig                    *LaunchInstanceShapeConfigDetails `json:"shapeConfig"`
		SourceDetails                  instancesourcedetails             `json:"sourceDetails"`
		IsPvEncryptionInTransitEnabled *bool                             `json:"isPvEncryptionInTransitEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AvailabilityDomain = model.AvailabilityDomain

	m.CapacityReservationId = model.CapacityReservationId

	m.CompartmentId = model.CompartmentId

	m.CreateVnicDetails = model.CreateVnicDetails

	m.DedicatedVmHostId = model.DedicatedVmHostId

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FaultDomain = model.FaultDomain

	m.FreeformTags = model.FreeformTags

	m.HostnameLabel = model.HostnameLabel

	m.IpxeScript = model.IpxeScript

	m.InstanceOptions = model.InstanceOptions

	m.PreemptibleInstanceConfig = model.PreemptibleInstanceConfig

	m.AgentConfig = model.AgentConfig

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(InstanceSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.IsPvEncryptionInTransitEnabled = model.IsPvEncryptionInTransitEnabled

	return
}
