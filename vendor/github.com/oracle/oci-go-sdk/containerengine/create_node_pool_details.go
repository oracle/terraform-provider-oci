// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateNodePoolDetails The properties that define a request to create a node pool.
type CreateNodePoolDetails struct {

	// The OCID of the compartment in which the node pool exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the cluster to which this node pool is attached.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// The name of the node pool. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The version of Kubernetes to install on the nodes in the node pool.
	KubernetesVersion *string `mandatory:"true" json:"kubernetesVersion"`

	// The name of the image running on the nodes in the node pool.
	NodeImageName *string `mandatory:"true" json:"nodeImageName"`

	// The name of the node shape of the nodes in the node pool.
	NodeShape *string `mandatory:"true" json:"nodeShape"`

	// A list of key/value pairs to add to each underlying OCI instance in the node pool.
	NodeMetadata map[string]string `mandatory:"false" json:"nodeMetadata"`

	// A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	InitialNodeLabels []KeyValue `mandatory:"false" json:"initialNodeLabels"`

	// The SSH public key to add to each node in the node pool.
	SshPublicKey *string `mandatory:"false" json:"sshPublicKey"`

	// Optional, default to 1. The number of nodes to create in each subnet specified in subnetIds property.
	// When used, subnetIds is required. This property is deprecated, use nodeConfigDetails instead.
	QuantityPerSubnet *int `mandatory:"false" json:"quantityPerSubnet"`

	// The OCIDs of the subnets in which to place nodes for this node pool. When used, quantityPerSubnet
	// can be provided. This property is deprecated, use nodeConfigDetails. Exactly one of the
	// subnetIds or nodeConfigDetails properties must be specified.
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// The configuration of nodes in the node pool. Exactly one of the
	// subnetIds or nodeConfigDetails properties must be specified.
	NodeConfigDetails *CreateNodePoolNodeConfigDetails `mandatory:"false" json:"nodeConfigDetails"`
}

func (m CreateNodePoolDetails) String() string {
	return common.PointerString(m)
}
