// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseOneoffPatchDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oneoff_patch_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseOneoffPatchResource(), fieldMap, readSingularDatabaseOneoffPatch)
}

func readSingularDatabaseOneoffPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseOneoffPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseOneoffPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetOneoffPatchResponse
}

func (s *DatabaseOneoffPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseOneoffPatchDataSourceCrud) Get() error {
	request := oci_database.GetOneoffPatchRequest{}

	if oneoffPatchId, ok := s.D.GetOkExists("oneoff_patch_id"); ok {
		tmp := oneoffPatchId.(string)
		request.OneoffPatchId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetOneoffPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseOneoffPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("one_off_patches", s.Res.OneOffPatches)
	s.D.Set("one_off_patches", s.Res.OneOffPatches)

	if s.Res.ReleaseUpdate != nil {
		s.D.Set("release_update", *s.Res.ReleaseUpdate)
	}

	if s.Res.Sha256Sum != nil {
		s.D.Set("sha256sum", *s.Res.Sha256Sum)
	}

	if s.Res.SizeInKBs != nil {
		s.D.Set("size_in_kbs", *s.Res.SizeInKBs)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfExpiration != nil {
		s.D.Set("time_of_expiration", s.Res.TimeOfExpiration.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
