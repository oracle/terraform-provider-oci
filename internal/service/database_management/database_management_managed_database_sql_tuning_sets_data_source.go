// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseSqlTuningSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseSqlTuningSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_tuning_set_collection": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"scheduled_job_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"statement_counts": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readDatabaseManagementManagedDatabaseSqlTuningSets(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningSetsResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningSetsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		tmp := owner.(string)
		request.Owner = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlTuningSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningSetsDataSource-", DatabaseManagementManagedDatabaseSqlTuningSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseSqlTuningSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningSetSummaryToMap(item))
	}
	managedDatabaseSqlTuningSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseSqlTuningSetsDataSource().Schema["sql_tuning_set_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseSqlTuningSet["items"] = items
	}

	resources = append(resources, managedDatabaseSqlTuningSet)
	if err := s.D.Set("sql_tuning_set_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlInSqlTuningSetToMap(obj oci_database_management.SqlInSqlTuningSet) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContainerDatabaseId != nil {
		result["container_database_id"] = strconv.FormatInt(*obj.ContainerDatabaseId, 10)
	}

	metrics := []interface{}{}
	for _, item := range obj.Metrics {
		metrics = append(metrics, SqlMetricsToMap(item))
	}
	result["metrics"] = metrics

	if obj.Module != nil {
		result["module"] = string(*obj.Module)
	}

	if obj.PlanHashValue != nil {
		result["plan_hash_value"] = strconv.FormatInt(*obj.PlanHashValue, 10)
	}

	if obj.Schema != nil {
		result["schema"] = string(*obj.Schema)
	}

	if obj.SqlId != nil {
		result["sql_id"] = string(*obj.SqlId)
	}

	if obj.SqlText != nil {
		result["sql_text"] = string(*obj.SqlText)
	}

	return result
}

func SqlMetricsToMap(obj oci_database_management.SqlMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BufferGets != nil {
		result["buffer_gets"] = strconv.FormatInt(*obj.BufferGets, 10)
	}

	if obj.CpuTime != nil {
		result["cpu_time"] = strconv.FormatInt(*obj.CpuTime, 10)
	}

	if obj.DirectWrites != nil {
		result["direct_writes"] = strconv.FormatInt(*obj.DirectWrites, 10)
	}

	if obj.DiskReads != nil {
		result["disk_reads"] = strconv.FormatInt(*obj.DiskReads, 10)
	}

	if obj.ElapsedTime != nil {
		result["elapsed_time"] = strconv.FormatInt(*obj.ElapsedTime, 10)
	}

	if obj.Executions != nil {
		result["executions"] = strconv.FormatInt(*obj.Executions, 10)
	}

	return result
}

func SqlTuningSetAdminCredentialDetailsToMap(obj *oci_database_management.SqlTuningSetAdminCredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_management.SqlTuningSetAdminPasswordCredentialDetails:
		result["sql_tuning_set_admin_credential_type"] = "PASSWORD"

		if v.Password != nil {
			result["password"] = string(*v.Password)
		}
	case oci_database_management.SqlTuningSetAdminSecretCredentialDetails:
		result["sql_tuning_set_admin_credential_type"] = "SECRET"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'sql_tuning_set_admin_credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SqlTuningSetSummaryToMap(obj oci_database_management.SqlTuningSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	if obj.Id != nil {
		result["id"] = int(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.ScheduledJobName != nil {
		result["scheduled_job_name"] = string(*obj.ScheduledJobName)
	}

	if obj.StatementCounts != nil {
		result["statement_counts"] = int(*obj.StatementCounts)
	}

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}
