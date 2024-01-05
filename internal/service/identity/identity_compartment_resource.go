// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityCompartmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Delete: tfresource.GetTimeoutDuration("90m"), // service team states: p50: 30 min, p90: 60 min, max: 180 min
		},
		Create: createIdentityCompartment,
		Read:   readIdentityCompartment,
		Update: updateIdentityCompartment,
		Delete: deleteIdentityCompartment,
		Schema: map[string]*schema.Schema{
			// Required
			// @next-break: remove customizations
			// The legacy provider exposed this as read-only/computed. The API requires this param. For legacy users who are
			// not supplying a value, make it optional, behind the scenes it will use the tenancy ocid if not supplied.
			// If a user supplies the value, then changes it, it requires forcing new.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"enable_delete": {
				Type:     schema.TypeBool,
				Optional: true,
			},
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"inactive_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_accessible": {
				Type:     schema.TypeBool,
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

func createIdentityCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.Configuration = m.(*client.OracleClients).Configuration

	return tfresource.CreateResource(d, sync)
}

func readIdentityCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityCompartment(d *schema.ResourceData, m interface{}) error {
	if enableDelete, ok := d.GetOkExists("enable_delete"); !ok || !enableDelete.(bool) {
		return nil
	}

	sync := &IdentityCompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityCompartmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Configuration          map[string]string
	Res                    *oci_identity.Compartment
	DisableNotFoundRetries bool
}

func (s *IdentityCompartmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityCompartmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateCreating),
	}
}

func (s *IdentityCompartmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateActive),
	}
}

func (s *IdentityCompartmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateDeleting),
	}
}

func (s *IdentityCompartmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.CompartmentLifecycleStateDeleted),
	}
}

func (s *IdentityCompartmentResourceCrud) Create() error {
	request := oci_identity.CreateCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else { // @next-break: remove
		// Prevent potentially inferring wrong TenancyOCID from InstancePrincipal
		if auth := s.Configuration["auth"]; strings.ToLower(auth) == strings.ToLower(globalvar.AuthInstancePrincipalSetting) {
			return fmt.Errorf("compartment_id must be specified for this resource when using with auth as '%s'", globalvar.AuthInstancePrincipalSetting)
		}
		// Maintain legacy contract of compartment_id defaulting to tenancy_ocid if not specified
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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateCompartment(context.Background(), request)
	if err != nil {
		if response.RawResponse != nil && response.RawResponse.StatusCode == 409 {

			// Return an error if 'enable_delete' was explicitly set to 'true' in case of automatic import on conflict
			if enableDelete, ok := s.D.GetOkExists("enable_delete"); ok && enableDelete.(bool) {
				return fmt.Errorf(`%s

If you define a compartment resource in your configurations with 
the same name as an existing compartment with 'enable_delete' set to 'true', 
the compartment will no longer be automatically imported. 
If you intended to manage an existing compartment, use terraform import instead.`, err)
			}

			// React to name collisions or conflict errors by importing pre-existing compartment into this plan if the name matches.
			if strings.Contains(err.Error(), "already exists") ||
				strings.Contains(err.Error(), "Maximum number of compartment") {
				// List all compartments using the datasource to find that compartment with the matching name.
				// CompartmentsDataSourceCrud requires a compartment_id, so forward whatever value was used in
				// the Create attempt above.
				s.D.Set("compartment_id", request.CompartmentId)
				log.Println(fmt.Sprintf("[DEBUG] The specified compartment with name '%s' may already exist, listing compartments to lookup with name instead.",
					*request.Name))
				dsCrud := &IdentityCompartmentsDataSourceCrud{s.D, s.Client, nil}
				if err := dsCrud.Get(); err != nil {
					return err
				}

				for _, compartment := range dsCrud.Res.Items {
					if *compartment.Name == *request.Name {
						s.Res = &compartment
						//Update with correct description
						s.D.SetId(s.ID())
						return s.Update()
					}
				}
				// Return an error if the lookup failed, to provide user with information on which compartment id and name were used for lookup
				return fmt.Errorf(`%s

failed to lookup the compartment with name: '%s' in compartment_id: '%s'.
Verify your configuration if the correct 'compartment_id' and 'name' were specified.
In most cases, the 'compartment_id' will be your 'tenancy_ocid' with the exception of nested compartments.
Refer to the 'oci_identity_compartment' documentation for more information.`, err, *request.Name, *request.CompartmentId)
			}
		}
		return err
	}
	s.Res = &response.Compartment
	return nil
}

func (s *IdentityCompartmentResourceCrud) Get() error {
	request := oci_identity.GetCompartmentRequest{}

	tmp := s.D.Id()
	request.CompartmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Compartment
	return nil
}

func (s *IdentityCompartmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_identity.UpdateCompartmentRequest{}

	tmp := s.D.Id()
	request.CompartmentId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Compartment
	return nil
}

func (s *IdentityCompartmentResourceCrud) Delete() error {
	request := oci_identity.DeleteCompartmentRequest{}

	tmp := s.D.Id()
	request.CompartmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteCompartment(context.Background(), request)
	return err
}

func (s *IdentityCompartmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.IsAccessible != nil {
		s.D.Set("is_accessible", *s.Res.IsAccessible)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *IdentityCompartmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_identity.MoveCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CompartmentId = &idTmp

	if targetCompartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := targetCompartmentId.(string)
		changeCompartmentRequest.TargetCompartmentId = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.MoveCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
