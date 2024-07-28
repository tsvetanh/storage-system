package product

type Product struct {
	ProductID          int      `gorm:"column:product_id;primaryKey" json:"product_id"`
	ProductName        string   `gorm:"column:product_name;size:50;not null" json:"product_name"`
	ProductDescription *string  `gorm:"column:product_description;size:200" json:"product_description,omitempty"`
	CategoryID         int      `gorm:"column:category_id;not null" json:"category_id"`
	SupplierID         int      `gorm:"column:supplier_id;not null" json:"supplier_id"`
	QuantityInStock    int      `gorm:"column:quantity_in_stock;not null" json:"quantity_in_stock"`
	Price              float64  `gorm:"column:price;type:numeric(10,2);not null" json:"price"`
	Category           Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Supplier           Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
}

type Category struct {
	CategoryID   int    `gorm:"column:category_id;primaryKey" json:"category_id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
	Description  string `gorm:"column:description" json:"description,omitempty"`
}

type Supplier struct {
	SupplierID      int    `gorm:"column:supplier_id;primaryKey" json:"supplier_id"`
	SupplierName    string `gorm:"column:supplier_name" json:"supplier_name"`
	ContactName     string `gorm:"column:contact_name" json:"contact_name"`
	ContactEmail    string `gorm:"column:contact_email" json:"contact_email"`
	ContactPhone    string `gorm:"column:contact_phone" json:"contact_phone"`
	SupplierAddress string `gorm:"column:supplier_address" json:"supplier_address"`
}

func (Product) TableName() string {
	return "storageuser.products"
}
