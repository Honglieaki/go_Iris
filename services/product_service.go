package services

import (
	"go_Iris/dapmodels"
	"go_Iris/repositories"
)

type IProductService interface {
	GetProductByID(int64) (*dapmodels.Product, error)
	GetAllProduct() ([]*dapmodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *dapmodels.Product) (int64, error)
	UpdateProduct(product *dapmodels.Product) error
}

type ProductService struct {
	productRepository repositories.IProduct
}

// 创建productservice实例
func NewProductService(productRepository repositories.IProduct) IProductService {
	return &ProductService{productRepository}
}

/*
	GetProductByID(int64) (*dapmodels.Product, error)
	GetAllProduct() ([]*dapmodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *dapmodels.Product) (int64, error)
	UpdateProduct(product *dapmodels.Product) error
*/

// 查询所有商品
func (p *ProductService) GetAllProduct() ([]*dapmodels.Product, error) {
	return p.productRepository.SelectAll()
}

// 根据ID查询单个商品
func (p *ProductService) GetProductByID(ProductId int64) (*dapmodels.Product, error) {
	return p.productRepository.SelectById(ProductId)
}

// 删除某个商品
func (p *ProductService) DeleteProductByID(ProductId int64) bool {
	return p.productRepository.Delete(ProductId)
}

func (p *ProductService) InsertProduct(product *dapmodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

func (p *ProductService) UpdateProduct(product *dapmodels.Product) error {
	return p.productRepository.Update(product)
}
