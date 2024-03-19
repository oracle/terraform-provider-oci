// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkUnpinStatus Response of a bulk attribute unpin operation.
type BulkUnpinStatus struct {

	// We preserve the order of the attribute items from the bulk unpin request in this collection.  The ith object in this collection represents the
	// bulk unpin operation status of the ith object in the BulkUnpinAttributeDetails object in the Bulk unpin request.  If the
	// bulk unpin operation results in a processing error or a validation error, the operationStatus property in the  BulkUnpinMetadata object will
	// contain the appropriate bulk error status for the bulk operation.
	AttributeStatuses []AttributeUnpinResponse `mandatory:"true" json:"attributeStatuses"`

	BulkUnpinMetadata *BulkUnpinMetadata `mandatory:"true" json:"bulkUnpinMetadata"`
}

func (m BulkUnpinStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkUnpinStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
