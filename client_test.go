package main

import (
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceClientTestSuite struct {
	suite.Suite
}

func (s *ResourceClientTestSuite) TestClientInit() {
	var c *baremetal.Client
	c = &baremetal.Client{}
	s.Require().True(c != nil) // This won't compile if it's not
}
