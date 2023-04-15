package datamodels

// 商品编号 商品名称 商品数量 商品图片 商品URL地址

type Product struct {
	ID           int64  `json:"id" sql:"ID" imooc:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" imooc:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" imooc:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" imooc:"ProductUrl"`
}
