package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/options"
	"github.com/hashicorp/terraform/helper/schema"
)

type VolumeAttachmentDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListVolumeAttachments
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
	if s.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]string{}
		for _, v := range s.Res.VolumeAttachments {
			res := map[string]string{
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
		s.D.Set("volume_attachments", resources)
	}
	return
}
