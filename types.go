package evotorrestogo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Состояние заказа
type OrderState string

const (
	OrderStateNew      OrderState = "new"      // новый
	OrderStatePaid     OrderState = "paid"     // оплачен
	OrderStateDone     OrderState = "done"     // выполнен
	OrderStateCanceled OrderState = "canceled" // отменен
)

// Цена в копейках
type Money uint64

func (m Money) MarshalJSON() ([]byte, error) {
	cents := m % 100
	rubbles := m / 100

	val := fmt.Sprintf("%d.%02d", rubbles, cents)

	return []byte(val), nil
}

func (m *Money) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	str := string(data)
	if str == "null" {
		return nil
	}

	str = strings.Trim(str, "\"")

	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*m = Money(v * 100)
		return nil
	}

	rubbles, err := strconv.ParseUint(str[:dotIndex], 10, 64)
	if err != nil {
		return err
	}
	cents, err := strconv.ParseUint(str[dotIndex+1:], 10, 64)
	if err != nil {
		return err
	}
	*m = Money(rubbles*100 + cents)
	return nil
}

// Количество в тысячных долях
type Quantity uint64

func (m Quantity) MarshalJSON() ([]byte, error) {
	frac := m % 1000
	integ := m / 1000

	val := fmt.Sprintf("%d.%03d", integ, frac)

	return []byte(val), nil
}

func (m *Quantity) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	str := string(data)
	if str == "null" {
		return nil
	}

	str = strings.Trim(str, "\"")

	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*m = Quantity(v * 1000)
		return nil
	}

	integ, err := strconv.ParseUint(str[:dotIndex], 10, 64)
	if err != nil {
		return err
	}
	frac, err := strconv.ParseUint(str[dotIndex+1:], 10, 64)
	if err != nil {
		return err
	}
	*m = Quantity(integ*1000 + frac)
	return nil
}

type Timestamp struct {
	time.Time
}

func (m Timestamp) MarshalJSON() ([]byte, error) {
	val := strconv.FormatInt(m.Unix(), 10)

	return []byte(val), nil
}

func (m *Timestamp) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	str := string(data)
	if str == "null" {
		return nil
	}

	str = strings.Trim(str, "\"")
	sec, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	m.Time = time.Unix(sec, 0)
	return nil
}

// Торговая точка
type Store struct {
	UUID    string `json:"uuid"`              // ИД
	Status  string `json:"status"`            // состояние
	Name    string `json:"name"`              // наименование
	Address string `json:"address,omitempty"` // адрес
}

// Элемент меню
type MenuItem struct {
	UUID          string `json:"uuid"`                    // ИД
	Group         bool   `json:"group"`                   // это группа
	Code          string `json:"code,omitempty"`          // код
	ArticleNumber string `json:"articleNumber,omitempty"` // артикул
	Name          string `json:"name"`                    // наименование
	Price         Money  `json:"price,omitempty"`         // цена в копейках
	MeasureName   string `json:"measureName,omitempty"`   // ЕИ
	AllowToSell   bool   `json:"allowToSell,omitempty"`   // разрешена продажа
	Description   string `json:"description,omitempty"`   // описание
	ParentUUID    string `json:"parentUuid,omitempty"`    // ИД родителя
	Type          string `json:"type,omitempty"`          // вид номенклатуры

	ImageURL      string `json:"image_url,omitempty"` // URL изображения
	IsUnavailable bool   `json:"isUnavailable"`       // присутствие в стоп-листе
}

// Заказ
type Order struct {
	UUID      string          `json:"uuid,omitempty"`      // ИД
	ID        string          `json:"id,omitempty"`        // внешний ИД
	Number    string          `json:"number,omitempty"`    // номер
	Contacts  Contacts        `json:"contacts"`            // контакты
	CreatedAt Timestamp       `json:"createdAt,omitempty"` // дата создания
	UpdatedAt Timestamp       `json:"updatedAt,omitempty"` // дата обновления
	Positions []OrderPosition `json:"positions"`           // позиции
	State     string          `json:"state,omitempty"`     // состояние
	Comment   string          `json:"comment,omitempty"`   // комментарий
}

// Контактные данные
type Contacts struct {
	Phone string `json:"phone"`           // номер телефона
	Email string `json:"email,omitempty"` // адрес электронной почты
}

// Позиция заказа
type OrderPosition struct {
	Position                int      `json:"position,omitempty"`                 // Номер п/п
	ProductUUID             string   `json:"product_uuid"`                       // ИД номенклатуры
	ProductName             string   `json:"product_name,omitempty"`             // наименование
	ProductCode             string   `json:"product_code,omitempty"`             // код
	ProductMeasureName      string   `json:"product_measureName,omitempty"`      // ЕИ
	ProductMeasurePrecision int      `json:"product_measurePrecision,omitempty"` // точность ЕИ
	ProductType             string   `json:"product_type,omitempty"`             // вид номенклатуры
	Price                   Money    `json:"price"`                              // цена
	PriceWithDiscount       Money    `json:"priceWithDiscount"`                  // цена со скидкой
	Quantity                Quantity `json:"quantity"`                           // количество
}

func MakeOrderPosition(productUuid, productName string, price, priceWithDiscount Money, quantity Quantity) OrderPosition {
	return OrderPosition{
		ProductUUID:       productUuid,
		ProductName:       productName,
		Price:             price,
		PriceWithDiscount: priceWithDiscount,
		Quantity:          quantity,
	}
}

func MakeOrder(id, comment string, contacts Contacts, positions []OrderPosition) Order {
	return Order{
		ID:        id,
		Contacts:  contacts,
		Positions: positions,
		Comment:   comment,
	}
}
