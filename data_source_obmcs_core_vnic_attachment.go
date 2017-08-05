// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func resourceVnicAttachment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func DatasourceCoreVnicAttachments() *schema.Resource {
	return &schema.Resource{
		Read: readVnicAttachments,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vnic_attachments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceVnicAttachment(),
			},
		},
	}
}

func readVnicAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &VnicAttachmentDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type VnicAttachmentDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListVnicAttachments
}

func (r *VnicAttachmentDatasourceCrud) Get() (e error) {
	compartmentID := r.D.Get("compartment_id").(string)

	opts := &baremetal.ListVnicAttachmentsOptions{}
	options.SetListOptions(r.D, &opts.ListOptions)
	if val, ok := r.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}
	if val, ok := r.D.GetOk("instance_id"); ok {
		opts.InstanceID = val.(string)
	}
	if val, ok := r.D.GetOk("vnic_id"); ok {
		opts.VnicID = val.(string)
	}

	r.Res = &baremetal.ListVnicAttachments{
		Attachments: []baremetal.VnicAttachment{},
	}

	for {
		var list *baremetal.ListVnicAttachments
		if list, e = r.Client.ListVnicAttachments(compartmentID, opts); e != nil {
			break
		}

		r.Res.Attachments = append(r.Res.Attachments, list.Attachments...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (r *VnicAttachmentDatasourceCrud) SetData() {

	if r.Res != nil {
		r.D.SetId(time.Now().UTC().String())
		attachments := []map[string]string{}

		for _, att := range r.Res.Attachments {
			attachment := map[string]string{}
			attachment["id"] = att.ID
			attachment["display_name"] = att.DisplayName
			attachment["availability_domain"] = att.AvailabilityDomain
			attachment["compartment_id"] = att.CompartmentID
			attachment["instance_id"] = att.InstanceID
			attachment["state"] = att.State
			attachment["subnet_id"] = att.SubnetID
			attachment["time_created"] = att.TimeCreated.Format(time.RFC1123)
			attachment["vnic_id"] = att.VnicID
			attachments = append(attachments, attachment)
		}

		r.D.Set("vnic_attachments", attachments)

	}

}
