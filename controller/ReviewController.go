package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	if _, err := repo.Books.FindBookById(review.BookId); err != nil {
		return c.JSON(fiber.Map{
			"message": "That book does not exist",
		})
	}

	reviewId := repo.Reviews.CreateNewReview(review)
	repo.UpdateBookRating(repo.Books.FindBookById2(review.BookId))
	return c.SendString(fmt.Sprintf("New review is created successfully with id = %d", reviewId))

}

func DeleteReviewById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	// repo.UpdateBookRating(repo.Books.FindBookById2(repo.Reviews.FindReviewById2(int64(id)).BookId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		return c.SendString("delete successfully")
	}

}

func AvgRating(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	_, err := repo.Books.FindBookById(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Not found book for this id",
		})
	}

	result := repo.Reviews.AvgRating()

	return c.JSON(result[int64(id)])
}
