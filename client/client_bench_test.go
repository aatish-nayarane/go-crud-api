package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vm-crud/internal/model"
)

func BenchmarkCreateVM(b *testing.B) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			w.Write([]byte(`{
				"id": 1,
				"name": "Test VM",
				"cpu": 2,
				"ram": 4096,
				"status": "stopped"
			}`))
		}),
	)

	defer server.Close()

	client := NewClient(server.URL)

	req := model.CreateVMRequest{
		Name: "Test VM",
		CPU:  2,
		RAM:  4096,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := client.CreateVM(req)
		if err != nil {
			b.Fatal(err)
		}
	}
}
