// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeAttachmentDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeAttachments,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeAttachmentResource(),
			},
		},
	}
}

func readVolumeAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &VolumeAttachmentDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type VolumeAttachmentDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVolumeAttachments
}

func (s *VolumeAttachmentDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListVolumeAttachmentsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}
	if val, ok := s.D.GetOk("instance_id"); ok {
		opts.InstanceID = val.(string)
	}
	if val, ok := s.D.GetOk("volume_id"); ok {
		opts.VolumeID = val.(string)
	}

	s.Res = &baremetal.ListVolumeAttachments{
		VolumeAttachments: []baremetal.VolumeAttachment{},
	}

	for {
		var list *baremetal.ListVolumeAttachments
		if list, e = s.Client.ListVolumeAttachments(compartmentID, opts); e != nil {
			break
		}

		s.Res.VolumeAttachments = append(s.Res.VolumeAttachments, list.VolumeAttachments...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *VolumeAttachmentDatasourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	// Important, if you don't have an ID, make one up for your datasource
	// or things will end in tears
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}
	for _, v := range s.Res.VolumeAttachments {
		res := map[string]interface{}{
			"attachment_type":     v.AttachmentType,
			"availability_domain": v.AvailabilityDomain,
			"compartment_id":      v.CompartmentID,
			"display_name":        v.DisplayName,
			"id":                  v.ID,
			"instance_id":         v.InstanceID,
			"state":               v.State,
			"time_created":        v.TimeCreated.String(),
			"volume_id":           v.VolumeID,
		}
		resources = append(resources, res)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("volume_attachments", resources); err != nil {
		panic(err)
	}

	return
}
