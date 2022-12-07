package entity

/*
	{
	  "id": 1,
	  "title": "Lemon cheesecake",
	  "description": "A cheesecake made of lemon",
	  "rating": 7,
	  "image": "<https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg>",
	  "created_at": "2020-02-01 10:56:31",
	  "updated_at": "2020-02-13 09:30:23"
	}
*/
// Cake is the struct for a cake
type Cake struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"required,min=1,max=10"`
	Image       string `json:"image" validate:"required,url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
