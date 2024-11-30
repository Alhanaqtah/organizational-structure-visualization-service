FROM node:18-alpine

RUN npm install -g pnpm

WORKDIR /usr/src/app

COPY package*.json pnpm-lock.yaml ./

RUN pnpm i --only=production

COPY . .

RUN pnpm run build 

CMD ["pnpm", "run", "start"]