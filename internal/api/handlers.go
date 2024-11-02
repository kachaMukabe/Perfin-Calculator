package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CompoundinterestBody struct {
	Principal  float32 `json:"principal,omitempty"`
	AnnualRate float32 `json:"annual_rate,omitempty"`
	Time       float32 `json:"time,omitempty"`
	Compound   string  `json:"compound,omitempty"`
}

func (c *CompoundinterestBody) Validate() error {
	if c.Principal <= 0 {
		return fmt.Errorf("principal must be greater than zero")
	}
	if c.AnnualRate <= 0 {
		return fmt.Errorf("annual rate must be greater than zero")
	}
	if c.Time <= 0 {
		return fmt.Errorf("time must be greater than zero")
	}
	if c.Compound == "" {
		return fmt.Errorf("compound frequency is required")
	}
	return nil
}

type FinancialIndependenceNumberBody struct {
	AnnualExpenses float32 `json:"annual_expenses,omitempty"`
	WithdrawalRate float32 `json:"withdrawal_rate,omitempty"`
}

type FixedDepositMaturityBody struct {
	Principal            float32 `json:"principal,omitempty"`
	Rate                 float32 `json:"rate,omitempty"`
	Time                 float32 `json:"time,omitempty"`
	CompoundingFrequency string  `json:"compounding_frequency,omitempty"`
}

type FutureValueBody struct {
	Principal float32 `json:"principal,omitempty"`
	Rate      float32 `json:"rate,omitempty"`
	Time      float32 `json:"time,omitempty"`
	Compound  string  `json:"compound,omitempty"`
}

type InflationAdjustedReturnBody struct {
	NominalRate   float32 `json:"nominal_rate,omitempty"`
	InflationRate float32 `json:"inflation_rate,omitempty"`
}

type InvestmentRoiBody struct {
	InitialInvestment float32 `json:"initial_investment,omitempty"`
	Earning           float32 `json:"earnings,omitempty"`
}

type LoanPaymentBody struct {
	Principal   float32 `json:"principal,omitempty"`
	Rate        float32 `json:"rate,omitempty"`
	NumPayments int32   `json:"num_payments,omitempty"`
}

type SalaryGrowthRateBody struct {
	Salary float32 `json:"salary,omitempty"`
	Year   float64 `json:"year,omitempty"`
}

type SimpleInterestBody struct {
	Principal float32 `json:"principal,omitempty"`
	Rate      float32 `json:"rate,omitempty"`
	Time      float32 `json:"time,omitempty"`
	Period    string  `json:"period,omitempty"`
}

func (f *FinancialIndependenceNumberBody) Validate() error {
	if f.AnnualExpenses <= 0 {
		return fmt.Errorf("annual expenses must be greater than zero")
	}
	if f.WithdrawalRate <= 0 || f.WithdrawalRate > 1 {
		return fmt.Errorf("withdrawal rate must be a positive value less than or equal to 1")
	}
	return nil
}

func (f *FixedDepositMaturityBody) Validate() error {
	if f.Principal <= 0 {
		return fmt.Errorf("principal must be greater than zero")
	}
	if f.Rate <= 0 {
		return fmt.Errorf("rate must be greater than zero")
	}
	if f.Time <= 0 {
		return fmt.Errorf("time must be greater than zero")
	}
	if f.CompoundingFrequency == "" {
		return fmt.Errorf("compounding frequency is required")
	}
	return nil
}

func (f *FutureValueBody) Validate() error {
	if f.Principal <= 0 {
		return fmt.Errorf("principal must be greater than zero")
	}
	if f.Rate <= 0 {
		return fmt.Errorf("rate must be greater than zero")
	}
	if f.Time <= 0 {
		return fmt.Errorf("time must be greater than zero")
	}
	if f.Compound == "" {
		return fmt.Errorf("compound frequency is required")
	}
	return nil
}

func (i *InflationAdjustedReturnBody) Validate() error {
	if i.NominalRate <= 0 {
		return fmt.Errorf("nominal rate must be greater than zero")
	}
	if i.InflationRate < 0 {
		return fmt.Errorf("inflation rate cannot be negative")
	}
	return nil
}

func (i *InvestmentRoiBody) Validate() error {
	if i.InitialInvestment <= 0 {
		return fmt.Errorf("initial investment must be greater than zero")
	}
	if i.Earning < 0 {
		return fmt.Errorf("earnings cannot be negative")
	}
	return nil
}

func (l *LoanPaymentBody) Validate() error {
	if l.Principal <= 0 {
		return fmt.Errorf("principal must be greater than zero")
	}
	if l.Rate <= 0 {
		return fmt.Errorf("rate must be greater than zero")
	}
	if l.NumPayments <= 0 {
		return fmt.Errorf("number of payments must be greater than zero")
	}
	return nil
}

func (s *SalaryGrowthRateBody) Validate() error {
	if s.Salary <= 0 {
		return fmt.Errorf("salary must be greater than zero")
	}
	if s.Year <= 0 {
		return fmt.Errorf("year must be greater than zero")
	}
	return nil
}

func (s *SimpleInterestBody) Validate() error {
	if s.Principal <= 0 {
		return fmt.Errorf("principal must be greater than zero")
	}
	if s.Rate <= 0 {
		return fmt.Errorf("rate must be greater than zero")
	}
	if s.Time <= 0 {
		return fmt.Errorf("time must be greater than zero")
	}
	if s.Period == "" {
		return fmt.Errorf("period is required")
	}
	return nil
}

// HandlePOSTCompoundInterest handles parsing input to pass to the POSTCompoundInterest operation and sends responses back to the client
func (h *APIHandler) HandlePOSTCompoundInterest(w http.ResponseWriter, r *http.Request) {
	var err error
	var reqBody CompoundinterestBody
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
	var reqBody FinancialIndependenceNumberBody
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
	var reqBody FixedDepositMaturityBody
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

// HandlePOSTInflationAdjustedReturn handles parsing input to pass to the POSTInflationAdjustedReturn operation and sends responses back to the client
func (h *APIHandler) HandlePOSTInflationAdjustedReturn(w http.ResponseWriter, r *http.Request) {
	var err error
	var reqBody InflationAdjustedReturnBody
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
	var reqBody InvestmentRoiBody
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
	var reqBody LoanPaymentBody
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
	var reqBody []SalaryGrowthRateBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully ''", w)
		return
	}
	response, err := h.POSTSalaryGrowthRate(r.Context(), reqBody)
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
	var reqBody SimpleInterestBody
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
