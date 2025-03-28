// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeShapeNodes Individual worker nodes groups details.
type ChangeShapeNodes struct {

	// Change shape of worker nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	Worker *string `mandatory:"false" json:"worker"`

	WorkerShapeConfig *ShapeConfigDetails `mandatory:"false" json:"workerShapeConfig"`

	// Change shape of compute only worker nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	ComputeOnlyWorker *string `mandatory:"false" json:"computeOnlyWorker"`

	ComputeOnlyWorkerShapeConfig *ShapeConfigDetails `mandatory:"false" json:"computeOnlyWorkerShapeConfig"`

	// Change shape of master nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	Master *string `mandatory:"false" json:"master"`

	MasterShapeConfig *ShapeConfigDetails `mandatory:"false" json:"masterShapeConfig"`

	// Change shape of utility nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	Utility *string `mandatory:"false" json:"utility"`

	UtilityShapeConfig *ShapeConfigDetails `mandatory:"false" json:"utilityShapeConfig"`

	// Change shape of the Cloud SQL node to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	Cloudsql *string `mandatory:"false" json:"cloudsql"`

	CloudsqlShapeConfig *ShapeConfigDetails `mandatory:"false" json:"cloudsqlShapeConfig"`

	// Change shape of edge nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	Edge *string `mandatory:"false" json:"edge"`

	EdgeShapeConfig *ShapeConfigDetails `mandatory:"false" json:"edgeShapeConfig"`

	// Change shape of Kafka Broker nodes to the desired target shape. Both VM_STANDARD and E4 Flex shapes are allowed here.
	KafkaBroker *string `mandatory:"false" json:"kafkaBroker"`

	KafkaBrokerShapeConfig *ShapeConfigDetails `mandatory:"false" json:"kafkaBrokerShapeConfig"`
}

func (m ChangeShapeNodes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeShapeNodes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
