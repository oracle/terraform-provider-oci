// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageObjectStorageLinkDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["object_storage_link_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(LustreFileStorageObjectStorageLinkResource(), fieldMap, readSingularLustreFileStorageObjectStorageLinkWithContext)
}

func readSingularLustreFileStorageObjectStorageLinkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type LustreFileStorageObjectStorageLinkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.GetObjectStorageLinkResponse
}

func (s *LustreFileStorageObjectStorageLinkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageObjectStorageLinkDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.GetObjectStorageLinkRequest{}

	if objectStorageLinkId, ok := s.D.GetOkExists("object_storage_link_id"); ok {
		tmp := objectStorageLinkId.(string)
		request.ObjectStorageLinkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.GetObjectStorageLink(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LustreFileStorageObjectStorageLinkDataSourceCrud) SetData() error {
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

	if s.Res.CurrentJobId != nil {
		s.D.Set("current_job_id", *s.Res.CurrentJobId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSystemPath != nil {
		s.D.Set("file_system_path", *s.Res.FileSystemPath)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOverwrite != nil {
		s.D.Set("is_overwrite", *s.Res.IsOverwrite)
	}

	if s.Res.LastJobId != nil {
		s.D.Set("last_job_id", *s.Res.LastJobId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LustreFileSystemId != nil {
		s.D.Set("lustre_file_system_id", *s.Res.LustreFileSystemId)
	}

	if s.Res.ObjectStoragePrefix != nil {
		s.D.Set("object_storage_prefix", *s.Res.ObjectStoragePrefix)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
