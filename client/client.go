package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func GetRecommendations(userID int32) ([]*Recommendation, error) {
	// Create to the Python gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := NewRecommendationServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and send request
	request := &RecommendationRequest{
		UserId: userID,
	}
	response, err := client.GetRecommendations(ctx, request)
	if err != nil {
		return nil, err
	}

	// Log recommendations
	// for _, rec := range response.Recommendations {
	// 	log.Printf("Recommendation: ID=%d, Name=%s, Score=%f", rec.Id, rec.Name, rec.Score)
	// }

	return response.Recommendations, nil
}
