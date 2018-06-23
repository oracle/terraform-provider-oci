// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"strings"

	"strconv"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func TagNamespaceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createTagNamespace,
		Read:     readTagNamespace,
		Update:   updateTagNamespace,
		Delete:   deleteTagNamespace,
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
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
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
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"id": {
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

func createTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &TagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &TagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateTagNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &TagNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteTagNamespace(d *schema.ResourceData, m interface{}) error {
	return nil
}

type TagNamespaceResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.TagNamespace
	DisableNotFoundRetries bool
}

func (s *TagNamespaceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *TagNamespaceResourceCrud) Create() error {
	request := oci_identity.CreateTagNamespaceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	contextToUse := context.Background()
	response, err := s.Client.CreateTagNamespace(contextToUse, request)
	if err == nil {
		s.Res = &response.TagNamespace
		//is_retired field is currently not supported in create so update to make server state same as config
		if updateError := s.Update(); updateError != nil {
			return updateError
		}
		return nil
	}
	// Tag Namespaces can't be deleted, so there is a work around here to react to name collisions
	// by basically importing that pre-existing namespace into this plan if tags_import_if_exists
	// flag is set to 'true'. This is ONLY for TESTING and should not be used elsewhere.
	// Use 'terraform import' for existing namespaces
	importIfExists, _ := strconv.ParseBool(getEnvSetting("tags_import_if_exists", "false"))
	if importIfExists && strings.Contains(err.Error(), "TagNamespaceAlreadyExists") {
		// List all namespaces using the datasource to find that namespace with the matching name.
		s.D.Set("compartment_id", request.CompartmentId)
		s.D.Set("name", request.Name)
		s.D.Set("include_subcompartments", false)
		dsCrud := &TagNamespacesDataSourceCrud{s.D, s.Client, nil}
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

func (s *TagNamespaceResourceCrud) Get() error {
	request := oci_identity.GetTagNamespaceRequest{}

	tmp := s.D.Id()
	request.TagNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetTagNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagNamespace
	return nil
}

func (s *TagNamespaceResourceCrud) Update() error {
	request := oci_identity.UpdateTagNamespaceRequest{}

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

	if isRetired, ok := s.D.GetOkExists("is_retired"); ok {
		tmp := isRetired.(bool)
		request.IsRetired = &tmp
	}

	tmp := s.D.Id()
	request.TagNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateTagNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TagNamespace
	return nil
}

func (s *TagNamespaceResourceCrud) SetData() {
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

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
