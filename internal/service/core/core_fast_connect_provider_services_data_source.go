// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreFastConnectProviderServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreFastConnectProviderServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fast_connect_provider_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bandwith_shape_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customer_asn_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_peering_bgp_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider_service_key_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider_service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_peering_bgp_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"required_total_cross_connects": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"supported_virtual_circuit_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreFastConnectProviderServices(d *schema.ResourceData, m interface{}) error {
	sync := &CoreFastConnectProviderServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreFastConnectProviderServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListFastConnectProviderServicesResponse
}

func (s *CoreFastConnectProviderServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreFastConnectProviderServicesDataSourceCrud) Get() error {
	request := oci_core.ListFastConnectProviderServicesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListFastConnectProviderServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFastConnectProviderServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreFastConnectProviderServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreFastConnectProviderServicesDataSource-", CoreFastConnectProviderServicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		fastConnectProviderService := map[string]interface{}{}

		fastConnectProviderService["bandwith_shape_management"] = r.BandwithShapeManagement

		fastConnectProviderService["customer_asn_management"] = r.CustomerAsnManagement

		if r.Description != nil {
			fastConnectProviderService["description"] = *r.Description
		}

		if r.Id != nil {
			fastConnectProviderService["id"] = *r.Id
		}

		fastConnectProviderService["private_peering_bgp_management"] = r.PrivatePeeringBgpManagement

		if r.ProviderName != nil {
			fastConnectProviderService["provider_name"] = *r.ProviderName
		}

		fastConnectProviderService["provider_service_key_management"] = r.ProviderServiceKeyManagement

		if r.ProviderServiceName != nil {
			fastConnectProviderService["provider_service_name"] = *r.ProviderServiceName
		}

		fastConnectProviderService["public_peering_bgp_management"] = r.PublicPeeringBgpManagement

		if r.RequiredTotalCrossConnects != nil {
			fastConnectProviderService["required_total_cross_connects"] = *r.RequiredTotalCrossConnects
		}

		fastConnectProviderService["supported_virtual_circuit_types"] = r.SupportedVirtualCircuitTypes

		fastConnectProviderService["type"] = r.Type

		resources = append(resources, fastConnectProviderService)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreFastConnectProviderServicesDataSource().Schema["fast_connect_provider_services"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("fast_connect_provider_services", resources); err != nil {
		return err
	}

	return nil
}
