// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
	"strings"
)

// ActionableInsightsContentTypesResourceEnum Enum with underlying type: string
type ActionableInsightsContentTypesResourceEnum string

// Set of constants representing the allowable values for ActionableInsightsContentTypesResourceEnum
const (
	ActionableInsightsContentTypesResourceNewHighs                                               ActionableInsightsContentTypesResourceEnum = "NEW_HIGHS"
	ActionableInsightsContentTypesResourceBigChanges                                             ActionableInsightsContentTypesResourceEnum = "BIG_CHANGES"
	ActionableInsightsContentTypesResourceCurrentInventory                                       ActionableInsightsContentTypesResourceEnum = "CURRENT_INVENTORY"
	ActionableInsightsContentTypesResourceInventoryChanges                                       ActionableInsightsContentTypesResourceEnum = "INVENTORY_CHANGES"
	ActionableInsightsContentTypesResourceFleetStatistics                                        ActionableInsightsContentTypesResourceEnum = "FLEET_STATISTICS"
	ActionableInsightsContentTypesResourceFleetAnalysisSummaryDbCount                            ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_SUMMARY_DB_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisSummarySqlAnalyzedCount                   ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_SUMMARY_SQL_ANALYZED_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisSummaryNewSqlCount                        ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_SUMMARY_NEW_SQL_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisSummaryBusiestDb                          ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_SUMMARY_BUSIEST_DB"
	ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlCount                         ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_DEGRADING_SQL_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlByDb                          ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_DEGRADING_SQL_BY_DB"
	ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlBySqlId                       ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_DEGRADING_SQL_BY_SQL_ID"
	ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesCount                          ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_PLAN_CHANGES_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesDbMostChanges                  ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_PLAN_CHANGES_DB_MOST_CHANGES"
	ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdImproved                ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_IMPROVED"
	ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdDegraded                ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_DEGRADED"
	ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsCount                   ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_INVALIDATION_STORMS_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsHighest                 ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_INVALIDATION_STORMS_HIGHEST"
	ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesCount                  ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_COUNT"
	ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesByDb                   ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_DB"
	ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesBySql                  ActionableInsightsContentTypesResourceEnum = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_SQL"
	ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDbCount                   ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SUMMARY_DB_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlAnalyzedCount          ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SUMMARY_SQL_ANALYZED_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlPerformanceTrendsCount ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SUMMARY_SQL_PERFORMANCE_TRENDS_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDegradedSqlCount          ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SUMMARY_DEGRADED_SQL_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationSummaryImprovedSqlCount          ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SUMMARY_IMPROVED_SQL_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationDbDegradedCount                  ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_DB_DEGRADED_COUNT"
	ActionableInsightsContentTypesResourcePerformanceDegradationSqlDegradedTable                 ActionableInsightsContentTypesResourceEnum = "PERFORMANCE_DEGRADATION_SQL_DEGRADED_TABLE"
	ActionableInsightsContentTypesResourcePlanChangesSummaryDbCount                              ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_SUMMARY_DB_COUNT"
	ActionableInsightsContentTypesResourcePlanChangesSummarySqlAnalyzedCount                     ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_SUMMARY_SQL_ANALYZED_COUNT"
	ActionableInsightsContentTypesResourcePlanChangesSummaryPlanChangesCount                     ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_SUMMARY_PLAN_CHANGES_COUNT"
	ActionableInsightsContentTypesResourcePlanChangesSummaryImprovementsCount                    ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_SUMMARY_IMPROVEMENTS_COUNT"
	ActionableInsightsContentTypesResourcePlanChangesSummaryDegradationCount                     ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_SUMMARY_DEGRADATION_COUNT"
	ActionableInsightsContentTypesResourcePlanChangesTopPlanChangesTable                         ActionableInsightsContentTypesResourceEnum = "PLAN_CHANGES_TOP_PLAN_CHANGES_TABLE"
	ActionableInsightsContentTypesResourceTopDbSummaryDbCount                                    ActionableInsightsContentTypesResourceEnum = "TOP_DB_SUMMARY_DB_COUNT"
	ActionableInsightsContentTypesResourceTopDbSummarySqlAnalyzedCount                           ActionableInsightsContentTypesResourceEnum = "TOP_DB_SUMMARY_SQL_ANALYZED_COUNT"
	ActionableInsightsContentTypesResourceTopDbSummaryBusiestDb                                  ActionableInsightsContentTypesResourceEnum = "TOP_DB_SUMMARY_BUSIEST_DB"
	ActionableInsightsContentTypesResourceTopTable                                               ActionableInsightsContentTypesResourceEnum = "TOP_TABLE"
	ActionableInsightsContentTypesResourceCollectionDelayCount                                   ActionableInsightsContentTypesResourceEnum = "COLLECTION_DELAY_COUNT"
	ActionableInsightsContentTypesResourceCollectionDelayPreviousWeekCount                       ActionableInsightsContentTypesResourceEnum = "COLLECTION_DELAY_PREVIOUS_WEEK_COUNT"
)

var mappingActionableInsightsContentTypesResourceEnum = map[string]ActionableInsightsContentTypesResourceEnum{
	"NEW_HIGHS":                                                    ActionableInsightsContentTypesResourceNewHighs,
	"BIG_CHANGES":                                                  ActionableInsightsContentTypesResourceBigChanges,
	"CURRENT_INVENTORY":                                            ActionableInsightsContentTypesResourceCurrentInventory,
	"INVENTORY_CHANGES":                                            ActionableInsightsContentTypesResourceInventoryChanges,
	"FLEET_STATISTICS":                                             ActionableInsightsContentTypesResourceFleetStatistics,
	"FLEET_ANALYSIS_SUMMARY_DB_COUNT":                              ActionableInsightsContentTypesResourceFleetAnalysisSummaryDbCount,
	"FLEET_ANALYSIS_SUMMARY_SQL_ANALYZED_COUNT":                    ActionableInsightsContentTypesResourceFleetAnalysisSummarySqlAnalyzedCount,
	"FLEET_ANALYSIS_SUMMARY_NEW_SQL_COUNT":                         ActionableInsightsContentTypesResourceFleetAnalysisSummaryNewSqlCount,
	"FLEET_ANALYSIS_SUMMARY_BUSIEST_DB":                            ActionableInsightsContentTypesResourceFleetAnalysisSummaryBusiestDb,
	"FLEET_ANALYSIS_DEGRADING_SQL_COUNT":                           ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlCount,
	"FLEET_ANALYSIS_DEGRADING_SQL_BY_DB":                           ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlByDb,
	"FLEET_ANALYSIS_DEGRADING_SQL_BY_SQL_ID":                       ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlBySqlId,
	"FLEET_ANALYSIS_PLAN_CHANGES_COUNT":                            ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesCount,
	"FLEET_ANALYSIS_PLAN_CHANGES_DB_MOST_CHANGES":                  ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesDbMostChanges,
	"FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_IMPROVED":               ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdImproved,
	"FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_DEGRADED":               ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdDegraded,
	"FLEET_ANALYSIS_INVALIDATION_STORMS_COUNT":                     ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsCount,
	"FLEET_ANALYSIS_INVALIDATION_STORMS_HIGHEST":                   ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsHighest,
	"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_COUNT":                   ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesCount,
	"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_DB":                   ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesByDb,
	"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_SQL":                  ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesBySql,
	"PERFORMANCE_DEGRADATION_SUMMARY_DB_COUNT":                     ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDbCount,
	"PERFORMANCE_DEGRADATION_SUMMARY_SQL_ANALYZED_COUNT":           ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlAnalyzedCount,
	"PERFORMANCE_DEGRADATION_SUMMARY_SQL_PERFORMANCE_TRENDS_COUNT": ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlPerformanceTrendsCount,
	"PERFORMANCE_DEGRADATION_SUMMARY_DEGRADED_SQL_COUNT":           ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDegradedSqlCount,
	"PERFORMANCE_DEGRADATION_SUMMARY_IMPROVED_SQL_COUNT":           ActionableInsightsContentTypesResourcePerformanceDegradationSummaryImprovedSqlCount,
	"PERFORMANCE_DEGRADATION_DB_DEGRADED_COUNT":                    ActionableInsightsContentTypesResourcePerformanceDegradationDbDegradedCount,
	"PERFORMANCE_DEGRADATION_SQL_DEGRADED_TABLE":                   ActionableInsightsContentTypesResourcePerformanceDegradationSqlDegradedTable,
	"PLAN_CHANGES_SUMMARY_DB_COUNT":                                ActionableInsightsContentTypesResourcePlanChangesSummaryDbCount,
	"PLAN_CHANGES_SUMMARY_SQL_ANALYZED_COUNT":                      ActionableInsightsContentTypesResourcePlanChangesSummarySqlAnalyzedCount,
	"PLAN_CHANGES_SUMMARY_PLAN_CHANGES_COUNT":                      ActionableInsightsContentTypesResourcePlanChangesSummaryPlanChangesCount,
	"PLAN_CHANGES_SUMMARY_IMPROVEMENTS_COUNT":                      ActionableInsightsContentTypesResourcePlanChangesSummaryImprovementsCount,
	"PLAN_CHANGES_SUMMARY_DEGRADATION_COUNT":                       ActionableInsightsContentTypesResourcePlanChangesSummaryDegradationCount,
	"PLAN_CHANGES_TOP_PLAN_CHANGES_TABLE":                          ActionableInsightsContentTypesResourcePlanChangesTopPlanChangesTable,
	"TOP_DB_SUMMARY_DB_COUNT":                                      ActionableInsightsContentTypesResourceTopDbSummaryDbCount,
	"TOP_DB_SUMMARY_SQL_ANALYZED_COUNT":                            ActionableInsightsContentTypesResourceTopDbSummarySqlAnalyzedCount,
	"TOP_DB_SUMMARY_BUSIEST_DB":                                    ActionableInsightsContentTypesResourceTopDbSummaryBusiestDb,
	"TOP_TABLE":                                                    ActionableInsightsContentTypesResourceTopTable,
	"COLLECTION_DELAY_COUNT":                                       ActionableInsightsContentTypesResourceCollectionDelayCount,
	"COLLECTION_DELAY_PREVIOUS_WEEK_COUNT":                         ActionableInsightsContentTypesResourceCollectionDelayPreviousWeekCount,
}

var mappingActionableInsightsContentTypesResourceEnumLowerCase = map[string]ActionableInsightsContentTypesResourceEnum{
	"new_highs":                                                    ActionableInsightsContentTypesResourceNewHighs,
	"big_changes":                                                  ActionableInsightsContentTypesResourceBigChanges,
	"current_inventory":                                            ActionableInsightsContentTypesResourceCurrentInventory,
	"inventory_changes":                                            ActionableInsightsContentTypesResourceInventoryChanges,
	"fleet_statistics":                                             ActionableInsightsContentTypesResourceFleetStatistics,
	"fleet_analysis_summary_db_count":                              ActionableInsightsContentTypesResourceFleetAnalysisSummaryDbCount,
	"fleet_analysis_summary_sql_analyzed_count":                    ActionableInsightsContentTypesResourceFleetAnalysisSummarySqlAnalyzedCount,
	"fleet_analysis_summary_new_sql_count":                         ActionableInsightsContentTypesResourceFleetAnalysisSummaryNewSqlCount,
	"fleet_analysis_summary_busiest_db":                            ActionableInsightsContentTypesResourceFleetAnalysisSummaryBusiestDb,
	"fleet_analysis_degrading_sql_count":                           ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlCount,
	"fleet_analysis_degrading_sql_by_db":                           ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlByDb,
	"fleet_analysis_degrading_sql_by_sql_id":                       ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlBySqlId,
	"fleet_analysis_plan_changes_count":                            ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesCount,
	"fleet_analysis_plan_changes_db_most_changes":                  ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesDbMostChanges,
	"fleet_analysis_plan_changes_by_sql_id_improved":               ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdImproved,
	"fleet_analysis_plan_changes_by_sql_id_degraded":               ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdDegraded,
	"fleet_analysis_invalidation_storms_count":                     ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsCount,
	"fleet_analysis_invalidation_storms_highest":                   ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsHighest,
	"fleet_analysis_cursor_sharing_issues_count":                   ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesCount,
	"fleet_analysis_cursor_sharing_issues_by_db":                   ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesByDb,
	"fleet_analysis_cursor_sharing_issues_by_sql":                  ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesBySql,
	"performance_degradation_summary_db_count":                     ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDbCount,
	"performance_degradation_summary_sql_analyzed_count":           ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlAnalyzedCount,
	"performance_degradation_summary_sql_performance_trends_count": ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlPerformanceTrendsCount,
	"performance_degradation_summary_degraded_sql_count":           ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDegradedSqlCount,
	"performance_degradation_summary_improved_sql_count":           ActionableInsightsContentTypesResourcePerformanceDegradationSummaryImprovedSqlCount,
	"performance_degradation_db_degraded_count":                    ActionableInsightsContentTypesResourcePerformanceDegradationDbDegradedCount,
	"performance_degradation_sql_degraded_table":                   ActionableInsightsContentTypesResourcePerformanceDegradationSqlDegradedTable,
	"plan_changes_summary_db_count":                                ActionableInsightsContentTypesResourcePlanChangesSummaryDbCount,
	"plan_changes_summary_sql_analyzed_count":                      ActionableInsightsContentTypesResourcePlanChangesSummarySqlAnalyzedCount,
	"plan_changes_summary_plan_changes_count":                      ActionableInsightsContentTypesResourcePlanChangesSummaryPlanChangesCount,
	"plan_changes_summary_improvements_count":                      ActionableInsightsContentTypesResourcePlanChangesSummaryImprovementsCount,
	"plan_changes_summary_degradation_count":                       ActionableInsightsContentTypesResourcePlanChangesSummaryDegradationCount,
	"plan_changes_top_plan_changes_table":                          ActionableInsightsContentTypesResourcePlanChangesTopPlanChangesTable,
	"top_db_summary_db_count":                                      ActionableInsightsContentTypesResourceTopDbSummaryDbCount,
	"top_db_summary_sql_analyzed_count":                            ActionableInsightsContentTypesResourceTopDbSummarySqlAnalyzedCount,
	"top_db_summary_busiest_db":                                    ActionableInsightsContentTypesResourceTopDbSummaryBusiestDb,
	"top_table":                                                    ActionableInsightsContentTypesResourceTopTable,
	"collection_delay_count":                                       ActionableInsightsContentTypesResourceCollectionDelayCount,
	"collection_delay_previous_week_count":                         ActionableInsightsContentTypesResourceCollectionDelayPreviousWeekCount,
}

// GetActionableInsightsContentTypesResourceEnumValues Enumerates the set of values for ActionableInsightsContentTypesResourceEnum
func GetActionableInsightsContentTypesResourceEnumValues() []ActionableInsightsContentTypesResourceEnum {
	values := make([]ActionableInsightsContentTypesResourceEnum, 0)
	for _, v := range mappingActionableInsightsContentTypesResourceEnum {
		values = append(values, v)
	}
	return values
}

// GetActionableInsightsContentTypesResourceEnumStringValues Enumerates the set of values in String for ActionableInsightsContentTypesResourceEnum
func GetActionableInsightsContentTypesResourceEnumStringValues() []string {
	return []string{
		"NEW_HIGHS",
		"BIG_CHANGES",
		"CURRENT_INVENTORY",
		"INVENTORY_CHANGES",
		"FLEET_STATISTICS",
		"FLEET_ANALYSIS_SUMMARY_DB_COUNT",
		"FLEET_ANALYSIS_SUMMARY_SQL_ANALYZED_COUNT",
		"FLEET_ANALYSIS_SUMMARY_NEW_SQL_COUNT",
		"FLEET_ANALYSIS_SUMMARY_BUSIEST_DB",
		"FLEET_ANALYSIS_DEGRADING_SQL_COUNT",
		"FLEET_ANALYSIS_DEGRADING_SQL_BY_DB",
		"FLEET_ANALYSIS_DEGRADING_SQL_BY_SQL_ID",
		"FLEET_ANALYSIS_PLAN_CHANGES_COUNT",
		"FLEET_ANALYSIS_PLAN_CHANGES_DB_MOST_CHANGES",
		"FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_IMPROVED",
		"FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_DEGRADED",
		"FLEET_ANALYSIS_INVALIDATION_STORMS_COUNT",
		"FLEET_ANALYSIS_INVALIDATION_STORMS_HIGHEST",
		"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_COUNT",
		"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_DB",
		"FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_SQL",
		"PERFORMANCE_DEGRADATION_SUMMARY_DB_COUNT",
		"PERFORMANCE_DEGRADATION_SUMMARY_SQL_ANALYZED_COUNT",
		"PERFORMANCE_DEGRADATION_SUMMARY_SQL_PERFORMANCE_TRENDS_COUNT",
		"PERFORMANCE_DEGRADATION_SUMMARY_DEGRADED_SQL_COUNT",
		"PERFORMANCE_DEGRADATION_SUMMARY_IMPROVED_SQL_COUNT",
		"PERFORMANCE_DEGRADATION_DB_DEGRADED_COUNT",
		"PERFORMANCE_DEGRADATION_SQL_DEGRADED_TABLE",
		"PLAN_CHANGES_SUMMARY_DB_COUNT",
		"PLAN_CHANGES_SUMMARY_SQL_ANALYZED_COUNT",
		"PLAN_CHANGES_SUMMARY_PLAN_CHANGES_COUNT",
		"PLAN_CHANGES_SUMMARY_IMPROVEMENTS_COUNT",
		"PLAN_CHANGES_SUMMARY_DEGRADATION_COUNT",
		"PLAN_CHANGES_TOP_PLAN_CHANGES_TABLE",
		"TOP_DB_SUMMARY_DB_COUNT",
		"TOP_DB_SUMMARY_SQL_ANALYZED_COUNT",
		"TOP_DB_SUMMARY_BUSIEST_DB",
		"TOP_TABLE",
		"COLLECTION_DELAY_COUNT",
		"COLLECTION_DELAY_PREVIOUS_WEEK_COUNT",
	}
}

// GetMappingActionableInsightsContentTypesResourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionableInsightsContentTypesResourceEnum(val string) (ActionableInsightsContentTypesResourceEnum, bool) {
	enum, ok := mappingActionableInsightsContentTypesResourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
