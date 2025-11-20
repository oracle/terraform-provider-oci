// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsaPrivateServiceAccessesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readPsaPrivateServiceAccessWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_service_access_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(PsaPrivateServiceAccessResource()),
						},
					},
				},
			},
		},
	}
}

func readPsaPrivateServiceAccessWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PsaPrivateServiceAccessDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PsaPrivateServiceAccessDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psa.PrivateServiceAccessClient
	Res    *oci_psa.ListPrivateServiceAccessesResponse
}

func (s *PsaPrivateServiceAccessDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaPrivateServiceAccessDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_psa.ListPrivateServiceAccessesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if serviceId, ok := s.D.GetOkExists("service_id"); ok {
		tmp := serviceId.(string)
		request.ServiceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_psa.PrivateServiceAccessLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.ListPrivateServiceAccesses(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateServiceAccesses(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsaPrivateServiceAccessDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsaPrivateServiceAccessDataSource-", PsaPrivateServiceAccessesDataSource(), s.D))
	resources := []map[string]interface{}{}
	privateServiceAcces := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivateServiceAccessSummaryToMap(item, true))
	}
	privateServiceAcces["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsaPrivateServiceAccessesDataSource().Schema["private_service_access_collection"].Elem.(*schema.Resource).Schema)
		privateServiceAcces["items"] = items
	}

	resources = append(resources, privateServiceAcces)
	if err := s.D.Set("private_service_access_collection", resources); err != nil {
		return err
	}

	return nil
}
