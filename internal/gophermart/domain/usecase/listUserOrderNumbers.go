package usecase

import (
	"context"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/core"
	"github.com/alexdyukov/gophermart/internal/sharedkernel"
)

type (
	ListUserOrdersRepository interface {
		FindAllOrders(context.Context, string) ([]core.UserOrderNumber, error)
	}

	ListUserOrdersPrimaryPort interface {
		Execute(context.Context, *sharedkernel.User) ([]ListUserOrdersOutputDTO, error)
	}

	ListUserOrdersOutputDTO struct {
		UploadedAt    time.Time          `json:"-"`
		UploadedAtStr string             `json:"uploaded_at"` // nolint:tagliatelle // ok
		Number        string             `json:"number"`
		Status        string             `json:"status"`
		Accrual       sharedkernel.Money `json:"accrual"`
	}

	ListUserOrders struct {
		Repo ListUserOrdersRepository
	}
)

func NewListUserOrders(repo ListUserOrdersRepository) *ListUserOrders {
	return &ListUserOrders{
		Repo: repo,
	}
}

func (l *ListUserOrders) Execute(ctx context.Context, user *sharedkernel.User) ([]ListUserOrdersOutputDTO, error) {
	orders, err := l.Repo.FindAllOrders(ctx, user.ID())
	if err != nil {
		return nil, err
	}

	log.Println(orders)

	lstOrdNumsDTO := make([]ListUserOrdersOutputDTO, 0)

	for _, order := range orders {
		lstOrdNumsDTO = append(lstOrdNumsDTO, ListUserOrdersOutputDTO{
			Number:        strconv.Itoa(order.Number),
			Status:        order.Status.String(),
			Accrual:       order.Accrual,
			UploadedAt:    order.DateAndTime,
			UploadedAtStr: order.DateAndTime.Format(time.RFC3339),
		})
	}

	sort.SliceStable(lstOrdNumsDTO, func(i, j int) bool {
		return lstOrdNumsDTO[i].UploadedAt.Before(lstOrdNumsDTO[j].UploadedAt)
	})

	return lstOrdNumsDTO, nil
}
