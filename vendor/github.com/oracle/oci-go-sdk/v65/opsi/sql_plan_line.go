// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlPlanLine SQL Plan Line type object.
type SqlPlanLine struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Plan hash value for the SQL Execution Plan
	PlanHash *int64 `mandatory:"true" json:"planHash"`

	// Collection time stamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Operation
	// Example: `"SELECT STATEMENT"`
	Operation *string `mandatory:"true" json:"operation"`

	// Identifier
	// Example: `3`
	Identifier *int64 `mandatory:"true" json:"identifier"`

	// Version
	// Example: `1`
	Version *float32 `mandatory:"false" json:"version"`

	// Force matching signature
	// Example: `"18067345456756876713"`
	ForceMatchingSignature *string `mandatory:"false" json:"forceMatchingSignature"`

	// Generation time stamp
	// Example: `"2020-05-05T02:10:00.000Z"`
	TimeGenerated *common.SDKTime `mandatory:"false" json:"timeGenerated"`

	// Remark
	// Example: `""`
	Remark *string `mandatory:"false" json:"remark"`

	// Options
	// Example: `"RANGE SCAN"`
	Options *string `mandatory:"false" json:"options"`

	// Object Node
	// Example: `"Q4000"`
	ObjectNode *string `mandatory:"false" json:"objectNode"`

	// Object Owner
	// Example: `"TENANT_A#SCHEMA"`
	ObjectOwner *string `mandatory:"false" json:"objectOwner"`

	// Object Name
	// Example: `"PLAN_LINES_PK"`
	ObjectName *string `mandatory:"false" json:"objectName"`

	// Object Alias
	// Example: `"PLAN_LINES@SEL$1"`
	ObjectAlias *string `mandatory:"false" json:"objectAlias"`

	// Object Instance
	// Example: `37472`
	ObjectInstance *int64 `mandatory:"false" json:"objectInstance"`

	// Object Type
	// Example: `"INDEX (UNIQUE)"`
	ObjectType *string `mandatory:"false" json:"objectType"`

	// Optimizer
	// Example: `"CLUSTER"`
	Optimizer *string `mandatory:"false" json:"optimizer"`

	// Search Columns
	// Example: `3`
	SearchColumns *int64 `mandatory:"false" json:"searchColumns"`

	// Parent Identifier
	// Example: `2`
	ParentIdentifier *int64 `mandatory:"false" json:"parentIdentifier"`

	// Depth
	// Example: `3`
	Depth *int64 `mandatory:"false" json:"depth"`

	// Position
	// Example: `1`
	Position *int64 `mandatory:"false" json:"position"`

	// Cost
	// Example: `1`
	Cost *int64 `mandatory:"false" json:"cost"`

	// Cardinality
	// Example: `1`
	Cardinality *int64 `mandatory:"false" json:"cardinality"`

	// Bytes
	// Example: `150`
	Bytes *int64 `mandatory:"false" json:"bytes"`

	// Other
	// Example: ``
	Other *string `mandatory:"false" json:"other"`

	// Other Tag
	// Example: `"PARALLEL_COMBINED_WITH_PARENT"`
	OtherTag *string `mandatory:"false" json:"otherTag"`

	// Partition start
	// Example: `1`
	PartitionStart *string `mandatory:"false" json:"partitionStart"`

	// Partition stop
	// Example: `2`
	PartitionStop *string `mandatory:"false" json:"partitionStop"`

	// Partition identifier
	// Example: `8`
	PartitionIdentifier *int64 `mandatory:"false" json:"partitionIdentifier"`

	// Distribution
	// Example: `"QC (RANDOM)"`
	Distribution *string `mandatory:"false" json:"distribution"`

	// CPU cost
	// Example: `7321`
	CpuCost *int64 `mandatory:"false" json:"cpuCost"`

	// IO cost
	// Example: `1`
	IoCost *int64 `mandatory:"false" json:"ioCost"`

	// Time space
	// Example: `15614000`
	TempSpace *int64 `mandatory:"false" json:"tempSpace"`

	// Access predicates
	// Example: `"\"RESOURCE_ID\"=:1 AND \"QUERY_ID\"=:2"`
	AccessPredicates *string `mandatory:"false" json:"accessPredicates"`

	// Filter predicates
	// Example: `"(INTERNAL_FUNCTION(\"J\".\"DATABASE_ROLE\") OR (\"J\".\"DATABASE_ROLE\" IS NULL AND SYS_CONTEXT('userenv','database_role')='PRIMARY'))"`
	FilterPredicates *string `mandatory:"false" json:"filterPredicates"`

	// Projection
	// Example: `"COUNT(*)[22]"`
	Projection *string `mandatory:"false" json:"projection"`

	// Qblock Name
	// Example: `"SEL$1"`
	QblockName *string `mandatory:"false" json:"qblockName"`

	// Total elapsed time
	// Example: `1.2`
	ElapsedTimeInSec *float32 `mandatory:"false" json:"elapsedTimeInSec"`

	// Other SQL
	// Example: `"<other_xml><info type=\"db_version\">18.0.0.0</info><info type=\"parse_schema\"><![CDATA[\"SYS\"]]></info><info type=\"plan_hash_full\">483892784</info><info type=\"plan_hash\">2709293936</info><info type=\"plan_hash_2\">483892784</info><outline_data><hint><![CDATA[IGNORE_OPTIM_EMBEDDED_HINTS]]></hint><hint><![CDATA[OPTIMIZER_FEATURES_ENABLE('18.1.0')]]></hint><hint><![CDATA[DB_VERSION('18.1.0')]]></hint><hint><![CDATA[OPT_PARAM('_b_tree_bitmap_plans' 'false')]]></hint><hint><![CDATA[OPT_PARAM('_optim_peek_user_binds' 'false')]]></hint><hint><![CDATA[OPT_PARAM('result_cache_mode' 'FORCE')]]></hint><hint><![CDATA[OPT_PARAM('_fix_control' '20648883:0 27745220:1 30001331:1 30142527:1 30539126:1')]]></hint><hint><![CDATA[OUTLINE_LEAF(@\"SEL$1\")]]></hint><hint><![CDATA[INDEX(@\"SEL$1\" \"USER$\"@\"SEL$1\" \"I_USER#\")]]></hint></outline_data></other_xml>"`
	OtherXML *string `mandatory:"false" json:"otherXML"`
}

func (m SqlPlanLine) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanLine) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
