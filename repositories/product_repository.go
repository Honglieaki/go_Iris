package repositories

import (
	"database/sql"
	"go_Iris/common"
	"go_Iris/dapmodels"
	"strconv"
)

// 第一步 开发对应的接口
// 第二步 实现接口中的方法

type IProduct interface {
	// 数据库连接
	Conn() error
	// 增删查改
	Insert(*dapmodels.Product) (int64, error)
	Delete(int64) bool
	Update(*dapmodels.Product) error
	SelectById(int64) (*dapmodels.Product, error)
	SelectAll() ([]*dapmodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlconn *sql.DB
}

//创建ProductManager实例

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlconn: db}
}

// 实现接口内的方法
/*
Conn() error
	// 增删查改
	Insert(*dapmodels.Product) (int, error)
	Delete(int64) bool
	Update(*dapmodels.Product) error
	SelectById(int64) (*dapmodels.Product, error)
	SecectAll() ([]*dapmodels.Product, error)
*/
// 数据库连接
func (p *ProductManager) Conn() (err error) {
	if p.mysqlconn == nil {
		conn, errconn := common.NewMysqlConn()
		if errconn != nil {
			return errconn
		}
		p.mysqlconn = conn
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// 新增商品
func (p *ProductManager) Insert(product *dapmodels.Product) (id int64, err error) {
	// 1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return
	}
	// 2.准备sql语句
	sql := "INSERT" + p.table + "SET productName=?,productNum=?,productImage=?,productUrl=?"
	stml, err := p.mysqlconn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	// 3.传入参数
	result, err := stml.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	// sql语句执行失败
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 删除商品
func (p *ProductManager) Delete(ProductId int64) bool {
	// 1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "delete from product where ID=?"
	stmt, err := p.mysqlconn.Prepare(sql)
	if err != nil {
		return false
	}

	_, err = stmt.Exec(strconv.FormatInt(ProductId, 10))
	if err != nil {
		return false
	}
	return true
}

// 修改商品
func (p *ProductManager) Update(product *dapmodels.Product) (err error) {
	// 1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return err
	}

	sql := "Update product set productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)
	stml, err := p.mysqlconn.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stml.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductImage)
	if err != nil {
		return err
	}
	return nil
}

// 根据商品ID查询商品
func (p *ProductManager) SelectById(ProductID int64) (productResult *dapmodels.Product, err error) {
	// 1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from" + p.table + "where ID =" + strconv.FormatInt(ProductID, 10)
	stml, err := p.mysqlconn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	row, err := stml.Query(sql)
	defer row.Close() // 查询完需要关闭
	if err != nil {
		return nil, err
	}

	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &dapmodels.Product{}, nil
	}
	productResult = &dapmodels.Product{}
	common.DataToStructByTagSql(result, productResult)
	return
}

// 查询所有商品

func (p *ProductManager) SelectAll() (Productslice []*dapmodels.Product, err error) {
	// 1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from" + p.table
	rows, err := p.mysqlconn.Query(sql)
	if err != nil {
		return nil, err
	}

	result := common.GetResultRows(rows)

	if len(result) == 0 {
		return nil, nil
	}

	for _, v := range result {
		product := &dapmodels.Product{}
		common.DataToStructByTagSql(v, product)
		Productslice = append(Productslice, product)
	}
	return Productslice, err
}
