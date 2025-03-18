package usecase

import "go-quickstart/domain"

type testUsecase struct {
	testRepository domain.TestRepository
}

func (t *testUsecase) TestService() bool {
	return t.testRepository.TestRepo()
}

func NewTestUsecase(dt domain.TestRepository) domain.TestUsecase {
	return &testUsecase{testRepository: dt}
}
