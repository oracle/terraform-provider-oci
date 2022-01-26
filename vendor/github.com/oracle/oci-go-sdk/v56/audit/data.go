// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.cloud.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Data The payload of the event. Information within `data` comes from the resource emitting the event.
type Data struct {

	// This value links multiple audit events that are part of the same API operation. For example,
	// a long running API operations that emit an event at the start and the end of an operation
	// would use the same value in this field for both events.
	EventGroupingId *string `mandatory:"false" json:"eventGroupingId"`

	// Name of the API operation that generated this event.
	// Example: `GetInstance`
	EventName *string `mandatory:"false" json:"eventName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the resource
	// emitting the event.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the compartment. This value is the friendly name associated with compartmentId.
	// This value can change, but the service logs the value that appeared at the time of the audit
	// event.
	// Example: `CompartmentA`
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The name of the resource emitting the event.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// An OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) or some other ID for the resource
	// emitting the event.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The availability domain where the resource resides.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name,
	// type, or namespace. Exists for cross-compatibility only. For more information,
	// see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Identity *Identity `mandatory:"false" json:"identity"`

	Request *Request `mandatory:"false" json:"request"`

	Response *Response `mandatory:"false" json:"response"`

	StateChange *StateChange `mandatory:"false" json:"stateChange"`

	// A container object for attribues unique to the resource emitting the event.
	// Example:
	//   -----
	//     {
	//       "imageId": "ocid1.image.oc1.phx.<unique_ID>",
	//       "shape": "VM.Standard1.1",
	//       "type": "CustomerVmi"
	//     }
	//   -----
	AdditionalDetails map[string]interface{} `mandatory:"false" json:"additionalDetails"`
}

func (m Data) String() string {
	return common.PointerString(m)
}
