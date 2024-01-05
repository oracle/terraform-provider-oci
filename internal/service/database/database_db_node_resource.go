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

func DatabaseDbNodeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDbNode,
		Read:     readDatabaseDbNode,
		Update:   updateDatabaseDbNode,
		Delete:   deleteDatabaseDbNode,
		Schema: map[string]*schema.Schema{
			// Required
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			/*"public_key": {
				Type:     schema.TypeString,
				Required: true,
			},*/

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"additional_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_vnic2id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_server_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_storage_size_in_gb": {
				Type:     schema.TypeInt,
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
			"time_maintenance_window_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_maintenance_window_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic2id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDbNode(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDbNode(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseDbNode(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseDbNode(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseDbNodeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DbNode
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDbNodeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDbNodeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DbNodeLifecycleStateProvisioning),
		string(oci_database.DbNodeLifecycleStateStarting),
	}
}

func (s *DatabaseDbNodeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DbNodeLifecycleStateAvailable),
	}
}

func (s *DatabaseDbNodeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DbNodeLifecycleStateTerminating),
	}
}

func (s *DatabaseDbNodeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DbNodeLifecycleStateTerminated),
	}
}

func (s *DatabaseDbNodeResourceCrud) Create() error {
	request := oci_database.UpdateDbNodeRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbNode(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.DbNode

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "node", oci_work_requests.WorkRequestResourceActionTypeInProgress, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatabaseDbNodeResourceCrud) Get() error {
	request := oci_database.GetDbNodeRequest{}

	tmp := s.D.Id()
	request.DbNodeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDbNode(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbNode
	return nil
}

func (s *DatabaseDbNodeResourceCrud) Update() error {
	request := oci_database.UpdateDbNodeRequest{}

	tmp := s.D.Id()
	request.DbNodeId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDbNode(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "node", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseDbNodeResourceCrud) SetData() error {
	if s.Res.AdditionalDetails != nil {
		s.D.Set("additional_details", *s.Res.AdditionalDetails)
	}

	if s.Res.BackupIpId != nil {
		s.D.Set("backup_ip_id", *s.Res.BackupIpId)
	}

	if s.Res.BackupVnic2Id != nil {
		s.D.Set("backup_vnic2id", *s.Res.BackupVnic2Id)
	}

	if s.Res.BackupVnicId != nil {
		s.D.Set("backup_vnic_id", *s.Res.BackupVnicId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerId != nil {
		s.D.Set("db_server_id", *s.Res.DbServerId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.FreeformTags != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}

	if s.Res.HostIpId != nil {
		s.D.Set("host_ip_id", *s.Res.HostIpId)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.SoftwareStorageSizeInGB != nil {
		s.D.Set("software_storage_size_in_gb", *s.Res.SoftwareStorageSizeInGB)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaintenanceWindowEnd != nil {
		s.D.Set("time_maintenance_window_end", s.Res.TimeMaintenanceWindowEnd.String())
	}

	if s.Res.TimeMaintenanceWindowStart != nil {
		s.D.Set("time_maintenance_window_start", s.Res.TimeMaintenanceWindowStart.String())
	}

	if s.Res.Vnic2Id != nil {
		s.D.Set("vnic2id", *s.Res.Vnic2Id)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}
	return nil
}
