// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package desktops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DesktopsDesktopPoolVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDesktopsDesktopPoolVolumes,
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
			"desktop_pool_id": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"desktop_pool_volume_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pool_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDesktopsDesktopPoolVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.ReadResource(sync)
}

type DesktopsDesktopPoolVolumesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_desktops.DesktopServiceClient
	Res    *oci_desktops.ListDesktopPoolVolumesResponse
}

func (s *DesktopsDesktopPoolVolumesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DesktopsDesktopPoolVolumesDataSourceCrud) Get() error {
	request := oci_desktops.ListDesktopPoolVolumesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if desktopPoolId, ok := s.D.GetOkExists("desktop_pool_id"); ok {
		tmp := desktopPoolId.(string)
		request.DesktopPoolId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "desktops")

	response, err := s.Client.ListDesktopPoolVolumes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDesktopPoolVolumes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DesktopsDesktopPoolVolumesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DesktopsDesktopPoolVolumesDataSource-", DesktopsDesktopPoolVolumesDataSource(), s.D))
	resources := []map[string]interface{}{}
	desktopPoolVolume := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DesktopPoolVolumeSummaryToMap(item))
	}
	desktopPoolVolume["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DesktopsDesktopPoolVolumesDataSource().Schema["desktop_pool_volume_collection"].Elem.(*schema.Resource).Schema)
		desktopPoolVolume["items"] = items
	}

	resources = append(resources, desktopPoolVolume)
	if err := s.D.Set("desktop_pool_volume_collection", resources); err != nil {
		return err
	}

	return nil
}

func DesktopPoolVolumeSummaryToMap(obj oci_desktops.DesktopPoolVolumeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PoolId != nil {
		result["pool_id"] = string(*obj.PoolId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
