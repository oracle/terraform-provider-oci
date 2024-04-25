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

func DatabaseManagementExternalExadataStorageGridResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseManagementExternalExadataStorageGrid,
		Read:     readDatabaseManagementExternalExadataStorageGrid,
		Update:   updateDatabaseManagementExternalExadataStorageGrid,
		Delete:   deleteDatabaseManagementExternalExadataStorageGrid,
		Schema: map[string]*schema.Schema{
			// Required
			"external_exadata_storage_grid_id": {
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
							Type:     schema.TypeFloat,
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

func createDatabaseManagementExternalExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseManagementExternalExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseManagementExternalExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageGridResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseManagementExternalExadataStorageGrid(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseManagementExternalExadataStorageGridResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_management.DbManagementClient
	Res                    *oci_database_management.ExternalExadataStorageGrid
	DisableNotFoundRetries bool
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateCreating),
	}
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateActive),
	}
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateDeleting),
	}
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_management.DbmResourceLifecycleStateDeleted),
	}
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) Create() error {
	request := oci_database_management.UpdateExternalExadataStorageGridRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if externalExadataStorageGridId, ok := s.D.GetOkExists("external_exadata_storage_grid_id"); ok {
		tmp := externalExadataStorageGridId.(string)
		request.ExternalExadataStorageGridId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageGrid
	return nil
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) Get() error {
	request := oci_database_management.GetExternalExadataStorageGridRequest{}

	tmp := s.D.Id()
	request.ExternalExadataStorageGridId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.GetExternalExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageGrid
	return nil
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) Update() error {
	request := oci_database_management.UpdateExternalExadataStorageGridRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.ExternalExadataStorageGridId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_management")

	response, err := s.Client.UpdateExternalExadataStorageGrid(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExternalExadataStorageGrid
	return nil
}

func (s *DatabaseManagementExternalExadataStorageGridResourceCrud) SetData() error {
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

	//s.D.Set("resource_type", s.Res.ResourceType)

	if s.Res.ServerCount != nil {
		s.D.Set("server_count", *s.Res.ServerCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	storageServers := []interface{}{}
	for _, item := range s.Res.StorageServers {
		storageServers = append(storageServers, ExternalExadataStorageServerSummaryToMap(item))
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
