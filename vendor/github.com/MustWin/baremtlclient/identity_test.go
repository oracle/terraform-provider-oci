package baremtlclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockRequestor struct {
	*Client
	mock.Mock
}

func (mr *mockRequestor) request(method, url string, body interface{}, headers http.Header) (respBody []byte, e error) {
	args := mr.Called(method, url, body, headers)
	return args.Get(0).([]byte), args.Error(1)
}

func (mr *mockRequestor) getRequest(urlStr string, headers http.Header) (resp *getResponse, e error) {
	args := mr.Called(urlStr, headers)
	return args.Get(0).(*getResponse), args.Error(1)

}

func (mr *mockRequestor) deleteRequest(urlStr string, headers http.Header) (e error) {
	args := mr.Called(urlStr, headers)
	return args.Error(0)
}

func newMockForTest() (m *mockRequestor) {
	m = new(mockRequestor)
	c := createClientForTest()

	m.Client = c
	m.identityAPI = m

	return
}

func createClientForTest() (c *Client) {
	key, _ := PrivateKeyFromBytes(testPrivateKey, "password")

	c = New(testUserOCID, testTenancyOCID, testKeyFingerPrint, key)
	return
}

func marshalObjectForTest(obj interface{}) []byte {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(obj)

	return buffer.Bytes()

}

type IdentityTestSuite struct {
	suite.Suite
	requestor *mockRequestor
	nilHeader http.Header
}

func (s *IdentityTestSuite) SetupTest() {
	s.requestor = newMockForTest()
}

func (s *IdentityTestSuite) TestListAvailabilityDomains() {
	response := &getResponse{
		body: marshalObjectForTest(
			[]AvailabilityDomain{
				AvailabilityDomain{
					Name:          "one",
					CompartmentID: "1",
				},
				AvailabilityDomain{
					Name:          "two",
					CompartmentID: "1",
				},
			},
		),
	}

	s.requestor.On(
		"getRequest",
		buildIdentityURL(
			resourceAvailabilityDomains,
			&url.Values{
				"compartmentId": []string{"1"},
			},
		),
		s.nilHeader,
	).Return(response, nil)

	actual, e := s.requestor.ListAvailablityDomains("1")
	s.requestor.AssertExpectations(s.T())

	s.Nil(e)
	s.NotNil(actual)
	s.Equal(len(actual), 2)

}

func (s *IdentityTestSuite) TestGetUserGroupMemberships() {
	gms := UserGroupMembership{}
	gms.ID = "1234"
	gms.GroupID = "4567"
	gms.UserID = "bob123"

	response := &getResponse{
		body: marshalObjectForTest(gms),
	}

	s.requestor.On(
		"getRequest",
		buildIdentityURL(resourceUserGroupMemberships, nil, "1234"),
		s.nilHeader,
	).Return(
		response,
		nil,
	)

	actual, e := s.requestor.GetUserGroupMembership("1234")
	s.requestor.AssertExpectations(s.T())

	s.Nil(e)
	s.NotNil(actual)
	s.Equal(actual.ID, "1234")
	s.Equal(actual.GroupID, gms.GroupID)
	s.Equal(actual.UserID, gms.UserID)
}

func (s *IdentityTestSuite) TestGetPolicy() {
	policy := Policy{}
	policy.ID = "1234"
	policy.Statements = []string{
		"statement 1",
		"statement 2",
	}

	response := &getResponse{
		body: marshalObjectForTest(policy),
	}

	s.requestor.On(
		"getRequest",
		buildIdentityURL(resourcePolicies, nil, "1234"),
		s.nilHeader,
	).Return(
		response,
		nil,
	)

	actual, e := s.requestor.GetPolicy("1234")
	s.requestor.AssertExpectations(s.T())

	s.Nil(e)
	s.NotNil(actual)
	s.Equal(actual.ID, "1234")
	s.Equal(len(actual.Statements), 2)
}

func (s *IdentityTestSuite) TestGetCompartment() {
	s.testGetResource(resourceCompartments, "123", s.requestor.GetCompartment)
}

func (s *IdentityTestSuite) testGetResource(resource resourceName, id string,
	funcUnderTest func(id string) (*Resource, error)) {

	expected := Resource{
		ID: id,
	}

	response := &getResponse{
		body: marshalObjectForTest(expected),
	}

	s.requestor.On(
		"getRequest",
		buildIdentityURL(resource, nil, id),
		s.nilHeader,
	).Return(
		response,
		nil,
	)

	actual, e := funcUnderTest(id)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
	s.NotNil(actual)
	s.Equal(expected.ID, actual.ID)
}

func (s *IdentityTestSuite) TestGetUser() {
	s.testGetResource(resourceUsers, "123", s.requestor.GetUser)
}

func (s *IdentityTestSuite) TestGetGroup() {
	s.testGetResource(resourceGroups, "345", s.requestor.GetGroup)
}

func (s *IdentityTestSuite) testListResources(resource resourceName,
	request ListResourceRequest,
	response ListResourceResponse,
	funcUnderTest func(*ListResourceRequest) (*ListResourceResponse, error)) {

	getResp := &getResponse{
		body: marshalObjectForTest(response.Items),
	}

	s.requestor.On(
		"getRequest",
		buildIdentityURL(resource, &url.Values{"compartmentId": []string{request.CompartmentID}}),
		s.nilHeader,
	).Return(
		getResp,
		nil,
	)

	actual, e := funcUnderTest(&request)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
	s.NotNil(actual)
	s.Equal(len(response.Items), len(actual.Items))
}

func (s *IdentityTestSuite) TestListGroups() {
	request := ListResourceRequest{
		CompartmentID: "1",
	}
	response := ListResourceResponse{
		Items: []Resource{
			Resource{
				ID:   "1",
				Name: "one",
			},
			Resource{
				ID:   "2",
				Name: "two",
			},
		},
	}

	s.testListResources(resourceGroups, request, response, s.requestor.ListGroups)

}

func (s *IdentityTestSuite) TestListCompartments() {
	request := ListResourceRequest{
		CompartmentID: "1",
	}

	response := ListResourceResponse{
		Items: []Resource{
			Resource{
				ID:            "123",
				CompartmentID: "1",
			},
			Resource{
				ID:            "234",
				CompartmentID: "1",
			},
			Resource{
				ID:            "567",
				CompartmentID: "1",
			},
		},
	}
	s.testListResources(resourceCompartments, request, response, s.requestor.ListCompartments)
}

func (s *IdentityTestSuite) TestListUsers() {
	request := ListResourceRequest{
		CompartmentID: "1",
	}

	response := ListResourceResponse{
		Items: []Resource{
			Resource{
				ID:            "123",
				CompartmentID: "1",
			},
			Resource{
				ID:            "234",
				CompartmentID: "1",
			},
			Resource{
				ID:            "567",
				CompartmentID: "1",
			},
		},
	}

	s.testListResources(resourceUsers, request, response, s.requestor.ListUsers)

}

var expectedID = "1234567890"

var createReq = CreateResourceRequest{
	CompartmentID: "ocid1.tenancy.oc1..aaaaaaaaq3hulfjvrouw3e6qx2ncxtp256aq7etiabqqtzunnhxjslzkfyxq",
	Name:          "name",
	Description:   "a name",
}

var resrc = Resource{
	Name:        "name",
	Description: "a name",
	ID:          expectedID,
}

func (s *IdentityTestSuite) GetCreateGroup() {
	s.testCreateResource(resourceGroups, createReq, resrc, s.requestor.CreateGroup)
}

func (s *IdentityTestSuite) TestCreateUser() {
	//	s.testCreateResource(resourceUsers, createReq, resrc, s.requestor.CreateUser)
	url := buildIdentityURL(resourceUsers, nil)
	headers := http.Header{
		"opc-idempotency-token": []string{"54321"},
	}

	s.requestor.On(
		"request",
		http.MethodPost,
		url,
		createReq,
		headers,
	).Return(
		marshalObjectForTest(resrc),
		nil,
	)

	actual, e := s.requestor.CreateUser(
		"name",
		"a name",
		Options{
			OPCIdempotencyToken: "54321",
		},
	)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
	s.NotNil(actual)
	s.Equal(actual.ID, expectedID)

}

func (s *IdentityTestSuite) testCreateResource(resource resourceName, req CreateResourceRequest, res Resource, f func(CreateResourceRequest, http.Header) (*Resource, error)) {

	url := buildIdentityURL(resource, nil)

	s.requestor.On(
		"request",
		http.MethodPost,
		url,
		createReq,
		s.nilHeader,
	).Return(
		marshalObjectForTest(res),
		nil,
	)

	actual, e := f(req, nil)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
	s.NotNil(actual)
	s.Equal(actual.ID, expectedID)

}

func (s *IdentityTestSuite) testUpdateResource(resource resourceName, id string, req UpdateResourceRequest, expected Resource,
	funcUnderTest func(string, UpdateResourceRequest, http.Header) (*Resource, error)) {

	s.requestor.On(
		"request",
		http.MethodPut,
		buildIdentityURL(resource, nil, id),
		req,
		s.nilHeader,
	).Return(
		marshalObjectForTest(expected),
		nil,
	)

	actual, e := funcUnderTest(id, req, nil)
	s.requestor.AssertExpectations(s.T())
	s.Nil(e)
	s.NotNil(actual)
	s.Equal(expected.Description, actual.Description)

}

func (s *IdentityTestSuite) TestUpdateCompartment() {
	req := UpdateResourceRequest{
		Description: "some description",
	}
	expected := Resource{
		ID:          "123",
		Description: "some desription",
	}

	s.testUpdateResource(resourceCompartments, "123", req, expected, s.requestor.UpdateCompartment)
}

func (s *IdentityTestSuite) TestCreateCompartment() {
	s.testCreateResource(resourceCompartments, createReq, resrc, s.requestor.CreateCompartment)
}

func (s *IdentityTestSuite) TestBuildQuery() {

	req := ListResourceRequest{}
	resp, e := req.buildQuery(resourceUsers)
	s.NotNil(e)
	s.Equal(resp, "")

	req = ListResourceRequest{
		CompartmentID: "123456",
	}
	expected := buildIdentityURL(resourceUsers, &url.Values{
		"compartmentId": []string{"123456"},
	})
	resp, e = req.buildQuery(resourceUsers)
	s.Nil(e)
	s.Equal(resp, expected)

	expected = buildIdentityURL(resourceUsers, &url.Values{
		"compartmentId": []string{"123456"},
		"limit":         []string{"100"},
		"page":          []string{"my/page"},
	})
	req = ListResourceRequest{
		CompartmentID: "123456",
		Page:          "my/page",
		Limit:         100,
	}
	resp, _ = req.buildQuery(resourceUsers)
	s.Equal(resp, expected)
}

func TestRunIdentityTests(t *testing.T) {
	suite.Run(t, new(IdentityTestSuite))
}
