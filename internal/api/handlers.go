package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CompoundinterestBody struct {
  Principal float32 `json:"principal,omitempty"`
  AnnualRate float32 `json:"annual_rate,omitempty"`
  Time float32 `json:"time,omitempty"`
  Compound string `json:"compound,omitempty"`
}

// HandlePOSTCompoundInterest handles parsing input to pass to the POSTCompoundInterest operation and sends responses back to the client
func (h *APIHandler) HandlePOSTCompoundInterest(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := CompoundinterestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTCompoundInterest(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTCompoundInterest returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTCompoundInterest was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTFinancialIndependenceNumber handles parsing input to pass to the POSTFinancialIndependenceNumber operation and sends responses back to the client
func (h *APIHandler) HandlePOSTFinancialIndependenceNumber(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTFinancialIndependenceNumber(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFinancialIndependenceNumber returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFinancialIndependenceNumber was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTFixedDepositMaturity handles parsing input to pass to the POSTFixedDepositMaturity operation and sends responses back to the client
func (h *APIHandler) HandlePOSTFixedDepositMaturity(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTFixedDepositMaturity(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFixedDepositMaturity returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFixedDepositMaturity was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTFutureValue handles parsing input to pass to the POSTFutureValue operation and sends responses back to the client
func (h *APIHandler) HandlePOSTFutureValue(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTFutureValue(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFutureValue returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTFutureValue was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTInflationAdjustedReturn handles parsing input to pass to the POSTInflationAdjustedReturn operation and sends responses back to the client
func (h *APIHandler) HandlePOSTInflationAdjustedReturn(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTInflationAdjustedReturn(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTInflationAdjustedReturn returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTInflationAdjustedReturn was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTInvestmentRoi handles parsing input to pass to the POSTInvestmentRoi operation and sends responses back to the client
func (h *APIHandler) HandlePOSTInvestmentRoi(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTInvestmentRoi(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTInvestmentRoi returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTInvestmentRoi was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTLoanPayment handles parsing input to pass to the POSTLoanPayment operation and sends responses back to the client
func (h *APIHandler) HandlePOSTLoanPayment(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTLoanPayment(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTLoanPayment returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTLoanPayment was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTSalaryGrowthRate handles parsing input to pass to the POSTSalaryGrowthRate operation and sends responses back to the client
func (h *APIHandler) HandlePOSTSalaryGrowthRate(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.POSTSalaryGrowthRate(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTSalaryGrowthRate returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTSalaryGrowthRate was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTSimpleInterest handles parsing input to pass to the POSTSimpleInterest operation and sends responses back to the client
func (h *APIHandler) HandlePOSTSimpleInterest(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := {}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTSimpleInterest(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTSimpleInterest returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTSimpleInterest was unable to send it's response, err: %s", err)
	}
}

