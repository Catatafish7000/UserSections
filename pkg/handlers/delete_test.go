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

func TestHandler_Delete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		sectionName string
		prepare     func(m *Mockrepo)
		expected    func(t assert.TestingT, w *httptest.ResponseRecorder)
	}{
		{
			name:        "Success",
			sectionName: "real",
			prepare: func(m *Mockrepo) {
				m.EXPECT().DeleteSection(gomock.Any()).Return(nil)
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
				//assert.Equal(t, gomock.Any(), w.Result())
			},
		},
		{
			name:        "Failed",
			sectionName: "real",
			prepare: func(m *Mockrepo) {
				m.EXPECT().DeleteSection(gomock.Any()).Return(errors.New("DB failed"))
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
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

			r, err := createDelRequest(tc.sectionName)
			assert.NoError(t, err)

			w := httptest.NewRecorder()

			handler.Delete(w, r)
			tc.expected(t, w)
		})
	}
}

func createDelRequest(param string) (*http.Request, error) {
	params := make(url.Values)
	params["section"] = []string{param}
	req, err := http.NewRequest(http.MethodPost, "fake_url"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
