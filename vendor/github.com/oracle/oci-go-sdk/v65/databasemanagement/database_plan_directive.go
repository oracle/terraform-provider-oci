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

// DatabasePlanDirective Manages resource allocation among databases. Besides the name, at least one other property must be available.
type DatabasePlanDirective struct {

	// The name of a database or a profile.
	Name *string `mandatory:"true" json:"name"`

	// The relative priority of a database in the database plan. A higher share value implies
	// higher priority and more access to the I/O resources.
	// Use either share or (level, allocation). All plan directives in a database plan
	// should use the same setting.
	// Share-based resource allocation is the recommended method for a database plan.
	Share *int `mandatory:"false" json:"share"`

	// The allocation level. Valid values are from 1 to 8. Resources are allocated to level 1 first,
	// and then remaining resources are allocated to level 2, and so on.
	Level *int `mandatory:"false" json:"level"`

	// The resource allocation as a percentage (0-100) within the level.
	Allocation *int `mandatory:"false" json:"allocation"`

	// The maximum I/O utilization limit as a percentage of the available resources.
	Limit *int `mandatory:"false" json:"limit"`

	// Controls use of Exadata Smart Flash Cache by a database.
	// This ensures that cache space is reserved for mission-critical databases.
	// flashcache=off is invalid in a directive that contains the flashcachemin, flashcachelimit, or flashcachesize attributes.
	IsFlashCacheOn *bool `mandatory:"false" json:"isFlashCacheOn"`

	// Controls use of the persistent memory (PMEM) cache by a database. This ensures that cache space
	// is reserved for mission-critical databases.
	// pmemcache=off is invalid in a directive that contains the pmemcachemin, pmemcachelimit, or pmemcachesize attributes.
	IsPmemCacheOn *bool `mandatory:"false" json:"isPmemCacheOn"`

	// Controls use of Exadata Smart Flash Log by a database.
	// This ensures that Exadata Smart Flash Log is reserved for mission-critical databases.
	IsFlashLogOn *bool `mandatory:"false" json:"isFlashLogOn"`

	// Controls use of persistent memory logging (PMEM log) by a database.
	// This ensures that PMEM log is reserved for mission-critical databases.
	IsPmemLogOn *bool `mandatory:"false" json:"isPmemLogOn"`

	// Defines a soft limit for space usage in Exadata Smart Flash Cache.
	// If the cache is not full, the limit can be exceeded.
	// You specify the value for flashcachelimit in bytes. You can also use the suffixes M (megabytes),
	// G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T.
	// The value for flashcachelimit must be at least 4 MB.
	// The flashcachelimit and flashcachesize attributes cannot be specified in the same directive.
	// The value for flashcachelimit cannot be smaller than flashcachemin, if it is specified.
	FlashCacheLimit *string `mandatory:"false" json:"flashCacheLimit"`

	// Specifies a minimum guaranteed space allocation in Exadata Smart Flash Cache.
	// You specify the value for flashcachemin in bytes. You can also use the suffixes
	// M (megabytes), G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T.
	// The value for flashcachemin must be at least 4 MB.
	// In any plan, the sum of all flashcachemin values cannot exceed the size of Exadata Smart Flash Cache.
	// If flashcachelimit is specified, then the value for flashcachemin cannot exceed flashcachelimit.
	// If flashcachesize is specified, then the value for flashcachemin cannot exceed flashcachesize.
	FlashCacheMin *string `mandatory:"false" json:"flashCacheMin"`

	// Defines a hard limit for space usage in Exadata Smart Flash Cache.
	// The limit cannot be exceeded, even if the cache is not full.
	// In an IORM plan, if the size of Exadata Smart Flash Cache can accommodate all of the flashcachemin
	// and flashcachesize allocations, then each flashcachesize definition represents a guaranteed space allocation.
	// However, starting with Oracle Exadata System Software release 19.2.0 you can use the flashcachesize
	// attribute to over-provision space in Exadata Smart Flash Cache. Consequently,
	// if the size of Exadata Smart Flash Cache cannot accommodate all of the flashcachemin and flashcachesize
	// allocations, then only flashcachemin is guaranteed.
	FlashCacheSize *string `mandatory:"false" json:"flashCacheSize"`

	// Defines a soft limit for space usage in the persistent memory (PMEM) cache.
	// If the cache is not full, the limit can be exceeded.
	// You specify the value for pmemcachelimit in bytes. You can also use the suffixes M (megabytes),
	// G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T.
	// The value for pmemcachelimit must be at least 4 MB.
	// The pmemcachelimit and pmemcachesize attributes cannot be specified in the same directive.
	// The value for pmemcachelimit cannot be smaller than pmemcachemin, if it is specified.
	PmemCacheLimit *string `mandatory:"false" json:"pmemCacheLimit"`

	// Specifies a minimum guaranteed space allocation in the persistent memory (PMEM) cache.
	PmemCacheMin *string `mandatory:"false" json:"pmemCacheMin"`

	// Defines a hard limit for space usage in the persistent memory (PMEM) cache.
	// The limit cannot be exceeded, even if the cache is not full.
	// In an IORM plan, if the size of the PMEM cache can accommodate all of the pmemcachemin and
	// pmemcachesize allocations, then each pmemcachesize definition represents a guaranteed space allocation.
	// However, you can use the pmemcachesize attribute to over-provision space in the PMEM cache.
	// Consequently, if the PMEM cache size cannot accommodate all of the pmemcachemin and pmemcachesize
	// allocations, then only pmemcachemin is guaranteed.
	PmemCacheSize *string `mandatory:"false" json:"pmemCacheSize"`

	// Starting with Oracle Exadata System Software release 19.1.0, you can use the asmcluster attribute to
	// distinguish between databases with the same name running in different Oracle ASM clusters.
	AsmCluster *string `mandatory:"false" json:"asmCluster"`

	// Enables you to create a profile or template, to ease management and configuration of resource plans
	// in environments with many databases.
	// - type=database: Specifies a directive that applies to a specific database.
	// If type in not specified, then the directive defaults to the database type.
	// - type=profile: Specifies a directive that applies to a profile rather than a specific database.
	//   To associate a database with an IORM profile, you must set the database initialization
	//   parameter db_performance_profile to the value of the profile name. Databases that map to a profile inherit the settings specified in the profile.
	Type DatabasePlanTypeEnumEnum `mandatory:"false" json:"type,omitempty"`

	// Enables you to specify different plan directives based on the Oracle Data Guard database role.
	Role DatabasePlanRoleEnumEnum `mandatory:"false" json:"role,omitempty"`
}

func (m DatabasePlanDirective) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabasePlanDirective) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabasePlanTypeEnumEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDatabasePlanTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabasePlanRoleEnumEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetDatabasePlanRoleEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
