package handlers_test

import (
	"Segments/pkg/handlers"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_AddUserSection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		userID   string
		ttl      string
		prepare  func(m *Mockrepo)
		expected func(t assert.TestingT, w *httptest.ResponseRecorder)
	}{
		{
			name:   "Success",
			userID: "1001",
			ttl:    "10",
			prepare: func(m *Mockrepo) {
				m.EXPECT().GetSectionList().Return([]byte(""), nil)
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, 500, w.Code)
				//assert.Equal(t, gomock.Any(), w.Result())
			},
		},
		{
			name:   "Failed",
			userID: "1002",
			ttl:    "1",
			prepare: func(m *Mockrepo) {

				m.EXPECT().GetSectionList().Return([]byte("afasras"), nil)
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, 500, w.Code)
				//assert.Equal(t, gomock.Any(), w.Result())
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := NewMockrepo(ctrl)
			handler := handlers.NewHandler(repoMock)
			tc.prepare(repoMock)

			r, err := createAddRequest(tc.userID, tc.ttl)
			assert.NoError(t, err)

			w := httptest.NewRecorder()

			handler.AddUserSection(w, r)
			tc.expected(t, w)
		})
	}
}

func createAddRequest(paramId string, paramTtl string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost/add/%s/%s", paramId, paramTtl), nil)
	if err != nil {
		return nil, err
	}
	req = mux.SetURLVars(req, map[string]string{"user": paramId, "ttl": paramTtl})
	return req, nil
}
