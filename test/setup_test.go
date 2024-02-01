package test

import (
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/controller"
	"jpnovelmtlgo/internal/repository"
	"jpnovelmtlgo/internal/service"
	"jpnovelmtlgo/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	// Mock
	MockTranslateRepository *mocks.TranslateRepository
	MockKakuyomuService     *mocks.KakuyomuService
	MockSyosetuService      *mocks.SyosetuService
	MockConfig              *mocks.Config
	// Repository
	translateRepository repository.TranslateRepository
	// Service
	syosetuService  service.SyosetuService
	kakuyomuService service.KakuyomuService
	// Controller
	kakuyomuController    controller.KakuyomuController
	syosetuController     controller.SyosetuController
	healthcheckController controller.HealthcheckController
	// Other
	suite.Suite
	app *fiber.App
}

func TestUnitTestSuite(t *testing.T) {
	t.Setenv("TRANSLATE_URL", "http://127.0.0.1:5001/translate")
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupSuite() {
	// Mock
	MockKakuyomuService := mocks.KakuyomuService{}
	MockSyosetuService := mocks.SyosetuService{}
	MockTranslateRepository := mocks.TranslateRepository{}
	MockConfig := mocks.Config{}
	// Repository
	translateRepository := repository.NewTranslateRepository(&MockConfig)
	// Service
	syosetuService := service.NewSyosetuService(&MockTranslateRepository)
	kakuyomuService := service.NewKakuyomuService(&MockTranslateRepository)
	// Controller
	kakuyomuController := controller.NewKakuyomuController(&MockKakuyomuService)
	syosetuController := controller.NewSyosetuController(&MockSyosetuService)
	healthcheckController := controller.NewHealthcheckController()

	// Mock
	uts.MockTranslateRepository = &MockTranslateRepository
	uts.MockKakuyomuService = &MockKakuyomuService
	uts.MockSyosetuService = &MockSyosetuService
	uts.MockConfig = &MockConfig
	// Repository
	uts.translateRepository = translateRepository
	// Service
	uts.syosetuService = syosetuService
	uts.kakuyomuService = kakuyomuService
	// Controller
	uts.app = fiber.New(config.NewFiberConfig())
	uts.kakuyomuController = kakuyomuController
	uts.syosetuController = syosetuController
	uts.healthcheckController = healthcheckController
	uts.kakuyomuController.Route(uts.app)
	uts.syosetuController.Route(uts.app)
	uts.healthcheckController.Route(uts.app)
}

func (uts *UnitTestSuite) TearDownSuite() {
	uts.MockTranslateRepository = nil
	uts.MockKakuyomuService = nil
	uts.kakuyomuService = nil
}
