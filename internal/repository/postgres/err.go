package postgresrep

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	NoCityError         = fmt.Errorf("no such city")
	NullCustomerIDError = fmt.Errorf("null id")
)

func dbError(err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil
	}

	return err
}
