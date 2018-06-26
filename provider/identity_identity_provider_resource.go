// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityProviderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createIdentityProvider,
		Read:     readIdentityProvider,
		Update:   updateIdentityProvider,
		Delete:   deleteIdentityProvider,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
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

func createIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProviderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

// 07-05-2018: Identity Providers support only SAML2 as the protocol
type IdentityProviderResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Saml2IdentityProvider
	DisableNotFoundRetries bool
}

func (s *IdentityProviderResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityProviderResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateCreating),
	}
}

func (s *IdentityProviderResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateActive),
	}
}

func (s *IdentityProviderResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateDeleting),
	}
}

func (s *IdentityProviderResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.IdentityProviderLifecycleStateDeleted),
	}
}

func (s *IdentityProviderResourceCrud) Create() error {
	request := oci_identity.CreateIdentityProviderRequest{}
	err := s.populateTopLevelPolymorphicCreateIdentityProviderRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	if provider, ok := response.IdentityProvider.(oci_identity.Saml2IdentityProvider); ok {
		s.Res = &provider
	}

	return nil
}

func (s *IdentityProviderResourceCrud) Get() error {
	request := oci_identity.GetIdentityProviderRequest{}

	tmp := s.D.Id()
	request.IdentityProviderId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	if provider, ok := response.IdentityProvider.(oci_identity.Saml2IdentityProvider); ok {
		s.Res = &provider
	}

	return nil
}

func (s *IdentityProviderResourceCrud) Update() error {
	request := oci_identity.UpdateIdentityProviderRequest{}
	err := s.populateTopLevelPolymorphicUpdateIdentityProviderRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	if provider, ok := response.IdentityProvider.(oci_identity.Saml2IdentityProvider); ok {
		s.Res = &provider
	}

	return nil
}

func (s *IdentityProviderResourceCrud) Delete() error {
	request := oci_identity.DeleteIdentityProviderRequest{}

	tmp := s.D.Id()
	request.IdentityProviderId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteIdentityProvider(context.Background(), request)
	return err
}

func (s *IdentityProviderResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("product_type", s.Res.ProductType)

	s.D.Set("protocol", string(oci_identity.ListIdentityProvidersProtocolSaml2))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}

func (s *IdentityProviderResourceCrud) populateTopLevelPolymorphicCreateIdentityProviderRequest(request *oci_identity.CreateIdentityProviderRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}

	switch protocol {
	case "SAML2":
		details := oci_identity.CreateSaml2IdentityProviderDetails{}
		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			details.Metadata = &tmp
		}
		if metadataUrl, ok := s.D.GetOkExists("metadata_url"); ok {
			tmp := metadataUrl.(string)
			details.MetadataUrl = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if productType, ok := s.D.GetOkExists("product_type"); ok {
			details.ProductType = oci_identity.CreateIdentityProviderDetailsProductTypeEnum(productType.(string))
		}
		request.CreateIdentityProviderDetails = details
	default:
		return fmt.Errorf("Unknown protocol '%v' was specified", protocol)
	}
	return nil
}

func (s *IdentityProviderResourceCrud) populateTopLevelPolymorphicUpdateIdentityProviderRequest(request *oci_identity.UpdateIdentityProviderRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}

	switch protocol {
	case "SAML2":
		details := oci_identity.UpdateSaml2IdentityProviderDetails{}
		if metadata, ok := s.D.GetOkExists("metadata"); ok {
			tmp := metadata.(string)
			details.Metadata = &tmp
		}
		if metadataUrl, ok := s.D.GetOkExists("metadata_url"); ok {
			tmp := metadataUrl.(string)
			details.MetadataUrl = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
		}

		tmp := s.D.Id()
		request.IdentityProviderId = &tmp

		request.UpdateIdentityProviderDetails = details
	default:
		return fmt.Errorf("Unknown protocol '%v' was specified", protocol)
	}
	return nil
}
