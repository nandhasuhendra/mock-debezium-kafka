package processor

import (
	"encoding/json"
	"log"
	"time"
)

type stringJSON struct {
	Payload string
}

type TransactionMessage struct {
	CompanyID string `json:"company_id"`
	Data      struct {
		ID                          int           `json:"id"`
		TransactionTypeID           int           `json:"transaction_type_id"`
		CompanyID                   int           `json:"company_id"`
		TransactionDate             string        `json:"transaction_date"`
		TransactionStatusID         int           `json:"transaction_status_id"`
		TransactionNo               string        `json:"transaction_no"`
		PersonID                    int           `json:"person_id"`
		Email                       string        `json:"email"`
		Address                     string        `json:"address"`
		ShippingAddress             interface{}   `json:"shipping_address"`
		Message                     string        `json:"message"`
		Memo                        string        `json:"memo"`
		DiscountTypeID              int           `json:"discount_type_id"`
		DiscountUnit                string        `json:"discount_unit"`
		DiscountPrice               string        `json:"discount_price"`
		ShippingPrice               string        `json:"shipping_price"`
		Deposit                     string        `json:"deposit"`
		AmountReceive               string        `json:"amount_receive"`
		Remaining                   string        `json:"remaining"`
		OriginalAmount              string        `json:"original_amount"`
		TermID                      int           `json:"term_id"`
		DueDate                     string        `json:"due_date"`
		ShipVia                     interface{}   `json:"ship_via"`
		ShippingDate                interface{}   `json:"shipping_date"`
		TrackingNo                  interface{}   `json:"tracking_no"`
		RefundFromID                interface{}   `json:"refund_from_id"`
		DepositToID                 interface{}   `json:"deposit_to_id"`
		ReferenceNo                 string        `json:"reference_no"`
		PaymentMethodID             interface{}   `json:"payment_method_id"`
		DeletedAt                   interface{}   `json:"deleted_at"`
		CreatedAt                   time.Time     `json:"created_at"`
		UpdatedAt                   time.Time     `json:"updated_at"`
		TaxAfterDiscount            bool          `json:"tax_after_discount"`
		TaxID                       interface{}   `json:"tax_id"`
		TaxAmount                   string        `json:"tax_amount"`
		Token                       string        `json:"token"`
		IsOpeningBalance            interface{}   `json:"is_opening_balance"`
		AccountReconcileID          interface{}   `json:"account_reconcile_id"`
		InvoiceID                   interface{}   `json:"invoice_id"`
		ReturnAmount                interface{}   `json:"return_amount"`
		CreditMemoInvoiceID         interface{}   `json:"credit_memo_invoice_id"`
		SelectedCreditMemoID        interface{}   `json:"selected_credit_memo_id"`
		IsImport                    bool          `json:"is_import"`
		SelectedPoID                interface{}   `json:"selected_po_id"`
		AmountReceivedFromPo        string        `json:"amount_received_from_po"`
		WitholdingAccountID         interface{}   `json:"witholding_account_id"`
		WitholdingValue             string        `json:"witholding_value"`
		WarehouseID                 interface{}   `json:"warehouse_id"`
		IsShipped                   bool          `json:"is_shipped"`
		IsShippingAddressAsBilling  interface{}   `json:"is_shipping_address_as_billing"`
		ClosingTheBookID            interface{}   `json:"closing_the_book_id"`
		IsIncomeSummaryTransaction  interface{}   `json:"is_income_summary_transaction"`
		IsTaxAdjustmentTransaction  interface{}   `json:"is_tax_adjustment_transaction"`
		SelectedPqID                interface{}   `json:"selected_pq_id"`
		ExpiryDate                  interface{}   `json:"expiry_date"`
		UseExpired                  bool          `json:"use_expired"`
		AssetID                     interface{}   `json:"asset_id"`
		DoNotValidateThis           bool          `json:"do_not_validate_this"`
		TransferAmount              string        `json:"transfer_amount"`
		Status                      string        `json:"status"`
		MultiCurrencyID             interface{}   `json:"multi_currency_id"`
		McRate                      interface{}   `json:"mc_rate"`
		DepositID                   interface{}   `json:"deposit_id"`
		LastMultiCurrencyID         interface{}   `json:"last_multi_currency_id"`
		CurrencyListID              interface{}   `json:"currency_list_id"`
		WitholdingType              string        `json:"witholding_type"`
		DeliveryInvoiceID           interface{}   `json:"delivery_invoice_id"`
		DiscountInputRaw            interface{}   `json:"discount_input_raw"`
		Source                      string        `json:"source"`
		ExpensePayable              bool          `json:"expense_payable"`
		TemporaryTransactionID      interface{}   `json:"temporary_transaction_id"`
		IsAutomatic                 bool          `json:"is_automatic"`
		IsPaid                      interface{}   `json:"is_paid"`
		IsPaidFromMoka              interface{}   `json:"is_paid_from_moka"`
		RealisedGainID              interface{}   `json:"realised_gain_id"`
		Checksum                    interface{}   `json:"checksum"`
		CustomID                    interface{}   `json:"custom_id"`
		ConversionBalanceID         interface{}   `json:"conversion_balance_id"`
		ReverseID                   interface{}   `json:"reverse_id"`
		IsCreateBeforeConversion    interface{}   `json:"is_create_before_conversion"`
		ComparativeBalanceID        interface{}   `json:"comparative_balance_id"`
		IsReversal                  bool          `json:"is_reversal"`
		UseTaxInclusive             bool          `json:"use_tax_inclusive"`
		ImportID                    interface{}   `json:"import_id"`
		UniqueTransactionNo         string        `json:"unique_transaction_no"`
		UserID                      int           `json:"user_id"`
		PayMethod                   interface{}   `json:"pay_method"`
		TagIds                      []interface{} `json:"tag_ids"`
		EventType                   string        `json:"event_type"`
		WebhookEvents               []interface{} `json:"webhook_events"`
		TransactionStatusName       string        `json:"transaction_status_name"`
		TransactionStatusNameBahasa string        `json:"transaction_status_name_bahasa"`
		AtDesc                      any           `json:"at_desc"`
		Payments                    []struct {
			ID                  int       `json:"id"`
			TransactionTypeID   int       `json:"transaction_type_id"`
			TransactionDate     string    `json:"transaction_date"`
			TransactionStatusID int       `json:"transaction_status_id"`
			TransactionNo       string    `json:"transaction_no"`
			OriginalAmount      string    `json:"original_amount"`
			CreatedAt           time.Time `json:"created_at"`
		} `json:"payments"`
		PaymentDate []string      `json:"payment_date"`
		Tags        []interface{} `json:"tags"`
		Person      struct {
			DisplayName        string      `json:"display_name"`
			Email              string      `json:"email"`
			Phone              interface{} `json:"phone"`
			Mobile             interface{} `json:"mobile"`
			TaxNo              interface{} `json:"tax_no"`
			AssociateCompany   interface{} `json:"associate_company"`
			FirstName          interface{} `json:"first_name"`
			LastName           interface{} `json:"last_name"`
			MiddleName         interface{} `json:"middle_name"`
			PeopleType         string      `json:"people_type"`
			DefaultArAccountID int         `json:"default_ar_account_id"`
			DefaultApAccountID int         `json:"default_ap_account_id"`
		} `json:"person"`
		PersonDefaultApAccount struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Number     string `json:"number"`
			CategoryID int    `json:"category_id"`
		} `json:"person_default_ap_account"`
		PersonDefaultArAccount struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Number     string `json:"number"`
			CategoryID int    `json:"category_id"`
		} `json:"person_default_ar_account"`
		HasAttachment               bool          `json:"has_attachment"`
		CurrencySymbol              string        `json:"currency_symbol"`
		WitholdingAmount            any           `json:"witholding_amount"`
		TaxDetails                  []interface{} `json:"tax_details"`
		CurrencyCode                string        `json:"currency_code"`
		DiscountTypeName            string        `json:"discount_type_name"`
		TermName                    string        `json:"term_name"`
		WarehouseName               interface{}   `json:"warehouse_name"`
		RateToBase                  string        `json:"rate_to_base"`
		HasDeliveries               bool          `json:"has_deliveries"`
		MergedInvoice               interface{}   `json:"merged_invoice"`
		Proforma                    []interface{} `json:"proforma"`
		ProformaLastDeleted         interface{}   `json:"proforma_last_deleted"`
		Returns                     []interface{} `json:"returns"`
		Deliveries                  []interface{} `json:"deliveries"`
		OrderID                     interface{}   `json:"order_id"`
		AllocatedUpfrontTaxesAmount int           `json:"allocated_upfront_taxes_amount"`
		Order                       struct{}      `json:"order"`
		UpfrontTaxPerAccounts       struct{}      `json:"upfront_tax_per_accounts"`
		TransactionAccountLines     struct {
			Data []struct {
				ID             int         `json:"id"`
				TransactionID  int         `json:"transaction_id"`
				AccountID      int         `json:"account_id"`
				Description    string      `json:"description"`
				Credit         string      `json:"credit"`
				Debit          string      `json:"debit"`
				Active         interface{} `json:"active"`
				CreatedAt      time.Time   `json:"created_at"`
				UpdatedAt      time.Time   `json:"updated_at"`
				DeletedAt      interface{} `json:"deleted_at"`
				ExpenseID      interface{} `json:"expense_id"`
				RealisedGainID interface{} `json:"realised_gain_id"`
				LineTaxID      interface{} `json:"line_tax_id"`
				Idx            int         `json:"idx"`
			} `json:"data"`
		} `json:"transaction_account_lines"`
		TransactionLines struct {
			Data []struct {
				ID                     int           `json:"id"`
				TransactionID          int           `json:"transaction_id"`
				ProductID              int           `json:"product_id"`
				Description            string        `json:"description"`
				Quantity               string        `json:"quantity"`
				UnitID                 int           `json:"unit_id"`
				Rate                   string        `json:"rate"`
				Amount                 string        `json:"amount"`
				Active                 interface{}   `json:"active"`
				CreatedAt              time.Time     `json:"created_at"`
				UpdatedAt              time.Time     `json:"updated_at"`
				Tax                    interface{}   `json:"tax"`
				Discount               string        `json:"discount"`
				DeletedAt              interface{}   `json:"deleted_at"`
				CustomID               interface{}   `json:"custom_id"`
				LineTaxID              interface{}   `json:"line_tax_id"`
				ConvertRate            string        `json:"convert_rate"`
				Idx                    int           `json:"idx"`
				UnitName               string        `json:"unit_name"`
				TaxableAmount          string        `json:"taxable_amount"`
				TaxPercent             interface{}   `json:"tax_percent"`
				TaxAmount              string        `json:"tax_amount"`
				ProductName            string        `json:"product_name"`
				ProductCode            interface{}   `json:"product_code"`
				BuyAccountID           int           `json:"buy_account_id"`
				SellAccountID          int           `json:"sell_account_id"`
				SourceTransactionLines []interface{} `json:"source_transaction_lines"`
			} `json:"data"`
			TotalData   int    `json:"total_data"`
			TotalPages  int    `json:"total_pages"`
			NextPageURL string `json:"next_page_url"`
		} `json:"transaction_lines"`
		ReturnWithholdingAmount string      `json:"return_withholding_amount"`
		Warehouse               interface{} `json:"warehouse"`
		UseShipping             bool        `json:"use_shipping"`
		UseDeposit              bool        `json:"use_deposit"`
		UseDiscount             bool        `json:"use_discount"`
		UseDiscountLine         bool        `json:"use_discount_line"`
		PaymentTransactions     []struct {
			ID               int         `json:"id"`
			PaymentID        int         `json:"payment_id"`
			TransactionID    int         `json:"transaction_id"`
			Amount           string      `json:"amount"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			AmountAsCurrency interface{} `json:"amount_as_currency"`
			McRate           interface{} `json:"mc_rate"`
			DeletedAt        interface{} `json:"deleted_at"`
			RealisedGainID   interface{} `json:"realised_gain_id"`
			Transaction      struct {
				ID            int    `json:"id"`
				TransactionNo string `json:"transaction_no"`
			} `json:"transaction"`
		} `json:"payment_transactions"`
		CreateWithMultipleWithholdings bool `json:"create_with_multiple_withholdings"`
		WithholdingTransactions        []struct {
			ID              int    `json:"id"`
			Label           string `json:"label"`
			WithholdingType string `json:"withholding_type"`
			Rate            string `json:"rate"`
			Amount          string `json:"amount"`
			AccountID       int    `json:"account_id"`
		} `json:"withholding_transactions"`
	} `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type DebeziumMessage struct {
	Payload struct {
		After struct {
			ID      int64  `json:"id"`
			Payload string `json:"payload"`
		} `json:"after"`
	} `json:"payload"`
}

func ProcessMessage(msg []byte) ([]byte, error) {
	var transactionMsg TransactionMessage
	err := json.Unmarshal(msg, &transactionMsg)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return nil, err
	}

	payload, err := transactionMsg.toStringJSON()
	if err != nil {
		return nil, err
	}

	debeziumMsg := DebeziumMessage{}
	debeziumMsg.Payload.After.ID = int64(transactionMsg.Data.ID)
	debeziumMsg.Payload.After.Payload = payload

	return json.Marshal(debeziumMsg)
}

func (trxMsg *TransactionMessage) toStringJSON() (string, error) {
	payload := trxMsg.Data

	jsonString, err := json.Marshal(&payload)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
