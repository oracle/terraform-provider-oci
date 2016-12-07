package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VnicAttachmentDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ListVnicAttachments
}

func (r *VnicAttachmentDatasourceCrud) Get() (e error) {
	compartmentID := r.D.Get("compartment_id").(string)

	opts := &baremetal.ListVnicAttachmentsOptions{}
	setListOptions(r.D, &opts.ListOptions)
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

		if hasNextPage := setNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
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
