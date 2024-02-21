### build

```sh
$ vim ent/entc.go
$ vim ent/schema/users.go
$ go generate ./...
$ go build
$ ./card

$ go generate ./...
$ PASS=`cat ./token.json|jq -r .password` TOKEN=`cat ./token.json|jq -r .token` go run -mod=mod main.go
$ curl -X POST -H "Content-Type: application/json" -d "{\"username\":\"syui\",\"password\":\"$pass\"}" localhost:8080/users
$ curl -X POST -H "Content-Type: application/json" -d "{\"owner\":1,\"password\":\"$pass\"}" localhost:8080/cards
$ curl -X POST -H "Content-Type: application/json" -d "{\"owner\":1,\"card\":1,\"cp\":11,\"status\":\"normal\",\"password\":\"$pass\"}" localhost:8080/cards
$ curl localhost:8080/users
$ curl localhost:8080/cards
$ curl localhost:8080/users/1
$ curl localhost:8080/users/1/card
```

### use

```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"username":"syui",\"password\":\"$pass\"}' https://api.syui.ai/users

# onconflict
$ !!

$ curl -sL https://api.syui.ai/users/1
```

```sh
# item select
$ curl -sL "https://api.syui.ai/users?itemsPerPage=255"
$ curl -sL "https://api.syui.ai/cards?itemsPerPage=255"
$ curl -sL "https://api.syui.ai/users/1/card?itemsPerPage=255"
```

### ref

```sh
$ vim ./ent/ogent/ogent.go
// 新規登録の停止
// CreateUsers handles POST /users-slice requests.
func (h *OgentHandler) CreateUsers(ctx context.Context, req CreateUsersReq) (CreateUsersRes, error) {
	b := h.client.Users.Create()
	//b.SetUser(req.User)
	b.SetUser("syui")
}

// 削除の無効
// DeleteUsers handles DELETE /users-slice/{id} requests.
func (h *OgentHandler) DeleteUsers(ctx context.Context, params DeleteUsersParams) (DeleteUsersRes, error) {
	if params.ID != 1 {
err := h.client.Users.DeleteOneID(params.ID).Exec(ctx)
	}
	return new(DeleteUsersNoContent), nil
}

// 要素の書き換えの禁止
// UpdateUsers handles PATCH /users-slice/{id} requests.
func (h *OgentHandler) UpdateUsers(ctx context.Context, req UpdateUsersReq, params UpdateUsersParams) (UpdateUsersRes, error) {
	b := h.client.Users.UpdateOneID(params.ID)
	// Add all fields.
	//if v, ok := req.Hp.Get(); ok {
	//	b.SetHp(v)
	//}
```

### link

- https://entgo.io/ja/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen/

- https://github.com/ariga/ogent/blob/main/example/todo/ent/entc.go

- https://github.com/ent/ent/blob/master/dialect/sql/schema/postgres_test.go

- https://github.com/go-kratos/beer-shop/tree/main/app/catalog/service/internal/data/ent


### update

```sh
$ curl --dump-header - 'https://api.syui.ai/users' -H 'Origin: https://card.syui.ai'|less
```

> ent/ogent/oas_response_encoders_gen.go

```go
func encodeCreateGroupResponse(response CreateGroupRes, w http.ResponseWriter, span trace.Span) error {
    w.Header().Set("Access-Control-Allow-Origin", "https://card.syui.ai")
        switch response := response.(type) {
            w.Header().Set("Access-Control-Allow-Origin", "https://card.syui.ai")
```

### northflank

#### backup sqlite

- `cron`, `repo(private)`, `pass(token)`

```sh
#!/bin/zsh

pass=password
/usr/bin/northflank exec service --project $project --service $service --cmd "/app/data/api/backup.sh $pass"

function f(){
    rm /app/data/api/backup.sh
    echo '#!/bin/bash
    pass=$1
    git config --global user.email syui@syui.ai
    git config --global user.name syui
    cp -rf /app/data/new.sqlite /app/data/api/latest.sqlite
    cp -rf /app/data/new.sqlite /app/data/api/`date '+%w'`.sqlite
    cd /app/data/api
    git remote add origin https://$pass@github.com/ai/api
    git add .
    git commit -m backup
    git push origin main
    git remote rm origin
    ' >> /app/data/api/backup.sh
    chmod +x /app/data/api/backup.sh
}
```

#### setting

- ports : http, 8080, dns=api.syui.ai
- env : PASS=xxx, TOKEN=xxx
- cmd-override : /bin/api
- volumes : /app/data

