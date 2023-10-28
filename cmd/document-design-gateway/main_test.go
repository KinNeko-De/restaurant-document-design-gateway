package main

import (
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

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
	assert.Equal(t, 40, exitCode)
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
	assert.Equal(t, 41, exitCode)
}

// test does not run on windows
// In case you broke something, the test will run forever
// In the pipeline you will see:
// panic: test timed out after 5m0s
// running tests:
// TestMain_ApplicationListenToInterrupt_GracefullShutdown (5m0s)
func TestMain_ApplicationListenToSIGTERM_AndGracefullyShutdown(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_ApplicationListenToSIGTERM_AndGracefullyShutdown")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Start()
	require.Nil(t, err)
	time.Sleep(1 * time.Second)
	cmd.Process.Signal(syscall.SIGTERM)
	err = cmd.Wait()
	require.Nil(t, err)
	exitCode := cmd.ProcessState.ExitCode()
	assert.Equal(t, 0, exitCode)
}

// test does not run on windows
// In case you broke something, the test will run forever
// In the pipeline you will see:
// panic: test timed out after 5m0s
// running tests:
// TestMain_ApplicationListenToInterrupt_GracefullShutdown (5m0s)
func TestMain_ProcessAlreadyListenToPort_AppCrash(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	blockingcmd := exec.Command(os.Args[0], "-test.run=TestMain_ProcessAlreadyListenToPort_AppCrash")
	blockingcmd.Env = append(os.Environ(), "EXECUTE=1")
	blockingErr := blockingcmd.Start()
	require.Nil(t, blockingErr)
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_ProcessAlreadyListenToPort_AppCrash")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 50, exitCode)
	blockingcmd.Process.Kill()
}
