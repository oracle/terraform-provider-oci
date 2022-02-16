// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v58/resourcemanager"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerStackDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularResourcemanagerStack,
		Schema: map[string]*schema.Schema{
			"stack_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_source": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"config_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ZIP_UPLOAD",
							}, true),
						},
						"zip_file_base64encoded": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"working_directory": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"variables": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func readSingularResourcemanagerStack(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

type ResourcemanagerStackDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resourcemanager.ResourceManagerClient
	Res    *oci_resourcemanager.GetStackResponse
}

func (s *ResourcemanagerStackDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourcemanagerStackDataSourceCrud) Get() error {
	request := oci_resourcemanager.GetStackRequest{}

	if stackId, ok := s.D.GetOkExists("stack_id"); ok {
		tmp := stackId.(string)
		request.StackId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.GetStack(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourcemanagerStackDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigSource != nil {
		configSourceArray := []interface{}{}
		if configSourceMap := ConfigSourceToMap(&s.Res.ConfigSource); configSourceMap != nil {
			configSourceArray = append(configSourceArray, configSourceMap)
		}
		s.D.Set("config_source", configSourceArray)
	} else {
		s.D.Set("config_source", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("variables", s.Res.Variables)

	return nil
}

func ConfigSourceToMap(obj *oci_resourcemanager.ConfigSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch (*obj).(type) {
	case oci_resourcemanager.ZipUploadConfigSource:
		result["config_source_type"] = "ZIP_UPLOAD"
	default:
		log.Printf("[WARN] Received 'config_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
