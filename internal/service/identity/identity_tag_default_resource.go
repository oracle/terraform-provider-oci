// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentityTagDefaultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityTagDefault,
		Read:     readIdentityTagDefault,
		Update:   updateIdentityTagDefault,
		Delete:   deleteIdentityTagDefault,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tag_definition_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"is_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_definition_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_namespace_id": {
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

func createIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityTagDefaultResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.TagDefault
	DisableNotFoundRetries bool
}

func (s *IdentityTagDefaultResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityTagDefaultResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IdentityTagDefaultResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.TagDefaultLifecycleStateActive),
	}
}

func (s *IdentityTagDefaultResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *IdentityTagDefaultResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *IdentityTagDefaultResourceCrud) Create() error {
	request := oci_identity.CreateTagDefaultRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isRequired, ok := s.D.GetOkExists("is_required"); ok {
		tmp := isRequired.(bool)
		request.IsRequired = &tmp
	}

	if tagDefinitionId, ok := s.D.GetOkExists("tag_definition_id"); ok {
		tmp := tagDefinitionId.(string)
		request.TagDefinitionId = &tmp
	}

	if value, ok := s.D.GetOkExists("value"); ok {
		tmp := value.(string)
		request.Value = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateTagDefault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagDefault

	// service takes some time for the isRequired effect to get stabilized
	if isRequired, ok := s.D.GetOkExists("is_required"); ok {
		tmp := isRequired.(bool)

		if tmp {
			time.Sleep(20 * time.Second)
		}
	}

	return nil
}

func (s *IdentityTagDefaultResourceCrud) Get() error {
	request := oci_identity.GetTagDefaultRequest{}

	tmp := s.D.Id()
	request.TagDefaultId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetTagDefault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagDefault
	return nil
}

func (s *IdentityTagDefaultResourceCrud) Update() error {
	request := oci_identity.UpdateTagDefaultRequest{}

	if isRequired, ok := s.D.GetOkExists("is_required"); ok {
		tmp := isRequired.(bool)
		request.IsRequired = &tmp
	}

	tmp := s.D.Id()
	request.TagDefaultId = &tmp

	if value, ok := s.D.GetOkExists("value"); ok {
		tmp := value.(string)
		request.Value = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateTagDefault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagDefault
	// service takes some time for the isRequired effect to get stabilized
	if isRequired, ok := s.D.GetOkExists("is_required"); ok && s.D.HasChange("is_required") {
		tmp := isRequired.(bool)

		if tmp {
			time.Sleep(20 * time.Second)
		}
	}
	return nil
}

func (s *IdentityTagDefaultResourceCrud) Delete() error {
	request := oci_identity.DeleteTagDefaultRequest{}

	tmp := s.D.Id()
	request.TagDefaultId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteTagDefault(context.Background(), request)
	return err
}

func (s *IdentityTagDefaultResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsRequired != nil {
		s.D.Set("is_required", *s.Res.IsRequired)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TagDefinitionId != nil {
		s.D.Set("tag_definition_id", *s.Res.TagDefinitionId)
	}

	if s.Res.TagDefinitionName != nil {
		s.D.Set("tag_definition_name", *s.Res.TagDefinitionName)
	}

	if s.Res.TagNamespaceId != nil {
		s.D.Set("tag_namespace_id", *s.Res.TagNamespaceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	return nil
}
