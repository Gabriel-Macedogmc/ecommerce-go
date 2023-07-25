package services

import (
	"encoding/json"

	"github.com/Gabriel-Macedogmc/products-backend/infra"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/dto"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/model"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/repository"
)

type ProductService struct {
	productRepo repository.ProductRepository
	mq          *infra.RabbitMQ
}

func NewProductService(productRepo repository.ProductRepository, mq *infra.RabbitMQ) *ProductService {
	return &ProductService{productRepo, mq}
}

func (p *ProductService) CreateProduct(data *dto.ProductDTO) (*model.Product, error) {
	product := &model.Product{
		Name:     data.Name,
		Price:    data.Price,
		Quantity: data.Quantity,
	}

	err := p.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetProductById(id uint) (*model.Product, error) {
	product, err := p.productRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) DeleteProduct(id uint) error {
	product, err := p.productRepo.GetById(id)
	if err != nil {
		return err
	}

	err = p.productRepo.Delete(product)
	return err
}

func (p *ProductService) UpdateProduct(id uint, data *dto.ProductDTO) (*model.Product, error) {

	product, err := p.productRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	product.Name = data.Name
	product.Price = data.Price
	product.Quantity = data.Quantity

	err = p.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) CreateOrder(order *dto.OrderDTO) error {
	product, err := p.productRepo.GetById(uint(order.ProductID))
	if err != nil {
		return err
	}

	if product.Quantity == 0 {
		return &domain.ErrInvalidEntity{EntityType: "Product", Message: "Product out of stock"}
	}

	payload, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = p.mq.Publish("", "orders", payload)
	if err != nil {
		return err
	}

	return nil
}
