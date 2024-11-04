package foreman

import (
	"net/http"
	"testing"
)

// ----------------------------------------------------------------------------
// Test Cases for the Unit Test Framework
// ----------------------------------------------------------------------------

// SEE: foreman_api_test.go#TestCRUDFunction_CorrectURLAndMethod()
func DataSourceForemanDiscoveryRuleCorrectURLAndMethodTestCases(t *testing.T) []TestCaseCorrectURLAndMethod {

	obj := RandForemanDiscoveryRule()
	s := ForemanDiscoveryRuleToInstanceState(obj)

	return []TestCaseCorrectURLAndMethod{
		{
			TestCase: TestCase{
				funcName:     "dataSourceForemanDiscoveryRuleRead",
				crudFunc:     dataSourceForemanDiscoveryRuleRead,
				resourceData: MockForemanDiscoveryRuleResourceData(s),
			},
			expectedURIs: []ExpectedUri{
				{
					expectedURI:    DiscoveryRuleURI,
					expectedMethod: http.MethodGet,
				},
			},
		},
	}

}

// SEE: foreman_api_test.go#TestCRUDFunction_RequestDataEmpty()
func DataSourceForemanDiscoveryRuleRequestDataEmptyTestCases(t *testing.T) []TestCase {
	obj := RandForemanDiscoveryRule()
	s := ForemanDiscoveryRuleToInstanceState(obj)

	return []TestCase{
		{
			funcName:     "dataSourceForemanDiscoveryRuleRead",
			crudFunc:     dataSourceForemanDiscoveryRuleRead,
			resourceData: MockForemanDiscoveryRuleResourceData(s),
		},
	}

}

// SEE: foreman_api_test.go#TestCRUDFunction_StatusCodeError()
func DataSourceForemanDiscoveryRuleStatusCodeTestCases(t *testing.T) []TestCase {

	obj := RandForemanDiscoveryRule()
	s := ForemanDiscoveryRuleToInstanceState(obj)

	return []TestCase{
		{
			funcName:     "dataSourceForemanDiscoveryRuleRead",
			crudFunc:     dataSourceForemanDiscoveryRuleRead,
			resourceData: MockForemanDiscoveryRuleResourceData(s),
		},
	}

}

// SEE: foreman_api_test.go#TestCRUDFunction_EmptyResponseError()
func DataSourceForemanDiscoveryRuleEmptyResponseTestCases(t *testing.T) []TestCase {

	obj := RandForemanDiscoveryRule()
	s := ForemanDiscoveryRuleToInstanceState(obj)

	return []TestCase{
		{
			funcName:     "dataSourceForemanDiscoveryRuleRead",
			crudFunc:     dataSourceForemanDiscoveryRuleRead,
			resourceData: MockForemanDiscoveryRuleResourceData(s),
		},
	}

}

// SEE: foreman_api_test.go#TestCRUDFunction_MockResponse()
func DataSourceForemanDiscoveryRuleMockResponseTestCases(t *testing.T) []TestCaseMockResponse {

	obj := RandForemanDiscoveryRule()
	s := ForemanDiscoveryRuleToInstanceState(obj)

	return []TestCaseMockResponse{
		// If the server responds with more than one search result for the data
		// source read, then the operation should return an error
		{
			TestCase: TestCase{
				funcName:     "dataSourceForemanDiscoveryRuleRead",
				crudFunc:     dataSourceForemanDiscoveryRuleRead,
				resourceData: MockForemanDiscoveryRuleResourceData(s),
			},
			responseFile: DiscoveryRuleTestDataPath + "/query_response_multi.json",
			returnError:  true,
		},
		// If the server responds with zero search results for the data source
		// read, then the operation should return an error
		{
			TestCase: TestCase{
				funcName:     "dataSourceForemanDiscoveryRuleRead",
				crudFunc:     dataSourceForemanDiscoveryRuleRead,
				resourceData: MockForemanDiscoveryRuleResourceData(s),
			},
			responseFile: TestDataPath + "/query_response_zero.json",
			returnError:  true,
		},
		// If the server responds with exactly one search result for the data source
		// read, then the operation should succeed and the attributes of the
		// ResourceData should be set properly.
		{
			TestCase: TestCase{
				funcName:     "dataSourceForemanDiscoveryRuleRead",
				crudFunc:     dataSourceForemanDiscoveryRuleRead,
				resourceData: MockForemanDiscoveryRuleResourceData(s),
			},
			responseFile: DiscoveryRuleTestDataPath + "/query_response_single.json",
			returnError:  false,
			expectedResourceData: MockForemanDiscoveryRuleResourceDataFromFile(
				t,
				DiscoveryRuleTestDataPath+"/query_response_single_state.json",
			),
			compareFunc: ForemanDiscoveryRuleResourceDataCompare,
		},
	}

}
