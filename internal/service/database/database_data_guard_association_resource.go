// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDataGuardAssociationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseDataGuardAssociation,
		Read:   readDatabaseDataGuardAssociation,
		Update: updateDatabaseDataGuardAssociation,
		Delete: deleteDatabaseDataGuardAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"creation_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ExistingDbSystem",
					"ExistingVmCluster",
					"NewDbSystem",
				}, true),
			},
			"database_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"delete_standby_db_home_on_delete": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"transport_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"create_async": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"backup_network_nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"database_defined_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"database_freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"data_collection_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
							ForceNew: true,
						},
						"is_health_monitoring_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_incident_logs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"database_software_image_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_system_defined_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"db_system_freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fault_domains": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_active_data_guard_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"peer_db_home_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"peer_db_unique_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_sid_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"storage_volume_performance_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"apply_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_data_guard_association_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
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
		},
	}
}

func createDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDataGuardAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DataGuardAssociation
	DisableNotFoundRetries bool
}

func (s *DatabaseDataGuardAssociationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDataGuardAssociationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateProvisioning),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) CreatedTarget() []string {
	if createAsyn, ok := s.D.GetOk("create_async"); ok {
		tmp := createAsyn.(bool)
		if tmp {
			return []string{
				string(oci_database.DataGuardAssociationLifecycleStateAvailable),
				string(oci_database.DataGuardAssociationLifecycleStateProvisioning),
			}
		}
	}
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateAvailable),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateTerminating),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DataGuardAssociationLifecycleStateTerminated),
	}
}

func (s *DatabaseDataGuardAssociationResourceCrud) Create() error {
	request := oci_database.CreateDataGuardAssociationRequest{}
	err := s.populateTopLevelPolymorphicCreateDataGuardAssociationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataGuardAssociation
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) Get() error {
	request := oci_database.GetDataGuardAssociationRequest{}

	tmp := s.D.Id()
	request.DataGuardAssociationId = &tmp

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataGuardAssociation
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) Update() error {
	request := oci_database.UpdateDataGuardAssociationRequest{}

	tmp := s.D.Id()
	request.DataGuardAssociationId = &tmp

	if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
		tmp := databaseAdminPassword.(string)
		request.DatabaseAdminPassword = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if isActiveDataGuardEnabled, ok := s.D.GetOkExists("is_active_data_guard_enabled"); ok {
		tmp := isActiveDataGuardEnabled.(bool)
		request.IsActiveDataGuardEnabled = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_database.UpdateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
	}

	if transportType, ok := s.D.GetOkExists("transport_type"); ok {
		request.TransportType = oci_database.UpdateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataGuardAssociation
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) SetData() error {

	if backupNetworkNsgIds, ok := s.D.GetOkExists("backup_network_nsg_ids"); ok {
		s.D.Set("backup_network_nsg_ids", backupNetworkNsgIds)
	} else {
		s.D.Set("backup_network_nsg_ids", nil)
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		s.D.Set("nsg_ids", nsgIds)
	} else {
		s.D.Set("nsg_ids", nil)
	}

	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.IsActiveDataGuardEnabled != nil {
		s.D.Set("is_active_data_guard_enabled", *s.Res.IsActiveDataGuardEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerDataGuardAssociationId != nil {
		s.D.Set("peer_data_guard_association_id", *s.Res.PeerDataGuardAssociationId)
	}

	if s.Res.PeerDatabaseId != nil {
		s.D.Set("peer_database_id", *s.Res.PeerDatabaseId)
	}

	if s.Res.PeerDbHomeId != nil {
		s.D.Set("peer_db_home_id", *s.Res.PeerDbHomeId)
	}

	if s.Res.PeerDbSystemId != nil {
		s.D.Set("peer_db_system_id", *s.Res.PeerDbSystemId)
	}

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("transport_type", s.Res.TransportType)

	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) mapToDataCollectionOptions(fieldKeyFormat string) (oci_database.DataCollectionOptions, error) {
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

func (s *DatabaseDataGuardAssociationResourceCrud) populateTopLevelPolymorphicCreateDataGuardAssociationRequest(request *oci_database.CreateDataGuardAssociationRequest) error {
	//discriminator
	creationTypeRaw, ok := s.D.GetOkExists("creation_type")
	var creationType string
	if ok {
		creationType = creationTypeRaw.(string)
	} else {
		creationType = "" // default value
	}
	switch strings.ToLower(creationType) {
	case strings.ToLower("ExistingDbSystem"):
		details := oci_database.CreateDataGuardAssociationToExistingDbSystemDetails{}
		if peerDbHomeId, ok := s.D.GetOkExists("peer_db_home_id"); ok {
			tmp := peerDbHomeId.(string)
			details.PeerDbHomeId = &tmp
		}
		if peerDbSystemId, ok := s.D.GetOkExists("peer_db_system_id"); ok {
			tmp := peerDbSystemId.(string)
			details.PeerDbSystemId = &tmp
		}
		if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
			tmp := databaseAdminPassword.(string)
			details.DatabaseAdminPassword = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}

		if isActiveDataGuardEnabled, ok := s.D.GetOkExists("is_active_data_guard_enabled"); ok {
			tmp := isActiveDataGuardEnabled.(bool)
			details.IsActiveDataGuardEnabled = &tmp
		}

		if peerDbUniqueName, ok := s.D.GetOkExists("peer_db_unique_name"); ok {
			tmp := peerDbUniqueName.(string)
			details.PeerDbUniqueName = &tmp
		}
		if peerSidPrefix, ok := s.D.GetOkExists("peer_sid_prefix"); ok {
			tmp := peerSidPrefix.(string)
			details.PeerSidPrefix = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
			details.ProtectionMode = oci_database.CreateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists("transport_type"); ok {
			details.TransportType = oci_database.CreateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
		}
		request.CreateDataGuardAssociationDetails = details
	case strings.ToLower("ExistingVmCluster"):
		details := oci_database.CreateDataGuardAssociationToExistingVmClusterDetails{}
		if peerDbHomeId, ok := s.D.GetOkExists("peer_db_home_id"); ok {
			tmp := peerDbHomeId.(string)
			details.PeerDbHomeId = &tmp
		}
		if peerVmClusterId, ok := s.D.GetOkExists("peer_vm_cluster_id"); ok {
			tmp := peerVmClusterId.(string)
			details.PeerVmClusterId = &tmp
		}
		if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
			tmp := databaseAdminPassword.(string)
			details.DatabaseAdminPassword = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}

		if isActiveDataGuardEnabled, ok := s.D.GetOkExists("is_active_data_guard_enabled"); ok {
			tmp := isActiveDataGuardEnabled.(bool)
			details.IsActiveDataGuardEnabled = &tmp
		}

		if peerDbUniqueName, ok := s.D.GetOkExists("peer_db_unique_name"); ok {
			tmp := peerDbUniqueName.(string)
			details.PeerDbUniqueName = &tmp
		}
		if peerSidPrefix, ok := s.D.GetOkExists("peer_sid_prefix"); ok {
			tmp := peerSidPrefix.(string)
			details.PeerSidPrefix = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
			details.ProtectionMode = oci_database.CreateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists("transport_type"); ok {
			details.TransportType = oci_database.CreateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
		}
		request.CreateDataGuardAssociationDetails = details
	case strings.ToLower("NewDbSystem"):
		details := oci_database.CreateDataGuardAssociationWithNewDbSystemDetails{}
		if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
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
				details.BackupNetworkNsgIds = tmp
			}
		}
		if cpuCoreCount, ok := s.D.GetOkExists("cpu_core_count"); ok {
			tmp := cpuCoreCount.(int)
			details.CpuCoreCount = &tmp
		}
		if databaseDefinedTags, ok := s.D.GetOkExists("database_defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(databaseDefinedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DatabaseDefinedTags = convertedDefinedTags
		}
		if databaseFreeformTags, ok := s.D.GetOkExists("database_freeform_tags"); ok {
			details.DatabaseFreeformTags = tfresource.ObjectMapToStringMap(databaseFreeformTags.(map[string]interface{}))
		}
		if dbSystemDefinedTags, ok := s.D.GetOkExists("db_system_defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(dbSystemDefinedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DbSystemDefinedTags = convertedDefinedTags
		}
		if dbSystemFreeformTags, ok := s.D.GetOkExists("db_system_freeform_tags"); ok {
			details.DbSystemFreeformTags = tfresource.ObjectMapToStringMap(dbSystemFreeformTags.(map[string]interface{}))
		}
		if dataCollectionOptions, ok := s.D.GetOkExists("data_collection_options"); ok {
			if tmpList := dataCollectionOptions.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_collection_options", 0)
				tmp, err := s.mapToDataCollectionOptions(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DataCollectionOptions = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if domain, ok := s.D.GetOkExists("domain"); ok {
			tmp := domain.(string)
			details.Domain = &tmp
		}
		if faultDomains, ok := s.D.GetOkExists("fault_domains"); ok {
			interfaces := faultDomains.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("fault_domains") {
				details.FaultDomains = tmp
			}
		}
		if hostname, ok := s.D.GetOkExists("hostname"); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if isActiveDataGuardEnabled, ok := s.D.GetOkExists("is_active_data_guard_enabled"); ok {
			tmp := isActiveDataGuardEnabled.(bool)
			details.IsActiveDataGuardEnabled = &tmp
		}
		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			details.LicenseModel = oci_database.CreateDataGuardAssociationWithNewDbSystemDetailsLicenseModelEnum(licenseModel.(string))
		}
		if nodeCount, ok := s.D.GetOkExists("node_count"); ok {
			tmp := nodeCount.(int)
			details.NodeCount = &tmp
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
				details.NsgIds = tmp
			}
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if shape, ok := s.D.GetOkExists("shape"); ok {
			tmp := shape.(string)
			details.Shape = &tmp
		}
		if storageVolumePerformanceMode, ok := s.D.GetOkExists("storage_volume_performance_mode"); ok {
			details.StorageVolumePerformanceMode = oci_database.CreateDataGuardAssociationWithNewDbSystemDetailsStorageVolumePerformanceModeEnum(storageVolumePerformanceMode.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
			tmp := databaseAdminPassword.(string)
			details.DatabaseAdminPassword = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			request.DatabaseId = &tmp
		}
		if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
			tmp := databaseSoftwareImageId.(string)
			details.DatabaseSoftwareImageId = &tmp
		}
		if peerDbUniqueName, ok := s.D.GetOkExists("peer_db_unique_name"); ok {
			tmp := peerDbUniqueName.(string)
			details.PeerDbUniqueName = &tmp
		}
		if peerSidPrefix, ok := s.D.GetOkExists("peer_sid_prefix"); ok {
			tmp := peerSidPrefix.(string)
			details.PeerSidPrefix = &tmp
		}
		if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
			details.ProtectionMode = oci_database.CreateDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
		}
		if transportType, ok := s.D.GetOkExists("transport_type"); ok {
			details.TransportType = oci_database.CreateDataGuardAssociationDetailsTransportTypeEnum(transportType.(string))
		}
		request.CreateDataGuardAssociationDetails = details
	default:
		return fmt.Errorf("unknown creation_type '%v' was specified", creationType)
	}
	return nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) Delete() error {
	if deleteStandbyDbHomeOnDelete, ok := s.D.GetOkExists("delete_standby_db_home_on_delete"); ok {
		tmp := deleteStandbyDbHomeOnDelete.(string)
		if tmp != "true" {
			return fmt.Errorf("we do not currently support deleting the dataguard association without deleting the standby dbHome. Please set delete_standby_db_home_on_delete to \"true\" if you want to continue with the destroy. Once you change the value of delete_standby_db_home_on_delete you must do a `terraform apply` before running a `terraform destroy` so that destroy operation would get the new value")
		}
	}

	err := s.Get()
	if err != nil {
		return err
	}

	creationType, ok := s.D.GetOkExists("creation_type")
	if !ok {
		return fmt.Errorf("creation_type could not be established during the delete")
	}
	if strings.ToLower(creationType.(string)) == strings.ToLower("ExistingDbSystem") {
		deleteDbHomeRequest := oci_database.DeleteDbHomeRequest{}
		var standbyDbHomeId *string
		if s.Res.PeerRole == oci_database.DataGuardAssociationPeerRoleStandby {
			standbyDbHomeId = s.Res.PeerDbHomeId
		} else if s.Res.Role == oci_database.DataGuardAssociationRoleStandby {
			standbyDbHomeId, _ = s.GetDbHomeIdFromDatabaseId(s.Res.DatabaseId)
		} else {
			return fmt.Errorf("could not delete the dataguard association as it is not possible to determine the standby db home")
		}

		if standbyDbHomeId == nil {
			return fmt.Errorf("could not delete the dataguard association as the standby Db Home Id could not be obtained")
		}

		deleteDbHomeRequest.DbHomeId = standbyDbHomeId
		deleteDbHomeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		if _, err = s.Client.DeleteDbHome(context.Background(), deleteDbHomeRequest); err != nil {
			return fmt.Errorf("failed to delete the standby db home")
		}

		getDbHomeRequest := oci_database.GetDbHomeRequest{}
		getDbHomeRequest.DbHomeId = standbyDbHomeId
		getDbHomeRequest.RequestMetadata.RetryPolicy = waitForDbHomeToTerminateRetryPolicy(2 * time.Hour)
		getDbHomeResponse, err := s.Client.GetDbHome(context.Background(), getDbHomeRequest)

		if getDbHomeResponse.LifecycleState == oci_database.DbHomeLifecycleStateAvailable {
			return fmt.Errorf("could not delete the dataguard association as the standby db home could not be deleted")
		}

		return err
	} else if strings.ToLower(creationType.(string)) == strings.ToLower("NewDbSystem") {
		var standbyDbSystemId *string
		if s.Res.PeerRole == oci_database.DataGuardAssociationPeerRoleStandby {
			standbyDbSystemId = s.Res.PeerDbSystemId
		} else if s.Res.Role == oci_database.DataGuardAssociationRoleStandby {
			standbyDbSystemId, err = s.GetDbSystemIdFromDatabaseId(s.Res.DatabaseId)
			if err != nil {
				return fmt.Errorf("could not delete the dataguard association as the standby DB System Id could not be obtained: %v", err)
			}
		} else {
			return fmt.Errorf("could not delete the dataguard association as it is not possible to determine the standby DB System")
		}

		if standbyDbSystemId == nil {
			return fmt.Errorf("could not delete the dataguard association as the standby DB System Id could not be obtained")
		}

		request := oci_database.TerminateDbSystemRequest{}
		request.DbSystemId = standbyDbSystemId
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		_, err := s.Client.TerminateDbSystem(context.Background(), request)
		if err != nil {
			return fmt.Errorf("could not delete standby DB System to delete the data guard association: %v", err)
		}

		getDbSystemRequest := oci_database.GetDbSystemRequest{}
		getDbSystemRequest.DbSystemId = standbyDbSystemId
		getDbSystemRequest.RequestMetadata.RetryPolicy = waitForDbSystemToTerminateRetryPolicy(2 * time.Hour)
		getDbSystemResponse, err := s.Client.GetDbSystem(context.Background(), getDbSystemRequest)

		if getDbSystemResponse.LifecycleState == oci_database.DbSystemLifecycleStateAvailable {
			return fmt.Errorf("could not delete the dataguard association as the dbSystem could not be deleted")
		}

		return err
	} else if strings.ToLower(creationType.(string)) == strings.ToLower("ExistingVmCluster") {
		deleteDBrequest := oci_database.DeleteDatabaseRequest{}

		var standbyDatabaseId *string
		if s.Res.PeerRole == oci_database.DataGuardAssociationPeerRoleStandby {
			standbyDatabaseId = s.Res.PeerDatabaseId
		} else if s.Res.Role == oci_database.DataGuardAssociationRoleStandby {
			standbyDatabaseId = s.Res.DatabaseId
		} else {
			return fmt.Errorf("could not delete the dataguard association as it is not possible to determine the standby database Id")
		}

		if standbyDatabaseId == nil {
			return fmt.Errorf("could not delete the dataguard association as the standby Database Id could not be obtained")
		}

		deleteDBrequest.DatabaseId = standbyDatabaseId
		deleteDBrequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		if _, err = s.Client.DeleteDatabase(context.Background(), deleteDBrequest); err != nil {
			return fmt.Errorf("failed to delete the standby database")
		}

		getDatabaseRequest := oci_database.GetDatabaseRequest{}
		getDatabaseRequest.DatabaseId = standbyDatabaseId
		getDatabaseRequest.RequestMetadata.RetryPolicy = waitForDatabaseToTerminateRetryPolicy(2 * time.Hour)
		getDatabaseResponse, err := s.Client.GetDatabase(context.Background(), getDatabaseRequest)

		if getDatabaseResponse.LifecycleState == oci_database.DatabaseLifecycleStateAvailable {
			return fmt.Errorf("could not delete the dataguard association as the standby database could not be deleted")
		}

		return err
	}
	return fmt.Errorf("unrecognized creation_type during delete")
}

func (s *DatabaseDataGuardAssociationResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return 10 * time.Millisecond
	}

	return time.Second * 30
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbHomeIdFromDatabaseId(databaseId *string) (*string, error) {
	request := oci_database.GetDatabaseRequest{}

	request.DatabaseId = databaseId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.DbHomeId, nil
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbSystemIdFromDatabaseId(databaseId *string) (*string, error) {
	dbHomeId, err := s.GetDbHomeIdFromDatabaseId(databaseId)
	if err != nil {
		return dbHomeId, err
	}
	return s.GetDbSystemIdFromDbHomeId(dbHomeId)
}

func (s *DatabaseDataGuardAssociationResourceCrud) GetDbSystemIdFromDbHomeId(dbHomeId *string) (*string, error) {
	request := oci_database.GetDbHomeRequest{}

	request.DbHomeId = dbHomeId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

	response, err := s.Client.GetDbHome(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.DbSystemId, nil
}

func waitForDbHomeToTerminateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()

	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if tfresource.ShouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDbHomeResponse, ok := response.Response.(oci_database.GetDbHomeResponse); ok {
				if getDbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateTerminated && getDbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateAvailable {
					timeWaited := tfresource.GetElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return tfresource.GetRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func waitForDbSystemToTerminateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()

	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if tfresource.ShouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDbSystemResponse, ok := response.Response.(oci_database.GetDbSystemResponse); ok {
				if getDbSystemResponse.LifecycleState != oci_database.DbSystemLifecycleStateTerminated && getDbSystemResponse.LifecycleState != oci_database.DbSystemLifecycleStateAvailable {
					timeWaited := tfresource.GetElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return tfresource.GetRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}

func waitForDatabaseToTerminateRetryPolicy(timeout time.Duration) *oci_common.RetryPolicy {
	startTime := time.Now()

	return &oci_common.RetryPolicy{
		ShouldRetryOperation: func(response oci_common.OCIOperationResponse) bool {
			if tfresource.ShouldRetry(response, false, "database", startTime) {
				return true
			}
			if getDatabaseResponse, ok := response.Response.(oci_database.GetDatabaseResponse); ok {
				if getDatabaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateTerminated && getDatabaseResponse.LifecycleState != oci_database.DatabaseLifecycleStateAvailable {
					timeWaited := tfresource.GetElapsedRetryDuration(startTime)
					return timeWaited < timeout
				}
			}
			return false
		},
		NextDuration: func(response oci_common.OCIOperationResponse) time.Duration {
			return tfresource.GetRetryBackoffDuration(response, false, "database", startTime)
		},
		MaximumNumberAttempts: 0,
	}
}
