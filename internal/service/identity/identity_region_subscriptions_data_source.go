// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityRegionSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityRegionSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region_subscriptions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tenancy_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"is_home_region": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"region_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityRegionSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityRegionSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityRegionSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListRegionSubscriptionsResponse
}

func (s *IdentityRegionSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityRegionSubscriptionsDataSourceCrud) Get() error {
	request := oci_identity.ListRegionSubscriptionsRequest{}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListRegionSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityRegionSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityRegionSubscriptionsDataSource-", IdentityRegionSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		regionSubscription := map[string]interface{}{}

		if r.IsHomeRegion != nil {
			regionSubscription["is_home_region"] = *r.IsHomeRegion
		}

		if r.RegionKey != nil {
			regionSubscription["region_key"] = *r.RegionKey
		}

		if r.RegionName != nil {
			regionSubscription["region_name"] = *r.RegionName
		}

		regionSubscription["state"] = r.Status

		resources = append(resources, regionSubscription)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityRegionSubscriptionsDataSource().Schema["region_subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("region_subscriptions", resources); err != nil {
		return err
	}

	return nil
}
