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

// Request A container object for request attributes.
// Example:
//   -----
//     {
//       "id": "<unique_ID>",
//       "path": "/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>",
//       "action": "GET",
//       "parameters": {},
//       "headers": {
//         "opc-principal": [
//           "{\"tenantId\":\"ocid1.tenancy.oc1..<unique_ID>\",\"subjectId\":\"ocid1.user.oc1..<unique_ID>\",\"claims\":[{\"key\":\"pstype\",\"value\":\"natv\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_host\",\"value\":\"iaas.r2.oracleiaas.com\",\"issuer\":\"h\"},{\"key\":\"h_opc-request-id\",\"value\":\"<unique_ID>\",\"issuer\":\"h\"},{\"key\":\"ptype\",\"value\":\"user\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_date\",\"value\":\"Wed, 18 Sep 2019 00:10:58 UTC\",\"issuer\":\"h\"},{\"key\":\"h_accept\",\"value\":\"application/json\",\"issuer\":\"h\"},{\"key\":\"authorization\",\"value\":\"Signature headers=\\\"date (request-target) host accept opc-request-id\\\",keyId=\\\"ocid1.tenancy.oc1..<unique_ID>/ocid1.user.oc1..<unique_ID>/8c:b4:5f:18:e7:ec:db:08:b8:fa:d2:2a:7d:11:76:ac\\\",algorithm=\\\"rsa-pss-sha256\\\",signature=\\\"<unique_ID>\\\",version=\\\"1\\\"\",\"issuer\":\"h\"},{\"key\":\"h_(request-target)\",\"value\":\"get /20160918/instances/ocid1.instance.oc1.phx.<unique_ID>\",\"issuer\":\"h\"}]}"
//         ],
//         "Accept": [
//           "application/json"
//         ],
//         "X-Oracle-Auth-Client-CN": [
//           "splat-proxy-se-02302.node.ad2.r2"
//         ],
//         "X-Forwarded-Host": [
//           "compute-api.svc.ad1.r2"
//         ],
//         "Connection": [
//           "close"
//         ],
//         "User-Agent": [
//           "Jersey/2.23 (HttpUrlConnection 1.8.0_212)"
//         ],
//         "X-Forwarded-For": [
//           "172.24.80.88"
//         ],
//         "X-Real-IP": [
//           "172.24.80.88"
//         ],
//         "oci-original-url": [
//           "https://iaas.r2.oracleiaas.com/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>"
//         ],
//         "opc-request-id": [
//           "<unique_ID>"
//         ],
//         "Date": [
//           "Wed, 18 Sep 2019 00:10:58 UTC"
//         ]
//       }
//     }
//   -----
type Request struct {

	// The opc-request-id of the request.
	Id *string `mandatory:"false" json:"id"`

	// The full path of the API request.
	// Example: `/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>`
	Path *string `mandatory:"false" json:"path"`

	// The HTTP method of the request.
	// Example: `GET`
	Action *string `mandatory:"false" json:"action"`

	// The parameters supplied by the caller during this operation.
	Parameters map[string][]string `mandatory:"false" json:"parameters"`

	// The HTTP header fields and values in the request.
	// Example:
	//   -----
	//     {
	//       "opc-principal": [
	//         "{\"tenantId\":\"ocid1.tenancy.oc1..<unique_ID>\",\"subjectId\":\"ocid1.user.oc1..<unique_ID>\",\"claims\":[{\"key\":\"pstype\",\"value\":\"natv\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_host\",\"value\":\"iaas.r2.oracleiaas.com\",\"issuer\":\"h\"},{\"key\":\"h_opc-request-id\",\"value\":\"<unique_ID>\",\"issuer\":\"h\"},{\"key\":\"ptype\",\"value\":\"user\",\"issuer\":\"authService.oracle.com\"},{\"key\":\"h_date\",\"value\":\"Wed, 18 Sep 2019 00:10:58 UTC\",\"issuer\":\"h\"},{\"key\":\"h_accept\",\"value\":\"application/json\",\"issuer\":\"h\"},{\"key\":\"authorization\",\"value\":\"Signature headers=\\\"date (request-target) host accept opc-request-id\\\",keyId=\\\"ocid1.tenancy.oc1..<unique_ID>/ocid1.user.oc1..<unique_ID>/8c:b4:5f:18:e7:ec:db:08:b8:fa:d2:2a:7d:11:76:ac\\\",algorithm=\\\"rsa-pss-sha256\\\",signature=\\\"<unique_ID>\\\",version=\\\"1\\\"\",\"issuer\":\"h\"},{\"key\":\"h_(request-target)\",\"value\":\"get /20160918/instances/ocid1.instance.oc1.phx.<unique_ID>\",\"issuer\":\"h\"}]}"
	//       ],
	//       "Accept": [
	//         "application/json"
	//       ],
	//       "X-Oracle-Auth-Client-CN": [
	//         "splat-proxy-se-02302.node.ad2.r2"
	//       ],
	//       "X-Forwarded-Host": [
	//         "compute-api.svc.ad1.r2"
	//       ],
	//       "Connection": [
	//         "close"
	//       ],
	//       "User-Agent": [
	//         "Jersey/2.23 (HttpUrlConnection 1.8.0_212)"
	//       ],
	//       "X-Forwarded-For": [
	//         "172.24.80.88"
	//       ],
	//       "X-Real-IP": [
	//         "172.24.80.88"
	//       ],
	//       "oci-original-url": [
	//         "https://iaas.r2.oracleiaas.com/20160918/instances/ocid1.instance.oc1.phx.<unique_ID>"
	//       ],
	//       "opc-request-id": [
	//         "<unique_ID>"
	//       ],
	//       "Date": [
	//         "Wed, 18 Sep 2019 00:10:58 UTC"
	//       ]
	//     }
	//   -----
	Headers map[string][]string `mandatory:"false" json:"headers"`
}

func (m Request) String() string {
	return common.PointerString(m)
}
