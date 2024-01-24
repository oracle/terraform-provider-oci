// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"strings"

	"strconv"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityTagNamespaceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityTagNamespace,
		Read:     readIdentityTagNamespace,
		Update:   updateIdentityTagNamespace,
		Delete:   deleteIdentityTagNamespace,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
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

func createIdentityTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityTagNamespace(d *schema.ResourceData, m interface{}) error {
	// Only empty tag namespaces can be deleted, to execute our tests we don't want to delete namespaces as we Create
	// namespaces with tags and deleting a tag is a sequential and time consuming operation allowed one per tenancy
	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists {
		return nil
	}

	sync := &IdentityTagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityTagNamespaceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.TagNamespace
	DisableNotFoundRetries bool
}

func (s *IdentityTagNamespaceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityTagNamespaceResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IdentityTagNamespaceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.TagNamespaceLifecycleStateActive),
	}
}

func (s *IdentityTagNamespaceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.TagNamespaceLifecycleStateDeleting),
	}
}

func (s *IdentityTagNamespaceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.TagNamespaceLifecycleStateDeleted),
	}
}

func (s *IdentityTagNamespaceResourceCrud) Create() error {
	request := oci_identity.CreateTagNamespaceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	contextToUse := context.Background()
	response, err := s.Client.CreateTagNamespace(contextToUse, request)
	if err == nil {
		s.Res = &response.TagNamespace
		s.D.SetId(*s.Res.Id)
		//is_retired field is currently not supported in Create so Update to make server state same as config
		if updateError := s.Update(); updateError != nil {
			return updateError
		}
		return nil
	}
	// Tag Namespaces can't be deleted, so there is a work around here to react to name collisions
	// by basically importing that pre-existing namespace into this plan if tags_import_if_exists
	// flag is set to 'true'. This is ONLY for TESTING and should not be used elsewhere.
	// Use 'terraform import' for existing namespaces
	importIfExists, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists && strings.Contains(err.Error(), "TagNamespaceAlreadyExists") {
		// List all namespaces using the datasource to find that namespace with the matching name.
		s.D.Set("compartment_id", request.CompartmentId)
		s.D.Set("name", request.Name)
		dsCrud := &IdentityTagNamespacesDataSourceCrud{s.D, s.Client, nil}
		if dsErr := dsCrud.Get(); dsErr != nil {
			//return original error when datasource call fails
			return err
		}

		for _, namespace := range dsCrud.Res.Items {
			if strings.EqualFold(*namespace.Name, *request.Name) {
				s.D.SetId(*namespace.Id)
				if updateError := s.Update(); updateError != nil {
					if getError := s.Get(); getError != nil {
						//Update to tags can only be done from home region, so do get in that case
						return getError
					}
				}
				return nil
			}
		}
	}

	return err
}

func (s *IdentityTagNamespaceResourceCrud) Get() error {
	request := oci_identity.GetTagNamespaceRequest{}

	tmp := s.D.Id()
	request.TagNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetTagNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagNamespace
	return nil
}

func (s *IdentityTagNamespaceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_identity.UpdateTagNamespaceRequest{}

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

	if isRetired, ok := s.D.GetOkExists("is_retired"); ok {
		tmp := isRetired.(bool)
		request.IsRetired = &tmp
	}

	tmp := s.D.Id()
	request.TagNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateTagNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagNamespace
	return nil
}

func (s *IdentityTagNamespaceResourceCrud) Delete() error {
	request := oci_identity.DeleteTagNamespaceRequest{}

	tmp := s.D.Id()
	request.TagNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteTagNamespace(context.Background(), request)
	return err
}

func (s *IdentityTagNamespaceResourceCrud) SetData() error {
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

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
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

func (s *IdentityTagNamespaceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_identity.ChangeTagNamespaceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.TagNamespaceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.ChangeTagNamespaceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
