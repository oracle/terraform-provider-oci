// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v27/database"
)

func init() {
	RegisterDatasource("oci_database_database_software_image", DatabaseDatabaseSoftwareImageDataSource())
}

func DatabaseDatabaseSoftwareImageDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_software_image_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseDatabaseSoftwareImageResource(), fieldMap, readSingularDatabaseDatabaseSoftwareImage)
}

func readSingularDatabaseDatabaseSoftwareImage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseDatabaseSoftwareImageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDatabaseSoftwareImageResponse
}

func (s *DatabaseDatabaseSoftwareImageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabaseSoftwareImageDataSourceCrud) Get() error {
	request := oci_database.GetDatabaseSoftwareImageRequest{}

	if databaseSoftwareImageId, ok := s.D.GetOkExists("database_software_image_id"); ok {
		tmp := databaseSoftwareImageId.(string)
		request.DatabaseSoftwareImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetDatabaseSoftwareImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDatabaseSoftwareImageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_software_image_included_patches", s.Res.DatabaseSoftwareImageIncludedPatches)

	s.D.Set("database_software_image_one_off_patches", s.Res.DatabaseSoftwareImageOneOffPatches)

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("image_shape_family", s.Res.ImageShapeFamily)

	s.D.Set("image_type", s.Res.ImageType)

	if s.Res.IncludedPatchesSummary != nil {
		s.D.Set("included_patches_summary", *s.Res.IncludedPatchesSummary)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LsInventory != nil {
		s.D.Set("ls_inventory", *s.Res.LsInventory)
	}

	if s.Res.PatchSet != nil {
		s.D.Set("patch_set", *s.Res.PatchSet)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
