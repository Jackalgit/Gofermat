package loyaltysystem

import (
	"github.com/Jackalgit/Gofermat/cmd/config"
	"github.com/Jackalgit/Gofermat/internal/jsondecoder"
	"github.com/Jackalgit/Gofermat/internal/models"
	"log"
	"net/http"
)

func CheckStatusOrder(orderList []models.OrderStatus) ([]models.OrderStatus, map[string]models.OrderStatus) {

	var orderListCheckStatus []models.OrderStatus
	dictOrderStatusForUpdateDB := make(map[string]models.OrderStatus)

	for _, v := range orderList {
		if v.Status != "INVALID" || v.Status != "PROCESSED" {

			UrlRequest := config.Config.AccrualSystem + "/api/orders/" + v.NumOrder
			response, err := http.Get(UrlRequest)
			if err != nil {
				log.Printf("[Get], %q", err)
				return nil, nil
			}

			responsLoyaltySystem, err := jsondecoder.ResponsLoyaltySystem(response.Body)
			if err != nil {
				log.Printf("[ResponsLoyaltySystem], %q", err)
				return nil, nil
			}
			if v.Status != responsLoyaltySystem.Status {
				dictOrderStatusForUpdateDB[v.NumOrder] = models.OrderStatus{
					Status:  responsLoyaltySystem.Status,
					Accrual: v.Accrual + responsLoyaltySystem.Accrual}
				v.Status = responsLoyaltySystem.Status
				v.Accrual += responsLoyaltySystem.Accrual
			}
		}

		orderListCheckStatus = append(
			orderListCheckStatus,
			models.OrderStatus{NumOrder: v.NumOrder, Status: v.Status, Accrual: v.Accrual, Uploaded_at: v.Uploaded_at})
	}

	return orderListCheckStatus, dictOrderStatusForUpdateDB

}
