// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func CompartmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createCompartment,
		Read:     readCompartment,
		Update:   updateCompartment,
		Delete:   deleteCompartment,
		Schema: map[string]*schema.Schema{
			// The legacy provider exposed this as read-only/computed. The API requires this param. For legacy users who are
			// not supplying a value, make it optional, behind the scenes it will use the tenancy ocid if not supplied.
			// If a user supplies the value, then changes it, it requires forcing new.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			// @Deprecated 01/2018: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Deprecated: crud.FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) error {
	return nil
}

type CompartmentResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Compartment
	DisableNotFoundRetries bool
}

func (s *CompartmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CompartmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateCreating),
	}
}

func (s *CompartmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateActive),
	}
}

func (s *CompartmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateDeleting),
	}
}

func (s *CompartmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateDeleted),
	}
}

func (s *CompartmentResourceCrud) Create() error {
	request := oci_identity.CreateCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else {
		c := *s.Client.ConfigurationProvider()
		if c == nil {
			return fmt.Errorf("cannot access tenancyOCID")
		}
		tenancy, err := c.TenancyOCID()
		if err != nil {
			return err
		}
		request.CompartmentId = &tenancy
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateCompartment(context.Background(), request)
	if err == nil {
		s.Res = &response.Compartment
		return nil
	}

	// Compartments can't be destroyed, so there is a work around here to react to name collisions
	// by basically importing that pre-existing compartment into this plan.
	if strings.Contains(err.Error(), "already exists") {
		// List all compartments using the datasource to find that compartment with the matching name.
		// CompartmentsDataSourceCrud requires a compartment_id, so forward whatever value was used in
		// the create attempt above.
		s.D.Set("compartment_id", request.CompartmentId)
		dsCrud := &CompartmentsDataSourceCrud{s.D, s.Client, nil}
		if err = dsCrud.Get(); err != nil {
			return err
		}

		for _, compartment := range dsCrud.Res.Items {
			if *compartment.Name == *request.Name {
				s.Res = &compartment
				//Update with correct description
				s.D.SetId(s.ID())
				s.Update()
				return nil
			}
		}
	}

	return err
}

func (s *CompartmentResourceCrud) Get() error {
	request := oci_identity.GetCompartmentRequest{}

	tmp := s.D.Id()
	request.CompartmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Compartment
	return nil
}

func (s *CompartmentResourceCrud) Update() error {
	request := oci_identity.UpdateCompartmentRequest{}

	tmp := s.D.Id()
	request.CompartmentId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Compartment
	return nil
}

func (s *CompartmentResourceCrud) Delete() error {
	// Compartments cannot be deleted. Just pretend it worked.
	return nil
}

func (s *CompartmentResourceCrud) SetData() {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
