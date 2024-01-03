// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AssetSourceConnection Descriptor of a connection to an asset source.
type AssetSourceConnection struct {

	// The type of connection for an asset source.
	ConnectionType AssetSourceConnectionTypeEnum `mandatory:"true" json:"connectionType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud bridge connector used for migration operations.
	ConnectorId *string `mandatory:"true" json:"connectorId"`

	// Type-specific identifier for an asset source.
	AssetSourceKey *string `mandatory:"true" json:"assetSourceKey"`

	// The current state of the connection.
	LifecycleState AssetSourceConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed sub-state of the connection.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`
}

func (m AssetSourceConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssetSourceConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetSourceConnectionTypeEnum(string(m.ConnectionType)); !ok && m.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", m.ConnectionType, strings.Join(GetAssetSourceConnectionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAssetSourceConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAssetSourceConnectionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
