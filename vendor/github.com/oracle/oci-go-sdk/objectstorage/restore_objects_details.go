// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object and Archive Storage APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RestoreObjectsDetails The representation of RestoreObjectsDetails
type RestoreObjectsDetails struct {

	// A object which was in an archived state and need to be restored.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m RestoreObjectsDetails) String() string {
	return common.PointerString(m)
}
