package handlers_test

import (
	"Segments/pkg/handlers"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_NewSection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		sectionName string
		percentage  string
		prepare     func(m *Mockrepo)
		expected    func(t assert.TestingT, w *httptest.ResponseRecorder)
	}{
		{
			name:        "Success",
			sectionName: "real",
			percentage:  "100",
			prepare: func(m *Mockrepo) {
				m.EXPECT().CreateSection(gomock.Any(), gomock.Any()).Return(nil)
			},
			expected: func(t assert.TestingT, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
				//assert.Equal(t, gomock.Any(), w.Result())
			},
		},
		{
			name:        "Failed",
			sectionName: "real",
			percentage:  "100",
			prepare: func(m *Mockrepo) {
				m.EXPECT().CreateSection(gomock.Any(), gomock.Any()).Return(errors.New("DB failed"))
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

			r, err := createNewRequest(tc.sectionName, tc.percentage)
			assert.NoError(t, err)

			w := httptest.NewRecorder()

			handler.NewSection(w, r)
			tc.expected(t, w)
		})
	}
}

func createNewRequest(section string, percentage string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:8080/new/%s/%s", section, percentage), nil)
	if err != nil {
		return nil, err
	}
	req = mux.SetURLVars(req, map[string]string{"section": section, "percentage": percentage})
	return req, nil
}
