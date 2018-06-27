// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readImages,
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
			"operating_system": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
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
				Elem:     ImageResource(),
			},
		},
	}
}

func readImages(d *schema.ResourceData, m interface{}) error {
	sync := &ImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type ImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListImagesResponse
}

func (s *ImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ImagesDataSourceCrud) Get() error {
	request := oci_core.ListImagesRequest{}

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

	if operatingSystem, ok := s.D.GetOkExists("operating_system"); ok {
		tmp := operatingSystem.(string)
		request.OperatingSystem = &tmp
	}

	if operatingSystemVersion, ok := s.D.GetOkExists("operating_system_version"); ok {
		tmp := operatingSystemVersion.(string)
		request.OperatingSystemVersion = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ImageLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *ImagesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		image := map[string]interface{}{}

		// The spec marks compartmentId as a required field, but the service doesn't return it for official images.
		if r.CompartmentId != nil {
			image["compartment_id"] = *r.CompartmentId
		}

		if r.BaseImageId != nil {
			image["base_image_id"] = *r.BaseImageId
		}

		if r.CreateImageAllowed != nil {
			image["create_image_allowed"] = *r.CreateImageAllowed
		}

		if r.DefinedTags != nil {
			image["defined_tags"] = definedTagsToMap(r.DefinedTags)
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
		}

		if r.OperatingSystem != nil {
			image["operating_system"] = *r.OperatingSystem
		}

		if r.OperatingSystemVersion != nil {
			image["operating_system_version"] = *r.OperatingSystemVersion
		}

		if r.SizeInMBs != nil {
			image["size_in_mbs"] = *r.SizeInMBs
		}

		image["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			image["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, image)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ImagesDataSource().Schema["images"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("images", resources); err != nil {
		panic(err)
	}

	return
}
