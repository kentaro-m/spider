package article

import "time"

type Article struct {
	ID        string    `json:"id" example:"faf9c3a7-b3ee-441f-baec-a5b668948381"`
	Title     string    `json:"title" example:"AWS CDKでサーバーレスアプリケーションのデプロイを試す"`
	URL       string    `json:"url" example:"https://blog.kentarom.com/learn-aws-cdk/"`
	PubDate   time.Time `json:"pub_date" example:"2019-01-19T14:13:01Z"`
	CreatedAt time.Time `json:"created_at" example:"2019-04-06T16:03:31Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-04-06T16:03:31Z"`
}
