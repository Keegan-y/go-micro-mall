package repository
import (
	"github.com/jinzhu/gorm"
	"git.imooc.com/keegan/payment/domain/model"
)
type IPaymentRepository interface{
    InitTable() error
    FindPaymentByID(int64) (*model.Payment, error)
	CreatePayment(*model.Payment) (int64, error)
	DeletePaymentByID(int64) error
	UpdatePayment(*model.Payment) error
	FindAll()([]model.Payment,error)

}
//创建paymentRepository
func NewPaymentRepository(db *gorm.DB) IPaymentRepository  {
	return &PaymentRepository{mysqlDb:db}
}

type PaymentRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *PaymentRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Payment{}).Error
}

//根据ID查找Payment信息
func (u *PaymentRepository)FindPaymentByID(paymentID int64) (payment *model.Payment,err error) {
	payment = &model.Payment{}
	return payment, u.mysqlDb.First(payment,paymentID).Error
}

//创建Payment信息
func (u *PaymentRepository) CreatePayment(payment *model.Payment) (int64, error) {
	return payment.ID, u.mysqlDb.Create(payment).Error
}

//根据ID删除Payment信息
func (u *PaymentRepository) DeletePaymentByID(paymentID int64) error {
	return u.mysqlDb.Where("id = ?",paymentID).Delete(&model.Payment{}).Error
}

//更新Payment信息
func (u *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	return u.mysqlDb.Model(payment).Update(payment).Error
}

//获取结果集
func (u *PaymentRepository) FindAll()(paymentAll []model.Payment,err error) {
	return paymentAll, u.mysqlDb.Find(&paymentAll).Error
}

