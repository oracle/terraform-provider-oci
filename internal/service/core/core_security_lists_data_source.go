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

func CoreSecurityListsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreSecurityLists,
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
			"security_lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreSecurityListResource()),
			},
		},
	}
}

func readCoreSecurityLists(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityListsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreSecurityListsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListSecurityListsResponse
}

func (s *CoreSecurityListsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreSecurityListsDataSourceCrud) Get() error {
	request := oci_core.ListSecurityListsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.SecurityListLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListSecurityLists(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityLists(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreSecurityListsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreSecurityListsDataSource-", CoreSecurityListsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityList := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			securityList["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			securityList["display_name"] = *r.DisplayName
		}

		egressSecurityRules := []interface{}{}
		for _, item := range r.EgressSecurityRules {
			egressSecurityRules = append(egressSecurityRules, EgressSecurityRuleToMap(item))
		}
		securityList["egress_security_rules"] = egressSecurityRules

		securityList["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			securityList["id"] = *r.Id
		}

		ingressSecurityRules := []interface{}{}
		for _, item := range r.IngressSecurityRules {
			ingressSecurityRules = append(ingressSecurityRules, IngressSecurityRuleToMap(item))
		}
		securityList["ingress_security_rules"] = ingressSecurityRules

		securityList["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			securityList["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			securityList["vcn_id"] = *r.VcnId
		}

		resources = append(resources, securityList)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreSecurityListsDataSource().Schema["security_lists"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("security_lists", resources); err != nil {
		return err
	}

	return nil
}
