// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DrgsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgResource(),
			},
		},
	}
}

func readDrgs(d *schema.ResourceData, m interface{}) error {
	sync := &DrgsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type DrgsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgsResponse
}

func (s *DrgsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DrgsDataSourceCrud) Get() error {
	request := oci_core.ListDrgsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListDrgs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DrgsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drg := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			drg["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			drg["display_name"] = *r.DisplayName
		}

		drg["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			drg["id"] = *r.Id
		}

		drg["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			drg["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, drg)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DrgsDataSource().Schema["drgs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drgs", resources); err != nil {
		panic(err)
	}

	return
}
