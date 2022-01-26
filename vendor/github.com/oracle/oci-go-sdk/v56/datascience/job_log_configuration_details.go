// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// JobLogConfigurationDetails Logging configuration for resource.
type JobLogConfigurationDetails struct {

	// If customer logging is enabled for job runs.
	EnableLogging *bool `mandatory:"false" json:"enableLogging"`

	// If automatic on-behalf-of log object creation is enabled for job runs.
	EnableAutoLogCreation *bool `mandatory:"false" json:"enableAutoLogCreation"`

	// The log group id for where log objects are for job runs.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// The log id the job run will push logs too.
	LogId *string `mandatory:"false" json:"logId"`
}

func (m JobLogConfigurationDetails) String() string {
	return common.PointerString(m)
}
