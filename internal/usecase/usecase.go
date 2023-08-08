package usecase

type usecase struct {
	r repository
}

func NewUsecase(r repository) usecase {
	return usecase{r}
}
