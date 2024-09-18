// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineClusterResource(), fieldMap, readSingularContainerengineCluster)
}

func readSingularContainerengineCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterResponse
}

func (s *ContainerengineClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterDataSourceCrud) Get() error {
	request := oci_containerengine.GetClusterRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("available_kubernetes_upgrades", s.Res.AvailableKubernetesUpgrades)

	clusterPodNetworkOptions := []interface{}{}
	for _, item := range s.Res.ClusterPodNetworkOptions {
		clusterPodNetworkOptions = append(clusterPodNetworkOptions, ClusterPodNetworkOptionDetailsToMap(item))
	}
	s.D.Set("cluster_pod_network_options", clusterPodNetworkOptions)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EndpointConfig != nil {
		s.D.Set("endpoint_config", []interface{}{ClusterEndpointConfigToMap(s.Res.EndpointConfig, true)})
	} else {
		s.D.Set("endpoint_config", nil)
	}

	if s.Res.Endpoints != nil {
		s.D.Set("endpoints", []interface{}{ClusterEndpointsToMap(s.Res.Endpoints)})
	} else {
		s.D.Set("endpoints", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImagePolicyConfig != nil {
		s.D.Set("image_policy_config", []interface{}{ImagePolicyConfigToMap(s.Res.ImagePolicyConfig)})
	} else {
		s.D.Set("image_policy_config", nil)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ClusterMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Options != nil {
		s.D.Set("options", []interface{}{ClusterCreateOptionsToMap(s.Res.Options)})
	} else {
		s.D.Set("options", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
