package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/davidgamero/bsdata-gql/graph/generated"
	"github.com/davidgamero/bsdata-gql/graph/model"
)

func (r *queryResolver) GetPlayerBrawlerStats(ctx context.Context, playerInput *model.PlayerInput) ([]*model.PlayerBrawlerStats, error) {
	pbs := &model.PlayerBrawlerStats{
		BrawlerName: playerInput.PlayerCode,
		WinLoss:     5,
	}

	resp := []*model.PlayerBrawlerStats{}
	resp = append(resp, pbs)

	getPlayerBrawlersWinloss, err := DB.Query("SELECT * FROM brawlstars.bois_recent_matches")

	if err == nil {
		//getPlayerBrawlersWinlossResult, err := getPlayerBrawlersWinloss.Exec()

		//rowsAffected, _ := getPlayerBrawlersWinloss.Columns()
		cols, err := getPlayerBrawlersWinloss.Columns()
		if err == nil {
			fmt.Print(cols)
		}
		//fmt.Sprintf("%d rows affected\n", rowsAffected)

	} else {

		panic(fmt.Errorf(err.Error()))
	}

	return resp, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
