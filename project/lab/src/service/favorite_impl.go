package service

import (
	"context"

	"github.com/golang-class/lab/connector"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/repository"
)

type RealFavoriteService struct {
	favoriteRepository repository.FavoriteRepository
	MovieAPIConnector connector.MovieAPIConnector
}

func (r *RealFavoriteService) GetFavorite(c context.Context) ([]model.FavoriteMovie, error) {
	favorite, err := r.favoriteRepository.GetFavorite(c)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func NewRealFavoriteService(favoriteRepository repository.FavoriteRepository, movieAPIConnector connector.MovieAPIConnector) FavoriteService {
	return &RealFavoriteService{
		favoriteRepository: favoriteRepository,
		MovieAPIConnector: movieAPIConnector,
	}
}

func (r *RealFavoriteService) AddFavorite(ctx context.Context, movieId string) error {
	movie, err := r.MovieAPIConnector.GetMovieDetail(ctx, movieId)
	if err != nil {
			return err
	}
	favoriteMovie := model.FavoriteMovie {
		MovieID: movie.MovieID,
		Title: movie.Title,
		Year:  movie.Year,
		Rating: movie.Rating,
		
	}
	err = r.favoriteRepository.AddFavorite(ctx, favoriteMovie)
	if err != nil {
		return err
	}	
	return nil
}
