// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CollectTracesData Information regarding the DB trace and alert log collection
type CollectTracesData struct {

	// Name of the bucket containing the file.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Object Storage namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Object name.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Status of trace collection process.
	CollectTracesState CollectTracesStatesEnum `mandatory:"true" json:"collectTracesState"`
}

func (m CollectTracesData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectTracesData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectTracesStatesEnum(string(m.CollectTracesState)); !ok && m.CollectTracesState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectTracesState: %s. Supported values are: %s.", m.CollectTracesState, strings.Join(GetCollectTracesStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
