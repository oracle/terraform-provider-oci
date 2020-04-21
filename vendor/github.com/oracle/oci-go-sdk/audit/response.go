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

// Response A container object for response attributes.
// Example:
//   -----
//     {
//       "status": "200",
//       "responseTime": "2019-09-18T00:10:59.278Z",
//       "headers": {
//         "ETag": [
//           "<unique_ID>"
//         ],
//         "Connection": [
//           "close"
//         ],
//         "Content-Length": [
//           "1828"
//         ],
//         "opc-request-id": [
//           "<unique_ID>"
//         ],
//         "Date": [
//           "Wed, 18 Sep 2019 00:10:59 GMT"
//         ],
//         "Content-Type": [
//           "application/json"
//         ]
//       },
//       "payload": {
//         "resourceName": "my_instance",
//         "id": "ocid1.instance.oc1.phx.<unique_ID>"
//       },
//       "message": null
//     }
//   -----
type Response struct {

	// The status code of the response.
	// Example: `200`
	Status *string `mandatory:"false" json:"status"`

	// The time of the response to the audited request, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-09-18T00:10:59.278Z`
	ResponseTime *common.SDKTime `mandatory:"false" json:"responseTime"`

	// The headers of the response.
	// Example:
	//   -----
	//     {
	//       "ETag": [
	//         "<unique_ID>"
	//       ],
	//       "Connection": [
	//         "close"
	//       ],
	//       "Content-Length": [
	//         "1828"
	//       ],
	//       "opc-request-id": [
	//         "<unique_ID>"
	//       ],
	//       "Date": [
	//         "Wed, 18 Sep 2019 00:10:59 GMT"
	//       ],
	//       "Content-Type": [
	//         "application/json"
	//       ]
	//     }
	//   -----
	Headers map[string][]string `mandatory:"false" json:"headers"`

	// This value is included for backward compatibility with the Audit version 1 schema, where
	// it contained metadata of interest from the response payload.
	// Example:
	//   -----
	//     {
	//       "resourceName": "my_instance",
	//       "id": "ocid1.instance.oc1.phx.<unique_ID>"
	//     }
	//   -----
	Payload map[string]interface{} `mandatory:"false" json:"payload"`

	// A friendly description of what happened during the operation. Use this for troubleshooting.
	Message *string `mandatory:"false" json:"message"`
}

func (m Response) String() string {
	return common.PointerString(m)
}
