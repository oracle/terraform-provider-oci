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

// Identity A container object for identity attributes.
// Example:
//   -----
//     {
//       "principalName": "ExampleName",
//       "principalId": "ocid1.user.oc1..<unique_ID>",
//       "authType": "natv",
//       "callerName": null,
//       "callerId": null,
//       "tenantId": "ocid1.tenancy.oc1..<unique_ID>",
//       "ipAddress": "172.24.80.88",
//       "credentials": null,
//       "userAgent": "Jersey/2.23 (HttpUrlConnection 1.8.0_212)",
//       "consoleSessionId": null
//     }
//   -----
type Identity struct {

	// The name of the user or service. This value is the friendly name associated with `principalId`.
	// Example: `ExampleName`
	PrincipalName *string `mandatory:"false" json:"principalName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the principal.
	PrincipalId *string `mandatory:"false" json:"principalId"`

	// The type of authentication used.
	// Example: `natv`
	AuthType *string `mandatory:"false" json:"authType"`

	// The name of the user or service. This value is the friendly name associated with `callerId`.
	CallerName *string `mandatory:"false" json:"callerName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the caller. The caller that made a
	// request on behalf of the prinicpal.
	CallerId *string `mandatory:"false" json:"callerId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenant.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The IP address of the source of the request.
	// Example: `172.24.80.88`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The credential ID of the user. This value is extracted from the HTTP 'Authorization' request
	// header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/).
	Credentials *string `mandatory:"false" json:"credentials"`

	// The user agent of the client that made the request.
	// Example: `Jersey/2.23 (HttpUrlConnection 1.8.0_212)`
	UserAgent *string `mandatory:"false" json:"userAgent"`

	// This value identifies any Console session associated with this request.
	ConsoleSessionId *string `mandatory:"false" json:"consoleSessionId"`
}

func (m Identity) String() string {
	return common.PointerString(m)
}
