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

func ContainerengineClusterNamespaceProfileVersionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_namespace_profile_version_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerengineClusterNamespaceProfileVersionResource(), fieldMap, readSingularContainerengineClusterNamespaceProfileVersion)
}

func readSingularContainerengineClusterNamespaceProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterNamespaceProfileVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterNamespaceProfileVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetClusterNamespaceProfileVersionResponse
}

func (s *ContainerengineClusterNamespaceProfileVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterNamespaceProfileVersionDataSourceCrud) Get() error {
	request := oci_containerengine.GetClusterNamespaceProfileVersionRequest{}

	if clusterNamespaceProfileVersionId, ok := s.D.GetOkExists("cluster_namespace_profile_version_id"); ok {
		tmp := clusterNamespaceProfileVersionId.(string)
		request.ClusterNamespaceProfileVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetClusterNamespaceProfileVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterNamespaceProfileVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdminClusterRoleName != nil {
		s.D.Set("admin_cluster_role_name", *s.Res.AdminClusterRoleName)
	}

	allowedNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.AllowedNamespaceAnnotations {
		allowedNamespaceAnnotations = append(allowedNamespaceAnnotations, AllowedNamespaceAnnotationToMap(item, true))
	}
	s.D.Set("allowed_namespace_annotations", allowedNamespaceAnnotations)

	allowedNamespaceLabels := []interface{}{}
	for _, item := range s.Res.AllowedNamespaceLabels {
		allowedNamespaceLabels = append(allowedNamespaceLabels, AllowedNamespaceLabelToMap(item, true))
	}
	s.D.Set("allowed_namespace_labels", allowedNamespaceLabels)

	if s.Res.ClusterNamespaceProfileId != nil {
		s.D.Set("cluster_namespace_profile_id", *s.Res.ClusterNamespaceProfileId)
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

	fixedNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.FixedNamespaceAnnotations {
		fixedNamespaceAnnotations = append(fixedNamespaceAnnotations, NamespaceAnnotationToMap(item))
	}
	s.D.Set("fixed_namespace_annotations", fixedNamespaceAnnotations)

	fixedNamespaceLabels := []interface{}{}
	for _, item := range s.Res.FixedNamespaceLabels {
		fixedNamespaceLabels = append(fixedNamespaceLabels, NamespaceLabelToMap(item))
	}
	s.D.Set("fixed_namespace_labels", fixedNamespaceLabels)

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDeprecated != nil {
		s.D.Set("is_deprecated", *s.Res.IsDeprecated)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	requiredNamespaceAnnotations := []interface{}{}
	for _, item := range s.Res.RequiredNamespaceAnnotations {
		requiredNamespaceAnnotations = append(requiredNamespaceAnnotations, RequiredNamespaceAnnotationToMap(item, true))
	}
	s.D.Set("required_namespace_annotations", requiredNamespaceAnnotations)

	requiredNamespaceLabels := []interface{}{}
	for _, item := range s.Res.RequiredNamespaceLabels {
		requiredNamespaceLabels = append(requiredNamespaceLabels, RequiredNamespaceLabelToMap(item, true))
	}
	s.D.Set("required_namespace_labels", requiredNamespaceLabels)

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
