// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package gdp

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GdpGdpPipelinesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readGdpGdpPipelinesWithContext,
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
			"env": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  gdpCommercialCode,
			},
			"gdp_pipeline_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GdpGdpPipelineResource()),
						},
					},
				},
			},
		},
	}
}

func readGdpGdpPipelinesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelinesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()
	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GdpGdpPipelinesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_gdp.GuardedDataPipelineClient
	Res    *oci_gdp.ListGdpPipelinesResponse
}

func (s *GdpGdpPipelinesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GdpGdpPipelinesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_gdp.ListGdpPipelinesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "gdp")

	response, err := s.Client.ListGdpPipelines(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGdpPipelines(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GdpGdpPipelinesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GdpGdpPipelinesDataSource-", GdpGdpPipelinesDataSource(), s.D))
	resources := []map[string]interface{}{}
	gdpPipeline := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, GdpPipelineSummaryToMap(item))
	}
	gdpPipeline["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GdpGdpPipelinesDataSource().Schema["gdp_pipeline_collection"].Elem.(*schema.Resource).Schema)
		gdpPipeline["items"] = items
	}

	resources = append(resources, gdpPipeline)
	if err := s.D.Set("gdp_pipeline_collection", resources); err != nil {
		return err
	}

	return nil
}
