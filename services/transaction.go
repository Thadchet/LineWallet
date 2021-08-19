package services

import (
	"line-wallet/config"
	"line-wallet/models"
	"line-wallet/repository"
	"line-wallet/utils"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionService struct {
	Conf           config.Config
	Repo           repository.Repository
	linebotService utils.ILineService
}

//go:generate mockgen -destination=../mocks/services/mock_transaction.go -source=transaction.go

type ITransactionService interface {
	Ping() string
	AddTransaction(req models.AddTransactionRequest, member *models.Member) error
	GetTreansactions(line_user_id string) ([]models.Transaction, error)
	GetTransactionByID(ID string) (*models.Transaction, error)
	EditTransactionByID(req models.AddTransactionRequest, id string, member models.Member) (*models.Transaction, error)
	AddIncome(req models.Income, member models.Member) error
	SummaryCurrentMonth(replyToken, line_user_id string)
	GetIncomeByID(ID string) (*models.Income, error)
	EditIncomeByID(req models.Income, id string, member models.Member) (*models.Income, error)
}

func (t TransactionService) Ping() string {
	if err := t.Repo.Transaction.Insert(); err != nil {
		return err.Error()
	}
	return "Pong"
}

func (t TransactionService) calculateTotalTxnCurrentMonth(line_user_id string) float64 {

	res, err := t.Repo.Transaction.FilterTransactionCurrentMonth(line_user_id)
	if err != nil {
		return 0.0
	}
	total := 0.0
	for _, txn := range res {
		amount, _ := strconv.ParseFloat(txn.Amount, 64)
		total += amount
	}
	return total
}

func (t TransactionService) AddTransaction(req models.AddTransactionRequest, member *models.Member) error {
	txnID := primitive.NewObjectID()
	transaction := models.Transaction{
		ID:         txnID,
		Amount:     req.Amount,
		Category:   req.Category,
		Memo:       req.Memo,
		LineUserId: member.LineUserID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Type:       "txn",
	}
	if err := t.Repo.Transaction.InsertTransaction(transaction); err != nil {
		return err
	}

	amount, _ := strconv.ParseFloat(req.Amount, 64)
	remaining := member.GetRemaining() - amount
	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, remaining); err != nil {
		return err
	}
	member.UpdateRemaining(req.Amount)

	totalTxn := t.calculateTotalTxnCurrentMonth(member.LineUserID)

	flexMessage := utils.TransactionCompleteFlex(req.Amount, req.Category, req.Memo, totalTxn, member.Remaining)
	_, err2 := t.linebotService.PushMessage(member.LineUserID, flexMessage)
	if err2 != nil {
		log.Println(err2.Error())
	}
	return nil
}

func (t TransactionService) GetTreansactions(line_user_id string) ([]models.Transaction, error) {
	transactions, err := t.Repo.Transaction.GetTransactions(line_user_id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t TransactionService) GetTransactionByID(ID string) (*models.Transaction, error) {
	transaction, err := t.Repo.Transaction.GetTransactionByID(ID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t TransactionService) EditTransactionByID(req models.AddTransactionRequest, id string, member models.Member) (*models.Transaction, error) {
	_, err := t.Repo.Transaction.UpdateTransactionByID(req.Amount, req.Category, req.Memo, id)
	if err != nil {
		return nil, err
	}

	totalTxn := t.calculateTotalTxnCurrentMonth(member.LineUserID)
	// Re calulate Remaining
	totalIncome := 0.0
	incomes, err := t.Repo.Transaction.FilterIncomeCurrentMonth(member.LineUserID)
	if err != nil {
		log.Printf("err Get income : %v", err.Error())
	}
	for _, income := range incomes {
		amount, _ := strconv.ParseFloat(income.Amount, 64)
		totalIncome += amount
	}
	remaining := totalIncome - totalTxn
	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, remaining); err != nil {
		log.Printf("UpdateRemainingBalance error %v", err.Error())
		return nil, err
	}

	flexMessage := utils.TransactionCompleteFlex(req.Amount, req.Category, req.Memo, totalTxn, remaining)
	_, err2 := t.linebotService.PushMessage(member.LineUserID, flexMessage)
	if err2 != nil {
		log.Println(err.Error())
	}

	return nil, nil
}

func (t TransactionService) AddIncome(req models.Income, member models.Member) error {
	txnID := primitive.NewObjectID()
	transaction := models.Income{
		ID:         txnID,
		Amount:     req.Amount,
		Month:      req.Month,
		Memo:       req.Memo,
		LineUserId: member.LineUserID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Type:       "income",
	}
	if err := t.Repo.Transaction.InsertTransaction(transaction); err != nil {
		return err
	}

	amount, _ := strconv.ParseFloat(req.Amount, 64)

	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, amount); err != nil {
		return err
	}

	incomes, err := t.Repo.Transaction.FilterIncomeCurrentMonth(member.LineUserID)
	if err != nil {
		log.Printf("FilterIncomeCurrentMonth error : %v", err.Error())
	}
	totalIncome := 0.0
	for _, income := range incomes {
		amount, _ := strconv.ParseFloat(income.Amount, 64)
		totalIncome += amount
	}

	flexMessage := utils.IncomeCompleteFlex(req.Amount, req.Month, req.Memo, totalIncome)
	res, err := t.linebotService.PushMessage(member.LineUserID, flexMessage)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("PushMessage success : %v", res.RequestID)

	return nil
}

func (t TransactionService) SummaryCurrentMonth(replyToken, line_user_id string) {

	txns, err := t.Repo.Transaction.FilterTransactionCurrentMonth(line_user_id)
	if err != nil {
		log.Println(err.Error())
	}

	total := 0.0
	for _, txn := range txns {
		amount, _ := strconv.ParseFloat(txn.Amount, 64)
		total += amount
	}

	incomes, err := t.Repo.Transaction.FilterIncomeCurrentMonth(line_user_id)
	if err != nil {
		log.Println(err.Error())
	}

	totalIncome := 0.0
	for _, income := range incomes {
		amount, _ := strconv.ParseFloat(income.Amount, 64)
		totalIncome += amount
	}
	remaining := totalIncome - total

	flexmessage := utils.SummaryCurrentMonth(txns, incomes, total, remaining)
	res, err := t.linebotService.ReplyMessage(replyToken, flexmessage)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Reply message done : %v", res.RequestID)
}

func (t TransactionService) GetIncomeByID(ID string) (*models.Income, error) {
	income, err := t.Repo.Transaction.GetIncomeByID(ID)
	if err != nil {
		return nil, err
	}
	return income, nil
}

func (t TransactionService) EditIncomeByID(req models.Income, id string, member models.Member) (*models.Income, error) {
	_, err := t.Repo.Transaction.UpdateIncomeByID(req.Amount, req.Month, req.Memo, id)
	if err != nil {
		return nil, err
	}

	totalTxn := t.calculateTotalTxnCurrentMonth(member.LineUserID)
	// Re calulate Remaining
	totalIncome := 0.0
	incomes, err := t.Repo.Transaction.FilterIncomeCurrentMonth(member.LineUserID)
	if err != nil {
		log.Printf("err Get income : %v", err.Error())
	}
	for _, income := range incomes {
		amount, _ := strconv.ParseFloat(income.Amount, 64)
		totalIncome += amount
	}
	remaining := totalIncome - totalTxn
	if err := t.Repo.Member.UpdateRemainingBalance(member.LineUserID, remaining); err != nil {
		log.Printf("UpdateRemainingBalance error %v", err.Error())
		return nil, err
	}

	flexMessage := utils.IncomeCompleteFlex(req.Amount, req.Month, req.Memo, totalIncome)
	_, err2 := t.linebotService.PushMessage(member.LineUserID, flexMessage)
	if err2 != nil {
		log.Println(err.Error())
	}

	return nil, nil
}
