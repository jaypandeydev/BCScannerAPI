package usecases

import (
	"BCScanner/domain"
	"context"
	"errors"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type QRGenerateUseCase struct {
	QRGenerateRepository domain.QRGenerateRepository
	contextTimeout       time.Duration
}

func NewQRGenerateUseCase(QRGenerateRepository domain.QRGenerateRepository, timeout time.Duration) domain.QRGenerateUseCase {
	return &QRGenerateUseCase{
		QRGenerateRepository: QRGenerateRepository,
		contextTimeout:       timeout,
	}
}

func (QR *QRGenerateUseCase) CreateQR(ctx context.Context, QRGenerate *domain.QRGenerateRequest) (string, error) {

	var InsertedID string
	// Generate QR code
	QrCode, err := qrcode.New(QRGenerate.OwnerName, qrcode.Medium)
	if err != nil {
		return InsertedID, errors.New("failed to generate QR code")
	}
	ctx, cancel := context.WithTimeout(ctx, QR.contextTimeout)
	defer cancel()

	// Ensure the directory exists
	dir := "content/qr_codes"
	err = ensureDir(dir)
	if err != nil {
		return InsertedID, errors.New("failed to create directory for QR codes")
	}

	// Save QR code to file
	//filePath := filepath.Join("qr_codes", QRGenerate.OwnerName+".png")
	filePath := filepath.Join("./content/qr_codes", strings.ReplaceAll(QRGenerate.OwnerName, `"`, "")+".png")
	//filePath1 := strings.ReplaceAll(filePath, `\"`, `"`)
	err = QrCode.WriteFile(256, filePath)
	if err != nil {
		return InsertedID, errors.New("failed to save QR code to file")
	}
	QRGenerate.QRImageURL = filePath
	InsertedID, err = QR.QRGenerateRepository.CreateQR(ctx, QRGenerate)
	return InsertedID, err
}

func (QR *QRGenerateUseCase) Fetch(ctx context.Context) ([]bson.D, error) {
	ctx, cancel := context.WithTimeout(ctx, QR.contextTimeout)
	defer cancel()
	return QR.QRGenerateRepository.Fetch(ctx)
}

func (QR *QRGenerateUseCase) GetByID(ctx context.Context, Id string) (domain.QRGenerateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, QR.contextTimeout)
	defer cancel()
	return QR.QRGenerateRepository.GetByID(ctx, Id)
}
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
