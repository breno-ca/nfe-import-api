package service

import (
	"context"
	"database/sql"
	"desafio-tecnico-backend/pkg/database"
	"desafio-tecnico-backend/pkg/entity"
	"errors"
	"log"
)

type APIServiceInterface interface {
	UserServiceInterface
	ProdutoServiceInterface
	NFeImportServiceInterface
}

type APIService struct {
	dbp database.DatabaseInterface
}

func NewAPIService(database_pool database.DatabaseInterface) *APIService {
	return &APIService{
		database_pool,
	}
}

type UserServiceInterface interface {
	// CreateUser(user *entity.User, ctx context.Context) (uint64, error)
	Login(user *entity.User, ctx context.Context) (string, error)
}

type ProdutoServiceInterface interface {
	CreateProduto(produto *entity.Prod, ctx context.Context) error
}

type NFeImportServiceInterface interface {
	ImportNFeXML(nfe *entity.NFe, ctx context.Context) error
}

func (s *APIService) Login(user *entity.User, ctx context.Context) (string, error) {
	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare("SELECT cnpj, senha FROM empresa WHERE cnpj = ?")
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error preparing statement")
	}

	defer stmt.Close()

	hash := ""
	err = stmt.QueryRow(user.CNPJ).Scan(&user.CNPJ, &hash)
	if err != nil {
		log.Println(err.Error())
		return "", nil
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return hash, nil

}

func (s *APIService) CreateProduto(produto *entity.Prod, ctx context.Context) error {
	database := s.dbp.GetDB()

	if produto == nil {
		return errors.New("Produto inválido")
	}

	if produto.CEAN == "" {
		return errors.New("O campo cEAN é obrigatório")
	}

	_, err := database.ExecContext(ctx, "INSERT INTO produtos (cEAN, vCusto, vPreco, qCom) VALUES (?, ?, ?, ?)", produto.CEAN, produto.VCusto, produto.VPreco, produto.QCom)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIService) ImportNFeXML(nfe *entity.NFe, ctx context.Context) error {
	for _, prod := range nfe.Prod {
		err := s.UpsertProduto(&prod, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *APIService) UpsertProduto(produto *entity.Prod, ctx context.Context) error {
	database := s.dbp.GetDB()

	if produto == nil {
		return errors.New("Produto inválido")
	}

	if produto.CEAN == "" {
		return errors.New("O campo CEAN é obrigatório")
	}

	var existingVCusto, existingVPreco, existingQCom float64

	err := database.QueryRowContext(ctx, "SELECT vCusto, vPreco, qCom FROM produtos WHERE cEAN = ?", produto.CEAN).Scan(&existingVCusto, &existingVPreco, &existingQCom)
	if err != nil {
		if err == sql.ErrNoRows {
			return s.CreateProduto(produto, ctx)
		}
		return err
	}

	newVCusto := existingVCusto + produto.VCusto
	newVPreco := existingVPreco + produto.VPreco
	newQCom := existingQCom + produto.QCom

	_, err = database.ExecContext(ctx, "UPDATE produtos SET vCusto = ?, vPreco = ?, qCom = ? WHERE cEAN = ?", newVCusto, newVPreco, newQCom, produto.CEAN)
	if err != nil {
		return err
	}
	return nil
}
