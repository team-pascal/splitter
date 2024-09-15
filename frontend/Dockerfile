FROM node:20.17-alpine3.19
WORKDIR /app/front-splitter

COPY front-splitter/package.json front-splitter/package-lock.json ./
RUN npm install

COPY front-splitter/ ./

CMD [ "npm", "run", "dev" ]
