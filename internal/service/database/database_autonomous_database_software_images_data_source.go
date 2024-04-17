// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabaseSoftwareImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseSoftwareImages,
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
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_database_software_image_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousDatabaseSoftwareImageResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseSoftwareImages(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSoftwareImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseSoftwareImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseSoftwareImagesResponse
}

func (s *DatabaseAutonomousDatabaseSoftwareImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseSoftwareImagesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseSoftwareImagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if imageShapeFamily, ok := s.D.GetOkExists("image_shape_family"); ok {
		request.ImageShapeFamily = oci_database.AutonomousDatabaseSoftwareImageImageShapeFamilyEnum(imageShapeFamily.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseSoftwareImageLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseSoftwareImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseSoftwareImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseSoftwareImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseSoftwareImagesDataSource-", DatabaseAutonomousDatabaseSoftwareImagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousDatabaseSoftwareImage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AutonomousDatabaseSoftwareImageSummaryToMap(item))
	}
	autonomousDatabaseSoftwareImage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousDatabaseSoftwareImagesDataSource().Schema["autonomous_database_software_image_collection"].Elem.(*schema.Resource).Schema)
		autonomousDatabaseSoftwareImage["items"] = items
	}

	resources = append(resources, autonomousDatabaseSoftwareImage)
	if err := s.D.Set("autonomous_database_software_image_collection", resources); err != nil {
		return err
	}

	return nil
}
