// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDrProtectionGroupMemberDetails Create properties for a member in a DR protection group.
type CreateDrProtectionGroupMemberDetails interface {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	GetMemberId() *string
}

type createdrprotectiongroupmemberdetails struct {
	JsonData   []byte
	MemberId   *string `mandatory:"true" json:"memberId"`
	MemberType string  `json:"memberType"`
}

// UnmarshalJSON unmarshals json
func (m *createdrprotectiongroupmemberdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedrprotectiongroupmemberdetails createdrprotectiongroupmemberdetails
	s := struct {
		Model Unmarshalercreatedrprotectiongroupmemberdetails
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
func (m *createdrprotectiongroupmemberdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MemberType {
	case "COMPUTE_INSTANCE_MOVABLE":
		mm := CreateDrProtectionGroupMemberComputeInstanceMovableDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_NON_MOVABLE":
		mm := CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE":
		mm := CreateDrProtectionGroupMemberComputeInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := CreateDrProtectionGroupMemberDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := CreateDrProtectionGroupMemberAutonomousDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NETWORK_LOAD_BALANCER":
		mm := CreateDrProtectionGroupMemberNetworkLoadBalancerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VOLUME_GROUP":
		mm := CreateDrProtectionGroupMemberVolumeGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE_SYSTEM":
		mm := CreateDrProtectionGroupMemberFileSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER":
		mm := CreateDrProtectionGroupMemberLoadBalancerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateDrProtectionGroupMemberDetails: %s.", m.MemberType)
		return *m, nil
	}
}

// GetMemberId returns MemberId
func (m createdrprotectiongroupmemberdetails) GetMemberId() *string {
	return m.MemberId
}

func (m createdrprotectiongroupmemberdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdrprotectiongroupmemberdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
