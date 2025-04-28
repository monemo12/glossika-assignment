package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"glossika-assignment/internal/model"
	"glossika-assignment/internal/repository"

	"github.com/golang/mock/gomock"
)

func TestRecommendationService_GetRecommendations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockIRecommendationRepository(ctrl)
	svc := NewRecommendationService(mockRepo)

	ctx := context.Background()
	req := &model.RecommendationRequest{Limit: 2, Offset: 0}

	t.Run("success", func(t *testing.T) {
		items := []*model.Recommendation{
			{ID: 1, Title: "A", Description: "descA", Score: 1.1},
			{ID: 2, Title: "B", Description: "descB", Score: 2.2},
		}
		total := 5
		mockRepo.EXPECT().FetchItemsByPagination(ctx, req.Limit, req.Offset).Return(items, nil)
		mockRepo.EXPECT().FetchItemsCount(ctx).Return(total, nil)

		resp, err := svc.GetRecommendations(ctx, req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(resp.Items, items) {
			t.Errorf("items not match: got %+v, want %+v", resp.Items, items)
		}
		if resp.Total != total {
			t.Errorf("total not match: got %d, want %d", resp.Total, total)
		}
		if !resp.NextPage {
			t.Errorf("NextPage should be true")
		}
	})

	t.Run("repo fetch error", func(t *testing.T) {
		errFetch := errors.New("fetch error")
		mockRepo.EXPECT().FetchItemsByPagination(ctx, req.Limit, req.Offset).Return(nil, errFetch)

		resp, err := svc.GetRecommendations(ctx, req)
		if err == nil || err.Error() != "fetch error" {
			t.Errorf("expected fetch error, got %v", err)
		}
		if resp != nil {
			t.Errorf("expected nil response, got %+v", resp)
		}
	})

	t.Run("repo count error", func(t *testing.T) {
		items := []*model.Recommendation{{ID: 1, Title: "A", Description: "descA", Score: 1.1}}
		errCount := errors.New("count error")
		mockRepo.EXPECT().FetchItemsByPagination(ctx, req.Limit, req.Offset).Return(items, nil)
		mockRepo.EXPECT().FetchItemsCount(ctx).Return(0, errCount)

		resp, err := svc.GetRecommendations(ctx, req)
		if err == nil || err.Error() != "count error" {
			t.Errorf("expected count error, got %v", err)
		}
		if resp != nil {
			t.Errorf("expected nil response, got %+v", resp)
		}
	})

	t.Run("no next page", func(t *testing.T) {
		items := []*model.Recommendation{{ID: 1, Title: "A", Description: "descA", Score: 1.1}}
		total := 1
		mockRepo.EXPECT().FetchItemsByPagination(ctx, req.Limit, req.Offset).Return(items, nil)
		mockRepo.EXPECT().FetchItemsCount(ctx).Return(total, nil)

		resp, err := svc.GetRecommendations(ctx, req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.NextPage {
			t.Errorf("NextPage should be false")
		}
	})
}
