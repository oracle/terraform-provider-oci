// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchClusterPipelineDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["opensearch_cluster_pipeline_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpensearchOpensearchClusterPipelineResource(), fieldMap, readSingularOpensearchOpensearchClusterPipeline)
}

func readSingularOpensearchOpensearchClusterPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterPipelineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterPipelineClient()

	return tfresource.ReadResource(sync)
}

type OpensearchOpensearchClusterPipelineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opensearch.OpensearchClusterPipelineClient
	Res    *oci_opensearch.GetOpensearchClusterPipelineResponse
}

func (s *OpensearchOpensearchClusterPipelineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpensearchOpensearchClusterPipelineDataSourceCrud) Get() error {
	request := oci_opensearch.GetOpensearchClusterPipelineRequest{}

	if opensearchClusterPipelineId, ok := s.D.GetOkExists("opensearch_cluster_pipeline_id"); ok {
		tmp := opensearchClusterPipelineId.(string)
		request.OpensearchClusterPipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opensearch")

	response, err := s.Client.GetOpensearchClusterPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpensearchOpensearchClusterPipelineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataPrepperConfigurationBody != nil {
		s.D.Set("data_prepper_configuration_body", *s.Res.DataPrepperConfigurationBody)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MemoryGB != nil {
		s.D.Set("memory_gb", *s.Res.MemoryGB)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	if s.Res.NodeShape != nil {
		s.D.Set("node_shape", *s.Res.NodeShape)
	}

	if s.Res.NsgId != nil {
		s.D.Set("nsg_id", *s.Res.NsgId)
	}

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.OpensearchPipelineFqdn != nil {
		s.D.Set("opensearch_pipeline_fqdn", *s.Res.OpensearchPipelineFqdn)
	}

	if s.Res.OpensearchPipelinePrivateIp != nil {
		s.D.Set("opensearch_pipeline_private_ip", *s.Res.OpensearchPipelinePrivateIp)
	}

	if s.Res.PipelineConfigurationBody != nil {
		s.D.Set("pipeline_configuration_body", *s.Res.PipelineConfigurationBody)
	}

	s.D.Set("pipeline_mode", s.Res.PipelineMode)

	reverseConnectionEndpoints := []interface{}{}
	for _, item := range s.Res.ReverseConnectionEndpoints {
		reverseConnectionEndpoints = append(reverseConnectionEndpoints, OpensearchPipelineReverseConnectionEndpointToMap(item))
	}
	s.D.Set("reverse_connection_endpoints", reverseConnectionEndpoints)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetCompartmentId != nil {
		s.D.Set("subnet_compartment_id", *s.Res.SubnetCompartmentId)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnCompartmentId != nil {
		s.D.Set("vcn_compartment_id", *s.Res.VcnCompartmentId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
