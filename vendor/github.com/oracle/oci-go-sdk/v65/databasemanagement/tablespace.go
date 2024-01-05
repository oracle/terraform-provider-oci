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

// Tablespace The details of a tablespace.
type Tablespace struct {

	// The name of the tablespace.
	Name *string `mandatory:"true" json:"name"`

	// The type of tablespace.
	Type TablespaceTypeEnum `mandatory:"true" json:"type"`

	// The status of the tablespace.
	Status TablespaceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The tablespace block size.
	BlockSizeBytes *float32 `mandatory:"false" json:"blockSizeBytes"`

	// The default logging attribute.
	Logging TablespaceLoggingEnum `mandatory:"false" json:"logging,omitempty"`

	// Indicates whether the tablespace is under Force Logging mode.
	IsForceLogging *bool `mandatory:"false" json:"isForceLogging"`

	// Indicates whether the extents in the tablespace are Locally managed or Dictionary managed.
	ExtentManagement TablespaceExtentManagementEnum `mandatory:"false" json:"extentManagement,omitempty"`

	// The type of extent allocation in effect for the tablespace.
	AllocationType TablespaceAllocationTypeEnum `mandatory:"false" json:"allocationType,omitempty"`

	// Indicates whether the tablespace is plugged in.
	IsPluggedIn *bool `mandatory:"false" json:"isPluggedIn"`

	// Indicates whether the free and used segment space in the tablespace is managed using free lists (MANUAL) or bitmaps (AUTO).
	SegmentSpaceManagement TablespaceSegmentSpaceManagementEnum `mandatory:"false" json:"segmentSpaceManagement,omitempty"`

	// Indicates whether default table compression is enabled or disabled.
	DefaultTableCompression TablespaceDefaultTableCompressionEnum `mandatory:"false" json:"defaultTableCompression,omitempty"`

	// Indicates whether undo retention guarantee is enabled for the tablespace.
	Retention TablespaceRetentionEnum `mandatory:"false" json:"retention,omitempty"`

	// Indicates whether the tablespace is a Bigfile tablespace or a Smallfile tablespace.
	IsBigfile *bool `mandatory:"false" json:"isBigfile"`

	// Indicates whether predicates are evaluated by Host or by Storage.
	PredicateEvaluation TablespacePredicateEvaluationEnum `mandatory:"false" json:"predicateEvaluation,omitempty"`

	// Indicates whether the tablespace is encrypted.
	IsEncrypted *bool `mandatory:"false" json:"isEncrypted"`

	// The operation type for which default compression is enabled.
	CompressFor TablespaceCompressForEnum `mandatory:"false" json:"compressFor,omitempty"`

	// Indicates whether the In-Memory Column Store (IM column store) is by default enabled or disabled for tables in the tablespace.
	DefaultInMemory TablespaceDefaultInMemoryEnum `mandatory:"false" json:"defaultInMemory,omitempty"`

	// Indicates the default priority for In-Memory Column Store (IM column store) population for the tablespace.
	DefaultInMemoryPriority TablespaceDefaultInMemoryPriorityEnum `mandatory:"false" json:"defaultInMemoryPriority,omitempty"`

	// Indicates how the IM column store is distributed by default for the tablespace in an Oracle Real Application Clusters (Oracle RAC) environment.
	DefaultInMemoryDistribute TablespaceDefaultInMemoryDistributeEnum `mandatory:"false" json:"defaultInMemoryDistribute,omitempty"`

	// Indicates the default compression level for the IM column store for the tablespace.
	DefaultInMemoryCompression TablespaceDefaultInMemoryCompressionEnum `mandatory:"false" json:"defaultInMemoryCompression,omitempty"`

	// Indicates the duplicate setting for the IM column store in an Oracle RAC environment.
	DefaultInMemoryDuplicate TablespaceDefaultInMemoryDuplicateEnum `mandatory:"false" json:"defaultInMemoryDuplicate,omitempty"`

	// Indicates whether the tablespace is for shared tablespace, or for local temporary tablespace for leaf (read-only) instances, or for local temporary tablespace for all instance types.
	Shared TablespaceSharedEnum `mandatory:"false" json:"shared,omitempty"`

	// Indicates whether default index compression is enabled or disabled.
	DefaultIndexCompression TablespaceDefaultIndexCompressionEnum `mandatory:"false" json:"defaultIndexCompression,omitempty"`

	// The operation type for which default index compression is enabled.
	IndexCompressFor TablespaceIndexCompressForEnum `mandatory:"false" json:"indexCompressFor,omitempty"`

	// This specifies the default value for the CELLMEMORY attribute that tables created in the tablespace will inherit unless the behavior is overridden explicitly. This column is intended for use with Oracle Exadata.
	DefaultCellMemory *string `mandatory:"false" json:"defaultCellMemory"`

	// Indicates how the IM column store is populated on various instances by default for the tablespace.
	DefaultInMemoryService TablespaceDefaultInMemoryServiceEnum `mandatory:"false" json:"defaultInMemoryService,omitempty"`

	// Indicates the service name for the service on which the IM column store should be populated by default for the tablespace. This column has a value only when the corresponding DEF_INMEMORY_SERVICE is USER_DEFINED. In all other cases, this column is null.
	DefaultInMemoryServiceName *string `mandatory:"false" json:"defaultInMemoryServiceName"`

	// The lost write protection setting for the tablespace.
	LostWriteProtect TablespaceLostWriteProtectEnum `mandatory:"false" json:"lostWriteProtect,omitempty"`

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

func (m Tablespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tablespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTablespaceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetTablespaceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTablespaceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTablespaceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceLoggingEnum(string(m.Logging)); !ok && m.Logging != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Logging: %s. Supported values are: %s.", m.Logging, strings.Join(GetTablespaceLoggingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceExtentManagementEnum(string(m.ExtentManagement)); !ok && m.ExtentManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExtentManagement: %s. Supported values are: %s.", m.ExtentManagement, strings.Join(GetTablespaceExtentManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceAllocationTypeEnum(string(m.AllocationType)); !ok && m.AllocationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllocationType: %s. Supported values are: %s.", m.AllocationType, strings.Join(GetTablespaceAllocationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSegmentSpaceManagementEnum(string(m.SegmentSpaceManagement)); !ok && m.SegmentSpaceManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SegmentSpaceManagement: %s. Supported values are: %s.", m.SegmentSpaceManagement, strings.Join(GetTablespaceSegmentSpaceManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultTableCompressionEnum(string(m.DefaultTableCompression)); !ok && m.DefaultTableCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultTableCompression: %s. Supported values are: %s.", m.DefaultTableCompression, strings.Join(GetTablespaceDefaultTableCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceRetentionEnum(string(m.Retention)); !ok && m.Retention != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Retention: %s. Supported values are: %s.", m.Retention, strings.Join(GetTablespaceRetentionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespacePredicateEvaluationEnum(string(m.PredicateEvaluation)); !ok && m.PredicateEvaluation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PredicateEvaluation: %s. Supported values are: %s.", m.PredicateEvaluation, strings.Join(GetTablespacePredicateEvaluationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceCompressForEnum(string(m.CompressFor)); !ok && m.CompressFor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompressFor: %s. Supported values are: %s.", m.CompressFor, strings.Join(GetTablespaceCompressForEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryEnum(string(m.DefaultInMemory)); !ok && m.DefaultInMemory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemory: %s. Supported values are: %s.", m.DefaultInMemory, strings.Join(GetTablespaceDefaultInMemoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryPriorityEnum(string(m.DefaultInMemoryPriority)); !ok && m.DefaultInMemoryPriority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryPriority: %s. Supported values are: %s.", m.DefaultInMemoryPriority, strings.Join(GetTablespaceDefaultInMemoryPriorityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryDistributeEnum(string(m.DefaultInMemoryDistribute)); !ok && m.DefaultInMemoryDistribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryDistribute: %s. Supported values are: %s.", m.DefaultInMemoryDistribute, strings.Join(GetTablespaceDefaultInMemoryDistributeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryCompressionEnum(string(m.DefaultInMemoryCompression)); !ok && m.DefaultInMemoryCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryCompression: %s. Supported values are: %s.", m.DefaultInMemoryCompression, strings.Join(GetTablespaceDefaultInMemoryCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryDuplicateEnum(string(m.DefaultInMemoryDuplicate)); !ok && m.DefaultInMemoryDuplicate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryDuplicate: %s. Supported values are: %s.", m.DefaultInMemoryDuplicate, strings.Join(GetTablespaceDefaultInMemoryDuplicateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceSharedEnum(string(m.Shared)); !ok && m.Shared != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shared: %s. Supported values are: %s.", m.Shared, strings.Join(GetTablespaceSharedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultIndexCompressionEnum(string(m.DefaultIndexCompression)); !ok && m.DefaultIndexCompression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultIndexCompression: %s. Supported values are: %s.", m.DefaultIndexCompression, strings.Join(GetTablespaceDefaultIndexCompressionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceIndexCompressForEnum(string(m.IndexCompressFor)); !ok && m.IndexCompressFor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IndexCompressFor: %s. Supported values are: %s.", m.IndexCompressFor, strings.Join(GetTablespaceIndexCompressForEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceDefaultInMemoryServiceEnum(string(m.DefaultInMemoryService)); !ok && m.DefaultInMemoryService != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultInMemoryService: %s. Supported values are: %s.", m.DefaultInMemoryService, strings.Join(GetTablespaceDefaultInMemoryServiceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTablespaceLostWriteProtectEnum(string(m.LostWriteProtect)); !ok && m.LostWriteProtect != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LostWriteProtect: %s. Supported values are: %s.", m.LostWriteProtect, strings.Join(GetTablespaceLostWriteProtectEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TablespaceTypeEnum Enum with underlying type: string
type TablespaceTypeEnum string

// Set of constants representing the allowable values for TablespaceTypeEnum
const (
	TablespaceTypeUndo                TablespaceTypeEnum = "UNDO"
	TablespaceTypeLostWriteProtection TablespaceTypeEnum = "LOST_WRITE_PROTECTION"
	TablespaceTypePermanent           TablespaceTypeEnum = "PERMANENT"
	TablespaceTypeTemporary           TablespaceTypeEnum = "TEMPORARY"
)

var mappingTablespaceTypeEnum = map[string]TablespaceTypeEnum{
	"UNDO":                  TablespaceTypeUndo,
	"LOST_WRITE_PROTECTION": TablespaceTypeLostWriteProtection,
	"PERMANENT":             TablespaceTypePermanent,
	"TEMPORARY":             TablespaceTypeTemporary,
}

var mappingTablespaceTypeEnumLowerCase = map[string]TablespaceTypeEnum{
	"undo":                  TablespaceTypeUndo,
	"lost_write_protection": TablespaceTypeLostWriteProtection,
	"permanent":             TablespaceTypePermanent,
	"temporary":             TablespaceTypeTemporary,
}

// GetTablespaceTypeEnumValues Enumerates the set of values for TablespaceTypeEnum
func GetTablespaceTypeEnumValues() []TablespaceTypeEnum {
	values := make([]TablespaceTypeEnum, 0)
	for _, v := range mappingTablespaceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceTypeEnumStringValues Enumerates the set of values in String for TablespaceTypeEnum
func GetTablespaceTypeEnumStringValues() []string {
	return []string{
		"UNDO",
		"LOST_WRITE_PROTECTION",
		"PERMANENT",
		"TEMPORARY",
	}
}

// GetMappingTablespaceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceTypeEnum(val string) (TablespaceTypeEnum, bool) {
	enum, ok := mappingTablespaceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceStatusEnum Enum with underlying type: string
type TablespaceStatusEnum string

// Set of constants representing the allowable values for TablespaceStatusEnum
const (
	TablespaceStatusOnline   TablespaceStatusEnum = "ONLINE"
	TablespaceStatusOffline  TablespaceStatusEnum = "OFFLINE"
	TablespaceStatusReadOnly TablespaceStatusEnum = "READ_ONLY"
)

var mappingTablespaceStatusEnum = map[string]TablespaceStatusEnum{
	"ONLINE":    TablespaceStatusOnline,
	"OFFLINE":   TablespaceStatusOffline,
	"READ_ONLY": TablespaceStatusReadOnly,
}

var mappingTablespaceStatusEnumLowerCase = map[string]TablespaceStatusEnum{
	"online":    TablespaceStatusOnline,
	"offline":   TablespaceStatusOffline,
	"read_only": TablespaceStatusReadOnly,
}

// GetTablespaceStatusEnumValues Enumerates the set of values for TablespaceStatusEnum
func GetTablespaceStatusEnumValues() []TablespaceStatusEnum {
	values := make([]TablespaceStatusEnum, 0)
	for _, v := range mappingTablespaceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceStatusEnumStringValues Enumerates the set of values in String for TablespaceStatusEnum
func GetTablespaceStatusEnumStringValues() []string {
	return []string{
		"ONLINE",
		"OFFLINE",
		"READ_ONLY",
	}
}

// GetMappingTablespaceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceStatusEnum(val string) (TablespaceStatusEnum, bool) {
	enum, ok := mappingTablespaceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceLoggingEnum Enum with underlying type: string
type TablespaceLoggingEnum string

// Set of constants representing the allowable values for TablespaceLoggingEnum
const (
	TablespaceLoggingLogging   TablespaceLoggingEnum = "LOGGING"
	TablespaceLoggingNologging TablespaceLoggingEnum = "NOLOGGING"
)

var mappingTablespaceLoggingEnum = map[string]TablespaceLoggingEnum{
	"LOGGING":   TablespaceLoggingLogging,
	"NOLOGGING": TablespaceLoggingNologging,
}

var mappingTablespaceLoggingEnumLowerCase = map[string]TablespaceLoggingEnum{
	"logging":   TablespaceLoggingLogging,
	"nologging": TablespaceLoggingNologging,
}

// GetTablespaceLoggingEnumValues Enumerates the set of values for TablespaceLoggingEnum
func GetTablespaceLoggingEnumValues() []TablespaceLoggingEnum {
	values := make([]TablespaceLoggingEnum, 0)
	for _, v := range mappingTablespaceLoggingEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceLoggingEnumStringValues Enumerates the set of values in String for TablespaceLoggingEnum
func GetTablespaceLoggingEnumStringValues() []string {
	return []string{
		"LOGGING",
		"NOLOGGING",
	}
}

// GetMappingTablespaceLoggingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceLoggingEnum(val string) (TablespaceLoggingEnum, bool) {
	enum, ok := mappingTablespaceLoggingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceExtentManagementEnum Enum with underlying type: string
type TablespaceExtentManagementEnum string

// Set of constants representing the allowable values for TablespaceExtentManagementEnum
const (
	TablespaceExtentManagementLocal      TablespaceExtentManagementEnum = "LOCAL"
	TablespaceExtentManagementDictionary TablespaceExtentManagementEnum = "DICTIONARY"
)

var mappingTablespaceExtentManagementEnum = map[string]TablespaceExtentManagementEnum{
	"LOCAL":      TablespaceExtentManagementLocal,
	"DICTIONARY": TablespaceExtentManagementDictionary,
}

var mappingTablespaceExtentManagementEnumLowerCase = map[string]TablespaceExtentManagementEnum{
	"local":      TablespaceExtentManagementLocal,
	"dictionary": TablespaceExtentManagementDictionary,
}

// GetTablespaceExtentManagementEnumValues Enumerates the set of values for TablespaceExtentManagementEnum
func GetTablespaceExtentManagementEnumValues() []TablespaceExtentManagementEnum {
	values := make([]TablespaceExtentManagementEnum, 0)
	for _, v := range mappingTablespaceExtentManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceExtentManagementEnumStringValues Enumerates the set of values in String for TablespaceExtentManagementEnum
func GetTablespaceExtentManagementEnumStringValues() []string {
	return []string{
		"LOCAL",
		"DICTIONARY",
	}
}

// GetMappingTablespaceExtentManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceExtentManagementEnum(val string) (TablespaceExtentManagementEnum, bool) {
	enum, ok := mappingTablespaceExtentManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceAllocationTypeEnum Enum with underlying type: string
type TablespaceAllocationTypeEnum string

// Set of constants representing the allowable values for TablespaceAllocationTypeEnum
const (
	TablespaceAllocationTypeSystem  TablespaceAllocationTypeEnum = "SYSTEM"
	TablespaceAllocationTypeUniform TablespaceAllocationTypeEnum = "UNIFORM"
	TablespaceAllocationTypeUser    TablespaceAllocationTypeEnum = "USER"
)

var mappingTablespaceAllocationTypeEnum = map[string]TablespaceAllocationTypeEnum{
	"SYSTEM":  TablespaceAllocationTypeSystem,
	"UNIFORM": TablespaceAllocationTypeUniform,
	"USER":    TablespaceAllocationTypeUser,
}

var mappingTablespaceAllocationTypeEnumLowerCase = map[string]TablespaceAllocationTypeEnum{
	"system":  TablespaceAllocationTypeSystem,
	"uniform": TablespaceAllocationTypeUniform,
	"user":    TablespaceAllocationTypeUser,
}

// GetTablespaceAllocationTypeEnumValues Enumerates the set of values for TablespaceAllocationTypeEnum
func GetTablespaceAllocationTypeEnumValues() []TablespaceAllocationTypeEnum {
	values := make([]TablespaceAllocationTypeEnum, 0)
	for _, v := range mappingTablespaceAllocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceAllocationTypeEnumStringValues Enumerates the set of values in String for TablespaceAllocationTypeEnum
func GetTablespaceAllocationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"UNIFORM",
		"USER",
	}
}

// GetMappingTablespaceAllocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceAllocationTypeEnum(val string) (TablespaceAllocationTypeEnum, bool) {
	enum, ok := mappingTablespaceAllocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSegmentSpaceManagementEnum Enum with underlying type: string
type TablespaceSegmentSpaceManagementEnum string

// Set of constants representing the allowable values for TablespaceSegmentSpaceManagementEnum
const (
	TablespaceSegmentSpaceManagementManual TablespaceSegmentSpaceManagementEnum = "MANUAL"
	TablespaceSegmentSpaceManagementAuto   TablespaceSegmentSpaceManagementEnum = "AUTO"
)

var mappingTablespaceSegmentSpaceManagementEnum = map[string]TablespaceSegmentSpaceManagementEnum{
	"MANUAL": TablespaceSegmentSpaceManagementManual,
	"AUTO":   TablespaceSegmentSpaceManagementAuto,
}

var mappingTablespaceSegmentSpaceManagementEnumLowerCase = map[string]TablespaceSegmentSpaceManagementEnum{
	"manual": TablespaceSegmentSpaceManagementManual,
	"auto":   TablespaceSegmentSpaceManagementAuto,
}

// GetTablespaceSegmentSpaceManagementEnumValues Enumerates the set of values for TablespaceSegmentSpaceManagementEnum
func GetTablespaceSegmentSpaceManagementEnumValues() []TablespaceSegmentSpaceManagementEnum {
	values := make([]TablespaceSegmentSpaceManagementEnum, 0)
	for _, v := range mappingTablespaceSegmentSpaceManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSegmentSpaceManagementEnumStringValues Enumerates the set of values in String for TablespaceSegmentSpaceManagementEnum
func GetTablespaceSegmentSpaceManagementEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTO",
	}
}

// GetMappingTablespaceSegmentSpaceManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSegmentSpaceManagementEnum(val string) (TablespaceSegmentSpaceManagementEnum, bool) {
	enum, ok := mappingTablespaceSegmentSpaceManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultTableCompressionEnum Enum with underlying type: string
type TablespaceDefaultTableCompressionEnum string

// Set of constants representing the allowable values for TablespaceDefaultTableCompressionEnum
const (
	TablespaceDefaultTableCompressionEnabled  TablespaceDefaultTableCompressionEnum = "ENABLED"
	TablespaceDefaultTableCompressionDisabled TablespaceDefaultTableCompressionEnum = "DISABLED"
)

var mappingTablespaceDefaultTableCompressionEnum = map[string]TablespaceDefaultTableCompressionEnum{
	"ENABLED":  TablespaceDefaultTableCompressionEnabled,
	"DISABLED": TablespaceDefaultTableCompressionDisabled,
}

var mappingTablespaceDefaultTableCompressionEnumLowerCase = map[string]TablespaceDefaultTableCompressionEnum{
	"enabled":  TablespaceDefaultTableCompressionEnabled,
	"disabled": TablespaceDefaultTableCompressionDisabled,
}

// GetTablespaceDefaultTableCompressionEnumValues Enumerates the set of values for TablespaceDefaultTableCompressionEnum
func GetTablespaceDefaultTableCompressionEnumValues() []TablespaceDefaultTableCompressionEnum {
	values := make([]TablespaceDefaultTableCompressionEnum, 0)
	for _, v := range mappingTablespaceDefaultTableCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultTableCompressionEnumStringValues Enumerates the set of values in String for TablespaceDefaultTableCompressionEnum
func GetTablespaceDefaultTableCompressionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceDefaultTableCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultTableCompressionEnum(val string) (TablespaceDefaultTableCompressionEnum, bool) {
	enum, ok := mappingTablespaceDefaultTableCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceRetentionEnum Enum with underlying type: string
type TablespaceRetentionEnum string

// Set of constants representing the allowable values for TablespaceRetentionEnum
const (
	TablespaceRetentionGuarantee   TablespaceRetentionEnum = "GUARANTEE"
	TablespaceRetentionNoguarantee TablespaceRetentionEnum = "NOGUARANTEE"
	TablespaceRetentionNotApply    TablespaceRetentionEnum = "NOT_APPLY"
)

var mappingTablespaceRetentionEnum = map[string]TablespaceRetentionEnum{
	"GUARANTEE":   TablespaceRetentionGuarantee,
	"NOGUARANTEE": TablespaceRetentionNoguarantee,
	"NOT_APPLY":   TablespaceRetentionNotApply,
}

var mappingTablespaceRetentionEnumLowerCase = map[string]TablespaceRetentionEnum{
	"guarantee":   TablespaceRetentionGuarantee,
	"noguarantee": TablespaceRetentionNoguarantee,
	"not_apply":   TablespaceRetentionNotApply,
}

// GetTablespaceRetentionEnumValues Enumerates the set of values for TablespaceRetentionEnum
func GetTablespaceRetentionEnumValues() []TablespaceRetentionEnum {
	values := make([]TablespaceRetentionEnum, 0)
	for _, v := range mappingTablespaceRetentionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceRetentionEnumStringValues Enumerates the set of values in String for TablespaceRetentionEnum
func GetTablespaceRetentionEnumStringValues() []string {
	return []string{
		"GUARANTEE",
		"NOGUARANTEE",
		"NOT_APPLY",
	}
}

// GetMappingTablespaceRetentionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceRetentionEnum(val string) (TablespaceRetentionEnum, bool) {
	enum, ok := mappingTablespaceRetentionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespacePredicateEvaluationEnum Enum with underlying type: string
type TablespacePredicateEvaluationEnum string

// Set of constants representing the allowable values for TablespacePredicateEvaluationEnum
const (
	TablespacePredicateEvaluationHost    TablespacePredicateEvaluationEnum = "HOST"
	TablespacePredicateEvaluationStorage TablespacePredicateEvaluationEnum = "STORAGE"
)

var mappingTablespacePredicateEvaluationEnum = map[string]TablespacePredicateEvaluationEnum{
	"HOST":    TablespacePredicateEvaluationHost,
	"STORAGE": TablespacePredicateEvaluationStorage,
}

var mappingTablespacePredicateEvaluationEnumLowerCase = map[string]TablespacePredicateEvaluationEnum{
	"host":    TablespacePredicateEvaluationHost,
	"storage": TablespacePredicateEvaluationStorage,
}

// GetTablespacePredicateEvaluationEnumValues Enumerates the set of values for TablespacePredicateEvaluationEnum
func GetTablespacePredicateEvaluationEnumValues() []TablespacePredicateEvaluationEnum {
	values := make([]TablespacePredicateEvaluationEnum, 0)
	for _, v := range mappingTablespacePredicateEvaluationEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespacePredicateEvaluationEnumStringValues Enumerates the set of values in String for TablespacePredicateEvaluationEnum
func GetTablespacePredicateEvaluationEnumStringValues() []string {
	return []string{
		"HOST",
		"STORAGE",
	}
}

// GetMappingTablespacePredicateEvaluationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespacePredicateEvaluationEnum(val string) (TablespacePredicateEvaluationEnum, bool) {
	enum, ok := mappingTablespacePredicateEvaluationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceCompressForEnum Enum with underlying type: string
type TablespaceCompressForEnum string

// Set of constants representing the allowable values for TablespaceCompressForEnum
const (
	TablespaceCompressForBasic            TablespaceCompressForEnum = "BASIC"
	TablespaceCompressForAdvanced         TablespaceCompressForEnum = "ADVANCED"
	TablespaceCompressForQueryLow         TablespaceCompressForEnum = "QUERY_LOW"
	TablespaceCompressForQueryHigh        TablespaceCompressForEnum = "QUERY_HIGH"
	TablespaceCompressForArchiveLow       TablespaceCompressForEnum = "ARCHIVE_LOW"
	TablespaceCompressForArchiveHigh      TablespaceCompressForEnum = "ARCHIVE_HIGH"
	TablespaceCompressForDirectLoadOnly   TablespaceCompressForEnum = "DIRECT_LOAD_ONLY"
	TablespaceCompressForForAllOperations TablespaceCompressForEnum = "FOR_ALL_OPERATIONS"
)

var mappingTablespaceCompressForEnum = map[string]TablespaceCompressForEnum{
	"BASIC":              TablespaceCompressForBasic,
	"ADVANCED":           TablespaceCompressForAdvanced,
	"QUERY_LOW":          TablespaceCompressForQueryLow,
	"QUERY_HIGH":         TablespaceCompressForQueryHigh,
	"ARCHIVE_LOW":        TablespaceCompressForArchiveLow,
	"ARCHIVE_HIGH":       TablespaceCompressForArchiveHigh,
	"DIRECT_LOAD_ONLY":   TablespaceCompressForDirectLoadOnly,
	"FOR_ALL_OPERATIONS": TablespaceCompressForForAllOperations,
}

var mappingTablespaceCompressForEnumLowerCase = map[string]TablespaceCompressForEnum{
	"basic":              TablespaceCompressForBasic,
	"advanced":           TablespaceCompressForAdvanced,
	"query_low":          TablespaceCompressForQueryLow,
	"query_high":         TablespaceCompressForQueryHigh,
	"archive_low":        TablespaceCompressForArchiveLow,
	"archive_high":       TablespaceCompressForArchiveHigh,
	"direct_load_only":   TablespaceCompressForDirectLoadOnly,
	"for_all_operations": TablespaceCompressForForAllOperations,
}

// GetTablespaceCompressForEnumValues Enumerates the set of values for TablespaceCompressForEnum
func GetTablespaceCompressForEnumValues() []TablespaceCompressForEnum {
	values := make([]TablespaceCompressForEnum, 0)
	for _, v := range mappingTablespaceCompressForEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceCompressForEnumStringValues Enumerates the set of values in String for TablespaceCompressForEnum
func GetTablespaceCompressForEnumStringValues() []string {
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

// GetMappingTablespaceCompressForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceCompressForEnum(val string) (TablespaceCompressForEnum, bool) {
	enum, ok := mappingTablespaceCompressForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryEnum Enum with underlying type: string
type TablespaceDefaultInMemoryEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryEnum
const (
	TablespaceDefaultInMemoryEnabled  TablespaceDefaultInMemoryEnum = "ENABLED"
	TablespaceDefaultInMemoryDisabled TablespaceDefaultInMemoryEnum = "DISABLED"
)

var mappingTablespaceDefaultInMemoryEnum = map[string]TablespaceDefaultInMemoryEnum{
	"ENABLED":  TablespaceDefaultInMemoryEnabled,
	"DISABLED": TablespaceDefaultInMemoryDisabled,
}

var mappingTablespaceDefaultInMemoryEnumLowerCase = map[string]TablespaceDefaultInMemoryEnum{
	"enabled":  TablespaceDefaultInMemoryEnabled,
	"disabled": TablespaceDefaultInMemoryDisabled,
}

// GetTablespaceDefaultInMemoryEnumValues Enumerates the set of values for TablespaceDefaultInMemoryEnum
func GetTablespaceDefaultInMemoryEnumValues() []TablespaceDefaultInMemoryEnum {
	values := make([]TablespaceDefaultInMemoryEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryEnum
func GetTablespaceDefaultInMemoryEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceDefaultInMemoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryEnum(val string) (TablespaceDefaultInMemoryEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryPriorityEnum Enum with underlying type: string
type TablespaceDefaultInMemoryPriorityEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryPriorityEnum
const (
	TablespaceDefaultInMemoryPriorityLow      TablespaceDefaultInMemoryPriorityEnum = "LOW"
	TablespaceDefaultInMemoryPriorityMedium   TablespaceDefaultInMemoryPriorityEnum = "MEDIUM"
	TablespaceDefaultInMemoryPriorityHigh     TablespaceDefaultInMemoryPriorityEnum = "HIGH"
	TablespaceDefaultInMemoryPriorityCritical TablespaceDefaultInMemoryPriorityEnum = "CRITICAL"
	TablespaceDefaultInMemoryPriorityNone     TablespaceDefaultInMemoryPriorityEnum = "NONE"
)

var mappingTablespaceDefaultInMemoryPriorityEnum = map[string]TablespaceDefaultInMemoryPriorityEnum{
	"LOW":      TablespaceDefaultInMemoryPriorityLow,
	"MEDIUM":   TablespaceDefaultInMemoryPriorityMedium,
	"HIGH":     TablespaceDefaultInMemoryPriorityHigh,
	"CRITICAL": TablespaceDefaultInMemoryPriorityCritical,
	"NONE":     TablespaceDefaultInMemoryPriorityNone,
}

var mappingTablespaceDefaultInMemoryPriorityEnumLowerCase = map[string]TablespaceDefaultInMemoryPriorityEnum{
	"low":      TablespaceDefaultInMemoryPriorityLow,
	"medium":   TablespaceDefaultInMemoryPriorityMedium,
	"high":     TablespaceDefaultInMemoryPriorityHigh,
	"critical": TablespaceDefaultInMemoryPriorityCritical,
	"none":     TablespaceDefaultInMemoryPriorityNone,
}

// GetTablespaceDefaultInMemoryPriorityEnumValues Enumerates the set of values for TablespaceDefaultInMemoryPriorityEnum
func GetTablespaceDefaultInMemoryPriorityEnumValues() []TablespaceDefaultInMemoryPriorityEnum {
	values := make([]TablespaceDefaultInMemoryPriorityEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryPriorityEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryPriorityEnum
func GetTablespaceDefaultInMemoryPriorityEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
		"NONE",
	}
}

// GetMappingTablespaceDefaultInMemoryPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryPriorityEnum(val string) (TablespaceDefaultInMemoryPriorityEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryPriorityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryDistributeEnum Enum with underlying type: string
type TablespaceDefaultInMemoryDistributeEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryDistributeEnum
const (
	TablespaceDefaultInMemoryDistributeAuto           TablespaceDefaultInMemoryDistributeEnum = "AUTO"
	TablespaceDefaultInMemoryDistributeByRowidRange   TablespaceDefaultInMemoryDistributeEnum = "BY_ROWID_RANGE"
	TablespaceDefaultInMemoryDistributeByPartition    TablespaceDefaultInMemoryDistributeEnum = "BY_PARTITION"
	TablespaceDefaultInMemoryDistributeBySubpartition TablespaceDefaultInMemoryDistributeEnum = "BY_SUBPARTITION"
)

var mappingTablespaceDefaultInMemoryDistributeEnum = map[string]TablespaceDefaultInMemoryDistributeEnum{
	"AUTO":            TablespaceDefaultInMemoryDistributeAuto,
	"BY_ROWID_RANGE":  TablespaceDefaultInMemoryDistributeByRowidRange,
	"BY_PARTITION":    TablespaceDefaultInMemoryDistributeByPartition,
	"BY_SUBPARTITION": TablespaceDefaultInMemoryDistributeBySubpartition,
}

var mappingTablespaceDefaultInMemoryDistributeEnumLowerCase = map[string]TablespaceDefaultInMemoryDistributeEnum{
	"auto":            TablespaceDefaultInMemoryDistributeAuto,
	"by_rowid_range":  TablespaceDefaultInMemoryDistributeByRowidRange,
	"by_partition":    TablespaceDefaultInMemoryDistributeByPartition,
	"by_subpartition": TablespaceDefaultInMemoryDistributeBySubpartition,
}

// GetTablespaceDefaultInMemoryDistributeEnumValues Enumerates the set of values for TablespaceDefaultInMemoryDistributeEnum
func GetTablespaceDefaultInMemoryDistributeEnumValues() []TablespaceDefaultInMemoryDistributeEnum {
	values := make([]TablespaceDefaultInMemoryDistributeEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryDistributeEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryDistributeEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryDistributeEnum
func GetTablespaceDefaultInMemoryDistributeEnumStringValues() []string {
	return []string{
		"AUTO",
		"BY_ROWID_RANGE",
		"BY_PARTITION",
		"BY_SUBPARTITION",
	}
}

// GetMappingTablespaceDefaultInMemoryDistributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryDistributeEnum(val string) (TablespaceDefaultInMemoryDistributeEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryDistributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryCompressionEnum Enum with underlying type: string
type TablespaceDefaultInMemoryCompressionEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryCompressionEnum
const (
	TablespaceDefaultInMemoryCompressionNoMemcompress   TablespaceDefaultInMemoryCompressionEnum = "NO_MEMCOMPRESS"
	TablespaceDefaultInMemoryCompressionForDml          TablespaceDefaultInMemoryCompressionEnum = "FOR_DML"
	TablespaceDefaultInMemoryCompressionForQueryLow     TablespaceDefaultInMemoryCompressionEnum = "FOR_QUERY_LOW"
	TablespaceDefaultInMemoryCompressionForQueryHigh    TablespaceDefaultInMemoryCompressionEnum = "FOR_QUERY_HIGH"
	TablespaceDefaultInMemoryCompressionForCapacityLow  TablespaceDefaultInMemoryCompressionEnum = "FOR_CAPACITY_LOW"
	TablespaceDefaultInMemoryCompressionForCapacityHigh TablespaceDefaultInMemoryCompressionEnum = "FOR_CAPACITY_HIGH"
)

var mappingTablespaceDefaultInMemoryCompressionEnum = map[string]TablespaceDefaultInMemoryCompressionEnum{
	"NO_MEMCOMPRESS":    TablespaceDefaultInMemoryCompressionNoMemcompress,
	"FOR_DML":           TablespaceDefaultInMemoryCompressionForDml,
	"FOR_QUERY_LOW":     TablespaceDefaultInMemoryCompressionForQueryLow,
	"FOR_QUERY_HIGH":    TablespaceDefaultInMemoryCompressionForQueryHigh,
	"FOR_CAPACITY_LOW":  TablespaceDefaultInMemoryCompressionForCapacityLow,
	"FOR_CAPACITY_HIGH": TablespaceDefaultInMemoryCompressionForCapacityHigh,
}

var mappingTablespaceDefaultInMemoryCompressionEnumLowerCase = map[string]TablespaceDefaultInMemoryCompressionEnum{
	"no_memcompress":    TablespaceDefaultInMemoryCompressionNoMemcompress,
	"for_dml":           TablespaceDefaultInMemoryCompressionForDml,
	"for_query_low":     TablespaceDefaultInMemoryCompressionForQueryLow,
	"for_query_high":    TablespaceDefaultInMemoryCompressionForQueryHigh,
	"for_capacity_low":  TablespaceDefaultInMemoryCompressionForCapacityLow,
	"for_capacity_high": TablespaceDefaultInMemoryCompressionForCapacityHigh,
}

// GetTablespaceDefaultInMemoryCompressionEnumValues Enumerates the set of values for TablespaceDefaultInMemoryCompressionEnum
func GetTablespaceDefaultInMemoryCompressionEnumValues() []TablespaceDefaultInMemoryCompressionEnum {
	values := make([]TablespaceDefaultInMemoryCompressionEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryCompressionEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryCompressionEnum
func GetTablespaceDefaultInMemoryCompressionEnumStringValues() []string {
	return []string{
		"NO_MEMCOMPRESS",
		"FOR_DML",
		"FOR_QUERY_LOW",
		"FOR_QUERY_HIGH",
		"FOR_CAPACITY_LOW",
		"FOR_CAPACITY_HIGH",
	}
}

// GetMappingTablespaceDefaultInMemoryCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryCompressionEnum(val string) (TablespaceDefaultInMemoryCompressionEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryDuplicateEnum Enum with underlying type: string
type TablespaceDefaultInMemoryDuplicateEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryDuplicateEnum
const (
	TablespaceDefaultInMemoryDuplicateNoDuplicate  TablespaceDefaultInMemoryDuplicateEnum = "NO_DUPLICATE"
	TablespaceDefaultInMemoryDuplicateDuplicate    TablespaceDefaultInMemoryDuplicateEnum = "DUPLICATE"
	TablespaceDefaultInMemoryDuplicateDuplicateAll TablespaceDefaultInMemoryDuplicateEnum = "DUPLICATE_ALL"
)

var mappingTablespaceDefaultInMemoryDuplicateEnum = map[string]TablespaceDefaultInMemoryDuplicateEnum{
	"NO_DUPLICATE":  TablespaceDefaultInMemoryDuplicateNoDuplicate,
	"DUPLICATE":     TablespaceDefaultInMemoryDuplicateDuplicate,
	"DUPLICATE_ALL": TablespaceDefaultInMemoryDuplicateDuplicateAll,
}

var mappingTablespaceDefaultInMemoryDuplicateEnumLowerCase = map[string]TablespaceDefaultInMemoryDuplicateEnum{
	"no_duplicate":  TablespaceDefaultInMemoryDuplicateNoDuplicate,
	"duplicate":     TablespaceDefaultInMemoryDuplicateDuplicate,
	"duplicate_all": TablespaceDefaultInMemoryDuplicateDuplicateAll,
}

// GetTablespaceDefaultInMemoryDuplicateEnumValues Enumerates the set of values for TablespaceDefaultInMemoryDuplicateEnum
func GetTablespaceDefaultInMemoryDuplicateEnumValues() []TablespaceDefaultInMemoryDuplicateEnum {
	values := make([]TablespaceDefaultInMemoryDuplicateEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryDuplicateEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryDuplicateEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryDuplicateEnum
func GetTablespaceDefaultInMemoryDuplicateEnumStringValues() []string {
	return []string{
		"NO_DUPLICATE",
		"DUPLICATE",
		"DUPLICATE_ALL",
	}
}

// GetMappingTablespaceDefaultInMemoryDuplicateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryDuplicateEnum(val string) (TablespaceDefaultInMemoryDuplicateEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryDuplicateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceSharedEnum Enum with underlying type: string
type TablespaceSharedEnum string

// Set of constants representing the allowable values for TablespaceSharedEnum
const (
	TablespaceSharedShared      TablespaceSharedEnum = "SHARED"
	TablespaceSharedLocalOnLeaf TablespaceSharedEnum = "LOCAL_ON_LEAF"
	TablespaceSharedLocalOnAll  TablespaceSharedEnum = "LOCAL_ON_ALL"
)

var mappingTablespaceSharedEnum = map[string]TablespaceSharedEnum{
	"SHARED":        TablespaceSharedShared,
	"LOCAL_ON_LEAF": TablespaceSharedLocalOnLeaf,
	"LOCAL_ON_ALL":  TablespaceSharedLocalOnAll,
}

var mappingTablespaceSharedEnumLowerCase = map[string]TablespaceSharedEnum{
	"shared":        TablespaceSharedShared,
	"local_on_leaf": TablespaceSharedLocalOnLeaf,
	"local_on_all":  TablespaceSharedLocalOnAll,
}

// GetTablespaceSharedEnumValues Enumerates the set of values for TablespaceSharedEnum
func GetTablespaceSharedEnumValues() []TablespaceSharedEnum {
	values := make([]TablespaceSharedEnum, 0)
	for _, v := range mappingTablespaceSharedEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceSharedEnumStringValues Enumerates the set of values in String for TablespaceSharedEnum
func GetTablespaceSharedEnumStringValues() []string {
	return []string{
		"SHARED",
		"LOCAL_ON_LEAF",
		"LOCAL_ON_ALL",
	}
}

// GetMappingTablespaceSharedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceSharedEnum(val string) (TablespaceSharedEnum, bool) {
	enum, ok := mappingTablespaceSharedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultIndexCompressionEnum Enum with underlying type: string
type TablespaceDefaultIndexCompressionEnum string

// Set of constants representing the allowable values for TablespaceDefaultIndexCompressionEnum
const (
	TablespaceDefaultIndexCompressionEnabled  TablespaceDefaultIndexCompressionEnum = "ENABLED"
	TablespaceDefaultIndexCompressionDisabled TablespaceDefaultIndexCompressionEnum = "DISABLED"
)

var mappingTablespaceDefaultIndexCompressionEnum = map[string]TablespaceDefaultIndexCompressionEnum{
	"ENABLED":  TablespaceDefaultIndexCompressionEnabled,
	"DISABLED": TablespaceDefaultIndexCompressionDisabled,
}

var mappingTablespaceDefaultIndexCompressionEnumLowerCase = map[string]TablespaceDefaultIndexCompressionEnum{
	"enabled":  TablespaceDefaultIndexCompressionEnabled,
	"disabled": TablespaceDefaultIndexCompressionDisabled,
}

// GetTablespaceDefaultIndexCompressionEnumValues Enumerates the set of values for TablespaceDefaultIndexCompressionEnum
func GetTablespaceDefaultIndexCompressionEnumValues() []TablespaceDefaultIndexCompressionEnum {
	values := make([]TablespaceDefaultIndexCompressionEnum, 0)
	for _, v := range mappingTablespaceDefaultIndexCompressionEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultIndexCompressionEnumStringValues Enumerates the set of values in String for TablespaceDefaultIndexCompressionEnum
func GetTablespaceDefaultIndexCompressionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTablespaceDefaultIndexCompressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultIndexCompressionEnum(val string) (TablespaceDefaultIndexCompressionEnum, bool) {
	enum, ok := mappingTablespaceDefaultIndexCompressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceIndexCompressForEnum Enum with underlying type: string
type TablespaceIndexCompressForEnum string

// Set of constants representing the allowable values for TablespaceIndexCompressForEnum
const (
	TablespaceIndexCompressForAdvancedLow  TablespaceIndexCompressForEnum = "ADVANCED_LOW"
	TablespaceIndexCompressForAdvancedHigh TablespaceIndexCompressForEnum = "ADVANCED_HIGH"
	TablespaceIndexCompressForNone         TablespaceIndexCompressForEnum = "NONE"
)

var mappingTablespaceIndexCompressForEnum = map[string]TablespaceIndexCompressForEnum{
	"ADVANCED_LOW":  TablespaceIndexCompressForAdvancedLow,
	"ADVANCED_HIGH": TablespaceIndexCompressForAdvancedHigh,
	"NONE":          TablespaceIndexCompressForNone,
}

var mappingTablespaceIndexCompressForEnumLowerCase = map[string]TablespaceIndexCompressForEnum{
	"advanced_low":  TablespaceIndexCompressForAdvancedLow,
	"advanced_high": TablespaceIndexCompressForAdvancedHigh,
	"none":          TablespaceIndexCompressForNone,
}

// GetTablespaceIndexCompressForEnumValues Enumerates the set of values for TablespaceIndexCompressForEnum
func GetTablespaceIndexCompressForEnumValues() []TablespaceIndexCompressForEnum {
	values := make([]TablespaceIndexCompressForEnum, 0)
	for _, v := range mappingTablespaceIndexCompressForEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceIndexCompressForEnumStringValues Enumerates the set of values in String for TablespaceIndexCompressForEnum
func GetTablespaceIndexCompressForEnumStringValues() []string {
	return []string{
		"ADVANCED_LOW",
		"ADVANCED_HIGH",
		"NONE",
	}
}

// GetMappingTablespaceIndexCompressForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceIndexCompressForEnum(val string) (TablespaceIndexCompressForEnum, bool) {
	enum, ok := mappingTablespaceIndexCompressForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceDefaultInMemoryServiceEnum Enum with underlying type: string
type TablespaceDefaultInMemoryServiceEnum string

// Set of constants representing the allowable values for TablespaceDefaultInMemoryServiceEnum
const (
	TablespaceDefaultInMemoryServiceDefault     TablespaceDefaultInMemoryServiceEnum = "DEFAULT"
	TablespaceDefaultInMemoryServiceNone        TablespaceDefaultInMemoryServiceEnum = "NONE"
	TablespaceDefaultInMemoryServiceAll         TablespaceDefaultInMemoryServiceEnum = "ALL"
	TablespaceDefaultInMemoryServiceUserDefined TablespaceDefaultInMemoryServiceEnum = "USER_DEFINED"
)

var mappingTablespaceDefaultInMemoryServiceEnum = map[string]TablespaceDefaultInMemoryServiceEnum{
	"DEFAULT":      TablespaceDefaultInMemoryServiceDefault,
	"NONE":         TablespaceDefaultInMemoryServiceNone,
	"ALL":          TablespaceDefaultInMemoryServiceAll,
	"USER_DEFINED": TablespaceDefaultInMemoryServiceUserDefined,
}

var mappingTablespaceDefaultInMemoryServiceEnumLowerCase = map[string]TablespaceDefaultInMemoryServiceEnum{
	"default":      TablespaceDefaultInMemoryServiceDefault,
	"none":         TablespaceDefaultInMemoryServiceNone,
	"all":          TablespaceDefaultInMemoryServiceAll,
	"user_defined": TablespaceDefaultInMemoryServiceUserDefined,
}

// GetTablespaceDefaultInMemoryServiceEnumValues Enumerates the set of values for TablespaceDefaultInMemoryServiceEnum
func GetTablespaceDefaultInMemoryServiceEnumValues() []TablespaceDefaultInMemoryServiceEnum {
	values := make([]TablespaceDefaultInMemoryServiceEnum, 0)
	for _, v := range mappingTablespaceDefaultInMemoryServiceEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceDefaultInMemoryServiceEnumStringValues Enumerates the set of values in String for TablespaceDefaultInMemoryServiceEnum
func GetTablespaceDefaultInMemoryServiceEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"NONE",
		"ALL",
		"USER_DEFINED",
	}
}

// GetMappingTablespaceDefaultInMemoryServiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceDefaultInMemoryServiceEnum(val string) (TablespaceDefaultInMemoryServiceEnum, bool) {
	enum, ok := mappingTablespaceDefaultInMemoryServiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TablespaceLostWriteProtectEnum Enum with underlying type: string
type TablespaceLostWriteProtectEnum string

// Set of constants representing the allowable values for TablespaceLostWriteProtectEnum
const (
	TablespaceLostWriteProtectEnabled    TablespaceLostWriteProtectEnum = "ENABLED"
	TablespaceLostWriteProtectProtectOff TablespaceLostWriteProtectEnum = "PROTECT_OFF"
	TablespaceLostWriteProtectSuspend    TablespaceLostWriteProtectEnum = "SUSPEND"
)

var mappingTablespaceLostWriteProtectEnum = map[string]TablespaceLostWriteProtectEnum{
	"ENABLED":     TablespaceLostWriteProtectEnabled,
	"PROTECT_OFF": TablespaceLostWriteProtectProtectOff,
	"SUSPEND":     TablespaceLostWriteProtectSuspend,
}

var mappingTablespaceLostWriteProtectEnumLowerCase = map[string]TablespaceLostWriteProtectEnum{
	"enabled":     TablespaceLostWriteProtectEnabled,
	"protect_off": TablespaceLostWriteProtectProtectOff,
	"suspend":     TablespaceLostWriteProtectSuspend,
}

// GetTablespaceLostWriteProtectEnumValues Enumerates the set of values for TablespaceLostWriteProtectEnum
func GetTablespaceLostWriteProtectEnumValues() []TablespaceLostWriteProtectEnum {
	values := make([]TablespaceLostWriteProtectEnum, 0)
	for _, v := range mappingTablespaceLostWriteProtectEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceLostWriteProtectEnumStringValues Enumerates the set of values in String for TablespaceLostWriteProtectEnum
func GetTablespaceLostWriteProtectEnumStringValues() []string {
	return []string{
		"ENABLED",
		"PROTECT_OFF",
		"SUSPEND",
	}
}

// GetMappingTablespaceLostWriteProtectEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceLostWriteProtectEnum(val string) (TablespaceLostWriteProtectEnum, bool) {
	enum, ok := mappingTablespaceLostWriteProtectEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
