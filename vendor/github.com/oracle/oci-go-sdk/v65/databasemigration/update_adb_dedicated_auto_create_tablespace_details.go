// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAdbDedicatedAutoCreateTablespaceDetails Auto create tablespace settings that are valid for Dedicated Autonomous Databases.
type UpdateAdbDedicatedAutoCreateTablespaceDetails struct {

	// Set this property to true to auto-create tablespaces in the target Database.
	// Note: This is not applicable for Autonomous Database Serverless databases.
	IsAutoCreate *bool `mandatory:"false" json:"isAutoCreate"`

	// Set this property to true to enable tablespace of the type big file.
	IsBigFile *bool `mandatory:"false" json:"isBigFile"`

	// Size to extend the tablespace in MB.
	// Note: Only applicable if 'isBigFile' property is set to true.
	ExtendSizeInMBs *int `mandatory:"false" json:"extendSizeInMBs"`

	// Size of Oracle database blocks in KB.
	BlockSizeInKBs DataPumpTablespaceBlockSizesInKbEnum `mandatory:"false" json:"blockSizeInKBs,omitempty"`
}

func (m UpdateAdbDedicatedAutoCreateTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAdbDedicatedAutoCreateTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataPumpTablespaceBlockSizesInKbEnum(string(m.BlockSizeInKBs)); !ok && m.BlockSizeInKBs != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BlockSizeInKBs: %s. Supported values are: %s.", m.BlockSizeInKBs, strings.Join(GetDataPumpTablespaceBlockSizesInKbEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAdbDedicatedAutoCreateTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAdbDedicatedAutoCreateTablespaceDetails UpdateAdbDedicatedAutoCreateTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeUpdateAdbDedicatedAutoCreateTablespaceDetails
	}{
		"ADB_D_AUTOCREATE",
		(MarshalTypeUpdateAdbDedicatedAutoCreateTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
