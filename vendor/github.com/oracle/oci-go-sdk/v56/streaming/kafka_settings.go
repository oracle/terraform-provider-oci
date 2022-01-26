// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
