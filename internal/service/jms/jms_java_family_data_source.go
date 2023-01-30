// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaFamilyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsJavaFamily,
		Schema: map[string]*schema.Schema{
			"family_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"doc_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_of_support_life_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsJavaFamily(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaFamilyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaFamilyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetJavaFamilyResponse
}

func (s *JmsJavaFamilyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaFamilyDataSourceCrud) Get() error {
	request := oci_jms.GetJavaFamilyRequest{}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetJavaFamily(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaFamilyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaFamilyDataSource-", JmsJavaFamilyDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DocUrl != nil {
		s.D.Set("doc_url", *s.Res.DocUrl)
	}

	if s.Res.EndOfSupportLifeDate != nil {
		s.D.Set("end_of_support_life_date", s.Res.EndOfSupportLifeDate.String())
	}

	s.D.Set("support_type", s.Res.SupportType)

	return nil
}
