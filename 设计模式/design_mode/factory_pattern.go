package main

import (
	"errors"
	"fmt"
)

// 实现示例
//	需要一个支付的产品，可以是微信支付，支付宝支付，银联支付
//- 抽象产品：支付
//- 具体产品：支付宝支付，微信支付，银联支付
//- 抽象工厂：创建产品的方法
//- 具体工厂：实例化支付宝支付类，实例化微信支付类，实例化银联支付类

// IPay define abstract product
type IPay interface {
	Pay(m int) error
}

// AliPayProd define concrete product class pf aliPay
type AliPayProd struct {
}

func (a *AliPayProd) Pay(m int) error {
	if m <= 0 {
		return errors.New("alipay amount can not less zero")
	}
	fmt.Println("alipay pay money: ", m)
	return nil
}

// WeiChatPayProd define concrete product class of weiChatPay
type WeiChatPayProd struct {
}

func (w *WeiChatPayProd) Pay(m int) error {
	if m <= 0 {
		return errors.New("weichat pay amount can not less zero")
	}
	fmt.Println("weichat pay money: ", m)
	return nil
}

type PayType uint8

const (
	PayTypeAli PayType = iota
	PayTypeWeiChat
	PayTypeUnionCard
)

// INewPay define abstract factory of pay
type INewPay interface {
	newPay(t PayType) IPay
}

// AliPayFactory define concrete factory of pay
type AliPayFactory struct {
}

// NewPay new AliPayProd instance
func (a AliPayFactory) newPay() IPay {
	return &AliPayProd{}
}

// WeiChatFactory defined concrete factory of weiChat
type WeiChatFactory struct {
}

func (w WeiChatFactory) newPay() IPay {
	return &WeiChatPayProd{}
}

//func newPay(t PayType) IPay {
//	switch t {
//	case PayTypeAli:
//		return AliPayFactory{}.newPay()
//	case PayTypeWeiChat:
//		return WeiChatFactory{}.newPay()
//	}
//	return nil
//}
//
//
//// RunFactoryPattern test factory pattern
//func RunFactoryPattern() {
//	p := newPay(PayTypeAli)
//	err := p.Pay(10)
//	if err != nil {
//		return
//	}
//
//	p = newPay(PayTypeWeiChat)
//	err = p.Pay(20)
//	if err != nil {
//		return
//	}
//}

// ----------------------------------------------加入新的产品和工厂--------------------

// UnionPayProd Add new pay mode, need defined concrete prod and factory
// define union pay prod class
type UnionPayProd struct {
}

func (u *UnionPayProd) Pay(m int) error {
	if m <= 0 {
		return errors.New("unionpay amount can not less zero")
	}
	fmt.Println("unionpay pay pay money: ", m)
	return nil
}

// UnionPayFactory defined concrete factory of union pay
type UnionPayFactory struct {
}

func (u UnionPayFactory) newPay() IPay {
	return &UnionPayProd{}
}

func NewPay(t PayType) IPay {
	switch t {
	case PayTypeAli:
		return AliPayFactory{}.newPay()
	case PayTypeWeiChat:
		return WeiChatFactory{}.newPay()
	case PayTypeUnionCard:
		return UnionPayFactory{}.newPay()
	}
	return nil
}

// RunFactoryPattern test factory pattern
func RunFactoryPattern() {
	p := NewPay(PayTypeAli)
	err := p.Pay(10)
	if err != nil {
		return
	}

	p = NewPay(PayTypeWeiChat)
	err = p.Pay(20)
	if err != nil {
		return
	}

	p = NewPay(PayTypeUnionCard)
	err = p.Pay(1000)
	if err != nil {
		return
	}
}
