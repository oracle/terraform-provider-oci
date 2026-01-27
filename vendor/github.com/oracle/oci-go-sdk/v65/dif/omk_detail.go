// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OmkDetail Details to reference an existing Oracle Managed Kubernetes environment.
type OmkDetail struct {

	// Unique identifier for an omk instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// OCID of existing OMK cluster-namespace.
	ClusterNamespaceId *string `mandatory:"true" json:"clusterNamespaceId"`

	// OCID of cluster assigned to OMK cluster-namespace.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// Kubernetes namespace-name of OMK cluster-namespace.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`
}

func (m OmkDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OmkDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
