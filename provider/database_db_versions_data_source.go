// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"supports_pdb": {
							Type:     schema.TypeBool,
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

func readDbVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DbVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbVersionsResponse
}

func (s *DbVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbVersionsDataSourceCrud) Get() error {
	request := oci_database.ListDbVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if dbSystemShape, ok := s.D.GetOkExists("db_system_shape"); ok {
		tmp := dbSystemShape.(string)
		request.DbSystemShape = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbVersionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbVersion := map[string]interface{}{}

		if r.SupportsPdb != nil {
			dbVersion["supports_pdb"] = *r.SupportsPdb
		}

		if r.Version != nil {
			dbVersion["version"] = *r.Version
		}

		resources = append(resources, dbVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DbVersionsDataSource().Schema["db_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_versions", resources); err != nil {
		panic(err)
	}

	return
}
