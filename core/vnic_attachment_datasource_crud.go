package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VnicAttachmentsSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.VnicAttachmentList
}

func (r *VnicAttachmentsSync) Get() (e error) {
	compartmentID := r.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(
		r.D,
		"availability_doresource",
		"instance_id",
		"vnic_id",
	)

	r.Res, e = r.Client.ListVnicAttachments(compartmentID, opts...)
	return
}

func (r *VnicAttachmentsSync) SetData() {

	if r.Res != nil {
		r.D.SetId(time.Now().UTC().String())
		attachments := []map[string]string{}

		for _, att := range r.Res.Attachments {
			attachment := map[string]string{}
			attachment["id"] = att.ID
			attachment["display_name"] = att.DisplayName
			attachment["availability_doresource"] = att.AvailabilityDomain
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
