// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreVnicAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVnicAttachments,
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
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVnicAttachmentResource()),
			},
		},
	}
}

func readCoreVnicAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVnicAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreVnicAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListVnicAttachmentsResponse
}

func (s *CoreVnicAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVnicAttachmentsDataSourceCrud) Get() error {
	request := oci_core.ListVnicAttachmentsRequest{}

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

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVnicAttachments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVnicAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVnicAttachmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVnicAttachmentsDataSource-", CoreVnicAttachmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vnicAttachment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			vnicAttachment["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DisplayName != nil {
			vnicAttachment["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			vnicAttachment["id"] = *r.Id
		}

		if r.InstanceId != nil {
			vnicAttachment["instance_id"] = *r.InstanceId
		}

		if r.NicIndex != nil {
			vnicAttachment["nic_index"] = *r.NicIndex
		}

		vnicAttachment["state"] = r.LifecycleState

		if r.SubnetId != nil {
			vnicAttachment["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			vnicAttachment["time_created"] = r.TimeCreated.String()
		}

		if r.VlanId != nil {
			vnicAttachment["vlan_id"] = *r.VlanId
		}

		if r.VlanTag != nil {
			vnicAttachment["vlan_tag"] = *r.VlanTag
		}

		if r.VnicId != nil {
			vnicAttachment["vnic_id"] = *r.VnicId
		}

		resources = append(resources, vnicAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVnicAttachmentsDataSource().Schema["vnic_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vnic_attachments", resources); err != nil {
		return err
	}

	return nil
}
