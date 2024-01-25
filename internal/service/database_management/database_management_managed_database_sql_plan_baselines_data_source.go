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

func DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseSqlPlanBaselines,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_accepted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_adaptive": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_auto_purged": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_fixed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_never_executed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_reproduced": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"plan_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_handle": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sql_plan_baseline_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"accepted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"adaptive": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"auto_purge": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"execution_plan": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fixed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"module": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"origin": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"plan_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reproduced": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_handle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_executed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_modified": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabaseSqlPlanBaselines(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListSqlPlanBaselinesResponse
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlPlanBaselinesRequest{}

	if isAccepted, ok := s.D.GetOkExists("is_accepted"); ok {
		tmp := isAccepted.(bool)
		request.IsAccepted = &tmp
	}

	if isAdaptive, ok := s.D.GetOkExists("is_adaptive"); ok {
		tmp := isAdaptive.(bool)
		request.IsAdaptive = &tmp
	}

	if isAutoPurged, ok := s.D.GetOkExists("is_auto_purged"); ok {
		tmp := isAutoPurged.(bool)
		request.IsAutoPurged = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if isFixed, ok := s.D.GetOkExists("is_fixed"); ok {
		tmp := isFixed.(bool)
		request.IsFixed = &tmp
	}

	if isNeverExecuted, ok := s.D.GetOkExists("is_never_executed"); ok {
		tmp := isNeverExecuted.(bool)
		request.IsNeverExecuted = &tmp
	}

	if isReproduced, ok := s.D.GetOkExists("is_reproduced"); ok {
		tmp := isReproduced.(bool)
		request.IsReproduced = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	if origin, ok := s.D.GetOkExists("origin"); ok {
		request.Origin = oci_database_management.ListSqlPlanBaselinesOriginEnum(origin.(string))
	}

	if planName, ok := s.D.GetOkExists("plan_name"); ok {
		tmp := planName.(string)
		request.PlanName = &tmp
	}

	if sqlHandle, ok := s.D.GetOkExists("sql_handle"); ok {
		tmp := sqlHandle.(string)
		request.SqlHandle = &tmp
	}

	if sqlText, ok := s.D.GetOkExists("sql_text"); ok {
		tmp := sqlText.(string)
		request.SqlText = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlPlanBaselines(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlPlanBaselines(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource-", DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseSqlPlanBaseline := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlPlanBaselineSummaryToMap(item))
	}
	managedDatabaseSqlPlanBaseline["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource().Schema["sql_plan_baseline_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseSqlPlanBaseline["items"] = items
	}

	resources = append(resources, managedDatabaseSqlPlanBaseline)
	if err := s.D.Set("sql_plan_baseline_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlPlanBaselineSummaryToMap(obj oci_database_management.SqlPlanBaselineSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["accepted"] = string(*obj.Accepted)

	result["adaptive"] = string(*obj.Adaptive)

	result["auto_purge"] = string(*obj.AutoPurge)

	result["enabled"] = string(*obj.Enabled)

	result["fixed"] = string(*obj.Fixed)

	result["origin"] = string(obj.Origin)

	if obj.PlanName != nil {
		result["plan_name"] = string(*obj.PlanName)
	}

	result["reproduced"] = string(*obj.Reproduced)

	if obj.SqlHandle != nil {
		result["sql_handle"] = string(*obj.SqlHandle)
	}

	if obj.SqlText != nil {
		result["sql_text"] = string(*obj.SqlText)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastExecuted != nil {
		result["time_last_executed"] = obj.TimeLastExecuted.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}
