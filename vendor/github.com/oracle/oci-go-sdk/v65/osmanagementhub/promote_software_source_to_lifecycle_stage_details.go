// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PromoteSoftwareSourceToLifecycleStageDetails A versioned custom software source OCID (softwareSourceId)
// is required when promoting software source content to
// lifecycle stage rank one. Software source content must be
// promoted to lifecycle stage rank one before being
// eligible for promotion to subsequent lifecycle stages,
// else an error is returned. Software source content is
// expected to be promoted in order starting with
// lifecycle stage rank one, followed by rank two, then rank
// three and so on.
// When promoting software source content to lifecycle stage
// rank two, three, four or five, softwareSourceId is optional.
// If a softwareSourceId is provided for a lifecycle stage
// between two and five, the system validates that the
// softwareSourceId is already promoted to the previous lifecycle stage.
// If the softwareSourceId from the previous lifecycle stage
// does not match the provided softwareSourceId an error returns.
// If a softwareSourceId is not provided for a lifecycle stage
// between two and five, the system promotes the
// softwareSourceId from the previous lifecycle stage. If the
// previous lifecycle stage has no SourceSource content
// an error returns.
type PromoteSoftwareSourceToLifecycleStageDetails struct {
	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
}

func (m PromoteSoftwareSourceToLifecycleStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PromoteSoftwareSourceToLifecycleStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
