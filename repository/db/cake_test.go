package db

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guntoroyk/cake-store-api/entity"
)

func Test_cakeRepo_GetCakes(t *testing.T) {
	columns := []string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Cake
		wantErr bool
	}{
		{
			name: "success GetCakes",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectQuery("SELECT id, title, description, rating, image, created_at, updated_at FROM cakes").WillReturnRows(
						sqlmock.NewRows(columns).
							AddRow(1, "title", "description", 1, "image", "2006-01-01 00:00:00", "2006-01-01 00:00:00"),
					)

					return db
				}(),
			},
			want: []*entity.Cake{
				{
					ID:          1,
					Title:       "title",
					Description: "description",
					Rating:      1,
					Image:       "image",
					CreatedAt:   "2006-01-01 00:00:00",
					UpdatedAt:   "2006-01-01 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cakeRepo{
				db: tt.fields.db,
			}
			got, err := c.GetCakes()
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeRepo.GetCakes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeRepo.GetCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeRepo_GetCake(t *testing.T) {
	columns := []string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}

	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Cake
		wantErr bool
	}{
		{
			name: "success GetCake",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectQuery("SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?").
						WithArgs(1).
						WillReturnRows(
							sqlmock.NewRows(columns).
								AddRow(1, "title", "description", 1, "image", "2006-01-01 00:00:00", "2006-01-01 00:00:00"),
						)

					return db
				}(),
			},
			args: args{
				id: 1,
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "title",
				Description: "description",
				Rating:      1,
				Image:       "image",
				CreatedAt:   "2006-01-01 00:00:00",
				UpdatedAt:   "2006-01-01 00:00:00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cakeRepo{
				db: tt.fields.db,
			}
			got, err := c.GetCake(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeRepo.GetCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeRepo.GetCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeRepo_CreateCake(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		cake *entity.Cake
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Cake
		wantErr bool
	}{
		{
			name: "success CreateCake",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("INSERT INTO cakes").
						WithArgs("title", "description", 1, "image", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				cake: &entity.Cake{
					Title:       "title",
					Description: "description",
					Rating:      1,
					Image:       "image",
				},
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "title",
				Description: "description",
				Rating:      1,
				Image:       "image",
				CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cakeRepo{
				db: tt.fields.db,
			}
			got, err := c.CreateCake(tt.args.cake)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeRepo.CreateCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeRepo.CreateCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeRepo_UpdateCake(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		cake *entity.Cake
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Cake
		wantErr bool
	}{
		{
			name: "success UpdateCake",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("UPDATE cakes").
						WithArgs("title", "description", 1, "image", time.Now().Format("2006-01-02 15:04:05")).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				cake: &entity.Cake{
					ID:          1,
					Title:       "title",
					Description: "description",
					Rating:      1,
					Image:       "image",
				},
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "title",
				Description: "description",
				Rating:      1,
				Image:       "image",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cakeRepo{
				db: tt.fields.db,
			}
			got, err := c.UpdateCake(tt.args.cake)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeRepo.UpdateCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeRepo.UpdateCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeRepo_DeleteCake(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success DeleteCake",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("DELETE FROM cakes").
						WithArgs(1).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cakeRepo{
				db: tt.fields.db,
			}
			if err := c.DeleteCake(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("cakeRepo.DeleteCake() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
