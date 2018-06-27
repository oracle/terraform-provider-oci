// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeAttachments(context.Background(), request)
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
		resources = append(resources, volumeAttachmentToMap(r))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeAttachmentsDataSource().Schema["volume_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_attachments", resources); err != nil {
		panic(err)
	}

	return
}

func volumeAttachmentToMap(r oci_core.VolumeAttachment) map[string]interface{} {
	volumeAttachment := map[string]interface{}{
		"compartment_id": *r.GetCompartmentId(),
	}

	if availabilityDomain := r.GetAvailabilityDomain(); availabilityDomain != nil {
		volumeAttachment["availability_domain"] = *availabilityDomain
	}

	if displayName := r.GetDisplayName(); displayName != nil {
		volumeAttachment["display_name"] = *displayName
	}

	if id := r.GetId(); id != nil {
		volumeAttachment["id"] = *id
	}

	if instanceId := r.GetInstanceId(); instanceId != nil {
		volumeAttachment["instance_id"] = *instanceId
	}

	if isReadOnly := r.GetIsReadOnly(); isReadOnly != nil {
		volumeAttachment["is_read_only"] = *isReadOnly
	}

	volumeAttachment["state"] = string(r.GetLifecycleState())

	if timeCreated := r.GetTimeCreated(); timeCreated != nil {
		volumeAttachment["time_created"] = timeCreated.String()
	}

	if volumeId := r.GetVolumeId(); volumeId != nil {
		volumeAttachment["volume_id"] = *volumeId
	}

	switch typedValue := r.(type) {
	case oci_core.IScsiVolumeAttachment:
		volumeAttachment["attachment_type"] = IScsiVolumeAttachmentDiscriminator

		// IScsiVolumeAttachment-specific fields:
		if typedValue.ChapSecret != nil {
			volumeAttachment["chap_secret"] = *typedValue.ChapSecret
		}

		if typedValue.ChapUsername != nil {
			volumeAttachment["chap_username"] = *typedValue.ChapUsername
		}

		if typedValue.Ipv4 != nil {
			volumeAttachment["ipv4"] = *typedValue.Ipv4
		}

		if typedValue.Iqn != nil {
			volumeAttachment["iqn"] = *typedValue.Iqn
		}

		if typedValue.Port != nil {
			volumeAttachment["port"] = *typedValue.Port
		}
	case oci_core.ParavirtualizedVolumeAttachment:
		volumeAttachment["attachment_type"] = ParavirtualizedVolumeAttachmentDiscriminator
	default:
		volumeAttachment["attachment_type"] = "Unknown"
		log.Printf("[WARNING] Retrieved a volume attachment of unknown type.")
	}

	return volumeAttachment
}
