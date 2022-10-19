package service

import (
	"go-transfer/internal/datastruct"
	"go-transfer/internal/dto"
	"go-transfer/internal/repository"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProcessTransfer(transfer *dto.Transfer) error {

	if transfer.Account_destination == transfer.Account_origin {
		return echo.NewHTTPError(http.StatusBadRequest, "You can not make a transfer to yourself")
	}

	origin, err := repository.RetrieveAccountDB(transfer.Account_origin)
	if err != nil {
		return err
	}
	destination, err := repository.RetrieveAccountDB(transfer.Account_destination)
	if err != nil {
		return err
	}

	result := origin.Balance - transfer.Amount
	if math.Signbit(result) || origin.Balance == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Insufficient amount to transfer")
	}

	origin.Balance = origin.Balance - transfer.Amount
	destination.Balance = destination.Balance + transfer.Amount

	transferDB := new(datastruct.TransferValues)
	transferDB.Origin = *origin
	transferDB.Destination = *destination

	transfer_validated := new(datastruct.Transfer)
	transfer_validated.Account_origin_id = origin.Id
	transfer_validated.Account_destination_id = destination.Id
	transfer_validated.Amount = transfer.Amount
	transfer_validated.Created_at = transfer.Created_at

	err = repository.TransferDB(transferDB, transfer_validated)

	return err

}
