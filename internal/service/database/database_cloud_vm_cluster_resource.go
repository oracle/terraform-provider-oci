// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseCloudVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseCloudVmCluster,
		Read:   readDatabaseCloudVmCluster,
		Update: updateDatabaseCloudVmCluster,
		Delete: deleteDatabaseCloudVmCluster,
		Schema: map[string]*schema.Schema{
			// Required
			"backup_subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gi_version": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.GiVersionDiffSuppress,
			},
			"hostname": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DbSystemHostnameDiffSuppress,
			},
			"ssh_public_keys": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"create_async": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"backup_network_nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_collection_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_diagnostics_events_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_health_monitoring_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_incident_logs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"data_storage_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"file_system_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"file_system_size_gb": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mount_point": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_local_backup_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_sparse_diskgroup_enabled": {
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
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"ocpu_count": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"private_zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"scan_listener_port_tcp": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"scan_listener_port_tcp_ssl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"system_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_redundancy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iorm_config_cache": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_plans": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flash_cache_limit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"share": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"objective": {
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
			"last_update_history_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"listener_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scan_dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_dns_record_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_ip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_size_in_gbs": {
				Type:     schema.TypeInt,
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
			"vip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseCloudVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseCloudVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseCloudVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseCloudVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseCloudVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.CloudVmCluster
	Infra                  *oci_database.CloudExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseCloudVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCloudVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) CreatedTarget() []string {
	if createAsyn, ok := s.D.GetOk("create_async"); ok {
		tmp := createAsyn.(bool)
		if tmp {
			return []string{
				string(oci_database.CloudVmClusterLifecycleStateAvailable),
				string(oci_database.CloudVmClusterLifecycleStateProvisioning),
			}
		}
	}
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateProvisioning),
		string(oci_database.CloudVmClusterLifecycleStateUpdating),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.CloudVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseCloudVmClusterResourceCrud) Create() error {
	request := oci_database.CreateCloudVmClusterRequest{}

	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		set := backupNetworkNsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
			request.BackupNetworkNsgIds = tmp
		}
	}

	if backupSubnetId, ok := s.D.GetOkExists("backup_subnet_id"); ok {
		tmp := backupSubnetId.(string)
		request.BackupSubnetId = &tmp
	}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
		tmp := clusterName.(string)
		request.ClusterName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
		if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
			tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataCollectionOptions = &tmp
		}
	}

	if dataStoragePercentage, ok := s.D.GetOkExists("data_storage_percentage"); ok {
		tmp := dataStoragePercentage.(int)
		request.DataStoragePercentage = &tmp
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(float64)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbNodeStorageSizeInGBs, ok := s.D.GetOkExists("db_node_storage_size_in_gbs"); ok {
		tmp := dbNodeStorageSizeInGBs.(int)
		request.DbNodeStorageSizeInGBs = &tmp
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

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if fileSystemConfigurationDetails, ok := s.D.GetOkExists("file_system_configuration_details"); ok {
		interfaces := fileSystemConfigurationDetails.([]interface{})
		tmp := make([]oci_database.FileSystemConfigurationDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "file_system_configuration_details", stateDataIndex)
			converted, err := s.mapToFileSystemConfigurationDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("file_system_configuration_details") {
			request.FileSystemConfigurationDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if giVersion, ok := s.D.GetOkExists("gi_version"); ok {
		tmp := giVersion.(string)
		request.GiVersion = &tmp
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if isLocalBackupEnabled, ok := s.D.GetOkExists("is_local_backup_enabled"); ok {
		tmp := isLocalBackupEnabled.(bool)
		request.IsLocalBackupEnabled = &tmp
	}

	if isSparseDiskgroupEnabled, ok := s.D.GetOkExists("is_sparse_diskgroup_enabled"); ok {
		tmp := isSparseDiskgroupEnabled.(bool)
		request.IsSparseDiskgroupEnabled = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateCloudVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(int)
		request.MemorySizeInGBs = &tmp
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

	if privateZoneId, ok := s.D.GetOkExists("private_zone_id"); ok {
		tmp := privateZoneId.(string)
		request.PrivateZoneId = &tmp
	}

	if scanListenerPortTcp, ok := s.D.GetOkExists("scan_listener_port_tcp"); ok {
		tmp := scanListenerPortTcp.(int)
		request.ScanListenerPortTcp = &tmp
	}

	if scanListenerPortTcpSsl, ok := s.D.GetOkExists("scan_listener_port_tcp_ssl"); ok {
		tmp := scanListenerPortTcpSsl.(int)
		request.ScanListenerPortTcpSsl = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
			request.SshPublicKeys = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if systemVersion, ok := s.D.GetOkExists("system_version"); ok {
		tmp := systemVersion.(string)
		request.SystemVersion = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateCloudVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudVmCluster", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.CloudVmCluster
	return nil
}

func (s *DatabaseCloudVmClusterResourceCrud) Get() error {
	request := oci_database.GetCloudVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetCloudVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudVmCluster
	return nil
}

func (s *DatabaseCloudVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateCloudVmClusterRequest{}

	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		set := backupNetworkNsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("backup_network_nsg_ids") {
			request.BackupNetworkNsgIds = tmp
		}
	}

	tmp := s.D.Id()
	request.CloudVmClusterId = &tmp

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		if s.Infra == nil || s.Infra.Id == nil {
			err := s.getInfraInfo(cloudExadataInfrastructureId.(string))
			if err != nil {
				log.Printf("[ERROR] Could not get Cloud Exadata Infrastructure info for the : %v", err)
			}
		}
	}
	if !utils.IsMultiVm(*s.Infra.Shape, s.Infra.MaxDataStorageInTBs) {
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			if s.Infra.ComputeCount != nil && *s.Infra.ComputeCount != nodeCount {
				request.ComputeNodes = []string{"ALL"}
			} else {
				request.ComputeNodes = []string{}
				if s.Infra.StorageCount != nil {
					if shape, ok := s.D.GetOkExists("shape"); ok {
						flexShape := shape.(string) + ".StorageServer"

						if compartmentId, compOk := s.D.GetOkExists("compartment_id"); compOk {
							flex, err := s.flexAvailableDbStorageInGBs(compartmentId.(string), flexShape)

							if err == nil {
								if storageSizeInGBs, ok := s.D.GetOkExists("storage_size_in_gbs"); ok {
									tmp := flex**s.Infra.StorageCount - storageSizeInGBs.(int)
									request.StorageSizeInGBs = &tmp
								}
							}
						}
					}
				}
			}
		}
	}

	if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok && s.D.HasChange("cpu_core_count") {
		tmp := cpuCoreCount.(int)
		request.CpuCoreCount = &tmp
	}

	if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
		if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
			tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataCollectionOptions = &tmp
		}
	}

	if dataStorageSizeInTBs, ok := s.D.GetOkExists("data_storage_size_in_tbs"); ok {
		tmp := dataStorageSizeInTBs.(float64)
		request.DataStorageSizeInTBs = &tmp
	}

	if dbNodeStorageSizeInGBs, ok := s.D.GetOkExists("db_node_storage_size_in_gbs"); ok {
		tmp := dbNodeStorageSizeInGBs.(int)
		request.DbNodeStorageSizeInGBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileSystemConfigurationDetails, ok := s.D.GetOkExists("file_system_configuration_details"); ok {
		interfaces := fileSystemConfigurationDetails.([]interface{})
		tmp := make([]oci_database.FileSystemConfigurationDetail, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "file_system_configuration_details", stateDataIndex)
			converted, err := s.mapToFileSystemConfigurationDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("file_system_configuration_details") {
			request.FileSystemConfigurationDetails = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		request.LicenseModel = oci_database.UpdateCloudVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if memorySizeInGBs, ok := s.D.GetOkExists("memory_size_in_gbs"); ok {
		tmp := memorySizeInGBs.(int)
		request.MemorySizeInGBs = &tmp
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

	if ocpuCount, ok := s.D.GetOkExists("ocpu_count"); ok && s.D.HasChange("ocpu_count") {
		tmp := float32(ocpuCount.(float64))
		request.OcpuCount = &tmp
	}

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok {
		interfaces := sshPublicKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ssh_public_keys") {
			request.SshPublicKeys = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateCloudVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "cloudVmCluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.CloudVmCluster
	return nil
}

func (s *DatabaseCloudVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteCloudVmClusterRequest{}

	tmp := s.D.Id()
	request.CloudVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteCloudVmCluster(context.Background(), request)
	return err
}

func (s *DatabaseCloudVmClusterResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	backupNetworkNsgIds := []interface{}{}
	for _, item := range s.Res.BackupNetworkNsgIds {
		backupNetworkNsgIds = append(backupNetworkNsgIds, item)
	}
	s.D.Set("backup_network_nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, backupNetworkNsgIds))

	if s.Res.BackupSubnetId != nil {
		s.D.Set("backup_subnet_id", *s.Res.BackupSubnetId)
	}

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.ClusterName != nil {
		s.D.Set("cluster_name", *s.Res.ClusterName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
	}
	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.DataStoragePercentage != nil {
		s.D.Set("data_storage_percentage", *s.Res.DataStoragePercentage)
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

	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	fileSystemConfigurationDetails := []interface{}{}
	for _, item := range s.Res.FileSystemConfigurationDetails {
		fileSystemConfigurationDetails = append(fileSystemConfigurationDetails, FileSystemConfigurationDetailToMap(item))
	}
	s.D.Set("file_system_configuration_details", fileSystemConfigurationDetails)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GiVersion != nil {
		s.D.Set("gi_version", *s.Res.GiVersion)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{ExadataIormConfigToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", nil)
	}

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	if s.Res.IsSparseDiskgroupEnabled != nil {
		s.D.Set("is_sparse_diskgroup_enabled", *s.Res.IsSparseDiskgroupEnabled)
	}

	if s.Res.LastUpdateHistoryEntryId != nil {
		s.D.Set("last_update_history_entry_id", *s.Res.LastUpdateHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", strconv.FormatInt(*s.Res.ListenerPort, 10))
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.ScanDnsName != nil {
		s.D.Set("scan_dns_name", *s.Res.ScanDnsName)
	}

	if s.Res.ScanDnsRecordId != nil {
		s.D.Set("scan_dns_record_id", *s.Res.ScanDnsRecordId)
	}

	s.D.Set("scan_ip_ids", s.Res.ScanIpIds)

	if s.Res.ScanListenerPortTcp != nil {
		s.D.Set("scan_listener_port_tcp", *s.Res.ScanListenerPortTcp)
	}

	if s.Res.ScanListenerPortTcpSsl != nil {
		s.D.Set("scan_listener_port_tcp_ssl", *s.Res.ScanListenerPortTcpSsl)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageSizeInGBs != nil {
		s.D.Set("storage_size_in_gbs", *s.Res.StorageSizeInGBs)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.SystemVersion != nil {
		s.D.Set("system_version", *s.Res.SystemVersion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	s.D.Set("vip_ids", s.Res.VipIds)

	if s.Res.ZoneId != nil {
		s.D.Set("zone_id", *s.Res.ZoneId)
	}

	return nil
}

func (s *DatabaseCloudVmClusterResourceCrud) mapToDataCollectionOptions(fieldKeyFormat string) (oci_database.DataCollectionOptions, error) {
	result := oci_database.DataCollectionOptions{}

	if isDiagnosticsEventsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_diagnostics_events_enabled")); ok {
		tmp := isDiagnosticsEventsEnabled.(bool)
		result.IsDiagnosticsEventsEnabled = &tmp
	}

	if isHealthMonitoringEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_health_monitoring_enabled")); ok {
		tmp := isHealthMonitoringEnabled.(bool)
		result.IsHealthMonitoringEnabled = &tmp
	}

	if isIncidentLogsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_incident_logs_enabled")); ok {
		tmp := isIncidentLogsEnabled.(bool)
		result.IsIncidentLogsEnabled = &tmp
	}

	return result, nil
}

func (s *DatabaseCloudVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeCloudVmClusterCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CloudVmClusterId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeCloudVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseCloudVmClusterResourceCrud) mapToFileSystemConfigurationDetail(fieldKeyFormat string) (oci_database.FileSystemConfigurationDetail, error) {
	result := oci_database.FileSystemConfigurationDetail{}

	if fileSystemSizeGb, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_system_size_gb")); ok {
		tmp := fileSystemSizeGb.(int)
		result.FileSystemSizeGb = &tmp
	}

	if mountPoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_point")); ok {
		tmp := mountPoint.(string)
		result.MountPoint = &tmp
	}

	return result, nil
}

func FileSystemConfigurationDetailToMap(obj oci_database.FileSystemConfigurationDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileSystemSizeGb != nil {
		result["file_system_size_gb"] = int(*obj.FileSystemSizeGb)
	}

	if obj.MountPoint != nil {
		result["mount_point"] = string(*obj.MountPoint)
	}

	return result
}

func (s *DatabaseCloudVmClusterResourceCrud) flexAvailableDbStorageInGBs(compartmentId string, shapeName string) (int, error) {
	request := oci_database.ListFlexComponentsRequest{}
	request.CompartmentId = &compartmentId
	request.Name = &shapeName
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListFlexComponents(context.Background(), request)
	if err != nil {
		return 0, err
	}
	for _, item := range response.FlexComponentCollection.Items {
		return *item.AvailableDbStorageInGBs, nil
	}

	return 0, fmt.Errorf("No flex component found for compartment")
}

func (s *DatabaseCloudVmClusterResourceCrud) getInfraInfo(ceiId string) error {
	request := oci_database.GetCloudExadataInfrastructureRequest{}

	request.CloudExadataInfrastructureId = &ceiId

	response, err := s.Client.GetCloudExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Infra = &response.CloudExadataInfrastructure

	return nil
}
