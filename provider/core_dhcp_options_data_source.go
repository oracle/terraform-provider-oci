// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DhcpOptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDhcpOptionsList,
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
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DhcpOptionsResource(),
			},
		},
	}
}

func readDhcpOptionsList(d *schema.ResourceData, m interface{}) error {
	sync := &DhcpOptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type DhcpOptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDhcpOptionsResponse
}

func (s *DhcpOptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DhcpOptionsDataSourceCrud) Get() error {
	request := oci_core.ListDhcpOptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.DhcpOptionsLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *DhcpOptionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dhcpOptions := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"vcn_id":         *r.VcnId,
		}

		if r.DefinedTags != nil {
			dhcpOptions["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			dhcpOptions["display_name"] = *r.DisplayName
		}

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

		resources = append(resources, dhcpOptions)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DhcpOptionsDataSource().Schema["options"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("options", resources); err != nil {
		panic(err)
	}

	return
}
