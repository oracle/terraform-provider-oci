// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditEvent All the attributes of an audit event. For more information, see Viewing Audit Log Events (https://docs.oracle.com/iaas/Content/Audit/Tasks/viewinglogevents.htm).
type AuditEvent struct {

	// The type of event that happened.
	// The service that produces the event can also add, remove, or change the meaning of a field.
	// A service implementing these type changes would publish a new version of an `eventType` and
	// revise the `eventTypeVersion` field.
	// Example: `com.oraclecloud.ComputeApi.GetInstance`
	EventType *string `mandatory:"true" json:"eventType"`

	// The version of the CloudEvents specification. The structure of the envelope follows the
	// CloudEvents (https://github.com/cloudevents/spec) industry standard format hosted by the
	// Cloud Native Computing Foundation ( CNCF) (https://www.cncf.io/).
	// Audit uses version 0.1 specification of the CloudEvents event envelope.
	// Example: `0.1`
	CloudEventsVersion *string `mandatory:"true" json:"cloudEventsVersion"`

	// The version of the event type. This version applies to the payload of the event, not the envelope.
	// Use `cloudEventsVersion` to determine the version of the envelope.
	// Example: `2.0`
	EventTypeVersion *string `mandatory:"true" json:"eventTypeVersion"`

	// The source of the event.
	// Example: `ComputeApi`
	Source *string `mandatory:"true" json:"source"`

	// The GUID of the event.
	EventId *string `mandatory:"true" json:"eventId"`

	// The time the event occurred, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-09-18T00:10:59.252Z`
	EventTime *common.SDKTime `mandatory:"true" json:"eventTime"`

	// The content type of the data contained in `data`.
	// Example: `application/json`
	ContentType *string `mandatory:"true" json:"contentType"`

	Data *Data `mandatory:"true" json:"data"`
}

func (m AuditEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
