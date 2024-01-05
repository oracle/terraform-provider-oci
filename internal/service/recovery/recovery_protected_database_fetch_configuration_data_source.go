// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectedDatabaseFetchConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularRecoveryProtectedDatabaseFetchConfiguration,
		Schema: map[string]*schema.Schema{
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"configuration_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protected_database_id": {
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

func readSingularRecoveryProtectedDatabaseFetchConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseFetchConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

type RecoveryProtectedDatabaseFetchConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_recovery.DatabaseRecoveryClient
	Res    *oci_recovery.FetchProtectedDatabaseConfigurationResponse
}

func (s *RecoveryProtectedDatabaseFetchConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RecoveryProtectedDatabaseFetchConfigurationDataSourceCrud) Get() error {
	request := oci_recovery.FetchProtectedDatabaseConfigurationRequest{}

	if configurationType, ok := s.D.GetOkExists("configuration_type"); ok {
		request.ConfigurationType = oci_recovery.FetchProtectedDatabaseConfigurationDetailsConfigurationTypeEnum(configurationType.(string))
	}

	if protectedDatabaseId, ok := s.D.GetOkExists("protected_database_id"); ok {
		tmp := protectedDatabaseId.(string)
		request.ProtectedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "recovery")

	response, err := s.Client.FetchProtectedDatabaseConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RecoveryProtectedDatabaseFetchConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RecoveryProtectedDatabaseFetchConfigurationDataSource-", RecoveryProtectedDatabaseFetchConfigurationDataSource(), s.D))

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
