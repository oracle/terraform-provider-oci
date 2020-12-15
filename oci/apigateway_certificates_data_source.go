// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v31/apigateway"
)

func init() {
	RegisterDatasource("oci_apigateway_certificates", ApigatewayCertificatesDataSource())
}

func ApigatewayCertificatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApigatewayCertificates,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(ApigatewayCertificateResource()),
						},
					},
				},
			},
		},
	}
}

func readApigatewayCertificates(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayCertificatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).apiGatewayClient()

	return ReadResource(sync)
}

type ApigatewayCertificatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.ApiGatewayClient
	Res    *oci_apigateway.ListCertificatesResponse
}

func (s *ApigatewayCertificatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayCertificatesDataSourceCrud) Get() error {
	request := oci_apigateway.ListCertificatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apigateway.CertificateLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "apigateway")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCertificates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApigatewayCertificatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ApigatewayCertificatesDataSource-", ApigatewayCertificatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	certificate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CertificateSummaryToMap(item))
	}
	certificate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, ApigatewayCertificatesDataSource().Schema["certificate_collection"].Elem.(*schema.Resource).Schema)
		certificate["items"] = items
	}

	resources = append(resources, certificate)
	if err := s.D.Set("certificate_collection", resources); err != nil {
		return err
	}

	return nil
}
