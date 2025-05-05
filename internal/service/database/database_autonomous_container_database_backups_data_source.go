// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabaseBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_remote": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_container_database_backup_collection": {
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
									"acd_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_container_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_databases": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"infrastructure_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_automatic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_remote_backup": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"retention_period_in_days": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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

func readDatabaseAutonomousContainerDatabaseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousContainerDatabaseBackupsResponse
}

func (s *DatabaseAutonomousContainerDatabaseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousContainerDatabaseBackupsRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if infrastructureType, ok := s.D.GetOkExists("infrastructure_type"); ok {
		request.InfrastructureType = oci_database.AutonomousContainerDatabaseBackupSummaryInfrastructureTypeEnum(infrastructureType.(string))
	}

	if isRemote, ok := s.D.GetOkExists("is_remote"); ok {
		tmp := isRemote.(bool)
		request.IsRemote = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousContainerDatabaseBackupSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousContainerDatabaseBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousContainerDatabaseBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousContainerDatabaseBackupsDataSource-", DatabaseAutonomousContainerDatabaseBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousContainerDatabaseBackup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AutonomousContainerDatabaseBackupSummaryToMap(item))
	}
	autonomousContainerDatabaseBackup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousContainerDatabaseBackupsDataSource().Schema["autonomous_container_database_backup_collection"].Elem.(*schema.Resource).Schema)
		autonomousContainerDatabaseBackup["items"] = items
	}

	resources = append(resources, autonomousContainerDatabaseBackup)
	if err := s.D.Set("autonomous_container_database_backup_collection", resources); err != nil {
		return err
	}

	return nil
}

func AutonomousContainerDatabaseBackupSummaryToMap(obj oci_database.AutonomousContainerDatabaseBackupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AcdDisplayName != nil {
		result["acd_display_name"] = string(*obj.AcdDisplayName)
	}

	if obj.AutonomousContainerDatabaseId != nil {
		result["autonomous_container_database_id"] = string(*obj.AutonomousContainerDatabaseId)
	}

	autonomousDatabases := []interface{}{}
	for _, item := range obj.AutonomousDatabases {
		autonomousDatabases = append(autonomousDatabases, AutonomousDatabaseInBackupToMap(item))
	}
	result["autonomous_databases"] = autonomousDatabases

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["infrastructure_type"] = string(obj.InfrastructureType)

	if obj.IsAutomatic != nil {
		result["is_automatic"] = bool(*obj.IsAutomatic)
	}

	if obj.IsRemoteBackup != nil {
		result["is_remote_backup"] = bool(*obj.IsRemoteBackup)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.RetentionPeriodInDays != nil {
		result["retention_period_in_days"] = int(*obj.RetentionPeriodInDays)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func AutonomousDatabaseInBackupToMap(obj oci_database.AutonomousDatabaseInBackup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}
