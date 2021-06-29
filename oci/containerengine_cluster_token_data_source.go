// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v42/containerengine"
)

func init() {
	RegisterDatasource("oci_containerengine_cluster_token", ContainerengineClusterTokenDataSource())
}

func ContainerengineClusterTokenDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterToken,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
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

func readSingularContainerengineClusterToken(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterTokenDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient()

	return ReadResource(sync)
}

type ContainerengineClusterTokenDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    string
}

func (s *ContainerengineClusterTokenDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterTokenDataSourceCrud) Get() error {
	var clusterId string
	if clusterIdD, ok := s.D.GetOk("cluster_id"); ok {
		clusterId = clusterIdD.(string)
	}
	clusterUrl := fmt.Sprintf("%s/cluster_request/%s", s.Client.Endpoint(), clusterId)

	signRequest, err := http.NewRequest("GET", clusterUrl, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	signRequest.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	err = s.Client.Signer.Sign(signRequest)

	tokenUrl, err := url.Parse(clusterUrl)
	if err != nil {
		return fmt.Errorf("failed to parse url for token: %w", err)
	}
	q := tokenUrl.Query()
	q.Set("authorization", signRequest.Header.Get("authorization"))
	q.Set("date", signRequest.Header.Get("date"))
	tokenUrl.RawQuery = q.Encode()

	urlBytes, err := tokenUrl.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal url for token: %w", err)
	}
	urlEncoded := base64.URLEncoding.EncodeToString(urlBytes)

	s.Res = urlEncoded
	return nil
}

func (s *ContainerengineClusterTokenDataSourceCrud) SetData() error {
	if s.Res == "" {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ContainerengineClusterTokenDataSource-", ContainerengineClusterTokenDataSource(), s.D))

	s.D.Set("content", s.Res)

	return nil
}
