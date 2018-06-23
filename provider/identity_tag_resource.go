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

func TagResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createTag,
		Read:     readTag,
		Update:   updateTag,
		Delete:   deleteTag,
		Schema: map[string]*schema.Schema{
			// Required
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
			"tag_namespace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func createTag(d *schema.ResourceData, m interface{}) error {
	sync := &TagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readTag(d *schema.ResourceData, m interface{}) error {
	sync := &TagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateTag(d *schema.ResourceData, m interface{}) error {
	sync := &TagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteTag(d *schema.ResourceData, m interface{}) error {
	return nil
}

type TagResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Tag
	DisableNotFoundRetries bool
}

func (s *TagResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *TagResourceCrud) Create() error {
	request := oci_identity.CreateTagRequest{}

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

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	contextToUse := context.Background()
	response, err := s.Client.CreateTag(contextToUse, request)
	if err == nil {
		s.Res = &response.Tag
		//is_retired field is currently not supported in create so update to make server state same as config
		if updateError := s.Update(); updateError != nil {
			return updateError
		}
		return nil
	}

	// Tag definitions can't be deleted, so this is a work around here to react to collisions by
	// basically importing that pre-existing namespace into this plan if tags_import_if_exists
	// flag is set to 'true'. This is ONLY for TESTING and should not be used elsewhere.
	// Use 'terraform import' for existing tag definitions
	importIfExists, _ := strconv.ParseBool(getEnvSetting("tags_import_if_exists", "false"))
	if importIfExists && strings.Contains(err.Error(), "TagDefinitionAlreadyExists") {
		// List all tag definitions using the datasource to find that tag definition which matches
		s.D.Set("tag_namespace_id", request.TagNamespaceId)
		s.D.Set("name", request.Name)
		dsCrud := &TagsDataSourceCrud{s.D, s.Client, nil}
		if dsErr := dsCrud.Get(); dsErr != nil {
			//return original error when datasource call fails
			return err
		}

		for _, tag := range dsCrud.Res.Items {
			if strings.EqualFold(*tag.Name, *request.Name) {
				s.D.SetId(*tag.Id)
				if updateError := s.Update(); updateError != nil {
					//Update to tags can only be done from home region, so do get in that case
					if getError := s.Get(); getError != nil {
						return getError
					}
				}
				return nil
			}
		}
	}

	return err

}

func (s *TagResourceCrud) Get() error {
	request := oci_identity.GetTagRequest{}

	if tagName, ok := s.D.GetOkExists("name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tag
	return nil
}

func (s *TagResourceCrud) Update() error {
	request := oci_identity.UpdateTagRequest{}

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

	if tagName, ok := s.D.GetOkExists("name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tag
	return nil
}

func (s *TagResourceCrud) SetData() {
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

	if s.Res.TagNamespaceId != nil {
		s.D.Set("tag_namespace_id", *s.Res.TagNamespaceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
