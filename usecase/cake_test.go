package usecase

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/cake-store-api/entity"
	"github.com/guntoroyk/cake-store-api/mocks"
	"github.com/guntoroyk/cake-store-api/repository"
)

func TestNewCakeUsecase(t *testing.T) {
	type args struct {
		cakeRepo repository.CakeRepoItf
	}
	tests := []struct {
		name string
		args args
		want CakeUsecaseItf
	}{
		{
			name: "TestNewCakeUsecase",
			args: args{
				cakeRepo: nil,
			},
			want: &cakeUsecase{
				cakeRepo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCakeUsecase(tt.args.cakeRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCakeUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeUsecase_GetCakes(t *testing.T) {
	type fields struct {
		cakeRepo func(ctrl *gomock.Controller) repository.CakeRepoItf
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Cake
		wantErr bool
	}{
		{
			name: "success get cakes",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().GetCakes().Return([]*entity.Cake{
						{
							ID:          1,
							Title:       "Cake 1",
							Description: "Desc 1",
							Rating:      1,
							Image:       "https://dummyimage.com/600x400/000/fff",
							CreatedAt:   "2021-01-01 00:00:00",
							UpdatedAt:   "2021-01-01 00:00:00",
						},
					}, nil)
					return mockCakeRepo
				},
			},
			want: []*entity.Cake{
				{
					ID:          1,
					Title:       "Cake 1",
					Description: "Desc 1",
					Rating:      1,
					Image:       "https://dummyimage.com/600x400/000/fff",
					CreatedAt:   "2021-01-01 00:00:00",
					UpdatedAt:   "2021-01-01 00:00:00",
				},
			},
			wantErr: false,
		},
		{
			name: "failed get cakes",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().GetCakes().Return(nil, errors.New("error"))
					return mockCakeRepo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &cakeUsecase{
				cakeRepo: tt.fields.cakeRepo(ctrl),
			}
			got, err := c.GetCakes()
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeUsecase.GetCakes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeUsecase.GetCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeUsecase_GetCake(t *testing.T) {
	type fields struct {
		cakeRepo func(ctrl *gomock.Controller) repository.CakeRepoItf
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
			name: "success get cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().GetCake(1).Return(&entity.Cake{
						ID:          1,
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
						CreatedAt:   "2021-01-01 00:00:00",
						UpdatedAt:   "2021-01-01 00:00:00",
					}, nil)
					return mockCakeRepo
				},
			},
			args: args{
				id: 1,
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "Cake 1",
				Description: "Desc 1",
				Rating:      1,
				Image:       "https://dummyimage.com/600x400/000/fff",
				CreatedAt:   "2021-01-01 00:00:00",
				UpdatedAt:   "2021-01-01 00:00:00",
			},
			wantErr: false,
		},
		{
			name: "failed get cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().GetCake(1).Return(nil, errors.New("error"))
					return mockCakeRepo
				},
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get cake not found",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().GetCake(1).Return(nil, sql.ErrNoRows)
					return mockCakeRepo
				},
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &cakeUsecase{
				cakeRepo: tt.fields.cakeRepo(ctrl),
			}
			got, err := c.GetCake(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeUsecase.GetCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeUsecase.GetCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeUsecase_CreateCake(t *testing.T) {
	type fields struct {
		cakeRepo func(ctrl *gomock.Controller) repository.CakeRepoItf
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
			name: "success create cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().CreateCake(&entity.Cake{
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
					}).Return(&entity.Cake{
						ID:          1,
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
						CreatedAt:   "2021-01-01 00:00:00",
						UpdatedAt:   "2021-01-01 00:00:00",
					}, nil)
					return mockCakeRepo
				},
			},
			args: args{
				cake: &entity.Cake{
					Title:       "Cake 1",
					Description: "Desc 1",
					Rating:      1,
					Image:       "https://dummyimage.com/600x400/000/fff",
				},
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "Cake 1",
				Description: "Desc 1",
				Rating:      1,
				Image:       "https://dummyimage.com/600x400/000/fff",
				CreatedAt:   "2021-01-01 00:00:00",
				UpdatedAt:   "2021-01-01 00:00:00",
			},
		},
		{
			name: "failed create cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().CreateCake(&entity.Cake{
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
					}).Return(nil, errors.New("error"))
					return mockCakeRepo
				},
			},
			args: args{
				cake: &entity.Cake{
					Title:       "Cake 1",
					Description: "Desc 1",
					Rating:      1,
					Image:       "https://dummyimage.com/600x400/000/fff",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &cakeUsecase{
				cakeRepo: tt.fields.cakeRepo(ctrl),
			}
			got, err := c.CreateCake(tt.args.cake)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeUsecase.CreateCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeUsecase.CreateCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeUsecase_UpdateCake(t *testing.T) {
	type fields struct {
		cakeRepo func(ctrl *gomock.Controller) repository.CakeRepoItf
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
			name: "success update cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().UpdateCake(&entity.Cake{
						ID:          1,
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
					}).Return(&entity.Cake{
						ID:          1,
						Title:       "Cake 1",
						Description: "Desc 1",
						Rating:      1,
						Image:       "https://dummyimage.com/600x400/000/fff",
						CreatedAt:   "2021-01-01 00:00:00",
						UpdatedAt:   "2021-01-01 00:00:00",
					}, nil)
					return mockCakeRepo
				},
			},
			args: args{
				cake: &entity.Cake{
					ID:          1,
					Title:       "Cake 1",
					Description: "Desc 1",
					Rating:      1,
					Image:       "https://dummyimage.com/600x400/000/fff",
				},
			},
			want: &entity.Cake{
				ID:          1,
				Title:       "Cake 1",
				Description: "Desc 1",
				Rating:      1,
				Image:       "https://dummyimage.com/600x400/000/fff",
				CreatedAt:   "2021-01-01 00:00:00",
				UpdatedAt:   "2021-01-01 00:00:00",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &cakeUsecase{
				cakeRepo: tt.fields.cakeRepo(ctrl),
			}
			got, err := c.UpdateCake(tt.args.cake)
			if (err != nil) != tt.wantErr {
				t.Errorf("cakeUsecase.UpdateCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cakeUsecase.UpdateCake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cakeUsecase_DeleteCake(t *testing.T) {
	type fields struct {
		cakeRepo func(ctrl *gomock.Controller) repository.CakeRepoItf
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
			name: "success delete cake",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().DeleteCake(1).Return(nil)
					return mockCakeRepo
				},
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "error delete not found",
			fields: fields{
				cakeRepo: func(ctrl *gomock.Controller) repository.CakeRepoItf {
					mockCakeRepo := mocks.NewMockCakeRepoItf(ctrl)
					mockCakeRepo.EXPECT().DeleteCake(1).Return(sql.ErrNoRows)
					return mockCakeRepo
				},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &cakeUsecase{
				cakeRepo: tt.fields.cakeRepo(ctrl),
			}
			if err := c.DeleteCake(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("cakeUsecase.DeleteCake() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
