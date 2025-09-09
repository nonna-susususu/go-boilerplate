package todo_test

import (
	"context"

	"github.com/fastworkco/go-boilerplate/internal/domain"
	"github.com/fastworkco/go-boilerplate/internal/service/todo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Todo", func() {
	var (
		ctx                context.Context
		todoService        todo.TodoService
		mockTodoRepository *todo.MockTodoRepository
	)

	BeforeEach(func() {
		ctx = context.Background()

		mockTodoRepository = todo.NewMockTodoRepository(GinkgoT())

		todoService = todo.NewTodoService(todo.TodoServiceDependencies{
			TodoRepository: mockTodoRepository,
		})
	})

	It("should return todo", func() {
		mockTodoRepository.EXPECT().GetAll(mock.Anything).Return([]domain.Todo{
			{
				Task:   "kingdod",
				IsDone: false,
			},
			{
				Task:   "coding",
				IsDone: false,
			},
		}, nil)

		result, err := todoService.GetAllTodo(ctx)

		Expect(result).To(HaveLen(2))
		Expect(err).ToNot(HaveOccurred())
	})

	It("should return error when failed to get todo from repository", func() {
		mockTodoRepository.EXPECT().GetAll(mock.Anything).Return([]domain.Todo{}, assert.AnError)

		result, err := todoService.GetAllTodo(ctx)

		Expect(result).To(HaveLen(0))
		Expect(err).To(HaveOccurred())
	})
})
