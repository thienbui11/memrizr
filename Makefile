.PHONY: create-keypair migrate-create migrate-up migrate-down migrate-force

# Các biến mặc định
ACCTPATH := account
MPATH := account/migrations      # Dùng đường dẫn tương đối để tránh lỗi migrate trên Windows
PORT := 5432
DB_NAME ?= postgres              # Cho phép override khi gọi make
DB_USER ?= postgres
DB_PASS ?= password
N ?= 1

# Tạo khóa RSA cho môi trường (ENV = dev/test)
create-keypair:
	@echo Creating an RSA 2048-bit key pair for $(ENV)
	mkdir -p $(ACCTPATH)
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

# Tạo migration mới (NAME = tên migration)
migrate-create:
	@echo Creating new migration file: $(NAME)
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

# Chạy migrate up
migrate-up:
	migrate -source file://$(MPATH) -database "postgres://postgres:password@localhost:5432/mindstack?sslmode=disable" up $(N)

# Chạy migrate down
migrate-down:
	migrate -source file://$(MPATH) -database "postgres://postgres:password@localhost:5432/mindstack?sslmode=disable" down $(N)

# Ép trạng thái migration
migrate-force:
	migrate -source file://$(MPATH) -database "postgres://postgres:password@localhost:5432/mindstack?sslmode=disable" force $(VERSION)
