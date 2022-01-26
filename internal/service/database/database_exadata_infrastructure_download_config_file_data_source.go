// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseExadataInfrastructureDownloadConfigFileDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseExadataInfrastructureDownloadConfigFile,
		Schema: map[string]*schema.Schema{
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseExadataInfrastructureDownloadConfigFile(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureDownloadConfigFileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadataInfrastructureDownloadConfigFileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.DownloadExadataInfrastructureConfigFileResponse
}

func (s *DatabaseExadataInfrastructureDownloadConfigFileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadataInfrastructureDownloadConfigFileDataSourceCrud) Get() error {
	request := oci_database.DownloadExadataInfrastructureConfigFileRequest{}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.DownloadExadataInfrastructureConfigFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExadataInfrastructureDownloadConfigFileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadataInfrastructureDownloadConfigFileDataSource-", DatabaseExadataInfrastructureDownloadConfigFileDataSource(), s.D))

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)

	if err != nil {
		log.Printf("unable to read 'content' from response. Error: %v", err)
	} else if base64EncodeContent {
		s.D.Set("content", base64.StdEncoding.EncodeToString(contentArray))
	} else {
		s.D.Set("content", string(contentArray))
	}

	return nil
}
