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

// CreateDrProtectionGroupMemberOkeClusterDetails Create properties for an OKE member.
type CreateDrProtectionGroupMemberOkeClusterDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The OCID of the peer OKE cluster.
	// This property applies to the OKE cluster member in both the primary and standby region.
	// Example: `ocid1.cluster.oc1..uniqueID`
	PeerClusterId *string `mandatory:"false" json:"peerClusterId"`

	// The OCID of the compute instance member that is designated as a jump host.
	// This compute instance will be used to perform DR operations on the cluster using Oracle Cloud Agent's Run Command feature.
	// Example: `ocid1.instance.oc1..uniqueID`
	JumpHostId *string `mandatory:"false" json:"jumpHostId"`

	BackupLocation *CreateOkeBackupLocationDetails `mandatory:"false" json:"backupLocation"`

	BackupConfig *CreateOkeClusterBackupConfigDetails `mandatory:"false" json:"backupConfig"`

	// The list of source-to-destination load balancer mappings required for DR operations.
	// This property applies to the OKE cluster member in primary region.
	LoadBalancerMappings []CreateOkeClusterLoadBalancerMappingDetails `mandatory:"false" json:"loadBalancerMappings"`

	// The list of source-to-destination network load balancer mappings required for DR operations.
	// This property applies to the OKE cluster member in primary region.
	NetworkLoadBalancerMappings []CreateOkeClusterNetworkLoadBalancerMappingDetails `mandatory:"false" json:"networkLoadBalancerMappings"`

	// The list of source-to-destination vault mappings required for DR operations.
	// This property applies to the OKE cluster member in primary region.
	VaultMappings []CreateOkeClusterVaultMappingDetails `mandatory:"false" json:"vaultMappings"`

	// The list of managed node pools with configurations for minimum and maximum node counts.
	// This property applies to the OKE cluster member in both the primary and standby region.
	ManagedNodePoolConfigs []CreateOkeClusterManagedNodePoolConfigurationDetails `mandatory:"false" json:"managedNodePoolConfigs"`

	// The list of virtual node pools with configurations for minimum and maximum node counts.
	// This property applies to the OKE cluster member in both the primary and standby region.
	VirtualNodePoolConfigs []CreateOkeClusterVirtualNodePoolConfigurationDetails `mandatory:"false" json:"virtualNodePoolConfigs"`

	// The list of config maps along with their corresponding namespaces.
	// This property applies to the OKE cluster member in primary region.
	ResourceModifierMappings []CreateOkeClusterResourceModifierMappingDetails `mandatory:"false" json:"resourceModifierMappings"`
}

// GetMemberId returns MemberId
func (m CreateDrProtectionGroupMemberOkeClusterDetails) GetMemberId() *string {
	return m.MemberId
}

func (m CreateDrProtectionGroupMemberOkeClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupMemberOkeClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDrProtectionGroupMemberOkeClusterDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDrProtectionGroupMemberOkeClusterDetails CreateDrProtectionGroupMemberOkeClusterDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeCreateDrProtectionGroupMemberOkeClusterDetails
	}{
		"OKE_CLUSTER",
		(MarshalTypeCreateDrProtectionGroupMemberOkeClusterDetails)(m),
	}

	return json.Marshal(&s)
}
