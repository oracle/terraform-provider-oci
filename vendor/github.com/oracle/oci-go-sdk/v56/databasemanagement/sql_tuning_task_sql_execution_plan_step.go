// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SqlTuningTaskSqlExecutionPlanStep A step of a SQL execution plan.
type SqlTuningTaskSqlExecutionPlanStep struct {

	// Numerical representation of the execution plan
	PlanHashValue *int64 `mandatory:"false" json:"planHashValue"`

	// Identification number for this step in the execution plan. It is unique within the execution plan.
	// It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	StepId *int `mandatory:"false" json:"stepId"`

	// ID of the next step that operates on the results of this step.
	// It is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ParentStepId *int `mandatory:"false" json:"parentStepId"`

	// Order of processing for steps with the same parent ID.
	Position *int `mandatory:"false" json:"position"`

	// Name of the operation performed at this step
	Operation *string `mandatory:"false" json:"operation"`

	// Options used for the operation performed at this step.
	Options *string `mandatory:"false" json:"options"`

	// Current mode of the optimizer, such as all_rows, first_rows_n (where n = 1, 10, 100, 1000 etc).
	OptimizerMode *string `mandatory:"false" json:"optimizerMode"`

	// Cost of the current operation estimated by the cost-based optimizer (CBO).
	Cost *float64 `mandatory:"false" json:"cost"`

	// Number of rows returned by the current operation (estimated by the CBO).
	Cardinality *int64 `mandatory:"false" json:"cardinality"`

	// Number of bytes returned by the current operation.
	Bytes *int64 `mandatory:"false" json:"bytes"`

	// The CPU cost of the current operation.
	CpuCost *float64 `mandatory:"false" json:"cpuCost"`

	// The I/O cost of the current operation.
	IoCost *float64 `mandatory:"false" json:"ioCost"`

	// Temporary space usage (in bytes) of the operation (sort or hash-join) as estimated by the CBO.
	TempSpace *int64 `mandatory:"false" json:"tempSpace"`

	// Elapsed time (in seconds) of the operation as estimated by the CBO.
	Time *int64 `mandatory:"false" json:"time"`

	// Name of the database link used to reference the object.
	ObjectNode *string `mandatory:"false" json:"objectNode"`

	// Owner of the object.
	ObjectOwner *string `mandatory:"false" json:"objectOwner"`

	// Name of the object.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// Numbered position of the object name in the original SQL statement.
	ObjectPosition *int `mandatory:"false" json:"objectPosition"`

	// Descriptive modifier that further describes the type of object.
	ObjectType *string `mandatory:"false" json:"objectType"`

	// A step may get data from a range of partitions of a partitioned object, such table or index,
	// based on predicates and sorting order. The partionStart is the starting partition of the range.
	// The partitionStop is the ending partition of the range
	PartitionStart *string `mandatory:"false" json:"partitionStart"`

	// A step may get data from a range of partitions of a partitioned object, such table or index,
	// based on predicates and sorting order. The partionStart is the starting partition of the range.
	// The partitionStop is the ending partition of the range
	PartitionStop *string `mandatory:"false" json:"partitionStop"`

	// The id of the step in the execution plan that has computed the pair of values of the partitionStart and partitionStop
	PartitionId *int `mandatory:"false" json:"partitionId"`

	// Place for comments that can be added to the steps of the execution plan.
	Remarks *string `mandatory:"false" json:"remarks"`

	// Number of index columns with start and stop keys (that is, the number of columns with matching predicates)
	NumberOfSearchColumn *int `mandatory:"false" json:"numberOfSearchColumn"`

	// Information about parallel execution servers and parallel queries
	Other *string `mandatory:"false" json:"other"`

	// Describes the function of the SQL text in the OTHER column.
	OtherTag *string `mandatory:"false" json:"otherTag"`

	// Text string identifying the type of the execution plan.
	Attribute *string `mandatory:"false" json:"attribute"`

	// Predicates used to locate rows in an access structure. For example,
	// start or stop predicates for an index range scan.
	AccessPredicates *string `mandatory:"false" json:"accessPredicates"`

	// Predicates used to filter rows before producing them.
	FilterPredicates *string `mandatory:"false" json:"filterPredicates"`
}

func (m SqlTuningTaskSqlExecutionPlanStep) String() string {
	return common.PointerString(m)
}
