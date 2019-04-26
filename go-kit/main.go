package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kkoehler/golang/go-kit/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kkoehler/golang/go-kit/model"
)

func makeEndpoint(sf service.StringFlipper) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.StringFlipperRequest)
		v := sf.FlipString(req.Value)
		return model.StringFlipperResponse{v}, nil
	}
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.StringFlipperRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	flipper := service.StringFlipperImpl{}

	handler := httptransport.NewServer(
		makeEndpoint(flipper),
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
