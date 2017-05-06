// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"



)

type ResourceLoadBalancerCertificateTestSuite struct {
	suite.Suite
	Client       mockableClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerCertificateTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": Provider(
			func(d *schema.ResourceData) (interface{}, error) {
				return s.Client, nil
			},
		),
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
}

func (s *ResourceLoadBalancerCertificateTestSuite) TestCreateResourceLoadBalancerCertificateMaximal() {
	resourceName := "baremetal_load_balancer_certificate.t"
	config := `
resource "baremetal_load_balancer_certificate" "t" {
  load_balancer_id   = "ocid1.loadbalancer.stub_id"
  ca_certificate     = "stub_ca_certificate"
  certificate_name   = "stub_certificate_name"
  passphrase         = "stub_passphrase"
  private_key        = "stub_private_key"
  public_certificate = "stub_public_certificate"
}
`
	config += testProviderConfig()

	loadBalancerID := "ocid1.loadbalancer.stub_id"
	res := &baremetal.Certificate{
		CACertificate:     "stub_ca_certificate",
		CertificateName:   "stub_certificate_name",
		Passphrase:        "stub_passphrase",
		PrivateKey:        "stub_private_key",
		PublicCertificate: "stub_public_certificate",
	}
	res.RequestID = "stub_opc_request_id"
	opts := &baremetal.LoadBalancerOptions{}
	// opts := (*baremetal.LoadBalancerOptions)(nil)

	deletedRes := &baremetal.Certificate{}
	*deletedRes = *res

	workReqID := "stub_work_req_id"
	s.Client.On(
		"CreateCertificate",
		loadBalancerID,
		res.CertificateName,
		res.CACertificate,
		res.PrivateKey,
		res.Passphrase,
		res.PublicCertificate,
		opts,
	).Return(workReqID, nil)

	workReqCreated := &baremetal.WorkRequest{
		ID:             workReqID,
		LoadBalancerID: loadBalancerID,
		State:          baremetal.WorkRequestSucceeded,
	}
	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	certs := &baremetal.ListCertificates{
		Certificates: []baremetal.Certificate{*res},
	}
	s.Client.On(
		"ListCertificates",
		loadBalancerID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(certs, nil)

	s.Client.On(
		"DeleteCertificate",
		loadBalancerID,
		res.CertificateName,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqID, nil)

	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadBalancerID),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", res.CertificateName),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", res.CACertificate),
					resource.TestCheckResourceAttr(resourceName, "private_key", res.PrivateKey),
					resource.TestCheckResourceAttr(resourceName, "passphrase", res.Passphrase),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", res.PublicCertificate),
				),
			},
		},
	})
}

func TestResourceLoadBalancerCertificateTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerCertificateTestSuite))
}
