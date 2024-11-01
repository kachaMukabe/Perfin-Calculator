package api

import (
	"context"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// Calculate compound interest
func (h *APIHandler) POSTCompoundInterest(ctx context.Context, reqBody CompoundinterestBody) (Response, error) {
	// TODO: implement the POSTCompoundInterest function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil
	var n float32
	switch reqBody.Compound {
	case "annually":
		n = 1
	case "semiannually":
		n = 2
	case "quarterly":
		n = 4
	case "monthly":
		n = 12
	default:
		return NewResponse(http.StatusBadRequest, ErrorMsg{"invalid compounding frequency"}, "application/json", nil), nil
	}

	// Compound Interest formula: A = P * (1 + r/n)^(nt)
	// where A = final amount, P = principal, r = annual rate, t = time, and n = frequency
	interest_rate := reqBody.AnnualRate / 100
	finalAmount := reqBody.Principal * float32(math.Pow(float64(1+interest_rate/n), float64(n*reqBody.Time)))

	// Prepare response body
	responseBody := map[string]float32{
		"total": finalAmount,
	}

	// Create and return the response
	return NewResponse(http.StatusOK, responseBody, "application/json", nil), nil
}

// Calculate Financial Independence (FI) Number
func (h *APIHandler) POSTFinancialIndependenceNumber(ctx context.Context, reqBody FinancialIndependenceNumberBody) (Response, error) {
	// TODO: implement the POSTFinancialIndependenceNumber function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	financialIndependenceNumber := reqBody.AnnualExpenses / reqBody.WithdrawalRate

	// Round FIN to two decimal places
	financialIndependenceNumber = float32(math.Round(float64(financialIndependenceNumber)*100) / 100)

	// Prepare response body
	responseBody := map[string]float32{
		"fi_number": financialIndependenceNumber,
	}

	// Create and return the response
	return NewResponse(http.StatusOK, responseBody, "application/json", nil), nil
}

// Calculate fixed deposit maturity value
func (h *APIHandler) POSTFixedDepositMaturity(ctx context.Context, reqBody FixedDepositMaturityBody) (Response, error) {
	// TODO: implement the POSTFixedDepositMaturity function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTFixedDepositMaturity operation has not been implemented yet"}, "application/json", nil), nil
}

// Calculate future value
func (h *APIHandler) POSTFutureValue(ctx context.Context, reqBody FutureValueBody) (Response, error) {
	// TODO: implement the POSTFutureValue function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTFutureValue operation has not been implemented yet"}, "application/json", nil), nil
}

// Calculate inflation-adjusted return
func (h *APIHandler) POSTInflationAdjustedReturn(ctx context.Context, reqBody InflationAdjustedReturnBody) (Response, error) {
	// TODO: implement the POSTInflationAdjustedReturn function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTInflationAdjustedReturn operation has not been implemented yet"}, "application/json", nil), nil
}

// Calculate return on investment (ROI)
func (h *APIHandler) POSTInvestmentRoi(ctx context.Context, reqBody InvestmentRoiBody) (Response, error) {
	// TODO: implement the POSTInvestmentRoi function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTInvestmentRoi operation has not been implemented yet"}, "application/json", nil), nil
}

// Calculate loan payment
func (h *APIHandler) POSTLoanPayment(ctx context.Context, reqBody LoanPaymentBody) (Response, error) {
	// TODO: implement the POSTLoanPayment function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTLoanPayment operation has not been implemented yet"}, "application/json", nil), nil
}

// Allows you to compare your salary growth rate to inflation
// Calculate your salary growth rate over a period of years
func (h *APIHandler) POSTSalaryGrowthRate(ctx context.Context) (Response, error) {
	// TODO: implement the POSTSalaryGrowthRate function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTSalaryGrowthRate operation has not been implemented yet"}, "application/json", nil), nil
}

// Calculate simple interest
func (h *APIHandler) POSTSimpleInterest(ctx context.Context, reqBody SimpleInterestBody) (Response, error) {
	// TODO: implement the POSTSimpleInterest function to return the following responses

	// return NewResponse(200, {}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(422, {}, "", responseHeaders), nil

	simpleInterest := reqBody.Principal * (reqBody.Rate / 100) * reqBody.Time
	total := reqBody.Principal + simpleInterest

	// Prepare response body
	responseBody := map[string]float32{
		"total": total,
	}

	return NewResponse(http.StatusOK, responseBody, "application/json", nil), nil
}
