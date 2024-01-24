// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CorePublicIpsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCorePublicIps,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lifetime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CorePublicIpResource()),
			},
		},
	}
}

func readCorePublicIps(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CorePublicIpsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListPublicIpsResponse
}

func (s *CorePublicIpsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePublicIpsDataSourceCrud) Get() error {
	request := oci_core.ListPublicIpsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if lifetime, ok := s.D.GetOkExists("lifetime"); ok {
		request.Lifetime = oci_core.ListPublicIpsLifetimeEnum(lifetime.(string))
	}

	if publicIpPoolId, ok := s.D.GetOkExists("public_ip_pool_id"); ok {
		tmp := publicIpPoolId.(string)
		request.PublicIpPoolId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_core.ListPublicIpsScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CorePublicIpsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CorePublicIpsDataSource-", CorePublicIpsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		publicIp := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"scope":          r.Scope,
		}

		if r.AssignedEntityId != nil {
			publicIp["assigned_entity_id"] = *r.AssignedEntityId
		}

		publicIp["assigned_entity_type"] = r.AssignedEntityType

		if r.AvailabilityDomain != nil {
			publicIp["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			publicIp["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
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

		if r.PublicIpPoolId != nil {
			publicIp["public_ip_pool_id"] = *r.PublicIpPoolId
		}

		publicIp["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			publicIp["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, publicIp)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CorePublicIpsDataSource().Schema["public_ips"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("public_ips", resources); err != nil {
		return err
	}

	return nil
}
