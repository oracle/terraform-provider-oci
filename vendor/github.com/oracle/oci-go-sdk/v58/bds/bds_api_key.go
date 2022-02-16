// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BdsApiKey The API key information.
type BdsApiKey struct {

	// Identifier of the user's API key.
	Id *string `mandatory:"true" json:"id"`

	// The user OCID for which this API key was created.
	UserId *string `mandatory:"true" json:"userId"`

	// User friendly identifier used to uniquely differentiate between different API keys.
	// Only ASCII alphanumeric characters with no spaces allowed.
	KeyAlias *string `mandatory:"true" json:"keyAlias"`

	// The name of the region to establish the Object Storage endpoint. Example us-phoenix-1 .
	DefaultRegion *string `mandatory:"true" json:"defaultRegion"`

	// The OCID of your tenancy.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The fingerprint that corresponds to the public API key requested.
	Fingerprint *string `mandatory:"true" json:"fingerprint"`

	// The full path and file name of the private key used for authentication. This location will be automatically selected
	// on the BDS local file system.
	Pemfilepath *string `mandatory:"true" json:"pemfilepath"`

	// The state of the key.
	LifecycleState BdsApiKeyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the API key was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m BdsApiKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsApiKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsApiKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsApiKeyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsApiKeyLifecycleStateEnum Enum with underlying type: string
type BdsApiKeyLifecycleStateEnum string

// Set of constants representing the allowable values for BdsApiKeyLifecycleStateEnum
const (
	BdsApiKeyLifecycleStateCreating BdsApiKeyLifecycleStateEnum = "CREATING"
	BdsApiKeyLifecycleStateActive   BdsApiKeyLifecycleStateEnum = "ACTIVE"
	BdsApiKeyLifecycleStateDeleting BdsApiKeyLifecycleStateEnum = "DELETING"
	BdsApiKeyLifecycleStateDeleted  BdsApiKeyLifecycleStateEnum = "DELETED"
	BdsApiKeyLifecycleStateFailed   BdsApiKeyLifecycleStateEnum = "FAILED"
)

var mappingBdsApiKeyLifecycleStateEnum = map[string]BdsApiKeyLifecycleStateEnum{
	"CREATING": BdsApiKeyLifecycleStateCreating,
	"ACTIVE":   BdsApiKeyLifecycleStateActive,
	"DELETING": BdsApiKeyLifecycleStateDeleting,
	"DELETED":  BdsApiKeyLifecycleStateDeleted,
	"FAILED":   BdsApiKeyLifecycleStateFailed,
}

// GetBdsApiKeyLifecycleStateEnumValues Enumerates the set of values for BdsApiKeyLifecycleStateEnum
func GetBdsApiKeyLifecycleStateEnumValues() []BdsApiKeyLifecycleStateEnum {
	values := make([]BdsApiKeyLifecycleStateEnum, 0)
	for _, v := range mappingBdsApiKeyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsApiKeyLifecycleStateEnumStringValues Enumerates the set of values in String for BdsApiKeyLifecycleStateEnum
func GetBdsApiKeyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingBdsApiKeyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsApiKeyLifecycleStateEnum(val string) (BdsApiKeyLifecycleStateEnum, bool) {
	mappingBdsApiKeyLifecycleStateEnumIgnoreCase := make(map[string]BdsApiKeyLifecycleStateEnum)
	for k, v := range mappingBdsApiKeyLifecycleStateEnum {
		mappingBdsApiKeyLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBdsApiKeyLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
