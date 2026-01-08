package main

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		service string
		params  map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "ecommerce service",
			args: args{
				service: "ecommerce",
				params: map[string]interface{}{
					"amount":   5.0,
					"customer": "vip",
				},
			},
			want: 4.0,
		},
		{
			name: "banking service",
			args: args{
				service: "banking",
				params: map[string]interface{}{
					"balance": 50.0,
					"account": "savings",
				},
			},
			want: 1.0,
		},
		{
			name: "healthcare service",
			args: args{
				service: "healthcare",
				params: map[string]interface{}{
					"cost":      20.0,
					"insured":   true,
					"emergency": true,
				},
			},
			want: 9.0,
		},
		{
			name: "logistics service",
			args: args{
				service: "logistics",
				params: map[string]interface{}{
					"weight":        20.0,
					"express":       true,
					"international": false,
				},
			},
			want: 30.0,
		},
		{
			name: "unknown service",
			args: args{
				service: "",
				params: map[string]interface{}{
					"weight":        0.0,
					"express":       false,
					"international": false,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.service, tt.args.params); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate_NilParamsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Test passed, panic was caught!")
		}
	}()

	Calculate("", nil)
	t.Errorf("Test failed, panic was expected")
}

func TestCalculateInterest(t *testing.T) {
	type args struct {
		balance     float64
		accountType string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "negative balance",
			args: args{
				balance:     -10,
				accountType: "",
			},
			want: 0,
		},
		{
			name: "savings account",
			args: args{
				balance:     10,
				accountType: "savings",
			},
			want: 0.2,
		},
		{
			name: "fixed account",
			args: args{
				balance:     10,
				accountType: "fixed",
			},
			want: 0.5,
		},
		{
			name: "checking account",
			args: args{
				balance:     10,
				accountType: "checking",
			},
			want: 0.05,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateInterest(tt.args.balance, tt.args.accountType); got != tt.want {
				t.Errorf("CalculateInterest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate_UnknownAccountPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Test passed, panic was caught!")
		}
	}()

	CalculateInterest(10.0, "")
	t.Errorf("Test failed, panic was expected")
}

func TestCalculateOrderTotal(t *testing.T) {
	type args struct {
		amount       float64
		customerType string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "negative amount",
			args: args{
				amount:       -100,
				customerType: "",
			},
			want: -1,
		},
		{
			name: "regular customer",
			args: args{
				amount:       100,
				customerType: "regular",
			},
			want: 100,
		},
		{
			name: "premium customer",
			args: args{
				amount:       100,
				customerType: "premium",
			},
			want: 90,
		},
		{
			name: "vip customer",
			args: args{
				amount:       100,
				customerType: "vip",
			},
			want: 80,
		},
		{
			name: "other customers",
			args: args{
				amount:       100,
				customerType: "",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateOrderTotal(tt.args.amount, tt.args.customerType); got != tt.want {
				t.Errorf("CalculateOrderTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateShipping(t *testing.T) {
	type args struct {
		weight        float64
		express       bool
		international bool
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "negative weight",
			args: args{
				weight:        -10,
				express:       false,
				international: false,
			},
			want: -1,
		},
		{
			name: "weight 0-10, no express, no international",
			args: args{
				weight:        5,
				express:       false,
				international: false,
			},
			want: 5,
		},
		{
			name: "weight > 10, no express, international",
			args: args{
				weight:        12,
				express:       false,
				international: false,
			},
			want: 15,
		},
		{
			name: "weight > 10 with express and international",
			args: args{
				weight:        12,
				express:       true,
				international: true,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateShipping(tt.args.weight, tt.args.express, tt.args.international); got != tt.want {
				t.Errorf("CalculateShipping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateTreatmentCost(t *testing.T) {
	type args struct {
		baseCost  float64
		insured   bool
		emergency bool
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "negative baseCost",
			args: args{
				baseCost:  -10,
				insured:   false,
				emergency: false,
			},
			want: 0,
		},
		{
			name: "not insured",
			args: args{
				baseCost:  10,
				insured:   false,
				emergency: false,
			},
			want: 10,
		},
		{
			name: "insured but not emergency",
			args: args{
				baseCost:  10,
				insured:   true,
				emergency: false,
			},
			want: 3,
		},
		{
			name: "insured and emergency",
			args: args{
				baseCost:  10,
				insured:   true,
				emergency: true,
			},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTreatmentCost(tt.args.baseCost, tt.args.insured, tt.args.emergency); got != tt.want {
				t.Errorf("CalculateTreatmentCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
