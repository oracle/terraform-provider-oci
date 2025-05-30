// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeEvictionSettings Node Eviction Configuration
type NodeEvictionSettings struct {

	// Duration after which OKE will give up eviction of the pods on the node. PT0M will indicate you want to perform node action without cordon and drain.
	// Default PT60M, Min PT0M, Max: PT60M. Format ISO 8601 e.g PT30M
	EvictionGraceDuration *string `mandatory:"false" json:"evictionGraceDuration"`

	// If the node action should be performed if not all the pods can be evicted in the grace period
	IsForceActionAfterGraceDuration *bool `mandatory:"false" json:"isForceActionAfterGraceDuration"`
}

func (m NodeEvictionSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeEvictionSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
