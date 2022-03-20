package rest

import (
	"bytes"
	"fmt"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/danilkaz/chartographer/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"golang.org/x/image/bmp"
	"image"
	"image/png"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func setup() (*Handler, string) {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "test_images")
	_ = os.MkdirAll(path, os.ModePerm)
	return NewHandler(service.NewService(repository.NewRepository(path))), path
}

func teardown(path string) {
	_ = os.RemoveAll(path)
}

func TestCRUD_CreateNewCharta(t *testing.T) {
	handler, path := setup()
	defer teardown(path)

	testCases := []struct {
		name       string
		statusCode int
		vars       map[string]string
	}{
		{
			name:       "Common test",
			statusCode: http.StatusCreated,
			vars: map[string]string{
				"width":  "123",
				"height": "456",
			},
		},
		{
			name:       "Incorrect parameters",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"width":  "123a",
				"height": "456",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST",
				fmt.Sprintf("/chartas/?width=%s&height=%s", tc.vars["width"], tc.vars["height"]), nil)
			request = mux.SetURLVars(request, tc.vars)
			h := http.HandlerFunc(handler.CreateNewCharta)
			h.ServeHTTP(recorder, request)
			assert.Equal(t, recorder.Code, tc.statusCode)
		})
	}
}

func TestCRUD_SaveRestoredFragmentOfCharta(t *testing.T) {
	handler, path := setup()
	defer teardown(path)

	id := createCharta(handler)

	type EncodeFunction func(w io.Writer, img image.Image) error

	testCases := []struct {
		name       string
		statusCode int
		vars       map[string]string
		function   EncodeFunction
	}{
		{
			name:       "Common test",
			statusCode: http.StatusOK,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "0",
				"y":      "0",
			},
			function: bmp.Encode,
		},
		{
			name:       "Incorrect parameters",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
			function: bmp.Encode,
		},
		{
			name:       "Parameter missing",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
			function: bmp.Encode,
		},
		{
			name:       "Id is not a UUID",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     "123",
				"width":  "200",
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
			function: bmp.Encode,
		},
		{
			name:       "Update out of bounds",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "-500",
				"y":      "-600",
			},
			function: bmp.Encode,
		},
		{
			name:       "Id does not exist",
			statusCode: http.StatusNotFound,
			vars: map[string]string{
				"id":     uuid.New().String(),
				"width":  "200",
				"height": "300",
				"x":      "0",
				"y":      "0",
			},
			function: bmp.Encode,
		},
		{
			name:       "Image is not decoded",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "0",
				"y":      "0",
			},
			function: png.Encode,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fragment := image.NewRGBA(image.Rect(0, 0, 200, 300))
			buffer := new(bytes.Buffer)
			_ = tc.function(buffer, fragment)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST",
				fmt.Sprintf("/chartas/%s/?width=%s&height=%s&x=%s&y=%s",
					tc.vars["id"], tc.vars["width"], tc.vars["height"], tc.vars["x"], tc.vars["y"]), buffer)
			request = mux.SetURLVars(request, tc.vars)
			h := http.HandlerFunc(handler.SaveRestoredFragmentOfCharta)
			h.ServeHTTP(recorder, request)
			assert.Equal(t, recorder.Code, tc.statusCode)
		})
	}
}

func TestCRUD_GetPartOfCharta(t *testing.T) {
	handler, path := setup()
	defer teardown(path)

	createRecorder := httptest.NewRecorder()
	createRequest := httptest.NewRequest("POST",
		fmt.Sprintf("/chartas/?width=%s&height=%s", "1000", "1000"), nil)
	createRequest = mux.SetURLVars(createRequest, map[string]string{"width": "1000", "height": "1000"})
	handler.CreateNewCharta(createRecorder, createRequest)
	id := createRecorder.Body.String()

	testCases := []struct {
		name       string
		statusCode int
		vars       map[string]string
	}{
		{
			name:       "Common test",
			statusCode: http.StatusOK,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "0",
				"y":      "0",
			},
		},
		{
			name:       "Incorrect parameters",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
		},
		{
			name:       "Parameter missing",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
		},
		{
			name:       "Id is not a UUID",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     "123",
				"width":  "200",
				"height": "300",
				"x":      "xyz",
				"y":      "0",
			},
		},
		{
			name:       "Update out of bounds",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id":     id,
				"width":  "200",
				"height": "300",
				"x":      "-500",
				"y":      "-600",
			},
		},
		{
			name:       "Id does not exist",
			statusCode: http.StatusNotFound,
			vars: map[string]string{
				"id":     uuid.New().String(),
				"width":  "200",
				"height": "300",
				"x":      "0",
				"y":      "0",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET",
				fmt.Sprintf("/chartas/%s/?width=%s&height=%s&x=%s&y=%s",
					tc.vars["id"], tc.vars["width"], tc.vars["height"], tc.vars["x"], tc.vars["y"]), nil)
			request = mux.SetURLVars(request, tc.vars)
			h := http.HandlerFunc(handler.GetPartOfCharta)
			h.ServeHTTP(recorder, request)
			assert.Equal(t, recorder.Code, tc.statusCode)
		})
	}
}

func TestCRUD_DeleteCharta(t *testing.T) {
	handler, path := setup()
	defer teardown(path)

	testCases := []struct {
		name       string
		statusCode int
		vars       map[string]string
	}{
		{
			name:       "Common test",
			statusCode: http.StatusOK,
			vars:       map[string]string{},
		},
		{
			name:       "Id is not a UUID",
			statusCode: http.StatusBadRequest,
			vars: map[string]string{
				"id": "123",
			},
		},
		{
			name:       "Id does not exist",
			statusCode: http.StatusNotFound,
			vars: map[string]string{
				"id": uuid.New().String(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id := createCharta(handler)
			if _, exists := tc.vars["id"]; !exists {
				tc.vars["id"] = id
			}
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE",
				fmt.Sprintf("/chartas/%s/", tc.vars["id"]), nil)
			request = mux.SetURLVars(request, tc.vars)
			h := http.HandlerFunc(handler.DeleteCharta)
			h.ServeHTTP(recorder, request)
			assert.Equal(t, recorder.Code, tc.statusCode)
		})
	}
}

func createCharta(handler *Handler) string {
	createRecorder := httptest.NewRecorder()
	createRequest := httptest.NewRequest("POST",
		fmt.Sprintf("/chartas/?width=%s&height=%s", "1000", "1000"), nil)
	createRequest = mux.SetURLVars(createRequest, map[string]string{"width": "1000", "height": "1000"})
	handler.CreateNewCharta(createRecorder, createRequest)
	return createRecorder.Body.String()
}
