# Bookmarks

저장하고 쌓여만 가던 북마크를 자동으로 요약, 정리해주는 서비스입니다.

---

- [Run Locally](#run-locally)
    * [Frontend](#frontend)
    * [API](#api)
        + [DB](#db)
            - [Env variables](#env-variables)
        + [API Reference](#api-reference)
            - [User](#user)
- [Release History](#release-history)
- [Roadmap](#roadmap)
- [Authors](#authors)

---

## Run Locally

Clone the project

```bash
git clone https://github.com/cheesecat47/Bookmarks.git
```

Go to the project directory

```bash
cd Bookmarks
```

### Frontend

- node v18
    - You can use nvm for convenience
        - <https://github.com/nvm-sh/nvm>
- React

```shell
nvm use # node v18
cd frontend
npm install
npm run format
npm run lint
npm run dev
```

### API

- Go 1.20
- Docker, Docker Compose

```shell
cd api
docker compose up -d
```

Go to <http://localhost:3000/docs>

#### DB

기본 DB는 MySQL을 사용하고, 도커 컨테이너로 제공됩니다.

##### Env variables

- `MYSQL_ROOT_PASSWORD`: `root` 계정 비밀번호
- `MYSQL_DATABASE`: 데이터베이스 이름
- `MYSQL_USER`, `MYSQL_PASSWORD`: `root`가 아닌 `MYSQL_DATABASE` 데이터베이스 권한을 갖는 유저명, 비밀번호
- `MYSQL_PORT`: 외부 접속 시 DB로 포트 매핑

#### API Reference

##### User

[//]: # (- Get all items)

[//]: # (  ```http)

[//]: # (  GET /api/items)

[//]: # (  ```)

[//]: # (  | Parameter |   Type   | Description                |  )

[//]: # (  |:---------:|:--------:|:---------------------------|)

[//]: # (  | `api_key` | `string` | **Required**. Your API key |)

[//]: # ()

[//]: # (- Get item )

[//]: # (  ```http)

[//]: # (  GET /api/items/${id})

[//]: # (  ```)

[//]: # (  | Parameter |   Type   | Description                       |)

[//]: # (  |:---------:|:--------:|:----------------------------------|)

[//]: # (  |   `id`    | `string` | **Required**. Id of item to fetch |)

[//]: # ()

[//]: # (- add&#40;num1, num2&#41;)

[//]: # (  Takes two numbers and returns the sum.)

## Release History

- Version 1.0.0
    - Board service

## Roadmap

[//]: # (- [ ] Additional browser support)

## Authors

- [@cheesecat47](https://www.github.com/cheesecat47)
- [@appletail](https://github.com/appletail)
- [@yeni28](https://github.com/yeni28)
