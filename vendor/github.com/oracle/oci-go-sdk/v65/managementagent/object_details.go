// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectDetails Details of the Objectstorage object
type ObjectDetails struct {

	// Objectstorage namespace reference providing the original location of this object
	ObjectNamespace *string `mandatory:"true" json:"objectNamespace"`

	// Objectstorage bucket reference providing the original location of this object
	ObjectBucket *string `mandatory:"true" json:"objectBucket"`

	// Objectstorage object name reference providing the original location of this object
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Object storage URL for download
	ObjectUrl *string `mandatory:"false" json:"objectUrl"`

	// Object content SHA256 Hash
	Checksum *string `mandatory:"false" json:"checksum"`
}

func (m ObjectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
