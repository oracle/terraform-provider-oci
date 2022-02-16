// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package sch

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_sch "github.com/oracle/oci-go-sdk/v58/sch"
)

func SchServiceConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSchServiceConnectors,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(SchServiceConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readSchServiceConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceConnectorClient()

	return tfresource.ReadResource(sync)
}

type SchServiceConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_sch.ServiceConnectorClient
	Res    *oci_sch.ListServiceConnectorsResponse
}

func (s *SchServiceConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SchServiceConnectorsDataSourceCrud) Get() error {
	request := oci_sch.ListServiceConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_sch.ListServiceConnectorsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "sch")

	response, err := s.Client.ListServiceConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SchServiceConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SchServiceConnectorsDataSource-", SchServiceConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	serviceConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceConnectorSummaryToMap(item))
	}
	serviceConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, SchServiceConnectorsDataSource().Schema["service_connector_collection"].Elem.(*schema.Resource).Schema)
		serviceConnector["items"] = items
	}

	resources = append(resources, serviceConnector)
	if err := s.D.Set("service_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
