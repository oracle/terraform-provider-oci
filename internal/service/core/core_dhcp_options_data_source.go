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

func CoreDhcpOptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDhcpOptionsList,
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
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDhcpOptionsResource()),
			},
		},
	}
}

func readCoreDhcpOptionsList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDhcpOptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDhcpOptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDhcpOptionsResponse
}

func (s *CoreDhcpOptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDhcpOptionsDataSourceCrud) Get() error {
	request := oci_core.ListDhcpOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.DhcpOptionsLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDhcpOptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDhcpOptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDhcpOptionsDataSource-", CoreDhcpOptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dhcpOptions := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			dhcpOptions["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			dhcpOptions["display_name"] = *r.DisplayName
		}

		dhcpOptions["domain_name_type"] = r.DomainNameType

		dhcpOptions["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dhcpOptions["id"] = *r.Id
		}

		options := []interface{}{}
		for _, item := range r.Options {
			options = append(options, DhcpOptionToMap(item))
		}
		dhcpOptions["options"] = options

		dhcpOptions["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dhcpOptions["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			dhcpOptions["vcn_id"] = *r.VcnId
		}

		resources = append(resources, dhcpOptions)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDhcpOptionsDataSource().Schema["options"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("options", resources); err != nil {
		return err
	}

	return nil
}
