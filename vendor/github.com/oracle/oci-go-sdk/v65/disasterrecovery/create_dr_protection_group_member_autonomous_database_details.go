// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDrProtectionGroupMemberAutonomousDatabaseDetails Create properties for an Autonomous Database Serverless member.
type CreateDrProtectionGroupMemberAutonomousDatabaseDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The OCID of the vault secret where the database SYSDBA password is stored.
	// This password is required and used for performing database DR Drill operations when using full clone.
	// Example: `ocid1.vaultsecret.oc1..uniqueID`
	PasswordVaultSecretId *string `mandatory:"false" json:"passwordVaultSecretId"`

	// This specifies the mechanism used to create a temporary Autonomous Database instance for DR Drills.
	// See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-clone-about.html for information about these clone types.
	// See https://docs.oracle.com/en/cloud/paas/autonomous-database/serverless/adbsb/autonomous-data-guard-snapshot-standby.html for information about snapshot standby.
	AutonomousDatabaseStandbyTypeForDrDrills AutonomousDatabaseStandbyTypeForDrDrillsEnum `mandatory:"false" json:"autonomousDatabaseStandbyTypeForDrDrills,omitempty"`
}

// GetMemberId returns MemberId
func (m CreateDrProtectionGroupMemberAutonomousDatabaseDetails) GetMemberId() *string {
	return m.MemberId
}

func (m CreateDrProtectionGroupMemberAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupMemberAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseStandbyTypeForDrDrillsEnum(string(m.AutonomousDatabaseStandbyTypeForDrDrills)); !ok && m.AutonomousDatabaseStandbyTypeForDrDrills != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutonomousDatabaseStandbyTypeForDrDrills: %s. Supported values are: %s.", m.AutonomousDatabaseStandbyTypeForDrDrills, strings.Join(GetAutonomousDatabaseStandbyTypeForDrDrillsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDrProtectionGroupMemberAutonomousDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDrProtectionGroupMemberAutonomousDatabaseDetails CreateDrProtectionGroupMemberAutonomousDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeCreateDrProtectionGroupMemberAutonomousDatabaseDetails
	}{
		"AUTONOMOUS_DATABASE",
		(MarshalTypeCreateDrProtectionGroupMemberAutonomousDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
