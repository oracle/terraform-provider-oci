// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateExternalPublicationValidationDetails The task type contains the audit summary information and the definition of the task that is published externally.
type CreateExternalPublicationValidationDetails struct {

	// Generated key that can be used in API calls to identify the task. On scenarios where reference to the task is needed, a value can be passed in the create operation.
	Key *string `mandatory:"false" json:"key"`
}

func (m CreateExternalPublicationValidationDetails) String() string {
	return common.PointerString(m)
}
