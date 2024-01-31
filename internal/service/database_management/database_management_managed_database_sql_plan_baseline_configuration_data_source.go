// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlPlanBaselineConfiguration,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"auto_capture_filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"modified_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"values_to_exclude": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"values_to_include": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"auto_spm_evolve_task_parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"allowed_time_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"alternate_plan_baselines": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"alternate_plan_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"alternate_plan_sources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"are_plans_auto_accepted": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"is_auto_spm_evolve_task_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_automatic_initial_plan_capture_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_high_frequency_auto_spm_evolve_task_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_sql_plan_baselines_usage_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"plan_retention_weeks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"space_budget_mb": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"space_budget_percent": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"space_used_mb": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlPlanBaselineConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetSqlPlanBaselineConfigurationResponse
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSourceCrud) Get() error {
	request := oci_database_management.GetSqlPlanBaselineConfigurationRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetSqlPlanBaselineConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSource-", DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSource(), s.D))

	autoCaptureFilters := []interface{}{}
	for _, item := range s.Res.AutoCaptureFilters {
		autoCaptureFilters = append(autoCaptureFilters, AutomaticCaptureFilterToMap(item))
	}
	s.D.Set("auto_capture_filters", autoCaptureFilters)

	if s.Res.AutoSpmEvolveTaskParameters != nil {
		s.D.Set("auto_spm_evolve_task_parameters", []interface{}{SpmEvolveTaskParametersToMap(s.Res.AutoSpmEvolveTaskParameters)})
	} else {
		s.D.Set("auto_spm_evolve_task_parameters", nil)
	}

	if s.Res.IsAutoSpmEvolveTaskEnabled != nil {
		s.D.Set("is_auto_spm_evolve_task_enabled", *s.Res.IsAutoSpmEvolveTaskEnabled)
	}

	if s.Res.IsAutomaticInitialPlanCaptureEnabled != nil {
		s.D.Set("is_automatic_initial_plan_capture_enabled", *s.Res.IsAutomaticInitialPlanCaptureEnabled)
	}

	if s.Res.IsHighFrequencyAutoSpmEvolveTaskEnabled != nil {
		s.D.Set("is_high_frequency_auto_spm_evolve_task_enabled", *s.Res.IsHighFrequencyAutoSpmEvolveTaskEnabled)
	}

	if s.Res.IsSqlPlanBaselinesUsageEnabled != nil {
		s.D.Set("is_sql_plan_baselines_usage_enabled", *s.Res.IsSqlPlanBaselinesUsageEnabled)
	}

	if s.Res.PlanRetentionWeeks != nil {
		s.D.Set("plan_retention_weeks", *s.Res.PlanRetentionWeeks)
	}

	if s.Res.SpaceBudgetMB != nil {
		s.D.Set("space_budget_mb", *s.Res.SpaceBudgetMB)
	}

	if s.Res.SpaceBudgetPercent != nil {
		s.D.Set("space_budget_percent", *s.Res.SpaceBudgetPercent)
	}

	if s.Res.SpaceUsedMB != nil {
		s.D.Set("space_used_mb", *s.Res.SpaceUsedMB)
	}

	return nil
}

func AutomaticCaptureFilterToMap(obj oci_database_management.AutomaticCaptureFilter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ModifiedBy != nil {
		result["modified_by"] = string(*obj.ModifiedBy)
	}

	result["name"] = string(obj.Name)

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	result["values_to_exclude"] = obj.ValuesToExclude

	result["values_to_include"] = obj.ValuesToInclude

	return result
}

func SpmEvolveTaskParametersToMap(obj *oci_database_management.SpmEvolveTaskParameters) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedTimeLimit != nil {
		result["allowed_time_limit"] = int(*obj.AllowedTimeLimit)
	}

	result["alternate_plan_baselines"] = obj.AlternatePlanBaselines

	if obj.AlternatePlanLimit != nil {
		result["alternate_plan_limit"] = int(*obj.AlternatePlanLimit)
	}

	result["alternate_plan_sources"] = obj.AlternatePlanSources

	if obj.ArePlansAutoAccepted != nil {
		result["are_plans_auto_accepted"] = bool(*obj.ArePlansAutoAccepted)
	}

	return result
}
