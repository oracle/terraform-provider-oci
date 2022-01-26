// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseAutonomousDatabaseWalletDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousDatabaseWallet,
		Schema: map[string]*schema.Schema{
			"autonomous_database_id": {
				Type:       schema.TypeString,
				Required:   true,
				Deprecated: tfresource.ResourceDeprecatedForAnother("data.oci_database_autonomous_database_wallet", "oci_database_autonomous_database_wallet"),
			},
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"generate_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "SINGLE",
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseAutonomousDatabaseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseWalletDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseWalletDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *[]byte
}

func (s *DatabaseAutonomousDatabaseWalletDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseWalletDataSourceCrud) Get() error {
	request := oci_database.GenerateAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if generateType, ok := s.D.GetOkExists("generate_type"); ok {
		request.GenerateType = oci_database.GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum(generateType.(string))
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GenerateAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if response.Content != nil {
		defer response.Content.Close()
		if contentBytes, err := ioutil.ReadAll(response.Content); err == nil {
			s.Res = &contentBytes
		} else {
			return err
		}
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseWalletDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseWalletDataSource-", DatabaseAutonomousDatabaseWalletDataSource(), s.D))

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	if base64EncodeContent {
		s.D.Set("content", base64.StdEncoding.EncodeToString(*s.Res))
	} else {
		s.D.Set("content", string(*s.Res))
	}

	return nil
}
