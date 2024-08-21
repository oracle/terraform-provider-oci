// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.cloud.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestLogEntryCollection A list of logs for the specified work request.
type WorkRequestLogEntryCollection struct {

	// A collection of work request logs.
	Items []WorkRequestLogEntry `mandatory:"true" json:"items"`
}

func (m WorkRequestLogEntryCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestLogEntryCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
