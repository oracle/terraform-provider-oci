// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicationLocationDetail Replication location detail where the snapshots reside
type ReplicationLocationDetail struct {

	// The type of replication location
	ReplicationLocationType ReplicationLocationDetailReplicationLocationTypeEnum `mandatory:"false" json:"replicationLocationType,omitempty"`

	// Properties for each of the replication location types
	Metadata *interface{} `mandatory:"false" json:"metadata"`
}

func (m ReplicationLocationDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationLocationDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReplicationLocationDetailReplicationLocationTypeEnum(string(m.ReplicationLocationType)); !ok && m.ReplicationLocationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReplicationLocationType: %s. Supported values are: %s.", m.ReplicationLocationType, strings.Join(GetReplicationLocationDetailReplicationLocationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationLocationDetailReplicationLocationTypeEnum Enum with underlying type: string
type ReplicationLocationDetailReplicationLocationTypeEnum string

// Set of constants representing the allowable values for ReplicationLocationDetailReplicationLocationTypeEnum
const (
	ReplicationLocationDetailReplicationLocationTypeOciObjectStore    ReplicationLocationDetailReplicationLocationTypeEnum = "OCI_OBJECT_STORE"
	ReplicationLocationDetailReplicationLocationTypeOlvmStorageDomain ReplicationLocationDetailReplicationLocationTypeEnum = "OLVM_STORAGE_DOMAIN"
)

var mappingReplicationLocationDetailReplicationLocationTypeEnum = map[string]ReplicationLocationDetailReplicationLocationTypeEnum{
	"OCI_OBJECT_STORE":    ReplicationLocationDetailReplicationLocationTypeOciObjectStore,
	"OLVM_STORAGE_DOMAIN": ReplicationLocationDetailReplicationLocationTypeOlvmStorageDomain,
}

var mappingReplicationLocationDetailReplicationLocationTypeEnumLowerCase = map[string]ReplicationLocationDetailReplicationLocationTypeEnum{
	"oci_object_store":    ReplicationLocationDetailReplicationLocationTypeOciObjectStore,
	"olvm_storage_domain": ReplicationLocationDetailReplicationLocationTypeOlvmStorageDomain,
}

// GetReplicationLocationDetailReplicationLocationTypeEnumValues Enumerates the set of values for ReplicationLocationDetailReplicationLocationTypeEnum
func GetReplicationLocationDetailReplicationLocationTypeEnumValues() []ReplicationLocationDetailReplicationLocationTypeEnum {
	values := make([]ReplicationLocationDetailReplicationLocationTypeEnum, 0)
	for _, v := range mappingReplicationLocationDetailReplicationLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationLocationDetailReplicationLocationTypeEnumStringValues Enumerates the set of values in String for ReplicationLocationDetailReplicationLocationTypeEnum
func GetReplicationLocationDetailReplicationLocationTypeEnumStringValues() []string {
	return []string{
		"OCI_OBJECT_STORE",
		"OLVM_STORAGE_DOMAIN",
	}
}

// GetMappingReplicationLocationDetailReplicationLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationLocationDetailReplicationLocationTypeEnum(val string) (ReplicationLocationDetailReplicationLocationTypeEnum, bool) {
	enum, ok := mappingReplicationLocationDetailReplicationLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
