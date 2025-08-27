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

// KafkaClusterConfigVersionSummary Summary information about a KafkaClusterConfigVersion.
type KafkaClusterConfigVersionSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
	ConfigId *string `mandatory:"true" json:"configId"`

	// Version of the cluster configuration
	VersionNumber *int `mandatory:"true" json:"versionNumber"`

	// The date and time the KafkaClusterConfigVersion was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m KafkaClusterConfigVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KafkaClusterConfigVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
