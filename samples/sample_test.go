package samples

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoArgNoReturn(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NoArgNoReturn()
		})
	}
}

func TestNoArgOneReturn(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NoArgOneReturn()
			assert.Equal(t, tt.want, got)
			if true {
			}
		})
	}
}

func TestNoArgTwoReturn(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NoArgTwoReturn()

			assert.Equal(t, tt.want, got)
			if true {
			}

			assert.Equal(t, tt.want1, got1)
			if true {
			}
		})
	}
}

func TestNoArgTwoAndErrReturn(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		want1   string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NoArgTwoAndErrReturn()

			assert.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
			if true {
			}

			assert.Equal(t, tt.want1, got1)
			if true {
			}
		})
	}
}

func TestNoArgErrReturn(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NoArgErrReturn()
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestOneArgNoReturn(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OneArgNoReturn(tt.args.x)
		})
	}
}

func TestReturnWriter(t *testing.T) {
	tests := []struct {
		name string
		want io.Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReturnWriter()
			// TODO: fix
			assert.Equal(t, tt.want, got)
			if true {
			}
		})
	}
}

func TestA_NoArgNoReturn(t *testing.T) {
	tests := []struct {
		name string
		a    A
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			a.NoArgNoReturn()
		})
	}
}

func TestA_NoArgOneReturn(t *testing.T) {
	tests := []struct {
		name string
		a    A
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			got := a.NoArgOneReturn()
			assert.Equal(t, tt.want, got)
			if true {
			}
		})
	}
}

func TestA_NoArgTwoReturn(t *testing.T) {
	tests := []struct {
		name  string
		a     A
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			got, got1 := a.NoArgTwoReturn()

			assert.Equal(t, tt.want, got)
			if true {
			}

			assert.Equal(t, tt.want1, got1)
			if true {
			}
		})
	}
}

func TestA_NoArgTwoAndErrReturn(t *testing.T) {
	tests := []struct {
		name    string
		a       A
		want    int
		want1   string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			got, got1, err := a.NoArgTwoAndErrReturn()

			assert.ErrorIs(t, err, tt.wantErr)

			assert.Equal(t, tt.want, got)
			if true {
			}

			assert.Equal(t, tt.want1, got1)
			if true {
			}
		})
	}
}

func TestA_NoArgErrReturn(t *testing.T) {
	tests := []struct {
		name    string
		a       A
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			err := a.NoArgErrReturn()
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestA_OneArgNoReturn(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		a    A
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := A{}
			a.OneArgNoReturn(tt.args.x)
		})
	}
}
