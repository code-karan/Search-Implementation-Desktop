package schema

type ShippingAddress struct {
	Name string `json:"name"`
}

type Result struct {
	ClientName      string           `json:"client_name"`
	Code            int64            `json:"code"`
	OrderType       string           `json:"order_type"`
	Status          string           `json:"status"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type LocalSearchResponse struct {
	Success bool     `json:"success"`
	Results []Result `json:"payload"`
}

// Example
// results := &schema.LocalSearchResponse{
// 	Results: []schema.Result{
// 		{
// 			ClientName: "Apple",
// 			Code:       1414,
// 			OrderType:  "non-priority",
// 			Status:     "shipped",
// 			ShippingAddress: &schema.ShippingAddress{
// 				Name: "Karan S.",
// 			},
// 		},
// 		{
// 			ClientName: "Apple",
// 			Code:       1616,
// 			OrderType:  "non-priority",
// 			Status:     "open",
// 			ShippingAddress: &schema.ShippingAddress{
// 				Name: "Himanshu S.",
// 			},
// 		},
// 	},
// }
