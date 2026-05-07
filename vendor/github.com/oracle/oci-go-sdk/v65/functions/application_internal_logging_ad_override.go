// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationInternalLoggingAdOverride Per-AD override values for internal Lumberjack logging configuration. Used when a region's ADs have different namespace and/or log group values for the same application.
// For payload examples, see the operation examples on /applications/{applicationId}/internalLogging.
type ApplicationInternalLoggingAdOverride struct {

	// Override Lumberjack namespace for this specific AD (only if different from the top-level namespace).
	Namespace *string `mandatory:"false" json:"namespace"`

	// The Lumberjack log group for this specific AD (only if different from the top-level logGroup).
	LogGroup *string `mandatory:"false" json:"logGroup"`
}

func (m ApplicationInternalLoggingAdOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationInternalLoggingAdOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
