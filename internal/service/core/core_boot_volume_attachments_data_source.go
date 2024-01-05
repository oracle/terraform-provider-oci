// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreBootVolumeAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreBootVolumeAttachments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_volume_id": {
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
			"boot_volume_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"boot_volume_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encryption_in_transit_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreBootVolumeAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreBootVolumeAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListBootVolumeAttachmentsResponse
}

func (s *CoreBootVolumeAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBootVolumeAttachmentsDataSourceCrud) Get() error {
	request := oci_core.ListBootVolumeAttachmentsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		request.BootVolumeId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListBootVolumeAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBootVolumeAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreBootVolumeAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreBootVolumeAttachmentsDataSource-", CoreBootVolumeAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bootVolumeAttachment := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.BootVolumeId != nil {
			bootVolumeAttachment["boot_volume_id"] = *r.BootVolumeId
		}

		if r.DisplayName != nil {
			bootVolumeAttachment["display_name"] = *r.DisplayName
		}

		bootVolumeAttachment["encryption_in_transit_type"] = r.EncryptionInTransitType

		if r.Id != nil {
			bootVolumeAttachment["id"] = *r.Id
		}

		if r.InstanceId != nil {
			bootVolumeAttachment["instance_id"] = *r.InstanceId
		}

		if r.IsPvEncryptionInTransitEnabled != nil {
			bootVolumeAttachment["is_pv_encryption_in_transit_enabled"] = *r.IsPvEncryptionInTransitEnabled
		}

		bootVolumeAttachment["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bootVolumeAttachment["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, bootVolumeAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreBootVolumeAttachmentsDataSource().Schema["boot_volume_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("boot_volume_attachments", resources); err != nil {
		return err
	}

	return nil
}
