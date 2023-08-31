package handlers_test

import (
	"Segments/pkg/handlers"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
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
				m.EXPECT().AddSection(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, 400, w.Code)
				//assert.Equal(t, gomock.Any(), w.Result())
			},
		},
		{
			name:   "Failed",
			userID: "1002",
			ttl:    "1",
			prepare: func(m *Mockrepo) {
				m.EXPECT().AddSection(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("fail"))
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, 400, w.Code)
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
	params := make(url.Values)
	params["user_id"] = []string{paramId}
	params["ttl"] = []string{paramTtl}
	req, err := http.NewRequest(http.MethodPost, "fake_url"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
