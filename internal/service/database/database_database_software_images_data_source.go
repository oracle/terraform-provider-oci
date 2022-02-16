// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseDatabaseSoftwareImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDatabaseSoftwareImages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_shape_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_upgrade_supported": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_software_images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseDatabaseSoftwareImageResource()),
			},
		},
	}
}

func readDatabaseDatabaseSoftwareImages(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSoftwareImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDatabaseSoftwareImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDatabaseSoftwareImagesResponse
}

func (s *DatabaseDatabaseSoftwareImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabaseSoftwareImagesDataSourceCrud) Get() error {
	request := oci_database.ListDatabaseSoftwareImagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if imageShapeFamily, ok := s.D.GetOkExists("image_shape_family"); ok {
		request.ImageShapeFamily = oci_database.DatabaseSoftwareImageSummaryImageShapeFamilyEnum(imageShapeFamily.(string))
	}

	if imageType, ok := s.D.GetOkExists("image_type"); ok {
		request.ImageType = oci_database.DatabaseSoftwareImageSummaryImageTypeEnum(imageType.(string))
	}

	if isUpgradeSupported, ok := s.D.GetOkExists("is_upgrade_supported"); ok {
		tmp := isUpgradeSupported.(bool)
		request.IsUpgradeSupported = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DatabaseSoftwareImageSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDatabaseSoftwareImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseSoftwareImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDatabaseSoftwareImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDatabaseSoftwareImagesDataSource-", DatabaseDatabaseSoftwareImagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		databaseSoftwareImage := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		databaseSoftwareImage["database_software_image_included_patches"] = r.DatabaseSoftwareImageIncludedPatches

		databaseSoftwareImage["database_software_image_one_off_patches"] = r.DatabaseSoftwareImageOneOffPatches

		if r.DatabaseVersion != nil {
			databaseSoftwareImage["database_version"] = *r.DatabaseVersion
		}

		if r.DefinedTags != nil {
			databaseSoftwareImage["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			databaseSoftwareImage["display_name"] = *r.DisplayName
		}

		databaseSoftwareImage["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			databaseSoftwareImage["id"] = *r.Id
		}

		databaseSoftwareImage["image_shape_family"] = r.ImageShapeFamily

		databaseSoftwareImage["image_type"] = r.ImageType

		if r.IncludedPatchesSummary != nil {
			databaseSoftwareImage["included_patches_summary"] = *r.IncludedPatchesSummary
		}

		if r.IsUpgradeSupported != nil {
			databaseSoftwareImage["is_upgrade_supported"] = *r.IsUpgradeSupported
		}

		if r.LifecycleDetails != nil {
			databaseSoftwareImage["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.LsInventory != nil {
			databaseSoftwareImage["ls_inventory"] = *r.LsInventory
		}

		if r.PatchSet != nil {
			databaseSoftwareImage["patch_set"] = *r.PatchSet
		}

		databaseSoftwareImage["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			databaseSoftwareImage["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, databaseSoftwareImage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDatabaseSoftwareImagesDataSource().Schema["database_software_images"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("database_software_images", resources); err != nil {
		return err
	}

	return nil
}
