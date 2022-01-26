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

// TablespaceSummary The summary of a tablespace.
type TablespaceSummary struct {

	// The name of the tablespace.
	Name *string `mandatory:"true" json:"name"`

	// The type of tablespace.
	Type TablespaceSummaryTypeEnum `mandatory:"true" json:"type"`

	// The status of the tablespace.
	Status TablespaceSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The tablespace block size.
	BlockSizeBytes *float32 `mandatory:"false" json:"blockSizeBytes"`

	// The default logging attribute.
	Logging TablespaceSummaryLoggingEnum `mandatory:"false" json:"logging,omitempty"`

	// Indicates whether the tablespace is under Force Logging mode.
	IsForceLogging *bool `mandatory:"false" json:"isForceLogging"`

	// Indicates whether the extents in the tablespace are Locally managed or Dictionary managed.
	ExtentManagement TablespaceSummaryExtentManagementEnum `mandatory:"false" json:"extentManagement,omitempty"`

	// The type of extent allocation in effect for the tablespace.
	AllocationType TablespaceSummaryAllocationTypeEnum `mandatory:"false" json:"allocationType,omitempty"`

	// Indicates whether the tablespace is plugged in.
	IsPluggedIn *bool `mandatory:"false" json:"isPluggedIn"`

	// Indicates whether the free and used segment space in the tablespace is managed using free lists (MANUAL) or bitmaps (AUTO).
	SegmentSpaceManagement TablespaceSummarySegmentSpaceManagementEnum `mandatory:"false" json:"segmentSpaceManagement,omitempty"`

	// Indicates whether default table compression is enabled or disabled.
	DefaultTableCompression TablespaceSummaryDefaultTableCompressionEnum `mandatory:"false" json:"defaultTableCompression,omitempty"`

	// Indicates whether undo retention guarantee is enabled for the tablespace.
	Retention TablespaceSummaryRetentionEnum `mandatory:"false" json:"retention,omitempty"`

	// Indicates whether the tablespace is a Bigfile tablespace or a Smallfile tablespace.
	IsBigfile *bool `mandatory:"false" json:"isBigfile"`

	// Indicates whether predicates are evaluated by Host or by Storage.
	PredicateEvaluation TablespaceSummaryPredicateEvaluationEnum `mandatory:"false" json:"predicateEvaluation,omitempty"`

	// Indicates whether the tablespace is encrypted.
	IsEncrypted *bool `mandatory:"false" json:"isEncrypted"`

	// The operation type for which default compression is enabled.
	CompressFor TablespaceSummaryCompressForEnum `mandatory:"false" json:"compressFor,omitempty"`

	// Indicates whether the In-Memory Column Store (IM column store) is by default enabled or disabled for tables in the tablespace.
	DefaultInMemory TablespaceSummaryDefaultInMemoryEnum `mandatory:"false" json:"defaultInMemory,omitempty"`

	// Indicates the default priority for In-Memory Column Store (IM column store) population for the tablespace.
	DefaultInMemoryPriority TablespaceSummaryDefaultInMemoryPriorityEnum `mandatory:"false" json:"defaultInMemoryPriority,omitempty"`

	// Indicates how the IM column store is distributed by default for the tablespace in an Oracle Real Application Clusters (Oracle RAC) environment.
	DefaultInMemoryDistribute TablespaceSummaryDefaultInMemoryDistributeEnum `mandatory:"false" json:"defaultInMemoryDistribute,omitempty"`

	// Indicates the default compression level for the IM column store for the tablespace.
	DefaultInMemoryCompression TablespaceSummaryDefaultInMemoryCompressionEnum `mandatory:"false" json:"defaultInMemoryCompression,omitempty"`

	// Indicates the duplicate setting for the IM column store in an Oracle RAC environment.
	DefaultInMemoryDuplicate TablespaceSummaryDefaultInMemoryDuplicateEnum `mandatory:"false" json:"defaultInMemoryDuplicate,omitempty"`

	// Indicates whether the tablespace is for shared tablespace, or for local temporary tablespace for leaf (read-only) instances, or for local temporary tablespace for all instance types.
	Shared TablespaceSummarySharedEnum `mandatory:"false" json:"shared,omitempty"`

	// Indicates whether default index compression is enabled or disabled.
	DefaultIndexCompression TablespaceSummaryDefaultIndexCompressionEnum `mandatory:"false" json:"defaultIndexCompression,omitempty"`

	// The operation type for which default index compression is enabled.
	IndexCompressFor TablespaceSummaryIndexCompressForEnum `mandatory:"false" json:"indexCompressFor,omitempty"`

	// This specifies the default value for the CELLMEMORY attribute that tables created in the tablespace will inherit unless the behavior is overridden explicitly. This column is intended for use with Oracle Exadata.
	DefaultCellMemory *string `mandatory:"false" json:"defaultCellMemory"`

	// Indicates how the IM column store is populated on various instances by default for the tablespace.
	DefaultInMemoryService TablespaceSummaryDefaultInMemoryServiceEnum `mandatory:"false" json:"defaultInMemoryService,omitempty"`

	// Indicates the service name for the service on which the IM column store should be populated by default for the tablespace. This column has a value only when the corresponding DEF_INMEMORY_SERVICE is USER_DEFINED. In all other cases, this column is null.
	DefaultInMemoryServiceName *string `mandatory:"false" json:"defaultInMemoryServiceName"`

	// The lost write protection setting for the tablespace.
	LostWriteProtect TablespaceSummaryLostWriteProtectEnum `mandatory:"false" json:"lostWriteProtect,omitempty"`

	// Indicates whether this is a chunk tablespace.
	IsChunkTablespace *bool `mandatory:"false" json:"isChunkTablespace"`

	// The temporary tablespace group.
	TempGroup *string `mandatory:"false" json:"tempGroup"`

	// The maximum tablespace size in KB. If the tablespace contains any data files with Autoextend enabled, then this column displays the amount of underlying free storage space for the tablespace. For example, if the current tablespace size is 1 GB, the combined maximum size of all its data files is 32 GB, and its underlying storage (for example, ASM or file system storage) has 20 GB of free space, then this column will have a value of approximately 20 GB. If the tablespace contains only data files with autoextend disabled, then this column displays the allocated space for the entire tablespace, that is, the combined size of all data files in the tablespace.
	MaxSizeKB *float32 `mandatory:"false" json:"maxSizeKB"`

	// The allocated tablespace size in KB.
	AllocatedSizeKB *float32 `mandatory:"false" json:"allocatedSizeKB"`

	// The size of the tablespace available for user data in KB. The difference between tablespace size and user data size is used for storing metadata.
	UserSizeKB *float32 `mandatory:"false" json:"userSizeKB"`

	// The free space available in the tablespace in KB.
	FreeSpaceKB *float32 `mandatory:"false" json:"freeSpaceKB"`

	// The total space used by the tablespace in KB.
	UsedSpaceKB *float32 `mandatory:"false" json:"usedSpaceKB"`

	// The percentage of used space out of the maximum available space in the tablespace.
	UsedPercentAvailable *float64 `mandatory:"false" json:"usedPercentAvailable"`

	// The percentage of used space out of the total allocated space in the tablespace.
	UsedPercentAllocated *float64 `mandatory:"false" json:"usedPercentAllocated"`

	// A list of the data files associated with the tablespace.
	Datafiles []Datafile `mandatory:"false" json:"datafiles"`
}

func (m TablespaceSummary) String() string {
	return common.PointerString(m)
}

// TablespaceSummaryTypeEnum Enum with underlying type: string
type TablespaceSummaryTypeEnum string

// Set of constants representing the allowable values for TablespaceSummaryTypeEnum
const (
	TablespaceSummaryTypeUndo                TablespaceSummaryTypeEnum = "UNDO"
	TablespaceSummaryTypeLostWriteProtection TablespaceSummaryTypeEnum = "LOST_WRITE_PROTECTION"
	TablespaceSummaryTypePermanent           TablespaceSummaryTypeEnum = "PERMANENT"
	TablespaceSummaryTypeTemporary           TablespaceSummaryTypeEnum = "TEMPORARY"
)

var mappingTablespaceSummaryType = map[string]TablespaceSummaryTypeEnum{
	"UNDO":                  TablespaceSummaryTypeUndo,
	"LOST_WRITE_PROTECTION": TablespaceSummaryTypeLostWriteProtection,
	"PERMANENT":             TablespaceSummaryTypePermanent,
	"TEMPORARY":             TablespaceSummaryTypeTemporary,
}

// GetTablespaceSummaryTypeEnumValues Enumerates the set of values for TablespaceSummaryTypeEnum
func GetTablespaceSummaryTypeEnumValues() []TablespaceSummaryTypeEnum {
	values := make([]TablespaceSummaryTypeEnum, 0)
	for _, v := range mappingTablespaceSummaryType {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryStatusEnum Enum with underlying type: string
type TablespaceSummaryStatusEnum string

// Set of constants representing the allowable values for TablespaceSummaryStatusEnum
const (
	TablespaceSummaryStatusOnline   TablespaceSummaryStatusEnum = "ONLINE"
	TablespaceSummaryStatusOffline  TablespaceSummaryStatusEnum = "OFFLINE"
	TablespaceSummaryStatusReadOnly TablespaceSummaryStatusEnum = "READ_ONLY"
)

var mappingTablespaceSummaryStatus = map[string]TablespaceSummaryStatusEnum{
	"ONLINE":    TablespaceSummaryStatusOnline,
	"OFFLINE":   TablespaceSummaryStatusOffline,
	"READ_ONLY": TablespaceSummaryStatusReadOnly,
}

// GetTablespaceSummaryStatusEnumValues Enumerates the set of values for TablespaceSummaryStatusEnum
func GetTablespaceSummaryStatusEnumValues() []TablespaceSummaryStatusEnum {
	values := make([]TablespaceSummaryStatusEnum, 0)
	for _, v := range mappingTablespaceSummaryStatus {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryLoggingEnum Enum with underlying type: string
type TablespaceSummaryLoggingEnum string

// Set of constants representing the allowable values for TablespaceSummaryLoggingEnum
const (
	TablespaceSummaryLoggingLogging   TablespaceSummaryLoggingEnum = "LOGGING"
	TablespaceSummaryLoggingNologging TablespaceSummaryLoggingEnum = "NOLOGGING"
)

var mappingTablespaceSummaryLogging = map[string]TablespaceSummaryLoggingEnum{
	"LOGGING":   TablespaceSummaryLoggingLogging,
	"NOLOGGING": TablespaceSummaryLoggingNologging,
}

// GetTablespaceSummaryLoggingEnumValues Enumerates the set of values for TablespaceSummaryLoggingEnum
func GetTablespaceSummaryLoggingEnumValues() []TablespaceSummaryLoggingEnum {
	values := make([]TablespaceSummaryLoggingEnum, 0)
	for _, v := range mappingTablespaceSummaryLogging {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryExtentManagementEnum Enum with underlying type: string
type TablespaceSummaryExtentManagementEnum string

// Set of constants representing the allowable values for TablespaceSummaryExtentManagementEnum
const (
	TablespaceSummaryExtentManagementLocal      TablespaceSummaryExtentManagementEnum = "LOCAL"
	TablespaceSummaryExtentManagementDictionary TablespaceSummaryExtentManagementEnum = "DICTIONARY"
)

var mappingTablespaceSummaryExtentManagement = map[string]TablespaceSummaryExtentManagementEnum{
	"LOCAL":      TablespaceSummaryExtentManagementLocal,
	"DICTIONARY": TablespaceSummaryExtentManagementDictionary,
}

// GetTablespaceSummaryExtentManagementEnumValues Enumerates the set of values for TablespaceSummaryExtentManagementEnum
func GetTablespaceSummaryExtentManagementEnumValues() []TablespaceSummaryExtentManagementEnum {
	values := make([]TablespaceSummaryExtentManagementEnum, 0)
	for _, v := range mappingTablespaceSummaryExtentManagement {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryAllocationTypeEnum Enum with underlying type: string
type TablespaceSummaryAllocationTypeEnum string

// Set of constants representing the allowable values for TablespaceSummaryAllocationTypeEnum
const (
	TablespaceSummaryAllocationTypeSystem  TablespaceSummaryAllocationTypeEnum = "SYSTEM"
	TablespaceSummaryAllocationTypeUniform TablespaceSummaryAllocationTypeEnum = "UNIFORM"
	TablespaceSummaryAllocationTypeUser    TablespaceSummaryAllocationTypeEnum = "USER"
)

var mappingTablespaceSummaryAllocationType = map[string]TablespaceSummaryAllocationTypeEnum{
	"SYSTEM":  TablespaceSummaryAllocationTypeSystem,
	"UNIFORM": TablespaceSummaryAllocationTypeUniform,
	"USER":    TablespaceSummaryAllocationTypeUser,
}

// GetTablespaceSummaryAllocationTypeEnumValues Enumerates the set of values for TablespaceSummaryAllocationTypeEnum
func GetTablespaceSummaryAllocationTypeEnumValues() []TablespaceSummaryAllocationTypeEnum {
	values := make([]TablespaceSummaryAllocationTypeEnum, 0)
	for _, v := range mappingTablespaceSummaryAllocationType {
		values = append(values, v)
	}
	return values
}

// TablespaceSummarySegmentSpaceManagementEnum Enum with underlying type: string
type TablespaceSummarySegmentSpaceManagementEnum string

// Set of constants representing the allowable values for TablespaceSummarySegmentSpaceManagementEnum
const (
	TablespaceSummarySegmentSpaceManagementManual TablespaceSummarySegmentSpaceManagementEnum = "MANUAL"
	TablespaceSummarySegmentSpaceManagementAuto   TablespaceSummarySegmentSpaceManagementEnum = "AUTO"
)

var mappingTablespaceSummarySegmentSpaceManagement = map[string]TablespaceSummarySegmentSpaceManagementEnum{
	"MANUAL": TablespaceSummarySegmentSpaceManagementManual,
	"AUTO":   TablespaceSummarySegmentSpaceManagementAuto,
}

// GetTablespaceSummarySegmentSpaceManagementEnumValues Enumerates the set of values for TablespaceSummarySegmentSpaceManagementEnum
func GetTablespaceSummarySegmentSpaceManagementEnumValues() []TablespaceSummarySegmentSpaceManagementEnum {
	values := make([]TablespaceSummarySegmentSpaceManagementEnum, 0)
	for _, v := range mappingTablespaceSummarySegmentSpaceManagement {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultTableCompressionEnum Enum with underlying type: string
type TablespaceSummaryDefaultTableCompressionEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultTableCompressionEnum
const (
	TablespaceSummaryDefaultTableCompressionEnabled  TablespaceSummaryDefaultTableCompressionEnum = "ENABLED"
	TablespaceSummaryDefaultTableCompressionDisabled TablespaceSummaryDefaultTableCompressionEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultTableCompression = map[string]TablespaceSummaryDefaultTableCompressionEnum{
	"ENABLED":  TablespaceSummaryDefaultTableCompressionEnabled,
	"DISABLED": TablespaceSummaryDefaultTableCompressionDisabled,
}

// GetTablespaceSummaryDefaultTableCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultTableCompressionEnum
func GetTablespaceSummaryDefaultTableCompressionEnumValues() []TablespaceSummaryDefaultTableCompressionEnum {
	values := make([]TablespaceSummaryDefaultTableCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultTableCompression {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryRetentionEnum Enum with underlying type: string
type TablespaceSummaryRetentionEnum string

// Set of constants representing the allowable values for TablespaceSummaryRetentionEnum
const (
	TablespaceSummaryRetentionGuarantee   TablespaceSummaryRetentionEnum = "GUARANTEE"
	TablespaceSummaryRetentionNoguarantee TablespaceSummaryRetentionEnum = "NOGUARANTEE"
	TablespaceSummaryRetentionNotApply    TablespaceSummaryRetentionEnum = "NOT_APPLY"
)

var mappingTablespaceSummaryRetention = map[string]TablespaceSummaryRetentionEnum{
	"GUARANTEE":   TablespaceSummaryRetentionGuarantee,
	"NOGUARANTEE": TablespaceSummaryRetentionNoguarantee,
	"NOT_APPLY":   TablespaceSummaryRetentionNotApply,
}

// GetTablespaceSummaryRetentionEnumValues Enumerates the set of values for TablespaceSummaryRetentionEnum
func GetTablespaceSummaryRetentionEnumValues() []TablespaceSummaryRetentionEnum {
	values := make([]TablespaceSummaryRetentionEnum, 0)
	for _, v := range mappingTablespaceSummaryRetention {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryPredicateEvaluationEnum Enum with underlying type: string
type TablespaceSummaryPredicateEvaluationEnum string

// Set of constants representing the allowable values for TablespaceSummaryPredicateEvaluationEnum
const (
	TablespaceSummaryPredicateEvaluationHost    TablespaceSummaryPredicateEvaluationEnum = "HOST"
	TablespaceSummaryPredicateEvaluationStorage TablespaceSummaryPredicateEvaluationEnum = "STORAGE"
)

var mappingTablespaceSummaryPredicateEvaluation = map[string]TablespaceSummaryPredicateEvaluationEnum{
	"HOST":    TablespaceSummaryPredicateEvaluationHost,
	"STORAGE": TablespaceSummaryPredicateEvaluationStorage,
}

// GetTablespaceSummaryPredicateEvaluationEnumValues Enumerates the set of values for TablespaceSummaryPredicateEvaluationEnum
func GetTablespaceSummaryPredicateEvaluationEnumValues() []TablespaceSummaryPredicateEvaluationEnum {
	values := make([]TablespaceSummaryPredicateEvaluationEnum, 0)
	for _, v := range mappingTablespaceSummaryPredicateEvaluation {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryCompressForEnum Enum with underlying type: string
type TablespaceSummaryCompressForEnum string

// Set of constants representing the allowable values for TablespaceSummaryCompressForEnum
const (
	TablespaceSummaryCompressForBasic            TablespaceSummaryCompressForEnum = "BASIC"
	TablespaceSummaryCompressForAdvanced         TablespaceSummaryCompressForEnum = "ADVANCED"
	TablespaceSummaryCompressForQueryLow         TablespaceSummaryCompressForEnum = "QUERY_LOW"
	TablespaceSummaryCompressForQueryHigh        TablespaceSummaryCompressForEnum = "QUERY_HIGH"
	TablespaceSummaryCompressForArchiveLow       TablespaceSummaryCompressForEnum = "ARCHIVE_LOW"
	TablespaceSummaryCompressForArchiveHigh      TablespaceSummaryCompressForEnum = "ARCHIVE_HIGH"
	TablespaceSummaryCompressForDirectLoadOnly   TablespaceSummaryCompressForEnum = "DIRECT_LOAD_ONLY"
	TablespaceSummaryCompressForForAllOperations TablespaceSummaryCompressForEnum = "FOR_ALL_OPERATIONS"
)

var mappingTablespaceSummaryCompressFor = map[string]TablespaceSummaryCompressForEnum{
	"BASIC":              TablespaceSummaryCompressForBasic,
	"ADVANCED":           TablespaceSummaryCompressForAdvanced,
	"QUERY_LOW":          TablespaceSummaryCompressForQueryLow,
	"QUERY_HIGH":         TablespaceSummaryCompressForQueryHigh,
	"ARCHIVE_LOW":        TablespaceSummaryCompressForArchiveLow,
	"ARCHIVE_HIGH":       TablespaceSummaryCompressForArchiveHigh,
	"DIRECT_LOAD_ONLY":   TablespaceSummaryCompressForDirectLoadOnly,
	"FOR_ALL_OPERATIONS": TablespaceSummaryCompressForForAllOperations,
}

// GetTablespaceSummaryCompressForEnumValues Enumerates the set of values for TablespaceSummaryCompressForEnum
func GetTablespaceSummaryCompressForEnumValues() []TablespaceSummaryCompressForEnum {
	values := make([]TablespaceSummaryCompressForEnum, 0)
	for _, v := range mappingTablespaceSummaryCompressFor {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryEnum
const (
	TablespaceSummaryDefaultInMemoryEnabled  TablespaceSummaryDefaultInMemoryEnum = "ENABLED"
	TablespaceSummaryDefaultInMemoryDisabled TablespaceSummaryDefaultInMemoryEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultInMemory = map[string]TablespaceSummaryDefaultInMemoryEnum{
	"ENABLED":  TablespaceSummaryDefaultInMemoryEnabled,
	"DISABLED": TablespaceSummaryDefaultInMemoryDisabled,
}

// GetTablespaceSummaryDefaultInMemoryEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryEnum
func GetTablespaceSummaryDefaultInMemoryEnumValues() []TablespaceSummaryDefaultInMemoryEnum {
	values := make([]TablespaceSummaryDefaultInMemoryEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemory {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryPriorityEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryPriorityEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryPriorityEnum
const (
	TablespaceSummaryDefaultInMemoryPriorityLow      TablespaceSummaryDefaultInMemoryPriorityEnum = "LOW"
	TablespaceSummaryDefaultInMemoryPriorityMedium   TablespaceSummaryDefaultInMemoryPriorityEnum = "MEDIUM"
	TablespaceSummaryDefaultInMemoryPriorityHigh     TablespaceSummaryDefaultInMemoryPriorityEnum = "HIGH"
	TablespaceSummaryDefaultInMemoryPriorityCritical TablespaceSummaryDefaultInMemoryPriorityEnum = "CRITICAL"
	TablespaceSummaryDefaultInMemoryPriorityNone     TablespaceSummaryDefaultInMemoryPriorityEnum = "NONE"
)

var mappingTablespaceSummaryDefaultInMemoryPriority = map[string]TablespaceSummaryDefaultInMemoryPriorityEnum{
	"LOW":      TablespaceSummaryDefaultInMemoryPriorityLow,
	"MEDIUM":   TablespaceSummaryDefaultInMemoryPriorityMedium,
	"HIGH":     TablespaceSummaryDefaultInMemoryPriorityHigh,
	"CRITICAL": TablespaceSummaryDefaultInMemoryPriorityCritical,
	"NONE":     TablespaceSummaryDefaultInMemoryPriorityNone,
}

// GetTablespaceSummaryDefaultInMemoryPriorityEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryPriorityEnum
func GetTablespaceSummaryDefaultInMemoryPriorityEnumValues() []TablespaceSummaryDefaultInMemoryPriorityEnum {
	values := make([]TablespaceSummaryDefaultInMemoryPriorityEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryPriority {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryDistributeEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryDistributeEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryDistributeEnum
const (
	TablespaceSummaryDefaultInMemoryDistributeAuto           TablespaceSummaryDefaultInMemoryDistributeEnum = "AUTO"
	TablespaceSummaryDefaultInMemoryDistributeByRowidRange   TablespaceSummaryDefaultInMemoryDistributeEnum = "BY_ROWID_RANGE"
	TablespaceSummaryDefaultInMemoryDistributeByPartition    TablespaceSummaryDefaultInMemoryDistributeEnum = "BY_PARTITION"
	TablespaceSummaryDefaultInMemoryDistributeBySubpartition TablespaceSummaryDefaultInMemoryDistributeEnum = "BY_SUBPARTITION"
)

var mappingTablespaceSummaryDefaultInMemoryDistribute = map[string]TablespaceSummaryDefaultInMemoryDistributeEnum{
	"AUTO":            TablespaceSummaryDefaultInMemoryDistributeAuto,
	"BY_ROWID_RANGE":  TablespaceSummaryDefaultInMemoryDistributeByRowidRange,
	"BY_PARTITION":    TablespaceSummaryDefaultInMemoryDistributeByPartition,
	"BY_SUBPARTITION": TablespaceSummaryDefaultInMemoryDistributeBySubpartition,
}

// GetTablespaceSummaryDefaultInMemoryDistributeEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryDistributeEnum
func GetTablespaceSummaryDefaultInMemoryDistributeEnumValues() []TablespaceSummaryDefaultInMemoryDistributeEnum {
	values := make([]TablespaceSummaryDefaultInMemoryDistributeEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryDistribute {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryCompressionEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryCompressionEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryCompressionEnum
const (
	TablespaceSummaryDefaultInMemoryCompressionNoMemcompress   TablespaceSummaryDefaultInMemoryCompressionEnum = "NO_MEMCOMPRESS"
	TablespaceSummaryDefaultInMemoryCompressionForDml          TablespaceSummaryDefaultInMemoryCompressionEnum = "FOR_DML"
	TablespaceSummaryDefaultInMemoryCompressionForQueryLow     TablespaceSummaryDefaultInMemoryCompressionEnum = "FOR_QUERY_LOW"
	TablespaceSummaryDefaultInMemoryCompressionForQueryHigh    TablespaceSummaryDefaultInMemoryCompressionEnum = "FOR_QUERY_HIGH"
	TablespaceSummaryDefaultInMemoryCompressionForCapacityLow  TablespaceSummaryDefaultInMemoryCompressionEnum = "FOR_CAPACITY_LOW"
	TablespaceSummaryDefaultInMemoryCompressionForCapacityHigh TablespaceSummaryDefaultInMemoryCompressionEnum = "FOR_CAPACITY_HIGH"
)

var mappingTablespaceSummaryDefaultInMemoryCompression = map[string]TablespaceSummaryDefaultInMemoryCompressionEnum{
	"NO_MEMCOMPRESS":    TablespaceSummaryDefaultInMemoryCompressionNoMemcompress,
	"FOR_DML":           TablespaceSummaryDefaultInMemoryCompressionForDml,
	"FOR_QUERY_LOW":     TablespaceSummaryDefaultInMemoryCompressionForQueryLow,
	"FOR_QUERY_HIGH":    TablespaceSummaryDefaultInMemoryCompressionForQueryHigh,
	"FOR_CAPACITY_LOW":  TablespaceSummaryDefaultInMemoryCompressionForCapacityLow,
	"FOR_CAPACITY_HIGH": TablespaceSummaryDefaultInMemoryCompressionForCapacityHigh,
}

// GetTablespaceSummaryDefaultInMemoryCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryCompressionEnum
func GetTablespaceSummaryDefaultInMemoryCompressionEnumValues() []TablespaceSummaryDefaultInMemoryCompressionEnum {
	values := make([]TablespaceSummaryDefaultInMemoryCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryCompression {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryDuplicateEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryDuplicateEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryDuplicateEnum
const (
	TablespaceSummaryDefaultInMemoryDuplicateNoDuplicate  TablespaceSummaryDefaultInMemoryDuplicateEnum = "NO_DUPLICATE"
	TablespaceSummaryDefaultInMemoryDuplicateDuplicate    TablespaceSummaryDefaultInMemoryDuplicateEnum = "DUPLICATE"
	TablespaceSummaryDefaultInMemoryDuplicateDuplicateAll TablespaceSummaryDefaultInMemoryDuplicateEnum = "DUPLICATE_ALL"
)

var mappingTablespaceSummaryDefaultInMemoryDuplicate = map[string]TablespaceSummaryDefaultInMemoryDuplicateEnum{
	"NO_DUPLICATE":  TablespaceSummaryDefaultInMemoryDuplicateNoDuplicate,
	"DUPLICATE":     TablespaceSummaryDefaultInMemoryDuplicateDuplicate,
	"DUPLICATE_ALL": TablespaceSummaryDefaultInMemoryDuplicateDuplicateAll,
}

// GetTablespaceSummaryDefaultInMemoryDuplicateEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryDuplicateEnum
func GetTablespaceSummaryDefaultInMemoryDuplicateEnumValues() []TablespaceSummaryDefaultInMemoryDuplicateEnum {
	values := make([]TablespaceSummaryDefaultInMemoryDuplicateEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryDuplicate {
		values = append(values, v)
	}
	return values
}

// TablespaceSummarySharedEnum Enum with underlying type: string
type TablespaceSummarySharedEnum string

// Set of constants representing the allowable values for TablespaceSummarySharedEnum
const (
	TablespaceSummarySharedShared      TablespaceSummarySharedEnum = "SHARED"
	TablespaceSummarySharedLocalOnLeaf TablespaceSummarySharedEnum = "LOCAL_ON_LEAF"
	TablespaceSummarySharedLocalOnAll  TablespaceSummarySharedEnum = "LOCAL_ON_ALL"
)

var mappingTablespaceSummaryShared = map[string]TablespaceSummarySharedEnum{
	"SHARED":        TablespaceSummarySharedShared,
	"LOCAL_ON_LEAF": TablespaceSummarySharedLocalOnLeaf,
	"LOCAL_ON_ALL":  TablespaceSummarySharedLocalOnAll,
}

// GetTablespaceSummarySharedEnumValues Enumerates the set of values for TablespaceSummarySharedEnum
func GetTablespaceSummarySharedEnumValues() []TablespaceSummarySharedEnum {
	values := make([]TablespaceSummarySharedEnum, 0)
	for _, v := range mappingTablespaceSummaryShared {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultIndexCompressionEnum Enum with underlying type: string
type TablespaceSummaryDefaultIndexCompressionEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultIndexCompressionEnum
const (
	TablespaceSummaryDefaultIndexCompressionEnabled  TablespaceSummaryDefaultIndexCompressionEnum = "ENABLED"
	TablespaceSummaryDefaultIndexCompressionDisabled TablespaceSummaryDefaultIndexCompressionEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultIndexCompression = map[string]TablespaceSummaryDefaultIndexCompressionEnum{
	"ENABLED":  TablespaceSummaryDefaultIndexCompressionEnabled,
	"DISABLED": TablespaceSummaryDefaultIndexCompressionDisabled,
}

// GetTablespaceSummaryDefaultIndexCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultIndexCompressionEnum
func GetTablespaceSummaryDefaultIndexCompressionEnumValues() []TablespaceSummaryDefaultIndexCompressionEnum {
	values := make([]TablespaceSummaryDefaultIndexCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultIndexCompression {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryIndexCompressForEnum Enum with underlying type: string
type TablespaceSummaryIndexCompressForEnum string

// Set of constants representing the allowable values for TablespaceSummaryIndexCompressForEnum
const (
	TablespaceSummaryIndexCompressForLow  TablespaceSummaryIndexCompressForEnum = "ADVANCED_LOW"
	TablespaceSummaryIndexCompressForHigh TablespaceSummaryIndexCompressForEnum = "ADVANCED_HIGH"
)

var mappingTablespaceSummaryIndexCompressFor = map[string]TablespaceSummaryIndexCompressForEnum{
	"ADVANCED_LOW":  TablespaceSummaryIndexCompressForLow,
	"ADVANCED_HIGH": TablespaceSummaryIndexCompressForHigh,
}

// GetTablespaceSummaryIndexCompressForEnumValues Enumerates the set of values for TablespaceSummaryIndexCompressForEnum
func GetTablespaceSummaryIndexCompressForEnumValues() []TablespaceSummaryIndexCompressForEnum {
	values := make([]TablespaceSummaryIndexCompressForEnum, 0)
	for _, v := range mappingTablespaceSummaryIndexCompressFor {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryDefaultInMemoryServiceEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryServiceEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryServiceEnum
const (
	TablespaceSummaryDefaultInMemoryServiceDefault     TablespaceSummaryDefaultInMemoryServiceEnum = "DEFAULT"
	TablespaceSummaryDefaultInMemoryServiceNone        TablespaceSummaryDefaultInMemoryServiceEnum = "NONE"
	TablespaceSummaryDefaultInMemoryServiceAll         TablespaceSummaryDefaultInMemoryServiceEnum = "ALL"
	TablespaceSummaryDefaultInMemoryServiceUserDefined TablespaceSummaryDefaultInMemoryServiceEnum = "USER_DEFINED"
)

var mappingTablespaceSummaryDefaultInMemoryService = map[string]TablespaceSummaryDefaultInMemoryServiceEnum{
	"DEFAULT":      TablespaceSummaryDefaultInMemoryServiceDefault,
	"NONE":         TablespaceSummaryDefaultInMemoryServiceNone,
	"ALL":          TablespaceSummaryDefaultInMemoryServiceAll,
	"USER_DEFINED": TablespaceSummaryDefaultInMemoryServiceUserDefined,
}

// GetTablespaceSummaryDefaultInMemoryServiceEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryServiceEnum
func GetTablespaceSummaryDefaultInMemoryServiceEnumValues() []TablespaceSummaryDefaultInMemoryServiceEnum {
	values := make([]TablespaceSummaryDefaultInMemoryServiceEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryService {
		values = append(values, v)
	}
	return values
}

// TablespaceSummaryLostWriteProtectEnum Enum with underlying type: string
type TablespaceSummaryLostWriteProtectEnum string

// Set of constants representing the allowable values for TablespaceSummaryLostWriteProtectEnum
const (
	TablespaceSummaryLostWriteProtectEnabled    TablespaceSummaryLostWriteProtectEnum = "ENABLED"
	TablespaceSummaryLostWriteProtectProtectOff TablespaceSummaryLostWriteProtectEnum = "PROTECT_OFF"
	TablespaceSummaryLostWriteProtectSuspend    TablespaceSummaryLostWriteProtectEnum = "SUSPEND"
)

var mappingTablespaceSummaryLostWriteProtect = map[string]TablespaceSummaryLostWriteProtectEnum{
	"ENABLED":     TablespaceSummaryLostWriteProtectEnabled,
	"PROTECT_OFF": TablespaceSummaryLostWriteProtectProtectOff,
	"SUSPEND":     TablespaceSummaryLostWriteProtectSuspend,
}

// GetTablespaceSummaryLostWriteProtectEnumValues Enumerates the set of values for TablespaceSummaryLostWriteProtectEnum
func GetTablespaceSummaryLostWriteProtectEnumValues() []TablespaceSummaryLostWriteProtectEnum {
	values := make([]TablespaceSummaryLostWriteProtectEnum, 0)
	for _, v := range mappingTablespaceSummaryLostWriteProtect {
		values = append(values, v)
	}
	return values
}
