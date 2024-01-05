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

// ResizeVmClusterNetworkDetails Details of Db server network nodes to extend or shrink the VM cluster network. Applies to Exadata Cloud@Customer
// instances only.
type ResizeVmClusterNetworkDetails struct {

	// Actions that can be performed on the VM cluster network.
	// ADD_DBSERVER_NETWORK - Provide Db server network details of network nodes to be added to the VM cluster network.
	// REMOVE_DBSERVER_NETWORK - Provide Db server network details of network nodes to be removed from the VM cluster network.
	Action ResizeVmClusterNetworkDetailsActionEnum `mandatory:"true" json:"action"`

	// Details of the client and backup networks.
	VmNetworks []VmNetworkDetails `mandatory:"true" json:"vmNetworks"`
}

func (m ResizeVmClusterNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResizeVmClusterNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResizeVmClusterNetworkDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetResizeVmClusterNetworkDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResizeVmClusterNetworkDetailsActionEnum Enum with underlying type: string
type ResizeVmClusterNetworkDetailsActionEnum string

// Set of constants representing the allowable values for ResizeVmClusterNetworkDetailsActionEnum
const (
	ResizeVmClusterNetworkDetailsActionAddDbserverNetwork    ResizeVmClusterNetworkDetailsActionEnum = "ADD_DBSERVER_NETWORK"
	ResizeVmClusterNetworkDetailsActionRemoveDbserverNetwork ResizeVmClusterNetworkDetailsActionEnum = "REMOVE_DBSERVER_NETWORK"
)

var mappingResizeVmClusterNetworkDetailsActionEnum = map[string]ResizeVmClusterNetworkDetailsActionEnum{
	"ADD_DBSERVER_NETWORK":    ResizeVmClusterNetworkDetailsActionAddDbserverNetwork,
	"REMOVE_DBSERVER_NETWORK": ResizeVmClusterNetworkDetailsActionRemoveDbserverNetwork,
}

var mappingResizeVmClusterNetworkDetailsActionEnumLowerCase = map[string]ResizeVmClusterNetworkDetailsActionEnum{
	"add_dbserver_network":    ResizeVmClusterNetworkDetailsActionAddDbserverNetwork,
	"remove_dbserver_network": ResizeVmClusterNetworkDetailsActionRemoveDbserverNetwork,
}

// GetResizeVmClusterNetworkDetailsActionEnumValues Enumerates the set of values for ResizeVmClusterNetworkDetailsActionEnum
func GetResizeVmClusterNetworkDetailsActionEnumValues() []ResizeVmClusterNetworkDetailsActionEnum {
	values := make([]ResizeVmClusterNetworkDetailsActionEnum, 0)
	for _, v := range mappingResizeVmClusterNetworkDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetResizeVmClusterNetworkDetailsActionEnumStringValues Enumerates the set of values in String for ResizeVmClusterNetworkDetailsActionEnum
func GetResizeVmClusterNetworkDetailsActionEnumStringValues() []string {
	return []string{
		"ADD_DBSERVER_NETWORK",
		"REMOVE_DBSERVER_NETWORK",
	}
}

// GetMappingResizeVmClusterNetworkDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResizeVmClusterNetworkDetailsActionEnum(val string) (ResizeVmClusterNetworkDetailsActionEnum, bool) {
	enum, ok := mappingResizeVmClusterNetworkDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
