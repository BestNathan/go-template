package config_test

import (
	"go-template/internal/pkg/config"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("read condig from yaml", func(t *testing.T) {
		c := config.New(config.ConfigFilePathOption("test_data"), config.ConfigFileNameOption("config-test.yaml"))

		a := c.GetViper().GetInt("test.a")
		t.Logf("a value is %d\n", a)
		if a != 1 {
			t.Fatalf("a expect 1 but got: %d", a)
		}

		b := c.GetViper().GetInt("test.b")
		t.Logf("b value is %d\n", b)
		if b != 2 {
			t.Fatalf("b expect 2 but got: %d", b)
		}
	})

	t.Run("read condig from env", func(t *testing.T) {
		os.Setenv("ENV_TEST_A", "2")

		c := config.New(config.ConfigFilePathOption("test_data"), config.ConfigFileNameOption("config-test.yaml"))

		a := c.GetViper().GetInt("test.a")
		t.Logf("a value is %d\n", a)
		if a != 2 {
			t.Fatalf("a expect 2 but got: %d", a)
		}

		b := c.GetViper().GetInt("test.b")
		t.Logf("b value is %d\n", b)
		if b != 2 {
			t.Fatalf("b expect 2 but got: %d", b)
		}
	})
}

func TestUnmarshall(t *testing.T) {
	t.Run("test Unmarshal", func(t *testing.T) {
		c := config.New(config.ConfigFilePathOption("test_data"), config.ConfigFileNameOption("config-test.yaml"))

		type Test struct {
			A int
			B int
		}

		var tt Test

		err := c.Unmarshal("test", &tt)
		if err != nil {
			t.Fatalf("err should be nil but got: %v", err)
		}

		if tt.A != 1 {
			t.Fatalf("a expect 1 but got: %d", tt.A)
		}

		if tt.B != 2 {
			t.Fatalf("b expect 2 but got: %d", tt.B)
		}
	})
}
