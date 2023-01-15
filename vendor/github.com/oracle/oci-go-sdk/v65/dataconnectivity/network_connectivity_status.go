// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkConnectivityStatus The network validation status for a Private Endpoint - Data Asset pair.
type NetworkConnectivityStatus struct {

	// DataAsset key to which the NetworkValidationStatus belongs to.
	DataAssetKey *string `mandatory:"true" json:"dataAssetKey"`

	// PrivateEndpoint key, if any, to which the NetworkValidationStatus belongs to.
	PrivateEndPointKey *string `mandatory:"false" json:"privateEndPointKey"`

	// Exception or error message encountered while testing network reachability for the data asset.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The timestamp when the network validation was last updated for the given DataAsset-PrivateEndpoint pair.
	TimeLastUpdated *common.SDKTime `mandatory:"false" json:"timeLastUpdated"`

	// Exception or error message encountered while testing network reachability for the data asset.
	NetworkValidationStatusEnum NetworkConnectivityStatusNetworkValidationStatusEnumEnum `mandatory:"false" json:"networkValidationStatusEnum,omitempty"`
}

func (m NetworkConnectivityStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkConnectivityStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNetworkConnectivityStatusNetworkValidationStatusEnumEnum(string(m.NetworkValidationStatusEnum)); !ok && m.NetworkValidationStatusEnum != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkValidationStatusEnum: %s. Supported values are: %s.", m.NetworkValidationStatusEnum, strings.Join(GetNetworkConnectivityStatusNetworkValidationStatusEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NetworkConnectivityStatusNetworkValidationStatusEnumEnum Enum with underlying type: string
type NetworkConnectivityStatusNetworkValidationStatusEnumEnum string

// Set of constants representing the allowable values for NetworkConnectivityStatusNetworkValidationStatusEnumEnum
const (
	NetworkConnectivityStatusNetworkValidationStatusEnumReachable    NetworkConnectivityStatusNetworkValidationStatusEnumEnum = "REACHABLE"
	NetworkConnectivityStatusNetworkValidationStatusEnumNotReachable NetworkConnectivityStatusNetworkValidationStatusEnumEnum = "NOT_REACHABLE"
	NetworkConnectivityStatusNetworkValidationStatusEnumError        NetworkConnectivityStatusNetworkValidationStatusEnumEnum = "ERROR"
)

var mappingNetworkConnectivityStatusNetworkValidationStatusEnumEnum = map[string]NetworkConnectivityStatusNetworkValidationStatusEnumEnum{
	"REACHABLE":     NetworkConnectivityStatusNetworkValidationStatusEnumReachable,
	"NOT_REACHABLE": NetworkConnectivityStatusNetworkValidationStatusEnumNotReachable,
	"ERROR":         NetworkConnectivityStatusNetworkValidationStatusEnumError,
}

var mappingNetworkConnectivityStatusNetworkValidationStatusEnumEnumLowerCase = map[string]NetworkConnectivityStatusNetworkValidationStatusEnumEnum{
	"reachable":     NetworkConnectivityStatusNetworkValidationStatusEnumReachable,
	"not_reachable": NetworkConnectivityStatusNetworkValidationStatusEnumNotReachable,
	"error":         NetworkConnectivityStatusNetworkValidationStatusEnumError,
}

// GetNetworkConnectivityStatusNetworkValidationStatusEnumEnumValues Enumerates the set of values for NetworkConnectivityStatusNetworkValidationStatusEnumEnum
func GetNetworkConnectivityStatusNetworkValidationStatusEnumEnumValues() []NetworkConnectivityStatusNetworkValidationStatusEnumEnum {
	values := make([]NetworkConnectivityStatusNetworkValidationStatusEnumEnum, 0)
	for _, v := range mappingNetworkConnectivityStatusNetworkValidationStatusEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetNetworkConnectivityStatusNetworkValidationStatusEnumEnumStringValues Enumerates the set of values in String for NetworkConnectivityStatusNetworkValidationStatusEnumEnum
func GetNetworkConnectivityStatusNetworkValidationStatusEnumEnumStringValues() []string {
	return []string{
		"REACHABLE",
		"NOT_REACHABLE",
		"ERROR",
	}
}

// GetMappingNetworkConnectivityStatusNetworkValidationStatusEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNetworkConnectivityStatusNetworkValidationStatusEnumEnum(val string) (NetworkConnectivityStatusNetworkValidationStatusEnumEnum, bool) {
	enum, ok := mappingNetworkConnectivityStatusNetworkValidationStatusEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
