// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDatabaseSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseDatabaseSnapshotWithContext,
		ReadContext:   readDatabaseDatabaseSnapshotWithContext,
		UpdateContext: updateDatabaseDatabaseSnapshotWithContext,
		DeleteContext: deleteDatabaseDatabaseSnapshotWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
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
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func createDatabaseDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseDatabaseSnapshotResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.DatabaseSnapshot
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDatabaseSnapshotResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseSnapshotResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseSnapshotLifecycleStateCreating),
	}
}

func (s *DatabaseDatabaseSnapshotResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseSnapshotLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseSnapshotResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DatabaseSnapshotLifecycleStateTerminating),
	}
}

func (s *DatabaseDatabaseSnapshotResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DatabaseSnapshotLifecycleStateTerminated),
	}
}

func (s *DatabaseDatabaseSnapshotResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database.CreateDatabaseSnapshotRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
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

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.DatabaseSnapshot

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "databasesnapshot", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.GetWithContext(ctx)
}

func (s *DatabaseDatabaseSnapshotResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetDatabaseSnapshotRequest{}

	tmp := s.D.Id()
	request.DatabaseSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseSnapshot
	return nil
}

func (s *DatabaseDatabaseSnapshotResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_database.UpdateDatabaseSnapshotRequest{}

	tmp := s.D.Id()
	request.DatabaseSnapshotId = &tmp

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

	response, err := s.Client.UpdateDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseSnapshot
	return nil
}

func (s *DatabaseDatabaseSnapshotResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database.DeleteDatabaseSnapshotRequest{}

	tmp := s.D.Id()
	request.DatabaseSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "databasesnapshot", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseDatabaseSnapshotResourceCrud) SetData() error {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func DatabaseSnapshotSummaryToMap(obj oci_database.DatabaseSnapshotSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClusterId != nil {
		result["cluster_id"] = string(*obj.ClusterId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
