// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImageSourceViaVolumeBackupDetails Define an Image by providing a set of VolumeBackups that represent the BootVolume and additional, optional data Volumes.
// This is an alternative way of creating an Image from an OS image that is already represented as VolumeBackup[s] in OCI.
type ImageSourceViaVolumeBackupDetails struct {
	BootVolume *BootVolumeBackupSourceDetails `mandatory:"true" json:"bootVolume"`

	OperatingSystem *string `mandatory:"false" json:"operatingSystem"`

	OperatingSystemVersion *string `mandatory:"false" json:"operatingSystemVersion"`

	// Data volume backup source parameters.
	DataVolumes []DataVolumeBackupSourceDetails `mandatory:"false" json:"dataVolumes"`

	// The format of the image to be imported. Only monolithic
	// images are supported. This attribute is not used for exported Oracle images with the OCI image format.
	SourceImageType ImageSourceDetailsSourceImageTypeEnum `mandatory:"false" json:"sourceImageType,omitempty"`
}

// GetOperatingSystem returns OperatingSystem
func (m ImageSourceViaVolumeBackupDetails) GetOperatingSystem() *string {
	return m.OperatingSystem
}

// GetOperatingSystemVersion returns OperatingSystemVersion
func (m ImageSourceViaVolumeBackupDetails) GetOperatingSystemVersion() *string {
	return m.OperatingSystemVersion
}

// GetSourceImageType returns SourceImageType
func (m ImageSourceViaVolumeBackupDetails) GetSourceImageType() ImageSourceDetailsSourceImageTypeEnum {
	return m.SourceImageType
}

func (m ImageSourceViaVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageSourceViaVolumeBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImageSourceDetailsSourceImageTypeEnum(string(m.SourceImageType)); !ok && m.SourceImageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceImageType: %s. Supported values are: %s.", m.SourceImageType, strings.Join(GetImageSourceDetailsSourceImageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImageSourceViaVolumeBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageSourceViaVolumeBackupDetails ImageSourceViaVolumeBackupDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeImageSourceViaVolumeBackupDetails
	}{
		"volumeBackupDetails",
		(MarshalTypeImageSourceViaVolumeBackupDetails)(m),
	}

	return json.Marshal(&s)
}
