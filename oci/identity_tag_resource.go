// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityTagResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createIdentityTag,
		Read:     readIdentityTag,
		Update:   updateIdentityTag,
		Delete:   deleteIdentityTag,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
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
			"is_cost_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteIdentityTag(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityTagResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.Tag
	DisableNotFoundRetries bool
}

func (s *IdentityTagResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityTagResourceCrud) Create() error {
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

	if isCostTracking, ok := s.D.GetOkExists("is_cost_tracking"); ok {
		tmp := isCostTracking.(bool)
		request.IsCostTracking = &tmp
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
		s.D.SetId(*s.Res.Id)
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
	importIfExists, _ := strconv.ParseBool(getEnvSettingWithDefault("tags_import_if_exists", "false"))
	if importIfExists && strings.Contains(err.Error(), "TagDefinitionAlreadyExists") {
		// List all tag definitions using the datasource to find that tag definition which matches
		s.D.Set("tag_namespace_id", request.TagNamespaceId)
		s.D.Set("name", request.Name)
		dsCrud := &IdentityTagsDataSourceCrud{s.D, s.Client, nil}
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

func (s *IdentityTagResourceCrud) Get() error {
	request := oci_identity.GetTagRequest{}

	tagName, tagNamespaceId, parseTagCompositeIdErr := parseTagCompositeId(s.D.Id())
	if parseTagCompositeIdErr == nil {
		request.TagName = &tagName
		request.TagNamespaceId = &tagNamespaceId
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

	response, err := s.Client.GetTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Tag
	if parseTagCompositeIdErr == nil {
		// Import sets the ID to composite ID and hence overwriting ID to OCID from response
		id := response.Tag.Id
		if id == nil {
			return fmt.Errorf("error : received null value for id attribute for request %s, id attribute cannot be null", *response.OpcRequestId)
		}
		s.D.SetId(*id)
	}
	return nil
}

func (s *IdentityTagResourceCrud) Update() error {
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

	if isCostTracking, ok := s.D.GetOkExists("is_cost_tracking"); ok {
		tmp := isCostTracking.(bool)
		request.IsCostTracking = &tmp
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

func (s *IdentityTagResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCostTracking != nil {
		s.D.Set("is_cost_tracking", *s.Res.IsCostTracking)
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

	return nil
}

func parseTagCompositeId(compositeId string) (tagName string, tagNamespaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("tagNamespaces/.*/tags/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	tagNamespaceId = parts[1]
	tagName = parts[3]

	return
}
