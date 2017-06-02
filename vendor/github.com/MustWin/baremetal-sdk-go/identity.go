// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "fmt"

// Error is returned from unsuccessful API calls. The OPCRequestID if present
// is used to reference the failing requests for support.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/responses/Error
type Error struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	OPCRequestID string `json:"opc-request-id,omitempty"`
}

// Error returns a formatted description of an API error.
func (e *Error) Error() string {
	return fmt.Sprintf("Status: %s; Code: %s; OPC Request ID: %s; Message: %s",
		e.Status,
		e.Code,
		e.OPCRequestID,
		e.Message,
	)
}
