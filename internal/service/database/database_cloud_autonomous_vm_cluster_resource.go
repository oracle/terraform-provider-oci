// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseCloudAutonomousVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseCloudAutonomousVmCluster,
		Read:     readDatabaseCloudAutonomousVmCluster,
		Update:   updateDatabaseCloudAutonomousVmCluster,
		Delete:   deleteDatabaseCloudAutonomousVmCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"cluster_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cpu_core_count_per_node": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compute_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_mtls_enabled_vm_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maintenance_window_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"custom_action_timeout_in_mins": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"days_of_week": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"hours_of_day": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 20,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"is_custom_action_timeout_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"lead_time_in_weeks": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"months": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"patching_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preference": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"skip_ru": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Computed
					},
				},
			},
			"memory_per_oracle_compute_unit_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scan_listener_port_non_tls": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"scan_listener_port_tls": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"total_container_databases": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"autonomous_data_storage_percentage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"available_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"available_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu_percentage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeFloat,
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
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_storage_in_tbs_lowest_scaled_value": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_update_history_entry_id": {
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
			"max_acds_lowest_scaled_value": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_maintenance_run_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"non_provisionable_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"ocpus_lowest_scaled_value": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisionable_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisioned_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisioned_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reclaimable_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reserved_cpus": {
				Type:     schema.TypeFloat,
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_database_ssl_certificate_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ords_certificate_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_autonomous_data_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"total_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if rotateOrdsCertsTrigger, ok := sync.D.GetOkExists("rotate_ords_certs_trigger"); ok {
		tmp := rotateOrdsCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterOrdsCerts()
			if err != nil {
				return err
			}
		}
	}

	if rotateSslCertsTrigger, ok := sync.D.GetOkExists("rotate_ssl_certs_trigger"); ok {
		tmp := rotateSslCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterSslCerts()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func readDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if rotateOrdsCertsTrigger, ok := sync.D.GetOkExists("rotate_ords_certs_trigger"); ok && sync.D.HasChange("rotate_ords_certs_trigger") {
		tmp := rotateOrdsCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterOrdsCerts()
			if err != nil {
				return err
			}
		}
	}

	if rotateSslCertsTrigger, ok := sync.D.GetOkExists("rotate_ssl_certs_trigger"); ok && sync.D.HasChange("rotate_ssl_certs_trigger") {
		tmp := rotateSslCertsTrigger.(bool)
		if tmp {
			err := sync.RotateAutonomousCloudVmClusterSslCerts()
			if err != nil {
				return err
			}
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseCloudAutonomousVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudAutonomousVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.CloudAutonomousVmCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.CloudAutonomousVmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Create() error {
	request := oci_database.CreateCloudAutonomousVmClusterRequest{}

	if autonomousDataStorageSizeInTBs, ok := s.D.GetOkExists("autonomous_data_storage_size_in_tbs"); ok {
		tmp := autonomousDataStorageSizeInTBs.(float64)
		request.AutonomousDataStorageSizeInTBs = &tmp
	}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if clusterTimeZone, ok := s.D.GetOkExists("cluster_time_zone"); ok {
		tmp := clusterTimeZone.(string)
		request.ClusterTimeZone = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCountPerNode, ok := s.D.GetOkExists("cpu_core_count_per_node"); ok {
		tmp := cpuCoreCountPerNode.(int)
		request.CpuCoreCountPerNode = &tmp
	}

	if dbServers, ok := s.D.GetOkExists("db_servers"); ok {
		interfaces := dbServers.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("db_servers") {
			request.DbServers = tmp
		}
	}

	if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
		request.ComputeModel = oci_database.CreateCloudAutonomousVmClusterDetailsComputeModelEnum(computeModel.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isMtlsEnabledVmCluster, ok := s.D.GetOkExists("is_mtls_enabled_vm_cluster"); ok {
		tmp := isMtlsEnabledVmCluster.(bool)
		request.IsMtlsEnabledVmCluster = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateCloudAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if maintenanceWindowDetails, ok := s.D.GetOkExists("maintenance_window_details"); ok {
		if tmpList := maintenanceWindowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_details", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowDetails = &tmp
		}
	}

	if memoryPerOracleComputeUnitInGBs, ok := s.D.GetOkExists("memory_per_oracle_compute_unit_in_gbs"); ok {
		tmp := memoryPerOracleComputeUnitInGBs.(int)
		request.MemoryPerOracleComputeUnitInGBs = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if scanListenerPortNonTls, ok := s.D.GetOkExists("scan_listener_port_non_tls"); ok {
		tmp := scanListenerPortNonTls.(int)
		request.ScanListenerPortNonTls = &tmp
	}

	if scanListenerPortTls, ok := s.D.GetOkExists("scan_listener_port_tls"); ok {
		tmp := scanListenerPortTls.(int)
		request.ScanListenerPortTls = &tmp
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if totalContainerDatabases, ok := s.D.GetOkExists("total_container_databases"); ok {
		tmp := totalContainerDatabases.(int)
		request.TotalContainerDatabases = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudAutonomousVmCluster

	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudautonomousvmcluster", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Get() error {
	request := oci_database.GetCloudAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudAutonomousVmCluster
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateCloudAutonomousVmClusterRequest{}

	if autonomousDataStorageSizeInTBs, ok := s.D.GetOkExists("autonomous_data_storage_size_in_tbs"); ok {
		tmp := autonomousDataStorageSizeInTBs.(float64)
		request.AutonomousDataStorageSizeInTBs = &tmp
	}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	if cpuCoreCountPerNode, ok := s.D.GetOkExists("cpu_core_count_per_node"); ok {
		tmp := cpuCoreCountPerNode.(int)
		request.CpuCoreCountPerNode = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.UpdateCloudAutonomousVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if maintenanceWindowDetails, ok := s.D.GetOkExists("maintenance_window_details"); ok {
		if tmpList := maintenanceWindowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_details", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowDetails = &tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		request.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	if totalContainerDatabases, ok := s.D.GetOkExists("total_container_databases"); ok {
		tmp := totalContainerDatabases.(int)
		request.TotalContainerDatabases = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response.CloudAutonomousVmCluster

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteCloudAutonomousVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteCloudAutonomousVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) SetData() error {
	if s.Res.AutonomousDataStoragePercentage != nil {
		s.D.Set("autonomous_data_storage_percentage", *s.Res.AutonomousDataStoragePercentage)
	}

	if s.Res.AutonomousDataStorageSizeInTBs != nil {
		s.D.Set("autonomous_data_storage_size_in_tbs", *s.Res.AutonomousDataStorageSizeInTBs)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailableAutonomousDataStorageSizeInTBs != nil {
		s.D.Set("available_autonomous_data_storage_size_in_tbs", *s.Res.AvailableAutonomousDataStorageSizeInTBs)
	}

	if s.Res.AvailableContainerDatabases != nil {
		s.D.Set("available_container_databases", *s.Res.AvailableContainerDatabases)
	}

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.ClusterTimeZone != nil {
		s.D.Set("cluster_time_zone", *s.Res.ClusterTimeZone)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.CpuCoreCountPerNode != nil {
		s.D.Set("cpu_core_count_per_node", *s.Res.CpuCoreCountPerNode)
	}

	if s.Res.CpuPercentage != nil {
		s.D.Set("cpu_percentage", *s.Res.CpuPercentage)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	s.D.Set("db_servers", s.Res.DbServers)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	if s.Res.ExadataStorageInTBsLowestScaledValue != nil {
		s.D.Set("exadata_storage_in_tbs_lowest_scaled_value", *s.Res.ExadataStorageInTBsLowestScaledValue)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IsMtlsEnabledVmCluster != nil {
		s.D.Set("is_mtls_enabled_vm_cluster", *s.Res.IsMtlsEnabledVmCluster)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	if s.Res.LastUpdateHistoryEntryId != nil {
		s.D.Set("last_update_history_entry_id", *s.Res.LastUpdateHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MaxAcdsLowestScaledValue != nil {
		s.D.Set("max_acds_lowest_scaled_value", *s.Res.MaxAcdsLowestScaledValue)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NonProvisionableAutonomousContainerDatabases != nil {
		s.D.Set("non_provisionable_autonomous_container_databases", *s.Res.NonProvisionableAutonomousContainerDatabases)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.OcpusLowestScaledValue != nil {
		s.D.Set("ocpus_lowest_scaled_value", *s.Res.OcpusLowestScaledValue)
	}

	if s.Res.ProvisionableAutonomousContainerDatabases != nil {
		s.D.Set("provisionable_autonomous_container_databases", *s.Res.ProvisionableAutonomousContainerDatabases)
	}

	if s.Res.ProvisionedAutonomousContainerDatabases != nil {
		s.D.Set("provisioned_autonomous_container_databases", *s.Res.ProvisionedAutonomousContainerDatabases)
	}

	if s.Res.ProvisionedCpus != nil {
		s.D.Set("provisioned_cpus", *s.Res.ProvisionedCpus)
	}

	if s.Res.ReclaimableCpus != nil {
		s.D.Set("reclaimable_cpus", *s.Res.ReclaimableCpus)
	}

	if s.Res.ReservedCpus != nil {
		s.D.Set("reserved_cpus", *s.Res.ReservedCpus)
	}

	if s.Res.ScanListenerPortNonTls != nil {
		s.D.Set("scan_listener_port_non_tls", *s.Res.ScanListenerPortNonTls)
	}

	if s.Res.ScanListenerPortTls != nil {
		s.D.Set("scan_listener_port_tls", *s.Res.ScanListenerPortTls)
	}

	if s.Res.SecurityAttributes != nil {
		s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDatabaseSslCertificateExpires != nil {
		s.D.Set("time_database_ssl_certificate_expires", s.Res.TimeDatabaseSslCertificateExpires.String())
	}

	if s.Res.TimeOrdsCertificateExpires != nil {
		s.D.Set("time_ords_certificate_expires", s.Res.TimeOrdsCertificateExpires.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalAutonomousDataStorageInTBs != nil {
		s.D.Set("total_autonomous_data_storage_in_tbs", *s.Res.TotalAutonomousDataStorageInTBs)
	}

	if s.Res.TotalContainerDatabases != nil {
		s.D.Set("total_container_databases", *s.Res.TotalContainerDatabases)
	}

	if s.Res.TotalCpus != nil {
		s.D.Set("total_cpus", *s.Res.TotalCpus)
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func DayOfWeekToMap(obj oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_database.MaintenanceWindow, error) {
	result := oci_database.MaintenanceWindow{}

	if customActionTimeoutInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_action_timeout_in_mins")); ok {
		tmp := customActionTimeoutInMins.(int)
		result.CustomActionTimeoutInMins = &tmp
	}

	if daysOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_week")); ok {
		interfaces := daysOfWeek.([]interface{})
		tmp := make([]oci_database.DayOfWeek, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "days_of_week"), stateDataIndex)
			converted, err := s.mapToDayOfWeek(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_week")) {
			result.DaysOfWeek = tmp
		}
	}

	if hoursOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hours_of_day")); ok {
		interfaces := hoursOfDay.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hours_of_day")) {
			result.HoursOfDay = tmp
		}
	}

	if isCustomActionTimeoutEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_custom_action_timeout_enabled")); ok {
		tmp := isCustomActionTimeoutEnabled.(bool)
		result.IsCustomActionTimeoutEnabled = &tmp
	}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
	}

	if leadTimeInWeeks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lead_time_in_weeks")); ok {
		tmp := leadTimeInWeeks.(int)
		result.LeadTimeInWeeks = &tmp
	}

	if months, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "months")); ok {
		interfaces := months.([]interface{})
		tmp := make([]oci_database.Month, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "months"), stateDataIndex)
			converted, err := s.mapToMonth(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "months")) {
			result.Months = tmp
		}
	}

	if patchingMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patching_mode")); ok {
		result.PatchingMode = oci_database.MaintenanceWindowPatchingModeEnum(patchingMode.(string))
	}

	if preference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "preference")); ok {
		result.Preference = oci_database.MaintenanceWindowPreferenceEnum(preference.(string))
	}

	if skipRu, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "skip_ru")); ok {
		interfaces := skipRu.([]interface{})
		tmp := make([]bool, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(bool)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "skip_ru")) {
			result.SkipRu = tmp
		}
	}

	if weeksOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")); ok {
		interfaces := weeksOfMonth.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")) {
			result.WeeksOfMonth = tmp
		}
	}

	return result, nil
}

func MaintenanceWindowToMap(obj *oci_database.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomActionTimeoutInMins != nil {
		result["custom_action_timeout_in_mins"] = int(*obj.CustomActionTimeoutInMins)
	}

	daysOfWeek := []interface{}{}
	for _, item := range obj.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, DayOfWeekToMap(item))
	}
	result["days_of_week"] = daysOfWeek

	result["hours_of_day"] = obj.HoursOfDay
	result["hours_of_day"] = obj.HoursOfDay

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
	}

	if obj.IsMonthlyPatchingEnabled != nil {
		result["is_monthly_patching_enabled"] = bool(*obj.IsMonthlyPatchingEnabled)
	}

	if obj.LeadTimeInWeeks != nil {
		result["lead_time_in_weeks"] = int(*obj.LeadTimeInWeeks)
	}

	months := []interface{}{}
	for _, item := range obj.Months {
		months = append(months, MonthToMap(item))
	}
	result["months"] = months

	result["patching_mode"] = string(obj.PatchingMode)

	result["preference"] = string(obj.Preference)

	result["skip_ru"] = obj.SkipRu

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func MonthToMap(obj oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeCloudAutonomousVmClusterCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CloudAutonomousVmClusterId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeCloudAutonomousVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudAutonomousVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) RotateAutonomousCloudVmClusterOrdsCerts() error {
	request := oci_database.RotateCloudAutonomousVmClusterOrdsCertsRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateCloudAutonomousVmClusterOrdsCerts(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterResourceCrud) RotateAutonomousCloudVmClusterSslCerts() error {
	request := oci_database.RotateCloudAutonomousVmClusterSslCertsRequest{}

	tmp := s.D.Id()
	request.CloudAutonomousVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateCloudAutonomousVmClusterSslCerts(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
	if err != nil {
		return err
	}

	return nil
}
