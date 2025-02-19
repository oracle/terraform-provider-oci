// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseCloudExadataInfrastructureConfigureExascaleManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseCloudExadataInfrastructureConfigureExascaleManagement,
		Read:     readDatabaseCloudExadataInfrastructureConfigureExascaleManagement,
		Delete:   deleteDatabaseCloudExadataInfrastructureConfigureExascaleManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"total_storage_in_gbs": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"activated_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"additional_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cluster_placement_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"customer_contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_file_system_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_backup_partition": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_resizable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"min_size_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mount_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"exascale_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_storage_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_storage_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_scheduling_policy_associated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"months": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"patching_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"skip_ru": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"max_cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_data_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"max_db_node_storage_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_memory_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"monthly_db_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monthly_storage_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"storage_server_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": {
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
			"total_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDatabaseCloudExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseCloudExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseCloudExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.CloudExadataInfrastructure
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateProvisioning),
	}
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateTerminating),
	}
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.CloudExadataInfrastructureLifecycleStateTerminated),
	}
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) Create() error {
	request := oci_database.ConfigureExascaleCloudExadataInfrastructureRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if totalStorageInGBs, ok := s.D.GetOkExists("total_storage_in_gbs"); ok {
		tmp := totalStorageInGBs.(int)
		request.TotalStorageInGBs = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConfigureExascaleCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.CloudExadataInfrastructure

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudexadatainfrastructure", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceCrud) SetData() error {
	if s.Res.ActivatedStorageCount != nil {
		s.D.Set("activated_storage_count", *s.Res.ActivatedStorageCount)
	}

	if s.Res.AdditionalStorageCount != nil {
		s.D.Set("additional_storage_count", *s.Res.AdditionalStorageCount)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableStorageSizeInGBs != nil {
		s.D.Set("available_storage_size_in_gbs", *s.Res.AvailableStorageSizeInGBs)
	}

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	if s.Res.CpuCount != nil {
		s.D.Set("cpu_count", *s.Res.CpuCount)
	}

	customerContacts := []interface{}{}
	for _, item := range s.Res.CustomerContacts {
		customerContacts = append(customerContacts, CustomerContactToMap(item))
	}
	s.D.Set("customer_contacts", customerContacts)

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerVersion != nil {
		s.D.Set("db_server_version", *s.Res.DbServerVersion)
	}

	definedFileSystemConfigurations := []interface{}{}
	for _, item := range s.Res.DefinedFileSystemConfigurations {
		definedFileSystemConfigurations = append(definedFileSystemConfigurations, DefinedFileSystemConfigurationToMap(item))
	}
	s.D.Set("defined_file_system_configurations", definedFileSystemConfigurations)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExascaleConfig != nil {
		s.D.Set("exascale_config", []interface{}{ExascaleConfigDetailsToMap(s.Res.ExascaleConfig)})
	} else {
		s.D.Set("exascale_config", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSchedulingPolicyAssociated != nil {
		s.D.Set("is_scheduling_policy_associated", *s.Res.IsSchedulingPolicyAssociated)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MaxCpuCount != nil {
		s.D.Set("max_cpu_count", *s.Res.MaxCpuCount)
	}

	if s.Res.MaxDataStorageInTBs != nil {
		s.D.Set("max_data_storage_in_tbs", *s.Res.MaxDataStorageInTBs)
	}

	if s.Res.MaxDbNodeStorageInGBs != nil {
		s.D.Set("max_db_node_storage_in_gbs", *s.Res.MaxDbNodeStorageInGBs)
	}

	if s.Res.MaxMemoryInGBs != nil {
		s.D.Set("max_memory_in_gbs", *s.Res.MaxMemoryInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.MonthlyDbServerVersion != nil {
		s.D.Set("monthly_db_server_version", *s.Res.MonthlyDbServerVersion)
	}

	if s.Res.MonthlyStorageServerVersion != nil {
		s.D.Set("monthly_storage_server_version", *s.Res.MonthlyStorageServerVersion)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageCount != nil {
		s.D.Set("storage_count", *s.Res.StorageCount)
	}

	if s.Res.StorageServerVersion != nil {
		s.D.Set("storage_server_version", *s.Res.StorageServerVersion)
	}

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	total_storage_in_gbs := s.D.Get("total_storage_in_gbs")
	s.D.Set("total_storage_size_in_gbs", total_storage_in_gbs)

	return nil
}
