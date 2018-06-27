// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VnicAttachmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVnicAttachments,
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
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VnicAttachmentResource(),
			},
		},
	}
}

func readVnicAttachments(d *schema.ResourceData, m interface{}) error {
	sync := &VnicAttachmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type VnicAttachmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListVnicAttachmentsResponse
}

func (s *VnicAttachmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VnicAttachmentsDataSourceCrud) Get() error {
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

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

func (s *VnicAttachmentsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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

		if r.VlanTag != nil {
			vnicAttachment["vlan_tag"] = *r.VlanTag
		}

		if r.VnicId != nil {
			vnicAttachment["vnic_id"] = *r.VnicId
		}

		resources = append(resources, vnicAttachment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VnicAttachmentsDataSource().Schema["vnic_attachments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vnic_attachments", resources); err != nil {
		panic(err)
	}

	return
}
