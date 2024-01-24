// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
)

func BdsBdsInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceResource(), fieldMap, readSingularBdsBdsInstance)
}

func readSingularBdsBdsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetBdsInstanceResponse
}

func (s *BdsBdsInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceDataSourceCrud) Get() error {
	request := oci_bds.GetBdsInstanceRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetBdsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BootstrapScriptUrl != nil {
		s.D.Set("bootstrap_script_url", *s.Res.BootstrapScriptUrl)
	}

	if s.Res.CloudSqlDetails != nil {
		s.D.Set("cloud_sql_details", []interface{}{CloudSqlDetailsToMap(s.Res.CloudSqlDetails)})
	} else {
		s.D.Set("cloud_sql_details", nil)
	}

	if s.Res.ClusterDetails != nil {
		s.D.Set("cluster_details", []interface{}{ClusterDetailsToMap(s.Res.ClusterDetails)})
	} else {
		s.D.Set("cluster_details", nil)
	}

	s.D.Set("cluster_profile", s.Res.ClusterProfile)

	s.D.Set("cluster_version", s.Res.ClusterVersion)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCloudSqlConfigured != nil {
		s.D.Set("is_cloud_sql_configured", *s.Res.IsCloudSqlConfigured)
	}

	if s.Res.IsHighAvailability != nil {
		s.D.Set("is_high_availability", *s.Res.IsHighAvailability)
	}

	if s.Res.IsKafkaConfigured != nil {
		s.D.Set("is_kafka_configured", *s.Res.IsKafkaConfigured)
	}

	if s.Res.IsSecure != nil {
		s.D.Set("is_secure", *s.Res.IsSecure)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.NetworkConfig != nil {
		s.D.Set("network_config", []interface{}{NetworkConfigToMap(s.Res.NetworkConfig)})
	} else {
		s.D.Set("network_config", nil)
	}

	nodes := []interface{}{}
	nodeMap := make(map[string]map[string]interface{})
	for _, item := range s.Res.Nodes {
		node := BdsNodeToMap(item)
		nodes = append(nodes, node)
		PopulateNodeTemplate(item, nodeMap)
	}
	s.D.Set("nodes", nodes)
	s.D.Set("master_node", []interface{}{nodeMap["MASTER"]})
	s.D.Set("util_node", []interface{}{nodeMap["UTILITY"]})
	s.D.Set("worker_node", []interface{}{nodeMap["WORKER"]})
	s.D.Set("compute_only_worker_node", []interface{}{nodeMap["COMPUTE_ONLY_WORKER"]})
	s.D.Set("edge_node", []interface{}{nodeMap["EDGE"]})

	if s.Res.NumberOfNodes != nil {
		s.D.Set("number_of_nodes", *s.Res.NumberOfNodes)
	}

	if s.Res.NumberOfNodesRequiringMaintenanceReboot != nil {
		s.D.Set("number_of_nodes_requiring_maintenance_reboot", *s.Res.NumberOfNodesRequiringMaintenanceReboot)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
