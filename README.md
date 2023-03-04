# AcademicTeamManager - ATM
A academic team management system, which can be used for making academic discussion plan and team status visualization, written in Go and Vue3.

> [Demo is available here](https://piaodazhu.github.io/) Due to legal restrictions, the front-end part of this demo is deployed on Github Pages, and you **must set your browser to allow unsafe content** in order to experience it properly.

# Environment Requirement

| Env | Version | Dowload |
|---|---|---|
| go | >= 1.18.1 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 6.0.16 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |

# Fast Deployment
```sh
# Clone this project
git clone https://github.com/piaodazhu/AcademicTeamManager.git

# Go to the backend directory, edit the configuration yaml file, run the server
cd backend/
vim config.yaml
go mod tidy
go build -o atmserver . && ./atmserver

# Go to the frontend directory
cd frontend/
npm install .
npm run dev

# Open browser and access http:127.0.0.1:8062
```

# Tech Stack

ATM mainly use Vue3 as frontend framework and Golang as backend framework.

### Frontend

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Vue.js | Frontend framework | https://v3.cn.vuejs.org |
| Vue Router | Page router | https://router.vuejs.org |
| Axios | Async request | https://axios-http.com |
| Pinia | State management | https://pinia.vuejs.org |
| Vite | Build tool | https://vitejs.dev |
| Ant Design Vue | frontend UI | https://www.antdv.com |
| Apache ECharts | Visualization figure | https://echarts.apache.org |

### Backend

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Gin | Web framework | https://gin-gonic.com |
| Gorm | ORM framework | https://gorm.io |
| Jwt | Authentication | https://github.com/golang-jwt/jwt |
| Viper | Configuration management | https://github.com/spf13/viper |
| Redis | Cache | https://github.com/go-redis/redis |
| Gomail | Email service SDK | https://github.com/go-gomail/gomail |
