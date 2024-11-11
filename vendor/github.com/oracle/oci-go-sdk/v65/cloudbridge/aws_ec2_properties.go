// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwsEc2Properties AWS virtual machine related properties.
type AwsEc2Properties struct {

	// The architecture of the image.
	Architecture *string `mandatory:"true" json:"architecture"`

	// The ID of the instance.
	InstanceKey *string `mandatory:"true" json:"instanceKey"`

	// The instance type.
	InstanceType *string `mandatory:"true" json:"instanceType"`

	// The device name of the root device volume.
	RootDeviceName *string `mandatory:"true" json:"rootDeviceName"`

	State *InstanceState `mandatory:"true" json:"state"`

	// The boot mode of the instance.
	BootMode *string `mandatory:"false" json:"bootMode"`

	// The ID of the Capacity Reservation.
	CapacityReservationKey *string `mandatory:"false" json:"capacityReservationKey"`

	// Indicates if the elastic inference accelerators attached to an instance
	AreElasticInferenceAcceleratorsPresent *bool `mandatory:"false" json:"areElasticInferenceAcceleratorsPresent"`

	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	IsEnclaveOptions *bool `mandatory:"false" json:"isEnclaveOptions"`

	// Indicates whether the instance is enabled for hibernation.
	IsHibernationOptions *bool `mandatory:"false" json:"isHibernationOptions"`

	// The ID of the AMI used to launch the instance.
	ImageKey *string `mandatory:"false" json:"imageKey"`

	// Indicates whether this is a Spot Instance or a Scheduled Instance.
	InstanceLifecycle *string `mandatory:"false" json:"instanceLifecycle"`

	// The public IPv4 address, or the Carrier IP address assigned to the instance.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The IPv6 address assigned to the instance.
	Ipv6Address *string `mandatory:"false" json:"ipv6Address"`

	// The kernel associated with this instance, if applicable.
	KernelKey *string `mandatory:"false" json:"kernelKey"`

	// The time the instance was launched.
	TimeLaunch *common.SDKTime `mandatory:"false" json:"timeLaunch"`

	// The license configurations for the instance.
	Licenses []string `mandatory:"false" json:"licenses"`

	// Provides information on the recovery and maintenance options of your instance.
	MaintenanceOptions *string `mandatory:"false" json:"maintenanceOptions"`

	// The monitoring for the instance.
	Monitoring *string `mandatory:"false" json:"monitoring"`

	// The network interfaces for the instance.
	NetworkInterfaces []InstanceNetworkInterface `mandatory:"false" json:"networkInterfaces"`

	Placement *Placement `mandatory:"false" json:"placement"`

	// (IPv4 only) The private DNS hostname name assigned to the instance.
	PrivateDnsName *string `mandatory:"false" json:"privateDnsName"`

	// The private IPv4 address assigned to the instance.
	PrivateIpAddress *string `mandatory:"false" json:"privateIpAddress"`

	// The root device type used by the AMI. The AMI can use an EBS volume or an instance store volume.
	RootDeviceType *string `mandatory:"false" json:"rootDeviceType"`

	// The security groups for the instance.
	SecurityGroups []GroupIdentifier `mandatory:"false" json:"securityGroups"`

	// Indicates whether source/destination checking is enabled.
	IsSourceDestCheck *bool `mandatory:"false" json:"isSourceDestCheck"`

	// If the request is a Spot Instance request, this value will be true.
	IsSpotInstance *bool `mandatory:"false" json:"isSpotInstance"`

	// Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.
	SriovNetSupport *string `mandatory:"false" json:"sriovNetSupport"`

	// EC2-VPC The ID of the subnet in which the instance is running.
	SubnetKey *string `mandatory:"false" json:"subnetKey"`

	// Any tags assigned to the instance.
	Tags []Tag `mandatory:"false" json:"tags"`

	// If the instance is configured for NitroTPM support, the value is v2.0.
	TpmSupport *string `mandatory:"false" json:"tpmSupport"`

	// The virtualization type of the instance.
	VirtualizationType *string `mandatory:"false" json:"virtualizationType"`

	// EC2-VPC The ID of the VPC in which the instance is running.
	VpcKey *string `mandatory:"false" json:"vpcKey"`
}

func (m AwsEc2Properties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsEc2Properties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
