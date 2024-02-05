/*
Candlepin

Testing OwnerContentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package caliri

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func Test_caliri_OwnerContentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OwnerContentAPIService CreateContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerContentAPI.CreateContent(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerContentAPIService CreateContentBatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerContentAPI.CreateContentBatch(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerContentAPIService GetContentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string
		var contentId string

		resp, httpRes, err := apiClient.OwnerContentAPI.GetContentById(context.Background(), ownerKey, contentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerContentAPIService GetContentsByOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerContentAPI.GetContentsByOwner(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerContentAPIService RemoveContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string
		var contentId string

		httpRes, err := apiClient.OwnerContentAPI.RemoveContent(context.Background(), ownerKey, contentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerContentAPIService UpdateContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string
		var contentId string

		resp, httpRes, err := apiClient.OwnerContentAPI.UpdateContent(context.Background(), ownerKey, contentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}