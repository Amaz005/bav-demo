package orm

import (
	g_ "aimi-landing-back-go/domain/repository/gorm"

	"gorm.io/gorm"

	"aimi-landing-back-go/domain/entity"
)

// db, _ := gorm.NewDb()

type SocietyOrm struct {
	db *gorm.DB
}

func (s *SocietyOrm) NewsocietyOrm() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]

}

func (s *SocietyOrm) GetAllPost() ([]*entity.APIPostListAll, error) {
	posts := make([]*entity.APIPostListAll, 0)

	fields := `*`

	rows, _ := s.db.Table("posts").Select(fields).Rows()

	for rows.Next() {
		post := new(entity.APIPostListAll)
		if err := rows.Scan(
			&post.PostId,
			&post.Title,
		); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// photo:,
// person:,
