package unit_testing

import "testing"

func TestAddition(t *testing.T) {
	type args struct {
		number  int
		number2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addition Case",
			args: args{
				number:  -75,
				number2: 6,
			},
			want: -69,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.number, tt.args.number2); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		number  int
		number2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Subtraction Case",
			args: args{
				number:  120,
				number2: 10,
			},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtraction(tt.args.number, tt.args.number2); got != tt.want {
				t.Errorf("Subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	type args struct {
		number  int
		number2 int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Positive Division Case",
			args:    args{number: 20, number2: 5},
			want:    4,
			wantErr: false,
		},
		{
			name:    "Negative Division Case",
			args:    args{number: 20, number2: 0},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Division(tt.args.number, tt.args.number2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Division() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	type args struct {
		number  int
		number2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiplication Case",
			args: args{
				number:  28,
				number2: 12,
			},
			want: 336,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplication(tt.args.number, tt.args.number2); got != tt.want {
				t.Errorf("Multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
