// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseWallet The Autonomous Database wallet details.
type AutonomousDatabaseWallet struct {

	// The current lifecycle state of the Autonomous Database wallet.
	LifecycleState AutonomousDatabaseWalletLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the wallet was last rotated.
	TimeRotated *common.SDKTime `mandatory:"false" json:"timeRotated"`
}

func (m AutonomousDatabaseWallet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseWallet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseWalletLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseWalletLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseWalletLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseWalletLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseWalletLifecycleStateEnum
const (
	AutonomousDatabaseWalletLifecycleStateActive   AutonomousDatabaseWalletLifecycleStateEnum = "ACTIVE"
	AutonomousDatabaseWalletLifecycleStateUpdating AutonomousDatabaseWalletLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseWalletLifecycleStateEnum = map[string]AutonomousDatabaseWalletLifecycleStateEnum{
	"ACTIVE":   AutonomousDatabaseWalletLifecycleStateActive,
	"UPDATING": AutonomousDatabaseWalletLifecycleStateUpdating,
}

var mappingAutonomousDatabaseWalletLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseWalletLifecycleStateEnum{
	"active":   AutonomousDatabaseWalletLifecycleStateActive,
	"updating": AutonomousDatabaseWalletLifecycleStateUpdating,
}

// GetAutonomousDatabaseWalletLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseWalletLifecycleStateEnum
func GetAutonomousDatabaseWalletLifecycleStateEnumValues() []AutonomousDatabaseWalletLifecycleStateEnum {
	values := make([]AutonomousDatabaseWalletLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseWalletLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseWalletLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseWalletLifecycleStateEnum
func GetAutonomousDatabaseWalletLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseWalletLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseWalletLifecycleStateEnum(val string) (AutonomousDatabaseWalletLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseWalletLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
