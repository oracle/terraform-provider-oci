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

func DatabaseGiVersionMinorVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseGiVersionMinorVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_gi_version_for_provisioning": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gi_minor_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"grid_image_id": {
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
		},
	}
}

func readDatabaseGiVersionMinorVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseGiVersionMinorVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseGiVersionMinorVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListGiVersionMinorVersionsResponse
}

func (s *DatabaseGiVersionMinorVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseGiVersionMinorVersionsDataSourceCrud) Get() error {
	request := oci_database.ListGiVersionMinorVersionsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isGiVersionForProvisioning, ok := s.D.GetOkExists("is_gi_version_for_provisioning"); ok {
		tmp := isGiVersionForProvisioning.(bool)
		request.IsGiVersionForProvisioning = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if shapeFamily, ok := s.D.GetOkExists("shape_family"); ok {
		request.ShapeFamily = oci_database.ListGiVersionMinorVersionsShapeFamilyEnum(shapeFamily.(string))
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListGiVersionMinorVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGiVersionMinorVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseGiVersionMinorVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseGiVersionMinorVersionsDataSource-", DatabaseGiVersionMinorVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		giVersionMinorVersion := map[string]interface{}{
			"version": *r.Version,
		}

		if r.GridImageId != nil {
			giVersionMinorVersion["grid_image_id"] = *r.GridImageId
		}

		resources = append(resources, giVersionMinorVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseGiVersionMinorVersionsDataSource().Schema["gi_minor_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("gi_minor_versions", resources); err != nil {
		return err
	}

	return nil
}
