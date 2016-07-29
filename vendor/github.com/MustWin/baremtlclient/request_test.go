package baremtlclient

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func getTestClient() (c *Client, e error) {
	gopath := os.Getenv("GOPATH")
	// TODO: Note if the repository name changes the path will also have to be
	// changed in order to pass tests
	pemPath := path.Join(
		gopath,
		"src",
		"github.com",
		"MustWin",
		"baremtlclient",
		"test",
		"data",
		"private.pem",
	)

	c, e = NewFromKeyPath("userOCID", "tenancyOCID", "fingerprint", pemPath, "password")
	return
}

type RequestTestSuite struct {
	suite.Suite
	client    *Client
	transport *http.Transport
}

func (s *RequestTestSuite) SetupTest() {
	var e error
	s.client, e = getTestClient()
	if e != nil {
		panic("Test setup failed")
	}

	s.transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

}

func TestRunRequestTests(t *testing.T) {
	suite.Run(t, new(RequestTestSuite))
}

func (s *RequestTestSuite) TestDeleteRequest() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Deleted group 1"))
		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	url, _ := url.Parse(fmt.Sprintf("%s/%s/%s", ts.URL, resourceGroups, "123"))

	e := s.client.identityAPI.deleteRequest(url.String(), nil)
	s.Nil(e)

}

func (s *RequestTestSuite) TestUnsuccessfulDeleteRequest() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("opc-request-id", "1234567890")

		w.WriteHeader(http.StatusNonAuthoritativeInfo)

		error := Error{
			Code:    "42",
			Message: "ultimate answer",
		}

		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.Encode(error)

		w.Write(buffer.Bytes())

		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	url, _ := url.Parse(fmt.Sprintf("%s/%s/%s", ts.URL, resourceGroups, "123"))

	e := s.client.identityAPI.deleteRequest(url.String(), nil)
	s.NotNil(e)
	s.Equal(e.Error(), "Code: 42; OPC Request ID: 1234567890; Message: ultimate answer")

}

func (s *RequestTestSuite) TestGetRequest() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		u := []Resource{
			Resource{
				ID:            "0123456789",
				CompartmentID: r.URL.Query().Get("compartmentId"),
				Name:          "Bob",
				Description:   "Bob's name",
				TimeCreated:   time.Now(),
				TimeModified:  time.Now(),
				State:         ResourceCreated,
			},
		}

		buff, _ := json.Marshal(u)

		w.Write(buff)
		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	request := ListResourceRequest{
		CompartmentID: s.client.authInfo.tenancyOCID,
	}

	query, _ := request.buildQuery(resourceUsers)
	url := ts.URL + "/users?" + query

	resp, e := s.client.identityAPI.getRequest(url, nil)
	s.Nil(e)

	buffer := bytes.NewBuffer(resp.body)
	decoder := json.NewDecoder(buffer)
	var users []Resource
	e = decoder.Decode(&users)
	s.Nil(e)

	s.Equal(len(users), 1)
	s.Equal(users[0].ID, "0123456789")

}

func (s *RequestTestSuite) TestUnsuccessfulGetRequest() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("opc-request-id", "1234567890")

		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		err := Error{
			Message: "foo",
			Code:    "bar",
		}

		buff, _ := json.Marshal(err)

		w.Write(buff)
		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	url := ts.URL + "/users?compartmentId=1"

	resp, e := s.client.identityAPI.getRequest(url, nil)
	s.NotNil(e)
	s.Nil(resp)
	s.Equal(e.Error(), "Code: bar; OPC Request ID: 1234567890; Message: foo")

}

func (s *RequestTestSuite) TestUnsuccessfulRequest() {

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("opc-request-id", "1234567890")
		// anything other than 200 is an error
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		err := Error{
			Message: "foo",
			Code:    "bar",
		}

		buff, _ := json.Marshal(err)

		w.Write(buff)
		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	request := CreateResourceRequest{
		CompartmentID: "xyz",
		Name:          "Bob",
		Description:   "123abc",
	}

	_, e := s.client.identityAPI.request(http.MethodPost, ts.URL, request, nil)

	s.NotNil(e)
	s.Equal(e.Error(), "Code: bar; OPC Request ID: 1234567890; Message: foo")

}

func (s *RequestTestSuite) TestSuccessfulRequest() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var createUser CreateResourceRequest
		decoder := json.NewDecoder(r.Body)

		decoder.Decode(&createUser)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		u := Resource{
			ID:            "0123456789",
			CompartmentID: createUser.CompartmentID,
			Name:          createUser.Name,
			Description:   createUser.Description,
			TimeCreated:   time.Now(),
			TimeModified:  time.Now(),
			State:         ResourceCreated,
		}

		buff, _ := json.Marshal(u)

		w.Write(buff)
		return

	}))

	s.client.identityAPI = newAPIRequestor(s.client.authInfo, s.transport)

	request := CreateResourceRequest{
		CompartmentID: "xyz",
		Name:          "Bob",
		Description:   "123abc",
	}

	response, e := s.client.identityAPI.request(http.MethodPost, ts.URL, request, nil)

	if !s.Nil(e) {
		s.T().Log(e.Error())
	}

	var user Resource
	e = json.Unmarshal(response, &user)

	s.Nil(e)
	s.Equal(user.CompartmentID, request.CompartmentID)
	s.Equal(user.Name, request.Name)
	s.Equal(user.Description, request.Description)

}

func TestAddRequestHeaders(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://www.goo.com", nil)

	addRequiredRequestHeaders(request, []byte{})
	assert.Equal(t, request.Header.Get("content-type"), "application/json")

	request, _ = http.NewRequest(http.MethodGet, "http://www.goo.com", nil)

	request.Header.Add("content-type", "something")

	addRequiredRequestHeaders(request, []byte{})

	assert.Equal(t, request.Header.Get("content-type"), "something")

	buffer := []byte("12345")
	body := bytes.NewBuffer(buffer)
	request, _ = http.NewRequest(http.MethodPost, "http://www.postme.com", body)

	addRequiredRequestHeaders(request, buffer)

	assert.Equal(t, request.Header.Get("content-length"), "5")

}

func TestGetSigningString(t *testing.T) {
	request, _ := http.NewRequest(
		http.MethodGet,
		"https://core.us-az-phoenix-1.oracleiaas.com/v1/instances?availabilityDomain=Pjwf%3A%20PHX-AD-1",
		nil,
	)

	request.Header.Add("date", "Thu, 05 Jan 2014 21:31:40 GMT")
	addRequiredRequestHeaders(request, []byte{})
	actual := getSigningString(request)
	expected := `date: Thu, 05 Jan 2014 21:31:40 GMT
(request-target): get /v1/instances?availabilityDomain=Pjwf%3A%20PHX-AD-1`

	if !assert.Equal(t, actual, expected) {
		t.Log("Actual   ", actual)
		t.Log("Expected ", expected)
	}

	buffer := []byte("{'foo':'bar'}")
	body := bytes.NewBuffer(buffer)
	request, _ = http.NewRequest(
		http.MethodPost,
		"https://core.us-az-phoenix-1.oracleiaas.com/v1/instances?availabilityDomain=Pjwf%3A%20PHX-AD-1",
		body,
	)

	request.Header.Add("date", "Thu, 05 Jan 2014 21:31:40 GMT")
	addRequiredRequestHeaders(request, buffer)
	actual = getSigningString(request)
	expected = "date: Thu, 05 Jan 2014 21:31:40 GMT\n" +
		"(request-target): post /v1/instances?availabilityDomain=Pjwf%3A%20PHX-AD-1\n" +
		"content-length: 13\n" +
		"content-type: application/json\n" +
		"x-content-sha256: " + getBodyHash([]byte("{'foo':'bar'}"))

	if !assert.Equal(t, actual, expected) {
		t.Log("Actual   ", actual)
		t.Log("Expected ", expected)
	}

}

func TestCreateAuthorizationHeader(t *testing.T) {

	testGetURI := "https://core.us-az-phoenix-1.oracleiaas.com/v1/instances" +
		"?availabilityDomain=Pjwf%3A%20PHX-AD-1" +
		"&compartmentId=ocid1.compartment.oc1..aaaaaaaayzim47sto5wqh5d4vugrsx566gjqmflvhlifte3p5ez3miy6e4lq" +
		"&displayName=TeamXInstances" +
		"&volumeId=ocid1.volume.oc1.phx.abyhqljrav2k323acohquoxszz2zyh5vj5v2gbvntg7ifd4ndusyvr332whq"

	expected := "Signature version=\"1\" signature=\"R5I2tf3iU5ExvtywRi2fj4YxjDBuh" +
		"JT7TQwiK1XOU5Wf/hLgq25iLdX8YbRJWvOaLHOuDShhZeisODl/ksVSJISDArLe+cLailYmYPWB7T" +
		"3987U7IgtbhgucHw4bY09MGoRn3rHEfWYTj16C4O2y7zMRmdUwt3f2ioAe1EFrn8bixEM+AavCU/" +
		"ydLFCcxXr13pDSP+NAPvJ0dsyRyBzkYbuYPRulncBYEmFqVxFRARHzIAO7z0OBv8lkoGQTJhKI/5" +
		"ZZxnYmYfwgvM6djK57QdoBSXyrcwi2BdeiBjdhLRphjbWmB5l0OlWeQo6sEEFVcGOzuxazO0XTRw" +
		"baiJYfng==\",headers=\"date (request-target)\",algorithm=\"rsa-sha256\"" +
		",keyId=\"ocid1.tenancy.oc1..aaaaaaaaq3hulfjvrouw3e6qx2ncxtp256aq7etiabqqtz" +
		"unnhxjslzkfyxq/ocid1.user.oc1..aaaaaaaaflxvsdpjs5ztahmsf7vjxy5kdqnuzyqpvwnn" +
		"cbkfhavexwd4w5ra/b4:8a:7d:54:e6:81:04:b2:99:8e:b3:ed:10:e2:12:2b\""

	var privateKey *rsa.PrivateKey
	var e error

	if privateKey, e = PrivateKeyFromBytes(testPrivateKey, "password"); e != nil {
		t.Error("Couldn't create private key", e)
	}

	ai := &authenticationInfo{
		privateRSAKey:  privateKey,
		tenancyOCID:    testTenancyOCID,
		userOCID:       testUserOCID,
		keyFingerPrint: testKeyFingerPrint,
	}

	request, _ := http.NewRequest(
		http.MethodGet,
		testGetURI,
		nil,
	)
	// Set date to be the same date we used to generate test auth header, otherwise
	// sig will be different every time
	request.Header.Add("date", "Thu, 05 Jan 2014 21:31:40 GMT")

	e = createAuthorizationHeader(request, ai, []byte{})

	assert.Nil(t, e)

	authHeader := request.Header.Get("Authorization")

	if !assert.Equal(t, authHeader, expected) {
		t.Log("Actual   ", authHeader)
		t.Log("Expected ", expected)
	}

}

func TestConcatenateHeaders(t *testing.T) {
	headers := []string{
		"foo",
		"bar",
		"baz",
	}
	expected := "foo bar baz"
	actual := concantenateHeaders(headers)

	assert.Equal(t, actual, expected)

	headers = []string{"foo"}

	expected = "foo"
	actual = concantenateHeaders(headers)

	assert.Equal(t, actual, expected)

	headers = []string{}
	expected = ""
	actual = concantenateHeaders(headers)

	assert.Equal(t, actual, expected)

}
