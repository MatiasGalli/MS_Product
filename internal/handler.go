package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/MatiasGalli/MS_Product/controllers"
	"github.com/MatiasGalli/MS_Product/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Handler(d amqp.Delivery, channel *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mq models.MessageQueue
	err := json.Unmarshal(d.Body, &mq)
	failOnError(err, "Failed to unmarshal body")
	var response interface{}

	actionType := mq.Pattern
	log.Println("Action type: ", actionType)
	switch actionType {
	case "CREATE_PRODUCT":
		fmt.Print("CREATE_PRODUCT")
		log.Println("Creating product")

		productData, err := json.Marshal(mq.Data)
		failOnError(err, "Failed to marshal product data")

		var product models.Product
		err = json.Unmarshal(productData, &product)
		failOnError(err, "Failed to unmarshal product")

		createdProduct, err := controllers.CreateProduct(product)
		if err != nil {
			response = models.Product{}
		} else {
			response = createdProduct
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
			ID int `json:"id"`
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

		// Convierte el mapa mq.Data en un []byte utilizando json.Marshal
		categoryData, err := json.Marshal(mq.Data)
		failOnError(err, "Failed to marshal category data")

		// Deserializa los datos JSON en un models.Category
		var category models.Category
		err = json.Unmarshal(categoryData, &category)
		failOnError(err, "Failed to unmarshal category")

		// Llama a la función CreateCategory con la categoría deserializada
		createdCategory, err := controllers.CreateCategory(category)
		if err != nil {
			response = models.Category{}

		} else {

			response = createdCategory
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
			ID int `json:"id"`
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
