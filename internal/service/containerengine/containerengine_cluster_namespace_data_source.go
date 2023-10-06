// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterNamespaceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_namespace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineClusterNamespaceResource(), fieldMap, readSingularContainerengineClusterNamespace)
}

func readSingularContainerengineClusterNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterNamespaceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterNamespaceResponse
}

func (s *ContainerengineClusterNamespaceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterNamespaceDataSourceCrud) Get() error {
	request := oci_containerengine.GetClusterNamespaceRequest{}

	if clusterNamespaceId, ok := s.D.GetOkExists("cluster_namespace_id"); ok {
		tmp := clusterNamespaceId.(string)
		request.ClusterNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetClusterNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterNamespaceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("cluster_ids", s.Res.ClusterIds)
	s.D.Set("cluster_ids", s.Res.ClusterIds)

	if s.Res.ClusterNamespaceProfileVersionId != nil {
		s.D.Set("cluster_namespace_profile_version_id", *s.Res.ClusterNamespaceProfileVersionId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	namespaceAnnotations := []interface{}{}
	for _, item := range s.Res.NamespaceAnnotations {
		namespaceAnnotations = append(namespaceAnnotations, NamespaceAnnotationToMap(item))
	}
	s.D.Set("namespace_annotations", namespaceAnnotations)

	namespaceLabels := []interface{}{}
	for _, item := range s.Res.NamespaceLabels {
		namespaceLabels = append(namespaceLabels, NamespaceLabelToMap(item))
	}
	s.D.Set("namespace_labels", namespaceLabels)

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
