package main

import (
	"context"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthV1 "google.golang.org/grpc/health/grpc_health_v1"
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

	serviceToCheck := "readiness" // wait until the service is ready
	expectedStatus := healthV1.HealthCheckResponse_SERVING
	_, healthError := waitForStatus(t, serviceToCheck, expectedStatus)
	require.Nil(t, healthError)

	cmd.Process.Signal(syscall.SIGTERM)
	err = cmd.Wait()
	require.Nil(t, err)
	exitCode := cmd.ProcessState.ExitCode()
	assert.Equal(t, 0, exitCode)
}

func TestMain_StartHttpServer_ProcessAlreadyListenToPort_AppCrash(t *testing.T) {
	httpServerStarted := make(chan struct{})
	if os.Getenv("EXECUTE") == "1" {
		startHttpServer(httpServerStarted, make(chan struct{}), ":8080")
		return
	}

	blockingcmd := exec.Command(os.Args[0], "-test.run=TestMain_StartHttpServer_ProcessAlreadyListenToPort_AppCrash")
	blockingcmd.Env = append(os.Environ(), "EXECUTE=1")
	blockingErr := blockingcmd.Start()
	defer blockingcmd.Process.Kill()
	require.Nil(t, blockingErr)

	<-httpServerStarted

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_StartHttpServer_ProcessAlreadyListenToPort_AppCrash")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 50, exitCode)

}

func TestMain_StartGrpcServer_ProcessAlreadyListenToPort_AppCrash(t *testing.T) {
	grpcServerStarted := make(chan struct{})
	if os.Getenv("EXECUTE") == "1" {
		startHttpServer(grpcServerStarted, make(chan struct{}), ":3110")
		return
	}

	blockingcmd := exec.Command(os.Args[0], "-test.run=TestMain_StartGrpcServer_ProcessAlreadyListenToPort_AppCrash")
	blockingcmd.Env = append(os.Environ(), "EXECUTE=1")
	blockingErr := blockingcmd.Start()
	defer blockingcmd.Process.Kill()
	require.Nil(t, blockingErr)

	<-grpcServerStarted

	cmd := exec.Command(os.Args[0], "-test.run=TestMain_StartGrpcServer_ProcessAlreadyListenToPort_AppCrash")
	cmd.Env = append(os.Environ(), "EXECUTE=1")
	err := cmd.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 50, exitCode)
}

func TestMain_HealtProbeIsServing_liveness(t *testing.T) {
	serviceToCheck := "liveness"

	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	runningApp := exec.Command(os.Args[0], "-test.run=TestMain_HealthProbeIsServing")
	runningApp.Env = append(os.Environ(), "EXECUTE=1")
	blockingErr := runningApp.Start()
	require.Nil(t, blockingErr)
	defer runningApp.Process.Kill()

	expectedStatus := healthV1.HealthCheckResponse_SERVING
	healthResponse, err := waitForStatus(t, serviceToCheck, expectedStatus)

	require.Nil(t, err)
	require.NotNil(t, healthResponse)
	assert.Equal(t, expectedStatus, healthResponse.Status)
}

func TestMain_HealthProbeIsServing_readiness(t *testing.T) {
	serviceToCheck := "readiness"

	if os.Getenv("EXECUTE") == "1" {
		main()
		return
	}

	t.Setenv(document.HostEnv, "http://localhost")
	t.Setenv(document.PortEnv, "8080")
	t.Setenv(oauth.ClientIdEnv, "1234567890")
	t.Setenv(oauth.ClientSecretEnv, "1234567890")
	runningApp := exec.Command(os.Args[0], "-test.run=TestMain_HealthProbeIsServing")
	runningApp.Env = append(os.Environ(), "EXECUTE=1")
	blockingErr := runningApp.Start()
	require.Nil(t, blockingErr)
	defer runningApp.Process.Kill()

	expectedStatus := healthV1.HealthCheckResponse_SERVING
	healthResponse, err := waitForStatus(t, serviceToCheck, expectedStatus)

	require.Nil(t, err)
	require.NotNil(t, healthResponse)
	assert.Equal(t, expectedStatus, healthResponse.Status)
}

func waitForStatus(t *testing.T, serviceToCheck string, expectedStatus healthV1.HealthCheckResponse_ServingStatus) (*healthV1.HealthCheckResponse, error) {
	conn, dialErr := grpc.Dial("localhost:3110", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.Nil(t, dialErr)
	defer conn.Close()

	client := healthV1.NewHealthClient(conn)
	count := 0
	const iterations = 50
	const interval = time.Millisecond * 100
	var healthResponse *healthV1.HealthCheckResponse
	var err error
	for count < iterations {
		healthResponse, err = client.Check(context.Background(), &healthV1.HealthCheckRequest{Service: serviceToCheck})
		if healthResponse != nil && healthResponse.Status == expectedStatus {
			t.Logf("health check succeeded after %v iterations", count)
			break
		} else {
			t.Logf("health check failed after %v iterations", count)
		}
		time.Sleep(interval)
		count++
	}
	return healthResponse, err
}

func TestMain_StartGrpcServer_PortMalformed(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		startGrpcServer(make(chan struct{}), make(chan struct{}), "malformedPort")
		return
	}

	runningApp := exec.Command(os.Args[0], "-test.run=TestMain_StartGrpcServer_PortMalformed")
	runningApp.Env = append(os.Environ(), "EXECUTE=1")
	err := runningApp.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 51, exitCode)
}

func TestMain_StartHttpServer_PortMalformed(t *testing.T) {
	if os.Getenv("EXECUTE") == "1" {
		startHttpServer(make(chan struct{}), make(chan struct{}), "malformedPort")
		return
	}

	runningApp := exec.Command(os.Args[0], "-test.run=TestMain_StartHttpServer_PortMalformed")
	runningApp.Env = append(os.Environ(), "EXECUTE=1")
	err := runningApp.Run()
	require.NotNil(t, err)
	exitCode := err.(*exec.ExitError).ExitCode()
	assert.Equal(t, 50, exitCode)
}
