// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicaVaultMetadata Metadata for the replica vault, needed if different from primary vault
type ReplicaVaultMetadata interface {
}

type replicavaultmetadata struct {
	JsonData  []byte
	VaultType string `json:"vaultType"`
}

// UnmarshalJSON unmarshals json
func (m *replicavaultmetadata) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreplicavaultmetadata replicavaultmetadata
	s := struct {
		Model Unmarshalerreplicavaultmetadata
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VaultType = s.Model.VaultType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *replicavaultmetadata) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.VaultType {
	case "EXTERNAL":
		mm := ReplicaExternalVaultMetadata{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ReplicaVaultMetadata: %s.", m.VaultType)
		return *m, nil
	}
}

func (m replicavaultmetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m replicavaultmetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicaVaultMetadataVaultTypeEnum Enum with underlying type: string
type ReplicaVaultMetadataVaultTypeEnum string

// Set of constants representing the allowable values for ReplicaVaultMetadataVaultTypeEnum
const (
	ReplicaVaultMetadataVaultTypeExternal ReplicaVaultMetadataVaultTypeEnum = "EXTERNAL"
)

var mappingReplicaVaultMetadataVaultTypeEnum = map[string]ReplicaVaultMetadataVaultTypeEnum{
	"EXTERNAL": ReplicaVaultMetadataVaultTypeExternal,
}

var mappingReplicaVaultMetadataVaultTypeEnumLowerCase = map[string]ReplicaVaultMetadataVaultTypeEnum{
	"external": ReplicaVaultMetadataVaultTypeExternal,
}

// GetReplicaVaultMetadataVaultTypeEnumValues Enumerates the set of values for ReplicaVaultMetadataVaultTypeEnum
func GetReplicaVaultMetadataVaultTypeEnumValues() []ReplicaVaultMetadataVaultTypeEnum {
	values := make([]ReplicaVaultMetadataVaultTypeEnum, 0)
	for _, v := range mappingReplicaVaultMetadataVaultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicaVaultMetadataVaultTypeEnumStringValues Enumerates the set of values in String for ReplicaVaultMetadataVaultTypeEnum
func GetReplicaVaultMetadataVaultTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL",
	}
}

// GetMappingReplicaVaultMetadataVaultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicaVaultMetadataVaultTypeEnum(val string) (ReplicaVaultMetadataVaultTypeEnum, bool) {
	enum, ok := mappingReplicaVaultMetadataVaultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
