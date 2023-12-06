// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseAutonomousContainerDatabaseDataguardAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Read:     readDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Update:   updateDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Delete:   deleteDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"peer_autonomous_container_database_display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"peer_cloud_autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"fast_start_fail_over_lag_limit_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"is_automatic_failover_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"peer_autonomous_container_database_backup_config": {
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
						"backup_destination_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vpc_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										ForceNew:  true,
										Sensitive: true,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"peer_autonomous_container_database_compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"standby_maintenance_buffer_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"peer_autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_autonomous_container_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_lifecycle_state": {
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
			"time_last_role_changed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transport_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabaseDataguardAssociation
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Create() error {

	request := oci_database.CreateAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists("fast_start_fail_over_lag_limit_in_seconds"); ok {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		request.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
	}

	if peerAutonomousContainerDatabaseBackupConfig, ok := s.D.GetOkExists("peer_autonomous_container_database_backup_config"); ok {
		if tmpList := peerAutonomousContainerDatabaseBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "peer_autonomous_container_database_backup_config", 0)
			tmp, err := s.mapToPeerAutonomousContainerDatabaseBackupConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PeerAutonomousContainerDatabaseBackupConfig = &tmp
		}
	}

	if peerAutonomousContainerDatabaseCompartmentId, ok := s.D.GetOkExists("peer_autonomous_container_database_compartment_id"); ok {
		tmp := peerAutonomousContainerDatabaseCompartmentId.(string)
		request.PeerAutonomousContainerDatabaseCompartmentId = &tmp
	}

	if peerAutonomousContainerDatabaseDisplayName, ok := s.D.GetOkExists("peer_autonomous_container_database_display_name"); ok {
		tmp := peerAutonomousContainerDatabaseDisplayName.(string)
		request.PeerAutonomousContainerDatabaseDisplayName = &tmp
	}

	if peerCloudAutonomousVmClusterId, ok := s.D.GetOkExists("peer_cloud_autonomous_vm_cluster_id"); ok {
		tmp := peerCloudAutonomousVmClusterId.(string)
		request.PeerCloudAutonomousVmClusterId = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_database.CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum(protectionMode.(string))
	}

	if standbyMaintenanceBufferInDays, ok := s.D.GetOkExists("standby_maintenance_buffer_in_days"); ok {
		tmp := standbyMaintenanceBufferInDays.(int)
		request.StandbyMaintenanceBufferInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousContainerDatabaseDataguardAssociation

	var dgAssociationId *string
	dgAssociationId = response.Id
	s.D.SetId(*dgAssociationId)

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomouscontainerdatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.Set("autonomous_container_database_id", *identifier)
		}
		if err != nil {
		}
	}

	return s.Get()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	} else {
		tmp := s.D.Id()
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabaseDataguardAssociation
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	} else {
		tmp := s.D.Id()
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if fastStartFailOverLagLimitInSeconds, ok := s.D.GetOkExists("fast_start_fail_over_lag_limit_in_seconds"); ok && s.D.HasChange("fast_start_fail_over_lag_limit_in_seconds") {
		tmp := fastStartFailOverLagLimitInSeconds.(int)
		request.FastStartFailOverLagLimitInSeconds = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_database.UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum(protectionMode.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) SetData() error {

	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	if s.Res.FastStartFailOverLagLimitInSeconds != nil {
		s.D.Set("fast_start_fail_over_lag_limit_in_seconds", *s.Res.FastStartFailOverLagLimitInSeconds)
	}

	if s.Res.IsAutomaticFailoverEnabled != nil {
		s.D.Set("is_automatic_failover_enabled", *s.Res.IsAutomaticFailoverEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId != nil {
		s.D.Set("peer_autonomous_container_database_dataguard_association_id", *s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId)
	}

	if s.Res.PeerAutonomousContainerDatabaseId != nil {
		s.D.Set("peer_autonomous_container_database_id", *s.Res.PeerAutonomousContainerDatabaseId)
	}

	s.D.Set("peer_lifecycle_state", s.Res.PeerLifecycleState)

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRoleChanged != nil {
		s.D.Set("time_last_role_changed", s.Res.TimeLastRoleChanged.String())
	}

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	if s.Res.TransportLag != nil {
		s.D.Set("transport_lag", *s.Res.TransportLag)
	}

	return nil
}

func GetAutonomousContainerDatabaseDataguardAssociationCompositeId(autonomousContainerDatabaseDataguardAssociationId string, autonomousContainerDatabaseId string) string {
	autonomousContainerDatabaseDataguardAssociationId = url.PathEscape(autonomousContainerDatabaseDataguardAssociationId)
	autonomousContainerDatabaseId = url.PathEscape(autonomousContainerDatabaseId)
	compositeId := "autonomousContainerDatabases/" + autonomousContainerDatabaseId + "/autonomousContainerDatabaseDataguardAssociations/" + autonomousContainerDatabaseDataguardAssociationId
	return compositeId
}

// func parseAutonomousContainerDatabaseDataguardAssociationCompositeId(compositeId string) (autonomousContainerDatabaseDataguardAssociationId string, autonomousContainerDatabaseId string, err error) {
// 	parts := strings.Split(compositeId, "/")
// 	match, _ := regexp.MatchString("autonomousContainerDatabases/.*/autonomousContainerDatabaseDataguardAssociations/.*", compositeId)
// 	if !match || len(parts) != 4 {
// 		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
// 		return
// 	}
// 	autonomousContainerDatabaseId, _ = url.PathUnescape(parts[1])
// 	autonomousContainerDatabaseDataguardAssociationId, _ = url.PathUnescape(parts[3])
//
// 	return
// }

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if dbrsPolicyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dbrs_policy_id")); ok {
		tmp := dbrsPolicyId.(string)
		result.DbrsPolicyId = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if internetProxy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "internet_proxy")); ok {
		tmp := internetProxy.(string)
		result.InternetProxy = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result.VpcPassword = &tmp
	}

	if vpcUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_user")); ok {
		tmp := vpcUser.(string)
		result.VpcUser = &tmp
	}

	return result, nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) mapToPeerAutonomousContainerDatabaseBackupConfig(fieldKeyFormat string) (oci_database.PeerAutonomousContainerDatabaseBackupConfig, error) {
	result := oci_database.PeerAutonomousContainerDatabaseBackupConfig{}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_database.BackupDestinationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToBackupDestinationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}
