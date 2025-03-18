package domain

type TestDomain struct {
}

type TestUsecase interface {
	TestService() bool
}

type TestRepository interface {
	TestRepo() bool
}

type TestConsume interface {
	TestConsume()
	TestConsumeWithFunc()
}

type TestCron interface {
	TestCron()
}
