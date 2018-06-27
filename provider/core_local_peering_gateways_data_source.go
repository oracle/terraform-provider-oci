// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LocalPeeringGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLocalPeeringGateways,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"local_peering_gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LocalPeeringGatewayResource(),
			},
		},
	}
}

func readLocalPeeringGateways(d *schema.ResourceData, m interface{}) error {
	sync := &LocalPeeringGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type LocalPeeringGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListLocalPeeringGatewaysResponse
}

func (s *LocalPeeringGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LocalPeeringGatewaysDataSourceCrud) Get() error {
	request := oci_core.ListLocalPeeringGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListLocalPeeringGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLocalPeeringGateways(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LocalPeeringGatewaysDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		localPeeringGateway := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"vcn_id":         *r.VcnId,
		}

		if r.DefinedTags != nil {
			localPeeringGateway["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			localPeeringGateway["display_name"] = *r.DisplayName
		}

		localPeeringGateway["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			localPeeringGateway["id"] = *r.Id
		}

		if r.IsCrossTenancyPeering != nil {
			localPeeringGateway["is_cross_tenancy_peering"] = *r.IsCrossTenancyPeering
		}

		if r.PeerAdvertisedCidr != nil {
			localPeeringGateway["peer_advertised_cidr"] = *r.PeerAdvertisedCidr
		}

		localPeeringGateway["peering_status"] = r.PeeringStatus

		if r.PeeringStatusDetails != nil {
			localPeeringGateway["peering_status_details"] = *r.PeeringStatusDetails
		}

		localPeeringGateway["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			localPeeringGateway["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, localPeeringGateway)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LocalPeeringGatewaysDataSource().Schema["local_peering_gateways"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("local_peering_gateways", resources); err != nil {
		panic(err)
	}

	return
}
