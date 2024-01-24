// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageReplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["replication_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FileStorageReplicationResource(), fieldMap, readSingularFileStorageReplication)
}

func readSingularFileStorageReplication(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageReplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.GetReplicationResponse
}

func (s *FileStorageReplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageReplicationDataSourceCrud) Get() error {
	request := oci_file_storage.GetReplicationRequest{}

	if replicationId, ok := s.D.GetOkExists("replication_id"); ok {
		tmp := replicationId.(string)
		request.ReplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.GetReplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FileStorageReplicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeltaProgress != nil {
		s.D.Set("delta_progress", strconv.FormatInt(*s.Res.DeltaProgress, 10))
	}

	s.D.Set("delta_status", s.Res.DeltaStatus)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastSnapshotId != nil {
		s.D.Set("last_snapshot_id", *s.Res.LastSnapshotId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RecoveryPointTime != nil {
		s.D.Set("recovery_point_time", s.Res.RecoveryPointTime.String())
	}

	if s.Res.ReplicationInterval != nil {
		s.D.Set("replication_interval", strconv.FormatInt(*s.Res.ReplicationInterval, 10))
	}

	if s.Res.ReplicationTargetId != nil {
		s.D.Set("replication_target_id", *s.Res.ReplicationTargetId)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
