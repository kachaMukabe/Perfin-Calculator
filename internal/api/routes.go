package api

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

func (h *APIHandler) AddRoutesToGorillaMux(router *mux.Router) {
	for _, route := range h.GetRoutes() {
		router.
			Name(route.Name).
			Path(route.Path).
			Methods(route.Method).
			Handler(route.Handler)
	}
}

func (h *APIHandler) GetRoutes() Routes {
	return Routes{
		{
			"POSTCompoundInterest",
			"/compound-interest",
			"POST",
			h.HandlePOSTCompoundInterest,
		},{
			"POSTFinancialIndependenceNumber",
			"/financial-independence-number",
			"POST",
			h.HandlePOSTFinancialIndependenceNumber,
		},{
			"POSTFixedDepositMaturity",
			"/fixed-deposit-maturity",
			"POST",
			h.HandlePOSTFixedDepositMaturity,
		},{
			"POSTFutureValue",
			"/future-value",
			"POST",
			h.HandlePOSTFutureValue,
		},{
			"POSTInflationAdjustedReturn",
			"/inflation-adjusted-return",
			"POST",
			h.HandlePOSTInflationAdjustedReturn,
		},{
			"POSTInvestmentRoi",
			"/investment-roi",
			"POST",
			h.HandlePOSTInvestmentRoi,
		},{
			"POSTLoanPayment",
			"/loan-payment",
			"POST",
			h.HandlePOSTLoanPayment,
		},{
			"POSTSalaryGrowthRate",
			"/salary-growth-rate",
			"POST",
			h.HandlePOSTSalaryGrowthRate,
		},{
			"POSTSimpleInterest",
			"/simple-interest",
			"POST",
			h.HandlePOSTSimpleInterest,
		},
	}
}

