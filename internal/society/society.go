package society

import (
	"context"
	"log"

	"aimi-landing-back-go/domain/entity"
	g_ "aimi-landing-back-go/domain/repository/gorm"

	"gorm.io/gorm"

	"errors"

	util "aimi-landing-back-go/util"

	pb "aimi-landing-back-go/internal/protorender"
)

type Server struct {
	db *gorm.DB
	pb.UnimplementedSocietyServiceServer
}

func (s *Server) NewServer() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}

func (s *Server) GetAllPostRPC(ctx context.Context, in *pb.RequestPosts) (*pb.ResponsePosts, error) {

	log.Printf("Receive message body from client: %v", in)

	posts := make([]*entity.Post, 0)
	posts_ := make([]*pb.Post, 0)

	if s.db.Find(&posts).Error != nil {
		return nil, errors.New("Societies not found")
	}

	for _, element := range posts {
		post := pb.Post{
			PostId: element.PostId,
			Title: element.Title,
		}
		posts_ = append(posts_, &post)
	}

	return &pb.ResponsePosts{Posts: posts_}, nil
}

func (s *Server) GetPostById(ctx context.Context, in *pb.RequestPost) (*pb.ResponsePost, error) {

	log.Println("Receive message body from client: %v", in)

	var postQuery *entity.Post

	if s.db.Where("post_id = ?", in.PostId).Find(&postQuery).Error != nil {
		return nil, errors.New("post not found")
	}

	post := pb.Post{
		PostId: postQuery.PostId,
		Title: postQuery.Title,
	}

	return &pb.ResponsePost{Post: &post}, nil
}

func (s *Server) GetAllCommentsFromPost(ctx context.Context, in *pb.RequestPost) (*pb.ResponseComments, error) {

	log.Printf("Receive message body from client: %v", in)

	comments := make([]*entity.Comment, 0)
	comments_ := make([]*pb.Comment, 0)

	if s.db.Where("post_id = ?", in.PostId).Find(&comments).Error != nil {
		return nil, errors.New("Societies not found")
	}

	for _, element := range comments {
		comment := pb.Comment{
			PostId: element.PostId,
			Content: element.Content,
			CommentId: element.CommentId,
		}
		comments_ = append(comments_, &comment)
	}

	return &pb.ResponseComments{Comments: comments_}, nil
}

func (s *Server) AddPost(ctx context.Context, in *pb.RequestAddPost) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Post = entity.Post{
		PostId: util.GenUuid(),
		Title: in.Title,
	}

	if s.db.Create(&post).Error != nil {
		return nil, errors.New("something went wrong")
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}

func (s *Server) AddComment(ctx context.Context, in *pb.RequestAddComment) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Comment = entity.Comment{
		CommentId: util.GenUuid(),
		PostId: in.PostId,
		Content: in.Content,
	}

	if s.db.Create(&post).Error != nil {
		return nil, errors.New("something went wrong")
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}

func (s *Server) UpdateComment(ctx context.Context, in *pb.RequestUpdateComment) (*pb.ResponseUpdate, error) {

	log.Printf("Receive message body from client: %v", in)

	var post entity.Comment

	if s.db.Where("comment_id = ?", in.CommentId).Find(&post).Error != nil {
		return nil, errors.New("comment not found")
	}

	post.Content = in.Content

	if s.db.Where("comment_id = ?", in.CommentId).Save(&post).Error != nil {
		return nil, errors.New("something went wrong")
	}

	return &pb.ResponseUpdate{Status: 200}, nil
}