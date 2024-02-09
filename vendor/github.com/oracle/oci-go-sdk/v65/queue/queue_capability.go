// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"strings"
)

// QueueCapabilityEnum Enum with underlying type: string
type QueueCapabilityEnum string

// Set of constants representing the allowable values for QueueCapabilityEnum
const (
	QueueCapabilityConsumerGroups QueueCapabilityEnum = "CONSUMER_GROUPS"
	QueueCapabilityLargeMessages  QueueCapabilityEnum = "LARGE_MESSAGES"
)

var mappingQueueCapabilityEnum = map[string]QueueCapabilityEnum{
	"CONSUMER_GROUPS": QueueCapabilityConsumerGroups,
	"LARGE_MESSAGES":  QueueCapabilityLargeMessages,
}

var mappingQueueCapabilityEnumLowerCase = map[string]QueueCapabilityEnum{
	"consumer_groups": QueueCapabilityConsumerGroups,
	"large_messages":  QueueCapabilityLargeMessages,
}

// GetQueueCapabilityEnumValues Enumerates the set of values for QueueCapabilityEnum
func GetQueueCapabilityEnumValues() []QueueCapabilityEnum {
	values := make([]QueueCapabilityEnum, 0)
	for _, v := range mappingQueueCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetQueueCapabilityEnumStringValues Enumerates the set of values in String for QueueCapabilityEnum
func GetQueueCapabilityEnumStringValues() []string {
	return []string{
		"CONSUMER_GROUPS",
		"LARGE_MESSAGES",
	}
}

// GetMappingQueueCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueueCapabilityEnum(val string) (QueueCapabilityEnum, bool) {
	enum, ok := mappingQueueCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
