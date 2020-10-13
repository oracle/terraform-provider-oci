// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v27/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_databases_clones", DatabaseAutonomousDatabasesClonesDataSource())
}

func DatabaseAutonomousDatabasesClonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabasesClones,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"clone_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_upgrade_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_strings": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"all_connection_strings": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"dedicated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"connection_urls": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apex_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machine_learning_user_management_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_dev_web_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"cpu_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data_safe_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_storage_size_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_workload": {
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
						"failed_data_recovery_in_seconds": {
							Type:     schema.TypeInt,
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
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_data_guard_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_dedicated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_free_tier": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_preview": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_refreshable_clone": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nsg_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"open_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"refreshable_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"refreshable_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_console_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"standby_db": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"lag_time_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_deletion_of_free_autonomous_database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_begin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_failover": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_refresh": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_refresh_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_switchover": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_next_refresh": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_reclamation_of_free_autonomous_database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"used_data_storage_size_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"whitelisted_ips": {
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
	}
}

func readDatabaseAutonomousDatabasesClones(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabasesClonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousDatabasesClonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseClonesResponse
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseClonesRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if cloneType, ok := s.D.GetOkExists("clone_type"); ok {
		request.CloneType = oci_database.ListAutonomousDatabaseClonesCloneTypeEnum(cloneType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseClones(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseClones(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabasesClone := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AutonomousContainerDatabaseId != nil {
			autonomousDatabasesClone["autonomous_container_database_id"] = *r.AutonomousContainerDatabaseId
		}

		autonomousDatabasesClone["available_upgrade_versions"] = r.AvailableUpgradeVersions

		if r.ConnectionStrings != nil {
			autonomousDatabasesClone["connection_strings"] = []interface{}{AutonomousDatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			autonomousDatabasesClone["connection_strings"] = nil
		}

		if r.ConnectionUrls != nil {
			autonomousDatabasesClone["connection_urls"] = []interface{}{AutonomousDatabaseConnectionUrlsToMap(r.ConnectionUrls)}
		} else {
			autonomousDatabasesClone["connection_urls"] = nil
		}

		if r.CpuCoreCount != nil {
			autonomousDatabasesClone["cpu_core_count"] = *r.CpuCoreCount
		}

		autonomousDatabasesClone["data_safe_status"] = r.DataSafeStatus

		if r.DataStorageSizeInTBs != nil {
			autonomousDatabasesClone["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbName != nil {
			autonomousDatabasesClone["db_name"] = *r.DbName
		}

		if r.DbVersion != nil {
			autonomousDatabasesClone["db_version"] = *r.DbVersion
		}

		autonomousDatabasesClone["db_workload"] = r.DbWorkload

		if r.DefinedTags != nil {
			autonomousDatabasesClone["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousDatabasesClone["display_name"] = *r.DisplayName
		}

		if r.FailedDataRecoveryInSeconds != nil {
			autonomousDatabasesClone["failed_data_recovery_in_seconds"] = *r.FailedDataRecoveryInSeconds
		}

		autonomousDatabasesClone["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousDatabasesClone["id"] = *r.Id
		}

		autonomousDatabasesClone["infrastructure_type"] = r.InfrastructureType

		if r.IsAutoScalingEnabled != nil {
			autonomousDatabasesClone["is_auto_scaling_enabled"] = *r.IsAutoScalingEnabled
		}

		if r.IsDataGuardEnabled != nil {
			autonomousDatabasesClone["is_data_guard_enabled"] = *r.IsDataGuardEnabled
		}

		if r.IsDedicated != nil {
			autonomousDatabasesClone["is_dedicated"] = *r.IsDedicated
		}

		if r.IsFreeTier != nil {
			autonomousDatabasesClone["is_free_tier"] = *r.IsFreeTier
		}

		if r.IsPreview != nil {
			autonomousDatabasesClone["is_preview"] = *r.IsPreview
		}

		if r.IsRefreshableClone != nil {
			autonomousDatabasesClone["is_refreshable_clone"] = *r.IsRefreshableClone
		}

		autonomousDatabasesClone["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousDatabasesClone["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousDatabasesClone["nsg_ids"] = r.NsgIds

		autonomousDatabasesClone["open_mode"] = r.OpenMode

		if r.PrivateEndpoint != nil {
			autonomousDatabasesClone["private_endpoint"] = *r.PrivateEndpoint
		}

		if r.PrivateEndpointIp != nil {
			autonomousDatabasesClone["private_endpoint_ip"] = *r.PrivateEndpointIp
		}

		if r.PrivateEndpointLabel != nil {
			autonomousDatabasesClone["private_endpoint_label"] = *r.PrivateEndpointLabel
		}

		autonomousDatabasesClone["refreshable_mode"] = r.RefreshableMode

		autonomousDatabasesClone["refreshable_status"] = r.RefreshableStatus

		autonomousDatabasesClone["role"] = r.Role

		if r.ServiceConsoleUrl != nil {
			autonomousDatabasesClone["service_console_url"] = *r.ServiceConsoleUrl
		}

		if r.SourceId != nil {
			autonomousDatabasesClone["source_id"] = *r.SourceId
		}

		if r.StandbyDb != nil {
			autonomousDatabasesClone["standby_db"] = []interface{}{AutonomousDatabaseStandbySummaryToMap(r.StandbyDb)}
		} else {
			autonomousDatabasesClone["standby_db"] = nil
		}

		autonomousDatabasesClone["state"] = r.LifecycleState

		if r.SubnetId != nil {
			autonomousDatabasesClone["subnet_id"] = *r.SubnetId
		}

		if r.SystemTags != nil {
			autonomousDatabasesClone["system_tags"] = systemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			autonomousDatabasesClone["time_created"] = r.TimeCreated.String()
		}

		if r.TimeDeletionOfFreeAutonomousDatabase != nil {
			autonomousDatabasesClone["time_deletion_of_free_autonomous_database"] = r.TimeDeletionOfFreeAutonomousDatabase.String()
		}

		if r.TimeMaintenanceBegin != nil {
			autonomousDatabasesClone["time_maintenance_begin"] = r.TimeMaintenanceBegin.String()
		}

		if r.TimeMaintenanceEnd != nil {
			autonomousDatabasesClone["time_maintenance_end"] = r.TimeMaintenanceEnd.String()
		}

		if r.TimeOfLastFailover != nil {
			autonomousDatabasesClone["time_of_last_failover"] = r.TimeOfLastFailover.String()
		}

		if r.TimeOfLastRefresh != nil {
			autonomousDatabasesClone["time_of_last_refresh"] = r.TimeOfLastRefresh.String()
		}

		if r.TimeOfLastRefreshPoint != nil {
			autonomousDatabasesClone["time_of_last_refresh_point"] = r.TimeOfLastRefreshPoint.String()
		}

		if r.TimeOfLastSwitchover != nil {
			autonomousDatabasesClone["time_of_last_switchover"] = r.TimeOfLastSwitchover.String()
		}

		if r.TimeOfNextRefresh != nil {
			autonomousDatabasesClone["time_of_next_refresh"] = r.TimeOfNextRefresh.String()
		}

		if r.TimeReclamationOfFreeAutonomousDatabase != nil {
			autonomousDatabasesClone["time_reclamation_of_free_autonomous_database"] = r.TimeReclamationOfFreeAutonomousDatabase.String()
		}

		if r.UsedDataStorageSizeInTBs != nil {
			autonomousDatabasesClone["used_data_storage_size_in_tbs"] = *r.UsedDataStorageSizeInTBs
		}

		autonomousDatabasesClone["whitelisted_ips"] = r.WhitelistedIps

		resources = append(resources, autonomousDatabasesClone)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabasesClonesDataSource().Schema["autonomous_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_databases", resources); err != nil {
		return err
	}

	return nil
}
