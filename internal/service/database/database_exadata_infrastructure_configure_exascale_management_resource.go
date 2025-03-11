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

func DatabaseExadataInfrastructureConfigureExascaleManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExadataInfrastructureConfigureExascaleManagement,
		Read:     readDatabaseExadataInfrastructureConfigureExascaleManagement,
		Delete:   deleteDatabaseExadataInfrastructureConfigureExascaleManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"exadata_infrastructure_id": {
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
			"additional_compute_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"additional_compute_system_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_storage_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_network_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_control_plane_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_control_plane_server2": {
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
			"contacts": {
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
						"is_contact_mos_validated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_primary": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"corporate_proxy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"csi_number": {
				Type:     schema.TypeString,
				Computed: true,
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
			"dns_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"infini_band_network_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cps_offline_report_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_multi_rack_deployment": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_scheduling_policy_associated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_slo_status": {
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
			"multi_rack_configuration_file": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_bonding_mode_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_network_bonding_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_network_bonding_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dr_network_bonding_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ntp_server": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"rack_serial_number": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExadataInfrastructureConfigureExascaleManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExadataInfrastructure
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateCreating),
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
	}
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
	}
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleting),
	}
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleted),
	}
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) Create() error {
	request := oci_database.ConfigureExascaleExadataInfrastructureRequest{}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if totalStorageInGBs, ok := s.D.GetOkExists("total_storage_in_gbs"); ok {
		tmp := totalStorageInGBs.(int)
		request.TotalStorageInGBs = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConfigureExascaleExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadatainfrastructure", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	s.Res = &response.ExadataInfrastructure
	return nil
}

func (s *DatabaseExadataInfrastructureConfigureExascaleManagementResourceCrud) SetData() error {
	if s.Res.ActivatedStorageCount != nil {
		s.D.Set("activated_storage_count", *s.Res.ActivatedStorageCount)
	}

	if s.Res.AdditionalComputeCount != nil {
		s.D.Set("additional_compute_count", *s.Res.AdditionalComputeCount)
	}

	s.D.Set("additional_compute_system_model", s.Res.AdditionalComputeSystemModel)

	if s.Res.AdditionalStorageCount != nil {
		s.D.Set("additional_storage_count", *s.Res.AdditionalStorageCount)
	}

	if s.Res.AdminNetworkCIDR != nil {
		s.D.Set("admin_network_cidr", *s.Res.AdminNetworkCIDR)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCount != nil {
		s.D.Set("compute_count", *s.Res.ComputeCount)
	}

	contacts := []interface{}{}
	for _, item := range s.Res.Contacts {
		contacts = append(contacts, ExadataInfrastructureContactToMap(item))
	}
	s.D.Set("contacts", contacts)

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.CsiNumber != nil {
		s.D.Set("csi_number", *s.Res.CsiNumber)
	}

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

	s.D.Set("dns_server", s.Res.DnsServer)

	if s.Res.ExascaleConfig != nil {
		s.D.Set("exascale_config", []interface{}{ExascaleConfigDetailsToMap(s.Res.ExascaleConfig)})
	} else {
		s.D.Set("exascale_config", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.InfiniBandNetworkCIDR != nil {
		s.D.Set("infini_band_network_cidr", *s.Res.InfiniBandNetworkCIDR)
	}

	if s.Res.IsCpsOfflineReportEnabled != nil {
		s.D.Set("is_cps_offline_report_enabled", *s.Res.IsCpsOfflineReportEnabled)
	}

	if s.Res.IsMultiRackDeployment != nil {
		s.D.Set("is_multi_rack_deployment", *s.Res.IsMultiRackDeployment)
	}

	if s.Res.IsSchedulingPolicyAssociated != nil {
		s.D.Set("is_scheduling_policy_associated", *s.Res.IsSchedulingPolicyAssociated)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_slo_status", s.Res.MaintenanceSLOStatus)

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

	if s.Res.MultiRackConfigurationFile != nil {
		s.D.Set("multi_rack_configuration_file", string(s.Res.MultiRackConfigurationFile))
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	if s.Res.NetworkBondingModeDetails != nil {
		s.D.Set("network_bonding_mode_details", []interface{}{NetworkBondingModeDetailsToMap(s.Res.NetworkBondingModeDetails)})
	} else {
		s.D.Set("network_bonding_mode_details", nil)
	}

	s.D.Set("ntp_server", s.Res.NtpServer)

	if s.Res.RackSerialNumber != nil {
		s.D.Set("rack_serial_number", *s.Res.RackSerialNumber)
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}
