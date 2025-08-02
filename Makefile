<<<<<<< HEAD
.PHONY: create-keypair migrate-create migrate-up migrate-down migrate-force

# Các biến mặc định
ACCTPATH := account
MPATH := account/migrations      # Dùng đường dẫn tương đối để tránh lỗi migrate trên Windows
PORT ?= 5432
DB_NAME ?= postgres
DB_USER ?= postgres
DB_PASS ?= password
DB_HOST ?= postgres-account
N ?= 1
URL ?= postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(PORT)/$(DB_NAME)?sslmode=disable

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
	migrate -source file://$(MPATH) -database "$(URL)" up $(N)

# Chạy migrate down
migrate-down:
	migrate -source file://$(MPATH) -database "$(URL)" down $(N)

# Ép trạng thái migration
migrate-force:
	migrate -source file://$(MPATH) -database "$(URL)" force $(VERSION)
=======
.PHONY: keypair migrate-create migrate-up migrate-down migrate-force init

PWD = $(shell pwd)
ACCTPATH = $(PWD)/account
PORT = 5432

# Default number of migrations to execute up or down
N = 1

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(PWD)/$(APPPATH)/migrations -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(PWD)/$(APPPATH)/migrations -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(PWD)/$(APPPATH)/migrations -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(PWD)/$(APPPATH)/migrations -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)


# create dev and test keys
# run postgres containers in docker-compose 
# migrate down 
# migrate up
# docker-compose dwon
init: 
	docker-compose up -d postgres-account && \
	$(MAKE) create-keypair ENV=dev && \
	$(MAKE) create-keypair ENV=test && \
	$(MAKE) migrate-down APPPATH=account N= && \
	$(MAKE) migrate-up APPPATH=account N= && \
	docker-compose down
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
