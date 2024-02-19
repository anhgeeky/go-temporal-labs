# Bank Transfer

## Workflow

### Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)
3. Kiểm tra số dư (`Check balance`)
4. Tạo lệnh chuyển tiền (`Create bank transfer`)

## APIs

- 1. Lấy DS giao dịch chuyển khoản: GET `/bank-transfer/transfers`
- 2. Kiểm tra số dư: GET `/accounts/:ID/balance`

## Quickstart

```bash
go run worker/main.go
go run api/main.go
```

## FAQ

- <https://github.com/temporalio/temporal-ecommerce>