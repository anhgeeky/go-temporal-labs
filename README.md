# Temporal Labs

## To-do

- [x] Build cấu trúc cho Temporal (Api + Worker)
- [x] Xây dựng luồng **Thông báo** (Notification Flow)
- [ ] Xây dựng luồng **Chuyển tiền** (Transfer Flow)
  - [ ] Bổ sung require before steps (`When step 2.4, 2.5 done -> Completed`)
- [ ] Saga for Temporal
- [ ] Add or Remove 1 activity
  - Follow: <https://community.temporal.io/t/update-activity-and-or-workflow-inputs/4972/5>

## Quickstart

```bash
go run ./pkg/banktransfer/cmd/worker/main.go
go run ./pkg/notification/cmd/worker/main.go
go run ./serivces/mcs-account/main.go
go run ./serivces/mcs-money-transfer/main.go
go run ./serivces/mcs-notification/main.go
# or 
sh start-worker.sh
sh start-api.sh
```

## Saga (Temporal + Kafka + Microservices)

![Screenshot](/docs/assets/saga-workflows-sample.png)

## Host APIs

- `mcs-account`: `localhost:3001`
- `mcs-money-transfer`: `localhost:3002`
- `mcs-notification`: `localhost:3003`

## Workers

- `banktransfer`
- `notification`
- `onboarding`

## Bank Transfer Workflow (Implement Saga)

### Before: Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)

### Start: Chuyển tiền

1. [**Transfer Flow**] Tạo lệnh YC chuyển tiền (`Create bank transfer`) (`Start`)
  - Run [**Notification Flow**] send OTP
2. [**Transfer Flow**] Xác thực OTP (`Trigger Signal`)
  - 2.1. Kiểm tra số dư (`Check balance account`) (`Synchronize`)
  - 2.2. Kiểm tra tra tài khoản đích (`Check target account`) (`Synchronize`)
  - 2.3. Tạo giao dịch chuyển tiền (`Create new transaction`) (`When step 2.1, 2.2 done -> Continue`)
  - 2.4. Tạo giao dịch ghi nợ (`Synchronize`)
  - 2.5. Tạo giao dịch ghi có (`Synchronize`)
  - 2.6. Transfer done  (`When step 2.4, 2.5 done -> Completed`) (`Trigger [Notification Flow]`)
  - 2.7. Call subflow [**Notification Flow**] Gửi thông báo đã chuyển tiền
    - 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
    - 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công` (`Parallel`)
    - 2.7.3 Push message notification vào `firebase` (`Parallel`)
    - 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công` (`Parallel`)

### End: Chuyển tiền

1. Nhận message internal app
2. Lấy thông tin kết quả chuyển tiền
3. Reload lại show kết quả `Chuyển tiền Thành công`

## APIs

- [x] Lấy DS giao dịch chuyển khoản: GET `/transfers`
- [ ] Kiểm tra số dư: GET `/accounts/:ID/balance`

## Temporal screenshot

![Screenshot](/docs/assets/bank-transfer-workflows.jpg)
![Screenshot](/docs/assets/bank-transfer-temporal-admin-log.png)
![Screenshot](/docs/assets/bank-transfer-sub-workflow-temporal-admin-log.png)

## Stack for Sample

- `fiber`
- `temporal`
- `viper`

## FAQ

- Temporal for Docker: <https://github.com/temporalio/docker-compose>
- <https://github.com/temporalio/samples-go>
- <https://github.com/temporalio/temporal-ecommerce>
