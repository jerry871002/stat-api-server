FROM node:20-alpine as builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .

RUN npm run build

FROM nginx:1.27.0-alpine as runner

COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80
