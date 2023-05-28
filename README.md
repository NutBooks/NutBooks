# Bookmarks

## Frontend

Vue.js + Vite + TailwindCSS

### Prerequisites

- node v18
  - You can use nvm for convenience
    - <https://github.com/nvm-sh/nvm>

```shell
nvm use # node v18
cd frontend
npm install
npm run format
npm run lint
npm run dev
```

## API

FastAPI

### Prerequisites

- Poetry
  - <https://python-poetry.org/docs/#installation>
  - <https://blog.gyus.me/2020/introduce-poetry/>

```shell
cd api
poetry init
poetry install
poetry run uvicorn main:app --reload

poetry add pip_package_name; poetry install # use this, not pip install
```

Go to <http://localhost:8000/docs>
