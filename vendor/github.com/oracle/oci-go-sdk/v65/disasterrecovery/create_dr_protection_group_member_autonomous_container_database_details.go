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

// CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails Create properties for an Autonomous Container Database member.
type CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The type of connection strings used to connect to an Autonomous Container Database snapshot standby created during a DR Drill operation.
	// See https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbcl/index.html for information about these service types.
	ConnectionStringType AutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum `mandatory:"false" json:"connectionStringType,omitempty"`
}

// GetMemberId returns MemberId
func (m CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails) GetMemberId() *string {
	return m.MemberId
}

func (m CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnum(string(m.ConnectionStringType)); !ok && m.ConnectionStringType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionStringType: %s. Supported values are: %s.", m.ConnectionStringType, strings.Join(GetAutonomousContainerDatabaseSnapshotStandbyConnectionStringTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails CreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeCreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails
	}{
		"AUTONOMOUS_CONTAINER_DATABASE",
		(MarshalTypeCreateDrProtectionGroupMemberAutonomousContainerDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
