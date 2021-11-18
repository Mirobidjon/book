package service

import (
	"context"

	"bitbucket.org/udevs/book_service/genproto/user_service"
	"bitbucket.org/udevs/book_service/pkg/helper"
	"bitbucket.org/udevs/book_service/pkg/logger"
	"bitbucket.org/udevs/book_service/storage/sqlc"
	"github.com/xtgo/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type bookService struct {
	db     *sqlc.Queries
	logger logger.Logger
}

func NewBookService(db *sqlc.Queries, log logger.Logger) *bookService {
	return &bookService{
		db:     db,
		logger: log,
	}
}

func (b *bookService) Create(ctx context.Context, req *user_service.CreateBookParams) (*user_service.BookId, error) {
	var book sqlc.CreateBookParams
	err := helper.ParseToStruct(&book, req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while parsing to struct, create book", req, codes.Internal)
	}

	book.ID = uuid.NewRandom().String()

	resp, err := b.db.CreateBook(ctx, book)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while creating book", req, codes.Internal)
	}

	return &user_service.BookId{
		Id: resp.ID,
	}, nil
}

func (b *bookService) Update(ctx context.Context, req *user_service.Book) (*emptypb.Empty, error) {
	var book sqlc.UpdateBookParams
	err := helper.ParseToStruct(&book, req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while parsing to struct, update book", req, codes.Internal)
	}

	err = b.db.UpdateBook(ctx, book)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while updating book", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (b *bookService) Get(ctx context.Context, req *user_service.BookId) (*user_service.Book, error) {
	var resp user_service.Book
	book, err := b.db.GetBook(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while getting book", req, codes.Internal)
	}

	err = helper.ParseToProto(&resp, book)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while parsing to proto, get book", req, codes.Internal)
	}

	return &resp, err
}

func (b *bookService) GetAll(ctx context.Context, req *user_service.GetAllRequest) (*user_service.GetAllBookResponse, error) {
	var (
		params sqlc.GetBooksParams
		books  user_service.GetAllBookResponse
	)

	err := helper.ParseToStruct(&params, req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while parsing to struct, get all book", req, codes.Internal)
	}

	// get counts
	count, err := b.db.GetCount(ctx, req.Search)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while getting all books count", req, codes.Internal)
	}

	// get books
	resp, err := b.db.GetBooks(ctx, params)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while getting all book", req, codes.Internal)
	}

	err = helper.ParseToProto(&books, struct {
		Books []sqlc.Book `json:"books"`
	}{
		Books: resp,
	})

	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while parsing to proto, get all book", req, codes.Internal)
	}

	books.Count = int32(count)

	return &books, nil
}

func (b *bookService) Delete(ctx context.Context, req *user_service.BookId) (*emptypb.Empty, error) {
	err := b.db.DeleteBook(ctx, req.Id)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while deleting book", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
