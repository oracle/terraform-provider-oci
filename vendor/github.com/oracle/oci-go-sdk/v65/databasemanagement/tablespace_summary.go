// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Indicates whether this is the default tablespace.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// A list of the data files associated with the tablespace.
	Datafiles []Datafile `mandatory:"false" json:"datafiles"`
}

func (m TablespaceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TablespaceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTablespaceSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetTablespaceSummaryTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTablespaceSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTablespaceSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryLoggingEnum(string(m.Logging)); !ok && m.Logging != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Logging: %s. Supported values are: %s.", m.Logging, strings.Join(GetTablespaceSummaryLoggingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryExtentManagementEnum(string(m.ExtentManagement)); !ok && m.ExtentManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExtentManagement: %s. Supported values are: %s.", m.ExtentManagement, strings.Join(GetTablespaceSummaryExtentManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryAllocationTypeEnum(string(m.AllocationType)); !ok && m.AllocationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllocationType: %s. Supported values are: %s.", m.AllocationType, strings.Join(GetTablespaceSummaryAllocationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummarySegmentSpaceManagementEnum(string(m.SegmentSpaceManagement)); !ok && m.SegmentSpaceManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SegmentSpaceManagement: %s. Supported values are: %s.", m.SegmentSpaceManagement, strings.Join(GetTablespaceSummarySegmentSpaceManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultTableCompressionEnum(string(m.DefaultTableCompression)); !ok && m.DefaultTableCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultTableCompression: %s. Supported values are: %s.", m.DefaultTableCompression, strings.Join(GetTablespaceSummaryDefaultTableCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryRetentionEnum(string(m.Retention)); !ok && m.Retention != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Retention: %s. Supported values are: %s.", m.Retention, strings.Join(GetTablespaceSummaryRetentionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryPredicateEvaluationEnum(string(m.PredicateEvaluation)); !ok && m.PredicateEvaluation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PredicateEvaluation: %s. Supported values are: %s.", m.PredicateEvaluation, strings.Join(GetTablespaceSummaryPredicateEvaluationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryCompressForEnum(string(m.CompressFor)); !ok && m.CompressFor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompressFor: %s. Supported values are: %s.", m.CompressFor, strings.Join(GetTablespaceSummaryCompressForEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryEnum(string(m.DefaultInMemory)); !ok && m.DefaultInMemory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemory: %s. Supported values are: %s.", m.DefaultInMemory, strings.Join(GetTablespaceSummaryDefaultInMemoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryPriorityEnum(string(m.DefaultInMemoryPriority)); !ok && m.DefaultInMemoryPriority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryPriority: %s. Supported values are: %s.", m.DefaultInMemoryPriority, strings.Join(GetTablespaceSummaryDefaultInMemoryPriorityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryDistributeEnum(string(m.DefaultInMemoryDistribute)); !ok && m.DefaultInMemoryDistribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryDistribute: %s. Supported values are: %s.", m.DefaultInMemoryDistribute, strings.Join(GetTablespaceSummaryDefaultInMemoryDistributeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryCompressionEnum(string(m.DefaultInMemoryCompression)); !ok && m.DefaultInMemoryCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryCompression: %s. Supported values are: %s.", m.DefaultInMemoryCompression, strings.Join(GetTablespaceSummaryDefaultInMemoryCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryDuplicateEnum(string(m.DefaultInMemoryDuplicate)); !ok && m.DefaultInMemoryDuplicate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryDuplicate: %s. Supported values are: %s.", m.DefaultInMemoryDuplicate, strings.Join(GetTablespaceSummaryDefaultInMemoryDuplicateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummarySharedEnum(string(m.Shared)); !ok && m.Shared != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shared: %s. Supported values are: %s.", m.Shared, strings.Join(GetTablespaceSummarySharedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultIndexCompressionEnum(string(m.DefaultIndexCompression)); !ok && m.DefaultIndexCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultIndexCompression: %s. Supported values are: %s.", m.DefaultIndexCompression, strings.Join(GetTablespaceSummaryDefaultIndexCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryIndexCompressForEnum(string(m.IndexCompressFor)); !ok && m.IndexCompressFor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IndexCompressFor: %s. Supported values are: %s.", m.IndexCompressFor, strings.Join(GetTablespaceSummaryIndexCompressForEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryDefaultInMemoryServiceEnum(string(m.DefaultInMemoryService)); !ok && m.DefaultInMemoryService != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryService: %s. Supported values are: %s.", m.DefaultInMemoryService, strings.Join(GetTablespaceSummaryDefaultInMemoryServiceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSummaryLostWriteProtectEnum(string(m.LostWriteProtect)); !ok && m.LostWriteProtect != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LostWriteProtect: %s. Supported values are: %s.", m.LostWriteProtect, strings.Join(GetTablespaceSummaryLostWriteProtectEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingTablespaceSummaryTypeEnum = map[string]TablespaceSummaryTypeEnum{
	"UNDO":                  TablespaceSummaryTypeUndo,
	"LOST_WRITE_PROTECTION": TablespaceSummaryTypeLostWriteProtection,
	"PERMANENT":             TablespaceSummaryTypePermanent,
	"TEMPORARY":             TablespaceSummaryTypeTemporary,
}

var mappingTablespaceSummaryTypeEnumLowerCase = map[string]TablespaceSummaryTypeEnum{
	"undo":                  TablespaceSummaryTypeUndo,
	"lost_write_protection": TablespaceSummaryTypeLostWriteProtection,
	"permanent":             TablespaceSummaryTypePermanent,
	"temporary":             TablespaceSummaryTypeTemporary,
}

// GetTablespaceSummaryTypeEnumValues Enumerates the set of values for TablespaceSummaryTypeEnum
func GetTablespaceSummaryTypeEnumValues() []TablespaceSummaryTypeEnum {
	values := make([]TablespaceSummaryTypeEnum, 0)
	for _, v := range mappingTablespaceSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryTypeEnumStringValues Enumerates the set of values in String for TablespaceSummaryTypeEnum
func GetTablespaceSummaryTypeEnumStringValues() []string {
	return []string{
		"UNDO",
		"LOST_WRITE_PROTECTION",
		"PERMANENT",
		"TEMPORARY",
	}
}

// GetMappingTablespaceSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryTypeEnum(val string) (TablespaceSummaryTypeEnum, bool) {
	enum, ok := mappingTablespaceSummaryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryStatusEnum Enum with underlying type: string
type TablespaceSummaryStatusEnum string

// Set of constants representing the allowable values for TablespaceSummaryStatusEnum
const (
	TablespaceSummaryStatusOnline   TablespaceSummaryStatusEnum = "ONLINE"
	TablespaceSummaryStatusOffline  TablespaceSummaryStatusEnum = "OFFLINE"
	TablespaceSummaryStatusReadOnly TablespaceSummaryStatusEnum = "READ_ONLY"
)

var mappingTablespaceSummaryStatusEnum = map[string]TablespaceSummaryStatusEnum{
	"ONLINE":    TablespaceSummaryStatusOnline,
	"OFFLINE":   TablespaceSummaryStatusOffline,
	"READ_ONLY": TablespaceSummaryStatusReadOnly,
}

var mappingTablespaceSummaryStatusEnumLowerCase = map[string]TablespaceSummaryStatusEnum{
	"online":    TablespaceSummaryStatusOnline,
	"offline":   TablespaceSummaryStatusOffline,
	"read_only": TablespaceSummaryStatusReadOnly,
}

// GetTablespaceSummaryStatusEnumValues Enumerates the set of values for TablespaceSummaryStatusEnum
func GetTablespaceSummaryStatusEnumValues() []TablespaceSummaryStatusEnum {
	values := make([]TablespaceSummaryStatusEnum, 0)
	for _, v := range mappingTablespaceSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryStatusEnumStringValues Enumerates the set of values in String for TablespaceSummaryStatusEnum
func GetTablespaceSummaryStatusEnumStringValues() []string {
	return []string{
		"ONLINE",
		"OFFLINE",
		"READ_ONLY",
	}
}

// GetMappingTablespaceSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryStatusEnum(val string) (TablespaceSummaryStatusEnum, bool) {
	enum, ok := mappingTablespaceSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryLoggingEnum Enum with underlying type: string
type TablespaceSummaryLoggingEnum string

// Set of constants representing the allowable values for TablespaceSummaryLoggingEnum
const (
	TablespaceSummaryLoggingLogging   TablespaceSummaryLoggingEnum = "LOGGING"
	TablespaceSummaryLoggingNologging TablespaceSummaryLoggingEnum = "NOLOGGING"
)

var mappingTablespaceSummaryLoggingEnum = map[string]TablespaceSummaryLoggingEnum{
	"LOGGING":   TablespaceSummaryLoggingLogging,
	"NOLOGGING": TablespaceSummaryLoggingNologging,
}

var mappingTablespaceSummaryLoggingEnumLowerCase = map[string]TablespaceSummaryLoggingEnum{
	"logging":   TablespaceSummaryLoggingLogging,
	"nologging": TablespaceSummaryLoggingNologging,
}

// GetTablespaceSummaryLoggingEnumValues Enumerates the set of values for TablespaceSummaryLoggingEnum
func GetTablespaceSummaryLoggingEnumValues() []TablespaceSummaryLoggingEnum {
	values := make([]TablespaceSummaryLoggingEnum, 0)
	for _, v := range mappingTablespaceSummaryLoggingEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryLoggingEnumStringValues Enumerates the set of values in String for TablespaceSummaryLoggingEnum
func GetTablespaceSummaryLoggingEnumStringValues() []string {
	return []string{
		"LOGGING",
		"NOLOGGING",
	}
}

// GetMappingTablespaceSummaryLoggingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryLoggingEnum(val string) (TablespaceSummaryLoggingEnum, bool) {
	enum, ok := mappingTablespaceSummaryLoggingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryExtentManagementEnum Enum with underlying type: string
type TablespaceSummaryExtentManagementEnum string

// Set of constants representing the allowable values for TablespaceSummaryExtentManagementEnum
const (
	TablespaceSummaryExtentManagementLocal      TablespaceSummaryExtentManagementEnum = "LOCAL"
	TablespaceSummaryExtentManagementDictionary TablespaceSummaryExtentManagementEnum = "DICTIONARY"
)

var mappingTablespaceSummaryExtentManagementEnum = map[string]TablespaceSummaryExtentManagementEnum{
	"LOCAL":      TablespaceSummaryExtentManagementLocal,
	"DICTIONARY": TablespaceSummaryExtentManagementDictionary,
}

var mappingTablespaceSummaryExtentManagementEnumLowerCase = map[string]TablespaceSummaryExtentManagementEnum{
	"local":      TablespaceSummaryExtentManagementLocal,
	"dictionary": TablespaceSummaryExtentManagementDictionary,
}

// GetTablespaceSummaryExtentManagementEnumValues Enumerates the set of values for TablespaceSummaryExtentManagementEnum
func GetTablespaceSummaryExtentManagementEnumValues() []TablespaceSummaryExtentManagementEnum {
	values := make([]TablespaceSummaryExtentManagementEnum, 0)
	for _, v := range mappingTablespaceSummaryExtentManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryExtentManagementEnumStringValues Enumerates the set of values in String for TablespaceSummaryExtentManagementEnum
func GetTablespaceSummaryExtentManagementEnumStringValues() []string {
	return []string{
		"LOCAL",
		"DICTIONARY",
	}
}

// GetMappingTablespaceSummaryExtentManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryExtentManagementEnum(val string) (TablespaceSummaryExtentManagementEnum, bool) {
	enum, ok := mappingTablespaceSummaryExtentManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryAllocationTypeEnum Enum with underlying type: string
type TablespaceSummaryAllocationTypeEnum string

// Set of constants representing the allowable values for TablespaceSummaryAllocationTypeEnum
const (
	TablespaceSummaryAllocationTypeSystem  TablespaceSummaryAllocationTypeEnum = "SYSTEM"
	TablespaceSummaryAllocationTypeUniform TablespaceSummaryAllocationTypeEnum = "UNIFORM"
	TablespaceSummaryAllocationTypeUser    TablespaceSummaryAllocationTypeEnum = "USER"
)

var mappingTablespaceSummaryAllocationTypeEnum = map[string]TablespaceSummaryAllocationTypeEnum{
	"SYSTEM":  TablespaceSummaryAllocationTypeSystem,
	"UNIFORM": TablespaceSummaryAllocationTypeUniform,
	"USER":    TablespaceSummaryAllocationTypeUser,
}

var mappingTablespaceSummaryAllocationTypeEnumLowerCase = map[string]TablespaceSummaryAllocationTypeEnum{
	"system":  TablespaceSummaryAllocationTypeSystem,
	"uniform": TablespaceSummaryAllocationTypeUniform,
	"user":    TablespaceSummaryAllocationTypeUser,
}

// GetTablespaceSummaryAllocationTypeEnumValues Enumerates the set of values for TablespaceSummaryAllocationTypeEnum
func GetTablespaceSummaryAllocationTypeEnumValues() []TablespaceSummaryAllocationTypeEnum {
	values := make([]TablespaceSummaryAllocationTypeEnum, 0)
	for _, v := range mappingTablespaceSummaryAllocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryAllocationTypeEnumStringValues Enumerates the set of values in String for TablespaceSummaryAllocationTypeEnum
func GetTablespaceSummaryAllocationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"UNIFORM",
		"USER",
	}
}

// GetMappingTablespaceSummaryAllocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryAllocationTypeEnum(val string) (TablespaceSummaryAllocationTypeEnum, bool) {
	enum, ok := mappingTablespaceSummaryAllocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummarySegmentSpaceManagementEnum Enum with underlying type: string
type TablespaceSummarySegmentSpaceManagementEnum string

// Set of constants representing the allowable values for TablespaceSummarySegmentSpaceManagementEnum
const (
	TablespaceSummarySegmentSpaceManagementManual TablespaceSummarySegmentSpaceManagementEnum = "MANUAL"
	TablespaceSummarySegmentSpaceManagementAuto   TablespaceSummarySegmentSpaceManagementEnum = "AUTO"
)

var mappingTablespaceSummarySegmentSpaceManagementEnum = map[string]TablespaceSummarySegmentSpaceManagementEnum{
	"MANUAL": TablespaceSummarySegmentSpaceManagementManual,
	"AUTO":   TablespaceSummarySegmentSpaceManagementAuto,
}

var mappingTablespaceSummarySegmentSpaceManagementEnumLowerCase = map[string]TablespaceSummarySegmentSpaceManagementEnum{
	"manual": TablespaceSummarySegmentSpaceManagementManual,
	"auto":   TablespaceSummarySegmentSpaceManagementAuto,
}

// GetTablespaceSummarySegmentSpaceManagementEnumValues Enumerates the set of values for TablespaceSummarySegmentSpaceManagementEnum
func GetTablespaceSummarySegmentSpaceManagementEnumValues() []TablespaceSummarySegmentSpaceManagementEnum {
	values := make([]TablespaceSummarySegmentSpaceManagementEnum, 0)
	for _, v := range mappingTablespaceSummarySegmentSpaceManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummarySegmentSpaceManagementEnumStringValues Enumerates the set of values in String for TablespaceSummarySegmentSpaceManagementEnum
func GetTablespaceSummarySegmentSpaceManagementEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTO",
	}
}

// GetMappingTablespaceSummarySegmentSpaceManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummarySegmentSpaceManagementEnum(val string) (TablespaceSummarySegmentSpaceManagementEnum, bool) {
	enum, ok := mappingTablespaceSummarySegmentSpaceManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryDefaultTableCompressionEnum Enum with underlying type: string
type TablespaceSummaryDefaultTableCompressionEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultTableCompressionEnum
const (
	TablespaceSummaryDefaultTableCompressionEnabled  TablespaceSummaryDefaultTableCompressionEnum = "ENABLED"
	TablespaceSummaryDefaultTableCompressionDisabled TablespaceSummaryDefaultTableCompressionEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultTableCompressionEnum = map[string]TablespaceSummaryDefaultTableCompressionEnum{
	"ENABLED":  TablespaceSummaryDefaultTableCompressionEnabled,
	"DISABLED": TablespaceSummaryDefaultTableCompressionDisabled,
}

var mappingTablespaceSummaryDefaultTableCompressionEnumLowerCase = map[string]TablespaceSummaryDefaultTableCompressionEnum{
	"enabled":  TablespaceSummaryDefaultTableCompressionEnabled,
	"disabled": TablespaceSummaryDefaultTableCompressionDisabled,
}

// GetTablespaceSummaryDefaultTableCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultTableCompressionEnum
func GetTablespaceSummaryDefaultTableCompressionEnumValues() []TablespaceSummaryDefaultTableCompressionEnum {
	values := make([]TablespaceSummaryDefaultTableCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultTableCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultTableCompressionEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultTableCompressionEnum
func GetTablespaceSummaryDefaultTableCompressionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceSummaryDefaultTableCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultTableCompressionEnum(val string) (TablespaceSummaryDefaultTableCompressionEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultTableCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryRetentionEnum Enum with underlying type: string
type TablespaceSummaryRetentionEnum string

// Set of constants representing the allowable values for TablespaceSummaryRetentionEnum
const (
	TablespaceSummaryRetentionGuarantee   TablespaceSummaryRetentionEnum = "GUARANTEE"
	TablespaceSummaryRetentionNoguarantee TablespaceSummaryRetentionEnum = "NOGUARANTEE"
	TablespaceSummaryRetentionNotApply    TablespaceSummaryRetentionEnum = "NOT_APPLY"
)

var mappingTablespaceSummaryRetentionEnum = map[string]TablespaceSummaryRetentionEnum{
	"GUARANTEE":   TablespaceSummaryRetentionGuarantee,
	"NOGUARANTEE": TablespaceSummaryRetentionNoguarantee,
	"NOT_APPLY":   TablespaceSummaryRetentionNotApply,
}

var mappingTablespaceSummaryRetentionEnumLowerCase = map[string]TablespaceSummaryRetentionEnum{
	"guarantee":   TablespaceSummaryRetentionGuarantee,
	"noguarantee": TablespaceSummaryRetentionNoguarantee,
	"not_apply":   TablespaceSummaryRetentionNotApply,
}

// GetTablespaceSummaryRetentionEnumValues Enumerates the set of values for TablespaceSummaryRetentionEnum
func GetTablespaceSummaryRetentionEnumValues() []TablespaceSummaryRetentionEnum {
	values := make([]TablespaceSummaryRetentionEnum, 0)
	for _, v := range mappingTablespaceSummaryRetentionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryRetentionEnumStringValues Enumerates the set of values in String for TablespaceSummaryRetentionEnum
func GetTablespaceSummaryRetentionEnumStringValues() []string {
	return []string{
		"GUARANTEE",
		"NOGUARANTEE",
		"NOT_APPLY",
	}
}

// GetMappingTablespaceSummaryRetentionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryRetentionEnum(val string) (TablespaceSummaryRetentionEnum, bool) {
	enum, ok := mappingTablespaceSummaryRetentionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryPredicateEvaluationEnum Enum with underlying type: string
type TablespaceSummaryPredicateEvaluationEnum string

// Set of constants representing the allowable values for TablespaceSummaryPredicateEvaluationEnum
const (
	TablespaceSummaryPredicateEvaluationHost    TablespaceSummaryPredicateEvaluationEnum = "HOST"
	TablespaceSummaryPredicateEvaluationStorage TablespaceSummaryPredicateEvaluationEnum = "STORAGE"
)

var mappingTablespaceSummaryPredicateEvaluationEnum = map[string]TablespaceSummaryPredicateEvaluationEnum{
	"HOST":    TablespaceSummaryPredicateEvaluationHost,
	"STORAGE": TablespaceSummaryPredicateEvaluationStorage,
}

var mappingTablespaceSummaryPredicateEvaluationEnumLowerCase = map[string]TablespaceSummaryPredicateEvaluationEnum{
	"host":    TablespaceSummaryPredicateEvaluationHost,
	"storage": TablespaceSummaryPredicateEvaluationStorage,
}

// GetTablespaceSummaryPredicateEvaluationEnumValues Enumerates the set of values for TablespaceSummaryPredicateEvaluationEnum
func GetTablespaceSummaryPredicateEvaluationEnumValues() []TablespaceSummaryPredicateEvaluationEnum {
	values := make([]TablespaceSummaryPredicateEvaluationEnum, 0)
	for _, v := range mappingTablespaceSummaryPredicateEvaluationEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryPredicateEvaluationEnumStringValues Enumerates the set of values in String for TablespaceSummaryPredicateEvaluationEnum
func GetTablespaceSummaryPredicateEvaluationEnumStringValues() []string {
	return []string{
		"HOST",
		"STORAGE",
	}
}

// GetMappingTablespaceSummaryPredicateEvaluationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryPredicateEvaluationEnum(val string) (TablespaceSummaryPredicateEvaluationEnum, bool) {
	enum, ok := mappingTablespaceSummaryPredicateEvaluationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTablespaceSummaryCompressForEnum = map[string]TablespaceSummaryCompressForEnum{
	"BASIC":              TablespaceSummaryCompressForBasic,
	"ADVANCED":           TablespaceSummaryCompressForAdvanced,
	"QUERY_LOW":          TablespaceSummaryCompressForQueryLow,
	"QUERY_HIGH":         TablespaceSummaryCompressForQueryHigh,
	"ARCHIVE_LOW":        TablespaceSummaryCompressForArchiveLow,
	"ARCHIVE_HIGH":       TablespaceSummaryCompressForArchiveHigh,
	"DIRECT_LOAD_ONLY":   TablespaceSummaryCompressForDirectLoadOnly,
	"FOR_ALL_OPERATIONS": TablespaceSummaryCompressForForAllOperations,
}

var mappingTablespaceSummaryCompressForEnumLowerCase = map[string]TablespaceSummaryCompressForEnum{
	"basic":              TablespaceSummaryCompressForBasic,
	"advanced":           TablespaceSummaryCompressForAdvanced,
	"query_low":          TablespaceSummaryCompressForQueryLow,
	"query_high":         TablespaceSummaryCompressForQueryHigh,
	"archive_low":        TablespaceSummaryCompressForArchiveLow,
	"archive_high":       TablespaceSummaryCompressForArchiveHigh,
	"direct_load_only":   TablespaceSummaryCompressForDirectLoadOnly,
	"for_all_operations": TablespaceSummaryCompressForForAllOperations,
}

// GetTablespaceSummaryCompressForEnumValues Enumerates the set of values for TablespaceSummaryCompressForEnum
func GetTablespaceSummaryCompressForEnumValues() []TablespaceSummaryCompressForEnum {
	values := make([]TablespaceSummaryCompressForEnum, 0)
	for _, v := range mappingTablespaceSummaryCompressForEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryCompressForEnumStringValues Enumerates the set of values in String for TablespaceSummaryCompressForEnum
func GetTablespaceSummaryCompressForEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
		"QUERY_LOW",
		"QUERY_HIGH",
		"ARCHIVE_LOW",
		"ARCHIVE_HIGH",
		"DIRECT_LOAD_ONLY",
		"FOR_ALL_OPERATIONS",
	}
}

// GetMappingTablespaceSummaryCompressForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryCompressForEnum(val string) (TablespaceSummaryCompressForEnum, bool) {
	enum, ok := mappingTablespaceSummaryCompressForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryDefaultInMemoryEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryEnum
const (
	TablespaceSummaryDefaultInMemoryEnabled  TablespaceSummaryDefaultInMemoryEnum = "ENABLED"
	TablespaceSummaryDefaultInMemoryDisabled TablespaceSummaryDefaultInMemoryEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultInMemoryEnum = map[string]TablespaceSummaryDefaultInMemoryEnum{
	"ENABLED":  TablespaceSummaryDefaultInMemoryEnabled,
	"DISABLED": TablespaceSummaryDefaultInMemoryDisabled,
}

var mappingTablespaceSummaryDefaultInMemoryEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryEnum{
	"enabled":  TablespaceSummaryDefaultInMemoryEnabled,
	"disabled": TablespaceSummaryDefaultInMemoryDisabled,
}

// GetTablespaceSummaryDefaultInMemoryEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryEnum
func GetTablespaceSummaryDefaultInMemoryEnumValues() []TablespaceSummaryDefaultInMemoryEnum {
	values := make([]TablespaceSummaryDefaultInMemoryEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryEnum
func GetTablespaceSummaryDefaultInMemoryEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryEnum(val string) (TablespaceSummaryDefaultInMemoryEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTablespaceSummaryDefaultInMemoryPriorityEnum = map[string]TablespaceSummaryDefaultInMemoryPriorityEnum{
	"LOW":      TablespaceSummaryDefaultInMemoryPriorityLow,
	"MEDIUM":   TablespaceSummaryDefaultInMemoryPriorityMedium,
	"HIGH":     TablespaceSummaryDefaultInMemoryPriorityHigh,
	"CRITICAL": TablespaceSummaryDefaultInMemoryPriorityCritical,
	"NONE":     TablespaceSummaryDefaultInMemoryPriorityNone,
}

var mappingTablespaceSummaryDefaultInMemoryPriorityEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryPriorityEnum{
	"low":      TablespaceSummaryDefaultInMemoryPriorityLow,
	"medium":   TablespaceSummaryDefaultInMemoryPriorityMedium,
	"high":     TablespaceSummaryDefaultInMemoryPriorityHigh,
	"critical": TablespaceSummaryDefaultInMemoryPriorityCritical,
	"none":     TablespaceSummaryDefaultInMemoryPriorityNone,
}

// GetTablespaceSummaryDefaultInMemoryPriorityEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryPriorityEnum
func GetTablespaceSummaryDefaultInMemoryPriorityEnumValues() []TablespaceSummaryDefaultInMemoryPriorityEnum {
	values := make([]TablespaceSummaryDefaultInMemoryPriorityEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryPriorityEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryPriorityEnum
func GetTablespaceSummaryDefaultInMemoryPriorityEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
		"NONE",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryPriorityEnum(val string) (TablespaceSummaryDefaultInMemoryPriorityEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryPriorityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTablespaceSummaryDefaultInMemoryDistributeEnum = map[string]TablespaceSummaryDefaultInMemoryDistributeEnum{
	"AUTO":            TablespaceSummaryDefaultInMemoryDistributeAuto,
	"BY_ROWID_RANGE":  TablespaceSummaryDefaultInMemoryDistributeByRowidRange,
	"BY_PARTITION":    TablespaceSummaryDefaultInMemoryDistributeByPartition,
	"BY_SUBPARTITION": TablespaceSummaryDefaultInMemoryDistributeBySubpartition,
}

var mappingTablespaceSummaryDefaultInMemoryDistributeEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryDistributeEnum{
	"auto":            TablespaceSummaryDefaultInMemoryDistributeAuto,
	"by_rowid_range":  TablespaceSummaryDefaultInMemoryDistributeByRowidRange,
	"by_partition":    TablespaceSummaryDefaultInMemoryDistributeByPartition,
	"by_subpartition": TablespaceSummaryDefaultInMemoryDistributeBySubpartition,
}

// GetTablespaceSummaryDefaultInMemoryDistributeEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryDistributeEnum
func GetTablespaceSummaryDefaultInMemoryDistributeEnumValues() []TablespaceSummaryDefaultInMemoryDistributeEnum {
	values := make([]TablespaceSummaryDefaultInMemoryDistributeEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryDistributeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryDistributeEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryDistributeEnum
func GetTablespaceSummaryDefaultInMemoryDistributeEnumStringValues() []string {
	return []string{
		"AUTO",
		"BY_ROWID_RANGE",
		"BY_PARTITION",
		"BY_SUBPARTITION",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryDistributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryDistributeEnum(val string) (TablespaceSummaryDefaultInMemoryDistributeEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryDistributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTablespaceSummaryDefaultInMemoryCompressionEnum = map[string]TablespaceSummaryDefaultInMemoryCompressionEnum{
	"NO_MEMCOMPRESS":    TablespaceSummaryDefaultInMemoryCompressionNoMemcompress,
	"FOR_DML":           TablespaceSummaryDefaultInMemoryCompressionForDml,
	"FOR_QUERY_LOW":     TablespaceSummaryDefaultInMemoryCompressionForQueryLow,
	"FOR_QUERY_HIGH":    TablespaceSummaryDefaultInMemoryCompressionForQueryHigh,
	"FOR_CAPACITY_LOW":  TablespaceSummaryDefaultInMemoryCompressionForCapacityLow,
	"FOR_CAPACITY_HIGH": TablespaceSummaryDefaultInMemoryCompressionForCapacityHigh,
}

var mappingTablespaceSummaryDefaultInMemoryCompressionEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryCompressionEnum{
	"no_memcompress":    TablespaceSummaryDefaultInMemoryCompressionNoMemcompress,
	"for_dml":           TablespaceSummaryDefaultInMemoryCompressionForDml,
	"for_query_low":     TablespaceSummaryDefaultInMemoryCompressionForQueryLow,
	"for_query_high":    TablespaceSummaryDefaultInMemoryCompressionForQueryHigh,
	"for_capacity_low":  TablespaceSummaryDefaultInMemoryCompressionForCapacityLow,
	"for_capacity_high": TablespaceSummaryDefaultInMemoryCompressionForCapacityHigh,
}

// GetTablespaceSummaryDefaultInMemoryCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryCompressionEnum
func GetTablespaceSummaryDefaultInMemoryCompressionEnumValues() []TablespaceSummaryDefaultInMemoryCompressionEnum {
	values := make([]TablespaceSummaryDefaultInMemoryCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryCompressionEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryCompressionEnum
func GetTablespaceSummaryDefaultInMemoryCompressionEnumStringValues() []string {
	return []string{
		"NO_MEMCOMPRESS",
		"FOR_DML",
		"FOR_QUERY_LOW",
		"FOR_QUERY_HIGH",
		"FOR_CAPACITY_LOW",
		"FOR_CAPACITY_HIGH",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryCompressionEnum(val string) (TablespaceSummaryDefaultInMemoryCompressionEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryDefaultInMemoryDuplicateEnum Enum with underlying type: string
type TablespaceSummaryDefaultInMemoryDuplicateEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultInMemoryDuplicateEnum
const (
	TablespaceSummaryDefaultInMemoryDuplicateNoDuplicate  TablespaceSummaryDefaultInMemoryDuplicateEnum = "NO_DUPLICATE"
	TablespaceSummaryDefaultInMemoryDuplicateDuplicate    TablespaceSummaryDefaultInMemoryDuplicateEnum = "DUPLICATE"
	TablespaceSummaryDefaultInMemoryDuplicateDuplicateAll TablespaceSummaryDefaultInMemoryDuplicateEnum = "DUPLICATE_ALL"
)

var mappingTablespaceSummaryDefaultInMemoryDuplicateEnum = map[string]TablespaceSummaryDefaultInMemoryDuplicateEnum{
	"NO_DUPLICATE":  TablespaceSummaryDefaultInMemoryDuplicateNoDuplicate,
	"DUPLICATE":     TablespaceSummaryDefaultInMemoryDuplicateDuplicate,
	"DUPLICATE_ALL": TablespaceSummaryDefaultInMemoryDuplicateDuplicateAll,
}

var mappingTablespaceSummaryDefaultInMemoryDuplicateEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryDuplicateEnum{
	"no_duplicate":  TablespaceSummaryDefaultInMemoryDuplicateNoDuplicate,
	"duplicate":     TablespaceSummaryDefaultInMemoryDuplicateDuplicate,
	"duplicate_all": TablespaceSummaryDefaultInMemoryDuplicateDuplicateAll,
}

// GetTablespaceSummaryDefaultInMemoryDuplicateEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryDuplicateEnum
func GetTablespaceSummaryDefaultInMemoryDuplicateEnumValues() []TablespaceSummaryDefaultInMemoryDuplicateEnum {
	values := make([]TablespaceSummaryDefaultInMemoryDuplicateEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryDuplicateEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryDuplicateEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryDuplicateEnum
func GetTablespaceSummaryDefaultInMemoryDuplicateEnumStringValues() []string {
	return []string{
		"NO_DUPLICATE",
		"DUPLICATE",
		"DUPLICATE_ALL",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryDuplicateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryDuplicateEnum(val string) (TablespaceSummaryDefaultInMemoryDuplicateEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryDuplicateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummarySharedEnum Enum with underlying type: string
type TablespaceSummarySharedEnum string

// Set of constants representing the allowable values for TablespaceSummarySharedEnum
const (
	TablespaceSummarySharedShared      TablespaceSummarySharedEnum = "SHARED"
	TablespaceSummarySharedLocalOnLeaf TablespaceSummarySharedEnum = "LOCAL_ON_LEAF"
	TablespaceSummarySharedLocalOnAll  TablespaceSummarySharedEnum = "LOCAL_ON_ALL"
)

var mappingTablespaceSummarySharedEnum = map[string]TablespaceSummarySharedEnum{
	"SHARED":        TablespaceSummarySharedShared,
	"LOCAL_ON_LEAF": TablespaceSummarySharedLocalOnLeaf,
	"LOCAL_ON_ALL":  TablespaceSummarySharedLocalOnAll,
}

var mappingTablespaceSummarySharedEnumLowerCase = map[string]TablespaceSummarySharedEnum{
	"shared":        TablespaceSummarySharedShared,
	"local_on_leaf": TablespaceSummarySharedLocalOnLeaf,
	"local_on_all":  TablespaceSummarySharedLocalOnAll,
}

// GetTablespaceSummarySharedEnumValues Enumerates the set of values for TablespaceSummarySharedEnum
func GetTablespaceSummarySharedEnumValues() []TablespaceSummarySharedEnum {
	values := make([]TablespaceSummarySharedEnum, 0)
	for _, v := range mappingTablespaceSummarySharedEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummarySharedEnumStringValues Enumerates the set of values in String for TablespaceSummarySharedEnum
func GetTablespaceSummarySharedEnumStringValues() []string {
	return []string{
		"SHARED",
		"LOCAL_ON_LEAF",
		"LOCAL_ON_ALL",
	}
}

// GetMappingTablespaceSummarySharedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummarySharedEnum(val string) (TablespaceSummarySharedEnum, bool) {
	enum, ok := mappingTablespaceSummarySharedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryDefaultIndexCompressionEnum Enum with underlying type: string
type TablespaceSummaryDefaultIndexCompressionEnum string

// Set of constants representing the allowable values for TablespaceSummaryDefaultIndexCompressionEnum
const (
	TablespaceSummaryDefaultIndexCompressionEnabled  TablespaceSummaryDefaultIndexCompressionEnum = "ENABLED"
	TablespaceSummaryDefaultIndexCompressionDisabled TablespaceSummaryDefaultIndexCompressionEnum = "DISABLED"
)

var mappingTablespaceSummaryDefaultIndexCompressionEnum = map[string]TablespaceSummaryDefaultIndexCompressionEnum{
	"ENABLED":  TablespaceSummaryDefaultIndexCompressionEnabled,
	"DISABLED": TablespaceSummaryDefaultIndexCompressionDisabled,
}

var mappingTablespaceSummaryDefaultIndexCompressionEnumLowerCase = map[string]TablespaceSummaryDefaultIndexCompressionEnum{
	"enabled":  TablespaceSummaryDefaultIndexCompressionEnabled,
	"disabled": TablespaceSummaryDefaultIndexCompressionDisabled,
}

// GetTablespaceSummaryDefaultIndexCompressionEnumValues Enumerates the set of values for TablespaceSummaryDefaultIndexCompressionEnum
func GetTablespaceSummaryDefaultIndexCompressionEnumValues() []TablespaceSummaryDefaultIndexCompressionEnum {
	values := make([]TablespaceSummaryDefaultIndexCompressionEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultIndexCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultIndexCompressionEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultIndexCompressionEnum
func GetTablespaceSummaryDefaultIndexCompressionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceSummaryDefaultIndexCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultIndexCompressionEnum(val string) (TablespaceSummaryDefaultIndexCompressionEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultIndexCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryIndexCompressForEnum Enum with underlying type: string
type TablespaceSummaryIndexCompressForEnum string

// Set of constants representing the allowable values for TablespaceSummaryIndexCompressForEnum
const (
	TablespaceSummaryIndexCompressForAdvancedLow  TablespaceSummaryIndexCompressForEnum = "ADVANCED_LOW"
	TablespaceSummaryIndexCompressForAdvancedHigh TablespaceSummaryIndexCompressForEnum = "ADVANCED_HIGH"
	TablespaceSummaryIndexCompressForNone         TablespaceSummaryIndexCompressForEnum = "NONE"
)

var mappingTablespaceSummaryIndexCompressForEnum = map[string]TablespaceSummaryIndexCompressForEnum{
	"ADVANCED_LOW":  TablespaceSummaryIndexCompressForAdvancedLow,
	"ADVANCED_HIGH": TablespaceSummaryIndexCompressForAdvancedHigh,
	"NONE":          TablespaceSummaryIndexCompressForNone,
}

var mappingTablespaceSummaryIndexCompressForEnumLowerCase = map[string]TablespaceSummaryIndexCompressForEnum{
	"advanced_low":  TablespaceSummaryIndexCompressForAdvancedLow,
	"advanced_high": TablespaceSummaryIndexCompressForAdvancedHigh,
	"none":          TablespaceSummaryIndexCompressForNone,
}

// GetTablespaceSummaryIndexCompressForEnumValues Enumerates the set of values for TablespaceSummaryIndexCompressForEnum
func GetTablespaceSummaryIndexCompressForEnumValues() []TablespaceSummaryIndexCompressForEnum {
	values := make([]TablespaceSummaryIndexCompressForEnum, 0)
	for _, v := range mappingTablespaceSummaryIndexCompressForEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryIndexCompressForEnumStringValues Enumerates the set of values in String for TablespaceSummaryIndexCompressForEnum
func GetTablespaceSummaryIndexCompressForEnumStringValues() []string {
	return []string{
		"ADVANCED_LOW",
		"ADVANCED_HIGH",
		"NONE",
	}
}

// GetMappingTablespaceSummaryIndexCompressForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryIndexCompressForEnum(val string) (TablespaceSummaryIndexCompressForEnum, bool) {
	enum, ok := mappingTablespaceSummaryIndexCompressForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTablespaceSummaryDefaultInMemoryServiceEnum = map[string]TablespaceSummaryDefaultInMemoryServiceEnum{
	"DEFAULT":      TablespaceSummaryDefaultInMemoryServiceDefault,
	"NONE":         TablespaceSummaryDefaultInMemoryServiceNone,
	"ALL":          TablespaceSummaryDefaultInMemoryServiceAll,
	"USER_DEFINED": TablespaceSummaryDefaultInMemoryServiceUserDefined,
}

var mappingTablespaceSummaryDefaultInMemoryServiceEnumLowerCase = map[string]TablespaceSummaryDefaultInMemoryServiceEnum{
	"default":      TablespaceSummaryDefaultInMemoryServiceDefault,
	"none":         TablespaceSummaryDefaultInMemoryServiceNone,
	"all":          TablespaceSummaryDefaultInMemoryServiceAll,
	"user_defined": TablespaceSummaryDefaultInMemoryServiceUserDefined,
}

// GetTablespaceSummaryDefaultInMemoryServiceEnumValues Enumerates the set of values for TablespaceSummaryDefaultInMemoryServiceEnum
func GetTablespaceSummaryDefaultInMemoryServiceEnumValues() []TablespaceSummaryDefaultInMemoryServiceEnum {
	values := make([]TablespaceSummaryDefaultInMemoryServiceEnum, 0)
	for _, v := range mappingTablespaceSummaryDefaultInMemoryServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryDefaultInMemoryServiceEnumStringValues Enumerates the set of values in String for TablespaceSummaryDefaultInMemoryServiceEnum
func GetTablespaceSummaryDefaultInMemoryServiceEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"NONE",
		"ALL",
		"USER_DEFINED",
	}
}

// GetMappingTablespaceSummaryDefaultInMemoryServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryDefaultInMemoryServiceEnum(val string) (TablespaceSummaryDefaultInMemoryServiceEnum, bool) {
	enum, ok := mappingTablespaceSummaryDefaultInMemoryServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSummaryLostWriteProtectEnum Enum with underlying type: string
type TablespaceSummaryLostWriteProtectEnum string

// Set of constants representing the allowable values for TablespaceSummaryLostWriteProtectEnum
const (
	TablespaceSummaryLostWriteProtectEnabled    TablespaceSummaryLostWriteProtectEnum = "ENABLED"
	TablespaceSummaryLostWriteProtectProtectOff TablespaceSummaryLostWriteProtectEnum = "PROTECT_OFF"
	TablespaceSummaryLostWriteProtectSuspend    TablespaceSummaryLostWriteProtectEnum = "SUSPEND"
)

var mappingTablespaceSummaryLostWriteProtectEnum = map[string]TablespaceSummaryLostWriteProtectEnum{
	"ENABLED":     TablespaceSummaryLostWriteProtectEnabled,
	"PROTECT_OFF": TablespaceSummaryLostWriteProtectProtectOff,
	"SUSPEND":     TablespaceSummaryLostWriteProtectSuspend,
}

var mappingTablespaceSummaryLostWriteProtectEnumLowerCase = map[string]TablespaceSummaryLostWriteProtectEnum{
	"enabled":     TablespaceSummaryLostWriteProtectEnabled,
	"protect_off": TablespaceSummaryLostWriteProtectProtectOff,
	"suspend":     TablespaceSummaryLostWriteProtectSuspend,
}

// GetTablespaceSummaryLostWriteProtectEnumValues Enumerates the set of values for TablespaceSummaryLostWriteProtectEnum
func GetTablespaceSummaryLostWriteProtectEnumValues() []TablespaceSummaryLostWriteProtectEnum {
	values := make([]TablespaceSummaryLostWriteProtectEnum, 0)
	for _, v := range mappingTablespaceSummaryLostWriteProtectEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSummaryLostWriteProtectEnumStringValues Enumerates the set of values in String for TablespaceSummaryLostWriteProtectEnum
func GetTablespaceSummaryLostWriteProtectEnumStringValues() []string {
	return []string{
		"ENABLED",
		"PROTECT_OFF",
		"SUSPEND",
	}
}

// GetMappingTablespaceSummaryLostWriteProtectEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSummaryLostWriteProtectEnum(val string) (TablespaceSummaryLostWriteProtectEnum, bool) {
	enum, ok := mappingTablespaceSummaryLostWriteProtectEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
