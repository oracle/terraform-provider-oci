// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlTuningTaskSqlExecutionPlanStep A step in the SQL execution plan.
type SqlTuningTaskSqlExecutionPlanStep struct {

	// The numerical representation of the SQL execution plan.
	PlanHashValue *int64 `mandatory:"false" json:"planHashValue"`

	// The identification number of a step in the SQL execution plan. This is unique within the SQL execution plan.
	// This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	StepId *int `mandatory:"false" json:"stepId"`

	// The ID of the next step that operates on the results of this step.
	// This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ParentStepId *int `mandatory:"false" json:"parentStepId"`

	// The order of processing for steps with the same parent ID.
	Position *int `mandatory:"false" json:"position"`

	// The name of the operation performed at this step.
	Operation *string `mandatory:"false" json:"operation"`

	// The options used for the operation performed at this step.
	Options *string `mandatory:"false" json:"options"`

	// The current mode of the optimizer, such as all_rows, first_rows_n (where n = 1, 10, 100, 1000, and so on).
	OptimizerMode *string `mandatory:"false" json:"optimizerMode"`

	// The cost of the current operation estimated by the cost-based optimizer (CBO).
	Cost *float64 `mandatory:"false" json:"cost"`

	// The number of rows returned by the current operation (estimated by the CBO).
	Cardinality *int64 `mandatory:"false" json:"cardinality"`

	// The number of bytes returned by the current operation.
	Bytes *int64 `mandatory:"false" json:"bytes"`

	// The CPU cost of the current operation.
	CpuCost *float64 `mandatory:"false" json:"cpuCost"`

	// The I/O cost of the current operation.
	IoCost *float64 `mandatory:"false" json:"ioCost"`

	// The temporary space usage (in bytes) of the operation (sort or hash-join) as estimated by the CBO.
	TempSpace *int64 `mandatory:"false" json:"tempSpace"`

	// The elapsed time (in seconds) of the operation as estimated by the CBO.
	Time *int64 `mandatory:"false" json:"time"`

	// The name of the database link used to reference the object.
	ObjectNode *string `mandatory:"false" json:"objectNode"`

	// The owner of the object.
	ObjectOwner *string `mandatory:"false" json:"objectOwner"`

	// The name of the object.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The numbered position of the object name in the original SQL statement.
	ObjectPosition *int `mandatory:"false" json:"objectPosition"`

	// The descriptive modifier that further describes the type of object.
	ObjectType *string `mandatory:"false" json:"objectType"`

	// A step may get data from a range of partitions of a partitioned object, such as table or index,
	// based on predicates and sorting order. The partionStart is the starting partition of the range.
	// The partitionStop is the ending partition of the range.
	PartitionStart *string `mandatory:"false" json:"partitionStart"`

	// A step may get data from a range of partitions of a partitioned object, such as table or index,
	// based on predicates and sorting order. The partionStart is the starting partition of the range.
	// The partitionStop is the ending partition of the range.
	PartitionStop *string `mandatory:"false" json:"partitionStop"`

	// The ID of the step in the execution plan that has computed the pair of values of partitionStart and partitionStop.
	PartitionId *int `mandatory:"false" json:"partitionId"`

	// The place for comments that can be added to the steps of the execution plan.
	Remarks *string `mandatory:"false" json:"remarks"`

	// Number of index columns with start and stop keys (that is, the number of columns with matching predicates).
	NumberOfSearchColumn *int `mandatory:"false" json:"numberOfSearchColumn"`

	// Information about parallel execution servers and parallel queries
	Other *string `mandatory:"false" json:"other"`

	// Describes the function of the SQL text in the OTHER column.
	OtherTag *string `mandatory:"false" json:"otherTag"`

	// The text string identifying the type of execution plan.
	Attribute *string `mandatory:"false" json:"attribute"`

	// The predicates used to locate rows in an access structure. For example,
	// start or stop predicates for an index range scan.
	AccessPredicates *string `mandatory:"false" json:"accessPredicates"`

	// The predicates used to filter rows before producing them.
	FilterPredicates *string `mandatory:"false" json:"filterPredicates"`
}

func (m SqlTuningTaskSqlExecutionPlanStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningTaskSqlExecutionPlanStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
