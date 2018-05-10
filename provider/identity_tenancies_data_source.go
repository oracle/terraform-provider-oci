// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func TenancyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenancies,
		Schema: map[string]*schema.Schema{
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_region_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readTenancies(d *schema.ResourceData, m interface{}) error {
	sync := &TenanciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type TenanciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetTenancyResponse
}

func (s *TenanciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenanciesDataSourceCrud) Get() error {
	request := oci_identity.GetTenancyRequest{}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetTenancy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenanciesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeRegionKey != nil {
		s.D.Set("home_region_key", *s.Res.HomeRegionKey)
	}

	if s.Res.Id != nil {
		s.D.Set("tenancy_id", *s.Res.Id)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return
}
