// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PublicIpsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPublicIps,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     PublicIpResource(),
			},
		},
	}
}

func readPublicIps(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type PublicIpsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListPublicIpsResponse
}

func (s *PublicIpsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PublicIpsDataSourceCrud) Get() error {
	request := oci_core.ListPublicIpsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_core.ListPublicIpsScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListPublicIps(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublicIps(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PublicIpsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		publicIp := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"scope":          r.Scope,
		}

		if r.AvailabilityDomain != nil {
			publicIp["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			publicIp["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			publicIp["display_name"] = *r.DisplayName
		}

		publicIp["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			publicIp["id"] = *r.Id
		}

		if r.IpAddress != nil {
			publicIp["ip_address"] = *r.IpAddress
		}

		publicIp["lifetime"] = r.Lifetime

		if r.PrivateIpId != nil {
			publicIp["private_ip_id"] = *r.PrivateIpId
		}

		publicIp["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			publicIp["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, publicIp)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, PublicIpsDataSource().Schema["public_ips"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("public_ips", resources); err != nil {
		panic(err)
	}

	return
}
