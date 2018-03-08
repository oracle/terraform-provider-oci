// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeAttachmentsDataSource() *schema.Resource {
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
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
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

func readVolumeAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type VolumeAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListVolumeAttachmentsResponse
}

func (s *VolumeAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VolumeAttachmentsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeAttachmentsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
		tmp := volumeId.(string)
		request.VolumeId = &tmp
	}

	response, err := s.Client.ListVolumeAttachments(context.Background(), request, getRetryOptions(false, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeAttachments(context.Background(), request, getRetryOptions(false, "core")...)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VolumeAttachmentsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		iscsiVolumeAttachment, castOk := r.(oci_core.IScsiVolumeAttachment)
		if !castOk {
			panic("unexpected VolumeAttachment type on item.")
		}

		volumeAttachment := map[string]interface{}{
			"compartment_id":  *r.GetCompartmentId(),
			"attachment_type": IScsiVolumeAttachmentDiscriminator,
		}

		if iscsiVolumeAttachment.AvailabilityDomain != nil {
			volumeAttachment["availability_domain"] = *iscsiVolumeAttachment.AvailabilityDomain
		}

		if iscsiVolumeAttachment.DisplayName != nil {
			volumeAttachment["display_name"] = *iscsiVolumeAttachment.DisplayName
		}

		if iscsiVolumeAttachment.Id != nil {
			volumeAttachment["id"] = *iscsiVolumeAttachment.Id
		}

		if iscsiVolumeAttachment.InstanceId != nil {
			volumeAttachment["instance_id"] = *iscsiVolumeAttachment.InstanceId
		}

		volumeAttachment["state"] = iscsiVolumeAttachment.LifecycleState

		volumeAttachment["time_created"] = iscsiVolumeAttachment.TimeCreated.String()

		if iscsiVolumeAttachment.VolumeId != nil {
			volumeAttachment["volume_id"] = *iscsiVolumeAttachment.VolumeId
		}

		// IScsiVolumeAttachment-specific fields:
		if iscsiVolumeAttachment.ChapSecret != nil {
			volumeAttachment["chap_secret"] = *iscsiVolumeAttachment.ChapSecret
		}

		if iscsiVolumeAttachment.ChapUsername != nil {
			volumeAttachment["chap_username"] = *iscsiVolumeAttachment.ChapUsername
		}

		if iscsiVolumeAttachment.Ipv4 != nil {
			volumeAttachment["ipv4"] = *iscsiVolumeAttachment.Ipv4
		}

		if iscsiVolumeAttachment.Iqn != nil {
			volumeAttachment["iqn"] = *iscsiVolumeAttachment.Iqn
		}

		if iscsiVolumeAttachment.Port != nil {
			volumeAttachment["port"] = *iscsiVolumeAttachment.Port
		}

		resources = append(resources, volumeAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("volume_attachments", resources); err != nil {
		panic(err)
	}

	return
}
