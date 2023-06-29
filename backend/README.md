# NutBooks/backend

## Env Variables

| 서비스 |         변수명         |   값(예시)   |                         비고                          |
|:---:|:-------------------:|:---------:|:---------------------------------------------------:|
| DB  |     MYSQL_HOST      | 127.0.0.1 | ipv4 또는 도커 컴포즈 서비스 명<br>localhost인 경우는 127.0.0.1 사용 |
| DB  |     MYSQL_PORT      |   3306    |               외부 접속 시 DB 컨테이너로 포트 매핑                |
| DB  |     MYSQL_USER      |   user    |                                                     |
| DB  |   MYSQL_PASSWORD    |   1234    |                                                     |
| DB  | MYSQL_ROOT_PASSWORD |   5678    |                    root 계정 비밀번호                     |

## API 서버 빌드

```bash
go build -o ./bin/main main.go

```

## DB 실행

### MySQL 컨테이너 실행

```bash
docker compose up -d db; docker compose logs -f --tail=1000 db
```

### Migrate DB

- DB 설치 후 최초 1회 실행
    ```bash
    $ ./bin/main migrate
    ```
