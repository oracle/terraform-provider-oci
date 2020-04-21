// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// KafkaSettings Settings for the Kafka compatibility layer.
type KafkaSettings struct {

	// Bootstrap servers.
	BootstrapServers *string `mandatory:"false" json:"bootstrapServers"`

	// Enable auto creation of topic on the server.
	AutoCreateTopicsEnable *bool `mandatory:"false" json:"autoCreateTopicsEnable"`

	// The number of hours to keep a log file before deleting it (in hours).
	LogRetentionHours *int `mandatory:"false" json:"logRetentionHours"`

	// The default number of log partitions per topic.
	NumPartitions *int `mandatory:"false" json:"numPartitions"`
}

func (m KafkaSettings) String() string {
	return common.PointerString(m)
}
