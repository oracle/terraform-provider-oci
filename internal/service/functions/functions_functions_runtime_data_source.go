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

// FunctionsFunctionsRuntimeDataSource defines the singular Functions runtime lookup schema.
func FunctionsFunctionsRuntimeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularFunctionsFunctionsRuntimeWithContext,
		Schema: map[string]*schema.Schema{
			"functions_runtime_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"current_functions_runtime_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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
			"time_decommissioned": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_deprecated": {
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

// readSingularFunctionsFunctionsRuntimeWithContext wires Terraform reads to the runtime data source CRUD.
func readSingularFunctionsFunctionsRuntimeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FunctionsFunctionsRuntimeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type FunctionsFunctionsRuntimeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.GetFunctionsRuntimeResponse
}

func (s *FunctionsFunctionsRuntimeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

// GetWithContext calls OCI to fetch one Functions runtime by OCID.
func (s *FunctionsFunctionsRuntimeDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_functions.GetFunctionsRuntimeRequest{}

	if functionsRuntimeId, ok := s.D.GetOkExists("functions_runtime_id"); ok {
		tmp := functionsRuntimeId.(string)
		request.FunctionsRuntimeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.GetFunctionsRuntime(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

// SetData maps the OCI runtime response into Terraform state.
func (s *FunctionsFunctionsRuntimeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CurrentFunctionsRuntimeVersionId != nil {
		s.D.Set("current_functions_runtime_version_id", *s.Res.CurrentFunctionsRuntimeVersionId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Language != nil {
		s.D.Set("language", *s.Res.Language)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Os != nil {
		s.D.Set("os", *s.Res.Os)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDecommissioned != nil {
		s.D.Set("time_decommissioned", s.Res.TimeDecommissioned.String())
	}

	if s.Res.TimeDeprecated != nil {
		s.D.Set("time_deprecated", s.Res.TimeDeprecated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
