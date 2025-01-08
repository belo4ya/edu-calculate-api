package service

import (
	"calculate-api/internal/httputil"
	"calculate-api/pkg/calc"
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Service) CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req calculateRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteJSONWithCode(w, errUnexpected, http.StatusInternalServerError)
		return
	}

	res, err := calc.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, calc.ErrInvalidExpr) {
			httputil.WriteJSONWithCode(w, errInvalidExpr, http.StatusUnprocessableEntity)
			return
		}
		httputil.WriteJSONWithCode(w, errUnexpected, http.StatusInternalServerError)
		return
	}

	httputil.WriteJSONWithCode(w, calculateResponse{Result: res}, http.StatusOK)
}

type calculateRequest struct {
	Expression string `json:"expression"`
}

type calculateResponse struct {
	Result float64 `json:"result"`
}

type calculateErrorResponse struct {
	Error string `json:"error"`
}

var (
	errInvalidExpr = calculateErrorResponse{Error: "Expression is not valid"}
	errUnexpected  = calculateErrorResponse{Error: "Internal server error"}
)
