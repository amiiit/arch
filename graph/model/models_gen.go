// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Offer struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type SetRolesInput struct {
	UserID string   `json:"userId"`
	Roles  []string `json:"roles"`
}

type Transaction struct {
	Sender   *User             `json:"sender"`
	Receiver *User             `json:"receiver"`
	SentAt   int               `json:"sentAt"`
	Status   TransactionStatus `json:"status"`
}

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone"`
	Region    *string    `json:"region"`
	Roles     *UserRoles `json:"roles"`
	Offers    []*Offer   `json:"offers"`
}

type UserInput struct {
	Username  string  `json:"username"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
	Region    *string `json:"region"`
}

type UserRoles struct {
	Admin  bool `json:"admin"`
	Member bool `json:"member"`
}

type TransactionStatus string

const (
	TransactionStatusAccepted TransactionStatus = "accepted"
	TransactionStatusFailed   TransactionStatus = "failed"
	TransactionStatusRejected TransactionStatus = "rejected"
)

var AllTransactionStatus = []TransactionStatus{
	TransactionStatusAccepted,
	TransactionStatusFailed,
	TransactionStatusRejected,
}

func (e TransactionStatus) IsValid() bool {
	switch e {
	case TransactionStatusAccepted, TransactionStatusFailed, TransactionStatusRejected:
		return true
	}
	return false
}

func (e TransactionStatus) String() string {
	return string(e)
}

func (e *TransactionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionStatus", str)
	}
	return nil
}

func (e TransactionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
