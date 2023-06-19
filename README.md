# Bookmarks

저장하고 쌓여만 가던 북마크를 자동으로 요약, 정리해주는 서비스입니다.

---

- [Run Locally](#run-locally)
    * [Frontend](#frontend)
    * [API](#api)
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
