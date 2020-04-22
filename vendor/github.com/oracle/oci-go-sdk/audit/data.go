// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// Data The payload of the event. Information within `data` comes from the resource emitting the event.
// Example:
//   -----
//     {
//       "eventGroupingId": null,
//       "eventName": "GetInstance",
//       "compartmentId": "ocid1.tenancy.oc1..<unique_ID>",
//       "compartmentName": "compartmentA",
//       "resourceName": "my_instance",
//       "resourceId": "ocid1.instance.oc1.phx.<unique_ID>",
//       "availabilityDomain": "<availability_domain>",
//       "freeformTags": null,
//       "definedTags": null,
//       "identity": {
//         "principalName": "ExampleName",
//         "principalId": "ocid1.user.oc1..<unique_ID>",
//         "authType": "natv",
//         "callerName": null,
//         "callerId": null,
//         "tenantId": "ocid1.tenancy.oc1..<unique_ID>",
//         "ipAddress": "172.24.80.88",
//         "credentials": null,
//         "userAgent": "Jersey/2.23 (HttpUrlConnection 1.8.0_212)",
//         "consoleSessionId": null
//       },
//       "request": {
//         "id": "<unique_ID>",
//         "path": "/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>",
//         "action": "GET",
//         "parameters": {},
//         "headers": {
//           "opc-principal": [
//             "{\"tenantId\":\"ocid1.tenancy.oc1..<unique_ID>\",\"subjectId\":\"ocid1.user.oc1..<unique_ID>\",\"claims\":[{\"key\":\"pstype\",\"value\":\"natv\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_host\",\"value\":\"iaas.r2.oracleiaas.com\",\"issuer\":\"h\"},{\"key\":\"h_opc-request-id\",\"value\":\"<unique_ID>\",\"issuer\":\"h\"},{\"key\":\"ptype\",\"value\":\"user\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_date\",\"value\":\"Wed, 18 Sep 2019 00:10:58 UTC\",\"issuer\":\"h\"},{\"key\":\"h_accept\",\"value\":\"application/json\",\"issuer\":\"h\"},{\"key\":\"authorization\",\"value\":\"Signature headers=\\\"date (request-target) host accept opc-request-id\\\",keyId=\\\"ocid1.tenancy.oc1..<unique_ID>/ocid1.user.oc1..<unique_ID>/8c:b4:5f:18:e7:ec:db:08:b8:fa:d2:2a:7d:11:76:ac\\\",algorithm=\\\"rsa-pss-sha256\\\",signature=\\\"<unique_ID>\\\",version=\\\"1\\\"\",\"issuer\":\"h\"},{\"key\":\"h_(request-target)\",\"value\":\"get /20160918/instances/ocid1.instance.oc1.phx.<unique_ID>\",\"issuer\":\"h\"}]}"
//           ],
//           "Accept": [
//             "application/json"
//           ],
//           "X-Oracle-Auth-Client-CN": [
//             "splat-proxy-se-02302.node.ad2.r2"
//           ],
//           "X-Forwarded-Host": [
//             "compute-api.svc.ad1.r2"
//           ],
//           "Connection": [
//             "close"
//           ],
//           "User-Agent": [
//             "Jersey/2.23 (HttpUrlConnection 1.8.0_212)"
//           ],
//           "X-Forwarded-For": [
//             "172.24.80.88"
//           ],
//           "X-Real-IP": [
//             "172.24.80.88"
//           ],
//           "oci-original-url": [
//             "https://iaas.r2.oracleiaas.com/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>"
//           ],
//           "opc-request-id": [
//             "<unique_ID>"
//           ],
//           "Date": [
//             "Wed, 18 Sep 2019 00:10:58 UTC"
//           ]
//         }
//       },
//       "response": {
//         "status": "200",
//         "responseTime": "2019-09-18T00:10:59.278Z",
//         "headers": {
//           "ETag": [
//             "<unique_ID>"
//           ],
//           "Connection": [
//             "close"
//           ],
//           "Content-Length": [
//             "1828"
//           ],
//           "opc-request-id": [
//             "<unique_ID>"
//           ],
//           "Date": [
//             "Wed, 18 Sep 2019 00:10:59 GMT"
//           ],
//           "Content-Type": [
//             "application/json"
//           ]
//         },
//         "payload": {
//           "resourceName": "my_instance",
//           "id": "ocid1.instance.oc1.phx.<unique_ID>"
//         },
//         "message": null
//       },
//       "stateChange": {
//         "previous": null,
//         "current": null
//       },
//       "additionalDetails": {
//         "imageId": "ocid1.image.oc1.phx.<unique_ID>",
//         "shape": "VM.Standard1.1",
//         "type": "CustomerVmi"
//       }
//     }
//   -----
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
