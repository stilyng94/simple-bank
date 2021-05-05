package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {
	account, err := createFakeAccountWithFakeUser(context.Background())
	require.NoError(t, err)

	testCases := []struct {
		name          string
		accountID     string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{name: "OK",
			accountID: account.ID.String(),
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{name: "NotFound",
			accountID: "e1994ed2-34c0-4d73-8a1c-936e2b6c9689",
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%s", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}

}
