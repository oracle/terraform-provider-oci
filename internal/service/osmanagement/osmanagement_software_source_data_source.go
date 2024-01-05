// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
)

func OsmanagementSoftwareSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["software_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsmanagementSoftwareSourceResource(), fieldMap, readSingularOsmanagementSoftwareSource)
}

func readSingularOsmanagementSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementSoftwareSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.GetSoftwareSourceResponse
}

func (s *OsmanagementSoftwareSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementSoftwareSourceDataSourceCrud) Get() error {
	request := oci_osmanagement.GetSoftwareSourceRequest{}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementSoftwareSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arch_type", s.Res.ArchType)

	associatedManagedInstances := []interface{}{}
	for _, item := range s.Res.AssociatedManagedInstances {
		associatedManagedInstances = append(associatedManagedInstances, IdToMap(item))
	}
	s.D.Set("associated_managed_instances", associatedManagedInstances)

	s.D.Set("checksum_type", s.Res.ChecksumType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.GpgKeyFingerprint != nil {
		s.D.Set("gpg_key_fingerprint", *s.Res.GpgKeyFingerprint)
	}

	if s.Res.GpgKeyId != nil {
		s.D.Set("gpg_key_id", *s.Res.GpgKeyId)
	}

	if s.Res.GpgKeyUrl != nil {
		s.D.Set("gpg_key_url", *s.Res.GpgKeyUrl)
	}

	if s.Res.MaintainerEmail != nil {
		s.D.Set("maintainer_email", *s.Res.MaintainerEmail)
	}

	if s.Res.MaintainerName != nil {
		s.D.Set("maintainer_name", *s.Res.MaintainerName)
	}

	if s.Res.MaintainerPhone != nil {
		s.D.Set("maintainer_phone", *s.Res.MaintainerPhone)
	}

	if s.Res.Packages != nil {
		s.D.Set("packages", *s.Res.Packages)
	}

	if s.Res.ParentId != nil {
		s.D.Set("parent_id", *s.Res.ParentId)
	}

	if s.Res.ParentName != nil {
		s.D.Set("parent_name", *s.Res.ParentName)
	}

	if s.Res.RepoType != nil {
		s.D.Set("repo_type", *s.Res.RepoType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}
