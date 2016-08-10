package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type VnicAttachmentsReader struct {
	resourceData *schema.ResourceData
	client       BareMetalClient
	response     *baremtlsdk.VnicAttachmentList
}

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

func ResourceCoreVnicAttachments() *schema.Resource {
	return &schema.Resource{
		Read: readVnicAttachments,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_attachments": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceVnicAttachment(),
			},
		},
	}
}

func readVnicAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	reader := &VnicAttachmentsReader{
		resourceData: d,
		client:       client,
	}

	return readResource(reader)
}

func (r *VnicAttachmentsReader) Get() (e error) {
	compartmentID := r.resourceData.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(
		r.resourceData,
		"availability_domain",
		"instance_id",
		"vnic_id",
	)

	r.response, e = r.client.ListVnicAttachments(compartmentID, opts...)
	return
}

func (r *VnicAttachmentsReader) SetData() {

	if r.response != nil {
		r.resourceData.SetId(time.Now().UTC().String())
		attachments := []map[string]string{}

		for _, att := range r.response.Attachments {
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

		r.resourceData.Set("vnic_attachments", attachments)

	}

}
