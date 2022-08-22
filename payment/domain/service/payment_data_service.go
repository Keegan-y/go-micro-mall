package service

import (
	"git.imooc.com/keegan/payment/domain/model"
	"git.imooc.com/keegan/payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64 , error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}


//创建
func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService{
	return &PaymentDataService{ paymentRepository }
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}


//插入
func (u *PaymentDataService) AddPayment(payment *model.Payment) (int64 ,error) {
	 return u.PaymentRepository.CreatePayment(payment)
}

//删除
func (u *PaymentDataService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}

//更新
func (u *PaymentDataService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdatePayment(payment)
}

//查找
func (u *PaymentDataService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindPaymentByID(paymentID)
}

//查找
func (u *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}

