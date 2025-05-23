// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementCloudExadataStorageGridResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementCloudExadataStorageGrid,
		Read:     readDatabaseManagementCloudExadataStorageGrid,
		Update:   updateDatabaseManagementCloudExadataStorageGrid,
		Delete:   deleteDatabaseManagementCloudExadataStorageGrid,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_exadata_storage_grid_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

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
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_count": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"additional_details": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"connector_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_count": {
							Type:     schema.TypeInt,
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
						"internal_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"make_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_flash_disk_iops": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_flash_disk_throughput": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_hard_disk_iops": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_hard_disk_throughput": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory_gb": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
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
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseManagementCloudExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementCloudExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementCloudExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementCloudExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementCloudExadataStorageGridResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.CloudExadataStorageGrid
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.LifecycleStatesCreating),
	}
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateActive),
	}
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) Create() error {
	request := oci_database_management.UpdateCloudExadataStorageGridRequest{}

	if cloudExadataStorageGridId, ok := s.D.GetOkExists("id"); ok {
		tmp := cloudExadataStorageGridId.(string)
		request.CloudExadataStorageGridId = &tmp
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

	if cloudExadataStorageGridId, ok := s.D.GetOkExists("cloud_exadata_storage_grid_id"); ok {
		tmp := cloudExadataStorageGridId.(string)
		request.CloudExadataStorageGridId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataStorageGrid
	return nil
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) Get() error {
	request := oci_database_management.GetCloudExadataStorageGridRequest{}

	tmp := s.D.Id()
	request.CloudExadataStorageGridId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetCloudExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataStorageGrid
	return nil
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) Update() error {
	request := oci_database_management.UpdateCloudExadataStorageGridRequest{}

	tmp := s.D.Id()
	request.CloudExadataStorageGridId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateCloudExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudExadataStorageGrid
	return nil
}

func (s *DatabaseManagementCloudExadataStorageGridResourceCrud) SetData() error {
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataInfrastructureId != nil {
		s.D.Set("exadata_infrastructure_id", *s.Res.ExadataInfrastructureId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InternalId != nil {
		s.D.Set("internal_id", *s.Res.InternalId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ServerCount != nil {
		s.D.Set("server_count", *s.Res.ServerCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	storageServers := []interface{}{}
	for _, item := range s.Res.StorageServers {
		storageServers = append(storageServers, CloudExadataStorageServerSummaryToMap(item))
	}
	s.D.Set("storage_servers", storageServers)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func CloudExadataStorageServerSummaryToMap(obj oci_database_management.CloudExadataStorageServerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_details"] = obj.AdditionalDetails

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	if obj.CpuCount != nil {
		result["cpu_count"] = int(*obj.CpuCount)
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

	if obj.InternalId != nil {
		result["internal_id"] = string(*obj.InternalId)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MakeModel != nil {
		result["make_model"] = string(*obj.MakeModel)
	}

	if obj.MaxFlashDiskIOPS != nil {
		result["max_flash_disk_iops"] = int(*obj.MaxFlashDiskIOPS)
	}

	if obj.MaxFlashDiskThroughput != nil {
		result["max_flash_disk_throughput"] = int(*obj.MaxFlashDiskThroughput)
	}

	if obj.MaxHardDiskIOPS != nil {
		result["max_hard_disk_iops"] = int(*obj.MaxHardDiskIOPS)
	}

	if obj.MaxHardDiskThroughput != nil {
		result["max_hard_disk_throughput"] = int(*obj.MaxHardDiskThroughput)
	}

	if obj.MemoryGB != nil {
		result["memory_gb"] = float64(*obj.MemoryGB)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
