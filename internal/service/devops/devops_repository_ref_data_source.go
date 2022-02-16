// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryRefDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ref_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["repository_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsRepositoryRefResource(), fieldMap, readSingularDevopsRepositoryRef)
}

func readSingularDevopsRepositoryRef(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryRefDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetRefResponse
}

func (s *DevopsRepositoryRefDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryRefDataSourceCrud) Get() error {
	request := oci_devops.GetRefRequest{}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetRef(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryRefDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryRefDataSource-", DevopsRepositoryRefDataSource(), s.D))

	switch v := (s.Res.RepositoryRef).(type) {
	case oci_devops.RepositoryBranch:

		s.D.Set("ref_type", "BRANCH")

		if v.CommitId != nil {
			s.D.Set("commit_id", *v.CommitId)
		}

		if v.FullRefName != nil {
			s.D.Set("full_ref_name", *v.FullRefName)
		}

		if v.RefName != nil {
			s.D.Set("ref_name", *v.RefName)
		}

		if v.RepositoryId != nil {
			s.D.Set("repository_id", *v.RepositoryId)
		}

	case oci_devops.RepositoryTag:

		s.D.Set("ref_type", "TAG")

		if v.ObjectId != nil {
			s.D.Set("object_id", *v.ObjectId)
		}

		if v.FullRefName != nil {
			s.D.Set("full_ref_name", *v.FullRefName)
		}

		if v.RefName != nil {
			s.D.Set("ref_name", *v.RefName)
		}

		if v.RepositoryId != nil {
			s.D.Set("repository_id", *v.RepositoryId)
		}

	default:
		log.Printf("[WARN] Received 'ref_type' of unknown type %v", v)
		return nil

	}

	return nil
}
