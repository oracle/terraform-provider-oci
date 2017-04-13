package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
)

type ResourceClientTestSuite struct {
	suite.Suite
}

func (s *ResourceClientTestSuite) TestClientInit() {
	var c client.BareMetalClient
	c = &baremetal.Client{}
	s.Require().True(c != nil) // This won't compile if it's not
}
