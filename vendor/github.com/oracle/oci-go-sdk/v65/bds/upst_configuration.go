// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpstConfiguration Information about the UPST configuration.
type UpstConfiguration struct {

	// The instance OCID of the node, which is the resource from which the node backup was acquired.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// Master Encryption key used for encrypting token exchange keytab.
	MasterEncryptionKeyId *string `mandatory:"true" json:"masterEncryptionKeyId"`

	// Secret ID for token exchange keytab
	SecretId *string `mandatory:"true" json:"secretId"`

	// Time when the keytab for token exchange principal is last refreshed, shown as an RFC 3339 formatted datetime string.
	TimeTokenExchangeKeytabLastRefreshed *common.SDKTime `mandatory:"true" json:"timeTokenExchangeKeytabLastRefreshed"`

	// Lifecycle state of the UPST config
	LifecycleState UpstConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time when this UPST config was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when this UPST config was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The kerberos keytab content used for creating identity propagation trust config, in base64 format
	KeytabContent *string `mandatory:"true" json:"keytabContent"`

	// Token exchange kerberos Principal name in cluster
	TokenExchangePrincipalName *string `mandatory:"false" json:"tokenExchangePrincipalName"`
}

func (m UpstConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpstConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpstConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpstConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpstConfigurationLifecycleStateEnum Enum with underlying type: string
type UpstConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for UpstConfigurationLifecycleStateEnum
const (
	UpstConfigurationLifecycleStateCreating UpstConfigurationLifecycleStateEnum = "CREATING"
	UpstConfigurationLifecycleStateActive   UpstConfigurationLifecycleStateEnum = "ACTIVE"
	UpstConfigurationLifecycleStateDeleting UpstConfigurationLifecycleStateEnum = "DELETING"
	UpstConfigurationLifecycleStateInactive UpstConfigurationLifecycleStateEnum = "INACTIVE"
	UpstConfigurationLifecycleStateUpdating UpstConfigurationLifecycleStateEnum = "UPDATING"
	UpstConfigurationLifecycleStateFailed   UpstConfigurationLifecycleStateEnum = "FAILED"
)

var mappingUpstConfigurationLifecycleStateEnum = map[string]UpstConfigurationLifecycleStateEnum{
	"CREATING": UpstConfigurationLifecycleStateCreating,
	"ACTIVE":   UpstConfigurationLifecycleStateActive,
	"DELETING": UpstConfigurationLifecycleStateDeleting,
	"INACTIVE": UpstConfigurationLifecycleStateInactive,
	"UPDATING": UpstConfigurationLifecycleStateUpdating,
	"FAILED":   UpstConfigurationLifecycleStateFailed,
}

var mappingUpstConfigurationLifecycleStateEnumLowerCase = map[string]UpstConfigurationLifecycleStateEnum{
	"creating": UpstConfigurationLifecycleStateCreating,
	"active":   UpstConfigurationLifecycleStateActive,
	"deleting": UpstConfigurationLifecycleStateDeleting,
	"inactive": UpstConfigurationLifecycleStateInactive,
	"updating": UpstConfigurationLifecycleStateUpdating,
	"failed":   UpstConfigurationLifecycleStateFailed,
}

// GetUpstConfigurationLifecycleStateEnumValues Enumerates the set of values for UpstConfigurationLifecycleStateEnum
func GetUpstConfigurationLifecycleStateEnumValues() []UpstConfigurationLifecycleStateEnum {
	values := make([]UpstConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingUpstConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpstConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for UpstConfigurationLifecycleStateEnum
func GetUpstConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"INACTIVE",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingUpstConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpstConfigurationLifecycleStateEnum(val string) (UpstConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingUpstConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
