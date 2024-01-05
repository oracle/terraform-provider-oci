// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageScriptLocation The details of an object storage script location for a user-defined step in a DR plan.
type ObjectStorageScriptLocation struct {

	// The namespace in object storage (Note - this is usually the tenancy name).
	// Example: `myocitenancy`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The bucket name inside the object storage namespace.
	// Example: `custom_dr_scripts`
	Bucket *string `mandatory:"true" json:"bucket"`

	// The object name inside the object storage bucket.
	// Example: `validate_app_start.sh`
	Object *string `mandatory:"true" json:"object"`
}

func (m ObjectStorageScriptLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageScriptLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
