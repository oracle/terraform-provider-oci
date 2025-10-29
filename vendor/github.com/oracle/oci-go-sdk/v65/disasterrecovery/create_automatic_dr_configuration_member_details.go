// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateAutomaticDrConfigurationMemberDetails Create properties for a member in an Automatic DR configuration.
type CreateAutomaticDrConfigurationMemberDetails interface {

	// The OCID of the member.
	// Example: `ocid1.database.oc1..uniqueID`
	GetMemberId() *string
}

type createautomaticdrconfigurationmemberdetails struct {
	JsonData   []byte
	MemberId   *string `mandatory:"true" json:"memberId"`
	MemberType string  `json:"memberType"`
}

// UnmarshalJSON unmarshals json
func (m *createautomaticdrconfigurationmemberdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautomaticdrconfigurationmemberdetails createautomaticdrconfigurationmemberdetails
	s := struct {
		Model Unmarshalercreateautomaticdrconfigurationmemberdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MemberId = s.Model.MemberId
	m.MemberType = s.Model.MemberType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautomaticdrconfigurationmemberdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MemberType {
	case "AUTONOMOUS_DATABASE":
		mm := CreateAutomaticDrConfigurationMemberAutonomousDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := CreateAutomaticDrConfigurationMemberDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_CONTAINER_DATABASE":
		mm := CreateAutomaticDrConfigurationMemberAutonomousContainerDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateAutomaticDrConfigurationMemberDetails: %s.", m.MemberType)
		return *m, nil
	}
}

// GetMemberId returns MemberId
func (m createautomaticdrconfigurationmemberdetails) GetMemberId() *string {
	return m.MemberId
}

func (m createautomaticdrconfigurationmemberdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createautomaticdrconfigurationmemberdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
