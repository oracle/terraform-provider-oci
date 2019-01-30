// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
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
			"volume_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(VolumeAttachmentResource()),
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: FieldDeprecated("page"),
			},
		},
	}
}

func readVolumeAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
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

	if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
		tmp := volumeId.(string)
		request.VolumeId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
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

func (s *VolumeAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_core.IScsiVolumeAttachment:
			result["attachment_type"] = "iscsi"

			if v.ChapSecret != nil {
				result["chap_secret"] = string(*v.ChapSecret)
			}

			if v.ChapUsername != nil {
				result["chap_username"] = string(*v.ChapUsername)
			}

			if v.Ipv4 != nil {
				result["ipv4"] = string(*v.Ipv4)
			}

			if v.Iqn != nil {
				result["iqn"] = string(*v.Iqn)
			}

			if v.Port != nil {
				result["port"] = int(*v.Port)
			}

			if v.AvailabilityDomain != nil {
				result["availability_domain"] = string(*v.AvailabilityDomain)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.Device != nil {
				result["device"] = string(*v.Device)
			}

			if v.DisplayName != nil {
				result["display_name"] = string(*v.DisplayName)
			}

			if v.Id != nil {
				result["id"] = string(*v.Id)
			}

			if v.InstanceId != nil {
				result["instance_id"] = string(*v.InstanceId)
			}

			if v.IsPvEncryptionInTransitEnabled != nil {
				result["is_pv_encryption_in_transit_enabled"] = bool(*v.IsPvEncryptionInTransitEnabled)
			}

			if v.IsReadOnly != nil {
				result["is_read_only"] = bool(*v.IsReadOnly)
			}

			result["state"] = string(v.LifecycleState)

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}

			if v.VolumeId != nil {
				result["volume_id"] = string(*v.VolumeId)
			}
		case oci_core.ParavirtualizedVolumeAttachment:
			result["attachment_type"] = "paravirtualized"

			if v.AvailabilityDomain != nil {
				result["availability_domain"] = string(*v.AvailabilityDomain)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.Device != nil {
				result["device"] = string(*v.Device)
			}

			if v.DisplayName != nil {
				result["display_name"] = string(*v.DisplayName)
			}

			if v.Id != nil {
				result["id"] = string(*v.Id)
			}

			if v.InstanceId != nil {
				result["instance_id"] = string(*v.InstanceId)
			}

			if v.IsPvEncryptionInTransitEnabled != nil {
				result["is_pv_encryption_in_transit_enabled"] = bool(*v.IsPvEncryptionInTransitEnabled)
			}

			if v.IsReadOnly != nil {
				result["is_read_only"] = bool(*v.IsReadOnly)
			}

			result["state"] = string(v.LifecycleState)

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}

			if v.VolumeId != nil {
				result["volume_id"] = string(*v.VolumeId)
			}
		default:
			log.Printf("[WARN] Received 'attachment_type' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeAttachmentsDataSource().Schema["volume_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_attachments", resources); err != nil {
		return err
	}

	return nil
}
