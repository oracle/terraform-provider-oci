// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
)

func FunctionsFunctionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["function_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FunctionsFunctionResource(), fieldMap, readSingularFunctionsFunction)
}

func readSingularFunctionsFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsFunctionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsFunctionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.GetFunctionResponse
}

func (s *FunctionsFunctionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsFunctionDataSourceCrud) Get() error {
	request := oci_functions.GetFunctionRequest{}

	if functionId, ok := s.D.GetOkExists("function_id"); ok {
		tmp := functionId.(string)
		request.FunctionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.GetFunction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FunctionsFunctionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplicationId != nil {
		s.D.Set("application_id", *s.Res.ApplicationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config", s.Res.Config)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Image != nil {
		s.D.Set("image", *s.Res.Image)
	}

	if s.Res.ImageDigest != nil {
		s.D.Set("image_digest", *s.Res.ImageDigest)
	}

	if s.Res.InvokeEndpoint != nil {
		s.D.Set("invoke_endpoint", *s.Res.InvokeEndpoint)
	}

	if s.Res.MemoryInMBs != nil {
		s.D.Set("memory_in_mbs", strconv.FormatInt(*s.Res.MemoryInMBs, 10))
	}

	if s.Res.ProvisionedConcurrencyConfig != nil {
		provisionedConcurrencyConfigArray := []interface{}{}
		if provisionedConcurrencyConfigMap := FunctionProvisionedConcurrencyConfigToMap(&s.Res.ProvisionedConcurrencyConfig); provisionedConcurrencyConfigMap != nil {
			provisionedConcurrencyConfigArray = append(provisionedConcurrencyConfigArray, provisionedConcurrencyConfigMap)
		}
		s.D.Set("provisioned_concurrency_config", provisionedConcurrencyConfigArray)
	} else {
		s.D.Set("provisioned_concurrency_config", nil)
	}

	s.D.Set("shape", s.Res.Shape)

	if s.Res.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := FunctionSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	if s.Res.TraceConfig != nil {
		s.D.Set("trace_config", []interface{}{FunctionTraceConfigToMap(s.Res.TraceConfig)})
	} else {
		s.D.Set("trace_config", nil)
	}

	return nil
}
