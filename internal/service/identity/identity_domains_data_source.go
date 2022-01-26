// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomains,
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
			"home_region_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_hidden_on_login": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainResource()),
			},
		},
	}
}

func readIdentityDomains(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListDomainsResponse
}

func (s *IdentityDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsDataSourceCrud) Get() error {
	request := oci_identity.ListDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if homeRegionUrl, ok := s.D.GetOkExists("home_region_url"); ok {
		tmp := homeRegionUrl.(string)
		request.HomeRegionUrl = &tmp
	}

	if isHiddenOnLogin, ok := s.D.GetOkExists("is_hidden_on_login"); ok {
		tmp := isHiddenOnLogin.(bool)
		request.IsHiddenOnLogin = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		tmp := licenseType.(string)
		request.LicenseType = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.DomainLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	if url, ok := s.D.GetOkExists("url"); ok {
		tmp := url.(string)
		request.Url = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDomains(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityDomainsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsDataSource-", IdentityDomainsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		domain := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			domain["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			domain["description"] = *r.Description
		}

		if r.DisplayName != nil {
			domain["display_name"] = *r.DisplayName
		}

		domain["freeform_tags"] = r.FreeformTags

		if r.HomeRegion != nil {
			domain["home_region"] = *r.HomeRegion
		}

		if r.HomeRegionUrl != nil {
			domain["home_region_url"] = *r.HomeRegionUrl
		}

		if r.Id != nil {
			domain["id"] = *r.Id
		}

		if r.IsHiddenOnLogin != nil {
			domain["is_hidden_on_login"] = *r.IsHiddenOnLogin
		}

		if r.LicenseType != nil {
			domain["license_type"] = *r.LicenseType
		}

		domain["lifecycle_details"] = r.LifecycleDetails

		replicaRegions := []interface{}{}
		for _, item := range r.ReplicaRegions {
			replicaRegions = append(replicaRegions, ReplicatedRegionDetailsToMap(item))
		}
		domain["replica_regions"] = replicaRegions

		domain["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			domain["time_created"] = r.TimeCreated.String()
		}

		domain["type"] = r.Type

		if r.Url != nil {
			domain["url"] = *r.Url
		}

		resources = append(resources, domain)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityDomainsDataSource().Schema["domains"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("domains", resources); err != nil {
		return err
	}

	return nil
}
