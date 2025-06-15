package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sigmatech-kredit-plus/internal/consumer/dto"
	"sigmatech-kredit-plus/internal/consumer/handler"
	"sigmatech-kredit-plus/internal/consumer/mocks"
	"sigmatech-kredit-plus/pkg"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateConsumer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pkg.InitValidator()

	testCases := map[string]struct {
		requestBody      any
		mockUsecaseErr   error
		setPhotoContext1 bool
		setPhotoContext2 bool
		expectedStatus   int
	}{
		"successfully create consumer": {
			requestBody: dto.CreateConsumer{
				NIK:        "1234567890",
				FullName:   "John Doe",
				LegalName:  "John D",
				BirthPlace: "Jakarta",
				BirthDate:  time.Now(),
				Salary:     10000000,
			},
			mockUsecaseErr:   nil,
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusCreated,
		},
		"usecase returns error": {
			requestBody: dto.CreateConsumer{
				NIK:        "999999999",
				FullName:   "Jane Doe",
				LegalName:  "Jane D",
				BirthPlace: "Bandung",
				BirthDate:  time.Now(),
				Salary:     9000000,
			},
			mockUsecaseErr:   errors.New("unknown error"),
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusInternalServerError,
		},
		"nik required": {
			requestBody: dto.CreateConsumer{
				FullName:   "Jane Doe",
				LegalName:  "Jane D",
				BirthPlace: "Bandung",
				BirthDate:  time.Now(),
				Salary:     9000000,
			},
			mockUsecaseErr:   errors.New("unknown error"),
			setPhotoContext1: true,
			setPhotoContext2: true,
			expectedStatus:   http.StatusBadRequest,
		},
		"missing photo context 1": {
			requestBody: dto.CreateConsumer{
				NIK:        "88888888",
				FullName:   "No Photo",
				LegalName:  "No Photo",
				BirthPlace: "Surabaya",
				BirthDate:  time.Now(),
				Salary:     1000000,
			},
			setPhotoContext1: false,
			setPhotoContext2: true,
			expectedStatus:   http.StatusInternalServerError,
		},
		"missing photo context 2": {
			requestBody: dto.CreateConsumer{
				NIK:        "88888888",
				FullName:   "No Photo",
				LegalName:  "No Photo",
				BirthPlace: "Surabaya",
				BirthDate:  time.Now(),
				Salary:     1000000,
			},
			setPhotoContext1: true,
			setPhotoContext2: false,
			expectedStatus:   http.StatusInternalServerError,
		},
		"invalid body": {
			requestBody:    "wrong-format",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockUsecase := new(mocks.ConsumerUsecaseMock)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest(http.MethodPost, "/consumer", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			if tc.setPhotoContext1 {
				ctx.Set("photo_ktp", "path/to/ktp.jpg")
			}
			if tc.setPhotoContext2 {
				ctx.Set("photo_selfie", "path/to/selfie.jpg")
			}

			mockUsecase.On("CreateConsumer", mock.Anything, mock.Anything).Return(tc.mockUsecaseErr).Once()

			consumerHandler := handler.NewConsumerHandler(mockUsecase)
			consumerHandler.CreateConsumer(ctx)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
