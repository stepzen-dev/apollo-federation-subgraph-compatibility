// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Product struct {
	ID         string            `json:"id"`
	Sku        *string           `json:"sku"`
	Package    *string           `json:"package"`
	Variation  *ProductVariation `json:"variation"`
	Dimensions *ProductDimension `json:"dimensions"`
	CreatedBy  *User             `json:"createdBy"`
	Notes      *string           `json:"notes"`
}

func (Product) IsEntity() {}

type ProductDimension struct {
	Size   *string  `json:"size"`
	Weight *float64 `json:"weight"`
	Unit   *string  `json:"unit"`
}

type ProductVariation struct {
	ID string `json:"id"`
}

type User struct {
	Email                string  `json:"email"`
	Name                 *string `json:"name"`
	TotalProductsCreated *int    `json:"totalProductsCreated"`
}

func (User) IsEntity() {}