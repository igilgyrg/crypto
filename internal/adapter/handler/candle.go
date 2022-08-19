package handler

import (
	"github.com/igilgyrg/crypto/internal/adapter"
	"github.com/igilgyrg/crypto/internal/domain/candle"
	exception "github.com/igilgyrg/crypto/internal/error"
	"github.com/igilgyrg/crypto/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const (
	candlesURL = "/candles"
	candleURL  = "/candle?start=&end=:"
)

type handler struct {
	logger        *logging.Logger
	candleService candle.Service
}

func NewHandler(logger *logging.Logger, candleService candle.Service) adapter.Handler {
	return &handler{
		logger:        logger,
		candleService: candleService,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, candlesURL, exception.MiddleWare(h.GetByDatetime))
	router.HandlerFunc(http.MethodGet, candleURL, exception.MiddleWare(h.GetBetweenDatetime))
	router.HandlerFunc(http.MethodPost, candlesURL, exception.MiddleWare(h.Create))
}

func (h *handler) GetByDatetime(w http.ResponseWriter, request *http.Request) (interface{}, error) {
	return h.candleService.Store(request.Context(), nil)
}

func (h *handler) GetBetweenDatetime(w http.ResponseWriter, request *http.Request) (interface{}, error) {
	return h.candleService.GetByDatetime(request.Context(), time.Now())
}

func (h *handler) Create(w http.ResponseWriter, request *http.Request) (interface{}, error) {
	return h.candleService.GetBetweenDatetime(request.Context(), time.Now(), time.Now().Add(30))
}
