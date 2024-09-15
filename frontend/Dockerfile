FROM node:lts-alpine
WORKDIR /app/front-splitter

RUN yarn install

COPY . .
CMD [ "yarn", "dev" ]
