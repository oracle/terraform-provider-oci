// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbSystemOsPatchHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseDbSystemOsPatchHistoryEntriesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_os_patch_history_entry_collection": {
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
									"action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_patch_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"db_node_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_reboot_required": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"rpms": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
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

func readDatabaseDbSystemOsPatchHistoryEntriesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDbSystemOsPatchHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseDbSystemOsPatchHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemOsPatchHistoryEntriesResponse
}

func (s *DatabaseDbSystemOsPatchHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemOsPatchHistoryEntriesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.ListDbSystemOsPatchHistoryEntriesRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_database.DbSystemOsPatchHistoryEntrySummaryActionEnum(action.(string))
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemOsPatchHistoryEntries(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystemOsPatchHistoryEntries(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbSystemOsPatchHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemOsPatchHistoryEntriesDataSource-", DatabaseDbSystemOsPatchHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dbSystemOsPatchHistoryEntry := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DbSystemOsPatchHistoryEntrySummaryToMap(item))
	}
	dbSystemOsPatchHistoryEntry["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseDbSystemOsPatchHistoryEntriesDataSource().Schema["db_system_os_patch_history_entry_collection"].Elem.(*schema.Resource).Schema)
		dbSystemOsPatchHistoryEntry["items"] = items
	}

	resources = append(resources, dbSystemOsPatchHistoryEntry)
	if err := s.D.Set("db_system_os_patch_history_entry_collection", resources); err != nil {
		return err
	}

	return nil
}

func DbSystemOsPatchDetailsCollectionToMap(obj *oci_database.DbSystemOsPatchDetailsCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DbSystemOsPatchDetailsSummaryToMap(item))
	}
	result["items"] = items

	return result
}

func DbSystemOsPatchDetailsSummaryToMap(obj oci_database.DbSystemOsPatchDetailsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbNodeId != nil {
		result["db_node_id"] = string(*obj.DbNodeId)
	}

	if obj.IsRebootRequired != nil {
		result["is_reboot_required"] = bool(*obj.IsRebootRequired)
	}

	result["rpms"] = obj.Rpms

	return result
}

func DbSystemOsPatchHistoryEntrySummaryToMap(obj oci_database.DbSystemOsPatchHistoryEntrySummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OsPatchDetails != nil {
		result["os_patch_details"] = []interface{}{DbSystemOsPatchDetailsCollectionToMap(obj.OsPatchDetails)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
