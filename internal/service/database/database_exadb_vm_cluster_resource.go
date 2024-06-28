// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExadbVmClusterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"), // TODO
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create:        createDatabaseExadbVmCluster,
		Read:          readDatabaseExadbVmCluster,
		Update:        updateDatabaseExadbVmCluster,
		Delete:        deleteDatabaseExadbVmCluster,
		CustomizeDiff: customValidationOnNodeResources,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"backup_subnet_id": {
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
			"exascale_db_storage_vault_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"grid_image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"node_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"enabled_ecpu_count_per_node": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"total_ecpu_count_per_node": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"vm_file_system_storage_size_gbs_per_node": {
							Type:     schema.TypeInt,
							Required: true,
						},
						// Computed
						"memory_size_in_gbs_per_node": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"snapshot_file_system_storage_size_gbs_per_node": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_file_system_storage_size_gbs_per_node": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"node_resource": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Set:      nodeResourceHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"node_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: stringContainsNoSpaceAndIsNotBlack(),
						},
						// Computed
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_hostname": {
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

			// Optional
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"license_model": {
				Type:     schema.TypeString,
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
			"system_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"gi_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grid_image_type": {
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
			"state": {
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

func createDatabaseExadbVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseExadbVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseExadbVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseExadbVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExadbVmClusterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExadbVmCluster
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExadbVmClusterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExadbVmClusterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateProvisioning),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateTerminating),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateTerminated),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) UpdatePending() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateUpdating),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) UpdateTarget() []string {
	return []string{
		string(oci_database.ExadbVmClusterLifecycleStateAvailable),
	}
}

func (s *DatabaseExadbVmClusterResourceCrud) Create() error {
	request := oci_database.CreateExadbVmClusterRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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

	if clusterName, ok := s.D.GetOkExists("cluster_name"); ok {
		tmp := clusterName.(string)
		request.ClusterName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	err := s.setNodeConfigInCreateExaDbVmClusterRequest(&request)
	if err != nil {
		return err
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

	if exascaleDbStorageVaultId, ok := s.D.GetOkExists("exascale_db_storage_vault_id"); ok {
		tmp := exascaleDbStorageVaultId.(string)
		request.ExascaleDbStorageVaultId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gridImageId, ok := s.D.GetOkExists("grid_image_id"); ok {
		tmp := gridImageId.(string)
		request.GridImageId = &tmp
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
		request.LicenseModel = oci_database.CreateExadbVmClusterDetailsLicenseModelEnum(licenseModel.(string))
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

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
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

	if systemVersion, ok := s.D.GetOkExists("system_version"); ok {
		tmp := systemVersion.(string)
		request.SystemVersion = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	response, err := s.Client.CreateExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.ExadbVmCluster

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadbvmcluster", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseExadbVmClusterResourceCrud) Get() error {
	request := oci_database.GetExadbVmClusterRequest{}

	tmp := s.D.Id()
	request.ExadbVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadbVmCluster
	return nil
}

func (s *DatabaseExadbVmClusterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateExadbVmClusterRequest{}

	updateRequired := false
	oldNodeCount := 1
	newNodeCount := 1
	nodeResourceKey := getNodeResourceKey()
	if _, ok := s.D.GetOkExists(nodeResourceKey); ok {
		oldNodeListRaw, newNodeListRaw := s.D.GetChange(nodeResourceKey)
		oldNodeListRawSet := oldNodeListRaw.(*schema.Set)
		oldNodeList := oldNodeListRawSet.List()
		newNodeListRawSet := newNodeListRaw.(*schema.Set)
		newNodeList := newNodeListRawSet.List()
		oldNodeCount = len(oldNodeList)
		newNodeCount = len(newNodeList)
		if newNodeCount < oldNodeCount {
			err := s.removeVirtualMachineFromExadbVmCluster(oldNodeList, newNodeList)
			if err != nil {
				return err
			}
		} else if newNodeCount > oldNodeCount {
			updateRequired = true
			request.NodeCount = &newNodeCount
		}
	}

	enableCpuCountKey := getEnableCpuCountKey()
	if _, ok := s.D.GetOkExists(enableCpuCountKey); ok {
		oldEnabledCpuCoreCountPerNodeRaw, newEnabledCpuCoreCountPerNodeRaw := s.D.GetChange(enableCpuCountKey)
		oldEnabledCpuCoreCountPerNode := oldEnabledCpuCoreCountPerNodeRaw.(int)
		newEnabledCpuCoreCountPerNode := newEnabledCpuCoreCountPerNodeRaw.(int)
		if oldEnabledCpuCoreCountPerNode != newEnabledCpuCoreCountPerNode {
			updateRequired = true
			tmp := newEnabledCpuCoreCountPerNode * newNodeCount
			request.EnabledECpuCount = &tmp
		}
	}

	totalCpuCountKey := getTotalCpuCountKey()
	if _, ok := s.D.GetOkExists(totalCpuCountKey); ok {
		oldTotalCpuCoreCountPerNodeRaw, newTotalCpuCoreCountPerNodeRaw := s.D.GetChange(totalCpuCountKey)
		oldTotalCpuCoreCountPerNode := oldTotalCpuCoreCountPerNodeRaw.(int)
		newTotalCpuCoreCountPerNode := newTotalCpuCoreCountPerNodeRaw.(int)
		if oldTotalCpuCoreCountPerNode != newTotalCpuCoreCountPerNode {
			updateRequired = true
			tmp := newTotalCpuCoreCountPerNode * newNodeCount
			request.TotalECpuCount = &tmp
		}
	}

	vmStorageSizeKey := getVmStorageSizeKey()
	if _, ok := s.D.GetOkExists(vmStorageSizeKey); ok {
		oldVmFileSystemStoragePerNodeRaw, newVmFileSystemStoragePerNodeRaw := s.D.GetChange(vmStorageSizeKey)
		oldVmFileSystemStoragePerNode := oldVmFileSystemStoragePerNodeRaw.(int)
		newVmFileSystemStoragePerNode := newVmFileSystemStoragePerNodeRaw.(int)
		if oldVmFileSystemStoragePerNode != newVmFileSystemStoragePerNode {
			updateRequired = true
			tmp := newVmFileSystemStoragePerNode * newNodeCount
			request.VmFileSystemStorage = &oci_database.ExadbVmClusterStorageDetails{
				TotalSizeInGbs: &tmp,
			}
		}
	}

	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok && s.D.HasChange("backup_network_nsg_ids") {
		updateRequired = true
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

	if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok && s.D.HasChange("data_collection_options") {
		updateRequired = true
		if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
			tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataCollectionOptions = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		if s.D.HasChange("defined_tags") {
			updateRequired = true
		}
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		// we need to set request.DefinedTags even when there is no change. request.DefinedTags=nil would clear the definedTags
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		updateRequired = true
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ExadbVmClusterId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		if s.D.HasChange("freeform_tags") {
			updateRequired = true
		}
		// we need to set request.DefinedTags even when there is no change. request.DefinedTags=nil would clear the definedTags
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gridImageId, ok := s.D.GetOkExists("grid_image_id"); ok && s.D.HasChange("grid_image_id") {
		updateRequired = true
		tmp := gridImageId.(string)
		request.GridImageId = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists("license_model"); ok && s.D.HasChange("license_model") {
		updateRequired = true
		request.LicenseModel = oci_database.UpdateExadbVmClusterDetailsLicenseModelEnum(licenseModel.(string))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok && s.D.HasChange("nsg_ids") {
		updateRequired = true
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

	if sshPublicKeys, ok := s.D.GetOkExists("ssh_public_keys"); ok && s.D.HasChange("ssh_public_keys") {
		updateRequired = true
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

	if systemVersion, ok := s.D.GetOkExists("system_version"); ok && s.D.HasChange("system_version") {
		updateRequired = true
		tmp := systemVersion.(string)
		request.SystemVersion = &tmp
	}

	if updateAction, ok := s.D.GetOkExists("update_action"); ok && s.D.HasChange("update_action") {
		updateRequired = true
		request.UpdateAction = oci_database.UpdateExadbVmClusterDetailsUpdateActionEnum(updateAction.(string))
	}

	if updateRequired {
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.UpdateExadbVmCluster(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadbvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
	}
	return s.Get()
}

func (s *DatabaseExadbVmClusterResourceCrud) Delete() error {
	request := oci_database.DeleteExadbVmClusterRequest{}

	tmp := s.D.Id()
	request.ExadbVmClusterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadbvmcluster", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseExadbVmClusterResourceCrud) SetData() error {
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

	if s.Res.ClusterName != nil {
		s.D.Set("cluster_name", *s.Res.ClusterName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	if s.Res.ExascaleDbStorageVaultId != nil {
		s.D.Set("exascale_db_storage_vault_id", *s.Res.ExascaleDbStorageVaultId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GiVersion != nil {
		s.D.Set("gi_version", *s.Res.GiVersion)
	}

	if s.Res.GridImageId != nil {
		s.D.Set("grid_image_id", *s.Res.GridImageId)
	}

	s.D.Set("grid_image_type", s.Res.GridImageType)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{ExadataIormConfigToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", nil)
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

	nodeConfigInResponse, nodeListInResponse := getNodeConfigAndNodeListInResponse(s.Res.Id, s.Res.CompartmentId, s.Res.EnabledECpuCount, s.Res.TotalECpuCount, s.Res.VmFileSystemStorage, s.Res.MemorySizeInGBs, s.Res.SnapshotFileSystemStorage, s.Res.TotalFileSystemStorage, s.Client)
	s.D.Set("node_config", []interface{}{nodeConfigInResponse})

	nodeListInState := []interface{}{}
	if tmp, ok := s.D.GetOkExists(getNodeResourceKey()); ok {
		set := tmp.(*schema.Set)
		tmpList := set.List()
		nodeListInState = make([]interface{}, len(tmpList))
		_ = copy(nodeListInState, tmpList)
	}

	// first try to map node in response to node in state by node_id
	unusedNodeListInState := nodeListInState
	for _, tmp := range nodeListInResponse {
		node := tmp.(map[string]interface{})
		// if node_name has not been set and node_id is set
		if _, nodeNameExistInResp := node["node_name"]; !nodeNameExistInResp {
			if nodeId, nodeIdExistInResp := node["node_id"]; nodeIdExistInResp {
				nodeIdInResp := nodeId.(string)
				if nodeIdInResp != "" {
					for j, nodeInState := range nodeListInState {
						nodeInState := nodeInState.(map[string]interface{})
						if nodeIdInState, nodeIdExistInState := nodeInState["node_id"]; nodeIdExistInState {
							nodeIdInState = nodeIdInState.(string)
							if nodeIdInState == nodeIdInResp {
								if nodeNameInState, nodeNameExistInState := nodeInState["node_name"]; nodeNameExistInState {
									nodeNameInState = nodeNameInState.(string)
									node["node_name"] = nodeNameInState
								} else {
									node["node_name"] = ""
								}
								unusedNodeListInState = make([]interface{}, 0)
								unusedNodeListInState = append(unusedNodeListInState, nodeListInState[:j]...)
								unusedNodeListInState = append(unusedNodeListInState, nodeListInState[j+1:]...)
								break
							}
						}
					} // for nodeListInState
					nodeListInState = unusedNodeListInState
				} // if nodeIdInResp != ""
			} // if node["node_id"]
		} // if node["node_name"]
	} // for nodeListInResponse

	// second: try to map node in response to node in state by position for nodes created outside Terraform
	if len(nodeListInState) > 0 {
		for _, tmp := range nodeListInResponse {
			node := tmp.(map[string]interface{})
			// if node_name has not been set
			if _, nodeNameExistInResp := node["node_name"]; !nodeNameExistInResp {
				nodeInState := nodeListInState[0].(map[string]interface{})
				if nodeNameInState, nodeNameExistInState := nodeInState["node_name"]; nodeNameExistInState {
					nodeNameInState = nodeNameInState.(string)
					node["node_name"] = nodeNameInState
				} else {
					node["node_name"] = ""
				}
				unusedNodeListInState = make([]interface{}, 0)
				unusedNodeListInState = append(unusedNodeListInState, nodeListInState[1:]...)

				nodeListInState = unusedNodeListInState
				if len(nodeListInState) == 0 {
					break
				}
			} // if node["node_name"]
		} // for nodeListInResponse
	} // if len(nodeListInState)

	s.D.Set("node_resource", schema.NewSet(nodeResourceHashCodeForSets, nodeListInResponse))

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.PrivateZoneId != nil {
		s.D.Set("private_zone_id", *s.Res.PrivateZoneId)
	}

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

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
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

func (s *DatabaseExadbVmClusterResourceCrud) mapToDataCollectionOptions(fieldKeyFormat string) (oci_database.DataCollectionOptions, error) {
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

func (s *DatabaseExadbVmClusterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExadbVmClusterCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExadbVmClusterId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeExadbVmClusterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadbvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseExadbVmClusterResourceCrud) setNodeConfigInCreateExaDbVmClusterRequest(request *oci_database.CreateExadbVmClusterRequest) error {

	nodeCount := 1
	if tmpList, ok := s.D.GetOkExists(getNodeResourceKey()); ok {
		set := tmpList.(*schema.Set)
		interfaces := set.List()
		nodeCount = len(interfaces)
	} else {
		return fmt.Errorf("node_resource is required but not provided in config")
	}
	request.NodeCount = &nodeCount

	if nodeConfig, ok := s.D.GetOkExists("node_config"); ok {
		if tmpList := nodeConfig.([]interface{}); len(tmpList) > 0 {
			if enabledCpuCoreCountPerNode, ok := s.D.GetOkExists(getEnableCpuCountKey()); ok {
				tmp := enabledCpuCoreCountPerNode.(int) * nodeCount
				request.EnabledECpuCount = &tmp
			}

			if totalCpuCoreCountPerNode, ok := s.D.GetOkExists(getTotalCpuCountKey()); ok {
				tmp := totalCpuCoreCountPerNode.(int) * nodeCount
				request.TotalECpuCount = &tmp
			}

			if vmFileSystemStoragePerNode, ok := s.D.GetOkExists(getVmStorageSizeKey()); ok {
				tmp := vmFileSystemStoragePerNode.(int) * nodeCount
				request.VmFileSystemStorage = &oci_database.ExadbVmClusterStorageDetails{
					TotalSizeInGbs: &tmp,
				}
			}
		}
	}

	return nil
}

func (s *DatabaseExadbVmClusterResourceCrud) removeVirtualMachineFromExadbVmCluster(oldNodeResourceList []interface{}, newNodeResourceList []interface{}) error {
	if len(oldNodeResourceList) <= len(newNodeResourceList) {
		// this method is only applicable for removeVM use case
		return nil
	}
	newNodeIdMap := make(map[string]bool)
	for _, nodeResource := range newNodeResourceList {
		nodeId := getNodeIdFromNodeResource(nodeResource)
		if nodeId != "" {
			newNodeIdMap[nodeId] = true
		}
	}
	removedNodeIdList := []oci_database.DbNodeDetails{}
	for _, nodeResource := range oldNodeResourceList {
		nodeId := getNodeIdFromNodeResource(nodeResource)
		if nodeId != "" {
			if _, exist := newNodeIdMap[nodeId]; !exist {
				tmp := oci_database.DbNodeDetails{}
				tmp.DbNodeId = &nodeId
				removedNodeIdList = append(removedNodeIdList, tmp)
			}
		}
	}

	if len(removedNodeIdList) == 0 {
		return nil
	}

	ExadbVmClusterId := s.D.Id()
	request := oci_database.RemoveVirtualMachineFromExadbVmClusterRequest{
		ExadbVmClusterId: &ExadbVmClusterId,
		RemoveVirtualMachineFromExadbVmClusterDetails: oci_database.RemoveVirtualMachineFromExadbVmClusterDetails{
			DbNodes: removedNodeIdList,
		},
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RemoveVirtualMachineFromExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "exadbvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}

func getNodeConfigKey() string {
	// node_config is an array with exactly one entry
	return "node_config.0"
}

func getEnableCpuCountKey() string {
	return getNodeConfigKey() + ".enabled_ecpu_count_per_node"
}

func getTotalCpuCountKey() string {
	return getNodeConfigKey() + ".total_ecpu_count_per_node"
}

func getVmStorageSizeKey() string {
	return getNodeConfigKey() + ".vm_file_system_storage_size_gbs_per_node"
}

func getNodeResourceKey() string {
	return "node_resource"
}

func stringContainsNoSpaceAndIsNotBlack() func(i interface{}, k string) (warnings []string, errors []error) {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		if strings.ContainsAny(v, " ") {
			errors = append(errors, fmt.Errorf("expected value of %s to not contain any space, got '%v'", k, i))
			return warnings, errors
		}

		if strings.TrimSpace(v) == "" {
			return warnings, []error{fmt.Errorf("expected %q to not be an empty string or whitespace", k)}
		}

		return warnings, errors
	}
}

// key are node_name in lower case, value are node_resource object.
func buildNodeNameToNodeResourceMap(nodeResourceList []interface{}) (map[string]interface{}, error) {
	nodeNameToNodeResourceMap := make(map[string]interface{})
	for _, tmp := range nodeResourceList {
		nodeResource := tmp.(map[string]interface{})
		// imported resource could have node_name = nil / ""
		if nodeName, nodeNameExist := nodeResource["node_name"]; nodeNameExist && nodeName != nil && nodeName != "" {
			nodeNameLowerCase := strings.ToLower(nodeName.(string))
			// validate uniqueness of all node_name
			if _, duplicate := nodeNameToNodeResourceMap[nodeNameLowerCase]; duplicate {
				return nil, fmt.Errorf("expected node_name to be unique (case-insensitive) in all node_resource blocks. Found duplicate: '%s'", nodeNameLowerCase)
			}
			nodeNameToNodeResourceMap[nodeNameLowerCase] = nodeResource
		}
	}
	return nodeNameToNodeResourceMap, nil
}

func getNodeIdFromNodeResource(nodeResource interface{}) string {

	if nodeResourceMap, ok := nodeResource.(map[string]interface{}); ok {
		if nodeId, nodeIdExist := nodeResourceMap["node_id"]; nodeIdExist {
			return nodeId.(string)
		}
	}

	return ""
}

func getNodeConfigAndNodeListInResponse(ExaDbVmClusterId *string, CompartmentId *string, EnabledECpuCount *int, TotalECpuCount *int, VmFileSystemStorage *oci_database.ExadbVmClusterStorageDetails, MemorySizeInGBs *int, SnapshotFileSystemStorage *oci_database.ExadbVmClusterStorageDetails, TotalFileSystemStorage *oci_database.ExadbVmClusterStorageDetails, Client *oci_database.DatabaseClient) (map[string]interface{}, []interface{}) {

	nodeResourceList, err := getNodeResourceList(ExaDbVmClusterId, CompartmentId, Client)
	if err != nil {
		log.Printf("WARNING: Unable to get the list of dbnodes of the Exadata DB VM cluster: %v", err)
		return nil, nil
	}

	nodeConfigMap := map[string]interface{}{}
	nodeCount := len(nodeResourceList)
	if nodeCount != 0 {

		if EnabledECpuCount != nil {
			nodeConfigMap["enabled_ecpu_count_per_node"] = *EnabledECpuCount / nodeCount
		}

		if TotalECpuCount != nil {
			nodeConfigMap["total_ecpu_count_per_node"] = *TotalECpuCount / nodeCount
		}

		if VmFileSystemStorage != nil && VmFileSystemStorage.TotalSizeInGbs != nil {
			nodeConfigMap["vm_file_system_storage_size_gbs_per_node"] = *VmFileSystemStorage.TotalSizeInGbs / nodeCount
		}

		if MemorySizeInGBs != nil {
			nodeConfigMap["memory_size_in_gbs_per_node"] = *MemorySizeInGBs / nodeCount
		}

		if SnapshotFileSystemStorage != nil && SnapshotFileSystemStorage.TotalSizeInGbs != nil {
			nodeConfigMap["snapshot_file_system_storage_size_gbs_per_node"] = *SnapshotFileSystemStorage.TotalSizeInGbs / nodeCount
		}

		if TotalFileSystemStorage != nil && TotalFileSystemStorage.TotalSizeInGbs != nil {
			nodeConfigMap["total_file_system_storage_size_gbs_per_node"] = *TotalFileSystemStorage.TotalSizeInGbs / nodeCount
		}

	}

	return nodeConfigMap, nodeResourceList
}

func getNodeResourceList(VmClusterId *string, CompartmentId *string, Client *oci_database.DatabaseClient) ([]interface{}, error) {
	request := oci_database.ListDbNodesRequest{}

	request.VmClusterId = VmClusterId
	request.CompartmentId = CompartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	dbNodes := []interface{}{}

	response, err := Client.ListDbNodes(context.Background(), request)
	if err != nil {
		return nil, err
	}
	dbNodes = addLiveNodeToNodeResourceList(dbNodes, response)

	request.Page = response.OpcNextPage
	for request.Page != nil {
		response, err := Client.ListDbNodes(context.Background(), request)
		if err != nil {
			return nil, err
		}
		dbNodes = addLiveNodeToNodeResourceList(dbNodes, response)
		request.Page = response.OpcNextPage
	}
	return dbNodes, nil
}

func addLiveNodeToNodeResourceList(dbNodes []interface{}, response oci_database.ListDbNodesResponse) []interface{} {
	for _, item := range response.Items {
		if item.LifecycleState != oci_database.DbNodeSummaryLifecycleStateTerminated {
			dbNodeResource := map[string]interface{}{}

			if item.Id != nil {
				dbNodeResource["node_id"] = *item.Id
			}

			if item.Hostname != nil {
				dbNodeResource["node_hostname"] = *item.Hostname
			}

			dbNodeResource["state"] = item.LifecycleState
			dbNodes = append(dbNodes, dbNodeResource)
		}
	}
	return dbNodes
}

// imported cluster might cause blank node_name which is allowed
func removeEntryWithBlankNodeNameFromOldList(oldNodeResourceList []interface{}) []interface{} {
	resultList := []interface{}{}
	for _, tmp := range oldNodeResourceList {
		nodeResource := tmp.(map[string]interface{})
		if tmp, nodeNameExist := nodeResource["node_name"]; nodeNameExist {
			nodeName := tmp.(string)
			if nodeName == "" {
				continue
			} else {
				resultList = append(resultList, nodeResource)
			}
		}
	}
	return resultList
}

func customValidationOnNodeResources(context context.Context, resourceDiff *schema.ResourceDiff, meta interface{}) error {
	oldNodeResourcesRaw, newNodeResourcesRaw := resourceDiff.GetChange(getNodeResourceKey())
	oldNodeResourcesRawSet := oldNodeResourcesRaw.(*schema.Set)
	oldNodeResourceList := oldNodeResourcesRawSet.List()

	newNodeResourcesRawSet := newNodeResourcesRaw.(*schema.Set)
	newNodeResourceList := newNodeResourcesRawSet.List()

	// keys are node_name in lower case
	newNodeMap, err := buildNodeNameToNodeResourceMap(newNodeResourceList)
	if err != nil {
		return err
	}

	// keys are node_name in lower case
	oldNodeMap, err := buildNodeNameToNodeResourceMap(oldNodeResourceList)
	if err != nil {
		return err
	}

	// !!! node_name (user input) only lives in Terraform and does not exist or relate to dbnode API response
	// !!! in order to build and maintain a one-to-one mapping between the user-provided node_name and the dbnode data
	// !!! node_name must be UNIQUE and can NOT be changed once set in state
	if len(oldNodeResourceList) < len(newNodeResourceList) {
		// node_name list: ["node1", "node2"] -> ["node1", "node2", "node3"]
		// new list must be a superset of the old list
		// compare old (state) with new (config)  to make sure no name change on existing nodes
		for oldNodeName, _ := range oldNodeMap {
			if _, match := newNodeMap[oldNodeName]; !match {
				return fmt.Errorf("node_name can not be changed, found node_name '%v' was changed", oldNodeName)
			}
		}
	} else if len(oldNodeResourceList) > len(newNodeResourceList) {
		// node_name list (add node case): ["node1", "node2"] -> ["node1"]
		// new list must be a subset or an equal of the old list
		// compare new (config) with old (state) to make sure no name change on existing nodes
		for newNodeName, _ := range newNodeMap {
			if _, match := oldNodeMap[newNodeName]; !match {
				return fmt.Errorf("node_name can not be changed, got a new node_name '%v'", newNodeName)
			}
		}
	} else {
		// node_name list (no change case): ["node1", "node2"] -> ["node1", "node2"] (order change should not result in a diff)
		// new list must be an equal of the old list
		// compare old (state) with new (config) to make sure no name change on existing nodes
		for oldNodeName, _ := range oldNodeMap {
			if _, match := newNodeMap[oldNodeName]; !match {
				return fmt.Errorf("node_name can not be changed, found node_name '%v' was changed", oldNodeName)
			}
		}
	}

	return nil
}

func nodeResourceHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if nodeName, ok := m["node_name"]; ok && nodeName != "" {
		buf.WriteString(fmt.Sprintf("%v-", nodeName))
	}
	if nodeId, ok := m["node_id"]; ok && nodeId != "" {
		buf.WriteString(fmt.Sprintf("%v-", nodeId))
	}
	return utils.GetStringHashcode(buf.String())
}
