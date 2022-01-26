// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// NodePoolNodeConfigDetails The size and placement configuration of nodes in the node pool.
type NodePoolNodeConfigDetails struct {

	// The number of nodes in the node pool.
	Size *int `mandatory:"false" json:"size"`

	// The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID of the Key Management Service key assigned to the boot volume.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// The placement configurations for the node pool. Provide one placement
	// configuration for each availability domain in which you intend to launch a node.
	// To use the node pool with a regional subnet, provide a placement configuration for
	// each availability domain, and include the regional subnet in each placement
	// configuration.
	PlacementConfigs []NodePoolPlacementConfigDetails `mandatory:"false" json:"placementConfigs"`
}

func (m NodePoolNodeConfigDetails) String() string {
	return common.PointerString(m)
}
