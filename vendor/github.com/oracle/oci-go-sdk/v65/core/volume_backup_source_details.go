// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// VolumeBackupSourceDetails Define a volume that will be created and attached on the creation of an Instance.
type VolumeBackupSourceDetails interface {

	// The ocid of the VolumeBackup that will be utilized as the source for creating the Volume to attach to the
	// Instance.
	// Warning: This VolumeBackup will not longer be usable or accessible once this operation is completed. It will
	// become a part of the Image.
	GetVolumeBackupId() *string
}

type volumebackupsourcedetails struct {
	JsonData       []byte
	VolumeBackupId *string `mandatory:"true" json:"volumeBackupId"`
	VolumeType     string  `json:"volumeType"`
}

// UnmarshalJSON unmarshals json
func (m *volumebackupsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervolumebackupsourcedetails volumebackupsourcedetails
	s := struct {
		Model Unmarshalervolumebackupsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VolumeBackupId = s.Model.VolumeBackupId
	m.VolumeType = s.Model.VolumeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *volumebackupsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.VolumeType {
	case "DATA_VOLUME":
		mm := DataVolumeBackupSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BOOT_VOLUME":
		mm := BootVolumeBackupSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for VolumeBackupSourceDetails: %s.", m.VolumeType)
		return *m, nil
	}
}

// GetVolumeBackupId returns VolumeBackupId
func (m volumebackupsourcedetails) GetVolumeBackupId() *string {
	return m.VolumeBackupId
}

func (m volumebackupsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m volumebackupsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VolumeBackupSourceDetailsVolumeTypeEnum Enum with underlying type: string
type VolumeBackupSourceDetailsVolumeTypeEnum string

// Set of constants representing the allowable values for VolumeBackupSourceDetailsVolumeTypeEnum
const (
	VolumeBackupSourceDetailsVolumeTypeBootVolume VolumeBackupSourceDetailsVolumeTypeEnum = "BOOT_VOLUME"
	VolumeBackupSourceDetailsVolumeTypeDataVolume VolumeBackupSourceDetailsVolumeTypeEnum = "DATA_VOLUME"
)

var mappingVolumeBackupSourceDetailsVolumeTypeEnum = map[string]VolumeBackupSourceDetailsVolumeTypeEnum{
	"BOOT_VOLUME": VolumeBackupSourceDetailsVolumeTypeBootVolume,
	"DATA_VOLUME": VolumeBackupSourceDetailsVolumeTypeDataVolume,
}

var mappingVolumeBackupSourceDetailsVolumeTypeEnumLowerCase = map[string]VolumeBackupSourceDetailsVolumeTypeEnum{
	"boot_volume": VolumeBackupSourceDetailsVolumeTypeBootVolume,
	"data_volume": VolumeBackupSourceDetailsVolumeTypeDataVolume,
}

// GetVolumeBackupSourceDetailsVolumeTypeEnumValues Enumerates the set of values for VolumeBackupSourceDetailsVolumeTypeEnum
func GetVolumeBackupSourceDetailsVolumeTypeEnumValues() []VolumeBackupSourceDetailsVolumeTypeEnum {
	values := make([]VolumeBackupSourceDetailsVolumeTypeEnum, 0)
	for _, v := range mappingVolumeBackupSourceDetailsVolumeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVolumeBackupSourceDetailsVolumeTypeEnumStringValues Enumerates the set of values in String for VolumeBackupSourceDetailsVolumeTypeEnum
func GetVolumeBackupSourceDetailsVolumeTypeEnumStringValues() []string {
	return []string{
		"BOOT_VOLUME",
		"DATA_VOLUME",
	}
}

// GetMappingVolumeBackupSourceDetailsVolumeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVolumeBackupSourceDetailsVolumeTypeEnum(val string) (VolumeBackupSourceDetailsVolumeTypeEnum, bool) {
	enum, ok := mappingVolumeBackupSourceDetailsVolumeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
