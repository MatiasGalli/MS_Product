package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MatiasGalli/MS_Product/controllers"
	"github.com/MatiasGalli/MS_Product/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Handler is a struct that contains the methods to handle the messages

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Handler(d amqp.Delivery, channel *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var response models.Response

	actionType := d.Type
	switch actionType {
	case "CREATE_PRODUCT":
		log.Println("Creating product")

		var product models.Product
		err := json.Unmarshal(d.Body, &product)
		failOnError(err, "Failed to unmarshal product")
		productJSON, err := json.Marshal(product)
		failOnError(err, "Failed to marshal product")

		_, err = controllers.CreateProduct(product)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Failed to create product",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "Product created successfully",
				Data:    productJSON,
			}
		}
	case "GET_PRODUCTS":
		log.Println("Getting products")

		product, err := controllers.GetProducts()
		failOnError(err, "Failed to get products")
		productsJSON, err := json.Marshal(product)
		failOnError(err, "Failed to marshal products")

		response = models.Response{
			Success: "success",
			Message: "Products retrieved successfully",
			Data:    productsJSON,
		}

	case "GET_PRODUCT":
		log.Println("Getting product")
		var data struct {
			ID string `json:"id"`
		}

		err := json.Unmarshal(d.Body, &data)
		failOnError(err, "Failed to unmarshal product")

		product, err := controllers.GetProduct(data.ID)
		failOnError(err, "Failed to get product")
		productJSON, err := json.Marshal(product)
		failOnError(err, "Failed to marshal product")

		response = models.Response{
			Success: "success",
			Message: "Product retrieved successfully",
			Data:    productJSON,
		}

	case "CREATE_CATEGORY":
		log.Println("Creating category")
		var category models.Category
		err := json.Unmarshal(d.Body, &category)
		failOnError(err, "Failed to unmarshal category")
		categoryJSON, err := json.Marshal(category)
		failOnError(err, "Failed to marshal category")

		_, err = controllers.CreateCategory(category)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Failed to create category",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "Category created successfully",
				Data:    categoryJSON,
			}
		}

	case "GET_CATEGORIES":
		log.Println("Getting categories")

		category, err := controllers.GetCategories()
		failOnError(err, "Failed to get category")
		categoriesJSON, err := json.Marshal(category)
		failOnError(err, "Failed to marshal category")

		response = models.Response{
			Success: "success",
			Message: "Category retrieved successfully",
			Data:    categoriesJSON,
		}

	case "GET_CATEGORY":
		log.Println("Getting category")
		var data struct {
			ID string `json:"id"`
		}

		err := json.Unmarshal(d.Body, &data)
		failOnError(err, "Failed to unmarshal category")

		category, err := controllers.GetCategory(data.ID)
		failOnError(err, "Failed to get category")
		categoryJSON, err := json.Marshal(category)
		failOnError(err, "Failed to marshal category")

		response = models.Response{
			Success: "success",
			Message: "Category retrieved successfully",
			Data:    categoryJSON,
		}

	}

	responseJSON, err := json.Marshal(response)
	failOnError(err, "Failed to marshal response")

	err = channel.PublishWithContext(ctx,
		"",
		d.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: d.CorrelationId,
			Body:          responseJSON,
		})
	failOnError(err, "Failed to publish a message")

	d.Ack(false)

}
