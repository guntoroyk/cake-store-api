package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/cake-store-api/entity"
	"github.com/guntoroyk/cake-store-api/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	successGetCakesJSON   = "{\"code\":200,\"data\":[{\"id\":1,\"title\":\"Cake 1\",\"description\":\"Desc 1\",\"rating\":1,\"image\":\"https://dummyimage.com/600x400/000/fff\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}]}\n"
	successGetCakeJSON    = "{\"code\":200,\"data\":{\"id\":1,\"title\":\"Cake 1\",\"description\":\"Desc 1\",\"rating\":1,\"image\":\"https://dummyimage.com/600x400/000/fff\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
	successUpdateCakeJSON = "{\"code\":200,\"data\":{\"id\":1,\"title\":\"Cake 1\",\"description\":\"Desc 1\",\"rating\":1,\"image\":\"https://dummyimage.com/600x400/000/fff\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
	successCreateCakeJSON = "{\"code\":201,\"data\":{\"id\":1,\"title\":\"Cake 1\",\"description\":\"Desc 1\",\"rating\":1,\"image\":\"https://dummyimage.com/600x400/000/fff\",\"created_at\":\"2021-01-01 00:00:00\",\"updated_at\":\"2021-01-01 00:00:00\"}}\n"
)

func Test_handler_GetCakes(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/cakes")

	t.Run("success get cakes", func(t *testing.T) {
		mockTeamUC := mocks.NewMockCakeUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetCakes().Return([]*entity.Cake{
			{
				ID:          1,
				Title:       "Cake 1",
				Description: "Desc 1",
				Rating:      1,
				Image:       "https://dummyimage.com/600x400/000/fff",
				CreatedAt:   "2021-01-01 00:00:00",
				UpdatedAt:   "2021-01-01 00:00:00",
			},
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetCakes(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetCakesJSON, rec.Body.String())
		}
	})
}

func Test_handler_GetCake(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/cakes/:id")

	t.Run("success get cake", func(t *testing.T) {
		mockTeamUC := mocks.NewMockCakeUsecaseItf(ctrl)
		mockTeamUC.EXPECT().GetCake(gomock.Any()).Return(&entity.Cake{

			ID:          1,
			Title:       "Cake 1",
			Description: "Desc 1",
			Rating:      1,
			Image:       "https://dummyimage.com/600x400/000/fff",
			CreatedAt:   "2021-01-01 00:00:00",
			UpdatedAt:   "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.GetCake(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successGetCakeJSON, rec.Body.String())
		}
	})
}

func Test_handler_UpdateCake(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/cakes/:id")

	t.Run("success update cake", func(t *testing.T) {
		mockTeamUC := mocks.NewMockCakeUsecaseItf(ctrl)
		mockTeamUC.EXPECT().UpdateCake(gomock.Any()).Return(&entity.Cake{

			ID:          1,
			Title:       "Cake 1",
			Description: "Desc 1",
			Rating:      1,
			Image:       "https://dummyimage.com/600x400/000/fff",
			CreatedAt:   "2021-01-01 00:00:00",
			UpdatedAt:   "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.UpdateCake(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, successUpdateCakeJSON, rec.Body.String())
		}
	})
}

func Test_handler_CreateCake(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/cakes")

	t.Run("success create cake", func(t *testing.T) {
		mockTeamUC := mocks.NewMockCakeUsecaseItf(ctrl)
		mockTeamUC.EXPECT().CreateCake(gomock.Any()).Return(&entity.Cake{

			ID:          1,
			Title:       "Cake 1",
			Description: "Desc 1",
			Rating:      1,
			Image:       "https://dummyimage.com/600x400/000/fff",
			CreatedAt:   "2021-01-01 00:00:00",
			UpdatedAt:   "2021-01-01 00:00:00",
		}, nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.CreateCake(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, successCreateCakeJSON, rec.Body.String())
		}
	})
}

func Test_handler_DeleteCake(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/cakes/:id")

	t.Run("success delete cake", func(t *testing.T) {
		mockTeamUC := mocks.NewMockCakeUsecaseItf(ctrl)
		mockTeamUC.EXPECT().DeleteCake(gomock.Any()).Return(nil)
		h := NewHandler(mockTeamUC)

		// Assertions
		if assert.NoError(t, h.DeleteCake(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})
}
