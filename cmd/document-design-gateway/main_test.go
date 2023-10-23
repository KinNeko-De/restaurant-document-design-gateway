package main

import (
	"os"
	"os/exec"
	"syscall"
	"testing"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain_GatewayConfigIsMissing(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_GatewayConfigIsMissing")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 1, exitCode)
}

func TestMain_OAuthConfigIsMissing(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_OAuthConfigIsMissing")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 1, exitCode)
}

func TestMain_ApplicationListenToInterrupt_GracefullShutdown(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_ApplicationListenToInterrupt_GracefullShutdown")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Start()
	require.Nil(t, err)
	cmd.Process.Signal(syscall.SIGTERM)
	err = cmd.Wait()
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 0, exitCode)
}
