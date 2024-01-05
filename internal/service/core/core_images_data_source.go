// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreImages,
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
			"operating_system": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreImageResource()),
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_core.ListImagesSortByTimecreated),
					string(oci_core.ListImagesSortByDisplayname),
				}, false),
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_core.ListImagesSortOrderAsc),
					string(oci_core.ListImagesSortOrderDesc),
				}, false),
			},
		},
	}
}

func readCoreImages(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListImagesResponse
}

func (s *CoreImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreImagesDataSourceCrud) Get() error {
	request := oci_core.ListImagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if operatingSystem, ok := s.D.GetOkExists("operating_system"); ok {
		tmp := operatingSystem.(string)
		request.OperatingSystem = &tmp
	}

	if operatingSystemVersion, ok := s.D.GetOkExists("operating_system_version"); ok {
		tmp := operatingSystemVersion.(string)
		request.OperatingSystemVersion = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ImageLifecycleStateEnum(state.(string))
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		request.SortBy = oci_core.ListImagesSortByEnum(sortBy.(string))
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		request.SortOrder = oci_core.ListImagesSortOrderEnum(sortOrder.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreImagesDataSource-", CoreImagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		image := map[string]interface{}{}

		// The spec marks compartmentId as a required field, but the service doesn't return it for official images.
		if r.CompartmentId != nil {
			image["compartment_id"] = *r.CompartmentId
		}

		if r.AgentFeatures != nil {
			image["agent_features"] = []interface{}{InstanceAgentFeaturesToMap(r.AgentFeatures)}
		} else {
			image["agent_features"] = nil
		}

		if r.BaseImageId != nil {
			image["base_image_id"] = *r.BaseImageId
		}

		if r.BillableSizeInGBs != nil {
			image["billable_size_in_gbs"] = strconv.FormatInt(*r.BillableSizeInGBs, 10)
		}

		if r.CreateImageAllowed != nil {
			image["create_image_allowed"] = *r.CreateImageAllowed
		}

		if r.DefinedTags != nil {
			image["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			image["display_name"] = *r.DisplayName
		}

		image["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			image["id"] = *r.Id
		}

		image["launch_mode"] = r.LaunchMode

		if r.LaunchOptions != nil {
			image["launch_options"] = []interface{}{LaunchOptionsToMap(r.LaunchOptions)}
		} else {
			image["launch_options"] = nil
		}

		image["listing_type"] = r.ListingType

		if r.OperatingSystem != nil {
			image["operating_system"] = *r.OperatingSystem
		}

		if r.OperatingSystemVersion != nil {
			image["operating_system_version"] = *r.OperatingSystemVersion
		}

		if r.SizeInMBs != nil {
			image["size_in_mbs"] = strconv.FormatInt(*r.SizeInMBs, 10)
		}

		image["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			image["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, image)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreImagesDataSource().Schema["images"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("images", resources); err != nil {
		return err
	}

	return nil
}
