// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// FunctionsFunctionsRuntimeVersionDataSource defines the singular Functions runtime version lookup schema.
func FunctionsFunctionsRuntimeVersionDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularFunctionsFunctionsRuntimeVersionWithContext,
		Schema: map[string]*schema.Schema{
			"functions_runtime_version_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
			"functions_runtime_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_architectures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// readSingularFunctionsFunctionsRuntimeVersionWithContext wires Terraform reads to the runtime version CRUD.
func readSingularFunctionsFunctionsRuntimeVersionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FunctionsFunctionsRuntimeVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FunctionsFunctionsRuntimeVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.GetFunctionsRuntimeVersionResponse
}

func (s *FunctionsFunctionsRuntimeVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

// GetWithContext calls OCI to fetch one Functions runtime version by OCID.
func (s *FunctionsFunctionsRuntimeVersionDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_functions.GetFunctionsRuntimeVersionRequest{}

	if functionsRuntimeVersionId, ok := s.D.GetOkExists("functions_runtime_version_id"); ok {
		tmp := functionsRuntimeVersionId.(string)
		request.FunctionsRuntimeVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.GetFunctionsRuntimeVersion(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

// SetData maps the OCI runtime version response into Terraform state.
func (s *FunctionsFunctionsRuntimeVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.FunctionsRuntimeId != nil {
		s.D.Set("functions_runtime_id", *s.Res.FunctionsRuntimeId)
	}

	if s.Res.LanguageVersion != nil {
		s.D.Set("language_version", *s.Res.LanguageVersion)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("supported_architectures", s.Res.SupportedArchitectures)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
