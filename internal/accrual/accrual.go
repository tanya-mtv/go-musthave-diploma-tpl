package accrual

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/models"

	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/logger"
)

const (
	RetryMax     int           = 3
	RetryWaitMin time.Duration = 1 * time.Second
	RetryMedium  time.Duration = 3 * time.Second
	RetryWaitMax time.Duration = 5 * time.Second
)

type ServiceAccrual struct {
	Storage    orders
	httpClient *http.Client
	log        logger.Logger
	addr       string
}

func NewServiceAccrual(stor orders, log logger.Logger, addr string) *ServiceAccrual {
	return &ServiceAccrual{
		Storage:    stor,
		httpClient: &http.Client{},
		log:        log,
		addr:       addr,
	}
}

func (s *ServiceAccrual) ProcessedAccrualData(ctx context.Context) {
	timer := time.NewTicker(15 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			orders, err := s.Storage.GetOrdersWithStatus()

			if err != nil {
				s.log.Error(err)
			}
			numjobs := len(orders)
			jobs := make(chan models.OrderResponse, numjobs)
			results := make(chan accrServiceResponce)

			for w := 1; w <= 5; w++ {
				go func(w int) {
					s.recieveChainData(ctx, jobs, results, w)
				}(w)
			}
			go func() {

				for j := 1; j <= numjobs; j++ {
					fmt.Println("get order  ", orders[j-1])
					jobs <- orders[j-1]
				}

				close(jobs)
			}()

			for res := range results {
				if res.err != nil {
					s.log.Error(err)

					if res.t != 0 {
						s.log.Info("Too Many Requests")
						timer.Reset(time.Duration(res.t) * time.Second)
					}
					continue
				}

				err = s.Storage.ChangeStatusAndSum(res.ord.Accrual, res.ord.Status, res.ord.Number)

				if err != nil {
					s.log.Error(err)
				}

			}
			close(results)

		case <-ctx.Done():
			return
		}
	}

}

type accrServiceResponce struct {
	ord models.OrderResponse
	t   int
	err error
}

func (s *ServiceAccrual) recieveChainData(ctx context.Context, jobs <-chan models.OrderResponse, res chan<- accrServiceResponce, w int) {
	var accrResponce accrServiceResponce
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-jobs:
			if !ok {
				fmt.Println("<-- loop broke!")
				return
			} else {
				fmt.Println("worker ", w, "send request", val)
				accrResponce.ord, accrResponce.t, accrResponce.err = s.RecieveOrder(ctx, val.Number)
				res <- accrResponce
			}
		}
	}
}

func (s *ServiceAccrual) RecieveOrder(ctx context.Context, number string) (models.OrderResponse, int, error) {
	var orderResp models.OrderResponse
	url := fmt.Sprintf("%s/api/orders/%s", s.addr, number)

	s.log.Info("Recieving order from accrual system ", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {

		s.log.Error(err)
		return orderResp, 0, err
	}

	resp, err := s.httpClient.Do(req)

	if err != nil {
		s.log.Debug("Can't get message")
		return orderResp, 0, err
	}
	defer resp.Body.Close()

	s.log.Info("Get response status ", resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK:

		jsonData, err := io.ReadAll(resp.Body)
		if err != nil {
			s.log.Error(err)
			return orderResp, 0, err
		}

		if err := json.Unmarshal(jsonData, &orderResp); err != nil {
			s.log.Error(err)
			return orderResp, 0, err
		}
		s.log.Info("Get data from accrual system  ", orderResp)

		if orderResp.Status == "REGISTERED" {
			orderResp.Status = "NEW"
		}
		s.log.Info("Get data", orderResp)
		return orderResp, 0, nil
	case http.StatusNoContent:
		s.log.Info("No content in request ")
		return orderResp, 0, errors.New("NoContent")
	case http.StatusTooManyRequests:
		s.log.Info("Too Many Requests ")

		retryHeder := resp.Header.Get("Retry-After")
		retryafter, err := strconv.Atoi(retryHeder)
		if err != nil {
			return orderResp, 0, errors.New("TooManyRequests")
		}

		return orderResp, retryafter, errors.New("TooManyRequests")
	}
	return orderResp, 0, nil
}
