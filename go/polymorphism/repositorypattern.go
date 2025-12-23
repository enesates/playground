package main

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price float64
}

type Repository interface {
	GetByID(id string) (Product, error)
	GetAll() (map[string]Product, error)
	Save(p Product) (Product, error)
	Delete(id string) (string, error)
}

/////////////////

type MemoryRepository struct {
	Products map[string]Product
}

func (mr MemoryRepository) GetByID(id string) (Product, error) {
	product, ok := mr.Products[id]

	if !ok {
		return Product{}, fmt.Errorf("product %s not found", id)
	}

	return product, nil
}

func (mr MemoryRepository) GetAll() (map[string]Product, error) {
	if mr.Products == nil {
		return nil, fmt.Errorf("products not found")
	}

	return mr.Products, nil
}

func (mr MemoryRepository) Save(p Product) (Product, error) {
	if mr.Products == nil {
		return Product{}, fmt.Errorf("products not found")
	}

	mr.Products[p.ID] = p

	return p, nil
}

func (mr MemoryRepository) Delete(id string) (string, error) {
	product, ok := mr.Products[id]

	if !ok {
		return product.ID, fmt.Errorf("product %s not found", id)
	}

	delete(mr.Products, id)
	return product.ID, nil
}

/////////////////

type MockRepository struct{}

/////////////////

func SaveProduct(r Repository, p Product) {
	_, err := r.Save(p)
	if err != nil {
		fmt.Println("ERROR:", p.ID, err)
		return
	}
	fmt.Printf("Product (%s) saved\n", p.ID)
}

func GetProductByID(r Repository, id string) {
	product, err := r.GetByID(id)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Product found: %+v\n", product)
}

func GetAllProducts(r Repository) {
	products, err := r.GetAll()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Number of products: %d\n", len(products))
}

func DeleteProduct(r Repository, id string) {
	pid, err := r.Delete(id)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Product (%s) deleted\n", pid)
}

func main() {
	mr := MemoryRepository{Products: make(map[string]Product)}

	product := Product{ID: "1", Name: "Book", Price: 10.00}
	product2 := Product{ID: "2", Name: "Table", Price: 80.00}

	SaveProduct(mr, product)
	GetAllProducts(mr)

	SaveProduct(mr, product2)
	GetAllProducts(mr)

	GetProductByID(mr, product.ID)

	DeleteProduct(mr, product.ID)
	GetAllProducts(mr)

	GetProductByID(mr, "3")
}
