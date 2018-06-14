// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func DrgAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDrgAttachment,
		Read:     readDrgAttachment,
		Update:   updateDrgAttachment,
		Delete:   deleteDrgAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				// Used to be required. Added back as optional to avoid showing errors (this field
				// cannot be set).
				Optional:   true,
				Deprecated: "No longer required. The DRG attachment is automatically placed into the same compartment as the VCN.",
			},
			"id": {
				Type:     schema.TypeString,
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
	}
}

func createDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &DrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DrgAttachmentResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgAttachment
	DisableNotFoundRetries bool
}

func (s *DrgAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DrgAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttaching),
	}
}

func (s *DrgAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttached),
	}
}

func (s *DrgAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetaching),
	}
}

func (s *DrgAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetached),
	}
}

func (s *DrgAttachmentResourceCrud) Create() error {
	request := oci_core.CreateDrgAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *DrgAttachmentResourceCrud) Get() error {
	request := oci_core.GetDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *DrgAttachmentResourceCrud) Update() error {
	request := oci_core.UpdateDrgAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *DrgAttachmentResourceCrud) Delete() error {
	request := oci_core.DeleteDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrgAttachment(context.Background(), request)
	return err
}

func (s *DrgAttachmentResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}
