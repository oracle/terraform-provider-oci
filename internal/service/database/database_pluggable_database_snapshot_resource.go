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

func DatabasePluggableDatabaseSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabasePluggableDatabaseSnapshot,
		Read:     readDatabasePluggableDatabaseSnapshot,
		Delete:   deleteDatabasePluggableDatabaseSnapshot,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func createDatabasePluggableDatabaseSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabasePluggableDatabaseSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func deleteDatabasePluggableDatabaseSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabasePluggableDatabaseSnapshotResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.PluggableDatabaseSnapshot
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseSnapshotLifecycleStateCreating),
	}
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseSnapshotLifecycleStateAvailable),
	}
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.PluggableDatabaseSnapshotLifecycleStateTerminating),
	}
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.PluggableDatabaseSnapshotLifecycleStateTerminated),
	}
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) Create() error {
	request := oci_database.CreatePluggableDatabaseSnapshotRequest{}

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

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreatePluggableDatabaseSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.PluggableDatabaseSnapshot

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggabledatabasesnapshot", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) Get() error {
	request := oci_database.GetPluggableDatabaseSnapshotRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetPluggableDatabaseSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PluggableDatabaseSnapshot
	return nil
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) Delete() error {
	request := oci_database.DeletePluggableDatabaseSnapshotRequest{}

	tmp := s.D.Id()
	request.PluggableDatabaseSnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeletePluggableDatabaseSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "pluggabledatabasesnapshot", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabasePluggableDatabaseSnapshotResourceCrud) SetData() error {
	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.PluggableDatabaseId != nil {
		s.D.Set("pluggable_database_id", *s.Res.PluggableDatabaseId)
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
