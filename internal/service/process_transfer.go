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

	origin, err := repository.RetrieveAccount(transfer.Account_origin)
	if err != nil {
		return err
	}
	destination, err := repository.RetrieveAccount(transfer.Account_destination)
	if err != nil {
		return err
	}

	// TO-DO: Write the logic that subtracts the balance, return error if invalid operation, check if balance is 0

	result := origin.Balance - transfer.Amount
	if math.Signbit(result) || origin.Balance == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Insuficient amount to transfer")
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
