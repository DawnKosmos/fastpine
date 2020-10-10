package ftx

import "time"

type NewOrder struct {
	Market                  string  `json:"market"`
	Side                    string  `json:"side"`
	Price                   float64 `json:"price"`
	Type                    string  `json:"type"`
	Size                    float64 `json:"size"`
	ReduceOnly              bool    `json:"reduceOnly"`
	Ioc                     bool    `json:"ioc"`
	PostOnly                bool    `json:"postOnly"`
	ExternalReferralProgram string  `json:"externalReferralProgram"`
	// ClientID                string  `json:"clientId"`
}

type Order struct {
	CreatedAt     time.Time `json:"createdAt"`
	FilledSize    float64   `json:"filledSize"`
	Future        string    `json:"future"`
	ID            int64     `json:"id"`
	Market        string    `json:"market"`
	Price         float64   `json:"price"`
	AvgFillPrice  float64   `json:"avgFillPrice"`
	RemainingSize float64   `json:"remainingSize"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	ReduceOnly    bool      `json:"reduceOnly"`
	Ioc           bool      `json:"ioc"`
	PostOnly      bool      `json:"postOnly"`
	ClientID      string    `json:"clientId"`
}

type Fills struct {
	Success bool   `json:"success"`
	Result  []Fill `json:"result"`
}

type OpenOrders struct {
	Success bool    `json:"success"`
	Result  []Order `json:"result"`
}

type Trades struct {
	Success bool    `json:"success"`
	Result  []Trade `json:"result"`
}

type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

type Positions struct {
	Success bool `json:"success"`
	Result  []struct {
		Cost                         float64 `json:"cost"`
		EntryPrice                   float64 `json:"entryPrice"`
		EstimatedLiquidationPrice    float64 `json:"estimatedLiquidationPrice"`
		Future                       string  `json:"future"`
		InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
		LongOrderSize                float64 `json:"longOrderSize"`
		MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
		NetSize                      float64 `json:"netSize"`
		OpenSize                     float64 `json:"openSize"`
		RealizedPnl                  float64 `json:"realizedPnl"`
		ShortOrderSize               float64 `json:"shortOrderSize"`
		Side                         string  `json:"side"`
		Size                         float64 `json:"size"`
		UnrealizedPnl                float64 `json:"unrealizedPnl"`
	} `json:"result"`
}
