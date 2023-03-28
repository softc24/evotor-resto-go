package evotorrestogo

import (
	"reflect"
	"testing"
)

func TestMoney_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       Money
		want    []byte
		wantErr bool
	}{
		{
			name:    "Zero",
			m:       0,
			want:    []byte("0.00"),
			wantErr: false,
		},
		{
			name:    "Int",
			m:       12300,
			want:    []byte("123.00"),
			wantErr: false,
		},
		{
			name:    "Float",
			m:       12345,
			want:    []byte("123.45"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Money.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Money.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Money
		wantErr bool
	}{
		{
			name: "Zero",
			args: args{
				data: []byte("0"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Integer",
			args: args{
				data: []byte("123"),
			},
			want:    12300,
			wantErr: false,
		},
		{
			name: "Float",
			args: args{
				data: []byte("123.45"),
			},
			want:    12345,
			wantErr: false,
		},
		{
			name: "Invalid",
			args: args{
				data: []byte("abc"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money(0)
			if err := m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Money.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != m {
				t.Errorf("Money.UnmarshalJSON() = %v, want %v", m, tt.want)
			}
		})
	}
}

func TestQuantity_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       Quantity
		want    []byte
		wantErr bool
	}{
		{
			name:    "Zero",
			m:       0,
			want:    []byte("0.000"),
			wantErr: false,
		},
		{
			name:    "Int",
			m:       123000,
			want:    []byte("123.000"),
			wantErr: false,
		},
		{
			name:    "Float",
			m:       123456,
			want:    []byte("123.456"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Quantity.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quantity.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuantity_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Quantity
		wantErr bool
	}{
		{
			name: "Zero",
			args: args{
				data: []byte("0"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Integer",
			args: args{
				data: []byte("123"),
			},
			want:    123000,
			wantErr: false,
		},
		{
			name: "Float",
			args: args{
				data: []byte("123.456"),
			},
			want:    123456,
			wantErr: false,
		},
		{
			name: "Invalid",
			args: args{
				data: []byte("abc"),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Quantity(0)
			if err := m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Quantity.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != m {
				t.Errorf("Quantity.UnmarshalJSON() = %v, want %v", m, tt.want)
			}
		})
	}
}

func TestMakeOrderPosition(t *testing.T) {
	type args struct {
		productUuid       string
		productName       string
		price             Money
		priceWithDiscount Money
		quantity          Quantity
	}
	tests := []struct {
		name string
		args args
		want OrderPosition
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeOrderPosition(tt.args.productUuid, tt.args.productName, tt.args.price, tt.args.priceWithDiscount, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeOrderPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeOrder(t *testing.T) {
	type args struct {
		id        string
		comment   string
		contacts  Contacts
		positions []OrderPosition
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeOrder(tt.args.id, tt.args.comment, tt.args.contacts, tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnavailable_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Bool
		wantErr bool
	}{
		{
			name: "true",
			args: args{
				data: []byte(`"1"`),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true",
			args: args{
				data: []byte(`1`),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true",
			args: args{
				data: []byte(`true`),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true",
			args: args{
				data: []byte(`"true"`),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "false",
			args: args{
				data: []byte(`"0"`),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "false",
			args: args{
				data: []byte(`0`),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "false",
			args: args{
				data: []byte(`false`),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "false",
			args: args{
				data: []byte(`"false"`),
			},
			want:    false,
			wantErr: false,
		},

		{
			name: "null",
			args: args{
				data: []byte(`null`),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Invalid",
			args: args{
				data: []byte("`q`"),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Bool(false)
			if err := m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Bool UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != m {
				t.Errorf("Bool UnmarshalJSON() = %v, want %v", m, tt.want)
			}
		})
	}
}
