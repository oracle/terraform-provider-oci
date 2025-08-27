// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDrProtectionGroupMemberVolumeGroupDetails Update properties for a volume group member.
type UpdateDrProtectionGroupMemberVolumeGroupDetails struct {

	// The OCID of the member.
	// Example: `ocid1.database.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The OCID of the backup policy to use in the destination region. This policy will be used to create backups for this volume group after it moves the destination region.
	// Example: `ocid1.volumebackuppolicy.oc1..uniqueID`
	DestinationBackupPolicyId *string `mandatory:"false" json:"destinationBackupPolicyId"`

	// A list of mappings between source volume IDs in the volume group and customer-managed encryption keys in the
	// destination region which will be used to encrypt the volume after it moves to the destination region.
	// If you add the entry for source volumes and its corresponding vault and encryption keys here, you can not use
	// 'commonDestinationKey' for encrypting all volumes with common encryption key. Similarly, if you specify common
	// vault and encryption key using 'commonDestinationKey', you cannot specify vaults and encryption keys individually
	// for each volume using 'sourceVolumeToDestinationEncryptionKeyMappings'.
	// An entry for each volume in volume group should be added in this list. The encryption key will not be updated
	// for the volumes that are part of volume group but missing in this list.
	SourceVolumeToDestinationEncryptionKeyMappings []UpdateSourceVolumeToDestinationEncryptionKeyMappingDetails `mandatory:"false" json:"sourceVolumeToDestinationEncryptionKeyMappings"`

	CommonDestinationKey *UpdateVaultAndEncryptionKeyDetails `mandatory:"false" json:"commonDestinationKey"`

	// The OCID of a compartment in the destination region in which the volume group should be launched.
	// Example: `ocid1.compartment.oc1..uniqueID`
	DestinationCompartmentId *string `mandatory:"false" json:"destinationCompartmentId"`
}

// GetMemberId returns MemberId
func (m UpdateDrProtectionGroupMemberVolumeGroupDetails) GetMemberId() *string {
	return m.MemberId
}

func (m UpdateDrProtectionGroupMemberVolumeGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrProtectionGroupMemberVolumeGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDrProtectionGroupMemberVolumeGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDrProtectionGroupMemberVolumeGroupDetails UpdateDrProtectionGroupMemberVolumeGroupDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeUpdateDrProtectionGroupMemberVolumeGroupDetails
	}{
		"VOLUME_GROUP",
		(MarshalTypeUpdateDrProtectionGroupMemberVolumeGroupDetails)(m),
	}

	return json.Marshal(&s)
}
