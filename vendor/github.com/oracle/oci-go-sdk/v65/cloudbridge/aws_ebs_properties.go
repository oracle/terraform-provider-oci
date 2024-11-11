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

// AwsEbsProperties AWS EBS volume related properties.
type AwsEbsProperties struct {

	// Indicates whether the volume is encrypted.
	IsEncrypted *bool `mandatory:"true" json:"isEncrypted"`

	// Indicates whether Amazon EBS Multi-Attach is enabled.
	IsMultiAttachEnabled *bool `mandatory:"true" json:"isMultiAttachEnabled"`

	// The size of the volume, in GiBs.
	SizeInGiBs *int `mandatory:"true" json:"sizeInGiBs"`

	// The ID of the volume.
	VolumeKey *string `mandatory:"true" json:"volumeKey"`

	// The volume type.
	VolumeType *string `mandatory:"true" json:"volumeType"`

	// Information about the volume attachments.
	Attachments []VolumeAttachment `mandatory:"false" json:"attachments"`

	// The Availability Zone for the volume.
	AvailabilityZone *string `mandatory:"false" json:"availabilityZone"`

	// The number of I/O operations per second.
	Iops *int `mandatory:"false" json:"iops"`

	// The volume state.
	Status *string `mandatory:"false" json:"status"`

	// Any tags assigned to the volume.
	Tags []Tag `mandatory:"false" json:"tags"`

	// The throughput that the volume supports, in MiB/s.
	Throughput *int `mandatory:"false" json:"throughput"`
}

func (m AwsEbsProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwsEbsProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
