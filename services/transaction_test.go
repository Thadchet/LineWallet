package services

import (
	mock_repository "line-wallet/mocks/repository"
	"line-wallet/models"
	"line-wallet/repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTransactionService_calculateTotalTxnCurrentMonth(t *testing.T) {

	type args struct {
		line_user_id string
	}
	tests := []struct {
		name            string
		args            args
		transactionMock []models.Transaction
		want            float64
	}{
		{
			name: "Test case 1",
			args: args{
				line_user_id: "ddffd",
			},
			transactionMock: []models.Transaction{
				{
					Amount: "20.0",
				},
				{
					Amount: "40.0",
				},
			},
			want: 60.0,
		},
		{
			name: "Test case 2",
			args: args{
				line_user_id: "ddffd",
			},
			transactionMock: []models.Transaction{
				{
					Amount: "10.0",
				},
				{
					Amount: "40.0",
				},
			},
			want: 50.0,
		},
		{
			name: "Test case 3",
			args: args{
				line_user_id: "ddffd",
			},
			transactionMock: []models.Transaction{
				{
					Amount: "40.0",
				},
				{
					Amount: "40.0",
				},
			},
			want: 80.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockITransactionRepo(ctrl)
			repo.EXPECT().FilterTransactionCurrentMonth(gomock.Any()).Return(tt.transactionMock, nil).AnyTimes()

			tr := TransactionService{
				Repo: repository.Repository{
					Transaction: repo,
				},
			}

			if got := tr.calculateTotalTxnCurrentMonth(tt.args.line_user_id); got != tt.want {
				t.Errorf("TransactionService.calculateTotalTxnCurrentMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
