// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BrokerShape Configuration of the broker node.
type BrokerShape struct {

	// Number of Kafka broker nodes
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// Number of OCPUs per nodes
	OcpuCount *int `mandatory:"true" json:"ocpuCount"`

	// Size of the storage per nodes.
	StorageSizeInGbs *int `mandatory:"false" json:"storageSizeInGbs"`
}

func (m BrokerShape) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BrokerShape) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
