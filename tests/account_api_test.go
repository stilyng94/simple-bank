package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple-bank/ent"
	"simple-bank/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetAccountApi(t *testing.T) {
	account := randomAccount()

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

func randomAccount() *ent.Account {
	acc := &ent.Account{ID: uuid.MustParse("e2994ed2-34c0-4d73-8a1c-936e2b6c9685"), Owner: util.RandomOwner(), Balance: util.RandomAmount(), Currency: util.RandomCurrency()}
	return acc
}
