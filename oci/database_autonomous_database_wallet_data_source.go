// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDatabaseWalletDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAutonomousDatabaseWallet,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
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

func readSingularAutonomousDatabaseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDatabaseWalletDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDatabaseWalletDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *[]byte
}

func (s *AutonomousDatabaseWalletDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDatabaseWalletDataSourceCrud) Get() error {
	request := oci_database.GenerateAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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

func (s *AutonomousDatabaseWalletDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("content", string(*s.Res))

	return nil
}
