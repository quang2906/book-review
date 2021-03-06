package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45312")
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) FindAllReviews() (result []*model.Review) {
	reviews := r.reviews
	for _, review := range reviews {
		result = append(result, review)
	}
	return result
}

func (r *ReviewRepo) FindReviewById2(Id int64) *model.Review {
	if review, ok := r.reviews[Id]; ok {
		return review //tìm được
	} else {
		return nil
	}
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}
