// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
	"github.com/hashicorp/terraform/helper/schema"
)

type ImageDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListImages
}

func (s *ImageDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListImagesOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("operating_system"); ok {
		opts.OperatingSystem = val.(string)
	}
	if val, ok := s.D.GetOk("operating_system_version"); ok {
		opts.OperatingSystemVersion = val.(string)
	}

	s.Res = &baremetal.ListImages{Images: []baremetal.Image{}}

	for {
		var list *baremetal.ListImages
		if list, e = s.Client.ListImages(compartmentID, opts); e != nil {
			break
		}

		s.Res.Images = append(s.Res.Images, list.Images...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *ImageDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Images {
			res := map[string]interface{}{
				"base_image_id":            v.BaseImageID,
				"compartment_id":           v.CompartmentID,
				"create_image_allowed":     v.CreateImageAllowed,
				"display_name":             v.DisplayName,
				"id":                       v.ID,
				"state":                    v.State,
				"operating_system":         v.OperatingSystem,
				"operating_system_version": v.OperatingSystemVersion,
				"time_created":             v.TimeCreated.String(),
			}
			resources = append(resources, res)
		}
		s.D.Set("images", resources)
	}
	return
}
