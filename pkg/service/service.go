package service

import (
	"context"
	"desafio-tecnico-backend/pkg/database"
	"desafio-tecnico-backend/pkg/entity"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type APIServiceInterface interface {
	UserServiceInterface
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
	Login(user *entity.User, ctx context.Context) (string, error)
}

type NFeImportServiceInterface interface {
	ImportNFeXML(xmlData []byte, ctx context.Context, userCNPJ string) error
}

// Função responsável pelo Login do usuário
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

// Verifica se há destinatário com o CNPJ informado no banco de dados
func (s *APIService) HasDest(cnpj string, ctx context.Context) (bool, error) {

	database := s.dbp.GetDB()

	sql := "SELECT COUNT(*) FROM destinatario WHERE cnpj = ?"

	var count int
	err := database.QueryRow(sql, cnpj).Scan(&count)
	if err != nil {
		log.Println("Erro ao verificar o destinatário:", err)
		return false, err
	}

	return count > 0, nil
}

// Realiza a inserção de um destinatário no banco de dados
func (s *APIService) InsertDest(destinatario *entity.Dest, ctx context.Context) error {

	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	sql := `INSERT IGNORE INTO destinatario (cnpj, xNome, email) VALUES (?, ?, ?)`
	_, err = tx.ExecContext(ctx, sql, destinatario.CNPJ, destinatario.XNome, "")
	if err != nil {
		log.Println(err, "error on insert")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Realiza a inserção do endereço de um destinatário presente no banco de dados
func (s *APIService) InsertEnderDest(CNPJ string, endereco *entity.EnderDest, ctx context.Context) error {

	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var idDest int
	sql := "SELECT destinatario_id FROM destinatario WHERE cnpj = ?"

	err = tx.QueryRow(sql, CNPJ).Scan(&idDest)
	if err != nil {
		return err
	}

	sql = `INSERT INTO endereco (destinatario_id, xLgr, nro, xCpl, xBairro, cMun, CEP, fone) VALUES (?, ?, ?, ?, ?, ? ,?, ?)`

	_, err = tx.ExecContext(ctx, sql, idDest, endereco.XLgr, endereco.Nro, endereco.XCpl, endereco.XBairro, endereco.CMun, endereco.CEP, "")
	if err != nil {
		log.Println(err, "error on insert")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Insere ou Atualiza o produto se existente
func (s *APIService) UpsertProduto(produtos []*entity.Prod, ctx context.Context) error {

	if len(produtos) == 0 {
		return errors.New("erro: lista de produtos vazia")
	}

	database := s.dbp.GetDB()

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var sql string

	var exists bool
	for _, produto := range produtos {

		err = tx.QueryRowContext(ctx, "SELECT 1 FROM produto WHERE cEAN = ?", produto.CEAN).Scan(&exists)
		if err != nil {
			log.Println(err, "error on check if exists")

		}

		if exists {

			sql = `UPDATE produto SET vCusto = ? WHERE cEAN = ?`

			result, err := tx.ExecContext(ctx, sql, (produto.VFrete + produto.VSeg + produto.VDesc + produto.VOutro), produto.CEAN)
			if err != nil {
				log.Println(err, "error on update")

			}

			_, err = result.RowsAffected()
			if err != nil {
				return errors.New("error rowAffected insert into database")
			}

		} else {

			sql = `INSERT INTO produto (cProd, cEAN, xProd, uCom, qCom, vUnCom, vCusto) VALUES (?, ?, ?, ?, ?, ?, ?)`

			result, err := tx.ExecContext(ctx, sql, produto.CProd, produto.CEAN, produto.XProd, produto.UCom, produto.QCom, produto.VUnCom, (produto.VFrete + produto.VSeg + produto.VDesc + produto.VOutro))
			if err != nil {
				log.Println(err, "error on insert")
			}

			_, err = result.RowsAffected()
			if err != nil {
				return errors.New("error rowAffected insert into database")
			}

		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Realiza a importação da NFe passada no corpo de uma requisição para o banco de dados
func (s *APIService) ImportNFeXML(xmlData []byte, ctx context.Context, userCNPJ string) error {

	var nfe entity.NfeProc

	if err := xml.Unmarshal(xmlData, &nfe); err != nil {
		fmt.Println("Erro ao decodificar XML:", err)
		return err
	}

	nfeCNPJ := nfe.NFe.InfNFe.Emit.CNPJ
	destCNPJ := nfe.NFe.InfNFe.Dest.CNPJ
	produtos := nfe.NFe.InfNFe.Det.Prod
	destinatario := nfe.NFe.InfNFe.Dest

	if nfeCNPJ != userCNPJ {
		return errors.New("nfe nao pertence a empresa logada")
	}

	existeDestinatario, err := s.HasDest(destCNPJ, ctx)
	if err != nil {
		log.Println("error on hasdest")
	}

	if !existeDestinatario {
		s.InsertDest(&destinatario, ctx)
		s.InsertEnderDest(destCNPJ, &destinatario.EnderDest, ctx)
	}

	s.UpsertProduto(produtos, ctx)

	return nil

}
