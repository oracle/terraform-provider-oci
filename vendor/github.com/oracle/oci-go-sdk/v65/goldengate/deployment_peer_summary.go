// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeploymentPeerSummary The summary of the deployment Peer.
type DeploymentPeerSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the default region.
	Region *string `mandatory:"true" json:"region"`

	// The availability domain of a placement.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The fault domain of a placement.
	FaultDomain *string `mandatory:"true" json:"faultDomain"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the deployment peer.
	PeerType DeploymentPeerTypeEnum `mandatory:"true" json:"peerType"`

	// The type of the deployment role.
	PeerRole DeploymentRoleEnum `mandatory:"true" json:"peerRole"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time of the last role change. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeRoleChanged *common.SDKTime `mandatory:"true" json:"timeRoleChanged"`

	// Possible lifecycle states for deployment peer.
	LifecycleState DeploymentPeerSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m DeploymentPeerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentPeerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentPeerTypeEnum(string(m.PeerType)); !ok && m.PeerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerType: %s. Supported values are: %s.", m.PeerType, strings.Join(GetDeploymentPeerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentRoleEnum(string(m.PeerRole)); !ok && m.PeerRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerRole: %s. Supported values are: %s.", m.PeerRole, strings.Join(GetDeploymentRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentPeerSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeploymentPeerSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeploymentPeerSummaryLifecycleStateEnum Enum with underlying type: string
type DeploymentPeerSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentPeerSummaryLifecycleStateEnum
const (
	DeploymentPeerSummaryLifecycleStateCreating DeploymentPeerSummaryLifecycleStateEnum = "CREATING"
	DeploymentPeerSummaryLifecycleStateActive   DeploymentPeerSummaryLifecycleStateEnum = "ACTIVE"
	DeploymentPeerSummaryLifecycleStateFailed   DeploymentPeerSummaryLifecycleStateEnum = "FAILED"
	DeploymentPeerSummaryLifecycleStateUpdating DeploymentPeerSummaryLifecycleStateEnum = "UPDATING"
	DeploymentPeerSummaryLifecycleStateDeleting DeploymentPeerSummaryLifecycleStateEnum = "DELETING"
)

var mappingDeploymentPeerSummaryLifecycleStateEnum = map[string]DeploymentPeerSummaryLifecycleStateEnum{
	"CREATING": DeploymentPeerSummaryLifecycleStateCreating,
	"ACTIVE":   DeploymentPeerSummaryLifecycleStateActive,
	"FAILED":   DeploymentPeerSummaryLifecycleStateFailed,
	"UPDATING": DeploymentPeerSummaryLifecycleStateUpdating,
	"DELETING": DeploymentPeerSummaryLifecycleStateDeleting,
}

var mappingDeploymentPeerSummaryLifecycleStateEnumLowerCase = map[string]DeploymentPeerSummaryLifecycleStateEnum{
	"creating": DeploymentPeerSummaryLifecycleStateCreating,
	"active":   DeploymentPeerSummaryLifecycleStateActive,
	"failed":   DeploymentPeerSummaryLifecycleStateFailed,
	"updating": DeploymentPeerSummaryLifecycleStateUpdating,
	"deleting": DeploymentPeerSummaryLifecycleStateDeleting,
}

// GetDeploymentPeerSummaryLifecycleStateEnumValues Enumerates the set of values for DeploymentPeerSummaryLifecycleStateEnum
func GetDeploymentPeerSummaryLifecycleStateEnumValues() []DeploymentPeerSummaryLifecycleStateEnum {
	values := make([]DeploymentPeerSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentPeerSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentPeerSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DeploymentPeerSummaryLifecycleStateEnum
func GetDeploymentPeerSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"UPDATING",
		"DELETING",
	}
}

// GetMappingDeploymentPeerSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentPeerSummaryLifecycleStateEnum(val string) (DeploymentPeerSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDeploymentPeerSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
