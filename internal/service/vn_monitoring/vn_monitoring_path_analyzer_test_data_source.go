// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vn_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VnMonitoringPathAnalyzerTestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["path_analyzer_test_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(VnMonitoringPathAnalyzerTestResource(), fieldMap, readSingularVnMonitoringPathAnalyzerTest)
}

func readSingularVnMonitoringPathAnalyzerTest(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.ReadResource(sync)
}

type VnMonitoringPathAnalyzerTestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vn_monitoring.VnMonitoringClient
	Res    *oci_vn_monitoring.GetPathAnalyzerTestResponse
}

func (s *VnMonitoringPathAnalyzerTestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VnMonitoringPathAnalyzerTestDataSourceCrud) Get() error {
	request := oci_vn_monitoring.GetPathAnalyzerTestRequest{}

	if pathAnalyzerTestId, ok := s.D.GetOkExists("path_analyzer_test_id"); ok {
		tmp := pathAnalyzerTestId.(string)
		request.PathAnalyzerTestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vn_monitoring")

	response, err := s.Client.GetPathAnalyzerTest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VnMonitoringPathAnalyzerTestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DestinationEndpoint != nil {
		destinationEndpointArray := []interface{}{}
		if destinationEndpointMap := EndpointToMap(&s.Res.DestinationEndpoint); destinationEndpointMap != nil {
			destinationEndpointArray = append(destinationEndpointArray, destinationEndpointMap)
		}
		s.D.Set("destination_endpoint", destinationEndpointArray)
	} else {
		s.D.Set("destination_endpoint", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}

	if s.Res.ProtocolParameters != nil {
		protocolParametersArray := []interface{}{}
		if protocolParametersMap := ProtocolParametersToMap(&s.Res.ProtocolParameters); protocolParametersMap != nil {
			protocolParametersArray = append(protocolParametersArray, protocolParametersMap)
		}
		s.D.Set("protocol_parameters", protocolParametersArray)
	} else {
		s.D.Set("protocol_parameters", nil)
	}

	if s.Res.QueryOptions != nil {
		s.D.Set("query_options", []interface{}{QueryOptionsToMap(s.Res.QueryOptions)})
	} else {
		s.D.Set("query_options", nil)
	}

	if s.Res.SourceEndpoint != nil {
		sourceEndpointArray := []interface{}{}
		if sourceEndpointMap := EndpointToMap(&s.Res.SourceEndpoint); sourceEndpointMap != nil {
			sourceEndpointArray = append(sourceEndpointArray, sourceEndpointMap)
		}
		s.D.Set("source_endpoint", sourceEndpointArray)
	} else {
		s.D.Set("source_endpoint", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

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
