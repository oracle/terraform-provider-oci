// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardSecurityZonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardSecurityZones,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_required_security_zones_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"security_recipe_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_zone_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudGuardSecurityZoneResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardSecurityZones(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardSecurityZonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListSecurityZonesResponse
}

func (s *CloudGuardSecurityZonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardSecurityZonesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListSecurityZonesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isRequiredSecurityZonesInSubtree, ok := s.D.GetOkExists("is_required_security_zones_in_subtree"); ok {
		tmp := isRequiredSecurityZonesInSubtree.(bool)
		request.IsRequiredSecurityZonesInSubtree = &tmp
	}

	if securityRecipeId, ok := s.D.GetOkExists("security_recipe_id"); ok {
		tmp := securityRecipeId.(string)
		request.SecurityRecipeId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.ListSecurityZonesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListSecurityZones(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityZones(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardSecurityZonesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardSecurityZonesDataSource-", CloudGuardSecurityZonesDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityZone := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SecurityZoneSummaryToMap(item))
	}
	securityZone["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardSecurityZonesDataSource().Schema["security_zone_collection"].Elem.(*schema.Resource).Schema)
		securityZone["items"] = items
	}

	resources = append(resources, securityZone)
	if err := s.D.Set("security_zone_collection", resources); err != nil {
		return err
	}

	return nil
}
