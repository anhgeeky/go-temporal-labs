# Temporal Labs

## To-do

- [x] Build cấu trúc cho Temporal (Api + Worker)
- [x] Xây dựng luồng **Thông báo** (Notification Flow)
- [ ] Xây dựng luồng **Chuyển tiền** (Transfer Flow)
  - [ ] Bổ sung require before steps (`When step 2.4, 2.5 done -> Completed`)
- [ ] Tái cấu trúc project for Temporal
  - [x] Chia nhỏ submodule cho `workflow`, `api`
  - [ ] Bổ sung thêm các features chung cho `workflow` core
- [ ] Saga for Temporal
  - [x] Saga sample with `REST Api`
  - [ ] Saga sample with `Kafka Event Driven`
- [ ] Add or Remove 1 activity
  - Follow: <https://community.temporal.io/t/update-activity-and-or-workflow-inputs/4972/5>
  - Temporal chỉ chạy từng activity, có `STOP` cluster, khi chạy lại vẫn còn `Running` thì sẽ chạy lại
  - Nếu có add or remove 1 activity thì sẽ load lại các activity đã update (add, remove, update) -> Chạy tiếp tục

## Quickstart

```bash
go run ./pkg/banktransfer/cmd/worker/main.go
go run ./pkg/notification/cmd/worker/main.go
go run ./serivces/mcs-account/main.go
go run ./serivces/mcs-money-transfer/main.go
go run ./serivces/mcs-notification/main.go
# or 
sh start-all.sh
sh start-worker.sh
sh start-api.sh
```

## Saga (Temporal + Kafka + Microservices)

![Screenshot](/docs/assets/saga-workflows-sample.png)

## Host APIs

- `mcs-account`: `localhost:6001`
- `mcs-money-transfer`: `localhost:6002`
- `mcs-notification`: `localhost:6003`

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
  - 2.1. Kiểm tra số dư (`CheckBalance`) (`Synchronize`)
  - 2.2. Kiểm tra tra tài khoản đích (`CheckTargetAccount`) (`Synchronize`)
  - 2.3. Tạo giao dịch chuyển tiền (`CreateTransferTransaction`) (`When step 2.1, 2.2 done -> Continue`)
  - 2.4. Tạo giao dịch ghi nợ (`WriteCreditAccount`) (`Synchronize`)
  - 2.5. Tạo giao dịch ghi có (`WriteDebitAccount`) (`Synchronize`)
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

- [x] Lấy DS giao dịch chuyển khoản: GET `/transfers/:workflowID`
- [x] Kiểm tra số dư: GET `/accounts/:workflowID/balance`
- [ ] Kiểm tra tra tài khoản đích (`CheckTargetAccount`)
- [x] Tạo giao dịch chuyển tiền (`CreateTransferTransaction`): POST `/transfers/:workflowID/transactions`
- [x] Tạo giao dịch ghi nợ (`WriteCreditAccount`): POST `/transfers/:workflowID/credit-accounts`
- [x] Tạo giao dịch ghi có (`WriteDebitAccount`): POST `/transfers/:workflowID/debit-accounts`
- [x] Add new activity for test: POST `/transfers/:workflowID/new-activity`
- [x] [Rollback] Tạo giao dịch chuyển tiền (`CreateTransferTransactionCompensation`): POST `/transfers/:workflowID/transactions/rollback`
- [x] [Rollback] Tạo giao dịch ghi nợ (`WriteCreditAccountCompensation`): POST `/transfers/:workflowID/credit-accounts/rollback`
- [x] [Rollback] Tạo giao dịch ghi có (`WriteDebitAccountCompensation`): POST `/transfers/:workflowID/debit-accounts/rollback`
- [x] [Rollback] Add new activity for test: POST `/transfers/:workflowID/new-activity/rollback`

## Saga

![Screenshot](/docs/assets/bank-transfer-saga-pattern-log.png)

## Temporal screenshot

![Screenshot](/docs/assets/bank-transfer-workflows.jpg)
![Screenshot](/docs/assets/bank-transfer-temporal-admin-log.png)
![Screenshot](/docs/assets/bank-transfer-sub-workflow-temporal-admin-log.png)

## Stack

- `fiber`: <https://github.com/gofiber/fiber>
- `temporal`: <https://github.com/temporalio/temporal>
- `viper`: <https://github.com/spf13/viper>
- `kafka-go`: <https://github.com/segmentio/kafka-go>

## FAQ

- Temporal for Docker: <https://github.com/temporalio/docker-compose>
- <https://github.com/temporalio/samples-go>
- <https://github.com/temporalio/temporal-ecommerce>
