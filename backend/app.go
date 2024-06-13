package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const timeout = 10

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		ctx: context.Background(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListGrpcSymbol(endpoint string, symbol string) ([]string, error) {
	timeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(a.ctx, timeout)
	defer cancel()

	client, err := NewGrpcClient(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if len(symbol) == 0 {
		return client.ListGrpcServices()
	}
	return client.ListGrpcMethods(symbol)
}

func (a *App) DescribeGrpcSymbol(endpoint string, symbol string) (string, error) {
	timeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(a.ctx, timeout)
	defer cancel()

	client, err := NewGrpcClient(ctx, endpoint)
	if err != nil {
		return "", err
	}
	defer client.Close()
	return client.DescribeGrpcSymbol(symbol)
}

func (a *App) GetGrpcMessageTemplate(endpoint string, symbol string) (string, error) {
	timeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(a.ctx, timeout)
	defer cancel()

	client, err := NewGrpcClient(ctx, endpoint)
	if err != nil {
		return "", err
	}
	defer client.Close()
	return client.GetGrpcMessageTemplate(symbol)
}

func (a *App) InvokeGrpc(endpoint string, grpcMethod string, body map[string]interface{}, headers map[string]interface{}) (*CommonResp, error) {
	timeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(a.ctx, timeout)
	defer cancel()

	client, err := NewGrpcClient(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	arr := make([]string, 0)
	for k, v := range headers {
		arr = append(arr, fmt.Sprintf("%s:%s", k, v))
	}

	return client.InvokeGrpc(grpcMethod, string(jsonBytes), arr)
}
